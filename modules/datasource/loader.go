package datasource

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	models "moneypro.kamontat.net/models-common"
	transaction "moneypro.kamontat.net/models-transaction"
	e "moneypro.kamontat.net/utils-error"
	logger "moneypro.kamontat.net/utils-logger"
	measure "moneypro.kamontat.net/utils-measure"
)

var logcode = 2000

// Loader will load csv and convert to transaction struct
func Loader(output *logger.Logger, filename string) (*models.Application, error) {
	timing := measure.NewTiming()

	stepname := "Load csv file"
	csvFile, err := os.Open(filename)
	defer csvFile.Close()
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}
	timing.LogSnapshot(stepname, output, logcode+10).Save(stepname)

	stepname = "Load csv file information"
	csvFileInfo, err := csvFile.Stat()
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	output.Info(logcode, "Reading %s %d bytes", csvFileInfo.Name(), csvFileInfo.Size()) // log information
	timing.LogSnapshot(stepname, output, logcode+11).Save(stepname)                     // log measurement timing

	stepname = "Read csv content"
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}
	timing.LogSnapshot(stepname, output, logcode+12).Save(stepname)

	stepname = "Convert content to golang object"
	var nameMapper []string
	application := models.NewApplication()

	output.Debug(logcode, "Loading csv size: %d line", len(csvLines)-1)
	for index, line := range csvLines {
		if index == 0 {
			nameMapper = line
		} else {
			mapper := make(map[string]string)

			mapper["Index"] = strconv.Itoa(index)
			for colIndex, col := range line {
				keyname := nameMapper[colIndex]
				value := strings.TrimSpace(col)

				if keyname != "" {
					mapper[keyname] = value
				}
			}

			transaction, err := transaction.Builder(mapper)
			e.When(err).Print(output, logcode).OnCompleted(func() {
				application.AddTransaction(transaction)
			})
		}
	}
	timing.LogSnapshot(stepname, output, logcode+13).Save(stepname)
	timing.LogAll("Loading csv file", output, logcode+100)

	return application, nil
}

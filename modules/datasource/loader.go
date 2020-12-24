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
)

var logcode = 2000

// Loader will load csv and convert to transaction struct
func Loader(output *logger.Logger, filename string) (*models.Application, error) {
	csvFile, err := os.Open(filename)
	defer csvFile.Close()
	if e.When(err).Exist() {
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if e.When(err).Exist() {
		return nil, err
	}

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

	return application, nil
}

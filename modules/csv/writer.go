package csv

import (
	"strings"

	connection "moneypro.kamontat.net/connection-common"
	models "moneypro.kamontat.net/models-common"
	profile "moneypro.kamontat.net/models-profile"
	logger "moneypro.kamontat.net/utils-logger"
	measure "moneypro.kamontat.net/utils-measure"
)

var logcode = 11000

// Writer is writer for csv format
type Writer struct {
	Connection connection.Interface
	Profile    *profile.Profile
}

// Header is csv header name
func (w *Writer) Header() string {
	outputs := []string{
		"Index",
		"Date",
		"Type Code",
		"Type Name",
		"Type Shortname",
		"Account Name",
		"Account Type Name",
		"Account Type Shortname",
		"Amount",
		"Auto Amount",
		"Expense",
		"Income",
		"Amount Currency Full",
		"Amount Currency",
		"Account To Name",
		"Account To Type Name",
		"Account To Type Shortname",
		"Amount To",
		"Amount To Currency Full",
		"Amount To Currency",
		"Balance",
		"Balance Currency Full",
		"Balance Currency",
		"Base Category",
		"Sub Category",
		"Full Category",
		"Description",
		"Agent",
		"Check number",
		"Class",
	}
	return strings.Join(outputs, ",")
}

// Start writing data to file via file creator
func (w *Writer) Start(output *logger.Logger) (int, error) {
	var byteSize int

	timing := measure.NewTiming()

	stepname := "      Writing header"
	size, err := w.Connection.Write(w.Header(), true)
	if err != nil {
		return byteSize, err
	}
	byteSize += size
	timing.LogSnapshot(stepname, output, logcode+10).Save(stepname)

	stepname = "      Writing transaction"
	w.Profile.ForEachTransaction(func(index int, transaction *models.Transaction) {
		csvFormat := Converter(transaction, w.Profile.GetAccount(transaction.Account), w.Profile.GetAccount(transaction.AccountTo))
		size, err := w.Connection.Write(csvFormat, true)
		if err != nil {
			byteSize += size
		}
	})
	timing.LogSnapshot(stepname, output, logcode+11).Save(stepname)

	return byteSize, nil
}

// NewWriter will return create writer object with input value
func NewWriter(file connection.Interface, profile *profile.Profile) *Writer {
	return &Writer{
		Profile:    profile,
		Connection: file,
	}
}

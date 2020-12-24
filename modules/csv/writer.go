package csv

import (
	"strings"

	models "moneypro.kamontat.net/models-common"
	e "moneypro.kamontat.net/utils-error"
	logger "moneypro.kamontat.net/utils-logger"
	writer "moneypro.kamontat.net/writer"
)

// Writer is writer for csv format
type Writer struct {
	writer.Writer
}

// Header is csv header name
func (w *Writer) Header() string {
	outputs := []string{
		"Date",
		"Type",
		"Account",
		"Amount",
		"Amount Currency",
		"Account To",
		"Amount To",
		"Amount To Currency",
		"Balance",
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

	// Write header to csv
	size, err := w.Writer.Creator.Write(w.Header(), true)
	if e.When(err).Print(output, 1).Exist() {
		return byteSize, err
	}
	byteSize += size

	w.Writer.Application.ForEachTransaction(func(index int, transaction *models.Transaction) {
		csvFormat := Converter(transaction)
		size, err := w.Writer.Creator.Write(csvFormat, true)
		if e.When(err).Print(output, 1).Empty() {
			byteSize += size
		}
	})

	return byteSize, nil
}

// NewWriter will return create writer object with input value
func NewWriter(creator *writer.FileCreator, application *models.Application) *Writer {
	return &Writer{
		writer.Writer{
			WriterType:  writer.CSV,
			Application: application,
			Creator:     creator,
		},
	}
}

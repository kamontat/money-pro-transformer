package writer

import (
	models "kamontat.net/money-pro-models"
)

// CsvWriter is writer for csv format
type CsvWriter struct {
	Writer
}

// Writing will write data to file by file creator
func (w *CsvWriter) Writing() (int, error) {
	var byteSize int

	s, e := w.Writer.creator.Write(w.Writer.account.CsvHeader())
	if e != nil {
		return byteSize, e
	}
	byteSize += s

	s, e = w.Writer.creator.WriteNewLine()
	if e != nil {
		return byteSize, e
	}
	byteSize += s

	for _, transaction := range w.Writer.account.GetTransactions() {
		s, e = w.Writer.creator.Write(transaction.CsvString())
		if e != nil {
			return byteSize, e
		}
		byteSize += s

		s, e = w.Writer.creator.WriteNewLine()
		if e != nil {
			return byteSize, e
		}
		byteSize += s
	}

	return byteSize, nil
}

// NewCsvWriter will return new writer for csv file
func NewCsvWriter(creator *FileCreator, ac *models.Accounts) *CsvWriter {
	return &CsvWriter{
		Writer{
			account:    ac,
			writerType: CSV,
			creator:    creator,
		},
	}
}

package writer

import (
	models "moneypro.kamontat.net/models-common"
	logger "moneypro.kamontat.net/utils-logger"
)

var logcode = 3000

// Writer is interface for writer data to file
type Writer struct {
	Application *models.Application
	WriterType  Type
	Creator     *FileCreator
}

// Debug will print debug info from account
func (w *Writer) Debug(output *logger.Logger) {
	output.Debug(logcode, w.Application.String())
}

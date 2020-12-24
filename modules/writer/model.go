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

// Info will print useful information
func (w *Writer) Info(output *logger.Logger) {
	output.Info(logcode, "Write to   : %s", w.Creator.FilePath)
	output.Info(logcode, "Result type: %s", w.WriterType)
}

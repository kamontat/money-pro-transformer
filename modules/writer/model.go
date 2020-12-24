package writer

import (
	profile "moneypro.kamontat.net/models-profile"
	logger "moneypro.kamontat.net/utils-logger"
)

var logcode = 3000

// Writer is interface for writer data to file
type Writer struct {
	Profile    *profile.Profile
	WriterType Type
	Creator    *FileCreator
}

// Info will print useful information
func (w *Writer) Info(output *logger.Logger) {
	output.Info(logcode, "Write to   : %s", w.Creator.FilePath)
	output.Info(logcode, "Result type: %s", w.WriterType)
}

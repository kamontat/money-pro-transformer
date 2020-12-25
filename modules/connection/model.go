package connection

import (
	"io"

	logger "moneypro.kamontat.net/utils-logger"
)

// Interface is model interface for connection
type Interface interface {
	Type() Type
	Close() error
	AutoClose()

	Write(data string, newline bool) (int, error)
	WriteNewLine() (int, error)

	GetReader() io.Reader

	Info(output *logger.Logger, code int)
}

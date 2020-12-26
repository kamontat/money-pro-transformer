package connection

import (
	"io"
)

// Interface is model interface for connection
type Interface interface {
	Type() Type
	Close() error
	AutoClose()

	Write(data string, newline bool) (int, error)
	WriteNewLine() (int, error)

	GetReader() io.Reader

	Info() string
}

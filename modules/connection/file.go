package connection

import (
	"fmt"
	"io"
	"os"
)

// File is inherite from connection.Interface
type File struct {
	Pathname string
	IO       *os.File
}

// Type of connection
func (f *File) Type() Type {
	return FILE
}

// Close will close connection
func (f *File) Close() error {
	return f.IO.Close()
}

// AutoClose is setting to close connection when not use
func (f *File) AutoClose() {
	defer f.Close()
}

// Write will write input string to file content
func (f *File) Write(msg string, newline bool) (int, error) {
	byteSize, err := f.IO.WriteString(msg)
	if err != nil {
		return 0, err
	}

	if newline {
		size, err := f.WriteNewLine()
		if err != nil {
			return 0, err
		}

		byteSize += size
	}

	return byteSize, nil
}

// WriteNewLine add new line to end of file content
func (f *File) WriteNewLine() (int, error) {
	return f.IO.WriteString("\n")
}

// GetReader will return file reader
func (f *File) GetReader() io.Reader {
	return f.IO
}

// Info will return useful information
func (f *File) Info() string {
	return fmt.Sprintf("Connection path: %s (type=%s)", f.Pathname, f.Type())
}

// NewInputFile will open new file as readonly
func NewInputFile(pathname string) (Interface, error) {
	file, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}

	return &File{
		Pathname: pathname,
		IO:       file,
	}, nil
}

// NewOutputFile will return create writable file
func NewOutputFile(pathname string) (Interface, error) {
	file, err := os.Create(pathname)
	if err != nil {
		return nil, err
	}

	return &File{
		Pathname: pathname,
		IO:       file,
	}, nil
}

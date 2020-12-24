package writer

import (
	"os"

	e "moneypro.kamontat.net/utils-error"
)

// FileCreator is a wrapper of os.File
type FileCreator struct {
	file *os.File
}

// AutoClose is setting to close connection when not use
func (f *FileCreator) AutoClose() {
	defer f.file.Close()
}

// Write will write input string to file content
func (f *FileCreator) Write(msg string, newline bool) (int, error) {
	byteSize, err := f.file.WriteString(msg)
	if e.When(err).Exist() {
		return 0, err
	}

	if newline {
		size, err := f.WriteNewLine()
		if e.When(err).Exist() {
			return 0, err
		}

		byteSize += size
	}

	return byteSize, nil
}

// WriteNewLine add new line to end of file content
func (f *FileCreator) WriteNewLine() (int, error) {
	return f.file.WriteString("\n")
}

// NewFileCreator is helper for create new FileCreator
func NewFileCreator(filepath string) (*FileCreator, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	return &FileCreator{
		file: file,
	}, nil
}

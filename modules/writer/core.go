package writer

import models "kamontat.net/money-pro-models"

// Writer is interface for writer data to file
type Writer struct {
	account    *models.Accounts
	writerType Type
	creator    *FileCreator
}

// Debug will print debug info from account
func (w *Writer) Debug() {
	w.account.Debug()
}

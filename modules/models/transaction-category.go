package models

import "strings"

// TransactionCategory is category modified string
type TransactionCategory struct {
	Name     string
	Sub      string
	FullName string
}

// NewCategory will split string to category and subcategory
func NewCategory(str string) *TransactionCategory {
	arr := strings.Split(str, ": ")

	baseCategory := ""
	if len(arr) > 0 {
		baseCategory = strings.TrimSpace(arr[0])
	}

	subcategory := ""
	if len(arr) > 1 {
		subcategory = strings.TrimSpace(arr[1])
	}

	return &TransactionCategory{
		Name:     baseCategory,
		Sub:      subcategory,
		FullName: strings.TrimSpace(str),
	}
}

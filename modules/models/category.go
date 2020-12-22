package models

import "strings"

// Category is the income and expense category
type Category struct {
	base string
	sub  string
	full string
}

// NewCategory will split string to category and subcategory
func NewCategory(str string) *Category {
	arr := strings.Split(str, ": ")

	baseCategory := ""
	if len(arr) > 0 {
		baseCategory = arr[0]
	}

	subcategory := ""
	if len(arr) > 1 {
		subcategory = arr[1]
	}

	return &Category{
		base: baseCategory,
		sub:  subcategory,
		full: strings.TrimSpace(str),
	}
}

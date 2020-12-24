package models

// Account is account option
type Account struct {
	Name         string
	Type         AccountType
	Transactions []*Transaction
}

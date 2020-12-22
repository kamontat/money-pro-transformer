package models

import (
	"fmt"
	"strings"
)

// Accounts is map of transaction per account
type Accounts struct {
	account      map[string][]*Transaction
	transactions []*Transaction
}

// CsvHeader is string of csv content from transaction CsvString()
func (a *Accounts) CsvHeader() string {
	outputs := []string{
		"Date",
		"Type",
		"Account",
		"Amount",
		"Amount Currency",
		"Account To",
		"Amount To",
		"Amount To Currency",
		"Balance",
		"Balance Currency",
		"Base Category",
		"Sub Category",
		"Full Category",
		"Description",
		"Agent",
		"Check number",
		"Class",
	}
	return strings.Join(outputs, ",")
}

// AddTransaction will add new transaction to spacify account
func (a *Accounts) AddTransaction(t *Transaction) {
	accountName := t.GetAccountName()
	a.account[accountName] = append(a.account[accountName], t)
	a.transactions = append(a.transactions, t)
}

// GetTransactions will return all transactions
func (a *Accounts) GetTransactions() []*Transaction {
	return a.transactions
}

// Debug will print debug information to console
func (a *Accounts) Debug() {
	for key, transactions := range a.account {
		fmt.Println("Account: " + key)
		for i, transaction := range transactions {
			fmt.Printf("  %d: %s\n", i+1, transaction.String())
		}
	}
}

// NewAccounts return empty account struct
func NewAccounts() *Accounts {
	return &Accounts{
		account:      make(map[string][]*Transaction),
		transactions: make([]*Transaction, 0),
	}
}

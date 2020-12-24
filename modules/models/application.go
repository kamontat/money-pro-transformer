package models

import (
	"fmt"
	"strings"
)

// Application is all processed information from csv file
type Application struct {
	transactions []*Transaction
	accounts     map[string]*Account
}

// AddTransaction will add transaction to application
func (a *Application) AddTransaction(t *Transaction) *Application {
	a.transactions = append(a.transactions, t)
	if a.accounts[t.Account] == nil {
		accountType := SA

		if t.Type == TOB && t.Amount == 0 && t.Balance != 0 {
			accountType = CC
		} else if t.Type == TOB && t.BalanceCurrency == USD {
			accountType = US
		} else if strings.Contains(strings.ToLower(t.Account), "fund") {
			accountType = FD
		} else if strings.Contains(strings.ToLower(t.Account), "stock") {
			accountType = SK
		} else if strings.Contains(strings.ToLower(t.Account), "true") ||
			strings.Contains(strings.ToLower(t.Account), "bluepay") ||
			strings.Contains(strings.ToLower(t.Account), "linepay") {
			accountType = OW
		}

		a.accounts[t.Account] = &Account{
			Type: accountType,
		}
	}

	return a
}

// ForEachTransaction will run for-loop on each transaction
func (a *Application) ForEachTransaction(fn func(int, *Transaction)) {
	for index, transaction := range a.transactions {
		fn(index, transaction)
	}
}

// MapEachTransaction will run for-loop on each transaction
func (a *Application) MapEachTransaction(fn func(int, *Transaction) interface{}) []interface{} {
	var result = make([]interface{}, 0)
	for index, transaction := range a.transactions {
		result = append(result, fn(index, transaction))
	}
	return result
}

// Length is size of transaction
func (a *Application) Length() int {
	return len(a.transactions)
}

// GetAccount will return account option
func (a *Application) GetAccount(key string) *Account {
	return a.accounts[key]
}

func (a *Application) String() string {
	var str strings.Builder

	for i, transaction := range a.transactions {
		str.WriteString(fmt.Sprintf("%3d: %s\n", i+1, transaction.String()))
	}

	return str.String()
}

// NewApplication will return empty application
func NewApplication() *Application {
	return &Application{
		transactions: make([]*Transaction, 0),
		accounts:     make(map[string]*Account),
	}
}

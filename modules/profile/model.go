package profile

import (
	"strings"

	models "moneypro.kamontat.net/models-common"
	logger "moneypro.kamontat.net/utils-logger"
)

// Profile is top level of data in Money Pro
type Profile struct {
	account      map[string]*models.Account
	transactions []*models.Transaction
}

// AddTransaction will add new transaction to this profile
func (p *Profile) AddTransaction(t *models.Transaction) {
	p.transactions = append(p.transactions, t)
	if p.account[t.Account] == nil {
		accountType := models.AUN
		if t.Type == models.TOB && t.Amount == 0 && t.Balance != 0 {
			accountType = models.ACC
		} else if t.Type == models.TOB && t.BalanceCurrency == models.USD {
			accountType = models.AUS
		} else if strings.Contains(strings.ToLower(t.Account), "fund") {
			accountType = models.AFD
		} else if strings.Contains(strings.ToLower(t.Account), "stock") {
			accountType = models.ASK
		} else if strings.Contains(strings.ToLower(t.Account), "true") ||
			strings.Contains(strings.ToLower(t.Account), "bluepay") ||
			strings.Contains(strings.ToLower(t.Account), "linepay") {
			accountType = models.AOW
		} else if t.Type == models.TOB && t.Amount != 0 && t.Balance != 0 {
			accountType = models.ASA
		}

		p.account[t.Account] = &models.Account{
			Name:         t.Account,
			Type:         accountType,
			Transactions: []*models.Transaction{t},
		}
	} else {
		p.account[t.Account].Transactions = append(p.account[t.Account].Transactions, t)
	}
}

// ForEachTransaction will run for-loop on each transaction
func (p *Profile) ForEachTransaction(fn func(int, *models.Transaction)) {
	for index, transaction := range p.transactions {
		fn(index, transaction)
	}
}

// MapEachTransaction will run for-loop on each transaction
func (p *Profile) MapEachTransaction(fn func(int, *models.Transaction) interface{}) []interface{} {
	var result = make([]interface{}, 0)
	for index, transaction := range p.transactions {
		result = append(result, fn(index, transaction))
	}
	return result
}

// GetAccount will return input account name
func (p *Profile) GetAccount(name string) *models.Account {
	var account = p.account[name]
	if account == nil {
		return &models.Account{
			Name:         "",
			Type:         models.AUN,
			Transactions: make([]*models.Transaction, 0),
		}
	}
	return account
}

// AccountSize will return total size of account
func (p *Profile) AccountSize() int {
	return len(p.account)
}

// TransactionSize will return total size of transaction
func (p *Profile) TransactionSize() int {
	return len(p.transactions)
}

// Info will print useful information
func (p *Profile) Info(output *logger.Logger, code int) {
	output.Info(code, "This Profile has %d account in %d transaction", p.AccountSize(), p.TransactionSize())
}

// Debug will print debug information
func (p *Profile) Debug(output *logger.Logger, code int) {
	output.Debug(code, "A profile statistic: ")
	for _, account := range p.account {
		output.Debug(code, "  Account %s (type=%s, size=%d)", account.Name, account.Type.Name, len(account.Transactions))
		for index, transaction := range account.Transactions {
			output.Debug(code, "   %3d) %s", index+1, transaction.String())
		}
	}
}

// NewProfile create empty profile object
func NewProfile() *Profile {
	return &Profile{
		account: make(map[string]*models.Account),
	}
}

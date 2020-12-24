package models

import "fmt"

// Transaction is modified data loading from csv
type Transaction struct {
	// Index is transaction index
	Index uint32
	// Datetime when transaction occurred
	Datetime string
	// Type of transaction
	Type TransactionType
	// Category of transaction
	Category *TransactionCategory
	// Amount is number of money receive/paid by this transaction
	Amount float64
	// AmountCurrency is currency of amount money
	AmountCurrency CurrencyType
	// AmountTo valid only when transaction type if transfer, amount to new account
	AmountTo float64
	// AmountToCurrency is currency of amount to money
	AmountToCurrency CurrencyType
	// Balance is current money holded in account after apply amount
	Balance float64
	// BalanceCurrency is currency of balance
	BalanceCurrency CurrencyType
	// Account is account name
	Account string
	// AccountTo is account that money transfer to
	AccountTo string
	// Description is optional description
	Description string
	// Agent is person this transaction paid for/receive from
	Agent string
	// Check is check number
	Check string
	// Class is classify of this transaction
	Class string
	// Raw is raw data from csv file
	Raw map[string]string
}

// String will return formatted string
func (t *Transaction) String() string {
	switch t.Type {
	case TIN, TEP:
		return fmt.Sprintf("%22s %s %.3f %s [%s]", t.Datetime, t.Category.FullName, t.Amount, t.AmountCurrency, t.Type.Name)
	case TOB:
		return fmt.Sprintf("%22s Opening new Account", t.Datetime)
	case TBJ:
		return fmt.Sprintf("%22s Update balance by %.3f %s", t.Datetime, t.Amount, t.AmountCurrency)
	case TMT:
		return fmt.Sprintf("%22s Move %.3f %s to '%s' (%.3f %s)", t.Datetime, t.Amount, t.AmountCurrency, t.AccountTo, t.AmountTo, t.AmountToCurrency)
	default:
		return fmt.Sprintf("%22s [%s]", t.Datetime, t.Type.Name)
	}
}

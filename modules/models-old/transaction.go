package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Transaction is object represent data on each line in csv
type Transaction struct {
	datetime         string
	ttype            TransactionType
	amount           float64
	amountCurrency   CurrencyUnit
	amountTo         float64
	amountToCurrency CurrencyUnit
	balance          float64
	balanceCurrency  CurrencyUnit
	account          string
	accountTo        string
	category         *Category
	description      string
	agent            string
	check            string
	class            string
	raw              map[string]string
}

// GetAccountName will return current transaction account from/use
func (t *Transaction) GetAccountName() string {
	return t.account
}

// String will return formatted string
func (t *Transaction) String() string {
	switch t.ttype {
	case IN, EP:
		return fmt.Sprintf("%22s %s %.3f %s [%s]", t.datetime, t.category, t.amount, t.amountCurrency, t.ttype)
	case OB:
		return fmt.Sprintf("%22s Opening new Account", t.datetime)
	case BJ:
		return fmt.Sprintf("%22s Update balance by %.3f %s", t.datetime, t.amount, t.amountCurrency)
	case MT:
		return fmt.Sprintf("%22s Move %.3f %s to '%s' (%.3f %s)", t.datetime, t.amount, t.amountCurrency, t.accountTo, t.amountTo, t.amountToCurrency)
	default:
		return fmt.Sprintf("%22s [%s]", t.datetime, t.ttype)
	}
}

// CsvString will return data as csv format
func (t *Transaction) CsvString() string {
	output := []string{
		t.datetime,
		t.ttype.String(),
		t.account,
		strconv.FormatFloat(t.amount, 'f', -1, 64),
		t.amountCurrency.String(),
		t.accountTo,
		strconv.FormatFloat(t.amountTo, 'f', -1, 64),
		t.amountToCurrency.String(),
		strconv.FormatFloat(t.balance, 'f', -1, 64),
		t.balanceCurrency.String(),
		t.category.base,
		t.category.sub,
		t.category.full,
		t.description,
		t.agent,
		t.check,
		t.class,
	}
	return strings.Join(output, ",")
}

// NewTransaction will create Transaction struct base on input map
func NewTransaction(mapper map[string]string) (*Transaction, error) {
	amount, amountCurrency, err := ToCurrency(AMOUNT.Get(mapper))
	if err != nil {
		return nil, err
	}

	amountTo, amountToCurrency, err := ToCurrency(AMOUNTTO.Get(mapper))
	if err != nil {
		return nil, err
	}

	balance, balanceCurrency, err := ToCurrency(BALANCE.Get(mapper))
	if err != nil {
		return nil, err
	}

	return &Transaction{
		datetime:         DATE.Get(mapper),
		ttype:            ToTransactionType(TRANSACTIONTYPE.Get(mapper)),
		account:          ACCOUNT.Get(mapper),
		accountTo:        ACCOUNTTO.Get(mapper),
		category:         NewCategory(CATEGORY.Get(mapper)),
		amount:           amount,
		amountCurrency:   amountCurrency,
		amountTo:         amountTo,
		amountToCurrency: amountToCurrency,
		balance:          balance,
		balanceCurrency:  balanceCurrency,
		description:      DESCRIPTION.Get(mapper),
		agent:            AGENT.Get(mapper),
		check:            CHECK.Get(mapper),
		class:            CLASS.Get(mapper),
		raw:              mapper,
	}, nil
}

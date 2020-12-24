package transaction

import (
	"strconv"

	models "moneypro.kamontat.net/models-common"
	currency "moneypro.kamontat.net/models-currency"
	csv "moneypro.kamontat.net/utils-csv"
	e "moneypro.kamontat.net/utils-error"
	logger "moneypro.kamontat.net/utils-logger"
)

var logcode = 5000

// Builder will build transaction object
func Builder(mapper map[string]string) (*models.Transaction, error) {
	output := logger.Get()

	amount, amountCurrency, err := currency.Builder(csv.AMOUNT.Get(mapper))
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	amountTo, amountToCurrency, err := currency.Builder(csv.AMOUNTTO.Get(mapper))
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	balance, balanceCurrency, err := currency.Builder(csv.BALANCE.Get(mapper))
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	// output.Debug(logcode, "Index number: %s", csv.INDEX.Get(mapper))
	index, err := strconv.ParseUint(csv.INDEX.Get(mapper), 10, 32)
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	return &models.Transaction{
		Index:            uint32(index),
		Datetime:         csv.DATE.Get(mapper),
		Type:             models.ToTransactionType(csv.TRANSACTIONTYPE.Get(mapper)),
		Category:         models.NewCategory(csv.CATEGORY.Get(mapper)),
		Amount:           amount,
		AmountCurrency:   amountCurrency,
		AmountTo:         amountTo,
		AmountToCurrency: amountToCurrency,
		Balance:          balance,
		BalanceCurrency:  balanceCurrency,
		Account:          csv.ACCOUNT.Get(mapper),
		AccountTo:        csv.ACCOUNTTO.Get(mapper),
		Description:      csv.DESCRIPTION.Get(mapper),
		Agent:            csv.AGENT.Get(mapper),
		Check:            csv.CHECK.Get(mapper),
		Class:            csv.CLASS.Get(mapper),
		Raw:              mapper,
	}, nil
}

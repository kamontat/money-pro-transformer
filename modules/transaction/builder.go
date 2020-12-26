package transaction

import (
	"strconv"

	models "moneypro.kamontat.net/models-common"
	currency "moneypro.kamontat.net/models-currency"
	e "moneypro.kamontat.net/utils-error"
	logger "moneypro.kamontat.net/utils-logger"
)

var logcode = 5000

// Builder will build transaction object
func Builder(mapper map[string]string) (*models.Transaction, error) {
	output := logger.Get()

	amount, amountCurrency, err := currency.Builder(models.AMOUNT.Get(mapper))
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	amountTo, amountToCurrency, err := currency.Builder(models.AMOUNTTO.Get(mapper))
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	balance, balanceCurrency, err := currency.Builder(models.BALANCE.Get(mapper))
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	// output.Debug(logcode, "Index number: %s", models.INDEX.Get(mapper))
	index, err := strconv.ParseUint(models.INDEX.Get(mapper), 10, 32)
	if e.When(err).Print(output, logcode).Exist() {
		return nil, err
	}

	return &models.Transaction{
		Index:            uint32(index),
		Datetime:         models.DATE.Get(mapper),
		Type:             models.ToTransactionType(models.TRANSACTIONTYPE.Get(mapper)),
		Category:         models.NewCategory(models.CATEGORY.Get(mapper)),
		Amount:           amount,
		AmountCurrency:   amountCurrency,
		AmountTo:         amountTo,
		AmountToCurrency: amountToCurrency,
		Balance:          balance,
		BalanceCurrency:  balanceCurrency,
		Account:          models.ACCOUNT.Get(mapper),
		AccountTo:        models.ACCOUNTTO.Get(mapper),
		Description:      models.DESCRIPTION.Get(mapper),
		Agent:            models.AGENT.Get(mapper),
		Check:            models.CHECK.Get(mapper),
		Class:            models.CLASS.Get(mapper),
		Raw:              mapper,
	}, nil
}

package csv

import (
	"strings"

	models "moneypro.kamontat.net/models-common"
	utils "moneypro.kamontat.net/utils-common"
)

// Converter will convert transaction to csv format
func Converter(t *models.Transaction) string {
	output := []string{
		t.Datetime,
		t.Type.Name,
		t.Account,
		utils.FloatToString(t.Amount),
		t.AmountCurrency.Short,
		t.AccountTo,
		utils.FloatToString(t.AmountTo),
		t.AmountToCurrency.Short,
		utils.FloatToString(t.Balance),
		t.BalanceCurrency.Short,
		t.Category.Name,
		t.Category.Sub,
		t.Category.FullName,
		t.Description,
		t.Agent,
		t.Check,
		t.Class,
	}

	return strings.Join(output, ",")
}

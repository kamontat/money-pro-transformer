package csv

import (
	"strings"

	models "moneypro.kamontat.net/models-common"
	utils "moneypro.kamontat.net/utils-common"
)

// Converter will convert transaction to csv format
func Converter(t *models.Transaction) string {
	output := []string{
		utils.UIntToString(t.Index),
		t.Datetime,
		utils.UIntToString(uint32(t.Type.Code)),
		t.Type.Name,
		t.Type.Short,
		t.Account,
		utils.FloatToString(t.Amount),
		t.AmountCurrency.Name,
		t.AmountCurrency.Short,
		t.AccountTo,
		utils.FloatToString(t.AmountTo),
		t.AmountToCurrency.Name,
		t.AmountToCurrency.Short,
		utils.FloatToString(t.Balance),
		t.BalanceCurrency.Name,
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

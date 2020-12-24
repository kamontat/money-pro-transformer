package csv

import (
	"strings"

	models "moneypro.kamontat.net/models-common"
	utils "moneypro.kamontat.net/utils-common"
)

// Converter will convert transaction to csv format
func Converter(t *models.Transaction, account *models.Account, accountTo *models.Account) string {
	output := []string{
		utils.UIntToString(t.Index),
		t.Datetime,
		utils.UIntToString(uint32(t.Type.Code)),
		t.Type.Name,
		t.Type.Short,
		account.Name,
		account.Type.Name,
		account.Type.Short,
		utils.FloatToString(t.Amount),
		utils.FloatToString(t.AutoAmount()),
		utils.FloatToString(t.Expense()),
		utils.FloatToString(t.Income()),
		t.AmountCurrency.Name,
		t.AmountCurrency.Short,
		accountTo.Name,
		accountTo.Type.Name,
		accountTo.Type.Short,
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

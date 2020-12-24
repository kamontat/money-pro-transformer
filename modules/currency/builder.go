package currency

import (
	"regexp"
	"strconv"
	"strings"

	models "moneypro.kamontat.net/models-common"
)

// Builder will return amount and currency of input string
func Builder(str string) (float64, models.CurrencyType, error) {
	negativeRegex := regexp.MustCompile(`^\(.*\)$`)
	numberRegex := regexp.MustCompile(`\d[\d,]*[\.]?[\d]*`)

	var err error
	amount := float64(0)
	currency := models.UKN

	if numberRegex.MatchString(str) {
		amountString := numberRegex.FindAllString(str, -1)[0]
		amount, err = strconv.ParseFloat(strings.ReplaceAll(amountString, ",", ""), 64)
		if err != nil {
			return amount, currency, err
		}

		if negativeRegex.MatchString(str) {
			amount = amount * -1
		}
	}

	if strings.Contains(str, models.USD.Symbol) {
		currency = models.USD
	} else if strings.Contains(str, models.THB.Symbol) {
		currency = models.THB
	} else if strings.Contains(str, models.HKD.Symbol) {
		currency = models.HKD
	} else if strings.Contains(str, models.GBP.Symbol) {
		currency = models.GBP
	}

	return amount, currency, nil
}

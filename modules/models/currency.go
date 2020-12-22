package models

import (
	"regexp"
	"strconv"
	"strings"
)

// CurrencyUnit is string enum of supported currency
type CurrencyUnit string

// String will return string represent current enum
func (c CurrencyUnit) String() string {
	return string(c)
}

// GetUnit return unit symbol using by money Pro application
func (c CurrencyUnit) GetUnit() string {
	switch c {
	case USD:
		return "US$"
	case THB:
		return "฿"
	case HKD:
		return "HK$"
	case GBP:
		return "£"
	}

	return ""
}

// ToCurrency will convert string to currency unit
func ToCurrency(str string) (float64, CurrencyUnit, error) {
	negativeRegex := regexp.MustCompile(`^\(.*\)$`)
	numberRegex := regexp.MustCompile(`\d[\d,]*[\.]?[\d]*`)

	var err error
	amount := float64(0)
	unit := NON

	if numberRegex.MatchString(str) {
		amountString := numberRegex.FindAllString(str, -1)[0]
		amount, err = strconv.ParseFloat(strings.ReplaceAll(amountString, ",", ""), 64)
		if err != nil {
			return amount, unit, err
		}

		if negativeRegex.MatchString(str) {
			amount = amount * -1
		}
	}

	if strings.Contains(str, USD.GetUnit()) {
		unit = USD
	} else if strings.Contains(str, THB.GetUnit()) {
		unit = THB
	} else if strings.Contains(str, HKD.GetUnit()) {
		unit = HKD
	} else if strings.Contains(str, GBP.GetUnit()) {
		unit = GBP
	}

	return amount, unit, nil
}

const (
	// USD is US Dollar
	USD CurrencyUnit = "USD"
	// THB is Thai Baht
	THB CurrencyUnit = "THB"
	// HKD is Hongkong Dollar
	HKD CurrencyUnit = "HKD"
	// GBP is England pound
	GBP CurrencyUnit = "GBP"

	// NON is unknown currency unit
	NON CurrencyUnit = "NON"
)

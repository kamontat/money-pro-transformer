package models

// CurrencyType is currency supported by Money Pro
type CurrencyType struct {
	Code  uint8
	Name  string
	Short string

	Symbol string
}

// String will return string represent current enum
func (c CurrencyType) String() string {
	return c.Name
}

var (
	// USD is US Dollar
	USD CurrencyType = CurrencyType{
		Code:   uint8(0),
		Name:   "US Dollar",
		Short:  "USD",
		Symbol: "US$",
	}

	// THB is Thai Baht
	THB CurrencyType = CurrencyType{
		Code:   uint8(1),
		Name:   "Thai Baht",
		Short:  "THB",
		Symbol: "฿",
	}

	// HKD is Hongkong Dollar
	HKD CurrencyType = CurrencyType{
		Code:   uint8(2),
		Name:   "Hong Kong Dollar",
		Short:  "HKD",
		Symbol: "HK$",
	}

	// GBP is England pound
	GBP CurrencyType = CurrencyType{
		Code:   uint8(3),
		Name:   "Pound sterling",
		Short:  "GBP",
		Symbol: "£",
	}

	// UKN is unknown currency unit
	UKN CurrencyType = CurrencyType{
		Code:   uint8(99),
		Name:   "Unknown",
		Short:  "UKN",
		Symbol: "",
	}
)

// CurrencyTypes is slice of supported currency
var CurrencyTypes = []CurrencyType{
	USD,
	THB,
	HKD,
	GBP,
}

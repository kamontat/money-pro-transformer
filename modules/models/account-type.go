package models

// AccountType is account type
type AccountType struct {
	Code  uint8
	Name  string
	Short string
}

var (
	// SA is Saving Account
	SA AccountType = AccountType{
		Code:  uint8(1),
		Name:  "Saving Account",
		Short: "SA",
	}

	// CC is Credit Card
	CC AccountType = AccountType{
		Code:  uint8(2),
		Name:  "Credit Card",
		Short: "CC",
	}

	// FD is Fund, one of possible asset in Money Pro
	FD AccountType = AccountType{
		Code:  uint8(3),
		Name:  "Fund",
		Short: "FD",
	}

	// SK is Stock, one of possible asset in Money Pro
	SK AccountType = AccountType{
		Code:  uint8(4),
		Name:  "Stock",
		Short: "SK",
	}

	// OW is Online wallet, another extra wallet
	OW AccountType = AccountType{
		Code:  uint8(5),
		Name:  "Online wallet",
		Short: "OW",
	}

	// US is international account (US)
	US AccountType = AccountType{
		Code:  uint8(6),
		Name:  "US Account",
		Short: "US",
	}
)

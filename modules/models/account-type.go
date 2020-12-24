package models

// AccountType is account type
type AccountType struct {
	Code  uint8
	Name  string
	Short string
}

var (
	// AUN is Unknown account type
	AUN AccountType = AccountType{
		Code:  uint8(0),
		Name:  "Unknown",
		Short: "AUN",
	}

	// ASA is Saving Account
	ASA AccountType = AccountType{
		Code:  uint8(1),
		Name:  "Saving Account",
		Short: "ASA",
	}

	// ACC is Credit Card
	ACC AccountType = AccountType{
		Code:  uint8(2),
		Name:  "Credit Card",
		Short: "ACC",
	}

	// AFD is Fund, one of possible asset in Money Pro
	AFD AccountType = AccountType{
		Code:  uint8(3),
		Name:  "Fund",
		Short: "AFD",
	}

	// ASK is Stock, one of possible asset in Money Pro
	ASK AccountType = AccountType{
		Code:  uint8(4),
		Name:  "Stock",
		Short: "ASK",
	}

	// AOW is Online wallet, another extra wallet
	AOW AccountType = AccountType{
		Code:  uint8(5),
		Name:  "Online wallet",
		Short: "AOW",
	}

	// AUS is international account (US)
	AUS AccountType = AccountType{
		Code:  uint8(6),
		Name:  "US Account",
		Short: "AUS",
	}
)

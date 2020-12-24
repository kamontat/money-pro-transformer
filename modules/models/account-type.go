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
)

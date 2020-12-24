package models

// AccountType is type of account
type AccountType string

func (t AccountType) String() string {
	return string(t)
}

const (
	// SV is Saving Account
	SV AccountType = "Saving Account"

	// CC is Credit card
	CC AccountType = "Credit card"

	// UNA is unknown account
	UNA AccountType = "Unknown"
)

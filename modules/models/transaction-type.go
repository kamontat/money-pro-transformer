package models

// TransactionType is type of transaction generate from money pro
type TransactionType string

func (t TransactionType) String() string {
	return string(t)
}

// ToTransactionType will convert string to TransactionType enum
func ToTransactionType(str string) TransactionType {
	switch str {
	case OB.String():
		return OB
	case EP.String():
		return EP
	case IN.String():
		return IN
	case BJ.String():
		return BJ
	case MT.String():
		return MT
	}
	return UN
}

const (
	// OB is Opening Balance, on each account should have only 1 OB transaction
	OB TransactionType = "Opening Balance"

	// EP is Expense
	EP TransactionType = "Expense"

	// IN is Income
	IN TransactionType = "Income"

	// BJ is Balance Adjustment, when something missing and current balance is not matches with actual money
	BJ TransactionType = "Balance Adjustment"

	// MT is Money Transfer, when we transfer money from one account to other
	MT TransactionType = "Money Transfer"

	// UN is Unknown, when data is not support from above list. It will fallback to this type
	UN TransactionType = "Unknown"
)

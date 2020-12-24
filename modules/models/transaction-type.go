package models

// TransactionType is type of transaction generate from money pro
type TransactionType struct {
	Code  uint8
	Name  string
	Short string
}

func (t TransactionType) String() string {
	return t.Name
}

// ToTransactionType will convert string to TransactionType enum
func ToTransactionType(str string) TransactionType {
	for _, t := range TransactionTypes {
		if str == t.String() {
			return t
		}
	}
	return TUN
}

var (
	// TOB is Opening Balance, on each account should have only 1 OB transaction
	TOB TransactionType = TransactionType{
		Code:  uint8(1),
		Name:  "Opening Balance",
		Short: "TOB",
	}

	// TEP is Expense
	TEP TransactionType = TransactionType{
		Code:  uint8(2),
		Name:  "Expense",
		Short: "TEP",
	}

	// TIN is Income
	TIN TransactionType = TransactionType{
		Code:  uint8(3),
		Name:  "Income",
		Short: "TIN",
	}

	// TBJ is Balance Adjustment, when something missing and current balance is not matches with actual money
	TBJ TransactionType = TransactionType{
		Code:  uint8(4),
		Name:  "Balance Adjustment",
		Short: "TBJ",
	}

	// TMT is Money Transfer, when we transfer money from one account to other
	TMT TransactionType = TransactionType{
		Code:  uint8(5),
		Name:  "Money Transfer",
		Short: "TMT",
	}

	// TUN is Unknown, when data is not support from above list. It will fallback to this type
	TUN TransactionType = TransactionType{
		Code:  uint8(99),
		Name:  "Unknown",
		Short: "TUN",
	}
)

// TransactionTypes is slice of all supported transaction type
var TransactionTypes = []TransactionType{
	TOB,
	TEP,
	TIN,
	TBJ,
	TMT,
}

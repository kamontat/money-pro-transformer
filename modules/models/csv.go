package models

// CsvKey is csv header generate from Money Pro
type CsvKey string

func (key CsvKey) String() string {
	return string(key)
}

// Get will return data from input map
func (key CsvKey) Get(mapper map[string]string) string {
	return mapper[key.String()]
}

const (
	// DATE is transaction date
	DATE CsvKey = "Date"
	// AMOUNT is account transaction amount
	AMOUNT CsvKey = "Amount"
	// ACCOUNT is current account name
	ACCOUNT CsvKey = "Account"
	// AMOUNTTO is the received amount from account (to)
	AMOUNTTO CsvKey = "Amount received"
	// ACCOUNTTO is a account that amount go to
	ACCOUNTTO CsvKey = "Account (to)"

	// BALANCE is total number after calculate amount
	BALANCE CsvKey = "Balance"

	// CATEGORY is transaction category in formatted '<category>: <subcategory>'
	CATEGORY CsvKey = "Category"

	// TRANSACTIONTYPE is type of transaction
	TRANSACTIONTYPE CsvKey = "Transaction Type"

	// DESCRIPTION is transaction description
	DESCRIPTION CsvKey = "Description"
	// AGENT is agent who spend to of receive from
	AGENT CsvKey = "Agent"
	// CHECK is check number
	CHECK CsvKey = "Check #"
	// CLASS is transaction class
	CLASS CsvKey = "Class"
)

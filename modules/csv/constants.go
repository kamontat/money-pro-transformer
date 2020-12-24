package csv

// Key is csv header generate from Money Pro
type Key string

func (key Key) String() string {
	return string(key)
}

// Get will return data from input map
func (key Key) Get(mapper map[string]string) string {
	return mapper[key.String()]
}

const (
	// INDEX is special key added after each row of csv line
	INDEX Key = "Index"
	// DATE is transaction date
	DATE Key = "Date"
	// AMOUNT is account transaction amount
	AMOUNT Key = "Amount"
	// ACCOUNT is current account name
	ACCOUNT Key = "Account"
	// AMOUNTTO is the received amount from account (to)
	AMOUNTTO Key = "Amount received"
	// ACCOUNTTO is a account that amount go to
	ACCOUNTTO Key = "Account (to)"

	// BALANCE is total number after calculate amount
	BALANCE Key = "Balance"

	// CATEGORY is transaction category in formatted '<category>: <subcategory>'
	CATEGORY Key = "Category"

	// TRANSACTIONTYPE is type of transaction
	TRANSACTIONTYPE Key = "Transaction Type"

	// DESCRIPTION is transaction description
	DESCRIPTION Key = "Description"
	// AGENT is agent who spend to of receive from
	AGENT Key = "Agent"
	// CHECK is check number
	CHECK Key = "Check #"
	// CLASS is transaction class
	CLASS Key = "Class"
)

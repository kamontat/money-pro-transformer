package models

// TransactionKey is string TransactionKey of csv format
type TransactionKey string

func (key TransactionKey) String() string {
	return string(key)
}

// Get will return data from input map
func (key TransactionKey) Get(mapper map[string]string) string {
	return mapper[key.String()]
}

const (
	// INDEX is special TransactionKey added after each row of csv line
	INDEX TransactionKey = "Index"
	// DATE is transaction date
	DATE TransactionKey = "Date"
	// AMOUNT is account transaction amount
	AMOUNT TransactionKey = "Amount"
	// ACCOUNT is current account name
	ACCOUNT TransactionKey = "Account"
	// AMOUNTTO is the received amount from account (to)
	AMOUNTTO TransactionKey = "Amount received"
	// ACCOUNTTO is a account that amount go to
	ACCOUNTTO TransactionKey = "Account (to)"

	// BALANCE is total number after calculate amount
	BALANCE TransactionKey = "Balance"

	// CATEGORY is transaction category in formatted '<category>: <subcategory>'
	CATEGORY TransactionKey = "Category"

	// TRANSACTIONTYPE is type of transaction
	TRANSACTIONTYPE TransactionKey = "Transaction Type"

	// DESCRIPTION is transaction description
	DESCRIPTION TransactionKey = "Description"
	// AGENT is agent who spend to of receive from
	AGENT TransactionKey = "Agent"
	// CHECK is check number
	CHECK TransactionKey = "Check #"
	// CLASS is transaction class
	CLASS TransactionKey = "Class"
)

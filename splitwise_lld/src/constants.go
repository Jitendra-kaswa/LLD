package src

type SplitStatus string

const (
	PAID   SplitStatus = "PAID"
	UNPAID SplitStatus = "UNPAID"
)

type ExpenseStatus string

const (
	CREATED           ExpenseStatus = "CREATED"
	PARTIALLY_SETTLED ExpenseStatus = "PARTIALLY_SETTLED"
	SETTLED           ExpenseStatus = "SETTLED"
)

type Currency string

const (
	INR Currency = "INR"
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

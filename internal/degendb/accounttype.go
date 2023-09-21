package degendb

type AccountType string

const (
	ExternallyOwnedAccount AccountType = "EOA"
	Contract               AccountType = "Contract"
)

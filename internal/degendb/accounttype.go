package degendb

type AccountType string

const (
	ExternallyOwnedAccount AccountType = "EOA"
	ContractAccount        AccountType = "Contract"
)

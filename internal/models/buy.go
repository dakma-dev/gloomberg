package models

import "github.com/ethereum/go-ethereum/common"

type BuyRules struct {
	Rules map[common.Address]*BuyRule
}

type BuyRule struct {
	ID              int
	ContractAddress common.Address
	Threshold       float64
	PrivateKey      string
	// ThresholdEth    *big.Int
}

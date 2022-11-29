package models

import "github.com/ethereum/go-ethereum/common"

type BuyRules struct {
	Rules map[common.Address]*BuyRule
}

type BuyRule struct {
	ID              int
	Name            string         `mapstructure:"name"`
	ContractAddress common.Address `mapstructure:"contract_address"`
	PrivateKey      string         `mapstructure:"private_key"`
	Threshold       float64        `mapstructure:"threshold"`
	MaxPrice        float64        `mapstructure:"max_price"`
	MinSales        uint64         `mapstructure:"min_sales"`
	MinListings     uint64         `mapstructure:"min_listings"`
}

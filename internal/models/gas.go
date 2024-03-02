package models

import "math/big"

type GasInfo struct {
	LastBlock         uint64   `json:"lastBlock"`
	LastBlockGasLimit uint64   `json:"lastBlockGasLimit"`
	GasPrice          *big.Int `json:"gasPrice"`
	GasTip            *big.Int `json:"gasTip"`
}

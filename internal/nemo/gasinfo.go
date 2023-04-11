package nemo

import "math/big"

type GasInfo struct {
	LastBlock         uint64   `json:"lastBlock"`
	LastBlockGasLimit uint64   `json:"lastBlockGasLimit"`
	GasPriceWei       *big.Int `json:"gasPrice"`
	GasTipWei         *big.Int `json:"gasTip"`
}

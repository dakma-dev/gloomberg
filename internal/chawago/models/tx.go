package models

import (
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Topic int64

type TxWithLogs struct {
	*types.Transaction
	*types.Receipt
	Pending bool
}

// getTxMessage is used to get the From field of a transaction.
func (t *TxWithLogs) Sender() *common.Address {
	sender, err := types.LatestSignerForChainID(t.ChainId()).Sender(t.Transaction)
	if err != nil {
		log.Warnf("could not get message for tx %s: %s", t.Hash().Hex(), err)

		return &common.Address{}
	}

	return &sender
}

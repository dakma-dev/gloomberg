package transactioncollector

import (
	"math"
	"math/big"
	"sync"

	"github.com/benleb/gloomberg/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionCollector struct {
	TxID          common.Hash
	LogIndices    []int
	Logs          map[int]*types.Log
	TokenIDs      []*big.Int
	FromAddresses []common.Address
	ToAddresses   []common.Address
	RWMu          *sync.RWMutex
	TX            *types.Transaction
}

func NewTransactionCollector(log *types.Log) *TransactionCollector {
	transco := &TransactionCollector{
		TxID:       log.TxHash,
		LogIndices: []int{},
		Logs:       map[int]*types.Log{},
		RWMu:       &sync.RWMutex{},
		TX:         nil,
	}

	transco.AddLog(log)

	// // parse log topics
	// _, fromAddress, toAddress, tokenID := utils.ParseTopics(log.Topics)

	// transco.LogIndices = append(transco.LogIndices, int(log.Index))
	// transco.TokenIDs = append(transco.TokenIDs, tokenID)

	// transco.FromAddresses = append(transco.FromAddresses, fromAddress)
	// transco.ToAddresses = append(transco.ToAddresses, toAddress)

	// transco.Logs[int(log.Index)] = log

	return transco
}

func (transco *TransactionCollector) AddLog(log *types.Log) {
	transco.RWMu.Lock()
	defer transco.RWMu.Unlock()

	// parse log topics
	_, fromAddress, toAddress, tokenID := utils.ParseTopics(log.Topics)

	transco.LogIndices = append(transco.LogIndices, int(log.Index))

	// transco.TokenIDs = append(transco.TokenIDs, getTokenIDFromTopics(log.Topics))
	transco.TokenIDs = append(transco.TokenIDs, tokenID)

	transco.FromAddresses = append(transco.FromAddresses, fromAddress)
	transco.ToAddresses = append(transco.ToAddresses, toAddress)

	transco.Logs[int(log.Index)] = log
}

func (transco *TransactionCollector) UniqueTokenIDs() uint64 {
	dupeMap := map[*big.Int]bool{}

	for _, tokenID := range transco.TokenIDs {
		if !dupeMap[tokenID] {
			dupeMap[tokenID] = true
		}
	}

	return uint64(math.Max(float64(len(dupeMap)), 1))
}

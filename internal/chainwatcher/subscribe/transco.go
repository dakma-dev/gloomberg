package subscribe

import (
	"math"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionCollector struct {
	TxID       common.Hash
	LogIndices []int
	Logs       map[int]*types.Log
	TokenIDs   []uint64
	RWMu       *sync.RWMutex
	TX         *types.Transaction
}

func NewTransactionCollector(log *types.Log) *TransactionCollector {
	transco := &TransactionCollector{
		TxID:       log.TxHash,
		LogIndices: []int{},
		Logs:       map[int]*types.Log{},
		RWMu:       &sync.RWMutex{},
		TX:         nil,
	}

	tokenID := getTokenIDFromTopics(log.Topics)

	transco.LogIndices = append(transco.LogIndices, int(log.Index))
	transco.TokenIDs = append(transco.TokenIDs, tokenID)

	transco.Logs[int(log.Index)] = log

	return transco
}

func (transco *TransactionCollector) AddLog(log *types.Log) {
	transco.RWMu.Lock()
	defer transco.RWMu.Unlock()

	transco.LogIndices = append(transco.LogIndices, int(log.Index))
	transco.TokenIDs = append(transco.TokenIDs, getTokenIDFromTopics(log.Topics))
	transco.Logs[int(log.Index)] = log
}

func (transco *TransactionCollector) UniqueTokenIDs() int {
	dupeMap := map[uint64]bool{}

	for _, tokenID := range transco.TokenIDs {
		if !dupeMap[tokenID] {
			dupeMap[tokenID] = true
		}
	}

	return int(math.Max(float64(len(dupeMap)), 1))
}

func getTokenIDFromTopics(topics []common.Hash) uint64 {
	// parse token id
	var tokenID uint64
	if len(topics) >= 4 {
		tokenID = topics[3].Big().Uint64()
	} else {
		tokenID = 0
	}

	return tokenID
}

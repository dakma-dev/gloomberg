package transactioncollector

import (
	"math"
	"math/big"
	"sync"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
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
	Style         lipgloss.Style
}

func NewTransactionCollector(log *types.Log) *TransactionCollector {
	transco := &TransactionCollector{
		TxID:       log.TxHash,
		LogIndices: []int{},
		Logs:       map[int]*types.Log{},
		RWMu:       &sync.RWMutex{},
		TX:         nil,
		Style:      lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(int64(log.TxHash.Big().Int64()))),
	}

	go transco.AddLog(log)

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
	transco.TokenIDs = append(transco.TokenIDs, tokenID)
	transco.FromAddresses = append(transco.FromAddresses, fromAddress)
	transco.ToAddresses = append(transco.ToAddresses, toAddress)

	transco.Logs[int(log.Index)] = log

	sameLen := (len(transco.TokenIDs) == len(transco.ToAddresses)) && (len(transco.FromAddresses) == len(transco.ToAddresses))

	gbl.Log.Debugf("collector %v: added log with idx %v | %v | IDs: %d | From: %d | To: %d", transco.Style.Render(transco.TxID.Hex()[:4]), log.Index, sameLen, len(transco.TokenIDs), len(transco.FromAddresses), len(transco.ToAddresses))
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

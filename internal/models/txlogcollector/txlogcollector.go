package txlogcollector

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxLogCollector struct {
	FromAddresses []common.Address
	LogIndices    []int
	Logs          map[int]*types.Log
	RWMu          *sync.RWMutex
	Style         lipgloss.Style
	ToAddresses   []common.Address
	TokenIDs      []*big.Int
	Tx            *types.Transaction
	TxID          common.Hash
}

func NewTxLogCollector(log *types.Log) *TxLogCollector {
	transco := &TxLogCollector{
		TxID:       log.TxHash,
		LogIndices: []int{},
		Logs:       map[int]*types.Log{},
		RWMu:       &sync.RWMutex{},
		Tx:         nil,
		Style:      lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(log.TxHash.Big().Int64())),
	}

	go transco.AddLog(log)

	return transco
}

func (transco *TxLogCollector) AddLog(log *types.Log) {
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

func (transco *TxLogCollector) UniqueTokenIDs() int {
	uniqueTokenIDs := make(map[uint64]*big.Int, 0)

	statusUniqueTokenIDs := strings.Builder{}
	statusUniqueTokenIDs.WriteString("\n")

	for _, tokenID := range transco.TokenIDs {
		uniqueTokenIDs[tokenID.Uint64()] = tokenID
	}

	statusUniqueTokenIDs.WriteString(fmt.Sprintf("transco.TokenIDs: %d | uniqueTokenIDs: %d\n", len(transco.TokenIDs), len(uniqueTokenIDs)))
	statusUniqueTokenIDs.WriteString(fmt.Sprintf("uniqueTokenIDs: %+v\n", uniqueTokenIDs))
	statusUniqueTokenIDs.WriteString("\n")

	if len(transco.TokenIDs) != len(uniqueTokenIDs) {
		gbl.Log.Debugf(statusUniqueTokenIDs.String())
	}

	if len(uniqueTokenIDs) <= 0 {
		return 1
	} else {
		return len(uniqueTokenIDs)
	}
}

func (transco *TxLogCollector) UniqueLogIndices() int {
	dupeMap := map[int]bool{}

	for _, logIdx := range transco.LogIndices {
		if !dupeMap[logIdx] {
			dupeMap[logIdx] = true
		}
	}

	return len(dupeMap)
}

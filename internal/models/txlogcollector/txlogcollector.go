package txlogcollector

import (
	"sync"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxLogs struct {
	RWMu          *sync.RWMutex
	FromAddresses []common.Address
	ToAddresses   []common.Address
	TokenSeller   map[uint64]common.Address
	style         lipgloss.Style
	txID          common.Hash
}

func NewTxLogCollector(log *types.Log) *TxLogs {
	transco := &TxLogs{
		txID:        log.TxHash,
		TokenSeller: make(map[uint64]common.Address, 0),
		RWMu:        &sync.RWMutex{},
		style:       lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(log.TxHash.Big().Int64())),
	}

	go transco.AddLog(log)

	return transco
}

func (transco *TxLogs) AddLog(log *types.Log) {
	transco.RWMu.Lock()
	defer transco.RWMu.Unlock()

	// parse log topics
	_, fromAddress, toAddress, tokenID := utils.ParseTopics(log.Topics)

	transco.FromAddresses = append(transco.FromAddresses, fromAddress)
	transco.ToAddresses = append(transco.ToAddresses, toAddress)

	transco.TokenSeller[tokenID.Uint64()] = fromAddress
}

// func (transco *TxLogs) GetTokenSeller() map[uint64]common.Address {
// 	return transco.TokenSeller
// }

// func (transco *TxLogs) UniqueTokenIDs() int {
// 	uniqueTokenIDs := make(map[uint64]*big.Int, 0)

// 	statusUniqueTokenIDs := strings.Builder{}
// 	statusUniqueTokenIDs.WriteString("\n")

// 	for _, tokenID := range transco.TokenIDs {
// 		uniqueTokenIDs[tokenID.Uint64()] = tokenID
// 	}

// 	statusUniqueTokenIDs.WriteString(fmt.Sprintf("transco.TokenIDs: %d | uniqueTokenIDs: %d\n", len(transco.TokenIDs), len(uniqueTokenIDs)))
// 	statusUniqueTokenIDs.WriteString(fmt.Sprintf("uniqueTokenIDs: %+v\n", uniqueTokenIDs))
// 	statusUniqueTokenIDs.WriteString("\n")

// 	if len(transco.TokenIDs) != len(uniqueTokenIDs) {
// 		gbl.Log.Debugf(statusUniqueTokenIDs.String())
// 	}

// 	if len(uniqueTokenIDs) <= 0 {
// 		return 1
// 	} else {
// 		return len(uniqueTokenIDs)
// 	}
// }

// func (transco *TxLogs) UniqueLogIndices() int {
// 	dupeMap := map[int]bool{}

// 	for _, logIdx := range transco.LogIndices {
// 		if !dupeMap[logIdx] {
// 			dupeMap[logIdx] = true
// 		}
// 	}

// 	return len(dupeMap)
// }

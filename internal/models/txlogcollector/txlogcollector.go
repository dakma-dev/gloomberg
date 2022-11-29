package txlogcollector

import (
	"sync"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/models/topic"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxLogs struct {
	RWMu             *sync.RWMutex
	FromAddresses    []common.Address
	ToAddresses      []common.Address
	TokenSeller      map[uint64]common.Address
	style            lipgloss.Style
	txID             common.Hash
	ERC20Logs        []types.Log // offer sales having erc20 txs
	MainLog          *types.Log  // leading log (singletransfer or orderfullfilled)
	ERC721Transfers  []abis.ERC721v3Transfer
	ERC1155Transfers []abis.ERC1155Transfer
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

	logTopic := topic.Topic(log.Topics[0].Hex())
	if logTopic == topic.Transfer && log.Address.Hex() == string(external.WETH) {
		// fmt.Println("WETH transfer found")
		transco.ERC20Logs = append(transco.ERC20Logs, *log)
		return
	}

	_, fromAddress, toAddress, tokenID := utils.ParseTopics(log.Topics)

	if logTopic == topic.Transfer && len(log.Topics) == 4 {
		transco.ERC721Transfers = append(transco.ERC721Transfers, abis.ERC721v3Transfer{
			From:    fromAddress,
			To:      toAddress,
			TokenId: tokenID,
		})
	}

	if logTopic == topic.TransferSingle && len(log.Topics) == 4 {
		transco.ERC1155Transfers = append(transco.ERC1155Transfers, abis.ERC1155Transfer{
			Id:   tokenID, // <-- Token Id
			From: fromAddress,
			To:   toAddress,
			// Value: nil, // <-- amount
			// Raw:   types.Log{},
		})
	}

	// OrderFullfilled Topic or ERC721 Transfer logs could occur multiple times. TODO enhance the log collector we need for each interesting log combination a possible notification, create new sale struct holding from, to and tokenId?
	// this code is just to whitelist the logs which are interesting
	if transco.MainLog == nil {
		mainlogCandidate := false

		// ERC721/1155 Transfer
		if logTopic == topic.Transfer && len(log.Topics) == 4 {
			mainlogCandidate = true
		}

		if logTopic == topic.OrderFulfilled && len(log.Topics) == 3 {
			mainlogCandidate = true
		}

		if logTopic == topic.TransferSingle && len(log.Topics) == 4 {
			mainlogCandidate = true
		}

		if logTopic == topic.OrdersMatched {
			mainlogCandidate = true
		}

		if mainlogCandidate {
			transco.MainLog = log
		}
	}

	// parse log topics

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

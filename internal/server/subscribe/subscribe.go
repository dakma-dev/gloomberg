package subscribe

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/domains"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models/topic"
	"github.com/benleb/gloomberg/internal/notifications"
	"github.com/benleb/gloomberg/internal/server/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

var (
	mu = &sync.Mutex{}

	totalNumReceived  = uint64(0)
	totalLastReceived = int64(0)

	knownTransactions     = make(map[common.Hash][]int)
	transactionCollectors = make(map[common.Hash]*TransactionCollector)
)

func WorkerLogsQueue(workerID int, cNode *node.Node, rawNodes []*node.Node, ownCollections *collections.Collections, queueLogs *chan types.Log, queueEvents *chan *collections.Event, queueOutWS *chan *collections.Event) {
	var cNodes node.Nodes = rawNodes

	gbl.Log.Infof("%d %d| queueLogs worker started - queueLogs: %d", cNode.NodeID, workerID, len(*queueLogs))

	for subLog := range *queueLogs {
		nanoNow := time.Now().UnixNano()

		// total
		atomic.AddUint64(&totalNumReceived, 1)
		atomic.StoreInt64(&totalLastReceived, nanoNow)

		// per node
		atomic.AddUint64(&cNode.NumLogsReceived, 1)
		atomic.StoreInt64(&cNode.LastLogReceived, nanoNow)

		if subLog.Address == common.HexToAddress("0x042874309Bf3F6C8E69Be4bf3D251fE9e41CF0d2") {
			fmt.Println("")
			fmt.Println("")
			fmt.Println(" ‼️ 🤳 Impostergram 💄 Impostergram 🤳 Impostergram 💄 Impostergram 🤳 ‼️")
			fmt.Println("")
			fmt.Println("  https://foundation.app/collection/impostergram")
			fmt.Println("")
			fmt.Println("")
			// notifications.SendAlert("Impostergram 🤳 💄", "https://foundation.app/collection/impostergram", true)
			if msg, err := notifications.SendTelegramMessage(1320669206, " @benleb ‼️ 🤳 Impostergram 💄 Impostergram 🤳 Impostergram 💄 Impostergram 🤳 !!\n\nhttps://foundation.app/collection/impostergram", ""); err != nil {
				gbl.Log.Errorf("failed to send telegram message: %v | imageURI: %s | err: %s\n", msg, "", err)
			}
		}

		// erc721 has 0-3, (erc1155 has topics 2?), erc20 has topics 0-2
		if len(subLog.Topics) != 4 {
			gbl.Log.Debugf("DiscardedOtherERC| %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
			continue
		}

		go parseTransferLog(cNode, &cNodes, ownCollections, subLog, queueEvents, queueOutWS)

		if totalNumReceived%1000 == 0 {
			// for _, node := range nodes {
			// 	gbl.Log.Infof("received logs %s: %d | %s", node.Name, node.NumLogsReceived, time.Unix(0, node.LastLogReceived).Format("15:04:05.000000000"))
			// }
			gbl.Log.Infof("  > received: %d | queueLogs: %d | queueEvents: %d | queueOutWS: %d <", totalNumReceived, len(*queueLogs), len(*queueEvents), len(*queueOutWS))
		}
	}
}

func parseTransferLog(cNode *node.Node, cNodes *node.Nodes, ownCollections *collections.Collections, subLog types.Log, queueEvents *chan *collections.Event, queueOutWS *chan *collections.Event) {
	// transaction collector to "recognize" multi-item txs
	var transco *TransactionCollector

	mu.Lock()
	if tc := transactionCollectors[subLog.TxHash]; tc != nil {
		transco = tc
		transco.AddLog(&subLog)
		mu.Unlock()

		return
	}

	transco = NewTransactionCollector(&subLog)
	transactionCollectors[subLog.TxHash] = transco

	mu.Unlock()

	time.Sleep(100 * time.Millisecond)

	// check if we have seen this logIndex for this transaction before
	// isMultiItemTx := multiItemNumber > 1
	logIndex := int(subLog.Index)

	mu.Lock()

	for _, lidx := range knownTransactions[subLog.TxHash] {
		if lidx == logIndex {
			mu.Unlock()

			// if we know the tx (from another node provider or ...) we don't need to do anything
			gbl.Log.Warnf("discarded already known tx: %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			return
		}
	}

	knownTransactions[subLog.TxHash] = append(knownTransactions[subLog.TxHash], logIndex)

	mu.Unlock()

	// parse topics
	logTopic, fromAddress, toAddress, rawTokenID := parseTopics(subLog.Topics)
	tokenID := rawTokenID.Uint64()

	// collection information
	ownCollections.RWMu.RLock()
	collection := ownCollections.UserCollections[subLog.Address]
	ownCollections.RWMu.RUnlock()

	if collection == nil && subLog.Address != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		collection = collections.NewCollection(subLog.Address, "", cNodes, collections.Stream)

		ownCollections.RWMu.Lock()
		ownCollections.UserCollections[subLog.Address] = collection
		ownCollections.RWMu.Unlock()

		if collection == nil {
			// atomic.AddUint64(&StatsBTV.DiscardedUnknownCollection, 1)
			gbl.Log.Warnf("discarded unknown collection: %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			return
		}
	}

	isMint := common.HexToAddress(subLog.Topics[1].Hex()).String() == "0x0000000000000000000000000000000000000000"
	showMints := viper.GetBool("show.mints") // || collection.Show.Mints
	isOwnCollection := collection.Source == collections.Wallet || collection.Source == collections.Configuration

	// value is just fetched for sales, not for mints
	value := big.NewInt(0)

	var eventType collections.EventType

	if isMint {
		// mint
		if !showMints && !collection.Show.Mints {
			// atomic.AddUint64(&StatsBTV.DiscardedMints, 1)
			gbl.Log.Debugf("‼️ mint not shown %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			return
		}

		eventType = collections.Mint
	} else {
		// sale | get the tx details - we don't do this for mints to save a lot of api calls

		// get the transaction details
		tx, _, err := cNodes.GetRandomNode().Client.TransactionByHash(context.Background(), subLog.TxHash)
		if err != nil {
			gbl.Log.Warnf("getting tx details failed: %s | %+v", subLog.TxHash.Hex(), err)
			// atomic.AddUint64(&StatsBTV.DiscardedTransactions, 1)

			return
		}

		// set to actual tx value
		value = tx.Value()

		eventType = collections.Sale
	}

	// if the tx has no 'value' (and is not a mint) it is a transfer
	isTransfer := value.Cmp(big.NewInt(0)) == 0 && !isMint
	showTransfers := viper.GetBool("show.transfers") || collection.Show.Transfers

	if !isMint && !isOwnCollection && node.WeiToEther(value).Cmp(big.NewFloat(viper.GetFloat64("show.min_price"))) < 0 {
		// atomic.AddUint64(&StatsBTV.DiscardedLowPrice, 1)
		gbl.Log.Debugf("‼️ DiscardedLowPrice| %v | TxHash: %v / %d | %+v", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

		return
	}

	if isTransfer {
		// transfer
		if !showTransfers {
			// atomic.AddUint64(&StatsBTV.DiscardedTransfers, 1)
			gbl.Log.Debugf("‼️ transfer not shown %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
			return
		}

		eventType = collections.Transfer
	}

	// if its an ENS nft, we try to get the name from the ens metadata service
	var domainENS string

	if collection.ContractAddress == common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85") {
		// set custom collection name
		collection.Name = "ENS"

		// get ens metadata, primarily the name
		ensMetadata, err := domains.GetMetadataForTokenID(rawTokenID)
		if err == nil && ensMetadata != nil {
			domainENS = ensMetadata.Name
		} else {
			gbl.Log.Warnf("getting ens metadata failed: %s | %s", err, fmt.Sprint(rawTokenID))
		}
	}

	event := &collections.Event{
		NodeID:      cNode.NodeID,
		EventType:   eventType,
		Topic:       logTopic.String(),
		TxHash:      subLog.TxHash,
		Collection:  collection,
		TokenID:     tokenID,
		DomainENS:   domainENS,
		PriceWei:    value,
		TxItemCount: uint(transco.UniqueTokenIDs()),
		Time:        time.Now(),
		From: collections.User{
			Address:       fromAddress,
			OpenseaUserID: "",
		},
		To: collections.User{
			Address:       toAddress,
			OpenseaUserID: "",
		},
	}

	// send to formatting
	*queueEvents <- event

	// send to websockets output
	if viper.GetBool("server.websockets.enabled") {
		*queueOutWS <- event
	}

	gbCache := cache.New()
	gbCache.StoreEvent(event.Collection.ContractAddress, event.Collection.Name, event.TokenID, event.PriceWei.Uint64(), event.TxItemCount, event.Time, int64(eventType))
}

func parseTopics(topics []common.Hash) (topic.Topic, common.Address, common.Address, *big.Int) {
	logTopic := topic.Topic(topics[0].Hex())

	// parse from/to addresses
	var fromAddress, toAddress common.Address
	if logTopic == topic.Transfer {
		fromAddress = common.HexToAddress(topics[1].Hex())
		toAddress = common.HexToAddress(topics[2].Hex())
	} else if logTopic == topic.TransferSingle {
		fromAddress = common.HexToAddress(topics[2].Hex())
		toAddress = common.HexToAddress(topics[3].Hex())
	}

	// parse token id
	var rawTokenID *big.Int
	if len(topics) >= 4 {
		rawTokenID = topics[3].Big()
	} else {
		rawTokenID = big.NewInt(0)
	}

	return logTopic, fromAddress, toAddress, rawTokenID
}

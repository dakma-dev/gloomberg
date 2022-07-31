package subscriptions

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
)

var (
	mu = &sync.Mutex{}

	knownTransactions     = make(map[common.Hash][]int)
	transactionCollectors = make(map[common.Hash]*TransactionCollector)

	ensContractAddress = common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")
)

func SubscriptionLogsHandler(nodes *gbnode.NodeCollection, gOwnCollections *collections.Collections, queueLogs chan types.Log, queueEvents chan *collections.Event) {
	for subLog := range queueLogs {
		// atomic.AddUint64(&stats.queueEvents, 1)
		gbl.Log.Debugf("%s | new subscription log (%d): %+v", time.Now().String(), len(queueLogs), subLog)

		// erc721 has 0-3, (erc1155 has topics 2?), erc20 has topics 0-2
		if len(subLog.Topics) != 4 {
			gbl.Log.Debugf("DiscardedOtherERC| %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			// atomic.AddUint64(&StatsBTV.DiscardedOtherERC, 1)

			gbl.Log.Debugf("DiscardedOtherERC| %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			continue
		}

		if subLog.Address == ensContractAddress {
			continue
		}

		go parseLog(nodes, gOwnCollections, subLog, queueEvents)
	}
}

func parseLog(nodes *gbnode.NodeCollection, ownCollections *collections.Collections, subLog types.Log, queueEvents chan *collections.Event) {
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
			// atomic.AddUint64(&StatsBTV.DiscardedAlreadyKnownTX, 1)

			gbl.Log.Warnf("DiscardedAlreadyKnownTX| %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			return
		}
	}

	knownTransactions[subLog.TxHash] = append(knownTransactions[subLog.TxHash], logIndex)

	mu.Unlock()

	// parse topics
	_, _, _, tokenID := parseTopics(subLog.Topics)

	// collection information
	ownCollections.RWMu.RLock()
	// collection := ownCollections.Collections[subLog.Address]
	collection := ownCollections.UserCollections[subLog.Address]
	// collection := gOwnCollections.DiscoveredCollections[subLog.Address]
	ownCollections.RWMu.RUnlock()

	if collection == nil && subLog.Address != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		collection = collections.NewCollection(subLog.Address, "", nodes, collections.Stream)

		ownCollections.RWMu.Lock()
		ownCollections.UserCollections[subLog.Address] = collection
		ownCollections.RWMu.Unlock()

		if collection == nil {
			// atomic.AddUint64(&StatsBTV.DiscardedUnknownCollection, 1)
			gbl.Log.Warnf("DiscardedUnknownCollection| %v | TxHash: %v / %d | %+v\n", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

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
			return
		}

		eventType = collections.Mint
	} else {
		// sale | get the tx details - we don't do this for mints to save a lot of api calls

		// get the transaction details
		tx, _, err := nodes.GetRandomNode().Client.TransactionByHash(context.Background(), subLog.TxHash)
		if err != nil {
			gbl.Log.Infof("getting tx details failed | %v | TxHash: %v / %d | %+v", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
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

	if !isMint && !isOwnCollection && WeiToEther(value).Cmp(big.NewFloat(viper.GetFloat64("show.min_price"))) < 0 {
		// atomic.AddUint64(&StatsBTV.DiscardedLowPrice, 1)
		gbl.Log.Debugf("DiscardedLowPrice| %v | TxHash: %v / %d | %+v", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

		return
	}

	if isTransfer {
		// transfer
		if !showTransfers {
			// atomic.AddUint64(&StatsBTV.DiscardedTransfers, 1)
			return
		}

		eventType = collections.Transfer
	}

	// if btvEventType == models.Sale {
	// 	atomic.AddUint64(&collection.LastRoundSales, uint64(math.Max(float64(len(transco.Logs)), 1)))
	// }

	// gbl.Log.Debugf(
	// 	"\nlen(data.Topics): %d\nStatsBTV.DiscardedOtherERC: %d\nStatsBTV.DiscardedAlreadyKnownTX: %d\nStatsBTV.DiscardedMints: %d\nStatsBTV.DiscardedTransfers: %d\nStatsBTV.DiscardedLowPrice: %d\nStatsBTV.DiscardedUnknownCollection: %d\nStatsBTV.DiscardedTransactions: %d",
	// 	len(subLog.Topics), StatsBTV.DiscardedOtherERC, StatsBTV.DiscardedAlreadyKnownTX, StatsBTV.DiscardedMints, StatsBTV.DiscardedTransfers, StatsBTV.DiscardedLowPrice, StatsBTV.DiscardedUnknownCollection, StatsBTV.DiscardedTransactions,
	// )

	logTopic := gbnode.Topic(subLog.Topics[0].Hex())

	// parse from/to addresses
	var fromAddress, toAddress common.Address
	if logTopic == gbnode.Transfer {
		fromAddress = common.HexToAddress(subLog.Topics[1].Hex())
		toAddress = common.HexToAddress(subLog.Topics[2].Hex())
	} else if logTopic == gbnode.TransferSingle {
		fromAddress = common.HexToAddress(subLog.Topics[2].Hex())
		toAddress = common.HexToAddress(subLog.Topics[3].Hex())
	}

	// tokenID := GetTokenIDFromTopics(subLog.Topics)

	event := &collections.Event{
		EventType:   eventType,
		Topic:       logTopic.String(),
		TxHash:      subLog.TxHash,
		Collection:  collection,
		TokenID:     tokenID,
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
		// WorkerID: workerID,
	}

	// send to formatting
	queueEvents <- event

	// *outputWs <- event

	// xAddArgs := &redis.XAddArgs{
	// 	Stream: "sales",
	// 	MaxLen: 100000,
	// 	Approx: true,
	// 	ID:     "*",
	// 	Values: map[string]any{
	// 		"collection": event.Collection.Name,
	// 		"tokenID":    int(event.TokenID),
	// 		"priceWei":   event.PriceWei.Uint64(),
	// 		"numItems":   len(transco.Logs),
	// 		"time":       event.Time,
	// 		"event_type": strconv.FormatInt(int64(event.EventType), 10),
	// 	},
	// }

	// go addSaleEventToCache(xAddArgs, event.Collection.Name, event.TokenID)

	// gbl.Log.Warnf("updating gui now... %p", g)

	// g.Update(func(g *gocui.Gui) error {
	// 	streamView, err := g.View("main")
	// 	if err != nil {
	// 		gbl.Log.Errorf("error getting streamView: %+v", err.Error())
	// 		return err
	// 	}

	// 	gbl.Log.Warnf("streamView: %p", streamView)

	// 	// logLine = fmt.Sprintf("%s | %s %d\n", time.Now().String(), subLog.Address.String(), len(*subscriptionLogs))
	// 	logLine := fmt.Sprintf("%s| %s #%d |  %s  ->  %s\n", topic.String(), collection.Name, tokenID, from.String(), to.String())

	// 	if _, err = streamView.Write([]byte(logLine)); err != nil {
	// 		gbl.Log.Errorf("error writing streamView: %+v", err.Error())
	// 		return err
	// 	}

	// 	return nil
	// })

	// gbl.Log.Warnf("...done %p\n", g)
	gbl.Log.Debugf("...done")
}

func GetTokenIDFromTopics(topics []common.Hash) uint64 {
	// parse token id
	var tokenID uint64
	if len(topics) >= 4 {
		tokenID = topics[3].Big().Uint64()
	} else {
		tokenID = 0
	}

	return tokenID
}

func parseTopics(topics []common.Hash) (gbnode.Topic, common.Address, common.Address, uint64) {
	topic := gbnode.Topic(topics[0].Hex())
	from := common.HexToAddress(topics[1].Hex())
	to := common.HexToAddress(topics[2].Hex())

	// parse token id
	var tokenID uint64
	if len(topics) >= 4 {
		tokenID = topics[3].Big().Uint64()
	} else {
		tokenID = 0
	}

	return topic, from, to, tokenID
}

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

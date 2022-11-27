package chainwatcher

import (
	"context"
	"fmt"
	"github.com/benleb/gloomberg/internal/abis"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/topic"
	"github.com/benleb/gloomberg/internal/models/txlogcollector"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/ws"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/spf13/viper"
)

type ChainWatcher struct {
	Nodes *nodes.Nodes

	CollectionDB *collections.CollectionDB
	// CollectionDB   *collections.CollectionDB

	queueLogs  *chan types.Log
	queueOutWS *chan *collections.Event

	WebsocketsServer *ws.WebsocketsServer

	wethContract *abis.ERC20
}

var (
	mu                = &sync.Mutex{}
	knownTransactions = make(map[common.Hash][]int)
	logCollectors     = make(map[common.Hash]*txlogcollector.TxLogs)

	txLogsReceivedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_transaction_logs_received_total",
		Help: "The total number of received transaction logs",
	})

	txReceivedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_transactions_received_total",
		Help: "The total number of received transactions",
	})

	txReceivedSale = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_transactions_received_sale",
		Help: "The total number of received sale transactions",
	})

	txReceivedMint = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_transactions_received_mint",
		Help: "The total number of received mint transactions",
	})

	txReceivedTransfer = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_transactions_received_transfer",
		Help: "The total number of received transfer transactions",
	})
)

type GItem struct {
	ItemType   uint8          `json:"itemType"`
	Token      common.Address `json:"token"`
	Identifier *big.Int       `json:"identifier"`
	Amount     *big.Int       `json:"amount"`
}

func New(nodes *nodes.Nodes, collectiondb *collections.CollectionDB) *ChainWatcher {
	// create a queue/channel for the received logs
	queueLogs := make(chan types.Log, 1024)
	// create a queue/channel for events to be sent out via ws
	queueOutWS := make(chan *collections.Event, 1024)

	// TODO move to new instance holding these?
	wethAddress := common.HexToAddress(string(external.WETH))
	wethContract, _ := abis.NewERC20(wethAddress, nodes.GetRandomNode().Client)

	return &ChainWatcher{
		CollectionDB: collectiondb,
		Nodes:        nodes,
		queueLogs:    &queueLogs,
		queueOutWS:   &queueOutWS,
		wethContract: wethContract,
	}
}
func (cw *ChainWatcher) GetLogsByBlockNumber(blockNumber int64) {
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
	}

	logs, err := cw.Nodes.GetRandomNode().Client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		return
	}

	for _, log := range logs {
		*cw.queueLogs <- log
	}

}

func (cw *ChainWatcher) SubscribeToSales(queueEvents *chan *collections.Event) {
	for _, node := range *cw.Nodes {
		gbl.Log.Debugf("%s: subscribing to chain events | QueueEvents: %d", node.Name, len(*queueEvents))

		// subscribe to all events where first topic is the "Transfer" topic
		if _, err := node.SubscribeToTransfers(*cw.queueLogs); err != nil {
			gbl.Log.Warnf("Transfers subscribe to %s failed: %s", node.WebsocketsEndpoint, err)
		}
		// subscribe to all events where first topic is the "SingleTransfer" topic
		if _, err := node.SubscribeToSingleTransfers(*cw.queueLogs); err != nil {
			gbl.Log.Warnf("SingleTransfers subscribe to %s failed: %s", node.WebsocketsEndpoint, err)
		}

		// create a defined number of workers/handlers per cNode to receive and process incoming events/logs
		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			go cw.logHandler(node, queueEvents)
		}
	}
}

func (cw *ChainWatcher) SubscribeToOrderFulfilled(queueEvents *chan *collections.Event) {
	for _, node := range *cw.Nodes {
		gbl.Log.Debugf("%s: subscribing to chain events | QueueEvents: %d", node.Name, len(*queueEvents))

		// subscribe to all events where first topic is the "Transfer" topic
		if _, err := node.SubscribeToOrderFulfilled(*cw.queueLogs); err != nil {
			gbl.Log.Warnf("Transfers subscribe to %s failed: %s", node.WebsocketsEndpoint, err)
		}

		// create a defined number of workers/handlers per cNode to receive and process incoming events/logs
		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			go cw.logHandler(node, queueEvents)
		}
	}
}

func (cw *ChainWatcher) logHandler(node *nodes.Node, queueEvents *chan *collections.Event) {
	// process new logs received via our subscriptions
	for subLog := range *cw.queueLogs {
		// increment prometheus counter
		txLogsReceivedTotal.Inc()

		// track & count
		nanoNow := time.Now().UnixNano()
		// logs per node
		atomic.AddUint64(&node.NumLogsReceived, 1)
		atomic.StoreInt64(&node.LastLogReceived, nanoNow)

		// discard Transfer/TransferSingle logs for non-NFT transfers | erc20: topics 0-2 | erc721/1155: 0-3
		// if (logTopic == topic.Transfer || logTopic == topic.TransferSingle) && len(subLog.Topics) < 4 {
		if len(subLog.Topics) < 3 {
			gbl.Log.Debugf("🗑️ number of topics in log is %d (!= 3) | %v | TxHash: %v / %d | %+v", len(subLog.Topics), subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
			continue
		}

		// parse log topics
		logTopic, _, _, _ := utils.ParseTopics(subLog.Topics)

		//
		// distribute to parser depending on log topic
		switch logTopic {
		case topic.Transfer, topic.TransferSingle:
			// parse generic transfer topics
			go cw.LogParserTransfers(node.NodeID, subLog, queueEvents)
			// case topic.OrderFulfilled:
			// 	// parse opensea seaport OrderFulfilled logs
			// 	go cw.logParserOrderFulfilled(node.NodeID, subLog, queueEvents)
		}

		//
		// distribute to parser depending on initiator addresses
		// if WatchedWallets.Contain(fromAddress) ...
		//
		// distribute to parser depending on target/contract addresses
		// if WatchedContracts.Contain(toAddress) ...
	}
}

func (cw *ChainWatcher) LogParserTransfers(nodeID int, subLog types.Log, queueEvents *chan *collections.Event) {
	var eventDiscarded *collections.EventDiscarded

	// parse log topics
	logTopic, fromAddress, toAddress, tokenID := utils.ParseTopics(subLog.Topics)

	// if logTopic == topic.OrderFulfilled {
	// 	// get the contractERC721 ABIs
	// 	_, err := abis.NewSeaport(subLog.Address, cw.Nodes.GetRandomLocalNode().Client)
	// 	if err != nil {
	// 		gbl.Log.Error(err)
	// 	}

	// 	// seaportABIFile, _ := os.Open("internal/abis/seaport11.json")
	// 	// seaport, _ := abi.JSON(seaportABIFile)

	// 	// dataMap := make(map[string]interface{})

	// 	// if err := seaport.UnpackIntoMap(dataMap, "OrderFulfilled", subLog.Data); err != nil {
	// 	// 	gbl.Log.Errorf("error unpacking into map: %s", err)
	// 	// 	fmt.Printf("error unpacking into map: %s\n", err)
	// 	// }

	// 	// orderFulfilled, _ := abiSeaport.ParseOrderFulfilled(subLog)
	// 	// fmt.Printf("orderFulfilled: %+v\n", orderFulfilled)
	// 	// return
	// }

	//
	// we use a "transaction collector" to "recognize" (wait for) multi-item tx logs
	mu.Lock()

	// check if we already have a collector for this tx hash
	if tc := logCollectors[subLog.TxHash]; tc != nil {
		// if we have a collector, we can add this log/log index to the collector
		tc.AddLog(&subLog)
		mu.Unlock()
		return
	}

	// if we don't have a collector, we create a new one for this tx hash
	txLogs := txlogcollector.NewTxLogCollector(&subLog)
	logCollectors[subLog.TxHash] = txLogs

	mu.Unlock()

	// wait for all logs of this tx to be received
	time.Sleep(97 * time.Millisecond)

	//
	// check if we have seen this logIndex for this transaction before
	logIndex := int(subLog.Index)

	mu.Lock()

	// check if the log is already known to us
	for _, lIdx := range knownTransactions[subLog.TxHash] {
		if lIdx == logIndex {
			mu.Unlock()
			return
		}
	}

	// if we don't have this logIndex, we add it to the list of known logs for this tx
	knownTransactions[subLog.TxHash] = append(knownTransactions[subLog.TxHash], logIndex)

	mu.Unlock()

	// get "main" log (ignore erc20 logs)

	if txLogs.MainLog == nil {
		gbl.Log.Infof("No main log found (!= erc02transfer). Skipping")
		return
	}

	subLog = *txLogs.MainLog

	// increment transactions counter
	txReceivedTotal.Inc()

	//
	// collection information
	cw.CollectionDB.RWMu.RLock()
	collection := cw.CollectionDB.Collections[subLog.Address]
	cw.CollectionDB.RWMu.RUnlock()

	if collection == nil && subLog.Address != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		name := ""

		if logTopic == topic.TransferSingle {
			if tokenName, err := cw.Nodes.GetERC1155TokenName(subLog.Address, tokenID); err == nil && tokenName != "" {
				name = tokenName
				gbl.Log.Debugf("found token name: %s | %s", name, subLog.Address.String())
			} else if err != nil {
				gbl.Log.Debugf("failed to get collection name: %s", err)
			}
		}

		collection = collections.NewCollection(subLog.Address, name, cw.Nodes, models.FromStream)

		cw.CollectionDB.RWMu.Lock()
		cw.CollectionDB.Collections[subLog.Address] = collection
		cw.CollectionDB.RWMu.Unlock()

		if collection == nil {
			// atomic.AddUint64(&StatsBTV.DiscardedUnknownCollection, 1)
			gbl.Log.Warnf("🗑️ collection is nil | cw.CollectionDB.UserCollections[subLog.Address] -> %v | %v | TxHash: %v / %d | %+v", cw.CollectionDB.Collections[subLog.Address], subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
			return
		}
	}

	// create a collection name if we can't find one
	if collection.Name == "" {
		// preSuffix := collection.StyleSecondary().Copy().Faint(true).Render("??")
		// name := collection.Style().Copy().Faint(true).Italic(true).Render("Unknown " + logTopic.String())
		// collection.Name = preSuffix + " " + name + " " + preSuffix
		preSuffix := "??"
		name := "Unknown " + logTopic.String()
		collection.Name = preSuffix + " " + name + " " + preSuffix
	}

	// further (tx) information as booleans
	isMint := fromAddress == utils.ZeroAddress
	showMints := viper.GetBool("show.mints") || collection.Show.Mints
	isOwnCollection := collection.Source == models.FromWallet || collection.Source == models.FromConfiguration

	var eventType collections.EventType

	var tx *types.Transaction

	var txErr error

	if isMint {
		// increment mints counter
		txReceivedMint.Inc()

		eventType = collections.Mint

		if !showMints {
			// atomic.AddUint64(&StatsBTV.DiscardedMints, 1)
			gbl.Log.Debugf("🗑️ showMints -> %v | collection.Show.Mints -> %v | %v | TxHash: %v / %d | %+v", showMints, collection.Show.Mints, subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
			// return

			// printEvent = false
			reason := "showing mints is disabled"

			if eventDiscarded == nil {
				eventDiscarded = &collections.EventDiscarded{
					DiscardedBy: "chainwatcher",
					Reasons:     []string{reason},
				}
			} else {
				eventDiscarded.Reasons = append(eventDiscarded.Reasons, reason)
			}
		}
	} else {
		eventType = collections.Sale
		if collection.IgnorePrinting {
			// printEvent = false
			reason := "IgnorePrinting is set"
			if eventDiscarded == nil {
				eventDiscarded = &collections.EventDiscarded{
					DiscardedBy: "chainwatcher",
					Reasons:     []string{reason},
				}
			} else {
				eventDiscarded.Reasons = append(eventDiscarded.Reasons, reason)
			}
		}
	}

	// get the transaction details - we do this for mints to save a lot of calls
	if isMint {
		tx, _, txErr = cw.Nodes.GetTransactionByHashFromLocalNode(context.Background(), subLog.TxHash)
	} else {
		tx, _, txErr = cw.Nodes.GetTransactionByHash(context.Background(), subLog.TxHash)
	}

	if txErr != nil {
		gbl.Log.Debugf("🗑️ getting tx details failed | %v | TxHash: %v / %d | %+v", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)
		// atomic.AddUint64(&StatsBTV.DiscardedTransactions, 1)
		return
	}

	// set to actual tx value
	value := tx.Value()

	isAcceptedOffer := false
	if len(txLogs.ERC20Logs) != 0 {
		for _, transferLog := range txLogs.ERC20Logs {
			if transferLog.Address.Hex() == string(external.WETH) {
				transfer, err := cw.wethContract.ParseTransfer(*transferLog)
				if err != nil {
					gbl.Log.Error(err)
					return
				}
				value.Add(value, transfer.Value)
				isAcceptedOffer = true
			}
		}
	}

	// if the tx has no 'value' (and is not a mint) it is a transfer
	isTransfer := value.Cmp(big.NewInt(0)) == 0 && !isMint // && logTopic != topic.TransferSingle
	showTransfers := viper.GetBool("show.transfers") || collection.Show.Transfers

	if !isMint && !isOwnCollection && nodes.WeiToEther(value).Cmp(big.NewFloat(viper.GetFloat64("show.min_value"))) < 0 {
		// atomic.AddUint64(&StatsBTV.DiscardedLowPrice, 1)
		gbl.Log.Debugf("‼️ DiscardedLowPrice| %v | TxHash: %v / %d | %+v", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

		// printEvent = false
		reason := "price below min_value"

		if eventDiscarded == nil {
			eventDiscarded = &collections.EventDiscarded{
				DiscardedBy: "chainwatcher",
				Reasons:     []string{reason},
			}
		} else {
			eventDiscarded.Reasons = append(eventDiscarded.Reasons, reason)
		}
	}

	if isTransfer {
		// increment transfers counter
		txReceivedTransfer.Inc()

		eventType = collections.Transfer

		if !showTransfers {
			// atomic.AddUint64(&StatsBTV.DiscardedTransfers, 1)
			gbl.Log.Debugf("‼️ transfer not shown %v | TxHash: %v / %d | %+v", subLog.Address.String(), subLog.TxHash, subLog.TxIndex, subLog)

			// printEvent = false
			reason := "showing transfers is disabled"

			if eventDiscarded == nil {
				eventDiscarded = &collections.EventDiscarded{
					DiscardedBy: "chainwatcher",
					Reasons:     []string{reason},
				}
			} else {
				eventDiscarded.Reasons = append(eventDiscarded.Reasons, reason)
			}
		}
	}

	// if it is an ENS nft, we try to get the name from the ens metadata service
	var ensMetadata *external.ENSMetadata = nil

	if collection.ContractAddress == external.ENSContract {
		// set custom collection name
		collection.Name = "ENS"

		// get ens token metadata
		metadata, err := external.GetENSMetadataForTokenID(tokenID)
		if err == nil && metadata != nil {
			ensMetadata = metadata
		} else {
			gbl.Log.Debugf("getting ens metadata failed: %s | %v", fmt.Sprint(tokenID), err)
		}
	}

	if logTopic == topic.TransferSingle {
		if tID := cw.Nodes.GetERC1155TokenID(subLog.Data); tID != nil {
			tokenID = tID
		}
	}

	// if collection.SupportedStandards.Contains(standard.ERC1155) { // && (tokenID.Cmp(big.NewInt(0)) > 0 && tokenID.Cmp(big.NewInt(999_999)) < 0) {
	// 	if tID := cw.Nodes.GetERC1155TokenID(subLog.Address, subLog.Data); tID != nil {
	// 		tokenID = tID
	// 	}
	// }

	numItems := len(txLogs.TokenSeller)
	// numItems := txLogs.UniqueTokenIDs()
	// numItems := len(txLogCollector.Logs)
	// numItems := len(txLogCollector.TokenIDs)
	// numItems := uint64(math.Max(float64(len(txLogCollector.UniqueTokenIDs())), 1))

	fromAddresses := make(map[common.Address]bool, 0)
	fromAddresses[fromAddress] = true

	for _, address := range txLogs.FromAddresses {
		fromAddresses[address] = true
	}

	toAddresses := make(map[common.Address]bool, 0)
	toAddresses[toAddress] = true

	for _, address := range txLogs.ToAddresses {
		toAddresses[address] = true
	}

	var priceArrowColor lipgloss.Color

	if eventType == collections.Sale {
		// increment sales counter
		txReceivedSale.Inc()

		// get a color with saturation depending on the tx price
		priceEther, _ := nodes.WeiToEther(value).Float64()
		priceArrowColor = style.GetPriceShadeColor(priceEther)
	} else {
		priceArrowColor = "#333333"
	}

	event := &collections.Event{
		NodeID:    nodeID,
		EventType: eventType,
		Topic:     logTopic.String(),
		TxHash:    subLog.TxHash,
		// TransactionCollector: txLogCollector,
		Collection:      collection,
		ContractAddress: collection.ContractAddress,
		TokenID:         tokenID,
		ENSMetadata:     ensMetadata,
		PriceWei:        value,
		PriceArrowColor: priceArrowColor,
		TxLogCount:      uint64(numItems),
		// UniqueTokenIDs:  txLogCollector.UniqueTokenIDs(),
		Time: time.Now(),
		From: collections.User{
			Address:       fromAddress,
			OpenseaUserID: "",
		},
		To: collections.User{
			Address:       toAddress,
			OpenseaUserID: "",
		},
		FromAddresses:    fromAddresses,
		ToAddresses:      toAddresses,
		Discarded:        eventDiscarded,
		IsAcceptedOffer:  isAcceptedOffer,
		ERC721Transfers:  txLogs.ERC721Transfers,
		ERC1155Transfers: txLogs.ERC1155Transfers,
	}

	// send to formatting
	*queueEvents <- event

	// gbCache := cache.New()
	// gbCache.StoreEvent(event.Collection.ContractAddress, event.Collection.Name, event.TokenID, event.PriceWei.Uint64(), event.TxLogCount, event.Time, int64(eventType))
}

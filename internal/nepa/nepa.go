package nepa

import (
	"context"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/pusu"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

type (
	NodeID int

	NePa struct {
		// channels to receive new blocks or logs from the chain/subscriptions
		newHeads chan *types.Header
		newLogs  chan types.Log
		// workQueue chan any

		Transactions           chan *types.Transaction
		QueueTokenTransactions chan *totra.TokenTransaction

		gb *gloomberg.Gloomberg

		knownTransactions   map[common.Hash]bool
		knownTransactionsMu *sync.RWMutex
	}
)

func NewNePa(gb *gloomberg.Gloomberg, queueTokenTransactions chan *totra.TokenTransaction) *NePa {
	if queueTokenTransactions == nil {
		queueTokenTransactions = make(chan *totra.TokenTransaction, 10240)
	}

	// create new np
	np := &NePa{
		newHeads: make(chan *types.Header, 10240),
		newLogs:  make(chan types.Log, 10240),

		Transactions:           make(chan *types.Transaction, 10240),
		QueueTokenTransactions: queueTokenTransactions,

		knownTransactions:   make(map[common.Hash]bool, 0),
		knownTransactionsMu: &sync.RWMutex{},

		gb: gb,
	}

	return np
}

func (np *NePa) Run() {
	// handle received logs
	for workerID := 1; workerID <= viper.GetInt("server.workers.newLogHandler"); workerID++ {
		go np.newLogHandler()
	}

	// // if preferred (formerly 'local') nodes are available, we just subscribe to them
	// var ethNodes []*nodes.Node

	// if len(np.gb.Nodes.GetLocalNodes()) > 0 {
	// 	ethNodes = np.gb.Nodes.GetLocalNodes()
	// } else {
	// 	ethNodes = *np.gb.Nodes
	// }

	// subscribe
	subscribedTo, err := np.gb.ProviderPool.Subscribe(np.newLogs)
	if err != nil {
		gbl.Log.Fatalf("❌ subscribing to logs failed: %s", err)

		return
	}

	// logs := np.gb.ProviderPool.GetLogsByBlockNumber(17184457)
	// for _, log := range logs {
	//	np.newLogs <- log
	//}

	gbl.Log.Infof("✍️ subscribed to logs via %d nodes", subscribedTo)

	// for _, node := range ethNodes {
	// 	// subscribe to all logs with "Tranfer" or "TransferSingle" as first topic
	// 	if _, err := node.SubscribeToAllTransfers(np.newLogs); err != nil {
	// 		gbl.Log.Warnf("subscribe to topic TransferSingle via node %d failed: %s", node.NodeID, err)
	// 	} else {
	// 		gbl.Log.Infof("✍️ subscribed to all transfer topics via node %s", style.Bold(node.Name))
	// 	}
	// }

	//
	// pubsub subscribe to sales
	if viper.GetBool("pubsub.sales.subscribe") {
		gbl.Log.Infof("🚇 subscribing to sales via redis on channel %s", internal.PubSubChannelSales)

		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			go pusu.SubscribeToSales(np.gb, internal.PubSubChannelSales, np.QueueTokenTransactions)
		}
	}

	select {}
}

// newLogHandler handles new logs from an ethNode and fetches the complete tx for it.
func (np *NePa) newLogHandler() {
	gbl.Log.Debugf("🧱 starting newLogHandler")

	for log := range np.newLogs {
		// filter logs without topics
		if len(log.Topics) == 0 {
			continue
		}

		// filter logs without erc721/1155 transfer (len(topics) == 4)
		if len(log.Topics) != 4 {
			continue
		}

		// skip if we already processed this logs tx
		np.knownTransactionsMu.Lock()
		known, ok := np.knownTransactions[log.TxHash]
		np.knownTransactionsMu.Unlock()

		if known && ok {
			// we already know this transaction
			gbl.Log.Debugf("already known log/transaction: %s", style.BoldStyle.Render(log.TxHash.String()))

			continue
		}

		np.knownTransactionsMu.Lock()
		np.knownTransactions[log.TxHash] = true
		np.knownTransactionsMu.Unlock()

		//
		// get the transaction
		tx, err := np.gb.ProviderPool.TransactionByHash(context.Background(), log.TxHash)
		if err != nil {
			gbl.Log.Infof("❌ getting %s failed: %s", style.TerminalLink("https://etherscan.io/tx/"+log.TxHash.String(), "transaction"), err)

			continue
		} else if tx == nil {
			gbl.Log.Infof("❌ %s is nil", style.TerminalLink("https://etherscan.io/tx/"+log.TxHash.String(), "transaction"))

			continue
		}

		// filter contract creation transactions (`to` is the zeroAddress)
		if tx.To() == nil {
			continue
		}

		// np.Transactions <- tx

		//
		// get the receipt including the logs
		receipt, err := np.gb.ProviderPool.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			gbl.Log.Infof("❗️ error getting %s receipt: %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"), err)

			continue
		} else if receipt == nil {
			gbl.Log.Infof("❗️ %s receipt is nil", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"))

			continue
		}

		gbl.Log.Debugf("🧱 queue: %d | txLogFetcher - logs: %d", len(np.Transactions), len(receipt.Logs))

		//
		// create a TokenTransaction
		if ttx := totra.NewTokenTransaction(tx, receipt, np.gb.ProviderPool); ttx != nil && ttx.IsMovingNFTs() {
			np.QueueTokenTransactions <- ttx

			// publish ttx via redis
			if viper.GetBool("pubsub.sales.publish") {
				go pusu.Publish(np.gb, internal.PubSubChannelSales, ttx)
			}
		}

		np.gb.ProviderPool.LastLogReceivedAt = time.Now()
	}
}

// func (np *NePa) txHandler() {
// 	gbl.Log.Info("🧱 starting txLogFetcher worker")

// 	for tx := range np.Transactions {
// 		receipt, err := np.gb.Nodes.TransactionReceipt(context.Background(), tx.Hash())
// 		if err != nil {
// 			gbl.Log.Infof("❗️ error getting %s receipt: %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"), err)

// 			continue
// 		} else if receipt == nil {
// 			gbl.Log.Infof("❗️ %s receipt is nil", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"))

// 			continue
// 		}

// 		gbl.Log.Debugf("🧱 queue: %d | txLogFetcher - logs: %d", len(np.Transactions), len(receipt.Logs))

// 		//
// 		// create a TokenTransaction
// 		if ttx := totra.NewTokenTransaction(tx, receipt, np.gb.Nodes.GetRandomNode()); ttx != nil && ttx.IsMovingNFTs() {
// 			np.QueueTokenTransactions <- ttx

// 			// publish ttx via redis
// 			if viper.GetBool("pubsub.sales.publish") {
// 				go pusu.Publish(np.gb, internal.PubSubChannelSales, ttx)
// 			}
// 		}
// 	}
// }

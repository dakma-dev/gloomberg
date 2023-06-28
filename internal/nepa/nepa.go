package nepa

import (
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/chawago"
	chawagoModels "github.com/benleb/gloomberg/internal/chawago/models"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/pusu"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

type (
	NodeID int

	NePa struct {
		// channels to receive new blocks or logs from the chain/subscriptions
		newHeads        chan *types.Header
		newLogs         chan types.Log
		newTransactions chan *chawagoModels.TxWithLogs
		// workQueue chan any

		Transactions           chan *types.Transaction
		QueueTokenTransactions chan *totra.TokenTransaction

		gb *gloomberg.Gloomberg

		knownTransactions   map[common.Hash]bool
		knownTransactionsMu *sync.RWMutex
	}
)

func NewNePa(gb *gloomberg.Gloomberg) *NePa {
	// create new np
	np := &NePa{
		newHeads:        make(chan *types.Header, 10240),
		newLogs:         make(chan types.Log, 10240),
		newTransactions: make(chan *chawagoModels.TxWithLogs, 10240),

		Transactions: make(chan *types.Transaction, 10240),
		// QueueTokenTransactions: queueTokenTransactions,
		QueueTokenTransactions: gb.In.TokenTransactions,

		knownTransactions:   make(map[common.Hash]bool),
		knownTransactionsMu: &sync.RWMutex{},

		gb: gb,
	}

	return np
}

func (np *NePa) Run() {
	newLogs := make(chan types.Log, 10240)

	//
	// subscribe via websocket/rpc
	subscribedTo, err := np.gb.ProviderPool.Subscribe(newLogs)
	if err != nil {
		gbl.Log.Fatalf("❌ subscribing to logs failed: %s", err)

		return
	}

	np.newTransactions = chawago.GetTransactionsForLogs(np.gb, newLogs)

	// handle received transactions
	qTxsWithLogs := np.gb.SubscribeTxWithLogs()
	for workerID := 1; workerID <= viper.GetInt("server.workers.newLogHandler"); workerID++ {
		go np.newLogHandler(qTxsWithLogs)
	}

	gbl.Log.Debugf("✍️ subscribed to logs via %d nodes", subscribedTo)

	//
	// subscribe via redis pubsub
	if viper.GetBool("pubsub.sales.subscribe") {
		gbl.Log.Infof("🚇 subscribing to sales via redis on channel %s", internal.PubSubChannelSales)

		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			go pusu.SubscribeToSales(np.gb, internal.PubSubChannelSales, np.QueueTokenTransactions)
		}
	}

	select {}
}

// newLogHandler handles new logs from an ethNode and fetches the complete tx for it.
func (np *NePa) newLogHandler(qTxsWithLogs chan *chawagoModels.TxWithLogs) {
	gbl.Log.Debugf("🧱 starting newLogHandler")

	// for tx := range np.newTransactions {
	for tx := range qTxsWithLogs {
		log.Debugf("📝 %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), tx.Hash().String()))

		//
		// create a TokenTransaction
		if ttx := totra.NewTokenTransaction(tx.Transaction, tx.Receipt, np.gb.ProviderPool); ttx != nil && ttx.IsMovingNFTs() {
			np.QueueTokenTransactions <- ttx

			// publish ttx via redis
			if viper.GetBool("pubsub.sales.publish") {
				go pusu.Publish(np.gb, internal.PubSubChannelSales, ttx)
			}
		}

		np.gb.ProviderPool.LastLogReceivedAt = time.Now()
	}
}

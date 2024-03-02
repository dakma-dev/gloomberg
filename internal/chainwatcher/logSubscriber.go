package chainwatcher

import (
	"errors"

	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum"
)

type logSubscriber struct {
	cw *ChainWatcher
}

func NewLogSubscriber(cw *ChainWatcher) *logSubscriber {
	ls := &logSubscriber{cw: cw}

	return ls
}

// func (ls *logSubscriber) startNewLogsWorker() {
// 	knownTransactions := mapset.NewSet[string]()

// 	// handle received logs
// 	for workerID := 1; workerID <= viper.GetInt("chainwatcher.worker.rawLogs"); workerID++ {
// 		log.Debugf("starting rawLogs worker %d", workerID)

// 		go func(workerID int) {
// 			for rawLog := range ls.cw.newLogs {
// 				if knownTransactions.Contains(rawLog.TxHash.Hex()) {
// 					// we already know this transaction
// 					continue
// 				}

// 				knownTransactions.Add(rawLog.TxHash.Hex())

// 				// create a "full" transaction with logs
// 				txw, err := ls.cw.FetchTransactionWithReceipt(rawLog.TxHash)
// 				if err != nil {
// 					log.Warnf("❌ fetching transaction with receipt failed: %s", err)

// 					continue
// 				}

// 				ls.cw.newTransactions <- txw

// 				// TODO: update last log received to a per-node basis to detect stalled providers
// 				ls.cw.lastLogReceivedAt = time.Now()
// 			}
// 		}(workerID)
// 	}
// }

func (ls *logSubscriber) Subscribe() (chan *models.TxWithLogs, error) {
	subscriptions := make([]ethereum.Subscription, 0)

	log.Infof(" subscribe: %+v", subscriptions)

	for _, node := range ls.cw.Nodes {
		subscription, err := node.SubscribeToNFTTransfers(ls.cw.newLogs)
		if err != nil {
			log.Error("❌ %s: subscribing to logs failed: %s", style.BoldAlmostWhite(node.Name), err)

			continue
		}

		subscriptions = append(subscriptions, subscription)

		log.Infof("✍️ %s: subscribed to all transfer logs: %+v", style.BoldAlmostWhite(node.Name), subscription)
	}

	if len(subscriptions) == 0 {
		return nil, errors.New("❌ subscribing to logs failed on all nodes")
	}

	return ls.cw.newTransactions, nil
}

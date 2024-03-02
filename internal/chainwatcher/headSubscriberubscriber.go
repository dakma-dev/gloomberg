package chainwatcher

import (
	"context"
	"errors"

	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type headSubscriber struct {
	cw *ChainWatcher
}

func NewHeadSubscriber(cw *ChainWatcher) *headSubscriber {
	hs := &headSubscriber{cw: cw}

	go hs.startNewHeadWorker()

	return hs
}

func (bs *headSubscriber) Subscribe() (chan *models.TxWithLogs, error) {
	subscriptions := make([]ethereum.Subscription, 0)

	for _, node := range bs.cw.Nodes {
		subscription, err := node.SubscribeToHeads(bs.cw.newHeads)
		if err != nil {
			log.Error("❌ %s: subscribing to new blocks failed: %s", style.BoldAlmostWhite(node.Name), err)

			continue
		}

		subscriptions = append(subscriptions, subscription)

		log.Infof("✍️ %s: subscribed to new blocks: %+v", style.BoldAlmostWhite(node.Name), subscription)
	}

	if len(subscriptions) == 0 {
		return nil, errors.New("❌ subscribing to new blocks failed on all nodes")
	}

	return bs.cw.newTransactions, nil
}

func (bs *headSubscriber) startNewHeadWorker() {
	log.Info("starting new head worker")

	for rawHead := range bs.cw.newHeads {
		blockHash := rawHead.Hash()

		topics := [][]common.Hash{
			{common.HexToHash(string(topic.Transfer)), common.HexToHash(string(topic.TransferSingle))},
			{},
			{},
			{},
		}

		filterQuery := ethereum.FilterQuery{
			BlockHash: &blockHash,
			Topics:    topics,
		}

		blockLogs, err := bs.cw.Node().EthClient.FilterLogs(context.Background(), filterQuery)
		if err != nil {
			log.Warnf("❌ getting block %d failed: %s", rawHead.Number, err)

			continue
		}

		log.Printf("🧱 new block: %+v", blockLogs)

		for _, newLog := range blockLogs {
			bs.cw.newLogs <- newLog
			// txWithReceipt, err := bs.cw.FetchTransactionWithReceipt(newLog.TxHash)
			// if err != nil {
			// 	log.Warnf("❌ fetching transaction with receipt failed: %s", err)

			// 	continue
			// }

			// bs.cw.newTransactions <- txWithReceipt
		}

		// log.Printf("🧱 new block: %+v", blockReceipts)
		// bs.cw.newBlocks <- &models.Block{
		// 	Block:        block,
		// 	Transactions: make([]*models.TxWithLogs, 0),
		// }
	}
}

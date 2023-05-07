package chawago

import (
	"context"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var numWorkersRawLogs = 4

type TxWithLogs struct {
	*types.Transaction
	*types.Receipt
	Pending bool
}

type Topic int64

const (
	Topic0 Topic = iota
	Topic1
	Topic2
	Topic3
)

// getTxMessage is used to get the From field of a transaction.
func (t *TxWithLogs) Sender() *common.Address {
	sender, err := types.LatestSignerForChainID(t.ChainId()).Sender(t.Transaction)
	if err != nil {
		log.Warnf("could not get message for tx %s: %s", t.Hash().Hex(), err)

		return &common.Address{}
	}

	return &sender
}

func (t *TxWithLogs) CheckAddress(address common.Address) Topic {
	if t.To() != nil && *t.To() == address {
		log.Debugf("✅ %s is recipient of %s", style.TerminalLink("https://etherscan.io/address/"+address.String(), address.String()), style.TerminalLink("https://etherscan.io/tx/"+t.Hash().String(), "transaction"))

		return -1
	}

	if t.Sender() != nil && *t.Sender() == address {
		log.Debugf("✅ %s is sender of %s", style.TerminalLink("https://etherscan.io/address/"+address.String(), address.String()), style.TerminalLink("https://etherscan.io/tx/"+t.Hash().String(), "transaction"))

		return -1
	}

	if !t.Pending {
		for _, rawLog := range t.Logs {
			switch {
			case len(rawLog.Topics) >= 1 && rawLog.Topics[0] == address.Hash():
				log.Debugf("✅ %s found in topic 0 of %s", style.TerminalLink("https://etherscan.io/address/"+address.String(), address.String()), style.TerminalLink("https://etherscan.io/tx/"+t.Hash().String(), "transaction"))

				return Topic0

			case len(rawLog.Topics) >= 2 && rawLog.Topics[1] == address.Hash():
				log.Debugf("✅ %s found in topic 1 of %s", style.TerminalLink("https://etherscan.io/address/"+address.String(), address.String()), style.TerminalLink("https://etherscan.io/tx/"+t.Hash().String(), "transaction"))

				return Topic1

			case len(rawLog.Topics) >= 3 && rawLog.Topics[2] == address.Hash():
				log.Debugf("✅ %s found in topic 2 of %s", style.TerminalLink("https://etherscan.io/address/"+address.String(), address.String()), style.TerminalLink("https://etherscan.io/tx/"+t.Hash().String(), "transaction"))

				return Topic2

			case len(rawLog.Topics) >= 4 && rawLog.Topics[3] == address.Hash():
				log.Debugf("✅ %s found in topic 3 of %s", style.TerminalLink("https://etherscan.io/address/"+address.String(), address.String()), style.TerminalLink("https://etherscan.io/tx/"+t.Hash().String(), "transaction"))

				return Topic3
			}
		}
	}

	return -1
}

// GetTransactionsForLogs utilizes the providerPool to fetch the transaction & receipt for logs from qRawLogs.
// The transaction with the receipt is then sent to qTxsWithLogs.
func GetTransactionsForLogs(qRawLogs chan types.Log, qTxsWithLogs chan TxWithLogs, providerPool *provider.Pool) {
	knownTransactions := make(map[common.Hash]bool, 0)
	knownTransactionsMu := &sync.RWMutex{}

	// handle received logs
	for workerID := 1; workerID <= numWorkersRawLogs; workerID++ {
		log.Printf("starting rawLogs worker %d", workerID)

		go func() {
			for rawLog := range qRawLogs {
				// filter logs without topics
				if len(rawLog.Topics) == 0 {
					log.Printf("❕ log without topics: %s", style.BoldStyle.Render(rawLog.TxHash.String()))

					continue
				}

				// skip if we already processed this logs tx
				knownTransactionsMu.Lock()
				known, ok := knownTransactions[rawLog.TxHash]
				knownTransactionsMu.Unlock()

				if known && ok {
					// we already know this transaction
					log.Debugf("❕ already known log/transaction: %s", style.BoldStyle.Render(rawLog.TxHash.String()))

					continue
				}

				knownTransactionsMu.Lock()
				knownTransactions[rawLog.TxHash] = true
				knownTransactionsMu.Unlock()

				// fetch the full transaction this log belongs to
				tx, err := providerPool.TransactionByHash(context.Background(), rawLog.TxHash)
				if err != nil {
					log.Printf("❌ getting %s failed: %s", style.TerminalLink("https://etherscan.io/tx/"+rawLog.TxHash.String(), "transaction"), err)

					continue
				} else if tx == nil {
					log.Printf("❌ %s is nil", style.TerminalLink("https://etherscan.io/tx/"+rawLog.TxHash.String(), "transaction"))

					continue
				}

				// fetch the receipt to get all logs for this transaction
				receipt, err := providerPool.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Printf("❗️ error getting %s receipt: %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"), err)

					continue
				} else if receipt == nil {
					log.Printf("❗️ %s receipt is nil", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"))

					continue
				}

				// queue lengths
				log.Debugf("qLogs: %d  |  qTxsWithLogs: %d", len(qRawLogs), len(qTxsWithLogs))

				// output TxWithLogs
				qTxsWithLogs <- TxWithLogs{
					Transaction: tx,
					Receipt:     receipt,
				}

				// update last log received at timestamp to detect stalled providers
				providerPool.LastLogReceivedAt = time.Now()
			}
		}()
	}
}

// GetPendingTransactions utilizes the providerPool to fetch the transaction & receipt for logs from qRawLogs.
// The transaction with the receipt is then sent to qTxsWithLogs.
func GetPendingTransactions(qPendingTx chan *types.Transaction, qTxsWithLogs chan TxWithLogs, providerPool *provider.Pool) {
	knownTransactions := make(map[common.Hash]bool, 0)
	knownTransactionsMu := &sync.RWMutex{}

	// handle received logs
	for workerID := 1; workerID <= numWorkersRawLogs; workerID++ {
		log.Printf("starting pending tx worker %d", workerID)

		go func() {
			for pendingTx := range qPendingTx {
				// skip if we already processed this logs tx
				knownTransactionsMu.Lock()
				known, ok := knownTransactions[pendingTx.Hash()]
				knownTransactionsMu.Unlock()

				if known && ok {
					// we already know this transaction
					log.Debugf("❕ already known log/transaction: %s", style.BoldStyle.Render(pendingTx.Hash().String()))

					continue
				}

				knownTransactionsMu.Lock()
				knownTransactions[pendingTx.Hash()] = true
				knownTransactionsMu.Unlock()

				// queue lengths
				log.Debugf("qPendingTx: %d  |  qTxsWithLogs: %d", len(qPendingTx), len(qTxsWithLogs))

				// output TxWithLogs
				qTxsWithLogs <- TxWithLogs{
					Transaction: pendingTx,
					Pending:     true,
				}

				// update last log received at timestamp to detect stalled providers
				providerPool.LastLogReceivedAt = time.Now()
			}
		}()
	}
}

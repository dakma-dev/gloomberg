package chawago

import (
	"context"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/chawago/models"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var numWorkersRawLogs = 4

const (
	Topic0 models.Topic = iota
	Topic1
	Topic2
	Topic3
)

// GetTransactionsForLogs utilizes the providerPool to fetch the transaction & receipt for logs from qRawLogs.
// The transaction with the receipt is then sent to qTxsWithLogs.
func GetTransactionsForLogs(gb *gloomberg.Gloomberg, qRawLogs chan types.Log) chan *models.TxWithLogs {
	return GetTransactionsForLogsWithChannel(gb, qRawLogs, make(chan *models.TxWithLogs, 10240))
}

func GetTransactionsForLogsWithChannel(gb *gloomberg.Gloomberg, qRawLogs chan types.Log, qTxsWithLogs chan *models.TxWithLogs) chan *models.TxWithLogs {
	knownTransactions := make(map[common.Hash]bool, 0)
	knownTransactionsMu := &sync.RWMutex{}

	if qTxsWithLogs == nil {
		qTxsWithLogs = make(chan *models.TxWithLogs, 10240)
	}

	// handle received logs
	for workerID := 1; workerID <= numWorkersRawLogs; workerID++ {
		log.Debugf("starting rawLogs worker %d", workerID)

		go func() {
			for rawLog := range qRawLogs {
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

				log.Debugf("🪵 %#v", rawLog)

				// fetch the full transaction this log belongs to
				tx, err := gb.ProviderPool.TransactionByHash(context.Background(), rawLog.TxHash)
				if err != nil {
					log.Printf("❌ getting %s failed: %s", style.TerminalLink("https://etherscan.io/tx/"+rawLog.TxHash.String(), "transaction"), err)

					continue
				} else if tx == nil {
					log.Printf("❌ %s is nil", style.TerminalLink("https://etherscan.io/tx/"+rawLog.TxHash.String(), "transaction"))

					continue
				}

				log.Debugf("📝 %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), "transaction"))

				// fetch the receipt to get all logs for this transaction
				receipt, err := gb.ProviderPool.TransactionReceipt(context.Background(), tx.Hash())
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
				txWithLogs := &models.TxWithLogs{
					Transaction: tx,
					Receipt:     receipt,
				}

				// qTxsWithLogs <- txWithLogs

				gb.In.TxWithLogs <- txWithLogs

				// update last log received at timestamp to detect stalled providers
				gb.ProviderPool.LastLogReceivedAt = time.Now()
			}
		}()
	}

	return qTxsWithLogs
}

// GetPendingTransactions utilizes the providerPool to fetch the transaction & receipt for logs from qRawLogs.
// The transaction with the receipt is then sent to qTxsWithLogs.
func GetPendingTransactions(qPendingTx chan *types.Transaction, qTxsWithLogs chan *models.TxWithLogs, providerPool *provider.Pool) {
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
				qTxsWithLogs <- &models.TxWithLogs{
					Transaction: pendingTx,
					Pending:     true,
				}

				// update last log received at timestamp to detect stalled providers
				providerPool.LastLogReceivedAt = time.Now()
			}
		}()
	}
}

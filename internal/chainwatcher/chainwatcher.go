package chainwatcher

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

type ChainWatcher struct {
	Nodes models.Nodes

	subscriber Subscriber

	// channels
	newLogs         chan types.Log
	newHeads        chan *types.Header
	newTransactions chan *models.TxWithLogs

	lastLogReceivedAt time.Time `mapstructure:"-"`
}

func (cw *ChainWatcher) String() string {
	return fmt.Sprintf("ChainWatcher [Nodes: %v]", strings.Join(cw.Nodes.Names(), ", "))
}

func (cw *ChainWatcher) Node() *models.Node {
	if len(cw.Nodes) == 0 {
		log.Error("cw: no node available")

		return nil
	}

	return cw.Nodes[0]
}

// func New(config *viper.Viper) (cw *ChainWatcher) {.
func New(nodes models.Nodes) (cw *ChainWatcher) {
	cw = &ChainWatcher{
		Nodes: make([]*models.Node, 0),

		newLogs:         make(chan types.Log, 1024),
		newHeads:        make(chan *types.Header, 1024),
		newTransactions: make(chan *models.TxWithLogs, 1024),
	}

	cw.subscriber = NewHeadSubscriber(cw)

	for _, node := range nodes {
		if err := node.Connect(); err == nil {
			cw.Nodes = append(cw.Nodes, node)
		} else {
			log.Error("❌ %s: connecting to eth client failed: %s", style.BoldAlmostWhite(node.Name), err)

			continue
		}
	}

	if len(cw.Nodes) == 0 {
		log.Error("❌ connection to all nodes failed")

		return nil
	}

	// start the newLogs worker
	go cw.newLogsHandler()

	return cw
}

func (cw *ChainWatcher) SubscribeToNFTTransfers() (chan *models.TxWithLogs, error) {
	log.Info(" cw SubscribeToNFTTransfers")

	return cw.subscriber.Subscribe()
}

// func (cw *ChainWatcher) SubscribeToNFTTransfers() (chan *models.TxWithLogs, error) {
// 	subscriptions := make([]ethereum.Subscription, 0)

// 	for _, node := range cw.Nodes {
// 		subscription, err := node.SubscribeToNFTTransfers(cw.newLogs)
// 		if err != nil {
// 			log.Error("❌ %s: subscribing to logs failed: %s", style.BoldAlmostWhite(node.Name), err)

// 			continue
// 		}

// 		subscriptions = append(subscriptions, subscription)

// 		log.Infof("✍️ %s: subscribed to all transfer logs: %+v", style.BoldAlmostWhite(node.Name), subscription)
// 	}

// 	if len(subscriptions) == 0 {
// 		return nil, errors.New("❌ subscribing to logs failed on all nodes")
// 	}

// 	return cw.newTransactions, nil
// }

//
// gas
//

func (cw *ChainWatcher) FetchTransactionWithReceipt(txHash common.Hash) (txw *models.TxWithLogs, err error) {
	// fetch the full transaction this log belongs to
	tx, _, err := cw.Node().EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Warnf("❌ getting %s failed: %s", style.TerminalLink("https://etherscan.io/tx/"+txHash.Hex(), "transaction"), err)

		return nil, err
	} else if tx == nil {
		log.Warnf("❌ %s is nil", style.TerminalLink("https://etherscan.io/tx/"+txHash.Hex(), "transaction"))

		return nil, errors.New("transaction is nil")
	}

	// fetch the receipt to get all logs for this transaction
	receipt, err := cw.Node().EthClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Warnf("❗️ error getting %s receipt: %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), "transaction"), err)

		return nil, err
	} else if receipt == nil {
		log.Warnf("❗️ %s receipt is nil", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), "transaction"))

		return nil, errors.New("receipt is nil")
	}

	// create a "full" transaction with logs
	txw = &models.TxWithLogs{
		Transaction: tx,
		Receipt:     receipt,
	}

	return txw, nil
}

// GetGasInfo returns the current gas price and tip.
func (cw *ChainWatcher) GetGasInfo(ctx context.Context) (*models.GasInfo, error) {
	if cw.Node() == nil {
		return nil, errors.New("GetGasInfo: no node available")
	}

	if cw.Node().EthClient == nil {
		return nil, errors.New("GetGasInfo: no eth client available")
	}

	// header, err := cw.Node().EthClient.BlockByNumber(ctx, nil)
	header, err := cw.Node().EthClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("header: %+v", header)

	gi := &models.GasInfo{
		LastBlock:         header.Number.Uint64(),
		LastBlockGasLimit: header.GasLimit,
	}

	gi.GasPrice, err = cw.Node().EthClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Errorf("❌ %s: getting gas price failed: %s", style.BoldAlmostWhite(cw.Node().Name), err)

		return gi, err
	}

	gi.GasTip, err = cw.Node().EthClient.SuggestGasTipCap(ctx)
	if err != nil {
		log.Errorf("❌ %s: getting gas tip failed: %s", style.BoldAlmostWhite(cw.Node().Name), err)

		return gi, err
	}

	log.Infof("gi: %+v", gi)

	return gi, nil
}

// newLogsHandler fetches the tx and its receipt for each received log and sends it to the newTransactions channel.
func (cw *ChainWatcher) newLogsHandler() {
	knownTransactions := mapset.NewSet[string]()

	// handle received logs
	for workerID := 1; workerID <= viper.GetInt("chainwatcher.worker.rawLogs"); workerID++ {
		log.Debugf("starting newLogs worker %d", workerID)

		go func(workerID int) {
			for newLog := range cw.newLogs {
				if knownTransactions.Contains(newLog.TxHash.Hex()) {
					// we already know this transaction
					continue
				}

				knownTransactions.Add(newLog.TxHash.Hex())

				// fetch the "full" transaction + the receipt
				txWithReceipt, err := cw.FetchTransactionWithReceipt(newLog.TxHash)
				if err != nil {
					log.Warnf("❌ fetching transaction with receipt failed: %s", err)

					continue
				}

				cw.newTransactions <- txWithReceipt

				// TODO: update last log received to a per-node basis to detect stalled providers
				cw.lastLogReceivedAt = time.Now()
			}
		}(workerID)
	}
}

// func readNodeConfiguration(config *viper.Viper) models.Nodes {
// 	// var nodes models.Nodes = make([]*models.Node, 0)
// 	var nodes models.Nodes = make([]*models.Node, 0)

// 	log.Infof("👋 config: %+v", config.AllSettings())

// 	// ...or just get a single key
// 	err := config.UnmarshalKey("nodes", &nodes)
// 	if err != nil {
// 		log.Errorf("failed to unmarshal configuration: %v", err)
// 	}

// 	log.Infof("👋 nodes: %+v", nodes)

// 	return nodes
// }

func GetLocalABI(path string) string {
	abiFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}

func DecodeTransactionInputData(contractABI *abi.ABI, data []byte) {
	// The first 4 bytes of the t represent the ID of the method in the ABI
	// https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#function-selector
	methodSigData := data[:4]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Fatal(err)
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Method Name: %s\n", method.Name)
	log.Print("Method Inputs")
	pretty.Println(inputsMap)
}

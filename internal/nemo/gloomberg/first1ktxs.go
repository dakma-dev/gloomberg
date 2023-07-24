package gloomberg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FirstTxsTask struct {
	CollectionName  string
	ContractAddress common.Address
}

var (
	firstTxsWorkQueue = make(chan *FirstTxsTask, 1024)
	fetchedContracts  = mapset.NewSet[common.Address]()
	ignoredContracts  = mapset.NewSet[common.Address]()

	firstTxsCollectionCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_firsttxs_collection_count_total",
		Help: "No of collections with first txs downloaded.",
	},
	)
)

func JobFirstTxsForContract(params ...any) {
	if len(params) != 2 {
		return
	}

	collectionName, ok := params[0].(string)
	if !ok {
		return
	}

	contractAddress, ok := params[1].(common.Address)
	if !ok {
		return
	}

	GetFirstTxsForContract(collectionName, contractAddress)
}

// GetFirstTxsForContract fetches the first 1337 txs for a contract from etherscan and saves them to a json file.
func GetFirstTxsForContract(collectionName string, contractAddress common.Address) {
	if collectionName == "" {
		return
	}

	if fetchedContracts.Contains(contractAddress) {
		gbl.Log.Debugf("🤷‍♀️ %s already fetched", style.AlmostWhiteStyle.Render(collectionName))

		return
	}

	if ignoredContracts.Contains(contractAddress) {
		gbl.Log.Debugf("🤷‍♀️ %s already ignored", style.AlmostWhiteStyle.Render(collectionName))

		return
	}

	collectionData := FirstTxsTask{CollectionName: collectionName, ContractAddress: contractAddress}
	// firstTxsWorkQueue <- &collectionData

	err := fetchfirstTxsForContract(collectionData.CollectionName, collectionData.ContractAddress)
	if err != nil && os.IsExist(err) {
		log.Debugf("file for %s already exists: %+v", collectionData.CollectionName, err)
	} else if err != nil {
		gbl.Log.Debugf("failed to get firstTxs for %s: %s", collectionData.CollectionName, err)

		// ignore for this session
		ignoredContracts.Add(collectionData.ContractAddress)
	}

	fetchedContracts.Add(collectionData.ContractAddress)

	gbl.Log.Debugf("⏳ added %s to firstTxs queue", style.AlmostWhiteStyle.Render(collectionData.CollectionName))
}

func FirstTxsWorker() {
	interval := viper.GetDuration("etherscan.fetchInterval")
	fetchTicker := time.NewTicker(interval)

	log.Printf("👷 firstTxs worker started with interval %.0fs", interval.Seconds())

	for task := range firstTxsWorkQueue {
		err := fetchfirstTxsForContract(task.CollectionName, task.ContractAddress)
		if err != nil && os.IsExist(err) {
			log.Debugf("file for %s already exists: %+v", task.CollectionName, err)

			continue
		} else if err != nil {
			gbl.Log.Debugf("failed to get firstTxs for %s: %s", task.CollectionName, err)

			// ignore for this session
			ignoredContracts.Add(task.ContractAddress)

			continue
		}

		fetchedContracts.Add(task.ContractAddress)

		gbl.Log.Infof("📝 firstTxs fetched for %s (%d total | queue: %d)", task.CollectionName, fetchedContracts.Cardinality(), len(firstTxsWorkQueue))

		<-fetchTicker.C
	}
}

func fetchfirstTxsForContract(collectionName string, contractAddress common.Address) error {
	numTxsToFecth := int64(1337)

	// get first 1k txs for contract
	transactions, err := external.GetFirstTransactionsByContract(numTxsToFecth, contractAddress)
	if err != nil {
		return err
	}

	// preparedName := cases.Lower(language.Und).String(strings.ReplaceAll(collectionName, " ", "_")) + "_" + contractAddress.Hex() + ".json"

	// TODO: find a better way to prepare the filename
	// pretty sure there exists something already
	specialCharPattern := regexp.MustCompile(`[!\/\[\]\':;.,<>?@#$%^&*()_+|{}~]`)
	preparedName := specialCharPattern.ReplaceAllString(collectionName, "")

	preparedName = cases.Lower(language.Und).String(strings.ReplaceAll(preparedName, " ", "_")) + "_" + contractAddress.Hex() + ".json"

	txsFile := filepath.Join(internal.PathDegendata, "first_txs", preparedName)

	// create file
	file, err := os.OpenFile(txsFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644) //nolint:nosnakecase
	if err != nil {
		return err
	}

	defer file.Close()

	marshalled, err := json.Marshal(transactions)
	if err != nil {
		log.Error(fmt.Errorf("failed to marshal firstTxs for %s: %w", collectionName, err))

		return err
	}

	// write result to file
	_, err = file.Write(marshalled)
	if err != nil {
		log.Error(fmt.Errorf("failed to write firstTxs for %s to file: %w", collectionName, err))

		return err
	}

	file.Close()

	firstTxsCollectionCounter.Inc()

	return nil
}

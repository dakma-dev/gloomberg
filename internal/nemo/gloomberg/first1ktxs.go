package gloomberg

import (
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
)

var (
	fetchedContracts = mapset.NewSet[common.Address]()
	ignoredContracts = mapset.NewSet[common.Address]()

	_ = make(chan map[common.Address]string, 137)
)

func fetchFirst1K(gb *Gloomberg, contractInfo *degendb.ContractInfo, slug string, contractAddress common.Address) error {
	if fetchedContracts.Contains(contractAddress) {
		log.Debugf("🤷‍♀️ %s already fetched", style.AlmostWhiteStyle.Render(slug))

		return nil
	}

	if ignoredContracts.Contains(contractAddress) {
		log.Debugf("🤷‍♀️ %s already ignored", style.AlmostWhiteStyle.Render(slug))

		return nil
	}

	if gb.DegenDB.First1KTxAlreadyFetchedFor(contractAddress) {
		log.Infof("🤷‍♀️ %s already fetched", style.AlmostWhiteStyle.Render(slug))

		return nil
	}

	numTxsToFecth := int64(1337)

	// get first 1k txs for contract
	transactions, err := external.GetFirstTransactionsByContract(numTxsToFecth, contractAddress)
	if err != nil {
		return err
	} else if len(transactions) == 0 {
		return nil
	}

	fetchedContracts.Add(contractAddress)

	// log.Printf("📝 firstTxs: received %d txs for %s", len(transactions), slug)

	count := 0
	for _, tx := range transactions {
		if tx.From == "" {
			continue
		}

		fromAddr := common.HexToAddress(tx.From)
		toAddr := common.HexToAddress(tx.To)

		if fromAddr != internal.ZeroAddress || tx.TokenID == "" {
			continue
		}

		gb.DegenDB.SaveAddressFirst1KSlugs(toAddr, []degendb.Tag{degendb.Tag(slug)})

		time.Sleep(37 * time.Millisecond)

		count++
	}

	gb.DegenDB.SaveAddressFirst1KFetchedAt(contractAddress, contractInfo)

	if len(transactions) == count {
		log.Printf("📝 firstTxs: stored all %d addresses for %s", len(transactions), slug)
	} else {
		log.Printf("📝 firstTxs: stored %d/%d addresses for %s", count, len(transactions), slug)
	}

	return nil
}

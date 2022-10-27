package subscribe

import (
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func StreamListingsHandler(workerID int, ownCollections *collections.CollectionDB, queueListings *chan *models.ItemListedEvent, queueEvents *chan *collections.Event) {
	gbl.Log.Debugf("workerListingsHandler %d/%d started", workerID, viper.GetInt("workers.listings"))

	for event := range *queueListings {
		patternContractAddress := regexp.MustCompile(`^ethereum/(.*?)/(.*)$`)
		contractAddress := patternContractAddress.ReplaceAllString(event.Payload.Item.NftID, "$1")
		gbl.Log.Debugf("contractAddress: %+v", contractAddress)

		collection := ownCollections.Collections[common.HexToAddress(contractAddress)]
		if collection == nil {
			gbl.Log.Infof("collection not found: %s", event.Payload.Item.Metadata.Name)

			continue
		}

		patternTokenID := regexp.MustCompile(`^(.*?) ?#?(\d*)(/.*)?$`)
		tokenIDRaw := strings.TrimPrefix(patternTokenID.ReplaceAllString(event.Payload.Item.Metadata.Name, "$2"), "#")

		var tokenID int64

		tokenID, err := strconv.ParseInt(tokenIDRaw, 10, 64)
		if err != nil && tokenIDRaw != "" {
			gbl.Log.Infof("error parsing token ID | payload: %+v | tokenIDRaw: %s | error: %s", event.Payload, tokenIDRaw, err.Error())
		}

		priceWeiRaw, _, err := big.ParseFloat(event.Payload.BasePrice, 10, 64, big.ToNearestEven)
		if err != nil {
			gbl.Log.Errorf("workerListingsHandler: %s | %s", event.BaseStreamMessage.StreamEvent, err)
			continue
		}

		priceWei, _ := priceWeiRaw.Int(nil)
		priceEther, _ := nodes.WeiToEther(priceWei).Float64()

		event := &collections.Event{
			EventType:   collections.Listing,
			Collection:  ownCollections.Collections[common.HexToAddress(contractAddress)],
			TokenID:     big.NewInt(tokenID),
			Permalink:   event.Payload.Item.Permalink,
			TxItemCount: 1,
			PriceWei:    priceWei,
			PriceEther:  priceEther,
			Time:        time.Now(),
			From:        collections.User{},
			To: collections.User{
				Address:       common.HexToAddress(event.Payload.Maker.Address),
				OpenseaUserID: "",
			},
			WorkerID:   workerID,
			PrintEvent: true,
		}

		*queueEvents <- event

		atomic.AddUint64(&collection.Counters.Listings, 1)
	}
}

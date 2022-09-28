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
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func StreamListingsHandler(workerID int, ownCollections *collections.Collections, queueListings *chan *models.ItemListedEvent, queueEvents *chan *collections.Event) {
	gbl.Log.Infof("workerListingsHandler %d/%d started", workerID, viper.GetInt("workers.listings_handler"))

	for event := range *queueListings {
		patternContractAddress := regexp.MustCompile(`^ethereum/(.*?)/(.*)$`)
		contractAddress := patternContractAddress.ReplaceAllString(event.Payload.Item.NftID, "$1")
		gbl.Log.Debugf("contractAddress: %+v", contractAddress)

		collection := ownCollections.UserCollections[common.HexToAddress(contractAddress)]
		if collection == nil {
			gbl.Log.Infof("collection not found: %s", event.Payload.Item.Metadata.Name)

			continue
		}

		patternTokenID := regexp.MustCompile(`^(.*?) ?#?(\d*)(/.*)?$`)
		tokenIDRaw := patternTokenID.ReplaceAllString(event.Payload.Item.Metadata.Name, "$2")
		gbl.Log.Debugf("tokenIDRaw: %+v", tokenIDRaw)

		tokenIDRaw = strings.TrimPrefix(tokenIDRaw, "#")

		tokenID, err := strconv.ParseInt(tokenIDRaw, 10, 64)
		if err != nil {
			gbl.Log.Infof("error parsing tokenIDRaw to big.int: %s | %s | %s", event.Payload.Item.Metadata.Name, tokenIDRaw, err.Error())
		}

		priceWeiRaw, _, err := big.ParseFloat(event.Payload.BasePrice, 10, 64, big.ToNearestEven)
		if err != nil {
			gbl.Log.Errorf("workerListingsHandler: %s | %s", event.BaseStreamMessage.StreamEvent, err)
			continue
		}

		priceWei, _ := priceWeiRaw.Int(nil)

		event := &collections.Event{
			EventType:   collections.Listing,
			Collection:  ownCollections.UserCollections[common.HexToAddress(contractAddress)],
			TokenID:     big.NewInt(tokenID),
			Permalink:   event.Payload.Item.Permalink,
			TxItemCount: 1,
			PriceWei:    priceWei,
			Time:        time.Now(),
			From:        collections.User{},
			To: collections.User{
				Address:       common.HexToAddress(event.Payload.Maker.Address),
				OpenseaUserID: "",
			},
			WorkerID: workerID,
		}

		*queueEvents <- event

		atomic.AddUint64(&collection.Counters.Listings, 1)
	}
}

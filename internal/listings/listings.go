package listings

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/gloomberg"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/seaport"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func StreamListingsHandler(gb *gloomberg.Gloomberg, workerID int, queueListings *chan *models.ItemListedEvent, queueEvents *chan *collections.Event) {
	gbl.Log.Debugf("workerListingsHandler %d/%d started", workerID, viper.GetInt("workers.listings"))

	for event := range *queueListings {
		patternContractAddress := regexp.MustCompile(`^ethereum/(.*?)/(.*)$`)
		contractAddress := patternContractAddress.ReplaceAllString(event.Payload.Item.NftID, "$1")

		collection := gb.CollectionDB.Collections[common.HexToAddress(contractAddress)]
		if collection == nil {
			gbl.Log.Infof("collection not found: %s", event.Payload.Item.Metadata.Name)

			continue
		}

		// patternTokenID := regexp.MustCompile(`^(.*?) ?#?(\d*)(/.*)?$`)
		// tokenIDRaw := strings.TrimPrefix(patternTokenID.ReplaceAllString(event.Payload.Item.Metadata.Name, "$2"), "#")

		// var tokenID int64

		// if tokenIDRaw != "" {
		// 	tID, err := strconv.ParseInt(tokenIDRaw, 10, 64)
		// 	if err != nil {
		// 		gbl.Log.Infof("error parsing token ID | payload: %+v | tokenIDRaw: %s | error: %s", event.Payload, tokenIDRaw, err.Error())
		// 	} else {
		// 		tokenID = tID
		// 	}
		// } else {
		// 	gbl.Log.Infof("🤷‍♀️ no tokenID found in: %s", event.Payload.Item.Metadata.Name)
		// 	gbl.Log.Infof(dump.Format(event.Payload))
		// }

		// nftID is a string in the format <chain>/<contract>/<tokenID>
		nftID := strings.Split(event.Payload.Item.NftID, "/")
		if len(nftID) != 3 {
			gbl.Log.Infof("🤷‍♀️ error parsing nftID: %s", event.Payload.Item.NftID)
		}

		// get tokenID from nftID
		tokenID, _ := strconv.ParseInt(nftID[2], 10, 64)

		// parse price
		priceWeiRaw, _, err := big.ParseFloat(event.Payload.BasePrice, 10, 64, big.ToNearestEven)
		if err != nil {
			gbl.Log.Errorf("workerListingsHandler: %s | %s", event.BaseStreamMessage.StreamEvent, err)
			continue
		}

		priceWei, _ := priceWeiRaw.Int(nil)
		priceEther, _ := nodes.WeiToEther(priceWei).Float64()

		// get our floor price
		collectionFP := (*collection.FloorPrice).Value()

		if tokenID > 0 && collectionFP > 0.0 {

			gbl.Log.Debugf("%s | priceEther %6.3f < %6.3f collectionFP == %v ~ %6.3f", event.Payload.Item.Metadata.Name, priceEther, collectionFP, priceEther < collectionFP, priceEther/collectionFP)

			if priceEther < collectionFP {
				fmt.Printf("")
				gbl.Log.Debugf("%s | %6.3f < %6.3f\n", event.Payload.Item.Metadata.Name, priceEther, collectionFP)

				if priceEther < (collectionFP * 0.1) {
					fmt.Printf("%s BUY! FATFINGER!!! BUY! | %6.3f < %6.3f * 0.6  !!  FATFINGER!!!\n", event.Payload.Item.Metadata.Name, priceEther, collectionFP)

					go func() {
						gbl.Log.Infof("trying to get lisings for tokenID %d", tokenID)

						listings := opensea.GetListings(collection.ContractAddress, tokenID)
						if len(listings) > 0 {
							gbl.Log.Infof("listing found for %s", event.Payload.Item.Metadata.Name)

							tx, err := seaport.FulfillBasicOrder(gb, &listings[0], viper.GetString("buy.privateKey"))
							if err != nil {
								gbl.Log.Errorf("❌ error fulfilling order: %s", err.Error())
							} else {
								gbl.Log.Infof("✅ order fulfilled: %s", tx.Hash().Hex())
							}
						}
					}()
				}

				if priceEther < (collectionFP * 0.6) {
					fmt.Printf("%s FATFINGER!!!  | %6.3f < %6.3f * 0.6  FATFINGER!!!\n", event.Payload.Item.Metadata.Name, priceEther, collectionFP)
				}

				if priceEther < (collectionFP * 0.9) {
					fmt.Printf("%s hmmm...  | %6.3f < %6.3f * 0.9 ~ %6.3f\n", event.Payload.Item.Metadata.Name, priceEther, collectionFP, priceEther/collectionFP)
				}

				// set price of listing as the new fp
				(*collection.FloorPrice).Set(priceEther)
				go cache.StoreFloor(collection.ContractAddress, priceEther)
			}
		}

		event := &collections.Event{
			EventType:  collections.Listing,
			Collection: gb.CollectionDB.Collections[common.HexToAddress(contractAddress)],
			TokenID:    big.NewInt(tokenID),
			Permalink:  event.Payload.Item.Permalink,
			TxLogCount: 1,
			PriceWei:   priceWei,
			Time:       time.Now(),
			From:       collections.User{},
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

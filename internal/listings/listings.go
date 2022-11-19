package listings

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/gloomberg"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/seaport"
	"github.com/benleb/gloomberg/internal/style"
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

		// get current floor price
		collectionFP := (*collection.FloorPrice).Value()

		// check listing for base requirements to trigger auto-buy
		if viper.IsSet("buy.privateKey") && collectionFP > 0.0 && priceEther < collectionFP && tokenID > 0 {
			buyEvent(gb, event)
		}

		atomic.AddUint64(&collection.Counters.Listings, 1)
	}
}

func buyEvent(gb *gloomberg.Gloomberg, event *collections.Event) {
	// tokenName consists of collection name and tokenID
	tokenName := event.Collection.Name + " #" + event.TokenID.String()

	priceEther, _ := nodes.WeiToEther(event.PriceWei).Float64()

	// get current floor price
	collectionFP := (*event.Collection.FloorPrice).Value()

	// listing price / floor price as percentage
	listingToFloorPriceRatio := priceEther / collectionFP
	fpRatioDifference := int((listingToFloorPriceRatio * 100) - 100)

	privateKey := viper.GetString("buy.privateKey")
	if privateKey == "" {
		gbl.Log.Infof("🤷‍♀️ %s| no private key set, skipping auto-buy", tokenName)
		return
	}

	// WEN...??
	timeNow := style.GrayStyle.Copy().Faint(true).Render(time.Now().Format("15:04:05"))

	divider := style.Sharrow.Copy().Foreground(style.DarkGray).String()

	// build the line to be displayed
	out := strings.Builder{}
	out.WriteString("  |" + timeNow)
	out.WriteString(" 🛍️ " + divider)

	var listingToFloorPriceThreshold float64
	if viper.IsSet("buy.listingToFloorPriceThreshold") {
		listingToFloorPriceThreshold = viper.GetFloat64("buy.listingToFloorPriceThreshold")

		if listingToFloorPriceThreshold < 0.0 || listingToFloorPriceThreshold > 1.0 {
			gbl.Log.Infof("🤷‍♀️ %s| invalid listingToFloorPriceThreshold (%.3f) value, skipping auto-buy", tokenName, listingToFloorPriceThreshold)
			return
		}
	} else {
		gbl.Log.Infof("🤷‍♀️ %s| no listingToFloorPriceThreshold value set, skipping auto-buy", tokenName)
		return
	}

	if listingToFloorPriceRatio > listingToFloorPriceThreshold {
		gbl.Log.Infof("🤷‍♀️ %s| listingToFloorPriceRatio (%.3f) > listingToFloorPriceThreshold (%.2f), skipping auto-buy", tokenName, listingToFloorPriceRatio, listingToFloorPriceThreshold)
		// fmt.Printf("🤷‍♀️ %s| listing: %.3f | floor: %.3f ~ ratio: %s > threshold: %.3f | skipping auto-buy", tokenName, priceEther, collectionFP, style.TrendRedStyle.Render(fmt.Sprintf("%.3f", listingToFloorPriceRatio)), listingToFloorPriceThreshold)

		out.WriteString(" " + style.TrendLightRedStyle.Render(fmt.Sprintf("%+d%%", fpRatioDifference)))
	} else {
		out.WriteString(" " + style.TrendGreenStyle.Render(fmt.Sprintf("%+d%%", fpRatioDifference)))
	}

	out.WriteString("  " + tokenName + " " + divider)

	// out.WriteString(" " + fmt.Sprintf("listing: %.3f | floor: %.3f ~ ratio: %.3f | threshold: %.2f", priceEther, collectionFP, listingToFloorPriceRatio, listingToFloorPriceThreshold))

	// fmt.Println(out.String())
	// gbl.Log.Info(out.String())

	if event.Collection.Counters.Sales > 15 && listingToFloorPriceRatio <= listingToFloorPriceThreshold {
		gbl.Log.Infof("trying to get lisings for tokenID %d", event.TokenID)

		// get listing details needed to fulfill order
		if listings := opensea.GetListings(event.Collection.ContractAddress, event.TokenID.Int64()); len(listings) > 0 {
			gbl.Log.Infof("listing found for %s", tokenName)

			tx, err := seaport.FulfillBasicOrder(gb, &listings[0], viper.GetString("buy.privateKey"))
			if err != nil {
				gbl.Log.Warnf("❌ %s| error fulfilling order: %s", tokenName, err.Error())
			} else {
				gbl.Log.Infof("✅ %s| successfully purchased 🛍️ | %s", tokenName, tx.Hash().String())
			}

			out.WriteString(" " + err.Error())

			fmt.Println(out.String())
		}
	}
}

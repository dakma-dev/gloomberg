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
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
		// priceEther, _ := nodes.WeiToEther(priceWei).Float64()

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
		gbl.Log.Debugf("🛍️ len(gb.BuyRules): %d | collectionFP: %f | tokenID: %d", len(gb.BuyRules), collectionFP, tokenID)

		if len(gb.BuyRules) > 0 && collectionFP > 0.0 && tokenID > 0 {
			gbl.Log.Debug("🛍️ checking listing for auto-buy")
			checkBuyRulesForEvent(gb, event)
		}

		atomic.AddUint64(&collection.Counters.Listings, 1)
	}
}

func checkBuyRulesForEvent(gb *gloomberg.Gloomberg, event *collections.Event) {
	// tokenName consists of collection name and tokenID
	tokenName := event.Collection.Name + " #" + event.TokenID.String()

	minSales := 1
	minListings := 3

	// filter events with non-accurate data
	if event.Collection.Counters.Sales < 1 || event.Collection.Counters.Listings < 3 {
		sales := fmt.Sprintf("(%d/%d)", event.Collection.Counters.Sales, minSales)
		listings := fmt.Sprintf("(%d/%d)", event.Collection.Counters.Listings, minListings)
		gbl.Log.Infof("🤷‍♀️ %s| too less sales %s and/or listings %s to calculate accurate floor price", tokenName, sales, listings)

		return
	}

	priceEther, _ := nodes.WeiToEther(event.PriceWei).Float64()

	// get current floor price
	collectionFP := (*event.Collection.FloorPrice).Value()

	// listing price / floor price as percentage
	listingToFloorPriceRatio := priceEther / collectionFP
	fpRatioDifference := int((listingToFloorPriceRatio * 100) - 100)

	// WEN...??
	timeNow := style.GrayStyle.Copy().Faint(true).Render(time.Now().Format("15:04:05"))
	divider := style.Sharrow.Copy().Foreground(style.DarkGray).String()

	// build the line to be displayed
	out := strings.Builder{}
	out.WriteString("  |" + timeNow)
	out.WriteString(" 🛍️ " + divider)

	for _, rule := range gb.BuyRules {
		if rule.ContractAddress != event.Collection.ContractAddress && rule.ContractAddress != utils.ZeroAddress {
			gbl.Log.Debugf("🤷‍♀️ %s| no rule matching contract address, skipping auto-buy", tokenName)
			continue
		}

		if listingToFloorPriceRatio > rule.Threshold {
			gbl.Log.Infof("🤷‍♀️ %s| rule matched contract address but price above threshold (%.3f > %.3f), skipping auto-buy", tokenName, listingToFloorPriceRatio, rule.Threshold)
			out.WriteString(" " + style.TrendLightRedStyle.Render(fmt.Sprintf("%+d%%", fpRatioDifference)))

			continue
		}

		out.WriteString(" " + style.TrendGreenStyle.Render(fmt.Sprintf("%+d%%", fpRatioDifference)))

		gbl.Log.Debugf("trying to get lisings for tokenID %d", event.TokenID)

		// get listing details needed to fulfill order
		if listings := opensea.GetListings(event.Collection.ContractAddress, event.TokenID.Int64()); len(listings) > 0 {
			gbl.Log.Infof("listing found for %s", tokenName)

			if tx, err := buy(gb, &listings[0], rule.PrivateKey, tokenName); err != nil {
				out.WriteString(" " + err.Error())
			} else {
				fmt.Printf("tx: %+v\n", tx)
			}

			fmt.Println(out.String())

			return
		}
	}
}

func buy(gb *gloomberg.Gloomberg, order *models.SeaportOrder, privateKey string, tokenName string) (*types.Transaction, error) {
	tx, err := seaport.FulfillBasicOrder(gb, order, privateKey)
	if err != nil {
		gbl.Log.Warnf("❌ %s| error fulfilling order: %s", tokenName, err.Error())
	} else {
		gbl.Log.Infof("✅ %s| successfully purchased 🛍️ | %s", tokenName, tx.Hash().String())
	}

	return tx, err
}

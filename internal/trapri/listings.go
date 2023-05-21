package trapri

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/ethereum/go-ethereum/common"
)

func FormatListing(gb *gloomberg.Gloomberg, event *osmodels.ItemListedEvent, queueTokenTransactions chan *totra.TokenTransaction) {
	// nftID is a string in the format <chain>/<contract>/<tokenID>
	nftID := strings.Split(event.Payload.Item.NftID, "/")
	if len(nftID) != 3 {
		gbl.Log.Infof("🤷‍♀️ error parsing nftID: %s", event.Payload.Item.NftID)
	}

	contractAddress := common.HexToAddress(nftID[1])

	// get tokenID from nftID
	tokenID, _ := strconv.ParseInt(nftID[2], 10, 64)

	// seller address
	sellerAddress := common.HexToAddress(event.Payload.Maker.Address)

	// parse price
	priceWeiRaw, _, err := big.ParseFloat(event.Payload.BasePrice, 10, 64, big.ToNearestEven)
	if err != nil {
		gbl.Log.Errorf("❌ error parsing price: %s", err.Error())

		return
	}

	priceWei, _ := priceWeiRaw.Int(nil)
	price := price.NewPrice(priceWei)

	// collection
	collection := tokencollections.GetCollection(gb, contractAddress, tokenID)

	// counting for salira and more...
	collection.AddListing(uint64(event.Payload.Quantity))

	var receivedAt time.Time
	if sentAt, err := time.Parse(time.RFC3339, event.Payload.EventTimestamp); err == nil {
		receivedAt = sentAt
	} else {
		gbl.Log.Debugf("❌ failed to parse sentAt: %s | %+v | %+v", err, event.Payload.EventTimestamp, event.Payload.ListingDate)
		receivedAt = time.Now()
	}

	itemName := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]
	//
	// create a TokenTransaction
	ttxListing := &totra.TokenTransaction{
		Tx:          nil,
		TxReceipt:   nil,
		From:        sellerAddress,
		AmountPaid:  price.Wei(),
		TotalTokens: int64(event.Payload.Quantity),
		Marketplace: &marketplace.OpenSea,
		Action:      totra.Listing,
		ReceivedAt:  receivedAt,
		DoNotPrint:  false,
		Transfers: []*totra.TokenTransfer{
			{
				From:         marketplace.OpenSea.ContractAddress(),
				To:           sellerAddress,
				AmountTokens: big.NewInt(int64(event.Payload.Quantity)),
				Token: &token.Token{
					Address: contractAddress,
					ID:      big.NewInt(tokenID),
					Name:    itemName,
				},
			},
		},
	}

	gbl.Log.Debugf("%s: %+v | %+v", event.StreamEvent, ttxListing, event)

	// format and print
	queueTokenTransactions <- ttxListing

	// 	// // publish ttx via redis
	// 	// if viper.GetBool("pubsub.sales.publish") {
	// 	// 	go pusu.Publish(gb, internal.PubSubChannelSales, ttx)
	// 	// }
	// }
}

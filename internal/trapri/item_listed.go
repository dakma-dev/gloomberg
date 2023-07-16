package trapri

import (
	"math/big"
	"strings"

	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
)

func HandleItemListed(gb *gloomberg.Gloomberg, event *models.ItemListed) {
	nftID := event.Payload.Item.NftID

	// seller address
	sellerAddress := event.Payload.Maker.Address

	// token name without id
	var itemName string
	if name := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]; name != "" {
		itemName = name
	} else {
		itemName = event.Payload.Item.Metadata.Name
	}

	// create a TokenTransaction
	ttxListing := &totra.TokenTransaction{
		Tx:          nil,
		TxReceipt:   nil,
		From:        sellerAddress,
		AmountPaid:  event.Payload.GetPrice().Wei(),
		TotalTokens: int64(event.Payload.Quantity),
		Marketplace: &marketplace.OpenSea,
		Action:      totra.Listing,
		ReceivedAt:  event.Payload.EventTimestamp,
		DoNotPrint:  false,
		Transfers: []*totra.TokenTransfer{
			{
				From:         marketplace.OpenSea.ContractAddress(),
				To:           sellerAddress,
				AmountTokens: big.NewInt(int64(event.Payload.Quantity)),
				Token: &token.Token{
					Address: nftID.ContractAddress(),
					ID:      big.NewInt(nftID.TokenID().Int64()),
					Name:    itemName,
				},
			},
		},
	}

	// format and print
	gb.In.TokenTransactions <- ttxListing

	// collection
	collection := tokencollections.GetCollection(gb, nftID.ContractAddress(), nftID.TokenID().Int64())
	// counting for salira and more...
	collection.AddListing(uint64(event.Payload.Quantity))

	// // highlight "rare" lawless listings
	// if contractAddress == common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790") {
	// 	tokenName := event.GetTokenName()

	// 	switch {
	// 	case strings.Contains(tokenName, "-qf"):
	// 		tokenName = strings.Replace(tokenName, "-qf", style.PinkBoldStyle.Render("-qf * * * "), 1)
	// 	case strings.Contains(tokenName, "-rq"):
	// 		tokenName = strings.Replace(tokenName, "-rq", style.PinkBoldStyle.Render("-rq"), 1)
	// 	case strings.Contains(tokenName, "-pq"):
	// 		tokenName = strings.Replace(tokenName, "-pq", style.Bold("-pq"), 1)
	// 	case strings.Contains(tokenName, "-qp"):
	// 		tokenName = strings.Replace(tokenName, "-qp", style.Bold("-qp"), 1)
	// 	case strings.Contains(tokenName, "-qr"):
	// 		tokenName = strings.Replace(tokenName, "-qr", style.Bold("-qr"), 1)

	// 	default:
	// 		log.Debugf("lawless listing but common token: %s", tokenName)

	// 		return
	// 	}

	// 	osLink := style.TerminalLink(event.GetPermalink(), event.GetPermalink())

	// 	go notify.SendMessageViaTelegram(fmt.Sprintf("lawless listing: %s \n price: %s  url: %s", tokenName, fmt.Sprintf("%5.3f", price.Ether()), osLink), viper.GetInt64("notifications.manifold.dakma"), "", 0, nil)

	// 	highlightMessage := strings.Builder{}
	// 	highlightMessage.WriteString("\n")
	// 	highlightMessage.WriteString(fmt.Sprintf("  lawless %s | %5.3fΞ | %s\n", tokenName, price.Ether(), osLink))
	// 	highlightMessage.WriteString("\n")

	// 	fmt.Println(highlightMessage.String()) //nolint:forbidigo
	// }
}

// func fetchFloorPrice(address common.Address, collectionSlug string) (float64, error) {
// 	gbl.Log.Debugf("requesting floor from OpenSea...")
// 	if osCollectionStats := opensea.GetCollectionStats(collectionSlug); osCollectionStats != nil && osCollectionStats.FloorPrice > 0.0 {
// 		return osCollectionStats.FloorPrice, nil
// 	}

// 	gbl.Log.Debugf("requesting floor from Alchemy...")
// 	if alchemyCollectionStats := external.GetFloorPriceFromAlchemy(address.Hex()); alchemyCollectionStats != nil && alchemyCollectionStats.Opensea.FloorPrice > 0.0 {
// 		return alchemyCollectionStats.Opensea.FloorPrice, nil
// 	}

// 	return 0.0, errors.New("no floor price found")
// }

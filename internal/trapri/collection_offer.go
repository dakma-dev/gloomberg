package trapri

import (
	"math/big"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

func HandleCollectionOffer(gb *gloomberg.Gloomberg, event *models.CollectionOffer) {
	contractAddress := common.HexToAddress(event.Payload.Payload.AssetContractCriteria.Address.Hex())

	// seller address
	sellerAddress := event.Payload.Payload.Maker.Address

	// parse tokenPrice
	var tokenPrice *price.Price
	if event.Payload.Payload.BasePrice != nil {
		tokenPrice = price.NewPrice(event.Payload.Payload.BasePrice)
	} else {
		tokenPrice = price.NewPrice(big.NewInt(0))

		gbl.Log.Warnf("🤷‍♀️ error parsing tokenPrice: %+v", event.Payload.Payload.BasePrice)
	}

	collectionTokens := gb.OwnWallets.GetCollectionTokens(contractAddress)

	// create a ttxCollectionOffer for each token of the collection...^^
	for _, collectionToken := range collectionTokens {
		//
		// create a TokenTransaction
		ttxCollectionOffer := &totra.TokenTransaction{
			Tx:          nil,
			TxReceipt:   nil,
			From:        sellerAddress,
			AmountPaid:  tokenPrice.Wei(),
			TotalTokens: int64(event.Payload.Payload.Quantity),
			Marketplace: &marketplace.OpenSea,
			Action:      degendb.CollectionOffer,
			ReceivedAt:  event.Payload.Payload.EventTimestamp,
			DoNotPrint:  false,
			Transfers: []*totra.TokenTransfer{
				{
					From:         marketplace.OpenSea.ContractAddress(),
					To:           sellerAddress,
					AmountTokens: big.NewInt(int64(event.Payload.Payload.Quantity)),
					Token: &token.Token{
						Address: contractAddress,
						ID:      big.NewInt(collectionToken.ID.Int64()),
						Name:    event.Payload.Payload.Collection.Slug,
					},
				},
			},
		}

		// format and print
		gb.In.TokenTransactions <- ttxCollectionOffer

		log.Printf("  🦄 collection offer: %s | %s | %s", contractAddress.String(), sellerAddress.String(), tokenPrice.String())
	}
}

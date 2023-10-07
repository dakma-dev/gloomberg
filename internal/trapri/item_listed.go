package trapri

import (
	"math/big"
	"strings"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
)

func HandleItemListed(gb *gloomberg.Gloomberg, event *models.ItemListed) {
	nftID := event.Payload.Item.NftID

	contractAddress := nftID.ContractAddress()

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
		Action:      degendb.Listing,
		ReceivedAt:  event.Payload.EventTimestamp,
		DoNotPrint:  false,
		Transfers: []*totra.TokenTransfer{
			{
				From:         marketplace.OpenSea.ContractAddress(),
				To:           sellerAddress,
				AmountTokens: big.NewInt(int64(event.Payload.Quantity)),
				Token: &token.Token{
					Address: contractAddress,
					ID:      big.NewInt(nftID.TokenID().Int64()),
					Name:    itemName,
				},
			},
		},
	}

	// format and print in stream
	gb.In.TokenTransactions <- ttxListing

	// // 💄 style
	// primaryColor, _ := style.GenerateAddressColors(&contractAddress)
	// collectionStyle := lipgloss.NewStyle().Foreground(primaryColor)
	// fmtItemName := collectionStyle.Bold(true).Render(event.Payload.Item.Name)
	// LogOpenSeaEvent(degendb.GetEventType(event.EventType), &contractAddress, price.NewPrice(event.Payload.BasePrice), fmtItemName, &event.Payload.Maker.Address)

	// collection
	if collection := tokencollections.GetCollection(gb, contractAddress, nftID.TokenID().Int64()); collection != nil {
		// counting for salira and more...
		collection.AddListing(uint64(event.Payload.Quantity))

		return
	}
}

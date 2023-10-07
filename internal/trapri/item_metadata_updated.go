package trapri

import (
	"math/big"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
)

func HandleItemMetdadataUpdated(gb *gloomberg.Gloomberg, event *models.ItemMetadataUpdated) {
	nftID := event.Payload.Item.NftID
	contractAddress := nftID.ContractAddress()

	// seller address
	sellerAddress := internal.ZeroAddress

	// token name without id
	var itemName string
	if name := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]; name != "" {
		itemName = name
	} else {
		itemName = event.Payload.Item.Metadata.Name
	}

	// create a TokenTransaction
	ttxMetadataUpdated := &totra.TokenTransaction{
		Tx:          nil,
		TxReceipt:   nil,
		From:        sellerAddress,
		AmountPaid:  big.NewInt(0),
		TotalTokens: int64(1),
		Marketplace: &marketplace.OpenSea,
		Action:      degendb.MetadataUpdated,
		ReceivedAt:  event.Payload.CreatedDate,
		DoNotPrint:  false,
		Transfers: []*totra.TokenTransfer{
			{
				From:         marketplace.OpenSea.ContractAddress(),
				To:           sellerAddress,
				AmountTokens: big.NewInt(int64(1)),
				Token: &token.Token{
					Address: contractAddress,
					ID:      big.NewInt(nftID.TokenID().Int64()),
					Name:    itemName,
				},
			},
		},
	}

	// format and print
	gb.In.TokenTransactions <- ttxMetadataUpdated

	// // 💄 style
	// primaryColor, _ := style.GenerateAddressColors(&contractAddress)
	// collectionStyle := lipgloss.NewStyle().Foreground(primaryColor)
	// fmtItemName := collectionStyle.Bold(true).Render(event.Payload.Item.Name)
	// LogOpenSeaEvent(degendb.GetEventType(event.Event), &contractAddress, price.NewPrice(big.NewInt(0)), fmtItemName, &sellerAddress)
}

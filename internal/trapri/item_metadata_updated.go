package trapri

// import (
// 	"math/big"
// 	"strings"

// 	"github.com/benleb/gloomberg/internal/degendb"
// 	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
// 	"github.com/benleb/gloomberg/internal/nemo/marketplace"
// 	"github.com/benleb/gloomberg/internal/nemo/token"
// 	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
// 	"github.com/benleb/gloomberg/internal/nemo/totra"
// 	"github.com/benleb/gloomberg/internal/seawa/models"
// )

// func HandleItemMetdadataUpdated(gb *gloomberg.Gloomberg, event *models.ItemMetadataUpdated) {
// 	nftID := event.Payload.Item.NftID

// 	// seller address
// 	sellerAddress := event.Payload.Maker.Address

// 	// token name without id
// 	var itemName string
// 	if name := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]; name != "" {
// 		itemName = name
// 	} else {
// 		itemName = event.Payload.Item.Metadata.Name
// 	}

// 	// create a TokenTransaction
// 	ttxListing := &totra.TokenTransaction{
// 		Tx:          nil,
// 		TxReceipt:   nil,
// 		From:        sellerAddress,
// 		AmountPaid:  big.NewInt(0),
// 		TotalTokens: int64(event.Payload.Quantity),
// 		Marketplace: &marketplace.OpenSea,
// 		Action:      degendb.MetadataUpdate,
// 		ReceivedAt:  event.Payload.CreatedDate,
// 		DoNotPrint:  false,
// 		Transfers: []*totra.TokenTransfer{
// 			{
// 				From:         marketplace.OpenSea.ContractAddress(),
// 				To:           sellerAddress,
// 				AmountTokens: big.NewInt(int64(event.Payload.Quantity)),
// 				Token: &token.Token{
// 					Address: nftID.ContractAddress(),
// 					ID:      big.NewInt(nftID.TokenID().Int64()),
// 					Name:    itemName,
// 				},
// 			},
// 		},
// 	}

// 	// format and print
// 	gb.In.TokenTransactions <- ttxListing

// 	// collection
// 	collkection := tokencollecthions.GetCollection(gb, nftID.ContractAddress(), nftID.TokenID().Int64())
// 	// counting for salira and more...
// 	collection.AddListing(uint64(event.Payload.Quantity))
// }
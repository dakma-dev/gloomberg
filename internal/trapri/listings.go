package trapri

// import (
// 	"fmt"
// 	"math/big"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/benleb/gloomberg/internal/degendb"
// 	"github.com/benleb/gloomberg/internal/gbl"
// 	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
// 	"github.com/benleb/gloomberg/internal/nemo/marketplace"
// 	"github.com/benleb/gloomberg/internal/nemo/osmodels"
// 	"github.com/benleb/gloomberg/internal/nemo/price"
// 	"github.com/benleb/gloomberg/internal/nemo/token"
// 	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
// 	"github.com/benleb/gloomberg/internal/nemo/totra"
// 	"github.com/benleb/gloomberg/internal/notify"
// 	"github.com/benleb/gloomberg/internal/style"
// 	"github.com/charmbracelet/log"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/spf13/viper"
// )

// func FormatListing(gb *gloomberg.Gloomberg, event *osmodels.ItemListedEvent) {
// 	// nftID is a string in the format <chain>/<contract>/<tokenID>
// 	nftID := strings.Split(event.Payload.Item.NftID, "/")
// 	if len(nftID) != 3 {
// 		gbl.Log.Infof("🤷‍♀️ error parsing nftID: %s", event.Payload.Item.NftID)
// 	}

// 	contractAddress := common.HexToAddress(nftID[1])

// 	// get tokenID from nftID
// 	tokenID, _ := strconv.ParseInt(nftID[2], 10, 64)

// 	// seller address
// 	sellerAddress := common.HexToAddress(event.Payload.Maker.Address)

// 	// parse amountPaid
// 	priceWeiRaw, _, err := big.ParseFloat(event.Payload.BasePrice.String(), 10, 64, big.ToNearestEven)
// 	if err != nil {
// 		gbl.Log.Errorf("❌ error parsing amountPaid: %s", err.Error())

// 		return
// 	}

// 	priceWei, _ := priceWeiRaw.Int(nil)
// 	amountPaid := price.NewPrice(priceWei)

// 	// collection
// 	collection := tokencollections.GetCollection(gb, contractAddress, tokenID)

// 	// counting for salira and more...
// 	collection.AddListing(uint64(event.Payload.Quantity))

// 	var receivedAt time.Time
// 	if sentAt, err := time.Parse(time.RFC3339, event.Payload.EventTimestamp); err == nil {
// 		receivedAt = sentAt
// 	} else {
// 		gbl.Log.Debugf("❌ failed to parse sentAt: %s | %+v | %+v", err, event.Payload.EventTimestamp, event.Payload.ListingDate)
// 		receivedAt = time.Now()
// 	}

// 	itemName := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]

// 	//
// 	// create a TokenTransaction
// 	ttxListing := &totra.TokenTransaction{
// 		Tx:          nil,
// 		TxReceipt:   nil,
// 		From:        sellerAddress,
// 		AmountPaid:  amountPaid.Wei(),
// 		TotalTokens: int64(event.Payload.Quantity),
// 		Marketplace: &marketplace.OpenSea,
// 		Action:      degendb.Listing,
// 		ReceivedAt:  receivedAt,
// 		DoNotPrint:  false,
// 		Transfers: []*totra.TokenTransfer{
// 			{
// 				From:         marketplace.OpenSea.ContractAddress(),
// 				To:           sellerAddress,
// 				AmountTokens: big.NewInt(int64(event.Payload.Quantity)),
// 				Token: &token.Token{
// 					Address: contractAddress,
// 					ID:      big.NewInt(tokenID),
// 					Name:    itemName,
// 				},
// 			},
// 		},
// 	}

// 	gbl.Log.Debugf("%s: %+v | %+v", event.StreamEvent, ttxListing, event)

// 	// format and print
// 	gb.In.TokenTransactions <- ttxListing

// 	// highlight "rare" lawless listings
// 	if contractAddress == common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790") {
// 		tokenName := event.Payload.Item.Metadata.Name

// 		switch {
// 		case strings.Contains(tokenName, "-qf"):
// 			tokenName = strings.Replace(tokenName, "-qf", style.PinkBoldStyle.Render("-qf * * * "), 1)
// 		case strings.Contains(tokenName, "-rq"):
// 			tokenName = strings.Replace(tokenName, "-rq", style.PinkBoldStyle.Render("-rq"), 1)
// 		case strings.Contains(tokenName, "-pq"):
// 			tokenName = strings.Replace(tokenName, "-pq", style.Bold("-pq"), 1)
// 		case strings.Contains(tokenName, "-qp"):
// 			tokenName = strings.Replace(tokenName, "-qp", style.Bold("-qp"), 1)
// 		case strings.Contains(tokenName, "-qr"):
// 			tokenName = strings.Replace(tokenName, "-qr", style.Bold("-qr"), 1)

// 		default:
// 			log.Debugf("lawless listing but common token: %s", tokenName)

// 			return
// 		}

// 		osLink := style.TerminalLink(event.Payload.Item.Permalink, event.Payload.Item.Permalink)

// 		go notify.SendMessageViaTelegram(fmt.Sprintf("lawless listing: %s \n amountPaid: %s  url: %s", tokenName, fmt.Sprintf("%5.3f", amountPaid.Ether()), osLink), viper.GetInt64("notifications.manifold.dakma"), "", 0, nil)

// 		highlightMessage := strings.Builder{}
// 		highlightMessage.WriteString("\n")
// 		highlightMessage.WriteString(fmt.Sprintf("  lawless %s | %5.3fΞ | %s\n", tokenName, amountPaid.Ether(), osLink))
// 		highlightMessage.WriteString("\n")

// 		fmt.Println(highlightMessage.String()) //nolint:forbidigo
// 	}
// }

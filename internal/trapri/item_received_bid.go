package trapri

import (
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/ethereum/go-ethereum/common"
)

var (
	tokenTopBids      = map[string]*models.ItemReceivedBid{}
	tokenTopBidsMutex = &sync.Mutex{}
)

func HandleItemReceivedBid(gb *gloomberg.Gloomberg, event *models.ItemReceivedBid) {
	nftID := event.Payload.NftID

	contractAddress := nftID.ContractAddress()

	// our token?
	isOwnToken := gb.OwnWallets.ContainsToken(contractAddress, nftID.TokenID().String())
	// did someone from us make a bid?
	isWatchUsersWallet := gb.Watcher != nil && gb.Watcher.Contains(event.Payload.Maker.Address)

	// check if we hold the token/got a bid
	if !isOwnToken && !isWatchUsersWallet {
		gbl.Log.Debugf("🤷‍♀️ %s | bid for token not held by any of our own wallets", nftID.LinkOS())

		return
	}

	// if it should be a new top bid, we highlight it when printing
	highlightBid := false

	// check if we already have a top bid for this token and if not, add it
	tokenTopBidsMutex.Lock()
	currentTopBid := tokenTopBids[nftID.TID()]
	tokenTopBidsMutex.Unlock()

	switch {
	// no or expired top bid - setting new top bid
	case currentTopBid == nil:
		gbl.Log.Debugf("🍭 no top bid, new top bid: %+v", event.Payload.GetPrice().Wei())

	case currentTopBid.Payload.ExpirationDate.Before(time.Now()):
		gbl.Log.Debugf("🍭 top bid expired, new top bid: %+v", event.Payload.GetPrice().Wei())

	// new bid is higher than current top bid
	case currentTopBid != nil:
		// we add a small amount (still researching how much :D) of ether/wei to the current top bid before comparing
		// to prevent printing a lot of backrunned (=doubled) bids all the time
		// amountToAdd := big.NewInt(13370000000000001) // ≈0.01337....Ξ
		amountToAdd := big.NewInt(7370000000000001) // ≈0.00737....Ξ

		currentTopBidWithBuffer := big.NewInt(0).Add(currentTopBid.Payload.GetPrice().Wei(), amountToAdd)

		if event.Payload.GetPrice().Wei().Cmp(currentTopBidWithBuffer) < 0 {
			gbl.Log.Debugf("🍭 current top bid (+buffer) higher than incoming bid: %+v > %+v", currentTopBidWithBuffer, event.Payload.GetPrice().Wei())

			return
		}

		highlightBid = true
	}

	tokenTopBidsMutex.Lock()
	tokenTopBids[nftID.TID()] = event
	tokenTopBidsMutex.Unlock()

	// seller address
	sellerAddress := common.HexToAddress(event.Payload.Maker.Address.Hex())

	itemName := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]

	//
	// create a TokenTransaction
	ttxBid := &totra.TokenTransaction{
		Tx:          nil,
		TxReceipt:   nil,
		From:        sellerAddress,
		AmountPaid:  event.Payload.GetPrice().Wei(),
		TotalTokens: int64(event.Payload.Quantity),
		Marketplace: &marketplace.OpenSea,
		Action:      degendb.Bid,
		ReceivedAt:  event.Payload.EventTimestamp,
		DoNotPrint:  false,
		Highlight:   highlightBid,
		Transfers: []*totra.TokenTransfer{
			{
				From:         marketplace.OpenSea.ContractAddress(),
				To:           sellerAddress,
				AmountTokens: big.NewInt(int64(event.Payload.Quantity)),
				Token: &token.Token{
					Address: contractAddress,
					ID:      nftID.TokenID(),
					Name:    itemName,
				},
			},
		},
	}

	// format and print
	gb.In.TokenTransactions <- ttxBid

	// send notification
	if isWatchUsersWallet {
		ttxBid.Action = degendb.OwnBid

		// send notification
		// go notify.SendNotification(gb, ttxBid)
	}
}

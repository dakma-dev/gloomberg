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
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

var (
	tokenTopBids      = map[string]*models.ItemReceivedBid{}
	tokenTopBidsMutex = &sync.Mutex{}
)

func HandleItemReceivedBid(gb *gloomberg.Gloomberg, event *models.ItemReceivedBid) {
	nftID := event.Payload.Item.NftID

	// our token?
	isOwnToken := gb.OwnWallets.ContainsToken(nftID.ContractAddress(), nftID.TokenID().String())
	// did someone from us make a bid?
	isWatchUsersWallet := gb.Watcher.Contains(event.Payload.Maker.Address)

	collection := tokencollections.GetCollection(gb, nftID.ContractAddress(), nftID.TokenID().Int64())

	collectionName := "unknown"
	if collection != nil {
		collectionName = collection.Name
	}

	// check if we hold the token/got a bid
	if !isOwnToken && !isWatchUsersWallet {
		gbl.Log.Debugf("🤷‍♀️ %s | bid for token not held by any of our own wallets", nftID.LinkOS())

		return
	}

	// check if we already have a top bid for this token
	// and if so, if the new bid is higher

	// if it should be a new top bid, we highlight it when printing
	highlightBid := false

	// check if we already have a top bid for this token and if not, add it
	tokenTopBidsMutex.Lock()
	if topBid, ok := tokenTopBids[nftID.TID()]; !ok || topBid == nil {
		tokenTopBids[nftID.TID()] = event
	}
	tokenTopBidsMutex.Unlock()

	// get the current top bid for this token
	topBid := tokenTopBids[nftID.TID()]

	log.Debugf("2 %s | own: %+v | watchUser: %+v", collectionName, isOwnToken, isWatchUsersWallet)

	switch {
	// no or expired top bid - setting new top bid
	case topBid == nil || topBid.Payload.ExpirationDate.Before(time.Now()):
		tokenTopBids[nftID.TID()] = event

	// new bid is higher than current top bid
	case topBid != nil:
		// we add a small amount of ether/wei to the current top bid before comparing
		// to prevent printing a lot of backrunned (=doubled) bids all the time
		amountToAdd := big.NewInt(13370000000000001) // ≈0.01337....Ξ
		topBidPriceWithBuffer := big.NewInt(0).Add(topBid.Payload.GetPrice().Wei(), amountToAdd)

		if event.Payload.GetPrice().Wei().Cmp(topBidPriceWithBuffer) < 0 {
			log.Debugf("🍭 bid lower than current topX:\n%+v\n>\n%+v", topBid, event)

			return
		}

		tokenTopBids[nftID.TID()] = event

		highlightBid = true

		log.Debugf("🍭 new topX: %+v", tokenTopBids[nftID.TID()])
	}

	// seller address
	sellerAddress := common.HexToAddress(event.Payload.Maker.Address.Hex())

	itemName := strings.Split(event.Payload.Item.Metadata.Name, " #")[0]

	log.Debugf("3 %s | own: %+v | watchUser: %+v", collectionName, isOwnToken, isWatchUsersWallet)
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
					Address: nftID.ContractAddress(),
					ID:      nftID.TokenID(),
					Name:    itemName,
				},
			},
		},
	}

	if isOwnToken {
		// format and print
		gb.In.TokenTransactions <- ttxBid
	}

	log.Debugf("4 %s | own: %+v | watchUser: %+v", collectionName, isOwnToken, isWatchUsersWallet)

	if isWatchUsersWallet {
		gbl.Log.Infof("🧱 sending telegram notification ItemReceivedBid 🧚 | isOwnToken: %+v | isWatchUsersWallet: %+v", isOwnToken, isWatchUsersWallet)

		ttxBid.Action = degendb.OwnBid

		// send notification
		go notify.SendNotification(gb, ttxBid)
	}
}

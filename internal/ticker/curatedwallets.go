package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	AlphaCaller          *AlphaScore
	alphaCallerKnownTX   = make(map[common.Hash]bool, 0)
	alphaCallerKnownTXMu = &sync.RWMutex{}
)

type AlphaScore struct {
	WalletMap      map[common.Address]*Wallet
	CollectionData map[common.Address]*CollectionStats

	RWMu *sync.RWMutex
	gb   *gloomberg.Gloomberg
}

type CollectionStats struct {
	Transactions         []*totra.TokenTransaction
	ArchivedTransactions []*totra.TokenTransaction
	Stats                *Counters
	// count of transactions when last notification was sent
	Score int32
}

func (s *AlphaScore) AlphaCallerTicker(gb *gloomberg.Gloomberg, alphaCallerTicker *time.Ticker) {

	for range alphaCallerTicker.C {

		for collectionAddress, collection := range AlphaCaller.CollectionData {

			// skip collections with no transactions
			if len(collection.Transactions) == 0 {
				continue
			}

			message := strings.Builder{}

			walletCount := len(collection.Transactions) + len(collection.ArchivedTransactions)

			collectionName := gb.CollectionDB.Collections[collectionAddress].Name
			message.WriteString(fmt.Sprintf("*%d Wallet(s) interacted with %s \n\n*", walletCount, collectionName))
			message.WriteString(fmt.Sprintf("*%s* Score: *%d* %s \n\n", collectionName, collection.Score, getScoreEmoji(collection.Score, walletCount)))
			message.WriteString(fmt.Sprintf("_Latest Transactions per Wallets:_\n"))
			var tokenID *big.Int
			var txHash common.Hash

			currentBlock, _ := gb.ProviderPool.BlockNumber(context.TODO())
			for _, tx := range collection.Transactions {
				wallet := AlphaCaller.WalletMap[tx.From]

				// get current block number

				blocksAgo := currentBlock - tx.TxReceipt.BlockNumber.Uint64()

				message.WriteString(fmt.Sprintf("%d Blocks ago | %s (%d) *%s*  \n", blocksAgo, wallet.Ens, wallet.Score, tx.Action.ActionName()))
				//tokenID = tx.Transfers[0].Token.ID
				_, tokenID = getFirstContractAddressAndTokenID(tx)
				txHash = tx.TxReceipt.TxHash
			}

			if len(collection.ArchivedTransactions) > 0 {
				message.WriteString(fmt.Sprintf("\n\nArchived Transactions per Wallets: \n"))
			}

			for _, tx := range collection.ArchivedTransactions {
				wallet := AlphaCaller.WalletMap[tx.From]
				gloombergReceivedEventAt := tx.ReceivedAt

				message.WriteString(fmt.Sprintf("%s | Wallet %s %s (%d) \n", gloombergReceivedEventAt.Format(time.RFC822), wallet.Ens, tx.Action.ActionName(), wallet.Score))
			}

			// move transactions to archived
			collection.ArchivedTransactions = append(collection.ArchivedTransactions, collection.Transactions...)
			collection.Transactions = make([]*totra.TokenTransaction, 0)

			// send notification via telegram
			if viper.GetString("notifications.smart_wallets.telegram_chat_id") != "" {

				etherscanURL, openseaURL, blurURL := utils.GetLinks(txHash, collectionAddress, tokenID.Int64())

				// emoji arrow up
				replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("ES", etherscanURL),
						tgbotapi.NewInlineKeyboardButtonURL("🧡Blur", blurURL),
						tgbotapi.NewInlineKeyboardButtonURL("🔵OS", openseaURL),
					),
				)

				notify.SendNotificationViaTelegram(message.String(), viper.GetInt64("notifications.smart_wallets.telegram_chat_id"), "", viper.GetInt("notifications.smart_wallets.telegram_reply_to_message_id"), replyMarkup)
			}

		}
	}
}

func getScoreEmoji(score int32, walletCount int) string {

	// walletCount to int32
	averageScore := 0
	if walletCount > 0 {
		averageScore = int(score / int32(walletCount))
	}

	if averageScore < 0 {
		return "🔴"
	}

	if averageScore > 8 {
		return "🟢"
	}

	return "🟡"
}

func NewAlphaScore(gb *gloomberg.Gloomberg) *AlphaScore {
	AlphaCaller = &AlphaScore{
		CollectionData: make(map[common.Address]*CollectionStats, 0),
		WalletMap:      make(map[common.Address]*Wallet, 0),
		RWMu:           &sync.RWMutex{},
		gb:             gb,
	}

	miwSpinner := style.GetSpinner("setting up curated wallets watcher ...")
	_ = miwSpinner.Start()

	fromJSON := ReadCuratedWalletsFromJSON("wallets/wallet_scores_edited_new.json")

	// build wallet map
	for _, address := range fromJSON.Addresses {

		// do a lookup address for ens name
		resolvedAddress, err := gb.ProviderPool.ResolveAddressForENS(context.TODO(), address.Ens)
		if err != nil {
			fmt.Println(fmt.Sprintf("ens resolve error: %s -> %s: %s", address.Ens, address, err))
			//gbl.Log.Info("ens resolve error")
			continue
		}
		//fmt.Println(fmt.Sprintf("ens resolve success: %s -> %s", address.Ens, resolvedAddress))
		address.Address = resolvedAddress
		AlphaCaller.WalletMap[address.Address] = address
	}

	if len(fromJSON.Addresses) > 0 {
		miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(AlphaCaller.WalletMap))), " curated wallets with alpha scores loaded", "\n")))
		_ = miwSpinner.Stop()
	} else {
		_ = miwSpinner.StopFail()
	}

	return AlphaCaller
}

func (s *AlphaScore) AddEvent(eventTx *totra.TokenTransaction) {
	if len(eventTx.GetTransfersByContract()) == 0 {
		return
	}

	if s == nil || s.gb == nil {
		return
	}

	contractAddress, tokenID := getFirstContractAddressAndTokenID(eventTx)
	if tokenID == nil {
		tokenID = big.NewInt(0)
	}

	currentCollection := tokencollections.GetCollection(s.gb, contractAddress, tokenID.Int64())

	s.UpdateScore(currentCollection, eventTx.From, eventTx)
}

func getFirstContractAddressAndTokenID(eventTx *totra.TokenTransaction) (common.Address, *big.Int) {
	var contractAddress common.Address
	var tokenID *big.Int
	for _, transfer := range eventTx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			contractAddress = transfer.Token.Address
			tokenID = transfer.Token.ID
			break
		}
	}
	return contractAddress, tokenID
}

func (s *AlphaScore) UpdateScore(collection *collections.Collection, recipientAddress common.Address, eventTx *totra.TokenTransaction) {
	// check if we already know the transaction the log belongs to
	alphaCallerKnownTXMu.Lock()
	known, ok := alphaCallerKnownTX[eventTx.TxReceipt.TxHash]
	alphaCallerKnownTXMu.Unlock()
	if known && ok {
		// we already know this transaction
		return
	}
	if collection == nil {
		return
	}

	if s.WalletMap[recipientAddress] == nil {
		return
	}

	if !allowedAction(eventTx.Action) {
		return
	}

	fmt.Println("Updating score for collection: ", collection.Name, recipientAddress.String(), eventTx.TxReceipt.TxHash.String())

	s.RWMu.Lock()
	if s.CollectionData[collection.ContractAddress] == nil {
		s.CollectionData[collection.ContractAddress] = &CollectionStats{
			Transactions: make([]*totra.TokenTransaction, 0),
			Score:        0,
		}
		s.CollectionData[collection.ContractAddress].Stats = &Counters{
			Sales:        0,
			Mints:        0,
			Wallets:      make([]*Wallet, 0),
			gbCollection: collection,
		}
	}
	s.CollectionData[collection.ContractAddress].Transactions = append(s.CollectionData[collection.ContractAddress].Transactions, eventTx)

	if eventTx.Action == totra.Purchase || eventTx.Action == totra.Mint {
		s.CollectionData[collection.ContractAddress].Score += s.WalletMap[recipientAddress].Score
	}
	if eventTx.Action == totra.Sale {
		s.CollectionData[collection.ContractAddress].Score -= s.WalletMap[recipientAddress].Score
	}

	s.RWMu.Unlock()
}

func ReadCuratedWalletsFromJSON(filePath string) *Wallets {
	// read json file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		//gbl.Log.Error(err)
	}
	defer file.Close()

	// decode json
	var blueChipWallets *Wallets

	err = json.NewDecoder(file).Decode(&blueChipWallets)
	if err != nil {
		fmt.Println("error decoding file")
		//gbl.Log.Error(err)
	}
	return blueChipWallets
}

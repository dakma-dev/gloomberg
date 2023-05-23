package ticker

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	BlueChips *BlueChipStats
	knownTX   = make(map[common.Hash]bool, 0)
	knownTXMu = &sync.RWMutex{}
)

type BlueChipStats struct {
	BlueChipEvents  []*totra.TokenTransaction
	WalletMap       map[common.Address]*Wallet
	CollectionStats map[common.Address]*Counters

	gb *gloomberg.Gloomberg

	sync.RWMutex
}

type Counters struct {
	Sales          uint64
	SalesTXs       uint64
	SalesVolume    *big.Int
	Mints          uint64
	MintsTXs       uint64
	MintsVolume    *big.Int
	Transfers      uint64
	gbCollection   *collections.Collection
	Wallets        []*Wallet
	RankingMap     map[HolderTypes]uint64
	BlueChipEvents []*totra.TokenTransaction
}

type BlueChipRanking struct {
	// bcType HolderTypes
	// count  uint64
}

func (s *BlueChipStats) BlueChipTicker(ticker *time.Ticker, queueOutput *chan string) {
	rowStyle := style.AlmostWhiteStyle

	for range ticker.C {
		// iterate over Counters
		for address, counters := range BlueChips.CollectionStats {
			if counters.Sales > viper.GetUint64("notifications.bluechip.threshold") {
				line := strings.Builder{}

				line.WriteString(rowStyle.Faint(true).Render(fmt.Sprintf("%s ", counters.gbCollection.Name)))
				line.WriteString(rowStyle.Faint(true).Render(fmt.Sprintf("%s: %d sales", address.String(), counters.Sales)))

				*queueOutput <- line.String()

				// send telegram message
				telegramMessage := strings.Builder{}
				telegramMessage.WriteString("🔵 bought: ")

				openseaURL := fmt.Sprintf("https://opensea.io/assets/ethereum/%s", counters.gbCollection.ContractAddress)

				telegramMessage.WriteString(fmt.Sprintf("%s: %d txs\n", "["+counters.gbCollection.Name+"]("+openseaURL+")", len(counters.BlueChipEvents)))

				rankingMap := counters.RankingMap
				// sort rankingMap by value
				keys := make([]HolderTypes, 0, len(rankingMap))

				for key := range rankingMap {
					keys = append(keys, key)
				}

				sort.SliceStable(keys, func(i, j int) bool {
					return rankingMap[keys[i]] > rankingMap[keys[j]]
				})

				for _, key := range keys {
					telegramMessage.WriteString(fmt.Sprintf("• %d %s •", rankingMap[key], GetEmojiMapping(key)))
				}

				//telegramMessage.WriteString(fmt.Sprintf("\n  %d tokens", counters.Sales))

				telegramMessage.WriteString("\n")

				groupByContracts := make(map[common.Address]string)

				groupByWallets := make(map[common.Address]string)

				groupByType := make(map[HolderTypes]int64)

				totalTokensTransferreToBlueChips := big.NewInt(0)

				for _, ttx := range counters.BlueChipEvents {
					firstNFTTransaction := s.getNFTInfo(ttx)

					// TODO check all involved wallets (blur bid dump case)
					recipientAddress := firstNFTTransaction.To

					groupByWallets[recipientAddress] = recipientAddress.String()

					transfers := ttx.GetNFTReceivers()[recipientAddress]

					amountTokens := s.getTransferredTokensCount(transfers)

					totalTokensTransferreToBlueChips = totalTokensTransferreToBlueChips.Add(totalTokensTransferreToBlueChips, amountTokens)

					etherscanURL := fmt.Sprintf("https://etherscan.io/tx/%s", ttx.Tx.Hash().String())

					telegramMessage.WriteString(fmt.Sprintf("  %s bought %d nfts (%s) \n", recipientAddress.String(), amountTokens, "[tx]("+etherscanURL+")"))

					wallet := s.WalletMap[recipientAddress]
					// print holdertypes
					for _, holderType := range wallet.Types {
						telegramMessage.WriteString(fmt.Sprintf("•%s•", GetEmojiMapping(holderType)))
						groupByType[holderType] = groupByType[holderType] + amountTokens.Int64()
					}

					telegramMessage.WriteString("\n")

					groupByContracts[*ttx.Tx.To()] = ttx.Tx.To().String()
				}

				for holderType, amount := range groupByType {
					telegramMessage.WriteString(fmt.Sprintf("• %s : %d •", GetEmojiMapping(holderType), amount))
				}

				telegramMessage.WriteString(fmt.Sprintf("\n Total:  %d tokens", totalTokensTransferreToBlueChips))

				telegramMessage.WriteString(fmt.Sprintf("\n Unique wallets:  %d ", len(groupByWallets)))

				//for contractAddress, contractNames := range groupByContracts {
				//telegramMessage.WriteString(fmt.Sprintf("  %s: %s\n", contractAddress.String(), contractNames))
				//}

				if telegramMessage.Len() > 0 {
					if viper.GetString("notifications.manifold.dakma") != "" {
						notify.SendMessageViaTelegram(telegramMessage.String(), viper.GetInt64("notifications.bluechip.telegram_chat_id"), "", viper.GetInt("notifications.bluechip.telegram_reply_to_message_id"), nil)

						s.resetCounters(counters)
					}
				}
			}
		}
	}
}

func (s *BlueChipStats) getTransferredTokensCount(transfers []*totra.TokenTransfer) *big.Int {
	amountTokens := big.NewInt(0)

	for _, transfer := range transfers {
		amountTokens = amountTokens.Add(amountTokens, transfer.AmountTokens)
	}

	return amountTokens
}

func (s *BlueChipStats) resetCounters(counters *Counters) {
	counters.Sales = 0
	counters.BlueChipEvents = make([]*totra.TokenTransaction, 0)
	counters.MintsTXs = 0
	counters.Mints = 0
	counters.MintsVolume = big.NewInt(0)
	counters.SalesTXs = 0
	counters.SalesVolume = big.NewInt(0)
	counters.Transfers = 0
	counters.RankingMap = make(map[HolderTypes]uint64, 0)
}

func GetEmojiMapping(holderType HolderTypes) string {
	switch holderType {
	case BAYC:
		return "🐵"
	case CryptoPunks:
		return "🅿️"
	case MAYC:
		return "🧟"
	case Azuki:
		return "⛩"
	case RLD:
		return "👯"
	case MOONBIRDS:
		return "🦉"
	case PUDGYPENGUINS:
		return "🐧"
	case DOODLES:
		return "🌈"
	case Goblintown:
		return "👹"
	case CYBERKONGZ:
		return "🦍"
	case Captainz:
		return "🏴‍☠️"
	case CloneX:
		return "👟"
	}

	return ""
}

func NewBlueChipTicker(gb *gloomberg.Gloomberg) *BlueChipStats {
	BlueChips = &BlueChipStats{
		BlueChipEvents:  make([]*totra.TokenTransaction, 0),
		CollectionStats: make(map[common.Address]*Counters, 0),
		WalletMap:       make(map[common.Address]*Wallet, 0),
		// RWMu:            &sync.RWMutex{},
		gb: gb,
	}
	BlueChips.RLock()
	defer BlueChips.RUnlock()

	miwSpinner := style.GetSpinner("setting up blue chip wallets...")
	_ = miwSpinner.Start()

	// bayc, mayc, cryptopunks, azuki, cool cats, world of women, clone x

	// fill bluechip wallet map
	// for _, address := range fromJSON.Addresses {
	//	BlueChips.WalletMap[address.Address] = address
	//}

	readBlueChipWalltesFromJSON("wallets/0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d.json", BAYC)
	readBlueChipWalltesFromJSON("wallets/0x60e4d786628fea6478f785a6d7e704777c86a7c6.json", MAYC)
	readBlueChipWalltesFromJSON("wallets/0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb.json", CryptoPunks)
	readBlueChipWalltesFromJSON("wallets/0xed5af388653567af2f388e6224dc7c4b3241c544.json", Azuki)
	readBlueChipWalltesFromJSON("wallets/0x513cd71defc801b9c1aa763db47b5df223da77a2.json", RLD)
	readBlueChipWalltesFromJSON("wallets/0x8a90cab2b38dba80c64b7734e58ee1db38b8992e.json", DOODLES)
	readBlueChipWalltesFromJSON("wallets/0x23581767a106ae21c074b2276d25e5c3e136a68b.json", MOONBIRDS)
	readBlueChipWalltesFromJSON("wallets/0xbd3531da5cf5857e7cfaa92426877b022e612cf8.json", PUDGYPENGUINS)
	readBlueChipWalltesFromJSON("wallets/0x769272677fab02575e84945f03eca517acc544cc.json", Captainz)

	readBlueChipWalltesFromJSON("wallets/0x49cf6f5d44e70224e2e23fdcdd2c053f30ada28b.json", CloneX)
	readBlueChipWalltesFromJSON("wallets/0x57a204aa1042f6e66dd7730813f4024114d74f37.json", CYBERKONGZ)

	if len(BlueChips.WalletMap) > 0 {
		miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(BlueChips.WalletMap))), " blue chip wallets loaded", "\n")))
	} else {
		_ = miwSpinner.StopFail()
	}

	_ = miwSpinner.Stop()

	return BlueChips
}

func readBlueChipWalltesFromJSON(file string, bluechipType HolderTypes) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		gbl.Log.Error("file %s does not exist", file)

		return
	}
	fromJSON := ReadWalletsFromJSON(file)
	for _, address := range fromJSON.OwnerAddresses {
		// to common.Address
		hexAddress := common.HexToAddress(address)
		if BlueChips.WalletMap[hexAddress] == nil {
			BlueChips.WalletMap[hexAddress] = &Wallet{
				Address: hexAddress,
				Types:   make([]HolderTypes, 0),
			}
		}
		//	if BlueChips.WalletMap[hexAddress].Types == nil {
		//		BlueChips.WalletMap[hexAddress].Types = make([]HolderTypes, 0)
		//	}
		BlueChips.WalletMap[hexAddress].Types = append(BlueChips.WalletMap[hexAddress].Types, bluechipType)
	}
}

func allowedAction(action totra.TxType) bool {
	switch action {
	case totra.Sale, totra.Purchase, totra.Mint:
		return true
	}

	return false
}

func (s *BlueChipStats) CheckForBlueChipInvolvment(eventTx *totra.TokenTransaction) {
	if len(eventTx.Transfers) < 1 {
		return
	}

	if !(allowedAction(eventTx.Action)) {
		return
	}

	if ignoreContract(eventTx) {
		return
	}

	if !s.isTransactionABluechipTX(eventTx) {
		return
	}

	if !eventTx.IsMovingNFTs() {
		return
	}
	// check if we already know the transaction the log belongs to
	knownTXMu.Lock()
	known, ok := knownTX[eventTx.TxReceipt.TxHash]
	knownTXMu.Unlock()

	if known && ok {
		// we already know this transaction
		return
	}

	if eventTx.Transfers == nil {
		return
	}

	firstNFTTransaction := s.getNFTInfo(eventTx)

	if firstNFTTransaction == nil {
		return
	}

	contractAddress := firstNFTTransaction.Token.Address

	// check if contractAddress is allowed
	if s.isContractIgnored(contractAddress) {
		return
	}

	s.Lock()
	defer s.Unlock()

	if s.CollectionStats[contractAddress] == nil {
		s.CollectionStats[contractAddress] = &Counters{
			Sales:          0,
			Mints:          0,
			Transfers:      0,
			SalesVolume:    big.NewInt(0),
			Wallets:        make([]*Wallet, 0),
			RankingMap:     make(map[HolderTypes]uint64, 0),
			BlueChipEvents: make([]*totra.TokenTransaction, 0),
		}

		currentCollection := tokencollections.GetCollection(s.gb, firstNFTTransaction.Token.Address, firstNFTTransaction.Token.ID.Int64())
		s.CollectionStats[contractAddress].gbCollection = currentCollection
	}

	// better check all ttx transfers
	recipientAddress := firstNFTTransaction.To
	s.CollectionStats[contractAddress].Wallets = append(s.CollectionStats[contractAddress].Wallets, s.WalletMap[recipientAddress])

	wallet := s.WalletMap[recipientAddress]
	for _, holderType := range wallet.Types {
		s.CollectionStats[contractAddress].RankingMap[holderType]++
	}

	// TODO get correct number of tokens
	numCollectionTokens := uint64(0)

	for _, transfer := range eventTx.Transfers {
		numCollectionTokens += transfer.AmountTokens.Uint64()
	}

	if eventTx.TotalTokens == 0 {
		eventTx.TotalTokens = 1
	}

	s.CollectionStats[contractAddress].BlueChipEvents = append(s.CollectionStats[contractAddress].BlueChipEvents, eventTx)

	switch eventTx.Action {
	case totra.Sale:
		s.CollectionStats[contractAddress].SalesTXs++
		s.CollectionStats[contractAddress].Sales += uint64(eventTx.TotalTokens)
	case totra.Mint:
		s.CollectionStats[contractAddress].Sales += uint64(eventTx.TotalTokens)
		s.CollectionStats[contractAddress].Mints += uint64(eventTx.TotalTokens)
	}
}

func (s *BlueChipStats) getNFTInfo(eventTx *totra.TokenTransaction) *totra.TokenTransfer {
	for _, transfer := range eventTx.Transfers {
		if transfer.Standard.IsERC721orERC1155() && s.ContainsWallet(transfer.To) && !s.isContractIgnored(transfer.Token.Address) {
			return transfer
		}
	}

	return nil
}

func (s *BlueChipStats) isTransactionABluechipTX(eventTx *totra.TokenTransaction) bool {
	// check if any transfers involves a blue chip wallet
	blueChipInvolved := false

	for _, transfer := range eventTx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			blueChipInvolved = blueChipInvolved || s.ContainsWallet(transfer.To)
		}
	}

	return blueChipInvolved
}

func (s *BlueChipStats) ContainsWallet(address common.Address) bool {
	if s == nil {
		return false
	}

	s.RLock()

	if s.WalletMap == nil {
		s.RUnlock()

		return false
	}

	if s.WalletMap[address] != nil {
		s.RUnlock()

		return true
	}

	s.RUnlock()

	return false
}

func (s *BlueChipStats) GetStats(address common.Address) *Counters {
	s.RLock()
	defer s.RUnlock()

	if s.CollectionStats == nil {
		return nil
	}

	return s.CollectionStats[address]
}

func (s *BlueChipStats) isContractIgnored(address common.Address) bool {
	if address == common.HexToAddress("0xc36442b4a4522e871399cd717abdd847ab11fe88") {
		return true
	}
	// Emblem Vault V4
	if address == common.HexToAddress("0x82C7a8f707110f5FBb16184A5933E9F78a34c6ab") {
		return true
	}

	// Blend
	if address == common.HexToAddress("0x29469395eAf6f95920E59F858042f0e28D98a20B") {
		return true
	}

	return false
}

func ignoreContract(ttx *totra.TokenTransaction) bool {
	if ttx.Tx == nil || ttx.Tx.To() == nil {
		return true
	}

	// Uniswap V3: Positions NFT
	if *ttx.Tx.To() == common.HexToAddress("0xc36442b4a4522e871399cd717abdd847ab11fe88") {
		return true
	}
	// Emblem Vault V4
	if *ttx.Tx.To() == common.HexToAddress("0x82C7a8f707110f5FBb16184A5933E9F78a34c6ab") {
		return true
	}

	// Blend
	if *ttx.Tx.To() == common.HexToAddress("0x29469395eAf6f95920E59F858042f0e28D98a20B") {
		return true
	}

	return false
}

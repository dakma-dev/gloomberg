package ticker

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

var (
	BlueChips                  *BlueChipStats
	knownTX                    = make(map[common.Hash]bool)
	knownTXMu                  = &sync.RWMutex{}
	amountOfBlueChipsToBeShown = 3

	readMutex = &sync.RWMutex{}
)

type BlueChipStats struct {
	BlueChipEvents  []*totra.TokenTransaction
	WalletMap       map[common.Address]*Wallet
	CollectionStats map[common.Address]*Counters

	gb *gloomberg.Gloomberg

	sync.RWMutex
}

func (s *BlueChipStats) GetCounterByAddress(address common.Address) *Counters {
	readMutex.RLock()
	counter := BlueChips.CollectionStats[address]
	readMutex.RUnlock()

	return counter
}

type Counters struct {
	Sales          uint64
	Mints          uint64
	gbCollection   *collections.Collection
	Wallets        []*Wallet
	RankingMap     map[HolderTypes]uint64
	BlueChipEvents []*totra.TokenTransaction

	TotalTokensTransferredToBlueChips *big.Int
	GroupByWallets                    map[common.Address]string
}

func (c *Counters) GetTXCount() uint64 {
	return c.Sales + c.Mints
}

type BlueChipRanking struct {
	// bcType HolderTypes
	// count  uint64
}

func (s *BlueChipStats) BlueChipTicker(ticker *time.Ticker, queueOutput *chan string) {
	rowStyle := style.AlmostWhiteStyle

	tokenTransactionsChannel := s.gb.SubscribeTokenTransactions()
	go func() {
		for ttx := range tokenTransactionsChannel {
			s.CheckForBlueChipInvolvment(ttx)
		}
	}()

	if !viper.GetBool("notifications.bluechip.enabled") {
		return
	}

	for range ticker.C {
		// iterate over Counters
		for address, counters := range BlueChips.CollectionStats {
			customThreshold := false
			// skull of luci
			if address == common.HexToAddress("0xc9041f80dce73721a5f6a779672ec57ef255d27c") {
				// custom thresholds
				if len(counters.GroupByWallets) >= 1 {
					customThreshold = true
				}
			}

			if len(counters.GroupByWallets) >= viper.GetInt("notifications.bluechip.threshold") || customThreshold {
				line := strings.Builder{}

				line.WriteString(rowStyle.Faint(true).Render(fmt.Sprintf("%s ", counters.gbCollection.Name)))
				line.WriteString(rowStyle.Faint(true).Render(fmt.Sprintf("%s: %d sales tx", address.String(), len(counters.BlueChipEvents))))

				*queueOutput <- line.String()

				// send telegram message
				telegramMessage := strings.Builder{}
				telegramMessage.WriteString("🔵 bought: ")

				openseaURL := fmt.Sprintf("https://opensea.io/assets/ethereum/%s", counters.gbCollection.ContractAddress)

				telegramMessage.WriteString(fmt.Sprintf("%s: (%d txs)\n", "["+counters.gbCollection.Name+"]("+openseaURL+")", len(counters.BlueChipEvents)))

				rankingMap := counters.RankingMap
				// sort rankingMap by value
				keys := make([]HolderTypes, 0, len(rankingMap))

				for key := range rankingMap {
					keys = append(keys, key)
				}

				sort.SliceStable(keys, func(i, j int) bool {
					return rankingMap[keys[i]] > rankingMap[keys[j]]
				})

				keys = sortKeysAsc(rankingMap)

				for _, key := range keys[:amountOfBlueChipsToBeShown] {
					telegramMessage.WriteString(fmt.Sprintf("%s: %d | ", GetEmojiMapping(key), rankingMap[key]))
				}
				telegramMessage.WriteString("...")

				// telegramMessage.WriteString(fmt.Sprintf("\n  %d tokens", counters.Sales))

				groupByContracts := make(map[common.Address]string)

				groupSalesByWallets := make(map[common.Address]int64)

				groupSalesByType := make(map[HolderTypes]uint64)

				groupUniqueWalletsByType := make(map[HolderTypes]uint64)

				totalTokensTransferreToBlueChips := big.NewInt(0)

				for _, ttx := range counters.BlueChipEvents {
					firstNFTTransaction := s.getNFTInfo(ttx)

					// TODO check all involved wallets (blur bid dump case)
					recipientAddress := firstNFTTransaction.To

					transfers := ttx.GetNFTReceivers()[recipientAddress]

					amountTokens := s.getTransferredTokensCount(transfers)

					totalTokensTransferreToBlueChips = totalTokensTransferreToBlueChips.Add(totalTokensTransferreToBlueChips, amountTokens)

					// telegramMessage.WriteString(fmt.Sprintf("  %s bought %d nfts (%s) \n", recipientAddress.String(), amountTokens, "[tx]("+etherscanURL+")"))

					wallet := s.WalletMap[recipientAddress]
					// print sales per bc type
					for _, holderType := range wallet.Types {
						groupSalesByType[holderType] += amountTokens.Uint64()

						// increment unique wallet counter for each type
						if groupSalesByWallets[recipientAddress] == 0 {
							groupUniqueWalletsByType[holderType]++
						}
					}

					groupSalesByWallets[recipientAddress] += amountTokens.Int64()

					groupByContracts[*ttx.Tx.To()] = ttx.Tx.To().String()
				}

				telegramMessage.WriteString(fmt.Sprintf("\n Total:  %d tokens \n", totalTokensTransferreToBlueChips))
				keys = sortKeysAsc(groupSalesByType)
				for _, key := range keys[:amountOfBlueChipsToBeShown] {
					telegramMessage.WriteString(fmt.Sprintf("%s : %d | ", GetEmojiMapping(key), groupSalesByType[key]))
				}
				telegramMessage.WriteString("...")

				telegramMessage.WriteString(fmt.Sprintf("\n Unique wallets:  %d \n", len(groupSalesByWallets)))

				// sorting unique wallets by type
				keys = sortKeysAsc(groupUniqueWalletsByType)

				for _, key := range keys[:amountOfBlueChipsToBeShown] {
					telegramMessage.WriteString(fmt.Sprintf("%s: %d | ", GetEmojiMapping(key), groupUniqueWalletsByType[key]))
				}
				telegramMessage.WriteString("...")
				for _, key := range keys[:amountOfBlueChipsToBeShown] {
					telegramMessage.WriteString(fmt.Sprintf("\n%s ", GetHashTags(key)))
				}
				if telegramMessage.Len() > 0 {
					if viper.GetString("notifications.manifold.dakma") != "" {
						notify.SendMessageViaTelegram(telegramMessage.String(), viper.GetInt64("notifications.bluechip.telegram_chat_id"), "", viper.GetInt("notifications.bluechip.telegram_reply_to_message_id"), nil)

						cred := &TwitterCredentials{
							ConsumerKey:       viper.GetString("twitter.consumer_key"),
							ConsumerSecret:    viper.GetString("twitter.consumer_secret"),
							AccessToken:       viper.GetString("twitter.access_token"),
							AccessTokenSecret: viper.GetString("twitter.access_token_secret"),
						}

						twitterClient := NewTwitterClient(cred)
						twitterClient.PostTweetV2(telegramMessage.String())

						s.resetCounters(counters)
					}
				}
			}
		}
	}
}

func sortKeysAsc(groupUniqueWalletsByType map[HolderTypes]uint64) []HolderTypes {
	keys := make([]HolderTypes, 0, len(groupUniqueWalletsByType))
	for key := range groupUniqueWalletsByType {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return groupUniqueWalletsByType[keys[i]] > groupUniqueWalletsByType[keys[j]]
	})

	return keys
}

func (s *BlueChipStats) getTransferredTokensCount(transfers []*totra.TokenTransfer) *big.Int {
	amountTokens := big.NewInt(0)

	for _, transfer := range transfers {
		amountTokens = amountTokens.Add(amountTokens, transfer.AmountTokens)
	}

	return amountTokens
}

func (s *BlueChipStats) resetCounters(counters *Counters) {
	counters.BlueChipEvents = make([]*totra.TokenTransaction, 0)
	counters.RankingMap = make(map[HolderTypes]uint64, 0)
	counters.TotalTokensTransferredToBlueChips = big.NewInt(0)
	counters.GroupByWallets = make(map[common.Address]string)
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
	case DeGods:
		return "⬜"
	case SKULLSOFLUCI:
		return "💀"
	case MILADY:
		return "👩"
	case NOUNS:
		return "🈁"
	}

	return ""
}

func GetHashTags(types HolderTypes) string {
	switch types {
	case BAYC:
		return "#BAYC"
	case CryptoPunks:
		return "#CryptoPunks"
	case MAYC:
		return "#MAYC"
	case Azuki:
		return "#Azuki"
	case RLD:
		return "#RLD"
	case MOONBIRDS:
		return "#MOONBIRDS"
	case PUDGYPENGUINS:
		return "#PUDGYPENGUINS"
	case DOODLES:
		return "#DOODLES"
	case Goblintown:
		return "#Goblintown"
	case CYBERKONGZ:
		return "#CYBERKONGZ"
	case Captainz:
		return "#Captainz"
	case CloneX:
		return "#CloneX"
	case DeGods:
		return "#DeGods"
	}

	return ""
}

func NewBlueChipTicker(gb *gloomberg.Gloomberg) *BlueChipStats {
	BlueChips = &BlueChipStats{
		BlueChipEvents:  make([]*totra.TokenTransaction, 0),
		CollectionStats: make(map[common.Address]*Counters),
		WalletMap:       make(map[common.Address]*Wallet),
		// RWMu:            &sync.RWMutex{},
		gb: gb,
	}
	BlueChips.RLock()
	defer BlueChips.RUnlock()

	miwSpinner := style.GetSpinner("setting up blue chip wallets...")
	_ = miwSpinner.Start()

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
	readBlueChipWalltesFromJSON("wallets/0x8821bee2ba0df28761afff119d66390d594cd280.json", DeGods)
	readBlueChipWalltesFromJSON("wallets/0xc9041f80dce73721a5f6a779672ec57ef255d27c.json", SKULLSOFLUCI)
	readBlueChipWalltesFromJSON("wallets/0x5af0d9827e0c53e4799bb226655a1de152a425a5.json", MILADY)
	readBlueChipWalltesFromJSON("wallets/0x9c8ff314c9bc7f6e59a9d9225fb22946427edc03.json", NOUNS)

	/**
	  bluechip art feature:
	  fidenza:         0xa7d8d9ef8d8ce8992df33d8b8cf4aebabd5bd270 projectID 78
	  ringers:         0xa7d8d9ef8d8ce8992df33d8b8cf4aebabd5bd270 projectID 13
	  The Eternal Pump 0xa7d8d9ef8d8ce8992df33d8b8cf4aebabd5bd270 projectID 22
	  squiggle 0x059edd72cd353df5106d2b9cc5ab83a52287ac3a projectID 0
	  autoglyphs: 0xd4e4078ca3495de5b1d4db434bebc5a986197782 (own contract)
	  sam spratt editions: 0xda6558fa1c2452938168ef79dfd29c45aba8a32b
	  skull of luci 0xc9041f80dce73721a5f6a779672ec57ef255d27c
	  twin flames 0x3c28de567d1412b06f43b15e9f75129625fa6e8c
	   qql 0x845dd2a7ee2a92a0518ab2135365ed63fdba0c88 ???
	*/

	if len(BlueChips.WalletMap) > 0 {
		miwSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(strconv.Itoa(len(BlueChips.WalletMap))), " blue chip wallets loaded", "\n"))
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

		BlueChips.WalletMap[hexAddress].Types = append(BlueChips.WalletMap[hexAddress].Types, bluechipType)
	}
}

func allowedAction(action degendb.EventType) bool {
	switch action {
	case degendb.Sale, degendb.Purchase, degendb.Mint:
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
			Wallets:                           make([]*Wallet, 0),
			RankingMap:                        make(map[HolderTypes]uint64),
			BlueChipEvents:                    make([]*totra.TokenTransaction, 0),
			TotalTokensTransferredToBlueChips: big.NewInt(0),
			GroupByWallets:                    make(map[common.Address]string),
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

	s.CollectionStats[contractAddress].GroupByWallets[recipientAddress] = recipientAddress.String()

	transfers := eventTx.GetNFTReceivers()[recipientAddress]
	// create slice for event
	amountTokens := s.getTransferredTokensCount(transfers)
	s.CollectionStats[contractAddress].TotalTokensTransferredToBlueChips.Add(s.CollectionStats[contractAddress].TotalTokensTransferredToBlueChips, amountTokens)

	switch eventTx.Action {
	case degendb.Sale:
		s.CollectionStats[contractAddress].Sales += amountTokens.Uint64()
	case degendb.Mint:
		s.CollectionStats[contractAddress].Mints += amountTokens.Uint64()
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
		if transfer.Standard.IsERC721orERC1155() && transfer.To != internal.ZeroAddress {
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

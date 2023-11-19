package ticker

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	Manifold                  *ManifoldStats
	manifoldContractAddresses = make(map[common.Address]bool)
	alreadyPrinted            = make(map[common.Hash]bool)
	alreadyPrintedMu          = &sync.RWMutex{}
)

type ManifoldStats struct {
	ManifoldEvents []*totra.TokenTransaction
	gb             *gloomberg.Gloomberg
}

func getManifoldAddressesFromConfig() {
	// get yaml object from viper
	rawWatchConfig, ok := viper.Get("contracts.manifold").([]interface{})
	if !ok {
		gbl.Log.Warnf("watch configuration is not an array, skipping")

		return
	}

	for _, walletConfig := range rawWatchConfig {
		var contract *wallet.Wallet

		decodeHooks := mapstructure.ComposeDecodeHookFunc(
			hooks.StringToAddressHookFunc(),
			hooks.StringToLipglossColorHookFunc(),
		)

		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: decodeHooks,
			Result:     &contract,
		})

		err := decoder.Decode(walletConfig)
		if err != nil {
			gbl.Log.Warnf("reading wallet group configuration failed: %+v", err)

			return
		}

		// print the wallet address
		gbl.Log.Infof("watching manifold contract: %s", contract.Address.String())

		manifoldContractAddresses[contract.Address] = true
	}
}

func (s *ManifoldStats) IsManifoldContractAddress(address common.Address) bool {
	if _, ok := manifoldContractAddresses[address]; ok {
		return true
	}

	return false
}

func (s *ManifoldStats) AppendManifoldEvent(event *totra.TokenTransaction) {
	if s != nil && s.ManifoldEvents != nil {
		// check if we already know the transaction the log belongs to
		alreadyPrintedMu.Lock()
		known, ok := alreadyPrinted[event.TxReceipt.TxHash]
		alreadyPrintedMu.Unlock()
		if known && ok {
			// we already know this transaction
			return
		}

		Manifold.ManifoldEvents = append(Manifold.ManifoldEvents, event)
	}
}

func (s *ManifoldStats) ManifoldTicker(manifoldTicker *time.Ticker, queueOutput *chan string) {
	rowStyle := style.DarkGrayStyle

	for range manifoldTicker.C {
		intro := style.DarkerGrayStyle.Render("~  ") + style.DarkGrayStyle.Render("manifold") + style.DarkerGrayStyle.Render("  ~   ")

		maxTickerStatsLines := 5

		// sort by sales
		sort.Slice(s.ManifoldEvents, func(i, j int) bool {
			return s.getCollection(s.ManifoldEvents[i]).Counters.Mints > s.getCollection(s.ManifoldEvents[j]).Counters.Mints
		})

		// every new tick -> new prints

		aggregrateEvents := make(map[common.Address]bool)

		telegramMessage := strings.Builder{}

		for _, event := range s.ManifoldEvents {
			collection := s.getCollection(event)
			if aggregrateEvents[collection.ContractAddress] {
				continue
			}

			aggregrateEvents[collection.ContractAddress] = true

			eventTimestamp := rowStyle.Render(event.ReceivedAt.Format("15:04:05"))
			manifoldLine := strings.Builder{}

			manifoldLine.WriteString(eventTimestamp)
			manifoldLine.WriteString(" " + event.Action.Icon())

			if event.TotalTokens == 0 {
				event.TotalTokens = 1
			}

			pricePerItem := big.NewInt(0).Div(event.AmountPaid, big.NewInt(event.TotalTokens))
			priceEtherPerItem, _ := utils.WeiToEther(pricePerItem).Float64()

			manifoldLine.WriteString(" " + rowStyle.Render(fmt.Sprintf("%6.3f", priceEtherPerItem)))
			telegramMessage.WriteString(fmt.Sprintf("%6.3f", priceEtherPerItem))

			collectionStyle := lipgloss.NewStyle().Foreground(collection.Colors.Primary)

			manifoldLine.WriteString(collectionStyle.Faint(true).Render("Ξ"))
			telegramMessage.WriteString("Ξ")

			var tokenInfo string

			if event.TotalTokens > 1 {
				tokenInfo = fmt.Sprintf("%s %s", rowStyle.Render(fmt.Sprintf("%dx", event.TotalTokens)), collectionStyle.Faint(true).Render(collection.Name))
			} else {
				tokenInfo = style.FormatTokenInfo(event.Transfers[0].Token.ID, collection.Name, collection.Style(), collection.StyleSecondary(), true, true)
			}

			manifoldLine.WriteString(" " + tokenInfo)

			openseaURL := utils.GetOpenseaItemLink(collection.ContractAddress.String(), event.Transfers[0].Token.ID.Int64())

			telegramMessage.WriteString(" · [" + collection.Name + "](" + openseaURL + ")")

			manifoldLine.WriteString(" | " + style.TrendLightGreenStyle.Render(strconv.FormatUint(collection.Counters.Mints, 10)))

			telegramMessage.WriteString(" | " + strconv.FormatUint(collection.Counters.Mints, 10) + "x")

			if collection.Counters.Mints > 200 {
				telegramMessage.WriteString(" " + "🚀")
			}

			mintVolumeEther, _ := utils.WeiToEther(collection.Counters.MintVolume).Float64()
			manifoldLine.WriteString(" | " + style.TrendLightGreenStyle.Render(fmt.Sprint(mintVolumeEther)))
			manifoldLine.WriteString(intro)
			telegramMessage.WriteString(" | " + fmt.Sprint(mintVolumeEther) + "Ξ")

			if mintVolumeEther > 10 {
				telegramMessage.WriteString(" " + "🚀")
			}

			if maxTickerStatsLines <= 0 {
				break
			}

			telegramMessage.WriteString("\n")

			*queueOutput <- manifoldLine.String()
			maxTickerStatsLines--
		}

		// send telegram message
		if telegramMessage.Len() > 0 && viper.GetBool("notifications.manifold.enabled") {
			// manifold ticker channel id -1001725324468
			// no styling information for telegram
			notify.SendMessageViaTelegram(telegramMessage.String(), viper.GetInt64("notifications.manifold.manifold_ticker_channel"), "", viper.GetInt("notifications.manifold.telegram_reply_to_message_id"), nil)
		}
	}
}

func (s *ManifoldStats) OneMinuteTicker(manifoldTicker *time.Ticker) {
	for range manifoldTicker.C {
		maxTickerStatsLines := 5

		// sort by sales
		sort.Slice(s.ManifoldEvents, func(i, j int) bool {
			return s.getCollection(s.ManifoldEvents[i]).Counters.Mints > s.getCollection(s.ManifoldEvents[j]).Counters.Mints
		})

		// every new tick -> new prints

		aggregrateEvents := make(map[common.Address]bool)

		telegramMessage := strings.Builder{}

		for _, event := range s.ManifoldEvents {
			collection := s.getCollection(event)
			if aggregrateEvents[collection.ContractAddress] {
				continue
			}

			salesVolumeEther, _ := utils.WeiToEther(collection.Counters.SalesVolume).Float64()
			if collection.Counters.Mints < 200 && salesVolumeEther < 10 {
				continue
			}

			// try to acquire the lock
			if viper.GetBool("redis.enabled") {
				notificationLock, err := s.gb.Rueidi.NotificationLockWtihDuration(collection.ContractAddress.Hex(), time.Hour*8)
				if notificationLock == nil || err != nil {
					gbl.Log.Infof("notification lock for %s already exists", style.BoldStyle.Render(event.TxReceipt.TxHash.String()))

					continue
				}

				gbl.Log.Infof("notification lock for %s acquired, trying to send...", style.BoldStyle.Render(event.TxReceipt.TxHash.String()))
			}

			aggregrateEvents[collection.ContractAddress] = true

			if event.TotalTokens > 0 {
				pricePerItem := big.NewInt(0).Div(event.AmountPaid, big.NewInt(event.TotalTokens))
				priceEtherPerItem, _ := utils.WeiToEther(pricePerItem).Float64()
				telegramMessage.WriteString(fmt.Sprintf("%6.3f", priceEtherPerItem))
				telegramMessage.WriteString("Ξ")
			}

			openseaURL := utils.GetOpenseaItemLink(collection.ContractAddress.String(), event.Transfers[0].Token.ID.Int64())

			telegramMessage.WriteString(" · [" + collection.Name + "](" + openseaURL + ")")
			// telegramMessage.WriteString(" " + collection.Name)
			telegramMessage.WriteString(" | " + strconv.FormatUint(collection.Counters.Mints, 10) + "x")

			if collection.Counters.Mints >= 200 {
				telegramMessage.WriteString(" " + "🚀")
			}

			telegramMessage.WriteString(" | " + fmt.Sprint(salesVolumeEther) + "Ξ")

			if salesVolumeEther >= 10 {
				telegramMessage.WriteString(" " + "🚀")
			}

			if maxTickerStatsLines <= 0 {
				break
			}

			telegramMessage.WriteString("\n")

			maxTickerStatsLines--

			// TODO own counter | yepp, own counter! will drive us crazy otherwise :D
			collection.Counters.Mints = 0
			collection.Counters.SalesVolume = big.NewInt(0)
		}

		// send telegram message
		if telegramMessage.Len() > 0 {
			if viper.GetString("notifications.manifold.dakma") != "" {
				notify.SendMessageViaTelegram(telegramMessage.String(), viper.GetInt64("notifications.manifold.dakma"), "", 0, nil)
			}

			notify.SendMessageViaTelegram(telegramMessage.String(), viper.GetInt64("notifications.manifold.manifold_ticker_channel"), "", 0, nil)
		}
	}
}

func (s *ManifoldStats) getCollection(ttx *totra.TokenTransaction) *collections.Collection {
	var currentCollection *collections.Collection

	if len(ttx.GetTransfersByContract()) >= 1 && currentCollection == nil {
		currentCollection = tokencollections.GetCollection(s.gb, ttx.Transfers[0].Token.Address, ttx.Transfers[0].Token.ID.Int64())
	}

	return currentCollection
}

// get twitter handle by contract address from alchemy api

func NewManifoldTicker(gb *gloomberg.Gloomberg) *ManifoldStats {
	stats := &ManifoldStats{
		ManifoldEvents: make([]*totra.TokenTransaction, 0),
		gb:             gb,
	}

	Manifold = stats

	getManifoldAddressesFromConfig()

	return Manifold
}

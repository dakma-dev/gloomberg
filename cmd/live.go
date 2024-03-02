package cmd

//  go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=../assets/openapi/opensea.oapi-codegen.yaml ../assets/openapi/opensea.json

import (
	"fmt"
	"net"
	"sort"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func countTopics(logs []*types.Log) (counts map[common.Hash]uint64) {
	counts = make(map[common.Hash]uint64)

	for _, l := range logs {
		for _, t := range l.Topics {
			counts[t]++
		}
	}

	return counts
}

// liveCmd represents the live command.
var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "watch the chain live",

	Run: func(cmd *cobra.Command, args []string) {
		// configure logging
		switch {
		case viper.GetBool("log.debug"):
			log.SetLevel(log.DebugLevel)
		case viper.GetBool("log.verbose"):
			log.SetLevel(log.InfoLevel)
		default:
			log.SetLevel(log.WarnLevel)
		}

		runGloomberg(cmd, args)
	},
}

func runGloomberg(_ *cobra.Command, _ []string) {
	termenv.DefaultOutput().ClearScreen()

	// print header
	header := style.GetHeader(internal.GloombergVersion)
	fmt.Println(header)

	// TODO print loaded configuration etc...
	log.Printf("  log level: %+v", log.DefaultStyles().Levels[log.GetLevel()].Render(log.GetPrefix()))

	gb := gloomberg.New()
	showBlockStats := false

	log.Printf("  gb: %+v", gb)

	newTransactions, err := gb.ChaWa.SubscribeToNFTTransfers()
	if err != nil {
		log.Errorf("❌ error subscribing to transfer logs: %s", err)
	}

	log.Printf("  newTransactions: %+v", newTransactions)

	//
	// queue for everything to print to the console
	// terminalPrinterQueue := make(chan string, viper.GetInt("gloomberg.eventhub.inQueuesSize"))

	go func() {
		var blockNumber atomic.Uint64

		allBlockTopics := make([]*types.Log, 0)

		for tx := range newTransactions {
			txBlockNumber := tx.BlockNumber.Uint64()

			//
			// block stats
			if showBlockStats {
				if currentBlockNumber := blockNumber.Load(); txBlockNumber > currentBlockNumber {
					blockNumber.Store(txBlockNumber)

					if currentBlockNumber > 0 {
						// sort allBlockTopics map by Value
						countedTopics := countTopics(allBlockTopics)
						// create a slice of keys
						topics := make([]common.Hash, 0, len(allBlockTopics))
						for k := range countedTopics {
							topics = append(topics, k)
						}

						// sort the slice of keys
						sort.Slice(topics, func(i, j int) bool {
							return countedTopics[topics[i]] > countedTopics[topics[j]]
						})

						// print sorted topics
						for _, topic := range topics {
							topicName := topic.Hex()
							// if eventSignature, err := external.GetEventSignature(topic); err == nil && eventSignature != nil {
							// 	// pretty.Println(eventSignature)
							// 	topicName = eventSignature.TextSignature
							// }

							log.Infof("  %s | %d", topicName, countedTopics[topic])
						}

						// previous block stats
						log.Infof("🧱 block %d | ...\n", currentBlockNumber)
					}

					allBlockTopics = make([]*types.Log, 0)

					// // new block
					// log.Infof("🧱 new block: %+v", txBlockNumber)
				}

				allBlockTopics = append(allBlockTopics, tx.Receipt.Logs...)
			}

			log.Debugf("📝 %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().String(), tx.Hash().String()))
			log.Infof("📝 %s | %s", tx.GetFunctionSignature().TerminalLinkShortAndStyled(), tx.TerminalLinkShortAndStyled())

			if tx.To().Hex() == common.HexToAddress("0x00000000000000adc04c56bf30ac9d3c0aaf14dc").Hex() {
				log.Infof("📝 OS!! %s", tx.Value().String())

				// contractABI, err := abi.JSON(strings.NewReader(chainwatcher.GetLocalABI("abis/0x00000000000000adc04c56bf30ac9d3c0aaf14dc.json")))
				// if err != nil {
				//	log.Fatal(err)
				//}

				// chainwatcher.DecodeTransactionInputData(&contractABI, tx.Data())
			}
			// //
			// // create a TokenTransaction
			// if ttx := totra.NewTokenTransaction(tx.Transaction, tx.Receipt, chaWa.Node(), gb.Rueidica()); ttx != nil && ttx.IsMovingNFTs() {
			// 	// np.QueueTokenTransactions <- ttx

			// 	gb.In.TokenTransactions <- ttx

			// 	// // publish ttx via redis
			// 	// if viper.GetBool("pubsub.sales.publish") {
			// 	// 	go pusu.Publish(np.gb, internal.PubSubChannelSales, ttx)
			// 	// }
			// }

			// chaWa.Node().LastLogReceivedAt = time.Now()
		}
	}()

	// // old printer active until fully migrated
	// go func() {
	// 	log.Debug("starting OLD terminal printer...")

	// 	for eventLine := range terminalPrinterQueue {
	// 		log.Debugf("OLD terminal printer eventLine: %s", eventLine)

	// 		if viper.GetBool("log.debug") {
	// 			debugPrefix := fmt.Sprintf("%d | ", len(terminalPrinterQueue))
	// 			eventLine = fmt.Sprint(debugPrefix, eventLine)
	// 		}

	// 		fmt.Println(eventLine)
	// 	}
	// }()

	//
	// gasline ticker
	var gasLineTicker *time.Ticker

	// if tickerInterval := viper.GetDuration("ticker.gasline"); gb.ProviderPool != nil && gb.ProviderPool.PreferredProviderAvailable() && tickerInterval > 0 {
	if tickerInterval := viper.GetDuration("ticker.gasline"); gb.Node() != nil && tickerInterval > 0 {
		// initial startup delay
		time.Sleep(tickerInterval / 5)

		// start gasline ticker
		gasLineTicker = time.NewTicker(tickerInterval)
		go gloomberg.GasLineTicker(gb, gasLineTicker, gb.PrintToTerminal)
	}

	//
	// statsbox
	gb.Stats = gloomberg.NewStats(gb, gasLineTicker, gb.OwnWallets, gb.GetRueidica().Rueidis())

	// if statsInterval := viper.GetDuration("ticker.statsbox"); viper.GetBool("stats.enabled") {
	if viper.GetBool("stats.enabled") {
		go gb.Stats.StartTicker(viper.GetDuration("ticker.statsbox"), gb.PrintToTerminal)
	}

	// //
	// // subscribe to redis pubsub channel to receive events from gloomberg central
	// if viper.GetBool("pubsub.client.enabled") {
	// 	gloomberg.Prf("starting redis pubsub client...")

	// 	// subscribe to redis pubsub channel
	// 	go pusu.SubscribeToListingsViaRedis(gb)

	// 	// initially send all our slugs & events to subscribe to
	// 	go gb.PublishOwnSlubSubscription()

	// 	// subscribe to redis pubsub mgmt channel to listen for "SendSlugs" events
	// 	go func() {
	// 		err := gb.Rueidica().Rueidis().Receive(context.Background(), gb.Rueidica().Rueidis().B().Subscribe().Channel(internal.PubSubSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
	// 			log.Debug(fmt.Sprintf("👔 received msg on %s: %s", msg.Channel, msg.Message))

	// 			var mgmtEvent *seawaModels.MgmtEvent

	// 			if err := json.Unmarshal([]byte(msg.Message), &mgmtEvent); err != nil {
	// 				log.Fatal(fmt.Sprintf("❌ error json.Unmarshal: %+v", err))
	// 			}

	// 			if mgmtEvent.Action == seawaModels.SendSlugs {
	// 				log.Info(fmt.Sprintf("👔 SendSlugs received on channel %s", msg.Channel))
	// 				gb.PublishOwnSlubSubscription()
	// 			}
	// 		})
	// 		if err != nil {
	// 			log.Errorf("❌ error subscribing to redis channels %s: %s", internal.PubSubSeaWatcherMgmt, err.Error())

	// 			return
	// 		}
	// 	}()
	// }

	// loop forever
	select {}
}

var degendataPath string

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(liveCmd)

	// main
	liveCmd.Flags().Bool("watch-sales", true, "get sales")
	_ = viper.BindPFlag("sales.enabled", liveCmd.Flags().Lookup("watch-sales"))
	liveCmd.Flags().Bool("watch-listings", false, "get (opensea) listings for own collections")
	_ = viper.BindPFlag("listings.enabled", liveCmd.Flags().Lookup("watch-listings"))

	// websockets server
	liveCmd.Flags().Bool("websockets", false, "enable websockets server")
	_ = viper.BindPFlag("websockets.server.enabled", liveCmd.Flags().Lookup("websockets"))

	liveCmd.Flags().IP("websockets-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	_ = viper.BindPFlag("websockets.server.host", liveCmd.Flags().Lookup("websockets-host"))
	liveCmd.Flags().Uint16("websockets-port", 42068, "websockets server port")
	_ = viper.BindPFlag("websockets.server.port", liveCmd.Flags().Lookup("websockets-port"))

	// metrics/prometheus
	liveCmd.Flags().Bool("metrics", false, "enable metrics server")
	_ = viper.BindPFlag("metrics.enabled", liveCmd.Flags().Lookup("metrics"))
	liveCmd.Flags().IP("metrics-host", net.IPv4(0, 0, 0, 0), "metrics listen address")
	_ = viper.BindPFlag("metrics.host", liveCmd.Flags().Lookup("metrics-host"))
	liveCmd.Flags().Uint16("metrics-port", 9090, "metrics server port")
	_ = viper.BindPFlag("metrics.port", liveCmd.Flags().Lookup("metrics-port"))

	// notifications
	liveCmd.Flags().Bool("telegram", false, "send telegram notifications")
	_ = viper.BindPFlag("notifications.telegram.enabled", liveCmd.Flags().Lookup("telegram"))

	liveCmd.Flags().Bool("manifold-notifications", false, "send manifold notifications")
	_ = viper.BindPFlag("notifications.manifold.enabled", liveCmd.Flags().Lookup("manifold-notifications"))

	// no ui
	liveCmd.Flags().Bool("headless", false, "run without terminal output")
	_ = viper.BindPFlag("ui.headless", liveCmd.Flags().Lookup("headless"))

	// web ui
	liveCmd.Flags().Bool("web-ui", false, "enable web ui")
	_ = viper.BindPFlag("web.enabled", liveCmd.Flags().Lookup("web-ui"))
	liveCmd.Flags().IP("web-ui-host", net.IPv4(0, 0, 0, 0), "web ui listen address")
	_ = viper.BindPFlag("web.host", liveCmd.Flags().Lookup("web-ui-host"))
	liveCmd.Flags().Uint16("web-ui-port", 42069, "web ui port")
	_ = viper.BindPFlag("web.port", liveCmd.Flags().Lookup("web-ui-port"))

	// // wallets
	// liveCmd.Flags().StringSliceVarP(&ownWallets, "wallets", "w", []string{}, "Own wallet addresses")
	// _ = viper.BindPFlag("wallets", liveCmd.Flags().Lookup("wallets"))

	// min value for sales to be shown (single item price)
	liveCmd.Flags().Float64("min-value", 0.0, "minimum value to show sales")
	_ = viper.BindPFlag("show.min_value", liveCmd.Flags().Lookup("min-value"))
	// multiplier for min_value applied to the total price of a tx
	viper.SetDefault("show.min_value_multiplier", 2.0)

	// what to show
	liveCmd.Flags().Bool("show-mints", false, "Show mints")
	_ = viper.BindPFlag("show.mints", liveCmd.Flags().Lookup("show-mints"))
	liveCmd.Flags().Bool("show-burns", false, "Show burns")
	_ = viper.BindPFlag("show.burns", liveCmd.Flags().Lookup("show-burns"))
	liveCmd.Flags().Bool("show-reburns", false, "Show re-burns")
	_ = viper.BindPFlag("show.reburns", liveCmd.Flags().Lookup("show-reburns"))
	liveCmd.Flags().Bool("show-airdrops", false, "Show airdrops")
	_ = viper.BindPFlag("show.airdrops", liveCmd.Flags().Lookup("show-airdrops"))
	liveCmd.Flags().Bool("show-transfers", false, "Show transfers")
	_ = viper.BindPFlag("show.transfers", liveCmd.Flags().Lookup("show-transfers"))
	liveCmd.Flags().Bool("show-unknown", false, "Show unknown")
	_ = viper.BindPFlag("show.unknown", liveCmd.Flags().Lookup("show-unknown"))

	// degendb
	liveCmd.Flags().StringVar(&degendataPath, "degendata", "degendata", "path to degendata repo")
	_ = viper.BindPFlag("degendata.path", liveCmd.Flags().Lookup("degendata"))

	// worker settings
	viper.SetDefault("trapri.numOpenSeaEventhandlers", 3)

	// eventhub
	viper.SetDefault("gloomberg.terminalPrinter.numWorker", 1)
	viper.SetDefault("gloomberg.eventhub.numHandler", 3)
	viper.SetDefault("gloomberg.eventhub.inQueuesSize", 512)
	viper.SetDefault("gloomberg.eventhub.outQueuesSize", 32)

	// first txs
	viper.SetDefault("gloomberg.firstTxs.min_value", 0.337)

	// job runner
	viper.SetDefault("jobs.numRunner", 3)
	viper.SetDefault("jobs.defaults.intervals", map[string]time.Duration{
		"opensea":   time.Millisecond * 7730,
		"etherscan": time.Millisecond * 3370,
		"node":      time.Millisecond * 337,
	})
	viper.SetDefault("jobs.status_every", 1337)

	// OLD worker settings OLD
	viper.SetDefault("server.workers.newHeadHandler", 2)
	viper.SetDefault("server.workers.newLogHandler", 6)
	viper.SetDefault("server.workers.ttxFormatter", 6)
	viper.SetDefault("server.workers.subscription_logs", 2)
	viper.SetDefault("server.workers.listings", 2)

	viper.SetDefault("server.pubsub.listings", 3)
	viper.SetDefault("server.workers.pubsub.listings", 2)

	// opensea settings
	viper.SetDefault("seawatcher.auto_subscribe_after_sales", 37)

	//
	// timeframes

	// sali.default_timeframe is used for the sales/listings counts shown
	viper.SetDefault("salira.default_timeframe", time.Minute*137)
	viper.SetDefault("salira.timeframes", []time.Duration{
		// time.Minute * 1,
		// time.Minute * 3,
		time.Minute * 13,
		time.Minute * 37,
		time.Minute * 137,
	})

	// ticker
	viper.SetDefault("ticker.statsbox", internal.BlockTime*9)
	viper.SetDefault("ticker.gasline", internal.BlockTime*3)

	// stats settings
	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.timeframe", time.Minute*13) // 13
	viper.SetDefault("stats.lines", 6)
	// high volume mints detection
	viper.SetDefault("stats.high_volume.check_interval", time.Second*17)
	viper.SetDefault("stats.high_volume.min_checks_below_threshold", 2)
	viper.SetDefault("stats.high_volume.mints.threshold", 47)
}

// func GetWalletTokens(gb *gloomberg.Gloomberg) map[common.Address]*token.Token {
// 	gbTokens := make([]*token.Token, 0)

// 	for _, w := range *gb.OwnWallets {
// 		tokensForWallet := opensea.GetTokensFor(w.Address, 2, "")
// 		gbTokens = append(gbTokens, tokensForWallet...)

// 		log.Debugf("Wallet %s has %d tokens: %+v", w.Address.String(), len(tokensForWallet), tokensForWallet)

// 		tokenMapForWallet := make(map[common.Address]map[string]*token.Token)
// 		for _, t := range tokensForWallet {
// 			if _, ok := tokenMapForWallet[t.Address]; !ok {
// 				tokenMapForWallet[t.Address] = make(map[string]*token.Token)
// 			}

// 			tokenMapForWallet[t.Address][t.ID.String()] = t
// 		}

// 		w.Tokens = tokenMapForWallet

// 		// honor the rate limit
// 		time.Sleep(time.Millisecond * 337)
// 	}

// 	gloomberg.PrMod("wawa", fmt.Sprintf("found %s tokens in our %s wallets", style.AlmostWhiteStyle.Render(strconv.Itoa(len(gbTokens))), style.AlmostWhiteStyle.Render(strconv.Itoa(len(*gb.OwnWallets)))))

// 	// create map
// 	gbTokensMap := make(map[common.Address]*token.Token)
// 	for _, t := range gbTokens {
// 		gbTokensMap[t.Address] = t
// 	}

// 	return gbTokensMap
// }

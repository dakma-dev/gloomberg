package cmd

import (
	"fmt"
	"math"
	"net"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/server/node"
	"github.com/benleb/gloomberg/internal/server/ws"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/server"
	"github.com/benleb/gloomberg/internal/server/subscribe"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/wwatcher"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cWatcher *server.ChainWatcher

// liveCmd represents the live command.
var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: live, // func(cmd *cobra.Command, args []string) { live() },
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(liveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// liveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// liveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// stats
	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.extended", false)
	viper.SetDefault("stats.gas", true)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.lines", 5)
	viper.SetDefault("stats.interval", time.Second*90)

	// worker settings
	viper.SetDefault("workers.log_handler", 7)
	viper.SetDefault("workers.listings_handler", 4)
	viper.SetDefault("workers.output", 8)

	// opensea settings
	viper.SetDefault("opensea.auto_list_min_sales", 50000)

	// number of retries to resolve an ens name to an address or vice versa
	viper.SetDefault("ens.resolve_max_retries", 5)

	// default collections and wallet
	viper.SetDefault("collections", []any{map[string]string{"name": "OSF's Red Lite District", "mark": "#6A0F27", "address": "0x513cd71defc801b9c1aa763db47b5df223da77a2"}})
	viper.SetDefault("wallets", []string{"pranksy.eth"})
	viper.SetDefault("wwatcher", []any{map[string]string{}})

	// show desktop notifications
	liveCmd.Flags().Bool("notifications", false, "Show notifications?")
	_ = viper.BindPFlag("show.notifications", liveCmd.Flags().Lookup("notifications"))

	// types of events to show
	liveCmd.Flags().Bool("sales", true, "Show sales?")
	_ = viper.BindPFlag("show.sales", liveCmd.Flags().Lookup("sales"))
	liveCmd.Flags().Bool("mints", false, "Show mints?")
	_ = viper.BindPFlag("show.mints", liveCmd.Flags().Lookup("mints"))
	liveCmd.Flags().Bool("listings", false, "Show listings?")
	_ = viper.BindPFlag("show.listings", liveCmd.Flags().Lookup("listings"))
	liveCmd.Flags().Bool("transfers", false, "Show transfers?")
	_ = viper.BindPFlag("show.transfers", liveCmd.Flags().Lookup("transfers"))
	liveCmd.Flags().Float64("min-price", 0.0, "Minimum price to show sales?")
	_ = viper.BindPFlag("show.min_price", liveCmd.Flags().Lookup("min-price"))

	// process & show just our own collections (from config/wallet)
	liveCmd.Flags().Bool("own", false, "Show only own collections (from config/wallet)")
	_ = viper.BindPFlag("show.own", liveCmd.Flags().Lookup("own"))

	// websockets server
	liveCmd.Flags().Bool("ws", false, "enable websockets server")
	liveCmd.Flags().IP("ws-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	liveCmd.Flags().Uint16("ws-port", 42069, "websockets server port")
	_ = viper.BindPFlag("server.websockets.enabled", liveCmd.Flags().Lookup("ws"))
	_ = viper.BindPFlag("server.websockets.host", liveCmd.Flags().Lookup("ws-host"))
	_ = viper.BindPFlag("server.websockets.port", liveCmd.Flags().Lookup("ws-port"))

	// telegram bot
	liveCmd.Flags().Bool("telegram", false, "Start telegram bot")
	_ = viper.BindPFlag("telegram.enabled", liveCmd.Flags().Lookup("telegram"))

	// ticker
	viper.SetDefault("ticker.statsbox", time.Second*89)
	viper.SetDefault("ticker.gasline", time.Second*13)

	viper.SetDefault("server.websockets.enabled", true)

	// ipfs gateway to fetch metadata/images
	viper.SetDefault("ipfs.gateway", "https://ipfs.io/ipfs/")
}

// wtf fix this
var queues = make(map[string]interface{})

func live(_ *cobra.Command, _ []string) {
	gbl.GetSugaredLogger()

	//
	// config
	if viper.GetFloat64("show.min_price") > 0.0 {
		numGlickerLines := math.Max(float64(viper.GetInt("stats.lines")), 4.0)
		viper.Set("stats.lines", numGlickerLines)
	}

	// show listings for own collections if an opensea api key is set
	if viper.IsSet("api_keys.opensea") && !viper.IsSet("show.listings") {
		viper.Set("show.listings", true)
	}

	// wallet watcher
	wwatcher.InitWatcher()

	// write settings to log
	gbl.Log.Debug(viper.AllSettings())

	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	// create a queue/channel for the received chain logs
	queueLogs := make(chan types.Log, 1024)
	queues["logs"] = &queueLogs
	// create a queue/channel for the received listings
	queueListings := make(chan *models.ItemListedEvent, 1024)
	queues["listings"] = &queueListings
	// create a queue/channel for the received & parsed events
	queueEvents := make(chan *collections.Event, 1024)
	queues["events"] = &queueEvents
	// create a queue/channel for events to be sent out via ws
	queueOutWS := make(chan *collections.Event, 1024)
	queues["outWS"] = &queueOutWS

	// create a new chainserver instance
	cWatcher = server.New()
	// subscribe to the chain logs/events and start the workers
	cWatcher.Subscribe(&queueEvents)
	// old type (while migrating to new snode.Node type (will be renamed from snode to node after))
	// nodes := cWatcher.GetNodesAsGBNodeCollection()

	// websockets server
	if viper.GetBool("server.websockets.enabled") {
		// websockets
		wsServer := ws.New(viper.GetString("server.websockets.host"), viper.GetUint("server.websockets.port"), &queueOutWS)
		go wsServer.Start()
	}

	//
	// initialize wallets
	wallets := getWallets(&cWatcher.Nodes)

	// gasline ticker
	var gasTicker *time.Ticker

	if tickerInterval := viper.GetDuration("ticker.gasline"); len(cWatcher.Nodes.GetLocalNodes()) > 0 && tickerInterval > 0 {
		// initial startup delay
		time.Sleep(tickerInterval / 5)

		// start gasline ticker
		gasTicker = time.NewTicker(tickerInterval)
		go ticker.GasTicker(gasTicker, &cWatcher.Nodes, &queueOutput)
	}

	//
	// initialize the statistics
	stats = ticker.New(gasTicker, wallets, &cWatcher.Nodes, len(cWatcher.OwnCollections.UserCollections))

	// start status indicator ticker
	if statsInterval := viper.GetDuration("stats.interval"); viper.GetBool("stats.enabled") {
		stats.StartTicker(statsInterval)
	}

	//
	// MIWs
	miwSpinner := style.GetSpinner("setting up MIWs...")
	_ = miwSpinner.Start()

	wwatcher.LoadMIWs()

	if len(wwatcher.MIWC.WeightedMIWs) > 0 {
		miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(wwatcher.MIWC.WeightedMIWs))), " MIWs loaded", "\n")))
		_ = miwSpinner.Stop()
	} else {
		_ = miwSpinner.StopFail()
	}

	//
	// initialize collections
	collectionsSpinner := style.GetSpinner("setting up collections...")
	_ = collectionsSpinner.Start()

	// collection from config file
	collectionsSpinner.Message("setting up config collections...")

	// read collections from config and store in ownCollections
	configCollections := collections.GetCollectionsFromConfiguration(&cWatcher.Nodes)
	configCollectionsNumber := len(configCollections)
	gbl.Log.Infof("collections from config: %d", configCollectionsNumber)

	for _, collection := range configCollections {
		cWatcher.OwnCollections.UserCollections[collection.ContractAddress] = collection
	}

	// collections from wallet holdings
	collectionsSpinner.Message("setting up wallet collections...")

	// read collections hold in wallets from opensea and store in currentCollections
	walletCollections := opensea.GetWalletCollections(*wallets, cWatcher.OwnCollections, &cWatcher.Nodes)

	for _, collection := range walletCollections {
		if cWatcher.OwnCollections.UserCollections[collection.ContractAddress] == nil {
			cWatcher.OwnCollections.UserCollections[collection.ContractAddress] = collection
		}
	}

	gbl.Log.Infof("collections from wallets: %d", len(cWatcher.OwnCollections.UserCollections)-configCollectionsNumber)

	// print collections from config & wallet holdings
	if len(cWatcher.OwnCollections.UserCollections) > 0 {
		collectionNames := cWatcher.OwnCollections.SortedAndColoredNames()
		collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections: ", strings.Join(collectionNames, ", "), "\n\n"))
	}

	_ = collectionsSpinner.Stop()

	//
	// specialized subscriptions
	if viper.GetBool("show.own") {
		for _, cNode := range cWatcher.Nodes {
			// subscribe to all "Transfer" events
			if _, err := cNode.SubscribeToTransfersFor(queueLogs, cWatcher.OwnCollections.Addresses()); err != nil {
				gbl.Log.Warnf("TransfersFor subscribe failed: %s", err)
			}

			// subscribe to all "SingleTransfer" events
			if _, err := cNode.SubscribeToSingleTransfersFor(queueLogs, cWatcher.OwnCollections.Addresses()); err != nil {
				gbl.Log.Warnf("SingleTransfersFor subscribe failed: %s", err)
			}
		}
	}

	//
	// format events and print them to stdout
	for outputWorkerID := 1; outputWorkerID <= viper.GetInt("workers.output"); outputWorkerID++ {
		// format events from queueEvents in a pretty way for terminal (and later other "outputs")
		go workerEventFormatter(outputWorkerID, &cWatcher.Nodes, wallets, &queueEvents, &queueOutput, &queueOutWS)

		// print formatted events to stdout
		go workerOutput(outputWorkerID, &queueOutput)
	}

	// processes new listings from the opensea stream api
	for listingsWorkerID := 1; listingsWorkerID <= viper.GetInt("workers.listings_handler"); listingsWorkerID++ {
		go subscribe.StreamListingsHandler(listingsWorkerID, cWatcher.OwnCollections, &queueListings, &queueEvents)
	}

	//
	// subscribe to sales on the stream api for all collections discovered in wallets and configuration
	if openseaToken := viper.GetString("api_keys.opensea"); openseaToken != "" {
		go func() {
			if client := opensea.NewStreamClient(openseaToken, func(err error) {
				gbl.Log.Error(err)
			}); client != nil {
				for _, collection := range cWatcher.OwnCollections.UserCollections {
					gbl.Log.Debugf("%s: collection.Show.Listings: %v | collection.OpenseaSlug: %s", collection.Name, collection.Show.Listings, collection.OpenseaSlug)

					if collection.Show.Listings {
						if collection.OpenseaSlug == "" {
							if slug := opensea.GetCollectionSlug(collection.ContractAddress); slug != "" {
								collection.OpenseaSlug = slug
							} else {
								gbl.Log.Warnf("no slug for collection %s", collection.ContractAddress)

								continue
							}
						}

						opensea.SubscribeToCollectionSlug(client, collection.OpenseaSlug, streamListingsReceiver)
						time.Sleep(337 * time.Millisecond)
					}
				}
			}
		}()
	}

	QueueStatsLive(&queueEvents, queueLogs, queueListings, &queueOutput, &queueOutWS)

	select {}
}

func streamListingsReceiver(response any) {
	var itemListedEvent models.ItemListedEvent

	err := mapstructure.Decode(response, &itemListedEvent)
	if err != nil {
		gbl.Log.Error("mapstructure.Decode failed for incoming stream api event", err)
	}

	gbl.Log.Debugf("%+v", itemListedEvent)

	// var queueLising chan *models.ItemListedEvent = queues["listings"].(*chan *models.ItemListedEvent)
	*queues["listings"].(*chan *models.ItemListedEvent) <- &itemListedEvent
}

// queueEvents -> queues for Output
// workerEventFormatter formats sale/listing events from queueEvents and pushes them to the queueOutput channel.
func workerEventFormatter(workerID int, nodes *node.Nodes, wallets *wallet.Wallets, queueEvents *chan *collections.Event, queueOutput *chan string, queueOutWS *chan *collections.Event) {
	gbl.Log.Infof("workerEventFormatter %d/%d started", workerID, viper.GetInt("workers.log_handler"))

	for event := range *queueEvents {
		gbl.Log.Debugf("%d ~ %d | workerEventFormatter event: %v", workerID, len(*queueEvents), event)

		// atomic.AddUint64(&stats.queueEvents, 1)
		go formatEvent(nil, event, nodes, wallets, queueOutput, queueOutWS)
	}
}

// queueOutput -> stdout
// workerOutput actually prints formatted lines from queueOutput to stdout.
func workerOutput(workerID int, queueOutput *chan string) {
	gbl.Log.Infof("workerOutput %d/%d started", workerID, viper.GetInt("workers.output"))

	for outputLine := range *queueOutput {
		gbl.Log.Debugf("%d ~ %d | workerOutput outputLine: %s", workerID, len(*queueOutput), outputLine)

		if viper.GetBool("log.debug") {
			outputLine = fmt.Sprintf("%s ~ %d | %s", style.BoldStyle.Render(fmt.Sprintf("%d", workerID)), len(*queueOutput), outputLine)
		}

		fmt.Println(outputLine)
	}
}

func QueueStatsLive(queueEvents *chan *collections.Event, queueLogs chan types.Log, queueListings chan *models.ItemListedEvent, queueOutput *chan string, queueOutWS *chan *collections.Event) {
	tickerPrintStats := time.NewTicker(time.Second * 5)
	for range tickerPrintStats.C {
		gbl.Log.Infof("  live > queueEvents: %d | queueLogs: %d | queueListings: %d | queueOutput: %d | queueOutWS: %d <", len(*queueEvents), len(queueLogs), len(queueListings), len(*queueOutput), len(*queueOutWS))
	}
}

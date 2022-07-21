package cmd

import (
	"context"
	"fmt"
	"math"
	"net"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/benleb/gloomberg/internal/glicker"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/notifications"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/subscriptions"
	"github.com/benleb/gloomberg/internal/wwatcher"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiKeyEtherscan, apiKeyOpensea string
	endpoints, ownWallets          []string
)

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
	viper.SetDefault("workers.log_handler", 8)
	viper.SetDefault("workers.listings_handler", 4)
	viper.SetDefault("workers.output", 8)

	// opensea settings
	viper.SetDefault("opensea.auto_list_min_sales", 500)

	// number of retries to resolve an ens name to an address or vice versa
	viper.SetDefault("ens.resolve_max_retries", 9)

	// ipfs gateway to fetch metadata/images
	viper.SetDefault("ipfs.gateway", "https://ipfs.io/ipfs/")

	// default collections and wallet
	viper.SetDefault("collections", []any{map[string]string{"name": "OSF's Red Lite District", "mark": "#6A0F27", "address": "0x513cd71defc801b9c1aa763db47b5df223da77a2"}})
	viper.SetDefault("wallets", []string{"pranksy.eth"})
	viper.SetDefault("wwatcher", []any{map[string]string{}})

	// logging
	liveCmd.PersistentFlags().BoolP("verbose", "v", false, "Show more output")
	_ = viper.BindPFlag("log.verbose", liveCmd.PersistentFlags().Lookup("verbose"))
	liveCmd.PersistentFlags().BoolP("debug", "d", false, "Show debug output")
	_ = viper.BindPFlag("log.debug", liveCmd.PersistentFlags().Lookup("debug"))

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

	// process & show all collection (we get events for) or just our own (from config/wallet)
	liveCmd.Flags().Bool("all", false, "Show all collections, not just own ones (from config & wallet)")
	_ = viper.BindPFlag("show.all", liveCmd.Flags().Lookup("all"))
	liveCmd.Flags().Bool("own-only", false, "Show own collections only (from config/wallet)")
	_ = viper.BindPFlag("show.own_only", liveCmd.Flags().Lookup("own-only"))

	// websockets server
	liveCmd.Flags().Bool("server", false, "Start websockets server")
	_ = viper.BindPFlag("server.enabled", liveCmd.Flags().Lookup("server"))
	liveCmd.Flags().IP("host", net.IPv4(0, 0, 0, 0), "Websockets server port")
	_ = viper.BindPFlag("server.host", liveCmd.Flags().Lookup("host"))
	liveCmd.Flags().Uint16("port", 42069, "Websockets server port")
	_ = viper.BindPFlag("server.port", liveCmd.Flags().Lookup("port"))

	// telegram bot
	liveCmd.Flags().Bool("telegram", false, "Start telegram bot")
	_ = viper.BindPFlag("telegram.enabled", liveCmd.Flags().Lookup("telegram"))

	// rpc node
	liveCmd.Flags().StringSliceVar(&endpoints, "endpoints", []string{}, "RPC endpoints")
	_ = viper.BindPFlag("endpoints", liveCmd.Flags().Lookup("endpoints"))

	// wallets
	liveCmd.Flags().StringSliceVar(&ownWallets, "wallets", []string{}, "Own wallet addresses")
	_ = viper.BindPFlag("wallets", liveCmd.Flags().Lookup("wallets"))

	// apis
	liveCmd.Flags().StringVar(&apiKeyEtherscan, "etherscan", "", "Etherscan API Key")
	_ = viper.BindPFlag("api_keys.etherscan", liveCmd.Flags().Lookup("etherscan"))
	liveCmd.Flags().StringVar(&apiKeyOpensea, "opensea", "", "Opensea API Key")
	_ = viper.BindPFlag("api_keys.opensea", liveCmd.Flags().Lookup("opensea"))
}

func live(_ *cobra.Command, _ []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*420*10))
	defer cancel()

	gbl.GetSugaredLogger()

	var (
		// wallets *models.Wallets
		stats *glicker.Stats

		ownCollections = collections.New()
	)

	//
	// config
	//
	if viper.GetString("server.host") != net.IPv4(0, 0, 0, 0).String() || viper.GetUint("server.port") != 42069 {
		viper.Set("server.enabled", true)
	}

	if viper.GetFloat64("show.min_price") > 0.0 {
		numGlickerLines := math.Max(float64(viper.GetInt("stats.lines")), float64(4.0))
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
	headerLogo := style.GetHeader()
	headerVersion := style.DarkGrayStyle.Copy().PaddingBottom(3).Render(fmt.Sprintf("                       gloomberg %s (%s) | %s", Version, Commit, BuildDate))
	header := headerLogo + "\n" + headerVersion

	fmt.Println(header)
	gbl.Log.Info(header)

	//
	// initialize node connections
	//
	nodes := getNodes()

	// if we subscribe to all chain-events, we can do it now
	if viper.GetBool("show.all") {
		nodes.SubscribeToAllTransfers(ctx, queueLogs)
	}

	//
	// initialize wallets
	//
	wallets := getWallets(nodes)

	//
	// initialize collections
	//
	collectionsSpinner := style.GetSpinner("setting up collections...")
	_ = collectionsSpinner.Start()

	// collection from config file
	collectionsSpinner.Message("setting up config collections...")

	// read collections from config and store in ownCollections
	configCollections := collections.GetCollectionsFromConfiguration(nodes)
	configCollectionsNumber := len(configCollections) // config.ReadCollectionsFromConfig(ctx, ownCollections)
	gbl.Log.Infof("collections from config: %d", configCollectionsNumber)

	for _, collection := range configCollections {
		ownCollections.UserCollections[collection.ContractAddress] = collection
	}

	// collections from wallet holdings
	collectionsSpinner.Message("setting up wallet collections...")

	// read collections hold in wallets from opensea and store in currentCollections
	walletCollections := opensea.GetWalletCollections(*wallets, ownCollections, nodes)

	for _, collection := range walletCollections {
		if ownCollections.UserCollections[collection.ContractAddress] == nil {
			ownCollections.UserCollections[collection.ContractAddress] = collection
		}
	}

	gbl.Log.Infof("collections from wallets: %d", len(ownCollections.UserCollections)-configCollectionsNumber)

	// print collections from config & wallet holdings
	if len(ownCollections.UserCollections) > 0 {
		collectionNames := ownCollections.SortedAndColoredNames()
		collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections: ", strings.Join(collectionNames, ", "), "\n\n"))
	}

	_ = collectionsSpinner.Stop()

	// specialized subscriptions
	if !viper.GetBool("show.all") {
		for _, node := range nodes.GetNodes() {
			// subscribe to all "Transfer" events
			if _, err := node.SubscribeToTransfersFor(ctx, queueLogs, ownCollections.Addresses()); err != nil {
				gbl.Log.Warnf("TransfersFor subscribe failed: ", err)
			}

			// subscribe to all "SingleTransfer" events
			if _, err := node.SubscribeToSingleTransfersFor(ctx, queueLogs, ownCollections.Addresses()); err != nil {
				gbl.Log.Warnf("SingleTransfersFor subscribe failed: ", err)
			}
		}
	}

	//
	// initialize the statistics
	//
	stats = glicker.New(ctx, wallets, len(ownCollections.UserCollections))

	//
	// set up workers print to process events from the chain
	//

	// receives formatted lines and prints them to stdout
	for outputWorkerID := 1; outputWorkerID <= viper.GetInt("workers.output"); outputWorkerID++ {
		go workerEventFormatter(ctx, outputWorkerID, nodes, wallets, &queueEvents, &queueOutput)
		go workerOutput(outputWorkerID, &queueOutput)
	}

	// processes logs from the ethereum chain from our nodes
	for workerID := 1; workerID <= viper.GetInt("workers.log_handler"); workerID++ {
		go subscriptions.SubscriptionLogsHandler(nodes, ownCollections, queueLogs, queueEvents)
	}

	// processes new listings from the opensea stream api
	for listingsWorkerID := 1; listingsWorkerID <= viper.GetInt("workers.listings_handler"); listingsWorkerID++ {
		go subscriptions.StreamListingsHandler(listingsWorkerID, ownCollections, &queueListings, &queueEvents)
	}

	// subscribe to sales on the stream api for all collections discovered in wallets and configuration
	if openseaToken := viper.GetString("api_keys.opensea"); openseaToken != "" {
		go func() {
			if client := opensea.NewStreamClient(openseaToken, func(err error) {
				gbl.Log.Error(err)
			}); client != nil {
				for _, collection := range ownCollections.UserCollections {
					if collection.Show.Listings {
						if collection.Metadata.OpenseaSlug == "" {
							if slug := opensea.GetCollectionSlug(collection.ContractAddress); slug != "" {
								collection.Metadata.OpenseaSlug = slug
							} else {
								gbl.Log.Warnf("no slug for collection %s", collection.ContractAddress)

								continue
							}
						}

						opensea.SubscribeToCollectionSlug(client, collection.Metadata.OpenseaSlug, streamListingsReceiver)
						time.Sleep(1370 * time.Millisecond)
					}
				}
			}
		}()
	}

	// start status indicator ticker
	if statsInterval := viper.GetDuration("stats.interval"); viper.GetBool("stats.enabled") {
		stats.StartTicker(statsInterval)
	}

	// telegram bot testing
	if viper.GetBool("telegram.enabled") {
		notifications.InitTelegramBot()
	}

	// // websockets server
	// if viper.GetBool("server.enabled") {
	// 	viper.Set("mode.server", true)
	// 	go server.StartWebsocketServer(&outputWs)
	// }

	select {}
}

func streamListingsReceiver(response any) {
	var itemListedEvent models.ItemListedEvent

	err := mapstructure.Decode(response, &itemListedEvent)
	if err != nil {
		gbl.Log.Error("mapstructure.Decode failed for incoming stream api event", err)
	}

	gbl.Log.Debugf("%+v", itemListedEvent)

	queueListings <- &itemListedEvent
}

// queueEvents -> queueOutput
// workerEventFormatter formats sale/listing events from queueEvents and pushes them to the queueOutput channel.
func workerEventFormatter(ctx context.Context, workerID int, nodes *gbnode.NodeCollection, wallets *models.Wallets, queueEvents *chan *collections.Event, queueOutput *chan string) {
	gbl.Log.Infof("workerEventFormatter %d/%d started", workerID, viper.GetInt("workers.log_handler"))

	for event := range *queueEvents {
		gbl.Log.Debugf("%s ~ %d | workerEventFormatter event: %s", workerID, len(*queueEvents), event)

		// atomic.AddUint64(&stats.queueEvents, 1)
		formatEvent(ctx, nil, event, nodes, wallets, queueOutput)
	}
}

// queueOutput -> stdout
// workerOutput actually prints formatted lines from queueOutput to stdout.
func workerOutput(workerID int, queueOutput *chan string) {
	gbl.Log.Infof("workerOutput %d/%d started", workerID, viper.GetInt("workers.output"))

	for outputLine := range *queueOutput {
		gbl.Log.Debugf("%s ~ %d | workerOutput outputLine: %s", workerID, len(*queueOutput), outputLine)

		if viper.GetBool("log.debug") {
			outputLine = fmt.Sprintf("%s ~ %d | %s", style.BoldStyle.Render(fmt.Sprintf("%d", workerID)), len(*queueOutput), outputLine)
		}

		// atomic.AddUint64(&stats.queueOutput, 1)
		fmt.Println(outputLine)
	}
}

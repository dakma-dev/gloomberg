package gloomberg

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/jobs"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/nemo/watch"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
)

type Gloomberg struct {
	// Nodes        *nodes.Nodes
	ProviderPool *provider.Pool
	Watcher      *watch.Watcher

	CollectionDB *collections.CollectionDB
	OwnWallets   *wallet.Wallets
	Stats        *Stats

	RecentOwnEvents mapset.Set[*degendb.PreformattedEvent]

	Ranks map[common.Address]map[int64]degendb.TokenRank

	Rdb    rueidis.Client
	Rueidi *rueidica.Rueidica

	QueueSlugs chan common.Address

	CurrentGasPriceGwei   uint64
	CurrentOnlineWebUsers uint64

	*eventHub
	// GloomHub *gloomhub
	Jobs *jobs.Runner
	*degendb.DegenDB

	// misc / maybe find better solutions...
	PrintConfigurations map[string]*printConfig
}

func (gb *Gloomberg) String() {
	fmt.Println("gloomberg | " + internal.GloombergVersion)
}

type printConfig struct {
	Icon     string
	Keywords []string
	Color    lipgloss.Color
}

var predefinedPrintConfigurations = []printConfig{
	{
		Icon:     "🖥️",
		Keywords: []string{"web", "ws"},
		Color:    lipgloss.Color("#662288"),
	},
	{
		Icon:     "💳",
		Keywords: []string{"wawa", "watch"},
		Color:    lipgloss.Color("#550933"),
	},
	{
		Icon:     "🤺",
		Keywords: []string{"ddb", "degendb"},
		Color:    lipgloss.Color("#095533"),
	},
	{
		Icon:     "🧪",
		Keywords: []string{"exp", "e6551", "eip6551"},
		Color:    lipgloss.Color("#5FF3B3"),
	},
	{
		Icon:     "🏃‍♂️",
		Keywords: []string{"jobs", "job"},
		Color:    lipgloss.Color("#4dc6e2"),
	},
}

var GB *Gloomberg

func New() *Gloomberg {
	// redis
	// rueidis / new redis library
	var connectAddr string

	if viper.IsSet("redis.address") {
		connectAddr = viper.GetString("redis.address")
	} else {
		// fallback to old config
		connectAddr = fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	}

	// use hostname as client name
	hostname, err := os.Hostname()
	if err != nil {
		log.Error(fmt.Sprintf("❗️ error getting hostname: %s", err))

		hostname = "unknown"
	}

	clientName := hostname + "_gloomberg_v" + internal.GloombergVersion
	redisClientOptions := rueidis.ClientOption{
		InitAddress: []string{connectAddr},
		ClientName:  clientName,
	}

	rdb := getRedisClient(redisClientOptions)

	gb := &Gloomberg{
		Rdb:    rdb,
		Rueidi: rueidica.NewRueidica(rdb),

		CollectionDB: collections.New(),

		RecentOwnEvents: mapset.NewSet[*degendb.PreformattedEvent](),

		Ranks: make(map[common.Address]map[int64]degendb.TokenRank),

		QueueSlugs: make(chan common.Address, 1024),

		eventHub: newEventHub(),

		// DegenDB:  degendb.NewDegenDB(),
	}

	//
	// start central terminal printer
	// printToTerminalChannel := gb.SubscribePrintToTerminal()

	// experimental: start multiple terminal printers
	for i := 0; i < viper.GetInt("gloomberg.terminalPrinter.numWorker"); i++ {
		go func() {
			gbl.Log.Debug("starting terminal printer...")

			for eventLine := range TerminalPrinterQueue {
				gbl.Log.Debugf("terminal printer eventLine: %s", eventLine)

				if viper.GetBool("log.debug") {
					debugPrefix := fmt.Sprintf("%d | ", len(TerminalPrinterQueue))
					eventLine = fmt.Sprint(debugPrefix, eventLine)
				}

				fmt.Println(eventLine)
			}
		}()
	}

	// load print configurations to pretty style prints from our different "modules"
	gb.PrintConfigurations = make(map[string]*printConfig)

	for idx, config := range predefinedPrintConfigurations {
		printConfiguration := predefinedPrintConfigurations[idx]

		for _, keyword := range config.Keywords {
			gb.PrintConfigurations[keyword] = &printConfiguration
		}
	}

	GB = gb

	return gb
}

func (gb *Gloomberg) PublishOwnCollectionsSlugs() {
	gb.publisheCollectionSlugsViaRedis(gb.CollectionDB.OpenseaSlugs())
}

func (gb *Gloomberg) PublishCollectionsSlugs(slugs []string) {
	gb.publisheCollectionSlugsViaRedis(slugs)
}

func (gb *Gloomberg) publisheCollectionSlugsViaRedis(slugs []string) {
	// to enable multiple users to use the central gloomberg instance for events from opensea,
	// we first send the slugs of 'our' collections to the events-subscriptions channel.
	// the central gloomberg instance then creates a subscription on the opensea
	// api and publishes upcoming incoming events to the pubsub channel
	// marshal event to json
	if !viper.GetBool("pubsub.client.enabled") {
		gbl.Log.Warn("❌ not sending slugs to server - pubsub client is not enabled")

		return
	}

	if len(slugs) == 0 {
		gbl.Log.Warn("❌ no slugs to send to gloomberg server")

		return
	}

	log.Debugf("👔 sending %s collection slugs to gloomberg server", style.BoldStyle.Render(fmt.Sprint(len(slugs))))

	mgmtEvent := &models.MgmtEvent{Action: models.Subscribe, Slugs: slugs}

	switch {
	case viper.GetBool("seawatcher.local"):
		// for local seawatcher
		gb.In.SeawatcherMgmt <- mgmtEvent
	case viper.GetBool("pubsub.client.enabled"):
		jsonMgmtEvent, err := json.Marshal(mgmtEvent)
		if err != nil {
			gbl.Log.Error("❌ marshal failed for outgoing list of collection slugs: %s | %v", err, gb.CollectionDB.OpenseaSlugs())

			return
		}

		if gb.Rdb.Do(context.Background(), gb.Rdb.B().Publish().Channel(internal.PubSubSeaWatcherMgmt).Message(string(jsonMgmtEvent)).Build()).Error() != nil {
			gbl.Log.Warnf("error publishing event to redis: %s", err.Error())
		} else {
			gbl.Log.Infof("👔 sent %s collection slugs to %s", style.BoldStyle.Render(fmt.Sprint(len(slugs))), style.BoldStyle.Render(internal.PubSubSeaWatcherMgmt))
		}
	}
}

func getRedisClient(redisClientOptions rueidis.ClientOption) rueidis.Client {
	// // use hostname as client name
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	log.Error(fmt.Sprintf("❗️ error getting hostname: %s", err))

	// 	hostname = "unknown"
	// }

	// // rueidis / new redis library
	// var connectAddr string

	// if viper.IsSet("redis.address") {
	// 	connectAddr = viper.GetString("redis.address")
	// } else {
	// 	// fallback to old config
	// 	connectAddr = fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	// }

	rdb, err := rueidis.NewClient(redisClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

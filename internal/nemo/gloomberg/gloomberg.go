package gloomberg

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

	RecentOwnEvents mapset.Set[*degendb.ParsedEvent]

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
	redisClientOptions := rueidis.ClientOption{InitAddress: []string{connectAddr}, ClientName: clientName}

	rdb := getRedisClient(redisClientOptions)

	gb := &Gloomberg{
		Rdb:    rdb,
		Rueidi: rueidica.NewRueidica(rdb),

		CollectionDB: collections.New(),

		RecentOwnEvents: mapset.NewSet[*degendb.ParsedEvent](),

		Ranks: make(map[common.Address]map[int64]degendb.TokenRank),

		QueueSlugs: make(chan common.Address, 1024),

		eventHub: newEventHub(),
		// GloomHub: newGloomhub(),
		// DegenDB:  degendb.NewDegenDB(),
	}

	//
	// start central terminal printer
	go func() {
		gbl.Log.Debug("starting terminal printer...")

		printToTerminalChannel := gb.SubscribePrintToTerminal()

		for eventLine := range printToTerminalChannel {
			gbl.Log.Debugf("terminal printer eventLine: %s", eventLine)

			if viper.GetBool("log.debug") {
				debugPrefix := fmt.Sprintf("%d | ", len(printToTerminalChannel))
				eventLine = fmt.Sprint(debugPrefix, eventLine)
			}

			fmt.Println(eventLine)
		}
	}()

	// load print configurations to pretty style prints from our different "modules"
	gb.PrintConfigurations = make(map[string]*printConfig)

	for idx, config := range predefinedPrintConfigurations {
		printConfiguration := predefinedPrintConfigurations[idx]

		for _, keyword := range config.Keywords {
			gb.PrintConfigurations[keyword] = &printConfiguration
		}
	}

	return gb
}

func (gb *Gloomberg) SendSlugsToServer() {
	// to enable multiple users to use the central gloomberg instance for events from opensea,
	// we first send the slugs of 'our' collections to the events-subscriptions channel.
	// the central gloomberg instance then creates a subscription on the opensea
	// api and publishes upcoming incoming events to the pubsub channel
	// marshal event to json
	slugs := gb.CollectionDB.OpenseaSlugs()
	if len(slugs) == 0 {
		gbl.Log.Warn("❌ no slugs to send to gloomberg server")

		return
	}

	log.Debugf("📢 sending %s collection slugs to gloomberg server", style.BoldStyle.Render(fmt.Sprint(len(slugs))))

	mgmtEvent := &models.MgmtEvent{Action: models.Subscribe, Slugs: slugs}

	gb.In.SeawatcherMgmt <- mgmtEvent

	if viper.GetBool("seawatcher.pubsub") {
		jsonMgmtEvent, err := json.Marshal(mgmtEvent)
		if err != nil {
			gbl.Log.Error("❌ marshal failed for outgoing list of collection slugs: %s | %v", err, gb.CollectionDB.OpenseaSlugs())

			return
		}

		if gb.Rdb.Do(context.Background(), gb.Rdb.B().Publish().Channel(internal.TopicSeaWatcherMgmt).Message(string(jsonMgmtEvent)).Build()).Error() != nil {
			gbl.Log.Warnf("error publishing event to redis: %s", err.Error())
		} else {
			gbl.Log.Infof("📢 sent %s collection slugs to %s", style.BoldStyle.Render(fmt.Sprint(len(slugs))), style.BoldStyle.Render(internal.TopicSeaWatcherMgmt))
		}
	}
}

// Pr prints messages from gloomberg to the terminal.
func (gb *Gloomberg) Pr(message string) {
	gb.printToTerminal("🧃", style.Gray5Style.Render("gb"), message) // style.PinkBoldStyle.Render("・"))
}

// Prf formats and prints messages from gloomberg to the terminal.
func (gb *Gloomberg) Prf(format string, a ...any) {
	gb.Pr(fmt.Sprintf(format, a...))
}

func (gb *Gloomberg) PrWarn(message string) {
	gb.printToTerminal("⚠️", "", message)
}

func (gb *Gloomberg) PrWithKeywordAndIcon(icon string, keyword string, message string) {
	gb.printToTerminal(icon, keyword, message)
}

func (gb *Gloomberg) PrMod(mod string, message string) {
	prConfig, ok := gb.PrintConfigurations[mod]
	if !ok {
		gbl.Log.Warnf("no print configuration for module %s | message: %s", mod, message)

		return
	}

	icon := prConfig.Icon

	color := style.DarkGray
	if prConfig.Color != "" {
		color = prConfig.Color
	}

	tag := lipgloss.NewStyle().Foreground(color).Render(mod)

	gb.printToTerminal(icon, tag, message)
}

// PrModf formats and prints messages from gloomberg to the terminal.
func (gb *Gloomberg) PrModf(mod string, format string, a ...any) {
	gb.PrMod(mod, fmt.Sprintf(format, a...))
}

// PrVMod prints messages from gloomberg to the terminal if verbose mode is enabled.
func (gb *Gloomberg) PrVMod(mod string, message string) {
	if viper.GetBool("log.verbose") {
		gb.PrMod(mod, message)
	}
}

// PrVModf formats and prints messages from gloomberg to the terminal if verbose mode is enabled.
func (gb *Gloomberg) PrVModf(mod string, format string, a ...any) {
	if viper.GetBool("log.verbose") {
		gb.PrModf(mod, format, a...)
	}
}

// PrDMod prints messages from gloomberg to the terminal if debug mode is enabled.
func (gb *Gloomberg) PrDMod(mod string, message string) {
	if viper.GetBool("log.debug") {
		gb.PrMod(mod, message)
	}
}

// PrDModf formats and prints messages from gloomberg to the terminal if debug mode is enabled.
func (gb *Gloomberg) PrDModf(mod string, format string, a ...any) {
	if viper.GetBool("log.debug") {
		gb.PrModf(mod, format, a...)
	}
}

func (gb *Gloomberg) printToTerminal(icon string, keyword string, message string) {
	if message == "" {
		return
	}

	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")

	out := strings.Builder{}
	out.WriteString(style.DarkGrayStyle.Render("|"))
	out.WriteString(style.Gray4Style.Render(currentTime))
	out.WriteString(" " + icon)
	out.WriteString(" " + lipgloss.NewStyle().Width(6).Align(lipgloss.Right).Render(keyword))
	out.WriteString("  " + message)

	gb.In.PrintToTerminal <- out.String()
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

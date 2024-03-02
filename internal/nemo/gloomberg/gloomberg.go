package gloomberg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/rueidica"
	seawaModels "github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
)

type Gloomberg struct {
	//
	//  new stuff ↓
	//
	ChaWa    *chainwatcher.ChainWatcher
	Rueidica *rueidica.Rueidica

	PrintToTerminal chan string

	// low level clients
	nodes         models.Nodes
	rueidisClient *rueidis.Client

	// // gloomberg modules
	// ChainWatcher *chainwatcher.ChainWatcher

	//
	//  old stuff ↓
	//

	CollectionDB *collections.CollectionDB
	OwnWallets   *wallet.Wallets
	Stats        *Stats

	RecentOwnEvents mapset.Set[*degendb.PreformattedEvent]

	// Ranks map[common.Address]map[int64]degendb.TokenRank

	// QueueSlugs chan common.Address

	CurrentGasPriceGwei   uint64
	CurrentOnlineWebUsers uint64

	*eventHub
	// GloomHub *gloomhub
	// Jobs *jobs.Runner
	*degendb.DegenDB

	// misc / maybe find better solutions...
	// PrintConfigurations map[string]*printConfig
}

func New() *Gloomberg {
	gb := &Gloomberg{
		PrintToTerminal: make(chan string, 128),
	}

	go gb.startTerminalPrinter()

	gb.nodes = readChainwatcherConfiguration(viper.Sub("chainwatcher"))
	gb.rueidisClient = loadRedisConfiguration(viper.Sub("redis"))

	gb.ChaWa = chainwatcher.New(gb.nodes)
	gb.Rueidica = rueidica.NewRueidica(gb.rueidisClient)

	return gb
}

func (gb *Gloomberg) Node() *models.Node {
	return gb.ChaWa.Node()
}

func (gb *Gloomberg) GetRueidica() *rueidica.Rueidica {
	// return rueidica.NewRueidica(gb.rueidis)
	return gb.Rueidica
}

func (gb *Gloomberg) String() {
	fmt.Println("gloomberg | " + internal.GloombergVersion)
}

// AddNodes is a migration helper function.
func (gb *Gloomberg) AddNodes(nodes models.Nodes) {
	gb.nodes = nodes
}

func (gb *Gloomberg) startTerminalPrinter() {
	log.Debug("starting terminal printer...")

	for eventLine := range gb.PrintToTerminal {
		if viper.GetBool("log.debug") {
			debugPrefix := fmt.Sprintf("%d | ", len(gb.PrintToTerminal))
			eventLine = fmt.Sprint(debugPrefix, eventLine)
		}

		fmt.Println(eventLine)
	}
}

func readChainwatcherConfiguration(chainwatcherConfig *viper.Viper) models.Nodes {
	var nodes models.Nodes = make([]*models.Node, 0)

	log.Infof("⛓️ config: %+v", chainwatcherConfig.AllSettings())

	// ...or just get a single key
	err := chainwatcherConfig.UnmarshalKey("nodes", &nodes)
	if err != nil {
		log.Errorf("failed to unmarshal configuration: %v", err)

		return nil
	}

	log.Infof("⛓️ nodes: %+v", nodes)

	return nodes
}

func loadRedisConfiguration(redisConfig *viper.Viper) *rueidis.Client {
	// default values set here due to a viper bug that
	// breaks using defaults in combination with Sub()
	redisConfig.SetDefault("enabled", false)
	redisConfig.SetDefault("address", "127.0.0.1:6379")
	redisConfig.SetDefault("database", 0)
	redisConfig.SetDefault("password", "")

	// use hostname as client name
	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("❗️ error getting hostname: %s", err)

		hostname = "unknown"
	}

	// create redis client name
	clientName := hostname + "_gloomberg_v" + internal.GloombergVersion
	redisClientOptions := rueidis.ClientOption{
		InitAddress: []string{redisConfig.GetString("address")},
		ClientName:  clientName,
	}

	rueidisClient, err := rueidis.NewClient(redisClientOptions)
	if err != nil {
		log.Errorf("error getting redis client: %+v", err)

		return nil
	}

	return &rueidisClient
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

func (gb *Gloomberg) PublishOwnSlubSubscription() {
	slugSubscriptions := make([]degendb.SlugSubscription, 0)
	for _, slug := range gb.CollectionDB.OpenseaSlugs() {
		// always subscribe to these events
		eventTypes := []degendb.EventType{degendb.Listing, degendb.CollectionOffer}

		// for collections from config or waller, we also want to subscribe to bids
		if collection := gb.CollectionDB.GetCollectionForSlug(slug); collection != nil {
			if collection.Source != degendb.FromStream {
				eventTypes = append(eventTypes, degendb.Bid)
			}
		}

		slugSubscriptions = append(slugSubscriptions, degendb.SlugSubscription{Slug: slug, Events: eventTypes})
	}

	gb.publishSlugSubscriptions(slugSubscriptions)
}

func (gb *Gloomberg) PublishSlubSubscription(slugSubscription degendb.SlugSubscription) {
	gb.publishSlugSubscriptions(degendb.SlugSubscriptions{slugSubscription})
}

func (gb *Gloomberg) PublishSlubSubscriptions(slugSubscriptions degendb.SlugSubscriptions) {
	gb.publishSlugSubscriptions(slugSubscriptions)
}

func (gb *Gloomberg) publishSlugSubscriptions(slugSubscriptions degendb.SlugSubscriptions) {
	// to enable multiple users to use the central gloomberg instance for events from opensea,
	// we first send the slugs of 'our' collections to the events-subscriptions channel.
	// the central gloomberg instance then creates a subscription on the opensea
	// api and publishes upcoming incoming events to the pubsub channel
	// marshal event to json
	if viper.Get("seawatcher.pubsubClient") == nil {
		log.Warn("❌ not sending slugs to server - pubsub client is not enabled")

		return
	}

	if len(slugSubscriptions) == 0 {
		log.Warn("❌ no slugs to send to gloomberg server")

		return
	}

	log.Debugf("👔 sending %s collection slugs to gloomberg server", style.BoldStyle.Render(strconv.Itoa(len(slugSubscriptions))))

	subscriptionEvent := &seawaModels.SubscriptionEvent{Action: seawaModels.Subscribe, Collections: slugSubscriptions}

	switch {
	case viper.Get("seawatcher.local") != nil:
		// runs on pubsub server side
		gb.In.SeawatcherSubscriptions <- subscriptionEvent

	case viper.Get("seawatcher.pubsubClient") != nil:
		// runs on pubsub client side
		jsonSubscriptionEvent, err := json.Marshal(subscriptionEvent)
		if err != nil {
			log.Error("❌ marshal failed for outgoing list of collection slugs: %s | %v", err, gb.CollectionDB.OpenseaSlugs())

			return
		}

		if gb == nil {
			log.Error("gb is nil")

			return
		}

		// log.Printf("jsonSubscriptionEvent: %+v", jsonSubscriptionEvent)
		// log.Warnf("gb: %+v", gb)
		// log.Printf("*gb: %+v", *gb)
		// log.Printf("gb.Rdb: %+v", gb.Rdb)

		// publish to redis
		if gb.GetRueidica().Rueidis().Do(context.Background(), gb.GetRueidica().Rueidis().B().Publish().Channel(internal.PubSubSeaWatcherMgmt).Message(string(jsonSubscriptionEvent)).Build()).Error() != nil {
			log.Warnf("error publishing event to redis: %s", err)
		} else {
			log.Infof("👔 sent %s collection subscriptions to %s", style.BoldStyle.Render(strconv.Itoa(len(slugSubscriptions))), style.BoldStyle.Render(internal.PubSubSeaWatcherMgmt))
		}
	}
}

func (gb *Gloomberg) IsContract(address common.Address) bool {
	// if its a marketplace address, its a contract
	if marketplace.Addresses().Contains(address) {
		return true
	}

	// check if we have a cached the account type already
	accountType, err := gb.GetRueidica().GetCachedAccountType(context.Background(), address)
	if err == nil {
		return degendb.AccountType(accountType) == degendb.ContractAccount
	}

	log.Debugf("❕ error getting cached account type: %s", err)

	// ok 🙄 seems we really need to check via a node if its a eoa or contract
	codeAt, err := gb.Node().EthClient.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Debugf("❕ failed to get codeAt for %s: %s", address.String(), err)

		return false
	}

	log.Debugf("codeAt(%s): %+v", address.Hex(), codeAt)

	// if there is deployed code at the address, it's a contract
	return len(codeAt) > 0
}

func (gb *Gloomberg) ResolveENS(ensName string) (common.Address, error) {
	if ensName == "" {
		return common.Address{}, errors.New("ensName is empty")
	}

	// reactivate me
	// address, err := gb.Node().ENSLookup(ensName)
	// if err == nil && address != (common.Address{}) {
	// 	err := gb.Rueidica().StoreENSName(ctx, address, ensName)
	// 	if err != nil {
	// 		log.Errorf("error storing ensName %s for address %s: %s", ensName, address, err)
	// 	}

	// 	return address, nil
	// }

	// if err != nil {
	// 	log.Errorf("pp.callMethod error - hex address for ensName %s is %+v", ensName, err)
	// }

	return common.Address{}, errors.New("ens address not found")
}

func (gb *Gloomberg) ReverseResolveAddressToENS(address common.Address) (string, error) {
	if address == (common.Address{}) {
		return "", errors.New("address is empty")
	}

	if address == internal.ZeroAddress {
		return "", errors.New("address is zero address")
	}

	ctx := context.Background()

	// if cachedName, err := cache.GetENSName(ctx, address); err == nil && cachedName != "" {
	if cachedName, err := gb.GetRueidica().GetCachedENSName(ctx, address); err == nil && cachedName != "" {
		log.Debugf("ens ensName for address %s is cached: %s", address.Hex(), cachedName)

		return cachedName, nil
	}

	// reactivate me
	// // name, err := p.callMethod(ctx, ReverseResolveENS, methodCallParams{Address: address})
	// if ensName, err := gb.Node().ReverseLookupAndValidate(address); err == nil && ensName != "" {
	// 	// cache.StoreENSName(ctx, address, ensName)
	// 	err := gb.Rueidica().StoreENSName(ctx, address, ensName)
	// 	if err != nil {
	// 		log.Errorf("error storing ensName %s for address %s: %s", ensName, address.Hex(), err)
	// 	}

	// 	return ensName, nil
	// }

	return "", errors.New("ens ensName not found")
}

func (gb *Gloomberg) FetchFirst1K(contractInfo *degendb.ContractInfo, collectionName string, contractAddress common.Address) {
	go fetchFirst1K(gb, contractInfo, collectionName, contractAddress)
}

// var alreadyFetchedContracts = mapset.NewSet[common.Address]()

// var (
// 	fetchedContracts = mapset.NewSet[common.Address]()
// 	ignoredContracts = mapset.NewSet[common.Address]()
// )

// func (gb *Gloomberg) FetchFirst1K(collectionName string, contractAddress common.Address) error {
// 	if fetchedContracts.Contains(contractAddress) {
// 		log.Debugf("🤷‍♀️ %s already fetched", style.AlmostWhiteStyle.Render(collectionName))

// 		return nil
// 	}

// 	if ignoredContracts.Contains(contractAddress) {
// 		log.Debugf("🤷‍♀️ %s already ignored", style.AlmostWhiteStyle.Render(collectionName))

// 		return nil
// 	}

// 	numTxsToFecth := int64(1337)

// 	// get first 1k txs for contract
// 	transactions, err := external.GetFirstTransactionsByContract(numTxsToFecth, contractAddress)
// 	if err != nil {
// 		return err
// 	}

// 	fetchedContracts.Add(contractAddress)

// 	log.Printf("📝 firstTxs: received %d txs for %s", len(transactions), collectionName)

// 	count := 0
// 	for _, tx := range transactions {
// 		if tx.From == "" {
// 			continue
// 		}

// 		fromAddr := common.HexToAddress(tx.From)
// 		toAddr := common.HexToAddress(tx.To)

// 		if fromAddr != internal.ZeroAddress || tx.TokenID == "" {
// 			pretty.Println(tx)

// 			continue
// 		}

// 		gb.DegenDB.SaveAddressWithFirst1KSlugs(toAddr, []degendb.Tag{degendb.Tag(collectionName)})

// 		time.Sleep(337 * time.Millisecond)

// 		count++

// 		log.Printf("📝 firstTxs: stored %s in db for %s (%d total)", toAddr.Hex(), collectionName, count)
// 	}

// 	return nil
// }

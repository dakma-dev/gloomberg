package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// seaWatcherCmd represents the seawa command.
var seaWatcherCmd = &cobra.Command{
	Use:     "seawatcher",
	Aliases: []string{"seawa", "oswwatcher", "osw"},
	Short:   "receives events from the OpenSea API and pushes them to the redis database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: run,
}

func init() {
	rootCmd.AddCommand(seaWatcherCmd)

	// metrics/prometheus
	viper.SetDefault("metrics.enabled", false)
	viper.SetDefault("metrics.host", net.IPv4(0, 0, 0, 0))
	viper.SetDefault("metrics.port", 9090)
}

func run(_ *cobra.Command, _ []string) {
	log.Info(fmt.Sprintf("⚓️ starting seawatcher %s…", Version))

	//
	// init redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + fmt.Sprint(viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	}).WithContext(context.Background())

	// init stream watcher
	sw := seawa.NewStreamWatcher(viper.GetString("api_keys.opensea"), rdb)

	// prometheus metrics
	if viper.GetBool("metrics.enabled") {
		go func() {
			listenHost := net.ParseIP(viper.GetString("metrics.host"))
			listenPort := viper.GetUint("metrics.port")
			listenAddress := net.JoinHostPort(listenHost.String(), strconv.Itoa(int(listenPort)))

			http.Handle("/metrics", promhttp.Handler())

			log.Info(fmt.Sprintf("⚓️ 📐 metrics: http://%s", listenAddress))

			if err := http.ListenAndServe(listenAddress, nil); err != nil { //nolint:gosec
				log.Error(fmt.Sprintf("⚓️ 📐 ❌ error starting metrics server: %s", err))
			}
		}()
	}

	// loop over incoming events
	for rawItemEvent := range sw.EventChannel() {
		log.Debug(fmt.Sprintf("⚓️ received rawItemEvent: %+v", rawItemEvent))

		itemEventType, ok := rawItemEvent["event_type"].(string)
		if !ok {
			log.Warn(fmt.Sprintf("⚓️ 🤷‍♀️ unknown event type: %s", rawItemEvent["event_type"]))

			continue
		}

		switch osmodels.EventType(itemEventType) {
		case osmodels.ItemSold, osmodels.ItemReceivedOffer:
			log.Debug(fmt.Sprintf("⚓️ received %s: %+v", itemEventType, rawItemEvent))
		case osmodels.ItemListed:
			var itemListedEvent osmodels.ItemListedEvent

			err := mapstructure.Decode(rawItemEvent, &itemListedEvent)
			if err != nil {
				log.Info("⚓️❌ decoding incoming opensea stream api event failed:", err)

				continue
			}

			// nftID is a identification string in the format <chain>/<contract>/<tokenID>
			nftID := strings.Split(itemListedEvent.Payload.Item.NftID, "/")
			if len(nftID) != 3 {
				log.Info(fmt.Sprintf("⚓️ 🤷‍♀️ error parsing nftID: %s", itemListedEvent.Payload.Item.NftID))
			}

			// just publish the event to redis if we have a valid api key (= may have it acquired via opensea api)
			if viper.GetString("api_keys.opensea") == "" {
				log.Debug(fmt.Sprintf("⚓️ 🤷‍♀️ no opensea api key set, skipping event: %s", itemListedEvent.Payload.Item.NftID))

				return
			}

			// marshal event to json
			jsonEvent, err := json.Marshal(itemListedEvent)
			if err != nil {
				log.Info("⚓️❌ json.Marshal failed for incoming stream api event", err)

				continue
			}

			// generate the redis pubsub channel name
			channel := internal.TopicSeaWatcher + "/" + common.HexToAddress(nftID[1]).String() + "/" + string(itemListedEvent.StreamEvent)

			// publish event to redis
			if err := rdb.Publish(context.Background(), channel, jsonEvent).Err(); err != nil {
				log.Warn(fmt.Sprintf("⚓️❕ error publishing event to redis: %s", err.Error()))
			}

			//
			// just for console output, not needed in general
			//

			// parse price
			priceWeiRaw, _, err := big.ParseFloat(itemListedEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
			if err != nil {
				log.Info(fmt.Sprintf("⚓️❌ error parsing price: %s", err.Error()))

				continue
			}

			priceWei, _ := priceWeiRaw.Int(nil)

			var listedBy string
			listedByAddress := common.HexToAddress(itemListedEvent.Payload.Maker.Address)
			listedByStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(listedByAddress.Big().Int64()))

			if itemListedEvent.Payload.Maker.User != "" {
				listedBy = listedByStyle.Render(itemListedEvent.Payload.Maker.User) + " (" + style.ShortenAddressStyled(&listedByAddress, listedByStyle) + ")"
			} else {
				listedBy = style.ShortenAddressStyled(&listedByAddress, listedByStyle)
			}

			eventType := osmodels.TxType[itemListedEvent.StreamEvent]

			log.Info(fmt.Sprintf("%s %s | %sΞ %s | %s", eventType.Icon(), eventType.String(), style.BoldStyle.Render(fmt.Sprintf("%5.3f", price.NewPrice(priceWei).Ether())), style.TerminalLink(itemListedEvent.Payload.Item.Permalink, style.BoldStyle.Render(itemListedEvent.Payload.Item.Metadata.Name)), listedBy))
		}
	}

	// loop
	select {}
}

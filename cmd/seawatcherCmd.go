package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/charmbracelet/log"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// seaWatcherCmd represents the seawa command.
var seaWatcherCmd = &cobra.Command{
	Use:     "seawatcher",
	Aliases: []string{"seawa"},
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
	// init metrics
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

	var redisAddress string
	network := "tcp"

	if viper.IsSet("redis.address") {
		redisAddress = viper.GetString("redis.address")
		if strings.HasPrefix(redisAddress, "unix://") {
			network = "unix"
			redisAddress = strings.Replace(redisAddress, "unix://", "", 1)
		}
	} else {
		// fallback to old config
		redisAddress = viper.GetString("redis.host") + ":" + fmt.Sprint(viper.GetInt("redis.port"))
	}
	//
	// init redis client
	rdb := redis.NewClient(&redis.Options{
		Network:  network,
		Addr:     redisAddress,
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	}).WithContext(context.Background())

	//
	// start sea watcher & loop forever
	seawa.NewSeaWatcher(viper.GetString("api_keys.opensea"), rdb).Start()
}

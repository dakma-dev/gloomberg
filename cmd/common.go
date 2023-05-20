package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
)

func GetRedisClient() rueidis.Client {
	// use hostname as client name
	hostname, err := os.Hostname()
	if err != nil {
		log.Error(fmt.Sprintf("❗️ error getting hostname: %s", err))

		hostname = "unknown"
	}

	// rueidis / new redis library
	var connectAddr string

	if viper.IsSet("redis.address") {
		connectAddr = viper.GetString("redis.address")
	} else {
		// fallback to old config
		connectAddr = fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	}

	rdb, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{connectAddr}, ClientName: hostname})
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

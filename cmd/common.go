package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

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
	var connectAddr net.Addr

	if viper.IsSet("redis.address") {
		splittedAddress := strings.Split(viper.GetString("redis.address"), ":")

		host := net.ParseIP(splittedAddress[0])
		port, _ := strconv.ParseInt(splittedAddress[1], 10, 64)

		connectAddr = &net.TCPAddr{IP: host, Port: int(port)}
	} else {
		// fallback to old config
		connectAddr = &net.TCPAddr{IP: net.ParseIP(viper.GetString("redis.host")), Port: viper.GetInt("redis.port")}
	}

	rdb, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{connectAddr.String()}, ClientName: hostname})
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

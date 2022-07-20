package cache

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	rdb *redis.Client
	mu  = &sync.RWMutex{}
)

type NamespaceKey string

const (
	PrefixContractName NamespaceKey = "contractName"
	PrefixENS          NamespaceKey = "ensDomain"
	KeyDelimiter       string       = ":"
)

func KeyContract(contractAddress common.Address) string {
	return fmt.Sprint(PrefixContractName, KeyDelimiter, contractAddress.Hex())
}

func KeyENS(address common.Address) string {
	return fmt.Sprint(PrefixENS, KeyDelimiter, address.Hex())
}

func GetRedisClient() *redis.Client {
	mu.Lock()
	if rdb == nil {
		redisHost := viper.GetString("redis.host")
		redisPort := viper.GetInt("redis.port")
		redisPassword := viper.GetString("redis.password")
		redisDatabase := viper.GetInt("redis.database")

		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprint(redisHost, ":", redisPort),
			Password: redisPassword,
			DB:       redisDatabase,
		})

		if viper.GetBool("log.debug") {
			fmt.Println(rdb.Info(rdb.Context()).Val())
		}
	}
	mu.Unlock()

	return rdb
}

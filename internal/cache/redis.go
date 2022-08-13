package cache

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

//	cache = &redisCache{
//		rdb: nil,
//	}
var mu = &sync.RWMutex{}

type NamespaceKey string

const (
	PrefixContractName NamespaceKey = "contractName"
	PrefixENS          NamespaceKey = "ensDomain"
	KeyDelimiter       string       = ":"
)

func keyContract(contractAddress common.Address) string {
	return fmt.Sprint(PrefixContractName, KeyDelimiter, contractAddress.Hex())
}

func keyENS(address common.Address) string {
	return fmt.Sprint(PrefixENS, KeyDelimiter, address.Hex())
}

func NewRCache(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: strings.Join([]string{
			viper.GetString("redis.host"),
			fmt.Sprint(viper.GetInt("redis.port")),
		}, ":"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	})

	rdb.WithContext(ctx)

	return rdb
}

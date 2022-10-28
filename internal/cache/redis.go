package cache

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

const (
	prefixContractName string = "contractName"
	prefixENS          string = "ensDomain"
	prefixOSSlug       string = "osslug"
	prefixBlurSlug     string = "blurslug"
	keyDelimiter       string = ":"
)

func keyContract(contractAddress common.Address) string {
	return fmt.Sprint(prefixContractName, keyDelimiter, contractAddress.Hex())
}

func keyENS(address common.Address) string {
	return fmt.Sprint(prefixENS, keyDelimiter, address.Hex())
}

func keyOSSlug(address common.Address) string {
	return fmt.Sprint(prefixOSSlug, keyDelimiter, address.Hex())
}

func keyBlurSlug(address common.Address) string {
	return fmt.Sprint(prefixBlurSlug, keyDelimiter, address.Hex())
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: strings.Join([]string{
			viper.GetString("redis.host"),
			fmt.Sprint(viper.GetInt("redis.port")),
		}, ":"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	})

	rdb.WithContext(context.Background())

	return rdb
}

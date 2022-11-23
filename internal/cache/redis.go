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
	keywordContractName string = "contractName"
	keywordENS          string = "ensDomain"
	keywordOSSlug       string = "osslug"
	keywordBlurSlug     string = "blurslug"
	keywordFloorPrice   string = "floor"
	keywordSalira       string = "salira"
	keyDelimiter        string = ":"
)

func keyContract(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordContractName)
}

func keyENS(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordENS)
}

func keyOSSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordOSSlug)
}

func keyBlurSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordBlurSlug)
}

func keyFloorPrice(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordFloorPrice)
}

func keySalira(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordSalira)
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

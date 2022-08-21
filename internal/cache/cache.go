package cache

import (
	"context"
	"errors"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// enable other cache/datastore backends besides redis?
// type ExtCache interface {
// 	cacheName(address common.Address, keyFunc func(common.Address) string, value string, duration time.Duration)
// 	getName(address common.Address, keyFunc func(common.Address) string) (string, error)
// }

var gbCache *GbCache

type GbCache struct {
	rdb           *redis.Client
	addressToName map[common.Address]string
}

func New(ctx context.Context) *GbCache {
	if gbCache != nil {
		return gbCache
	}

	gCache := &GbCache{
		addressToName: make(map[common.Address]string),
	}

	if viper.GetBool("redis.enabled") {
		if client := NewRedisClient(ctx); client != nil {
			gCache.rdb = client
		} else {
			viper.Set("redis.enabled", false)
		}
	}

	return gCache
}

func (c *GbCache) GetRDB() *redis.Client {
	return c.rdb
}

func (c *GbCache) CacheCollectionName(collectionAddress common.Address, collectionName string) {
	c.cacheName(collectionAddress, keyContract, collectionName, viper.GetDuration("cache.names_ttl"))
}

func (c *GbCache) GetCollectionName(collectionAddress common.Address) (string, error) {
	return c.getName(collectionAddress, keyContract)
}

func (c *GbCache) CacheENSName(walletAddress common.Address, ensName string) {
	c.cacheName(walletAddress, keyENS, ensName, viper.GetDuration("cache.ens_ttl"))
}

func (c *GbCache) GetENSName(walletAddress common.Address) (string, error) {
	return c.getName(walletAddress, keyENS)
}

func (c *GbCache) StoreEvent(contractAddress common.Address, collectionName string, tokenID uint64, priceWei uint64, numItems uint, eventTime time.Time, eventType int64) {
	xAddArgs := &redis.XAddArgs{
		Stream: "sales",
		MaxLen: 100000,
		Approx: true,
		ID:     "*",
		Values: map[string]any{
			"contractAddress": contractAddress.Hex(),
			"collectionName":  collectionName,
			"tokenID":         tokenID,
			"priceWei":        priceWei,
			"numItems":        numItems,
			"eventTime":       eventTime,
			"eventType":       eventType,
		},
	}

	if c.rdb != nil {
		gbl.Log.Debugf("redis | adding sale: %s #%d", collectionName, int(tokenID))

		if added, err := c.rdb.XAdd(c.rdb.Context(), xAddArgs).Result(); err == redis.Nil {
			gbl.Log.Errorf("redis | strange redis.Nil while adding to stream: %s %d -xxx-> %s: %s", collectionName, tokenID, xAddArgs.Stream, err)
		} else if err != nil {
			gbl.Log.Errorf("redis | could not add event: %s", err)
		} else {
			gbl.Log.Debugf("redis | added event (%d) to stream: %s %d | %s", eventType, collectionName, tokenID, added)
		}
	}
}

func (c *GbCache) cacheName(address common.Address, keyFunc func(common.Address) string, value string, duration time.Duration) {
	c.addressToName[address] = value

	if c.rdb != nil {
		gbl.Log.Debugf("redis | searching for: %s", keyFunc(address))

		err := c.rdb.SetEX(c.rdb.Context(), keyFunc(address), value, duration).Err()

		if err != nil {
			gbl.Log.Warnf("redis | error while adding: %s", err.Error())
		} else {
			gbl.Log.Debugf("redis | added: %s -> %s", keyFunc(address), value)
		}
	}
}

func (c *GbCache) getName(address common.Address, keyFunc func(common.Address) string) (string, error) {
	if name := c.addressToName[address]; name != "" {
		return name, nil
	}

	if c.rdb != nil {
		gbl.Log.Debugf("redis | searching for: %s", keyFunc(address))

		if name, err := c.rdb.Get(c.rdb.Context(), keyFunc(address)).Result(); err == nil && name != "" {
			gbl.Log.Debugf("redis | using cached name: %s", name)

			c.addressToName[address] = name

			return name, nil
		} else {
			gbl.Log.Debugf("redis | name not found in cache: %s", err)

			return "", err
		}
	}

	return "", errors.New("name not found in cache")
}

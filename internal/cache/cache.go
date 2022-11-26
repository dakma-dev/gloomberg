package cache

import (
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// enable other cache/datastore backends besides redis?
// type ExtCache interface {
// 	cacheName(address common.Address, keyFunc func(common.Address) string, value string, duration time.Duration)
// 	getName(address common.Address, keyFunc func(common.Address) string) (string, error)
// }

var gbCache *GbCache

const noENSName = "NO-ENS-NAME"

type GbCache struct {
	mu  *sync.RWMutex
	rdb *redis.Client
	// addressToName map[common.Address]string
	localCache      map[string]string
	localFloatCache map[string]float64
}

func New() *GbCache {
	if gbCache != nil {
		return gbCache
	}

	gbCache = &GbCache{
		mu: &sync.RWMutex{},
		// addressToName: make(map[common.Address]string),
		localCache:      make(map[string]string),
		localFloatCache: make(map[string]float64),
	}

	if viper.GetBool("redis.enabled") {
		if client := NewRedisClient(); client != nil {
			gbCache.rdb = client
		} else {
			viper.Set("redis.enabled", false)
		}
	}

	return gbCache
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

func (c *GbCache) StoreEvent(contractAddress common.Address, collectionName string, tokenID *big.Int, priceWei uint64, numItems uint64, eventTime time.Time, eventType int64) {
	xAddArgs := &redis.XAddArgs{
		Stream: "sales",
		MaxLen: 100000,
		Approx: true,
		ID:     "*",
		Values: map[string]any{
			"contractAddress": contractAddress.Hex(),
			"collectionName":  collectionName,
			"tokenID":         tokenID.Uint64(),
			"priceWei":        priceWei,
			"numItems":        numItems,
			"eventTime":       eventTime,
			"eventType":       eventType,
		},
	}

	if c.rdb != nil {
		gbl.Log.Debugf("redis | adding sale: %s #%d", collectionName, tokenID)

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
	if value == "" {
		value = noENSName
	}

	c.mu.Lock()
	// c.addressToName[address] = value
	c.localCache[keyFunc(address)] = value
	c.mu.Unlock()

	if c.rdb != nil {
		gbl.Log.Debugf("redis | caching %s -> %s", keyFunc(address), value)

		err := c.rdb.SetEX(c.rdb.Context(), keyFunc(address), value, duration).Err()

		if err != nil {
			gbl.Log.Warnf("redis | error while adding: %s", err.Error())
		} else {
			gbl.Log.Debugf("redis | added: %s -> %s", keyFunc(address), value)
		}
	}
}

func (c *GbCache) getName(address common.Address, keyFunc func(common.Address) string) (string, error) {
	c.mu.RLock()
	// name := c.addressToName[address]
	name := c.localCache[keyFunc(address)]
	c.mu.RUnlock()

	if name != "" {
		if name == noENSName {
			name = ""
		}

		gbl.Log.Debugf("cache | found name in in-memory cache: '%s'", name)

		return name, nil
	}

	if c.rdb != nil {
		gbl.Log.Debugf("redis | searching for: %s", keyFunc(address))

		if name, err := c.rdb.Get(c.rdb.Context(), keyFunc(address)).Result(); err == nil {
			gbl.Log.Debugf("redis | using cached name: %s", name)

			c.mu.Lock()
			// c.addressToName[address] = name
			c.localCache[keyFunc(address)] = name
			c.mu.Unlock()

			if name == noENSName {
				name = ""
			}

			return name, nil
		} else if errors.Is(err, redis.Nil) {
			gbl.Log.Debugf("redis | redis.Nil - name not found in cache: %s", keyFunc(address))
		} else {
			gbl.Log.Debugf("redis | get error: %s", err)

			return "", err
		}
	}

	return "", errors.New("name not found in cache")
}

func (c *GbCache) cacheFloat(address common.Address, keyFunc func(common.Address) string, value float64, duration time.Duration) {
	c.mu.Lock()
	// c.addressToName[address] = value
	c.localFloatCache[keyFunc(address)] = value
	c.mu.Unlock()

	if c.rdb != nil {
		gbl.Log.Debugf("redis | caching %s -> %s", keyFunc(address), value)

		err := c.rdb.SetEX(c.rdb.Context(), keyFunc(address), value, duration).Err()

		if err != nil {
			gbl.Log.Warnf("redis | error while adding: %s", err.Error())
		} else {
			gbl.Log.Debugf("redis | added: %s -> %s", keyFunc(address), value)
		}
	}
}

func (c *GbCache) getFloat(address common.Address, keyFunc func(common.Address) string) (float64, error) {
	c.mu.RLock()
	// value := c.addressToName[address]
	value := c.localFloatCache[keyFunc(address)]
	c.mu.RUnlock()

	if value != 0 {
		gbl.Log.Debugf("cache | found name in in-memory cache: '%s'", value)

		return value, nil
	}

	if c.rdb != nil {
		gbl.Log.Debugf("redis | searching for: %s", keyFunc(address))

		if value, err := c.rdb.Get(c.rdb.Context(), keyFunc(address)).Float64(); err == nil {
			gbl.Log.Debugf("redis | using cached value: %s", value)

			c.mu.Lock()
			// c.addressToName[address] = name
			c.localFloatCache[keyFunc(address)] = value
			c.mu.Unlock()

			return value, nil
		} else if errors.Is(err, redis.Nil) {
			gbl.Log.Debugf("redis | redis.Nil - value not found in cache: %s", keyFunc(address))
		} else {
			gbl.Log.Debugf("redis | get error: %s", err)

			return 0, err
		}
	}

	return 0, errors.New("value not found in cache")
}

//
// names

func StoreENSName(walletAddress common.Address, ensName string) {
	c := New()
	c.cacheName(walletAddress, keyENS, ensName, viper.GetDuration("cache.ens_ttl"))
}

func GetENSName(walletAddress common.Address) (string, error) {
	c := New()
	return c.getName(walletAddress, keyENS)
}

func StoreContractName(contractAddress common.Address, contractName string) {
	c := New()
	c.cacheName(contractAddress, keyContract, contractName, viper.GetDuration("cache.names_ttl"))
}

func GetContractName(contractAddress common.Address) (string, error) {
	c := New()
	return c.getName(contractAddress, keyContract)
}

//
// slugs

func StoreOSSlug(contractAddress common.Address, slug string) {
	c := New()
	c.cacheName(contractAddress, keyOSSlug, slug, viper.GetDuration("cache.slug_ttl"))
}

func StoreBlurSlug(contractAddress common.Address, slug string) {
	c := New()
	c.cacheName(contractAddress, keyBlurSlug, slug, viper.GetDuration("cache.slug_ttl"))
}

//
// numbers

func StoreFloor(address common.Address, value float64) {
	c := New()
	c.cacheFloat(address, keyFloorPrice, value, viper.GetDuration("cache.floor_ttl"))
}

func GetFloor(address common.Address) (float64, error) {
	c := New()
	return c.getFloat(address, keyFloorPrice)
}

func StoreSalira(address common.Address, value float64) {
	c := New()
	c.cacheFloat(address, keySalira, value, viper.GetDuration("cache.salira_ttl"))
}

func GetSalira(address common.Address) (float64, error) {
	c := New()
	return c.getFloat(address, keySalira)
}

// NotificationLock implements a lock to prevent sending multiple notifications for the same event
// see https://redis.io/docs/manual/patterns/distributed-locks/#correct-implementation-with-a-single-instance
func NotificationLock(txID common.Hash) (bool, error) {
	c := New()

	releaseKey := uuid.New()

	c.mu.Lock()
	c.localCache[keyNotificationsLock(txID)] = releaseKey.String()
	c.mu.Unlock()

	unlocked := false

	var err error

	if c.rdb != nil {
		unlocked, err = c.rdb.SetNX(c.rdb.Context(), keyNotificationsLock(txID), releaseKey.String(), viper.GetDuration("cache.notifications_lock_ttl")).Result()

		gbl.Log.Debugf("📣 %s | locked %+v | err %+v", txID.String(), unlocked)

		if err != nil {
			gbl.Log.Warnf("❌ redis | error while adding: %s", err.Error())
		} else {
			gbl.Log.Debugf("📣 redis | added: %s -> %s", keyNotificationsLock(txID), releaseKey)
		}
	}

	return unlocked, nil
}

func ReleaseNotificationLock(contractAddress common.Address) (string, error) {
	c := New()
	return c.getName(contractAddress, keyContract)
}

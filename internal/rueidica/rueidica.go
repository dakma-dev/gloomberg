package rueidica

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
)

const (
	keywordContractName string = "contractName"
	keywordENS          string = "ensDomain"
	keywordFloorOS      string = "floorOS"
	keywordOSSlug       string = "osslug"
	keywordBlurSlug     string = "blurslug"
	keywordSalira       string = "salira"
	keyDelimiter        string = ":"
)

type Rueidica struct {
	rueidis.Client
}

func NewRueidica(rdb rueidis.Client) *Rueidica {
	return &Rueidica{rdb}
}

//
// wrappers

// Wallet/Contract names.
func (r *Rueidica) GetCachedContractName(ctx context.Context, address common.Address) (string, error) {
	log.Debugf("rueidica.GetCachedContractName | %+v", address)

	return r.getCachedName(ctx, address, keyContract)
}

func (r *Rueidica) StoreContractName(ctx context.Context, address common.Address, name string) error {
	log.Debugf("rueidica.StoreContractName | %+v ⇄ %+v", address.Hex(), name)

	return r.cacheName(ctx, address, name, keyContract, viper.GetDuration("cache.names_ttl"))
}

// ENS.
func (r *Rueidica) GetCachedENSName(ctx context.Context, address common.Address) (string, error) {
	log.Debugf("rueidica.GetCachedENSName | %+v", address)

	return r.getCachedName(ctx, address, keyENS)
}

func (r *Rueidica) StoreENSName(ctx context.Context, address common.Address, name string) error {
	log.Debugf("rueidica.StoreENSName | %+v -> %+v", address.Hex(), name)

	return r.cacheName(ctx, address, name, keyENS, viper.GetDuration("cache.ens_ttl"))
}

// Floors.
func (r *Rueidica) GetCachedOSFloor(ctx context.Context, address common.Address) (float64, error) {
	log.Debugf("rueidica.GetCachedOSFloor | %+v", address)

	return r.getCachedNumber(ctx, address, keyFloorOS)
}

func (r *Rueidica) StoreOSFloor(ctx context.Context, address common.Address, value float64) error {
	log.Debugf("rueidica.StoreOSFloor | %+v -> %+v", address.Hex(), value)

	return r.cacheName(ctx, address, fmt.Sprint(value), keyFloorOS, viper.GetDuration("cache.floor_ttl"))
}

// Salira.
func (r *Rueidica) GetCachedSalira(ctx context.Context, address common.Address) (float64, error) {
	log.Debugf("rueidica.GetCachedSalira | %+v", address)

	return r.getCachedNumber(ctx, address, keySalira)
}

func (r *Rueidica) StoreSalira(ctx context.Context, address common.Address, value float64) error {
	log.Debugf("rueidica.StoreSalira | %+v -> %+v", address.Hex(), value)

	return r.cacheName(ctx, address, fmt.Sprint(value), keySalira, viper.GetDuration("cache.salira_ttl"))
}

// Slugs.
func (r *Rueidica) StoreOSSlug(ctx context.Context, address common.Address, slug string) error {
	log.Debugf("rueidica.StoreOSSlug | %+v -> %+v", address.Hex(), slug)

	return r.cacheName(ctx, address, slug, keyOSSlug, viper.GetDuration("cache.slug_ttl"))
}

func (r *Rueidica) StoreBlurSlug(ctx context.Context, address common.Address, slug string) error {
	log.Debugf("rueidica.StoreBlurSlug | %+v -Y %+v", address.Hex(), slug)

	return r.cacheName(ctx, address, slug, keyBlurSlug, viper.GetDuration("cache.slug_ttl"))
}

//
// implementations

func (r *Rueidica) getCachedName(ctx context.Context, address common.Address, keyFunc func(common.Address) string) (string, error) {
	clientCacheTTL := viper.GetDuration("cache.names_client_ttl")

	if r != nil {
		cachedName, err := r.DoCache(ctx, r.B().Get().Key(keyFunc(address)).Cache(), clientCacheTTL).ToString()

		switch {
		case err != nil && rueidis.IsRedisNil(err):
			gbl.Log.Debugf("rueidis | no name in cache for %s", address.Hex())
		case err != nil:
			gbl.Log.Errorf("rueidis | error getting cached name: %s", err)
		default:
			gbl.Log.Debugf("rueidis | found name: %s -> %s", keyFunc(address), cachedName)
		}

		return cachedName, err
	}

	return "", errors.New("value not found in cache")
}

func (r *Rueidica) getCachedNumber(ctx context.Context, address common.Address, keyFunc func(common.Address) string) (float64, error) {
	clientCacheTTL := viper.GetDuration("cache.names_client_ttl")

	if r != nil {
		cachedNumber, err := r.DoCache(ctx, r.B().Get().Key(keyFunc(address)).Cache(), clientCacheTTL).ToString()

		switch {
		case err != nil && rueidis.IsRedisNil(err):
			gbl.Log.Debugf("rueidis | no number in cache for %s", address.Hex())

			return 0, err
		case err != nil:
			gbl.Log.Errorf("rueidis | error getting cached number: %s", err)
		default:
			gbl.Log.Debugf("rueidis | found number: %s -> %f", keyFunc(address), cachedNumber)
		}

		num, err := strconv.ParseFloat(cachedNumber, 64)
		if err != nil {
			gbl.Log.Errorf("rueidis | error parsing cached number: %s", err)

			return 0, err
		}

		log.Debugf("cachedNumber: %+v | num: %+v", cachedNumber, num)

		return num, err
	}

	return 0, errors.New("value not found in cache")
}

func (r *Rueidica) cacheName(ctx context.Context, address common.Address, name string, keyFunc func(common.Address) string, duration time.Duration) error {
	err := r.Do(ctx, r.B().Set().Key(keyFunc(address)).Value(name).ExSeconds(int64(duration.Seconds())).Build()).Error()
	if err != nil {
		gbl.Log.Errorf("rueidis | error caching name: %s ⇄ %s | %s", address.Hex(), name, err)

		return err
	}

	return nil
}

//
// notifications lock

// NotificationLock implements a lock to prevent sending multiple notifications for the same event
// Refactored to use the Redlock algorithm as recommended in the [redis SET doc].
//
// [redis SET doc]: https://redis.io/commands/set/#patterns
func (r *Rueidica) NotificationLock(ctx context.Context, txID common.Hash) (bool, error) {
	return r.NotificationLockWtihDuration(ctx, txID, viper.GetDuration("cache.notifications_lock_ttl"))
}

// NotificationLockWtihDuration implements a lock to prevent sending multiple notifications for the same event
// Refactored to use the Redlock algorithm as recommended in the [redis SET doc].
//
// [redis SET doc]: https://redis.io/commands/set/#patterns
func (r *Rueidica) NotificationLockWtihDuration(ctx context.Context, txID common.Hash, duration time.Duration) (bool, error) {
	var connectAddr net.Addr

	if viper.IsSet("redis.address") {
		splittedAddress := strings.Split(viper.GetString("redis.address"), ":")

		host := net.ParseIP(splittedAddress[0])
		port, _ := strconv.ParseUint(splittedAddress[1], 10, 16)

		connectAddr = &net.TCPAddr{IP: host, Port: int(port)}
	} else {
		// fallback to old config
		connectAddr = &net.TCPAddr{IP: net.ParseIP(viper.GetString("redis.host")), Port: int(viper.GetUint16("redis.port"))}
	}

	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	client := goredislib.NewClient(&goredislib.Options{Addr: connectAddr.String()})
	pool := goredis.NewPool(client)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutex := rs.NewMutex("txID.Hex()", redsync.WithExpiry(duration))

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.LockContext(ctx); err != nil {
		gbl.Log.Errorf("rueidis | %s | %s", txID.Hex(), style.Bold("acquire lock failed"))

		return false, err
	}

	return true, nil
}

//
// keys

func keyContract(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordContractName)
}

func keyENS(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordENS)
}

func keyFloorOS(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordFloorOS)
}

func keyOSSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordOSSlug)
}

func keyBlurSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordBlurSlug)
}

func keySalira(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordSalira)
}

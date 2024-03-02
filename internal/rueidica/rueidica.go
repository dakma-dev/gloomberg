package rueidica

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidislock"
	"github.com/spf13/viper"
)

const (
	keywordContractName string = "contractName"
	keywordAccountType  string = "accountType"
	keywordENS          string = "ensDomain"
	keywordFloorOS      string = "floorOS"
	keywordOSSlug       string = "osslug"
	keywordAddress      string = "address"
	keywordBlurSlug     string = "blurslug"
	keywordSalira       string = "salira"
	keywordMethodSig    string = "methodsig"
	keyDelimiter        string = ":"
)

type Rueidica struct {
	rueidis.Client
}

func NewRueidica(rueidisClient *rueidis.Client) *Rueidica {
	return &Rueidica{*rueidisClient}
}

// func NewRueidica(redisConfig *viper.Viper) *Rueidica {
// 	// default values set here due to a viper bug that
// 	// breaks using defaults in combination with Sub()
// 	redisConfig.SetDefault("enabled", false)
// 	redisConfig.SetDefault("host", "127.0.0.1")
// 	redisConfig.SetDefault("port", 6379)
// 	redisConfig.SetDefault("database", 0)
// 	redisConfig.SetDefault("password", "")

// 	var connectAddr string

// 	if redisConfig.IsSet("address") {
// 		connectAddr = redisConfig.GetString("address")
// 	} else {
// 		// fallback to old config
// 		connectAddr = fmt.Sprintf("%s:%d", redisConfig.GetString("host"), redisConfig.GetInt("port"))
// 	}

// 	// use hostname as client name
// 	hostname, err := os.Hostname()
// 	if err != nil {
// 		log.Error(fmt.Sprintf("❗️ error getting hostname: %s", err))

// 		hostname = "unknown"
// 	}

// 	clientName := hostname + "_gloomberg_v" + internal.GloombergVersion
// 	redisClientOptions := rueidis.ClientOption{
// 		InitAddress: []string{connectAddr},
// 		ClientName:  clientName,
// 	}

// 	return &Rueidica{*getRueidis(redisClientOptions)}
// }

func (r *Rueidica) Rueidis() rueidis.Client {
	return r.Client
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

// Account type.
func (r *Rueidica) GetCachedAccountType(ctx context.Context, address common.Address) (string, error) {
	log.Debugf("rueidica.GetCachedAccountType | %+v", address)

	return r.getCachedName(ctx, address, keyAccountType)
}

func (r *Rueidica) StoreAccountType(ctx context.Context, address common.Address, accountType degendb.AccountType) error {
	log.Debugf("rueidica.StoreAccountType | %+v ⇄ %+v", address.Hex(), accountType)

	return r.cacheName(ctx, address, string(accountType), keyAccountType, viper.GetDuration("cache.names_ttl"))
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

// method sig.
func (r *Rueidica) GetCachedMethodSig(ctx context.Context, methodSig [4]byte) (string, error) {
	log.Debugf("rueidica.GetCachedMethodSig | %x", methodSig)

	return r.getCachedMethodSig(ctx, methodSig, keyMethodSig)
}

func (r *Rueidica) StoreMethodSig(ctx context.Context, methodSig [4]byte, name string) error {
	log.Debugf("rueidica.StoreENSName | %x -> %+v", methodSig, name)

	return r.cacheMethodSig(ctx, methodSig, name, keyMethodSig)
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
func (r *Rueidica) StoreOSSlugForAddress(ctx context.Context, address common.Address, slug string) error {
	log.Debugf("rueidica.StoreOSSlugForAddress | %+v -> %+v", address.Hex(), slug)

	return r.cacheName(ctx, address, slug, keyAddresToOSSlug, viper.GetDuration("cache.slug_ttl"))
}

func (r *Rueidica) GetOSSlugForAddress(ctx context.Context, address common.Address) (string, error) {
	log.Debugf("rueidica.GetCachedENSName | %+v", address)

	return r.getCachedName(ctx, address, keyENS)
}

func (r *Rueidica) StoreAddressForOSSlug(ctx context.Context, slug string, address common.Address) error {
	log.Debugf("rueidica.StoreAddressForOSSlug | %+v -> %+v", slug, address.Hex())

	return r.cacheAddressWithKey(ctx, keyOSSlugsToAddress(slug), address, viper.GetDuration("cache.slug_ttl"))
	// return r.cacheName(ctx, address, slug, keyOSSlugsToAddress, viper.GetDuration("cache.slug_ttl"))
}

func (r *Rueidica) GetAddressForOSSlug(ctx context.Context, slug string) (string, error) {
	log.Debugf("rueidica.GetAddressForOSSlug | %+v", slug)

	return r.getCachedStringValueWithKey(ctx, keyOSSlugsToAddress(slug))
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
			log.Debugf("rueidis | no name in cache for %s", address.Hex())
		case err != nil:
			log.Errorf("rueidis | error getting cached name: %s", err)
		default:
			log.Debugf("rueidis | found name: %s -> %s", keyFunc(address), cachedName)
		}

		return cachedName, err
	}

	return "", errors.New("value not found in cache")
}

func (r *Rueidica) getCachedMethodSig(ctx context.Context, methodSig [4]byte, keyFunc func([4]byte) string) (string, error) {
	clientCacheTTL := viper.GetDuration("cache.names_client_ttl")

	if r != nil {
		cachedName, err := r.DoCache(ctx, r.B().Get().Key(keyFunc(methodSig)).Cache(), clientCacheTTL).ToString()

		switch {
		case err != nil && rueidis.IsRedisNil(err):
			log.Debugf("rueidis | no name in cache for %x", methodSig)
		case err != nil:
			log.Errorf("rueidis | error getting cached name: %s", err)

			return "", err
		default:
			log.Debugf("rueidis | found name: %s -> %s", keyFunc(methodSig), cachedName)
		}

		return cachedName, err
	}

	return "", errors.New("value not found in cache")
}

func (r *Rueidica) getCachedStringValueWithKey(ctx context.Context, rKey string) (string, error) {
	clientCacheTTL := viper.GetDuration("cache.names_client_ttl")

	if r != nil {
		cachedValue, err := r.DoCache(ctx, r.B().Get().Key(rKey).Cache(), clientCacheTTL).ToString()

		switch {
		case err != nil && rueidis.IsRedisNil(err):
			log.Debugf("rueidis | no cachedValue in cache for %s", rKey)
		case err != nil:
			log.Errorf("rueidis | error getting cached name: %s", err)
		default:
			log.Debugf("rueidis | found name: %s -> %s", rKey, cachedValue)
		}

		return cachedValue, err
	}

	return "", errors.New("value not found in cache")
}

func (r *Rueidica) getCachedNumber(ctx context.Context, address common.Address, keyFunc func(common.Address) string) (float64, error) {
	clientCacheTTL := viper.GetDuration("cache.names_client_ttl")

	if r != nil {
		cachedNumber, err := r.DoCache(ctx, r.B().Get().Key(keyFunc(address)).Cache(), clientCacheTTL).ToString()

		switch {
		case err != nil && rueidis.IsRedisNil(err):
			log.Debugf("rueidis | no number in cache for %s", address.Hex())

			return 0, err
		case err != nil:
			log.Errorf("rueidis | error getting cached number: %s", err)
		default:
			log.Debugf("rueidis | found number: %s -> %s", keyFunc(address), cachedNumber)
		}

		num, err := strconv.ParseFloat(cachedNumber, 64)
		if err != nil {
			log.Errorf("rueidis | error parsing cached number: %s", err)

			return 0, err
		}

		log.Debugf("cachedNumber: %+v | num: %+v", cachedNumber, num)

		return num, err
	}

	return 0, errors.New("value not found in cache")
}

func (r *Rueidica) cacheName(ctx context.Context, address common.Address, name string, keyFunc func(common.Address) string, duration time.Duration) error {
	log.Debugf("🏪 rueidis | trying to cache: %s ⇄ %s | %s | %+v | %+v", address.Hex(), name, duration, keyFunc(address), r)

	if r != nil {
		err := r.Do(ctx, r.B().Set().Key(keyFunc(address)).Value(name).ExSeconds(int64(duration.Seconds())).Build()).Error()
		if err != nil {
			log.Errorf("rueidis | error caching name: %s ⇄ %s | %s", address.Hex(), name, err)
			log.Errorf("❌ rueidis | error caching name: %s ⇄ %s | %s", address.Hex(), name, err)

			return err
		}
	}

	return nil
}

func (r *Rueidica) cacheMethodSig(ctx context.Context, methodSig [4]byte, name string, keyFunc func([4]byte) string) error {
	err := r.Do(ctx, r.B().Set().Key(keyFunc(methodSig)).Value(name).Build()).Error()
	if err != nil {
		log.Errorf("rueidis | error caching method sig: %x ⇄ %s | %s", methodSig, name, err)

		return err
	}

	return nil
}

func (r *Rueidica) cacheAddressWithKey(ctx context.Context, rKey string, rValue common.Address, duration time.Duration) error {
	err := r.Do(ctx, r.B().Set().Key(rKey).Value(rValue.Hex()).ExSeconds(int64(duration.Seconds())).Build()).Error()
	if err != nil {
		log.Errorf("rueidis | error caching: %s ⇄ %s | %s", rKey, rValue, err)

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
func (r *Rueidica) NotificationLock(txHash common.Hash) (context.CancelFunc, error) {
	return r.NotificationLockWtihDuration(txHash.Hex(), viper.GetDuration("cache.notifications_lock_ttl"))
}

// NotificationLockWtihDuration implements a lock to prevent sending multiple notifications for the same event.
func (r *Rueidica) NotificationLockWtihDuration(identifier string, duration time.Duration) (context.CancelFunc, error) {
	var connectAddr string

	if viper.IsSet("redis.address") {
		connectAddr = viper.GetString("redis.address")
	} else {
		// fallback to old config
		connectAddr = fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	}

	// use hostname as client name
	hostname, err := os.Hostname()
	if err != nil {
		log.Error(fmt.Sprintf("❗️ error getting hostname: %s", err))

		hostname = "unknown"
	}

	clientName := hostname + "_gloomberg_v" + internal.GloombergVersion
	redisClientOptions := rueidis.ClientOption{InitAddress: []string{connectAddr}, ClientName: clientName}

	locker, err := rueidislock.NewLocker(rueidislock.LockerOption{
		ClientOption: redisClientOptions,
		KeyMajority:  2, // please make sure that all your `Locker`s share the same KeyMajority
	})
	if err != nil {
		panic(err)
	}
	defer locker.Close()

	// acquire the lock
	_, cancel, err := locker.WithContext(context.Background(), identifier)
	if err != nil {
		if errors.Is(err, rueidislock.ErrLockerClosed) {
			panic(err)
		}
	}
	// "my_lock" is acquired. use the ctx as normal.
	log.Debugf("🔒 %s | notification lock acquired (%.0fsec)", style.ShortenAddress(common.HexToAddress(identifier)), duration.Seconds())
	// doSomething(ctx)

	// invoke cancel() to release the lock.
	// cancel()
	// log.Printf("🔐 lock released")

	return cancel, nil
}

//
// keys

func keyAccountType(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordAccountType)
}

func keyContract(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordContractName)
}

func keyENS(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordENS)
}

func keyFloorOS(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordFloorOS)
}

func keyAddresToOSSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordOSSlug)
}

func keyOSSlugsToAddress(slug string) string {
	return fmt.Sprint(slug, keyDelimiter, keywordAddress)
}

func keyBlurSlug(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordBlurSlug)
}

func keySalira(address common.Address) string {
	return fmt.Sprint(address.Hex(), keyDelimiter, keywordSalira)
}

func keyMethodSig(methodSig [4]byte) string {
	return fmt.Sprint(keywordMethodSig, keyDelimiter, fmt.Sprintf("%x", methodSig))
}

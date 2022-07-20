package wwatcher

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/wealdtech/go-ens/v3"
)

// var bLog = btv.GetBLog()

var (
	addressesWithoutENSName = make(map[common.Address]bool)

	mu = &sync.RWMutex{}
)

// NamesCache Collection represents the collections configured by the user.
type NamesCache struct {
	Names map[common.Address]string // ens.Name
	RWMu  *sync.RWMutex
}

// MarshalBinary encodes the Collection into a binary format.
func (nc *NamesCache) MarshalBinary() ([]byte, error) { return json.Marshal(nc) }

// UnmarshalBinary decodes the Collection from a binary format.
func (nc *NamesCache) UnmarshalBinary(data []byte) error { return json.Unmarshal(data, nc) }

// GetENSForAddress returns the ENS name for an address or an empty string if not available.
func GetENSForAddress(nodes *gbnode.NodeCollection, address common.Address, namesCache *NamesCache) string {
	var ensName string

	mu.RLock()
	hasNoENSName := addressesWithoutENSName[address]
	mu.RUnlock()

	if hasNoENSName {
		gbl.Log.Debugf("address %s has no ens name - not checking", address.Hex())

		return ""
	}

	gbl.Log.Debugf("resolving address: %s", address.Hex())

	if ensName = ReverseLookupAndValidate(address, nodes); ensName == "" {
		mu.Lock()
		addressesWithoutENSName[address] = true
		mu.Unlock()

		return ""
	}

	namesCache.RWMu.Lock()
	namesCache.Names[address] = ensName
	namesCache.RWMu.Unlock()

	return ensName
}

func ReverseLookupAndValidate(address common.Address, nodes *gbnode.NodeCollection) string {
	var name, ensName string

	var err error

	client := nodes.GetRandomNode().Client

	// lookup the ens name for an address
	name, err = ens.ReverseResolve(client, address)
	if err != nil || common.IsHexAddress(name) {
		gbl.Log.Debugf("ens reverse resolve error: %s -> %s: %s", address, name, err)

		return ""
	}

	// do a lookup for the name to validate its authenticity
	resolvedAddress, err := ens.Resolve(client, name)
	if err != nil {
		gbl.Log.Warnf("ens resolve error: %s -> %s: %s", name, address, err)

		return ""
	}

	if resolvedAddress != address {
		gbl.Log.Warnf("addresses do not match: %s != %s", resolvedAddress.Hex(), address.Hex())
	} else {
		ensName = name

		if viper.GetBool("redis.enabled") {
			if rdb := cache.GetRedisClient(); rdb != nil {
				if err := rdb.SetEX(context.Background(), cache.KeyENS(address), ensName, viper.GetDuration("cache.ens_ttl")).Err(); err != nil {
					if errors.Is(err, redis.Nil) {
						gbl.Log.Warnf("redis | redis.Nil while adding to ensName: %s", err)
					} else {
						gbl.Log.Errorf("redis | could not add ensName: %s", err)
					}
				} else {
					gbl.Log.Infof("redis | added ensName: %s -> %s", ensName, address.Hex())
				}
			}
		}
	}

	return ensName
}

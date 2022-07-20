package collections

import (
	"sort"
	"sync"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/benleb/gloomberg/internal/hooks"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type CollectionSource int64

const (
	Configuration CollectionSource = iota
	Wallet
	Stream
)

type AddressCollection []common.Address

// Contains returns true if the given string is in the slice.
func (ac *AddressCollection) Contains(address common.Address) bool {
	for _, collectionAddress := range *ac {
		if address == collectionAddress {
			return true
		}
	}

	return false
}

type Collections struct {
	UserCollections map[common.Address]*GbCollection
	// DiscoveredCollections map[common.Address]*GbCollection

	// 'queue' to store collections to be processed
	// ListingAddresses chan common.Address
	RWMu *sync.RWMutex
}

func New() *Collections {
	return &Collections{
		UserCollections: make(map[common.Address]*GbCollection),
		RWMu:            &sync.RWMutex{},
	}
}

func (cs *Collections) Addresses() []common.Address {
	addresses := make([]common.Address, 0)
	addresses = append(addresses, cs.UserCollectionsAddresses()...)

	return addresses
}

// OpenseaSlugs returns a slice of slugs for collections with enabled listings.
func (cs *Collections) OpenseaSlugs() []string {
	slugs := make([]string, 0)

	for _, c := range cs.UserCollections {
		if c.Metadata != nil { // } && c.Show.Listings {
			if slug := c.Metadata.OpenseaSlug; slug != "" {
				slugs = append(slugs, c.Metadata.OpenseaSlug)
			}
		}
	}

	return slugs
}

// OpenseaSlugs returns a slice of slugs for collections with enabled listings.
func (cs *Collections) ListingsAddresses() []common.Address {
	addresses := make([]common.Address, 0)

	for _, c := range cs.UserCollections {
		if c.Show.Listings {
			addresses = append(addresses, c.ContractAddress)
		}
	}

	return addresses
}

func (cs *Collections) UserCollectionsAddresses() []common.Address {
	addresses := make([]common.Address, 0)
	for _, c := range cs.UserCollections {
		addresses = append(addresses, c.ContractAddress)
	}

	return addresses
}

func (cs *Collections) userCollectionNames() []string {
	namesIndex := make(map[string]bool, 0)
	names := make([]string, 0)

	for _, c := range cs.UserCollections {
		if !namesIndex[c.Name] {
			namesIndex[c.Name] = true

			names = append(names, c.Name)
		}
	}

	return names
}

func (cs *Collections) colorsByName() map[string]lipgloss.Color {
	colorNames := make(map[string]lipgloss.Color, 0)
	for _, c := range cs.UserCollections {
		colorNames[c.Name] = c.Colors.Primary
	}

	return colorNames
}

func (cs *Collections) SortedAndColoredNames() []string {
	names := make([]string, 0)
	colorNames := cs.colorsByName()

	keys := cs.userCollectionNames()

	sort.Strings(keys)

	for _, name := range keys {
		collectionStyle := lipgloss.NewStyle().Foreground(colorNames[name])
		names = append(names, collectionStyle.Render(name))
	}

	return names
}

func GetCollectionsFromConfiguration(nodes *gbnode.NodeCollection) []*GbCollection {
	collections := make([]*GbCollection, 0)

	if viper.IsSet("collections") {
		gbl.Log.Infof("config | reading collections from config")

		for address, collection := range viper.GetStringMap("collections") {
			contractAddress := common.HexToAddress(address)
			currentCollection := NewCollection(contractAddress, "", nodes, Configuration)

			if collection == nil && common.IsHexAddress(address) {
				gbl.Log.Infof("reading collection without details: %+v", address)

				currentCollection = NewCollection(contractAddress, "", nodes, Configuration)
				// global settings
				currentCollection.Show.Listings = viper.GetBool("show.listings")
				currentCollection.Show.Sales = viper.GetBool("show.sales")
				currentCollection.Show.Mints = viper.GetBool("show.mints")
				currentCollection.Show.Transfers = viper.GetBool("show.transfers")
			} else {
				gbl.Log.Infof("reading collection: %+v - %+v", address, collection)

				decodeHooks := mapstructure.ComposeDecodeHookFunc(
					hooks.StringToAddressHookFunc(),
					hooks.StringToDurationHookFunc(),
					hooks.StringToLipglossColorHookFunc(),
				)

				decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
					DecodeHook: decodeHooks,
					Result:     &currentCollection,
				})

				err := decoder.Decode(collection)
				if err != nil {
					gbl.Log.Errorf("error decoding collection: %+v", err)

					continue
				}
			}

			gbl.Log.Debugf("currentCollection: %+v", currentCollection)

			collections = append(collections, currentCollection)
		}
	}

	return collections
}

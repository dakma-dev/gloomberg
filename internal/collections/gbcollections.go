package collections

import (
	"sort"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
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

type CollectionDB struct {
	Collections map[common.Address]*GbCollection
	// DiscoveredCollections map[common.Address]*GbCollection

	// 'queue' to store collections to be processed
	// ListingAddresses chan common.Address
	RWMu *sync.RWMutex
}

func New() *CollectionDB {
	return &CollectionDB{
		Collections: make(map[common.Address]*GbCollection),
		RWMu:        &sync.RWMutex{},
	}
}

func (cs *CollectionDB) Addresses() []common.Address {
	addresses := make([]common.Address, 0)
	addresses = append(addresses, cs.UserCollectionsAddresses()...)

	return addresses
}

// OpenseaSlugs returns a slice of slugs for collections with enabled listings.
func (cs *CollectionDB) OpenseaSlugs() []string {
	slugs := make([]string, 0)

	for _, c := range cs.Collections {
		if slug := c.OpenseaSlug; slug != "" {
			slugs = append(slugs, c.OpenseaSlug)
		}
	}

	return slugs
}

// ListingsAddresses returns a slice of addresses.
func (cs *CollectionDB) ListingsAddresses() []common.Address {
	addresses := make([]common.Address, 0)

	for _, c := range cs.Collections {
		if c.Show.Listings {
			addresses = append(addresses, c.ContractAddress)
		}
	}

	return addresses
}

func (cs *CollectionDB) UserCollectionsAddresses() []common.Address {
	addresses := make([]common.Address, 0)
	for _, c := range cs.Collections {
		addresses = append(addresses, c.ContractAddress)
	}

	return addresses
}

func (cs *CollectionDB) userCollectionNames() []string {
	namesIndex := make(map[string]bool, 0)
	names := make([]string, 0)

	for _, c := range cs.Collections {
		if !namesIndex[c.Name] {
			namesIndex[c.Name] = true

			names = append(names, c.Name)
		}
	}

	return names
}

func (cs *CollectionDB) colorsByName() map[string]lipgloss.Color {
	colorNames := make(map[string]lipgloss.Color, 0)
	for _, c := range cs.Collections {
		colorNames[c.Name] = c.Colors.Primary
	}

	return colorNames
}

func (cs *CollectionDB) SortedAndColoredNames() []string {
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

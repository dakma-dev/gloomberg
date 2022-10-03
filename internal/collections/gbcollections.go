package collections

import (
	"sort"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type CollectionSource int64

const (
	Configuration CollectionSource = iota
	Wallet
	Stream
)

func (cs *CollectionSource) MarshalJSON() ([]byte, error) {
	return []byte(`"` + cs.String() + `"`), nil
}

func (cs *CollectionSource) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"configuration"`:
		*cs = Configuration
	case `"wallet"`:
		*cs = Wallet
	case `"stream"`:
		*cs = Stream
	}

	return nil
}

func (cs *CollectionSource) String() string {
	return map[CollectionSource]string{
		Configuration: "configuration",
		Wallet:        "wWallet",
		Stream:        "stream",
	}[*cs]
}

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
		if slug := c.OpenseaSlug; slug != "" {
			slugs = append(slugs, c.OpenseaSlug)
		}
	}

	return slugs
}

// ListingsAddresses returns a slice of addresses.
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

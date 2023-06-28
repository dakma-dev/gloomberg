package collections

import (
	"sort"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type CollectionDB struct {
	Collections map[common.Address]*Collection
	// DiscoveredCollections map[common.Address]*Collection

	// 'queue' to store collections to be processed
	// ListingAddresses chan common.Address
	RWMu *sync.RWMutex
}

func New() *CollectionDB {
	return &CollectionDB{
		Collections: make(map[common.Address]*Collection),
		RWMu:        &sync.RWMutex{},
	}
}

func (cs *CollectionDB) Addresses() []common.Address {
	addresses := make([]common.Address, 0)
	addresses = append(addresses, cs.UserCollectionsAddresses()...)

	return addresses
}

// OpenSeaSlugsAndAddresses returns a slice-to-address map of collections with enabled listings.
func (cs *CollectionDB) OpenSeaSlugsAndAddresses() map[string]common.Address {
	// slugs := make([]string, 0)
	slugAddresses := make(map[string]common.Address)

	for _, c := range cs.Collections {
		if slug := c.OpenseaSlug; slug != "" && !c.IgnorePrinting { // && c.Show.Listings {
			slugAddresses[slug] = c.ContractAddress
		}
	}

	return slugAddresses
}

// GetCollectionForSlug returns a collections for the given slug.
func (cs *CollectionDB) GetCollectionForSlug(slug string) *Collection {
	for _, c := range cs.Collections {
		if c.OpenseaSlug == slug {
			return c
		}
	}

	return nil
}

// OpenseaAddressToSlug returns a collectionAddressto-slug map of collections with enabled listings.
func (cs *CollectionDB) OpenseaAddressToSlug() map[common.Address]string {
	// slugs := make([]string, 0)
	slugAddresses := make(map[common.Address]string)

	for _, c := range cs.Collections {
		if slug := c.OpenseaSlug; slug != "" { // && c.Show.Listings {
			slugAddresses[c.ContractAddress] = slug
		}
	}

	return slugAddresses
}

// OpenseaSlugs returns a slice of slugs for collections with enabled listings.
func (cs *CollectionDB) OpenseaSlugs() []string {
	slugs := make([]string, 0)

	for s := range cs.OpenSeaSlugsAndAddresses() {
		slugs = append(slugs, s)
	}

	return slugs
}

func (cs *CollectionDB) OpenseaSlugAddresses() []common.Address {
	addresses := make([]common.Address, 0)

	for _, a := range cs.OpenSeaSlugsAndAddresses() {
		addresses = append(addresses, a)
	}

	return addresses
}

// ListingsAddresses returns a slice of addresses.
func (cs *CollectionDB) ListingsAddresses() []common.Address {
	addresses := make([]common.Address, 0)

	for _, c := range cs.Collections {
		// if c.Show.Listings {
		if c.FetchListings {
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
	namesIndex := make(map[string]bool)
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
	colorNames := make(map[string]lipgloss.Color)
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

package collections

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/VividCortex/ewma"
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/benleb/gloomberg/internal/nemo/collectionsource"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type BaseCollection struct{}

// Collection represents the collections configured by the user.
type Collection struct {
	ContractAddress common.Address `mapstructure:"address"`
	Name            string         `mapstructure:"name"`
	OpenseaSlug     string         `mapstructure:"slug"`

	FetchListings  bool `mapstructure:"fetchListings"`
	IgnorePrinting bool `mapstructure:"ignore"`

	Show struct {
		Sales     bool `mapstructure:"sales"`
		Mints     bool `mapstructure:"mints"`
		Transfers bool `mapstructure:"transfers"`
		Listings  bool `mapstructure:"listings"`
		History   bool `mapstructure:"history"`
	} `mapstructure:"show"`

	Highlight struct {
		Color              lipgloss.Color `mapstructure:"color"`
		Sales              lipgloss.Color `mapstructure:"show.sales"`
		Mints              lipgloss.Color `mapstructure:"mints"`
		Transfers          lipgloss.Color `mapstructure:"transfers"`
		Listings           lipgloss.Color `mapstructure:"listings"`
		ListingsBelowPrice float64        `mapstructure:"listings_below_price"`
	} `mapstructure:"highlight"`

	//
	// calculated/generated fields
	Metadata *nemo.CollectionMetadata `mapstructure:"metadata"`

	Source collectionsource.CollectionSource `mapstructure:"source"`

	Colors struct {
		Primary   lipgloss.Color `mapstructure:"primary"`
		Secondary lipgloss.Color `mapstructure:"secondary"`
	} `mapstructure:"colors"`

	Counters struct {
		Sales       uint64
		Mints       uint64
		Transfers   uint64
		Listings    uint64
		SalesVolume *big.Int
		MintVolume  *big.Int
	} `mapstructure:"counters"`

	SaLiRa                 ewma.MovingAverage  `mapstructure:"salira"`
	FloorPrice             *ewma.MovingAverage `mapstructure:"floorPrice"`
	PreviousFloorPrice     float64             `mapstructure:"previousFloorPrice"`
	HighestCollectionOffer float64
}

func NewCollection(contractAddress common.Address, name string, nodes *provider.Pool, source collectionsource.CollectionSource, rueidi *rueidica.Rueidica) *Collection {
	var collectionName string

	ctx := context.Background()

	switch {
	case name != "":
		collectionName = name
	case contractAddress == internal.ENSContractAddress, contractAddress == internal.ENSNameWrapperContractAddress:
		collectionName = "ENS"
	default:
		name, err := rueidi.GetCachedContractName(ctx, contractAddress)

		switch {
		case errors.Is(err, nil):
			gbl.Log.Debugf("cache | cached collection name: %s", name)

			if name != "" {
				collectionName = name
			}

		case nodes != nil:
			if name, err := nodes.ERC721CollectionName(ctx, contractAddress); err == nil {
				gbl.Log.Debugf("chain | collection name via chain call: %s", name)

				if name != "" {
					collectionName = name
				}

				// cache collection name
				// gbCache.CacheCollectionName(contractAddress, collectionName)
				err := rueidi.StoreContractName(ctx, contractAddress, collectionName)
				if err != nil {
					gbl.Log.Errorf("error storing contract name: %s | %s", style.ShortenAddress(&contractAddress), err)
				}
			}

		default:
			gbl.Log.Errorf("error getting collection name, using: %s | %s", style.ShortenAddress(&contractAddress), err)

			collectionName = style.ShortenAddress(&contractAddress)
		}
	}

	floorPrice := ewma.NewMovingAverage()

	collection := Collection{
		Name:            collectionName,
		ContractAddress: contractAddress,
		Metadata:        &nemo.CollectionMetadata{},

		Source: source,

		FloorPrice:             &floorPrice,
		PreviousFloorPrice:     0,
		SaLiRa:                 ewma.NewMovingAverage(),
		HighestCollectionOffer: 0,
	}

	if nodes != nil {
		go func() {
			rawMetaDatas, err := nodes.ERC721CollectionMetadata(context.Background(), contractAddress)
			if err != nil {
				gbl.Log.Errorf("error getting collection metadata, using: %s | %s", style.ShortenAddress(&contractAddress), err)

				return
			}

			metadata := &nemo.CollectionMetadata{}

			if name := rawMetaDatas["name"]; name != nil {
				name, ok := name.(string)
				if ok {
					metadata.ContractName = name
				}
			}

			if symbol := rawMetaDatas["symbol"]; symbol != nil {
				symbol, ok := symbol.(string)
				if ok {
					metadata.Symbol = symbol
				}
			}

			if totalSupply := rawMetaDatas["totalSupply"]; totalSupply != nil {
				totalSupply, ok := totalSupply.(uint64)
				if ok {
					metadata.TotalSupply = totalSupply
				}
			}

			if tokenURI := rawMetaDatas["tokenURI"]; tokenURI != nil {
				tokenURI, ok := tokenURI.(string)
				if ok {
					metadata.TokenURI = tokenURI
				}
			}

			collection.Metadata = metadata
		}()
	}

	if source == collectionsource.FromWallet || source == collectionsource.FromConfiguration {
		collection.Show.History = true
	}

	if source == collectionsource.FromWallet || source == collectionsource.FromStream {
		collection.Show.Sales = viper.GetBool("show.sales")
		collection.Show.Mints = viper.GetBool("show.mints")
		collection.Show.Transfers = viper.GetBool("show.transfers")

		if source == collectionsource.FromWallet {
			if viper.IsSet("api_keys.opensea") {
				collection.Show.Listings = viper.GetBool("listings.enabled")
			}
		}

		if source == collectionsource.FromStream {
			collection.Show.Listings = false
			collection.Show.History = false
		}
	}

	// generate the collection color based on the contract address if none given
	collection.generateColorsFromAddress()

	// initialize the counters
	collection.ResetStats()

	return &collection
}

func (uc *Collection) Style() lipgloss.Style {
	if uc.Colors.Primary == "" {
		gbl.Log.Infof("🎨 primary collection color missing for %s", uc.Name)
		uc.generateColorsFromAddress()
	}

	return lipgloss.NewStyle().Foreground(uc.Colors.Primary)
}

func (uc *Collection) StyleSecondary() lipgloss.Style {
	if uc.Colors.Secondary == "" {
		gbl.Log.Infof("🎨 secondary collection color missing for %s", uc.Name)
		uc.generateColorsFromAddress()
	}

	return lipgloss.NewStyle().Foreground(uc.Colors.Secondary)
}

func (uc *Collection) Render(text string) string {
	// generate the collection color based on the contract address if none given
	return uc.Style().Render(text)
}

func (uc *Collection) AddSale(value *big.Int, numItems uint64) float64 {
	uc.Counters.SalesVolume.Add(uc.Counters.SalesVolume, value)
	atomic.AddUint64(&uc.Counters.Sales, numItems)

	return float64((uc.Counters.Sales * 60) / uint64(viper.GetDuration("ticker.statsbox").Seconds()))
}

func (uc *Collection) AddMint() {
	atomic.AddUint64(&uc.Counters.Mints, 1)
}

func (uc *Collection) AddMintVolume(value *big.Int, numItems uint64) {
	atomic.AddUint64(&uc.Counters.Mints, numItems)
	uc.Counters.MintVolume.Add(uc.Counters.MintVolume, value)
}

func (uc *Collection) AddListing(numItems uint64) {
	atomic.AddUint64(&uc.Counters.Listings, numItems)
}

// CalculateSaLiRa updates the salira moving average of a given collection.
func (uc *Collection) CalculateSaLiRa(address common.Address, rueidica *rueidica.Rueidica) (float64, float64) {
	if uc.Counters.Listings <= 0 {
		return 0.0, 0.0
	}

	previousSaLiRa := uc.SaLiRa.Value()
	uc.SaLiRa.Add(float64(uc.Counters.Sales) / float64(uc.Counters.Listings))
	currentSaLiRa := uc.SaLiRa.Value()

	if address != internal.ZeroAddress {
		// go cache.StoreSalira(context.TODO(), address, currentSaLiRa)
		go rueidica.StoreSalira(context.Background(), address, currentSaLiRa)
	}

	return previousSaLiRa, currentSaLiRa
}

// CalculateFloorPrice updates the moving average of a given collection.
func (uc *Collection) CalculateFloorPrice(tokenPrice float64) (float64, float64) {
	// update the moving average
	uc.PreviousFloorPrice = (*uc.FloorPrice).Value()
	(*uc.FloorPrice).Add(tokenPrice)
	currentFloorPrice := (*uc.FloorPrice).Value()

	gbl.Log.Debugf("uc.PreviousFloorPrice: %f  |  currentFloorPrice: %f | uc.FloorPrice.Value(): %f", uc.PreviousFloorPrice, currentFloorPrice, (*uc.FloorPrice).Value())

	return uc.PreviousFloorPrice, currentFloorPrice
}

func (uc *Collection) ResetStats() {
	gbl.Log.Debugf("resetting collection statistics...")

	uc.Counters.Sales = 0
	uc.Counters.Mints = 0
	uc.Counters.Transfers = 0
	uc.Counters.Listings = 0
	uc.Counters.SalesVolume = big.NewInt(0)
	uc.Counters.MintVolume = big.NewInt(0)
}

// GenerateColors generates two colors based on contract address of the collection.
func (uc *Collection) generateColorsFromAddress() {
	if uc.Colors.Primary == "" {
		uc.Colors.Primary = style.GenerateColorWithSeed(uc.ContractAddress.Hash().Big().Int64())
	}

	if uc.Colors.Secondary == "" {
		uc.Colors.Secondary = style.GenerateColorWithSeed(uc.ContractAddress.Big().Int64() ^ 2)
	}
}

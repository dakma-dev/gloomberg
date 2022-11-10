package collections

import (
	"fmt"
	"math/big"
	"math/rand"
	"sync/atomic"

	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"

	"github.com/VividCortex/ewma"
	"github.com/benleb/gloomberg/internal/cache"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type BaseCollection struct{}

// GbCollection represents the collections configured by the user.
type GbCollection struct {
	//
	// configurable fields
	ContractAddress common.Address `mapstructure:"address"`
	Name            string         `mapstructure:"name"`
	OpenseaSlug     string         `mapstructure:"slug"`

	// SupportedStandards standard.Standards

	Show struct {
		Sales     bool `mapstructure:"sales"`
		Mints     bool `mapstructure:"mints"`
		Transfers bool `mapstructure:"transfers"`
		Listings  bool `mapstructure:"listings"`
	} `mapstructure:"show"`

	Highlight struct {
		Color              lipgloss.Color `mapstructure:"color"`
		Sales              lipgloss.Color `mapstructure:"show.sales"`
		Mints              lipgloss.Color `mapstructure:"mints"`
		Transfers          lipgloss.Color `mapstructure:"transfers"`
		Listings           lipgloss.Color `mapstructure:"listings.enabled"`
		ListingsBelowPrice float64        `mapstructure:"listings_below_price"`
	}

	//
	// calculated/generated fields
	Metadata *models.CollectionMetadata `mapstructure:"metadata"`

	Source models.CollectionSource `mapstructure:"source"`

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
	}

	SaLiRa ewma.MovingAverage `json:"salira"`

	// exponential moving average of the actual sale prices
	// FloorPrice         ewmabig.MovingAverage `mapstructure:"artificialFloor"`
	// PreviousFloorPrice *big.Int              `mapstructure:"artificialFloor"`

	FloorPrice         ewma.MovingAverage `mapstructure:"floorPrice"`
	PreviousFloorPrice float64            `mapstructure:"previousFloorPrice"`
}

func NewCollection(contractAddress common.Address, name string, nodes *nodes.Nodes, source models.CollectionSource) *GbCollection {
	var collectionName string

	gbCache := cache.New()

	if name != "" {
		collectionName = name
	} else {
		if name, err := gbCache.GetCollectionName(contractAddress); err == nil {
			gbl.Log.Debugf("cache | cached collection name: %s", name)

			if name != "" {
				collectionName = name
			}
		} else if nodes != nil {
			if name, err := nodes.GetERC721CollectionName(contractAddress); err == nil {
				gbl.Log.Debugf("chain | collection name via chain call: %s", name)

				if name != "" {
					collectionName = name
				}

				// cache collection name
				gbCache.CacheCollectionName(contractAddress, collectionName)
			}
		} else {
			gbl.Log.Errorf("error getting collection name, using: %s | %s", style.ShortenAddress(&contractAddress), err)

			collectionName = style.ShortenAddress(&contractAddress)
		}
	}

	collection := GbCollection{
		ContractAddress: contractAddress,
		Name:            collectionName,

		// OwnedTokenIDs: []uint64{},
		Metadata: &models.CollectionMetadata{},
		Source:   source,

		FloorPrice:         ewma.NewMovingAverage(),
		PreviousFloorPrice: 0,

		SaLiRa: ewma.NewMovingAverage(),
	}

	// go func() {
	// 	collection.SupportedStandards = nodes.GetSupportedStandards(contractAddress)
	// }()

	// go func() {
	// 	if nodes.ERC1155Supported(contractAddress) {
	// 		collection.SupportedStandards = append(collection.SupportedStandards, standard.ERC1155)
	// 	}
	// }()

	if nodes != nil {
		go func() {
			rawMetaDatas := nodes.GetCollectionMetadata(contractAddress)

			metadata := &models.CollectionMetadata{}

			if name := rawMetaDatas["name"]; name != nil {
				metadata.ContractName = name.(string)
			}

			if symbol := rawMetaDatas["symbol"]; symbol != nil {
				metadata.Symbol = symbol.(string)
			}

			if totalSupply := rawMetaDatas["totalSupply"]; totalSupply != nil {
				metadata.TotalSupply = totalSupply.(uint64)
			}

			if tokenURI := rawMetaDatas["tokenURI"]; tokenURI != nil {
				metadata.TokenURI = tokenURI.(string)
			}

			collection.Metadata = metadata
		}()
	}

	if source == models.FromWallet || source == models.FromStream {
		collection.Show.Sales = viper.GetBool("show.sales")
		collection.Show.Mints = viper.GetBool("show.mints")
		collection.Show.Transfers = viper.GetBool("show.transfers")

		if source == models.FromWallet {
			if viper.IsSet("api_keys.opensea") {
				collection.Show.Listings = viper.GetBool("listings.enabled")
			}
		}

		if source == models.FromStream {
			collection.Show.Listings = false
		}
	}

	// generate the collection color based on the contract address if none given
	collection.generateColorsFromAddress()

	// initialize the counters
	collection.ResetStats()

	return &collection
}

func (uc *GbCollection) Style() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(uc.Colors.Primary)
}

func (uc *GbCollection) StyleSecondary() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(uc.Colors.Secondary)
}

func (uc *GbCollection) Render(text string) string {
	// generate the collection color based on the contract address if none given
	return lipgloss.NewStyle().Foreground(uc.Colors.Primary).Render(text)
}

func (uc *GbCollection) AddSale(value *big.Int, numItems uint64) float64 {
	uc.Counters.SalesVolume.Add(uc.Counters.SalesVolume, value)
	atomic.AddUint64(&uc.Counters.Sales, numItems)

	return float64((uc.Counters.Sales * 60) / uint64(viper.GetDuration("ticker.statsbox").Seconds()))
}

func (uc *GbCollection) AddMint() {
	atomic.AddUint64(&uc.Counters.Mints, 1)
}

func (uc *GbCollection) SaLiRaAdd() float64 {
	if uc.Counters.Listings > 0 {
		return float64(uc.Counters.Sales) / float64(uc.Counters.Listings)
	}

	return 0.0
}

// CalculateSaLiRa updates the salira moving average of a given collection.
func (uc *GbCollection) CalculateSaLiRa() (float64, float64) {
	if uc.Counters.Listings <= 0 {
		return 0.0, 0.0
	}

	previousSaLiRa := uc.SaLiRa.Value()
	uc.SaLiRa.Add(float64(uc.Counters.Sales) / float64(uc.Counters.Listings))
	currentSaLiRa := uc.SaLiRa.Value()

	return previousSaLiRa, currentSaLiRa
}

// CalculateFloorPrice updates the moving average of a given collection.
func (uc *GbCollection) CalculateFloorPrice(tokenPrice float64) (float64, float64) {
	// update the moving average
	uc.PreviousFloorPrice = uc.FloorPrice.Value()
	uc.FloorPrice.Add(tokenPrice)
	currentFloorPrice := uc.FloorPrice.Value()

	return uc.PreviousFloorPrice, currentFloorPrice
}

func (uc *GbCollection) ResetStats() {
	gbl.Log.Debugf("resetting collection statistics...")

	uc.Counters.Sales = 0
	uc.Counters.Mints = 0
	uc.Counters.Transfers = 0
	uc.Counters.Listings = 0
	uc.Counters.SalesVolume = big.NewInt(0)
}

// GenerateColors generates two colors based on contract address of the collection.
func (uc *GbCollection) generateColorsFromAddress() {
	rand.Seed(uc.ContractAddress.Hash().Big().Int64())

	//nolint:gosec
	primary := lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", rand.Intn(256), rand.Intn(256), rand.Intn(256)))
	//nolint:gosec
	secondary := lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", rand.Intn(256), rand.Intn(256), rand.Intn(256)))

	if uc.Colors.Primary == "" {
		uc.Colors.Primary = primary
	}

	if uc.Colors.Secondary == "" {
		uc.Colors.Secondary = secondary
	}
}

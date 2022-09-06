package collections

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"sync/atomic"

	"github.com/VividCortex/ewma"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// GbCollection represents the collections configured by the user.
type GbCollection struct {
	//
	// user-configurable fields
	//

	ContractAddress common.Address `mapstructure:"address"`
	Name            string         `mapstructure:"name"`
	OpenseaSlug     string         `mapstructure:"slug"`

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
		Listings           lipgloss.Color `mapstructure:"show.listings"`
		ListingsBelowPrice float64        `mapstructure:"listings_below_price"`
	}

	//
	// calculated/generated fields
	//

	Metadata *gbnode.CollectionMetadata `mapstructure:"metadata"`

	Source CollectionSource `mapstructure:"source"`

	Colors struct {
		Primary   lipgloss.Color `mapstructure:"primary"`
		Secondary lipgloss.Color `mapstructure:"secondary"`
	} `mapstructure:"colors"`

	Counters struct {
		Sales       uint64   `mapstructure:"sales"`
		Mints       uint64   `mapstructure:"mints"`
		Transfers   uint64   `mapstructure:"transfers"`
		Listings    uint64   `mapstructure:"listings"`
		SalesVolume *big.Int `mapstructure:"salesVolume"`
	} `mapstructure:"counters"`

	SaLiRa ewma.MovingAverage `mapstructure:"salira"`

	// exponential moving average of the actual sale prices
	ArtificialFloor         ewma.MovingAverage `mapstructure:"artificialFloor"`
	PreviousArtificialFloor float64            `mapstructure:"artificialFloor"`
}

var ctx = context.Background()

// // MarshalBinary encodes the Collection into a binary format.
// func (uc *Collection) MarshalBinary() ([]byte, error) { return json.Marshal(uc) }

// // UnmarshalBinary decodes the Collection from a binary format.
// func (uc *Collection) UnmarshalBinary(data []byte) error { return json.Unmarshal(data, uc) }

func NewCollection(contractAddress common.Address, name string, nodes *gbnode.NodeCollection, source CollectionSource) *GbCollection {
	var collectionName string

	gbCache := cache.New(ctx)

	if name != "" {
		collectionName = name
	} else {
		if name, err := gbCache.GetCollectionName(contractAddress); err == nil {
			gbl.Log.Infof("cache | cached collection name: %s", name)

			if name != "" {
				collectionName = name
			}
		} else if name, err := nodes.GetRandomNode().GetCollectionName(contractAddress); err == nil {
			gbl.Log.Infof("chain | collection name via contract call: %s", name)

			if name != "" {
				collectionName = name
			}

			// cache collection name
			gbCache.CacheCollectionName(contractAddress, collectionName)
		} else {
			gbl.Log.Errorf("error getting collection name, using: %s | %s", style.ShortenAddress(&contractAddress), err)

			collectionName = style.ShortenAddress(&contractAddress)
		}
	}

	collection := GbCollection{
		ContractAddress: contractAddress,
		Name:            collectionName,

		Metadata: &gbnode.CollectionMetadata{},
		Source:   source,

		ArtificialFloor: ewma.NewMovingAverage(),
		SaLiRa:          ewma.NewMovingAverage(),
	}

	go func() {
		collection.Metadata = nodes.GetRandomNode().GetCollectionMetadata(contractAddress)
	}()

	if source == Wallet || source == Stream {
		collection.Show.Sales = viper.GetBool("show.sales")
		collection.Show.Mints = viper.GetBool("show.mints")
		collection.Show.Transfers = viper.GetBool("show.transfers")

		if source == Wallet {
			if viper.IsSet("api_keys.opensea") {
				collection.Show.Listings = viper.GetBool("show.listings")
			}
		}

		if source == Stream {
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

	return float64((uc.Counters.Sales * 60) / uint64(viper.GetDuration("stats.interval").Seconds()))
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

// CalculateArtificialFloor updates the moving average of a given collection.
func (uc *GbCollection) CalculateArtificialFloor(tokenPrice float64) (float64, float64) {
	// update the moving average
	uc.PreviousArtificialFloor = uc.ArtificialFloor.Value()
	uc.ArtificialFloor.Add(tokenPrice)
	currentMovingAverage := uc.ArtificialFloor.Value()

	return uc.PreviousArtificialFloor, currentMovingAverage
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

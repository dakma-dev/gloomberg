package collections

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync/atomic"
	"time"

	"github.com/VividCortex/ewma"
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/benleb/gloomberg/internal/nemo/collectionsource"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
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

	// better "counters"
	RecentEvents mapset.Set[*degendb.RecentEvent] `mapstructure:"recent_events"`
	degendb.SaLiRas

	Counters struct {
		Mints       uint64
		Transfers   uint64
		SalesVolume *big.Int
		MintVolume  *big.Int
	} `mapstructure:"counters"`

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

	saliraTimeframe, ok := viper.Get("salira.timeframes").([]time.Duration)
	if !ok {
		gbl.Log.Errorf("error getting SaLiRa timeframes")
	}

	collection := Collection{
		Name:            collectionName,
		ContractAddress: contractAddress,
		Metadata:        &nemo.CollectionMetadata{},

		Source: source,

		RecentEvents: mapset.NewSet[*degendb.RecentEvent](),
		SaLiRas:      degendb.NewSaLiRas(saliraTimeframe),

		FloorPrice:             &floorPrice,
		PreviousFloorPrice:     0,
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

func (uc *Collection) prettyOpenseaSlug() string {
	if uc.OpenseaSlug == "" {
		return ""
	}

	return uc.Render(uc.OpenseaSlug)
}

// GetPrettySaLiRas returns freshly calculated and beautifully formatted SaLiRas for all configured timeframes.
func (uc *Collection) GetPrettySaLiRas() []string {
	fmtSaLiRas := make([]string, 0)

	log.Debugf("%s | saliras: %+v", uc.prettyOpenseaSlug(), uc.SaLiRas)

	// sort by timeframe
	sort.Slice(uc.SaLiRas, func(i, j int) bool {
		return uc.SaLiRas[i].Timeframe < uc.SaLiRas[j].Timeframe
	})

	for _, salira := range uc.SaLiRas {
		// is gloomberg running long enough to have meaningful numbers for this timeframe?
		if time.Since(internal.RunningSince) > salira.Timeframe {
			continue
		}

		// get sale/listing counts for the current timeframe
		sales, listings := uc.getSaLiCountWithTimeframe(salira.Timeframe)

		// no numbers, no saLiRa 🤷‍♀️
		if sales == 0 || listings == 0 {
			break
		}

		log.Debugf("%s | salira %+v: %d / %d", uc.prettyOpenseaSlug(), salira.Timeframe, sales, listings)

		// 🧮 calculate the saLiRa
		salira.Previous = salira.Value()
		salira.Add(float64(sales) / float64(listings))

		// only add the salira if it is > 0
		// if current := salira.Value(); current > 0 {
		fmtSaLiRas = append(fmtSaLiRas, style.DarkGrayStyle.Render(fmt.Sprintf("%.0f", salira.Timeframe.Minutes())+"|"+salira.Pretty()))
		// }
	}

	return fmtSaLiRas
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

func (uc *Collection) GetSaLiCount() (int, int) {
	return uc.getSaLiCountWithTimeframe(viper.GetDuration("salira.default_timeframe"))
}

func (uc *Collection) getSaLiCountWithTimeframe(timeframe time.Duration) (int, int) {
	var recentSales, recentListings int

	for _, event := range uc.RecentEvents.ToSlice() {
		if time.Since(event.Timestamp) < timeframe || timeframe == 0 {
			switch event.Type {
			case degendb.Sale:
				recentSales++
			case degendb.Listing:
				recentListings++
			}
		}
	}

	return recentSales, recentListings
}

func (uc *Collection) AddSales(value *big.Int, numItems uint64) {
	uc.Counters.SalesVolume.Add(uc.Counters.SalesVolume, value)
	// atomic.AddUint64(&uc.Counters.Sales, numItems)

	uc.RecentEvents.Add(&degendb.RecentEvent{
		Timestamp: time.Now(),
		Type:      degendb.Sale,

		AmountWei:    value,
		AmountTokens: numItems,
	})
}

func (uc *Collection) AddMint() {
	atomic.AddUint64(&uc.Counters.Mints, 1)
}

func (uc *Collection) AddMintVolume(value *big.Int, numItems uint64) {
	atomic.AddUint64(&uc.Counters.Mints, numItems)
	uc.Counters.MintVolume.Add(uc.Counters.MintVolume, value)
}

func (uc *Collection) AddListing(numItems uint64) {
	// atomic.AddUint64(&uc.Counters.Listings, numItems)

	uc.RecentEvents.Add(&degendb.RecentEvent{
		Timestamp: time.Now(),
		Type:      degendb.Listing,

		// AmountWei:    value,
		AmountTokens: numItems,
	})
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

	uc.Counters.Mints = 0
	uc.Counters.Transfers = 0
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

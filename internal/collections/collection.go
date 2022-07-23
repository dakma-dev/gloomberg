package collections

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/VividCortex/ewma"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
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

	// exponential moving average of the actuall sale prices
	ArtificialFloor         ewma.MovingAverage `mapstructure:"artificialFloor"`
	PreviousArtificialFloor float64            `mapstructure:"artificialFloor"`
}

var ctx = context.Background()

// // MarshalBinary encodes the Collection into a binary format.
// func (uc *Collection) MarshalBinary() ([]byte, error) { return json.Marshal(uc) }

// // UnmarshalBinary decodes the Collection from a binary format.
// func (uc *Collection) UnmarshalBinary(data []byte) error { return json.Unmarshal(data, uc) }

func NewCollection(contractAddress common.Address, name string, nodes *gbnode.NodeCollection, source CollectionSource) *GbCollection {
	// name
	var collectionName string

	if name != "" {
		collectionName = name
	} else if viper.GetBool("redis.enabled") {
		// check if the collection is already in the redis cache
		if rdb := cache.GetRedisClient(); rdb != nil {
			gbl.Log.Debugf("redis | searching for contract address: %s", contractAddress.String())

			if redisName, err := rdb.Get(ctx, cache.KeyContract(contractAddress)).Result(); err == nil && redisName != "" {
				gbl.Log.Debugf("redis | using cached contractName: %s", redisName)

				collectionName = redisName
			}
		}
	} else {
		collectionName = nodes.GetRandomNode().GetCollectionName(contractAddress)
	}

	// redis
	if viper.GetBool("redis.enabled") {
		if rdb := cache.GetRedisClient(); rdb != nil {
			err := rdb.SetEX(ctx, cache.KeyContract(contractAddress), collectionName, time.Hour*48).Err()

			if err != nil {
				gbl.Log.Infof("redis | error while adding: %s", err.Error())
			} else {
				gbl.Log.Debugf("redis | added: %s -> %s", contractAddress.Hex(), collectionName)
			}
		}
	}

	// if uc.source == Wallet || uc.source == Stream {
	// 	uc.Show.Sales = viper.GetBool("show.sales")
	// 	uc.Show.Mints = viper.GetBool("show.mints")
	// 	uc.Show.Transfers = viper.GetBool("show.transfers")

	// 	if viper.IsSet("api_keys.opensea") {
	// 		uc.Show.Listings = viper.GetBool("show.listings")
	// 	}
	// }

	collection := GbCollection{
		ContractAddress: contractAddress,
		Name:            collectionName,
		// Metadata:        collectionMetadata,
		Metadata:        &gbnode.CollectionMetadata{},
		Source:          source,
		ArtificialFloor: ewma.NewMovingAverage(),
		SaLiRa:          ewma.NewMovingAverage(),
		// Show:            show,
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

	// // initialize the moving averages
	// uc.artificialFloor = ewma.NewMovingAverage()
	// uc.saLiRa = ewma.NewMovingAverage()

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

// CalculateMovingAverage updates the moving average of a given collection.
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

// func GetCollectionMetadata(contractAddress common.Address) *models.CollectionMetadata {
// 	// get the contractERC721 ABIs
// 	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
// 	if err != nil {
// 		gbl.Log.Error(err)
// 	}

// 	// collection name
// 	collectionName := ""

// 	// check if the collection is already in the redis cache
// 	if viper.GetBool("redis.enabled") {
// 		if rdb := cache.GetRedisClient(); rdb != nil {
// 			gbl.Log.Debugf("redis | searching for contract address: %s", contractAddress.String())

// 			if name, err := rdb.Get(ctx, cache.KeyContract(contractAddress)).Result(); err == nil && name != "" {
// 				gbl.Log.Debugf("redis | using cached contractName: %s", name)

// 				collectionName = name
// 			}
// 		}
// 	}

// 	// if not found in redis, we call the contract method to get the name
// 	if collectionName == "" {
// 		if name, err := contractERC721.Name(&bind.CallOpts{}); err == nil {
// 			collectionName = name

// 			if viper.GetBool("redis.enabled") {
// 				if rdb := cache.GetRedisClient(); rdb != nil {
// 					err := rdb.SetEX(ctx, cache.KeyContract(contractAddress), collectionName, time.Hour*48).Err()

// 					if err != nil {
// 						gbl.Log.Infof("redis | error while adding: %s", err.Error())
// 					} else {
// 						gbl.Log.Debugf("redis | added: %s -> %s", contractAddress.Hex(), collectionName)
// 					}
// 				}
// 			}
// 		} else {
// 			collectionName = style.ShortenAddress(&contractAddress)
// 		}
// 	}

// 	// collection total supply
// 	collectionTotalSupply := uint64(0)
// 	if totalSupply, err := contractERC721.TotalSupply(&bind.CallOpts{}); err == nil {
// 		collectionTotalSupply = totalSupply.Uint64()
// 	}

// 	// collection symbol
// 	collectionSymbol := ""
// 	if symbol, err := contractERC721.Symbol(&bind.CallOpts{}); err == nil {
// 		collectionSymbol = symbol
// 	}

// 	// collection token uri
// 	collectionTokenURI := ""
// 	if tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, big.NewInt(1)); err == nil {
// 		collectionTokenURI = tokenURI
// 	}

// 	return &models.CollectionMetadata{
// 		Name:        collectionName,
// 		Symbol:      collectionSymbol,
// 		TotalSupply: collectionTotalSupply,
// 		TokenURI:    collectionTokenURI,
// 	}
// }

// type CollectionMetadata struct {
// 	OpenseaSlug string `mapstructure:"slug"`
// 	Symbol      string `mapstructure:"symbol"`
// 	TotalSupply uint64 `mapstructure:"total_supply"`
// 	TokenURI    string `mapstructure:"token_uri"`
// }

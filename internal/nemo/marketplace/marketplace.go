package marketplace

import (
	"github.com/benleb/gloomberg/internal"
	"github.com/charmbracelet/lipgloss"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
)

type MarketPlace struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Color lipgloss.Color `json:"color"`
	// OldContractAddresses map[common.Address]bool    `json:"contractAddresses"`
	ContractAddresses mapset.Set[common.Address] `json:"-"`
	Tag               string                     `json:"tag"`
}

func Addresses() mapset.Set[common.Address] {
	return mapset.NewSetFromMapKeys[common.Address](AddressToMarketplace())
}

func AddressToMarketplace() map[common.Address]*MarketPlace {
	marketplaces := []*MarketPlace{
		&OpenSea,
		&Blur,
		&X2Y2,
		&LooksRare,
		&SuperRare,
		&NFTfi,
	}

	marketplaceAddresses := make(map[common.Address]*MarketPlace, len(marketplaces))

	for _, mp := range marketplaces {
		for _, addr := range mp.ContractAddresses.ToSlice() {
			marketplaceAddresses[addr] = mp
		}
	}

	return marketplaceAddresses
}

func (mp *MarketPlace) ContractAddress() common.Address {
	var contractAddress common.Address

	for _, addr := range mp.ContractAddresses.ToSlice() {
		contractAddress = addr

		break
	}

	return contractAddress
}

// RenderFaint renders the given text with the marketplace color.
func (mp *MarketPlace) Render(text string) string {
	return lipgloss.NewStyle().Foreground(mp.Color).Render(text)
}

// RenderFaint renders the given text with the marketplace color and faints it.
func (mp *MarketPlace) RenderFaint(text string) string {
	return lipgloss.NewStyle().Foreground(mp.Color).Faint(true).Render(text)
}

// RenderTag renders the marketplace tag with the marketplace color .
func (mp *MarketPlace) RenderTag() string { return mp.Render(mp.Tag) }

// RenderFaintTag renders the marketplace tag with the marketplace color and faints it.
func (mp *MarketPlace) RenderFaintTag() string { return mp.RenderFaint(mp.Tag) }

//
//
// marketplaces
//

var OpenSea = MarketPlace{
	ID:    "opensea",
	Name:  "OpenSea",
	Color: lipgloss.Color("#2C7BE5"),
	ContractAddresses: mapset.NewSet[common.Address](
		common.HexToAddress("0x00000000000001ad428e4906ae43d8f9852d0dd6"),
		common.HexToAddress("0x00000000006c3852cbef3e08e8df289169ede581"), // Seaport 1.1
		common.HexToAddress("0x00005ea00ac477b1030ce78506496e8c2de24bf5"), // SeaDrop
		common.HexToAddress("0x00000000000000adc04c56bf30ac9d3c0aaf14dc"), // Seaport 1.5
	),

	// ContractAddresses: map[common.Address]bool{
	// 	common.HexToAddress("0x00000000000001ad428e4906ae43d8f9852d0dd6"),
	// 	common.HexToAddress("0x00000000006c3852cbef3e08e8df289169ede581"), // Seaport 1.1
	// 	common.HexToAddress("0x00005ea00ac477b1030ce78506496e8c2de24bf5"), // SeaDrop
	// 	common.HexToAddress("0x00000000000000adc04c56bf30ac9d3c0aaf14dc"), // Seaport 1.5
	// },
	Tag: "|",
}

var Blur = MarketPlace{
	ID:    "blur",
	Name:  "Blur",
	Color: lipgloss.Color("#FF8700"),
	// ContractAddresses: map[common.Address]bool{
	ContractAddresses: mapset.NewSet[common.Address](
		common.HexToAddress("0x000000000000Ad05Ccc4F10045630fb830B95127"), // Blur.io: Marketplace
		common.HexToAddress("0x39da41747a83aee658334415666f3ef92dd0d541"), // Blur.io: Marketplace 2
		common.HexToAddress("0xb2ecfe4e4d61f8790bbb9de2d1259b9e2410cea5"), // Blur.io: Marketplace 3
		common.HexToAddress("0x29469395eaf6f95920e59f858042f0e28d98a20b"), // BLEND
		common.HexToAddress("0x0000000000a39bb272e79075ade125fd351887ac"), // Blur Pool Token
	),
	Tag: "|",
}

var X2Y2 = MarketPlace{
	ID:    "x2y2",
	Name:  "X2Y2",
	Color: lipgloss.Color("#acc2fa"),
	ContractAddresses: mapset.NewSet[common.Address](
		common.HexToAddress("0x74312363e45dcaba76c59ec49a7aa8a65a67eed3"),
	),
	Tag: "|",
}

var LooksRare = MarketPlace{
	ID:                "looksrare",
	Name:              "LooksRare",
	Color:             lipgloss.Color("#21E453"),
	ContractAddresses: mapset.NewSet[common.Address](common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a")),
	Tag:               "|",
}

var SuperRare = MarketPlace{
	ID:    "superrare",
	Name:  "SuperRare",
	Color: lipgloss.Color("#eeeeee"),
	ContractAddresses: mapset.NewSet[common.Address](
		common.HexToAddress("0x6D7c44773C52D396F43c2D511B81aa168E9a7a42"),
		common.HexToAddress("0x5e62454d6AA7392925ccd3E7cd33f2D7c2f33D97"),
		common.HexToAddress("0xDd867a8Eb1720185B3fdAD7F81Caed4E8132Be19"),
	),
	Tag: "|",
}

var NFTfi = MarketPlace{
	ID:    "nftfi",
	Name:  "NFTfi",
	Color: lipgloss.Color("#5C4CA0"),
	ContractAddresses: mapset.NewSet[common.Address](
		common.HexToAddress("0x8252Df1d8b29057d1Afe3062bf5a64D503152BC8"),
		common.HexToAddress("0xaDDE73498902F61BfCB702e94C31c13C534879AC"),
		common.HexToAddress("0x5A42d72372858E10Edc03B26bF449F78fF3c0e6F"),
		common.HexToAddress("0x0C90C8B4aa8549656851964d5fB787F0e4F54082"),
		common.HexToAddress("0x5660E206496808F7b5cDB8C56A696a96AE5E9b23"),
		common.HexToAddress("0xe73ECe5988FfF33a012CEA8BB6Fd5B27679fC481"),
		common.HexToAddress("0xE52Cec0E90115AbeB3304BaA36bc2655731f7934"),
	),
	Tag: "↩︎",
}

var Unknown = MarketPlace{
	ID:    "unknown",
	Name:  "Unknown",
	Color: lipgloss.Color("#777777"),
	ContractAddresses: mapset.NewSet[common.Address](
		internal.ZeroAddress,
	),
	Tag: "¦",
}

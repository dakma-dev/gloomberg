package marketplace

import (
	"github.com/benleb/gloomberg/internal"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type MarketPlace struct {
	ID                string                  `json:"id"`
	Name              string                  `json:"name"`
	Color             lipgloss.Color          `json:"color"`
	ContractAddresses map[common.Address]bool `json:"contractAddresses"`
	Tag               string                  `json:"tag"`
}

func (mp *MarketPlace) ContractAddress() common.Address {
	var contractAddress common.Address

	for addr := range mp.ContractAddresses {
		contractAddress = addr

		break
	}

	return contractAddress
}

func (mp *MarketPlace) Render(text string) string {
	// generate the collection color based on the contract address if none given
	return lipgloss.NewStyle().Foreground(mp.Color).Render(text)
}

func (mp *MarketPlace) RenderFaint(text string) string {
	// generate the collection color based on the contract address if none given
	return lipgloss.NewStyle().Foreground(mp.Color).Faint(true).Render(text)
}

func (mp *MarketPlace) RenderTag() string {
	return mp.Render(mp.Tag)
}

func (mp *MarketPlace) RenderFaintTag() string {
	return mp.RenderFaint(mp.Tag)
}

var OpenSea = MarketPlace{
	ID:    "opensea",
	Name:  "OpenSea",
	Color: lipgloss.Color("#2C7BE5"),
	ContractAddresses: map[common.Address]bool{
		common.HexToAddress("0x00000000000001ad428e4906ae43d8f9852d0dd6"): true,
		common.HexToAddress("0x00000000006c3852cbef3e08e8df289169ede581"): true,
		common.HexToAddress("0xc34349fbEDd527215aAE19B2E4626254ec29A13d"): true,
		// common.HexToAddress("0x0000000000c2d145a2526bd8c716263bfebe1a72"): true, // TransferHelper
		common.HexToAddress("0x00005ea00ac477b1030ce78506496e8c2de24bf5"): true,
	},
	Tag: "|",
}

var Blur = MarketPlace{
	ID:    "blur",
	Name:  "Blur",
	Color: lipgloss.Color("#FF8700"),
	ContractAddresses: map[common.Address]bool{
		common.HexToAddress("0x000000000000Ad05Ccc4F10045630fb830B95127"): true,
		// common.HexToAddress("0x0000000000a39bb272e79075ade125fd351887ac"): true, // blur pool token
		common.HexToAddress("0x39da41747a83aee658334415666f3ef92dd0d541"): true,
	},
	Tag: "|",
}

var X2Y2 = MarketPlace{
	ID:    "x2y2",
	Name:  "X2Y2",
	Color: lipgloss.Color("#acc2fa"),
	ContractAddresses: map[common.Address]bool{
		common.HexToAddress("0x74312363e45dcaba76c59ec49a7aa8a65a67eed3"): true,
	},
	Tag: "|",
}

var LooksRare = MarketPlace{
	ID:    "looksrare",
	Name:  "LooksRare",
	Color: lipgloss.Color("#21E453"),
	ContractAddresses: map[common.Address]bool{
		common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"): true,
	},
	Tag: "|",
}

var SuperRare = MarketPlace{
	ID:    "superrare",
	Name:  "SuperRare",
	Color: lipgloss.Color("#eeeeee"),
	ContractAddresses: map[common.Address]bool{
		common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a"): true,
		common.HexToAddress("0x6D7c44773C52D396F43c2D511B81aa168E9a7a42"): true,
		common.HexToAddress("0x5e62454d6AA7392925ccd3E7cd33f2D7c2f33D97"): true,
		common.HexToAddress("0xDd867a8Eb1720185B3fdAD7F81Caed4E8132Be19"): true,
	},
	Tag: "|",
}

var NFTfi = MarketPlace{
	ID:    "nftfi",
	Name:  "NFTfi",
	Color: lipgloss.Color("#5C4CA0"),
	ContractAddresses: map[common.Address]bool{
		common.HexToAddress("0x8252Df1d8b29057d1Afe3062bf5a64D503152BC8"): true,
		common.HexToAddress("0xaDDE73498902F61BfCB702e94C31c13C534879AC"): true,
		common.HexToAddress("0x5A42d72372858E10Edc03B26bF449F78fF3c0e6F"): true,
		common.HexToAddress("0x0C90C8B4aa8549656851964d5fB787F0e4F54082"): true,
		common.HexToAddress("0x5660E206496808F7b5cDB8C56A696a96AE5E9b23"): true,
		common.HexToAddress("0xe73ECe5988FfF33a012CEA8BB6Fd5B27679fC481"): true,
		common.HexToAddress("0xE52Cec0E90115AbeB3304BaA36bc2655731f7934"): true,
	},
	Tag: "↩︎",
}

var Unknown = MarketPlace{
	ID:    "unknown",
	Name:  "Unknown",
	Color: lipgloss.Color("#777777"),
	ContractAddresses: map[common.Address]bool{
		internal.ZeroAddress: true,
	},
	Tag: "¦",
}

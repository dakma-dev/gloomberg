package internal

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

const (
	PubSubChannelSales    = "sales"
	PubSubChannelListings = "listings"

	PubSubSeaWatcher         = "seawatcher"
	PubSubSeaWatcherMgmt     = PubSubSeaWatcher + "/mgmt"
	PubSubSeaWatcherListings = PubSubSeaWatcher + "/" + PubSubChannelListings

	BlockTime = 12 * time.Second

	NoENSName = "NO-ENS-NAME"

	PathDegendata = "degendata/"
)

var (
	GloombergVersion = "dev"
	RunningSince     = time.Now()

	WETHContractAddress           = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	BlurPoolTokenContractAddress  = common.HexToAddress("0x0000000000a39bb272e79075ade125fd351887ac")
	BlurBlendContractAddress      = common.HexToAddress("0x29469395eAf6f95920E59F858042f0e28D98a20B")
	ENSContractAddress            = common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")
	ENSNameWrapperContractAddress = common.HexToAddress("0xd4416b13d2b3a9abae7acd5d6c2bbdbe25686401")

	GrifterContractAddress = common.HexToAddress("0xc143bbfcdbdbed6d454803804752a064a622c1f3")

	// manifold.
	ManifoldCreatorCoreERC721  = common.HexToAddress("0x5133522ea5A0494EcB83F26311A095DDD7a9D4b6")
	ManifoldCreatorCoreERC1155 = common.HexToAddress("0xE9FF7CA11280553Af56d04Ecb8Be6B8c4468DCB2")
	ManifoldLazyClaimERC1155   = common.HexToAddress("0xDb8d79C775452a3929b86ac5DEaB3e9d38e1c006")
	ManifoldBurnRedeemERC1155  = common.HexToAddress("0xde659726CfD166aCa4867994d396EFeF386EAD68")

	// loan stuff.
	NFTfiContractAddress           = common.HexToAddress("0x5660E206496808F7b5cDB8C56A696a96AE5E9b23")
	NFTLoanTicketV2ContractAddress = common.HexToAddress("0x0E258c84Df0f8728ae4A6426EA5FD163Eb6b9D1B")
	BorrowerNoteTicket             = common.HexToAddress("0xbD85BF4C970b91984e6A2b8Ba9C577A58A8C20f9")

	// uniswapv2.

	UniswapV2FactoryContractAddress  = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	UniswapV2Router01ContractAddress = common.HexToAddress("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a")

	// uniswapv3.

	UniswapV3FactoryContractAddress       = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	UniswapUniversalRouterContractAddress = common.HexToAddress("0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B")
	UniswapV3QuoterContractAddress        = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
	UniswapV3QuoterV2ContractAddress      = common.HexToAddress("0x61fFE014bA17989E743c5F6cB21bF9697530B21e")

	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ZeroHash    = common.Hash{}

	// BaseLogger is the logger used to print to the terminal without reporting caller or timestamp.
	BaseLogger = log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
	})
)

// CountryFlags is a alphabetically sorted map from CountryCodes to flag emojis.
var CountryFlags = map[string]string{
	"00": "🌍",
	"AD": "🇦🇩",
	"AE": "🇦🇪",
	"AF": "🇦🇫",
	"AG": "🇦🇬",
	"AI": "🇦🇮",
	"AL": "🇦🇱",
	"AM": "🇦🇲",
	"AO": "🇦🇴",
	"AQ": "🇦🇶",
	"AR": "🇦🇷",
	"AS": "🇦🇸",
	"AT": "🇦🇹",
	"AU": "🇦🇺",
	"AW": "🇦🇼",
	"AX": "🇦🇽",
	"AZ": "🇦🇿",
	"BA": "🇧🇦",
	"BB": "🇧🇧",
	"BD": "🇧🇩",
	"BE": "🇧🇪",
	"BF": "🇧🇫",
	"BG": "🇧🇬",
	"BH": "🇧🇭",
	"BI": "🇧🇮",
	"BJ": "🇧🇯",
	"BL": "🇧🇱",
	"BM": "🇧🇲",
	"BN": "🇧🇳",
	"BO": "🇧🇴",
	"BQ": "🇧🇶",
	"BR": "🇧🇷",
	"BS": "🇧🇸",
	"BT": "🇧🇹",
	"BV": "🇧🇻",
	"BW": "🇧🇼",
	"BY": "🇧🇾",
	"BZ": "🇧🇿",
	"CA": "🇨🇦",
	"CC": "🇨🇨",
	"CD": "🇨🇩",
	"CF": "🇨🇫",
	"CG": "🇨🇬",
	"CH": "🇨🇭",
	"CI": "🇨🇮",
	"CK": "🇨🇰",
	"CL": "🇨🇱",
	"CM": "🇨🇲",
	"CN": "🇨🇳",
	"CO": "🇨🇴",
	"CR": "🇨🇷",
	"CU": "🇨🇺",
	"CV": "🇨🇻",
	"CW": "🇨🇼",
	"CX": "🇨🇽",
	"CY": "🇨🇾",
	"CZ": "🇨🇿",
	"DE": "🇩🇪",
	"DJ": "🇩🇯",
	"DK": "🇩🇰",
	"DM": "🇩🇲",
	"DO": "🇩🇴",
	"DZ": "🇩🇿",
	"EC": "🇪🇨",
	"EE": "🇪🇪",
	"EG": "🇪🇬",
	"EH": "🇪🇭",
	"ER": "🇪🇷",
	"ES": "🇪🇸",
	"ET": "🇪🇹",
	"FI": "🇫🇮",
	"FJ": "🇫🇯",
	"FK": "🇫🇰",
	"FM": "🇫🇲",
	"FO": "🇫🇴",
	"FR": "🇫🇷",
	"GA": "🇬🇦",
	"GB": "🇬🇧",
	"GD": "🇬🇩",
	"GE": "🇬🇪",
	"GF": "🇬🇫",
	"GG": "🇬🇬",
	"GH": "🇬🇭",
	"GI": "🇬🇮",
	"GL": "🇬🇱",
	"GM": "🇬🇲",
	"GN": "🇬🇳",
	"GP": "🇬🇵",
	"GQ": "🇬🇶",
	"GR": "🇬🇷",
	"GS": "🇬🇸",
	"GT": "🇬🇹",
	"GU": "🇬🇺",
	"GW": "🇬🇼",
	"GY": "🇬🇾",
	"HK": "🇭🇰",
	"HM": "🇭🇲",
	"HN": "🇭🇳",
	"HR": "🇭🇷",
	"HT": "🇭🇹",
	"HU": "🇭🇺",
	"ID": "🇮🇩",
	"IE": "🇮🇪",
	"IL": "🇮🇱",
	"IM": "🇮🇲",
	"IN": "🇮🇳",
	"IO": "🇮🇴",
	"IQ": "🇮🇶",
	"IR": "🇮🇷",
	"IS": "🇮🇸",
	"IT": "🇮🇹",
	"JE": "🇯🇪",
	"JM": "🇯🇲",
	"JO": "🇯🇴",
	"JP": "🇯🇵",
	"KE": "🇰🇪",
	"KG": "🇰🇬",
	"KH": "🇰🇭",
	"KI": "🇰🇮",
	"KM": "🇰🇲",
	"KN": "🇰🇳",
	"KP": "🇰🇵",
	"KR": "🇰🇷",
	"KW": "🇰🇼",
	"KY": "🇰🇾",
	"KZ": "🇰🇿",
	"LA": "🇱🇦",
	"LB": "🇱🇧",
	"LC": "🇱🇨",
	"LI": "🇱🇮",
	"LK": "🇱🇰",
	"LR": "🇱🇷",
	"LS": "🇱🇸",
	"LT": "🇱🇹",
	"LU": "🇱🇺",
	"LV": "🇱🇻",
	"LY": "🇱🇾",
	"MA": "🇲🇦",
	"MC": "🇲🇨",
	"MD": "🇲🇩",
	"ME": "🇲🇪",
	"MF": "🇲🇫",
	"MG": "🇲🇬",
	"MH": "🇲🇭",
	"MK": "🇲🇰",
	"ML": "🇲🇱",
	"MM": "🇲🇲",
	"MN": "🇲🇳",
	"MO": "🇲🇴",
	"MP": "🇲🇵",
	"MQ": "🇲🇶",
	"MR": "🇲🇷",
	"MS": "🇲🇸",
	"MT": "🇲🇹",
	"MU": "🇲🇺",
	"MV": "🇲🇻",
	"MW": "🇲🇼",
	"MX": "🇲🇽",
	"MY": "🇲🇾",
	"MZ": "🇲🇿",
	"NA": "🇳🇦",
	"NC": "🇳🇨",
	"NE": "🇳🇪",
	"NF": "🇳🇫",
	"NG": "🇳🇬",
	"NI": "🇳🇮",
	"NL": "🇳🇱",
	"NO": "🇳🇴",
	"NP": "🇳🇵",
	"NR": "🇳🇷",
	"NU": "🇳🇺",
	"NZ": "🇳🇿",
	"OM": "🇴🇲",
	"PA": "🇵🇦",
	"PE": "🇵🇪",
	"PF": "🇵🇫",
	"PG": "🇵🇬",
	"PH": "🇵🇭",
	"PK": "🇵🇰",
	"PL": "🇵🇱",
	"PM": "🇵🇲",
	"PN": "🇵🇳",
	"PR": "🇵🇷",
	"PS": "🇵🇸",
	"PT": "🇵🇹",
	"PW": "🇵🇼",
	"PY": "🇵🇾",
	"QA": "🇶🇦",
	"RE": "🇷🇪",
	"RO": "🇷🇴",
	"RS": "🇷🇸",
	"RU": "🇷🇺",
	"RW": "🇷🇼",
	"SA": "🇸🇦",
	"SB": "🇸🇧",
	"SC": "🇸🇨",
	"SD": "🇸🇩",
	"SE": "🇸🇪",
	"SG": "🇸🇬",
	"SH": "🇸🇭",
	"SI": "🇸🇮",
	"SJ": "🇸🇯",
	"SK": "🇸🇰",
	"SL": "🇸🇱",
	"SM": "🇸🇲",
	"SN": "🇸🇳",
	"SO": "🇸🇴",
	"SR": "🇸🇷",
	"SS": "🇸🇸",
	"ST": "🇸🇹",
	"SV": "🇸🇻",
	"SX": "🇸🇽",
	"SY": "🇸🇾",
	"SZ": "🇸🇿",
	"TC": "🇹🇨",
	"TD": "🇹🇩",
	"TF": "🇹🇫",
	"TG": "🇹🇬",
	"TH": "🇹🇭",
	"TJ": "🇹🇯",
	"TK": "🇹🇰",
	"TL": "🇹🇱",
	"TM": "🇹🇲",
	"TN": "🇹🇳",
	"TO": "🇹🇴",
	"TR": "🇹🇷",
	"TT": "🇹🇹",
	"TV": "🇹🇻",
	"TW": "🇹🇼",
	"TZ": "🇹🇿",
	"UA": "🇺🇦",
	"UG": "🇺🇬",
	"UM": "🇺🇲",
	"US": "🇺🇸",
	"UY": "🇺🇾",
	"UZ": "🇺🇿",
	"VA": "🇻🇦",
	"VC": "🇻🇨",
	"VE": "🇻🇪",
	"VG": "🇻🇬",
	"VI": "🇻🇮",
	"VN": "🇻🇳",
	"VU": "🇻🇺",
	"WF": "🇼🇫",
	"WS": "🇼🇸",
	"YE": "🇾🇪",
	"YT": "🇾🇹",
	"ZA": "🇿🇦",
	"ZM": "🇿🇲",
	"ZW": "🇿🇼",
}

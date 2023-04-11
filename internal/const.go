package internal

import "github.com/ethereum/go-ethereum/common"

var (
	WETHContractAddress          = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	BlurPoolTokenContractAddress = common.HexToAddress("0x0000000000a39bb272e79075ade125fd351887ac")

	ENSContractAddress = common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")

	PubSubChannelSales = "sales"

	TopicSeaWatcher     = "seawatcher"
	TopicSeaWatcherMgmt = TopicSeaWatcher + "/mgmt"
)

package flotscmd

import (
	"encoding/json"
	"fmt"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/r3labs/sse/v2"
	"github.com/spf13/cobra"
)

type mevsLog struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type mevsEvent struct {
	Hash      string      `json:"hash"`
	MethodSig string      `json:"methodSig"`
	Logs      []mevsLog   `json:"logs"`
	Txs       interface{} `json:"txs"`
}

var (
	mevShareEventStream = "https://mev-share.flashbots.net"
	syncTopic           = "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"

	eventSignatures = map[common.Hash]string{
		common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"): "Swap(address,uint256,uint256,uint256,uint256,address)",
		common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"): "Swap(address,address,int256,int256,uint160,uint128,int24)",
	}
)

// mevShareCmd represents the mevShare command
var mevShareCmd = &cobra.Command{
	Use:     "mev-share",
	Aliases: []string{"ms", "mevShare", "mevshare"},
	Short:   "watch the mev-share stream from ",
	Run: func(_ *cobra.Command, _ []string) {
		mevShareStream()
	},
}

func mevShareStream() {
	events := make(chan *sse.Event)

	client := sse.NewClient(mevShareEventStream)
	if err := client.SubscribeChan("mev-share", events); err != nil {
		log.Fatal(err)
	}

	for event := range events {
		fmt.Println()
		log.Debugf("data: %+v\n", string(event.Data))

		var mevsEvent mevsEvent

		err := json.Unmarshal(event.Data, &mevsEvent)
		if err != nil {
			log.Errorf("error: %s\n", err.Error())
		}

		for _, msLog := range mevsEvent.Logs {
			if msLog.Topics[0] != syncTopic {
				continue
			}

			switch {
			case len(msLog.Topics) > 1 && msLog.Topics[1] != common.Hash{}.String():
				log.Printf("logTopic[1]: %s\n", msLog.Topics[1])

				fallthrough
			case len(msLog.Topics) > 2 && msLog.Topics[2] != common.Hash{}.String():
				log.Printf("logTopic[2]: %s\n", msLog.Topics[2])

				fallthrough
			case len(msLog.Topics) > 3 && msLog.Topics[3] != common.Hash{}.String():
				log.Printf("logTopic[3]: %s\n", msLog.Topics)
			}

			if methodSig := eventSignatures[common.HexToHash(msLog.Topics[0])]; methodSig != "" {
				mevsEvent.MethodSig = methodSig
			} else {
				if signature, err := external.GetEventSignature(common.HexToHash(msLog.Topics[0])); err == nil {
					mevsEvent.MethodSig = signature.TextSignature
					eventSignatures[common.HexToHash(msLog.Topics[0])] = signature.TextSignature
				} else {
					mevsEvent.MethodSig = "unknown"
				}
			}
		}

		fmt.Printf("mevsEvent: %+v\n", mevsEvent)
	}
}

func init() { FlotsCmd.AddCommand(mevShareCmd) }

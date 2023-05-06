/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/r3labs/sse/v2"
	"github.com/spf13/cobra"
)

var eventSignatures = map[common.Hash]string{
	common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"): "Swap(address,uint256,uint256,uint256,uint256,address)",
	common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"): "Swap(address,address,int256,int256,uint160,uint128,int24)",
}

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

// shearcherCmd represents the shearcher command.
var shearcherCmd = &cobra.Command{
	Use:   "shearcher",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shearcher called")

		events := make(chan *sse.Event)

		client := sse.NewClient("https://mev-share.flashbots.net")
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
	},
}

func init() {
	rootCmd.AddCommand(shearcherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shearcherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shearcherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

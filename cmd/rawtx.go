/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/benleb/gloomberg/internal/config"
	"github.com/benleb/gloomberg/internal/seaport"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rawtxCmd represents the rawtx command
var rawtxCmd = &cobra.Command{
	Use:   "rawtx",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rawtx called")

		to := common.HexToAddress("0xf132f2c8f1eede27070e0850775436a0e6e7268a")

		ethNodes := config.GetNodesFromConfig()
		if ethNodes.ConnectAllNodes() == nil {
			fmt.Println("no nodes")
			return
		}

		privateKey := viper.GetString("buy.privateKey")
		if privateKey == "" {
			gbl.Log.Error("❌ private key is empty")
			return
		}

		rawData := "0x50733564000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000004cc0000000000000000000000000000000000000000000000000000000000000041999a8a1bb08e0cc824b7600b02611ffc4d194baae2d2bf43a42a77dbdb65c46e520455b7cd038209492fc1826f9199cf7938def6278294b2f3b67526248a02301c00000000000000000000000000000000000000000000000000000000000000"
		data, err := hex.DecodeString(rawData)
		if err != nil {
			panic(err)
		}

		seaport.SendRawTx(ethNodes, to, 0.003, data, privateKey)
	},
}

func init() {
	rootCmd.AddCommand(rawtxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rawtxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rawtxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

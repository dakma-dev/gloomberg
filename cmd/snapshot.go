package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

var (
	apikey string
)

type getOwnersForCollection struct {
	OwnerAddresses []string `json:"ownerAddresses"`
}

// generateCmd represents the generate command
var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Generate snapshot of wallets",
	Long:  `Generate snapshot of wallets using third API provider.`,
	Run: func(cmd *cobra.Command, args []string) {

		// check if argument is given
		if len(args) < 1 {
			fmt.Println("❌ missing argument: contract address")
			return
		}

		// get first argument
		contract := args[0]

		viper.Set("log.logFile", "/tmp/gloomberg-generate.log")
		viper.Set("log.verbose", true)

		fmt.Println("--")
		fmt.Println(fmt.Sprintf("🔍 getOwnersForCollection from alchemy · contract: %s", contract))

		// check if apikey is set
		if apikey == "" {
			fmt.Println(fmt.Sprintf("❌ missing argument: apikey"))
			return
		}

		// https://eth-mainnet.g.alchemy.com/nft/v2/{apiKey}/getOwnersForCollection
		//contract := "0x769272677fab02575e84945f03eca517acc544cc"
		url := "https://eth-mainnet.g.alchemy.com/nft/v2/" + apikey + "/getOwnersForCollection?contractAddress=" + contract
		response, err := utils.HTTP.GetWithTLS12(context.TODO(), url)
		if err != nil {
			if os.IsTimeout(err) {
				fmt.Println(fmt.Sprintf("⌛️ getContractMetadata from alchemy · timeout while fetching: %+v", err.Error()))
			} else {
				fmt.Println(fmt.Sprintf("❌ getContractMetadata from alchemy · error: %+v", err.Error()))
			}
			return
		}

		//gbl.Log.Debugf("getContractMetadata status: %s", response.Status)
		defer response.Body.Close()
		// read the response body
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(fmt.Sprintf("❌ getContractMetadata from alchemy · response read error: %+v", err.Error()))
			return
		}
		// decode

		var owners *getOwnersForCollection

		// decode the data
		dec := json.NewDecoder(bytes.NewReader(responseBody))

		if err := dec.Decode(&owners); err != nil {
			fmt.Println(fmt.Sprintf("❌  decode error: %s", err.Error()))
			//gbl.Log.Warnf("multiAccountBalance decode error: %s", err.Error())
		}

		fmt.Println(fmt.Sprintf("👛 %d owners found", len(owners.OwnerAddresses)))
		//fmt.Println(owners)
		// save struct as json
		jsonString, err := json.Marshal(owners)
		if err != nil {
			fmt.Println(fmt.Sprintf("❌  decode error: %s", err.Error()))
		}
		// write jsonString to file
		fileName := "./wallets/" + contract + ".json"
		fmt.Println(fmt.Sprintf("📝 writing file: %s", fileName))
		err = os.WriteFile(fileName, jsonString, 0644)
		if err != nil {
			fmt.Println(fmt.Sprintf("❌  writing file error: %s", err.Error()))
		}

	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)

	snapshotCmd.Flags().StringVar(&apikey, "alchemy-key", "", "alchemy api key")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

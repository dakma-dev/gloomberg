package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apikey          string
	snapshotAddress string
)

// generateCmd represents the generate command.
var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Generate snapshot of wallets",
	Long:  `Generate snapshot of wallets using third API provider.`,
	Run: func(_ *cobra.Command, args []string) {
		// check if argument is given
		if len(args) < 1 && snapshotAddress == "" {
			fmt.Println("❌ missing argument: contract address")

			return
		}

		var contract string

		if len(args) < 1 {
			contract = snapshotAddress
		} else {
			// get first argument
			contract = args[0]
		}

		// check if argument is given
		if snapshotAddress != "" {
			contract = snapshotAddress
		}

		viper.Set("log.logFile", "/tmp/gloomberg-generate.log")
		viper.Set("log.verbose", true)

		fmt.Println("--")

		fmt.Printf("🔍 getOwnersForCollection from alchemy · contract: %s\n", contract)

		// check if apikey is set
		if apikey == "" {
			fmt.Printf("❌ missing argument: apikey\n")

			return
		}

		// https://eth-mainnet.g.alchemy.com/nft/v2/{apiKey}/getOwnersForCollection
		url := "https://eth-mainnet.g.alchemy.com/nft/v2/" + apikey + "/getOwnersForCollection?contractAddress=" + contract
		response, err := utils.HTTP.GetWithTLS12(context.TODO(), url)
		if err != nil {
			if os.IsTimeout(err) {
				fmt.Printf("⌛️ getContractMetadata from alchemy · timeout while fetching: %+v\n", err.Error())
			} else {
				fmt.Printf("❌ getContractMetadata from alchemy · error: %+v\n", err.Error())
			}

			return
		}

		defer response.Body.Close()

		// read the response body
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("❌ getContractMetadata from alchemy · response read error: %+v\n", err.Error())

			return
		}

		// decode the data
		var owners *ticker.GetOwnersForCollectionResponse
		if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&owners); err != nil {
			fmt.Printf("❌  decode error: %s\n", err.Error())

			return
		}

		fmt.Printf("👛 %d owners found\n", len(owners.OwnerAddresses))

		// save struct as json
		jsonString, err := json.Marshal(owners)
		if err != nil {
			fmt.Printf("❌  decode error: %s\n", err.Error())
		}
		// write jsonString to file
		fileName := "./wallets/" + contract + ".json"

		fmt.Printf("📝 writing file: %s\n", fileName)

		err = os.WriteFile(fileName, jsonString, 0o644) //nolint:gosec
		if err != nil {
			fmt.Printf("❌  writing file error: %s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)

	// intentionally called the flag "snapshot.alchemy.key" to make the difference between the flag and the viper key clear
	// (usually the flag and the viper key have the same name to avoid confusion)
	snapshotCmd.Flags().StringVar(&apikey, "alchemy.key", "", "alchemy api key")
	// bind the cobra/pflags flag "alchemy.key" to the viper key "alchemy.apiKey
	_ = viper.BindPFlag("alchemy.apiKey", snapshotCmd.Flags().Lookup("alchemy.key"))

	// we do not need to bind the address flag to a viper key because we do not need it anywhere else
	snapshotCmd.Flags().StringVar(&snapshotAddress, "address", "", "contract address to snapshot")
}

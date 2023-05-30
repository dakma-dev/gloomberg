package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type GetFloorPriceAlchemyResponse struct {
	Opensea   FloorPriceAlchemyData `json:"openSea"`
	Looksrare FloorPriceAlchemyData `json:"looksRare"`
}

type FloorPriceAlchemyData struct {
	FloorPrice    float64 `json:"floorPrice"`
	PriceCurrency string  `json:"priceCurrency"`
	CollectionURL string  `json:"collectionUrl"`
	RetrievedAt   string  `json:"retrievedAt"`
	Error         string  `json:"error"`
}

func GetFloorPriceFromAlchemy(contract string) *GetFloorPriceAlchemyResponse {
	if contract == "" {
		fmt.Printf("❌ getContractMetadata from alchemy · error: contract address is empty\n")

		return nil
	}

	// https://eth-mainnet.g.alchemy.com/nft/v3/{apiKey}/getFloorPrice
	apikey := viper.GetString("api_keys.alchemy")
	url := "https://eth-mainnet.g.alchemy.com/nft/v3/" + apikey + "/getFloorPrice?contractAddress=" + contract
	response, err := utils.HTTP.GetWithTLS12(context.TODO(), url)
	if err != nil {
		if os.IsTimeout(err) {
			fmt.Printf("⌛️ getContractMetadata from alchemy · timeout while fetching: %+v\n", err.Error())
		} else {
			fmt.Printf("❌ getContractMetadata from alchemy · error: %+v\n", err.Error())
		}

		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("❌ getContractMetadata from alchemy · error: %+v\n", response.Status)

		return nil
	}

	// read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("❌ getContractMetadata from alchemy · response read error: %+v\n", err.Error())

		return nil
	}

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		log.Warnf("getContractMetadata invalid json: %s", err)

		return nil
	}

	// fmt.Println(string(responseBody))
	var decoded *GetFloorPriceAlchemyResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decoded); err != nil {
		fmt.Printf("❌  decode error: %s\n", err.Error())

		return nil
	}

	return decoded
}

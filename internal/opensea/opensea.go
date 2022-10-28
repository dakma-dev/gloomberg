package opensea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/utils"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func openSeaHeader() http.Header {
	header := http.Header{}
	if viper.IsSet("api_keys.opensea") {
		header.Add("X-API-KEY", viper.GetString("api_keys.opensea"))
	}

	return header
}

// func createOpenSeaHeaders() (*http.Header, error) {
// 	headers := &http.Header{}

// 	headers.Add("Accept", "application/json")
// 	headers.Add("Cache-Control", "no-cache")

// 	if viper.IsSet("api_keys.opensea") {
// 		headers.Add("X-API-KEY", viper.GetString("api_keys.opensea"))
// 	}

// 	return headers, nil
// }

// func createHTTPClient() (*http.Client, error) {
// 	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13}

// 	transport := &http.Transport{
// 		MaxIdleConnsPerHost:   25,
// 		TLSClientConfig:       tlsConfig,
// 		IdleConnTimeout:       17 * time.Second,
// 		ResponseHeaderTimeout: 7 * time.Second,
// 		TLSHandshakeTimeout:   5 * time.Second,
// 	}

// 	// explicitly use http2
// 	_ = http2.ConfigureTransport(transport)

// 	client := &http.Client{
// 		Timeout:   9 * time.Second,
// 		Transport: transport,
// 	}

// 	return client, nil
// }

// func createGetRequest(url string) (*http.Request, error) {
// 	req, _ := http.NewRequest("GET", url, nil)

// 	if httpHeader, _ := createOpenSeaHeaders(); httpHeader != nil {
// 		req.Header = *httpHeader
// 	}

// 	return req, nil
// }

// GetWalletCollections returns the collections a wallet owns at least one item of.
func GetWalletCollections(wallets *wallet.Wallets, userCollections *collections.CollectionDB, nodes *nodes.Nodes) []*collections.GbCollection {
	gbCollections := make([]*collections.GbCollection, 0)

	for _, w := range *wallets {
		gbCollections = append(gbCollections, GetCollectionsFor(w.Address, userCollections, nodes, 1)...)
	}

	return gbCollections
}

// GetCollectionsFor returns the collections a wallet owns at least one item of.
func GetCollectionsFor(walletAddress common.Address, userCollections *collections.CollectionDB, nodes *nodes.Nodes, try int) []*collections.GbCollection {
	receivedCollections := make([]*collections.GbCollection, 0)

	// create the http client & request
	// client, _ := createHTTPClient()

	// request, _ := createGetRequest(url)

	// response, err := client.Do(request)

	url := fmt.Sprintf("https://api.opensea.io/api/v1/collections?asset_owner=%s&offset=0&limit=300", walletAddress)

	response, err := utils.HTTP.GetWithHeader(url, openSeaHeader())
	if os.IsTimeout(err) {
		backoffSeconds := try * 2
		sleepTime := time.Duration(backoffSeconds) * time.Second
		time.Sleep(sleepTime)

		gbl.Log.Warnf("⌛️ timeout while fetching wallet collections for %s (try %d, sleep %ds)", walletAddress.Hex(), try, backoffSeconds)

		if try <= viper.GetInt("ens.resolve_max_retries") {
			GetCollectionsFor(walletAddress, userCollections, nodes, try+1)
		} else {
			gbl.Log.Warnf("⌛️ timeout while fetching wallet collections for %s, giving up after %d retries...", walletAddress.Hex(), try-1)
		}

		return receivedCollections
	} else if err != nil {
		return receivedCollections
	}
	defer response.Body.Close()

	// create a variable of the same type as our model
	var collectionResponse []*models.AssetCollection

	responseBody, _ := io.ReadAll(response.Body)

	// decode the data
	if !json.Valid(responseBody) {
		return receivedCollections
	}

	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&collectionResponse); err != nil {
		gbl.Log.Errorf("⌛️ error while decoding wallet collections for %s: %s", walletAddress.Hex(), err)
	}

	for _, collection := range collectionResponse {
		for _, contract := range collection.PrimaryAssetContracts {
			contractAddress := common.HexToAddress(contract.Address)

			if userCollections.Collections[contractAddress] != nil || contractAddress == external.ENSContract || collection.Name == "MegaCryptoPolis" {
				continue
			}

			if collection.Stats.AveragePrice <= 0.001 {
				continue
			}

			userCollection := collections.NewCollection(contractAddress, collection.Name, nodes, models.FromWallet)
			userCollection.OpenseaSlug = collection.Slug

			receivedCollections = append(receivedCollections, userCollection)
		}
	}

	return receivedCollections
}

func GetCollectionSlug(collectionAddress common.Address) string {
	assetContract := GetAssetContract(collectionAddress)

	if assetContract == nil {
		return ""
	}

	return assetContract.Collection.Slug
}

func GetAssetContract(contractAddress common.Address) *models.AssetContract {
	url := fmt.Sprintf("https://api.opensea.io/api/v1/asset_contract/%s", contractAddress.String())

	response, err := utils.HTTP.GetWithHeader(url, openSeaHeader())
	if err != nil {
		// if os.IsTimeout(err) {
		// 	// dalog.Warn("TIMEOUT while fetching listings, trying again next round... ", collectionSlug)
		// } else {
		// 	// dalog.Warn("ooopsss an error occurred while fetching asset events, please try again:", err)
		// }
		return nil
	}

	defer response.Body.Close()

	// create a variable of the same type as our model
	var assetContract *models.AssetContract

	responseBody, _ := io.ReadAll(response.Body)

	// fmt.Printf("response: %d | body: %s\n", response.StatusCode, responseBody)

	// decode the data
	if !json.Valid(responseBody) {
		return nil
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&assetContract); err != nil {
		return nil
	}

	if assetContract == nil {
		gbl.Log.Errorf("⌛️ error while decoding asset contract for %s: %s", contractAddress.Hex(), err)

		return nil
	}

	return assetContract
}

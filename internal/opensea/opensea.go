package opensea

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

func createOpenSeaHeaders() (*http.Header, error) {
	headers := &http.Header{}

	headers.Add("Accept", "application/json")
	headers.Add("Cache-Control", "no-cache")

	if viper.IsSet("api_keys.opensea") {
		headers.Add("X-API-KEY", viper.GetString("api_keys.opensea"))
	}

	return headers, nil
}

func createHTTPClient() (*http.Client, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13}

	transport := &http.Transport{
		MaxIdleConnsPerHost:   25,
		TLSClientConfig:       tlsConfig,
		IdleConnTimeout:       20 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
	}

	// explicitly use http2
	_ = http2.ConfigureTransport(transport)

	client := &http.Client{
		Timeout:   9 * time.Second,
		Transport: transport,
	}

	return client, nil
}

func createGetRequest(url string) (*http.Request, error) {
	req, _ := http.NewRequest("GET", url, nil)

	if httpHeader, _ := createOpenSeaHeaders(); httpHeader != nil {
		req.Header = *httpHeader
	}

	return req, nil
}

// GetWalletCollections returns the collections a wallet owns at least one item of.
func GetWalletCollections(wallets map[common.Address]*models.Wallet, userCollections *collections.Collections, nodes *gbnode.NodeCollection) []*collections.GbCollection {
	collections := make([]*collections.GbCollection, 0)

	for _, wallet := range wallets {
		collections = append(collections, GetCollectionsFor(wallet.Address, userCollections, nodes, 1)...)
	}

	return collections
}

// GetCollectionsFor returns the collections a wallet owns at least one item of.
func GetCollectionsFor(walletAddress common.Address, userCollections *collections.Collections, nodes *gbnode.NodeCollection, try int) []*collections.GbCollection {
	receivedCollections := make([]*collections.GbCollection, 0)

	// create the http client & request
	client, _ := createHTTPClient()

	url := fmt.Sprintf("https://api.opensea.io/api/v1/collections?asset_owner=%s&offset=0&limit=300", walletAddress)
	request, _ := createGetRequest(url)

	response, err := client.Do(request)
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

	responseBody, _ := ioutil.ReadAll(response.Body)

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

			if userCollections.UserCollections[contractAddress] != nil || contractAddress == common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85") || collection.Name == "MegaCryptoPolis" {
				continue
			}

			if collection.Stats.AveragePrice <= 0.01 {
				continue
			}

			userCollection := collections.NewCollection(contractAddress, collection.Name, nodes, collections.Wallet)
			userCollection.Metadata.OpenseaSlug = collection.Slug

			receivedCollections = append(receivedCollections, userCollection)
		}
	}

	return receivedCollections
}

// GetAssetEvents returns the events for a collection.
func GetAssetEvents(_ time.Time, userCollections *collections.Collections, contractAddress common.Address, newListings chan<- []models.AssetEvent) {
	collections := *userCollections
	collection := *collections.UserCollections[contractAddress]

	// create the http client & request
	client, _ := createHTTPClient()

	numberOfEventsToFetch := 30
	url := fmt.Sprintf("https://api.opensea.io/api/v1/events?asset_contract_address=%s&only_opensea=false&limit=%d&event_type=created", collection.ContractAddress, numberOfEventsToFetch)

	request, _ := createGetRequest(url)

	response, err := client.Do(request)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching listings for %s", collection.Name)
		} else {
			gbl.Log.Error("❌ opensea request error:", err)
		}

		return
	}

	defer response.Body.Close()

	// create a variable of the same type as our model
	var eventsResponse *models.AssetEventsResponse

	responseBody, _ := ioutil.ReadAll(response.Body)

	// decode the data
	if !json.Valid(responseBody) {
		fmt.Println(response.Status, "| GetAssetEvents expected a json but received something else, trying again next round...")

		return
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&eventsResponse); err != nil {
		fmt.Println(response.Status, "| GetAssetEvents ooops! an error occurred while decoding the events, please try again! error:", err)

		return
	}

	if len(eventsResponse.AssetEvents) > 0 {
		newListings <- eventsResponse.AssetEvents
	}
}

func GetCollectionSlug(collectionAddress common.Address) string {
	assetContract := GetAssetContract(collectionAddress)

	if assetContract == nil {
		return ""
	}

	return assetContract.Collection.Slug
}

func GetAssetContract(contractAddress common.Address) *models.AssetContract {
	// create the http client & request
	client, _ := createHTTPClient()

	url := fmt.Sprintf("https://api.opensea.io/api/v1/asset_contract/%s", contractAddress.String())

	request, _ := createGetRequest(url)

	response, err := client.Do(request)
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

	responseBody, _ := ioutil.ReadAll(response.Body)

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

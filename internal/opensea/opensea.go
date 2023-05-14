package opensea

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/collectionsource"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/utils"
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

// GetWalletCollections returns the collections a wallet owns at least one item of.
// func GetWalletCollections(wallets *wallet.Wallets, userCollections *collections.CollectionDB, nodes *nodes.Nodes) []*collections.Collection {.
func GetWalletCollections(gb *gloomberg.Gloomberg) []*collections.Collection {
	gbCollections := make([]*collections.Collection, 0)

	for _, w := range *gb.OwnWallets {
		gbCollections = append(gbCollections, GetCollectionsFor(w.Address, gb.CollectionDB, gb.ProviderPool, 1, gb.Rueidi)...)
	}

	return gbCollections
}

// GetCollectionsFor returns the collections a wallet owns at least one item of.
func GetCollectionsFor(walletAddress common.Address, userCollections *collections.CollectionDB, providerPool *provider.Pool, try int, rueidica *rueidica.Rueidica) []*collections.Collection {
	receivedCollections := make([]*collections.Collection, 0)

	url := fmt.Sprintf("https://api.opensea.io/api/v1/collections?asset_owner=%s&offset=0&limit=300", walletAddress)

	response, err := utils.HTTP.GetWithHeader(context.Background(), url, openSeaHeader())
	if os.IsTimeout(err) {
		backoffSeconds := try * 2
		sleepTime := time.Duration(backoffSeconds) * time.Second
		time.Sleep(sleepTime)

		gbl.Log.Warnf("⌛️ timeout while fetching wallet collections for %s (try %d, sleep %ds)", walletAddress.Hex(), try, backoffSeconds)

		if try <= viper.GetInt("ens.resolve_max_retries") {
			GetCollectionsFor(walletAddress, userCollections, providerPool, try+1, rueidica)
		} else {
			gbl.Log.Warnf("⌛️ timeout while fetching wallet collections for %s, giving up after %d retries...", walletAddress.Hex(), try-1)
		}

		return receivedCollections
	} else if err != nil {
		return receivedCollections
	}
	defer response.Body.Close()

	// create a variable of the same type as our model
	var collectionResponse []*osmodels.AssetCollection

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

			// if userCollections.Collections[contractAddress] != nil || contractAddress == external.ENSContract || collection.Name == "MegaCryptoPolis" {
			if userCollections.Collections[contractAddress] != nil {
				continue
			}

			if collection.Stats.AveragePrice <= 0.001 {
				continue
			}

			userCollection := collections.NewCollection(contractAddress, collection.Name, providerPool, collectionsource.FromWallet, nil)
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

func GetAssetContract(contractAddress common.Address) *osmodels.AssetContract {
	url := fmt.Sprintf("https://api.opensea.io/api/v1/asset_contract/%s", contractAddress.String())

	response, err := utils.HTTP.GetWithHeader(context.Background(), url, openSeaHeader())
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
	var assetContract *osmodels.AssetContract

	responseBody, _ := io.ReadAll(response.Body)

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

func GetListings(contractAddress common.Address, tokenID int64) []osmodels.SeaportOrder {
	url := fmt.Sprintf("https://api.opensea.io/v2/orders/ethereum/seaport/listings?asset_contract_address=%s&token_ids=%d&order_by=created_date&order_direction=desc", contractAddress.String(), tokenID)

	response, err := utils.HTTP.GetWithHeader(context.Background(), url, openSeaHeader())
	if err != nil {
		gbl.Log.Errorf("❌ error while fetching listings for %s/%d: %s", contractAddress.Hex(), tokenID, err)
		// if os.IsTimeout(err) {
		// 	// dalog.Warn("TIMEOUT while fetching listings, trying again next round... ", collectionSlug)
		// } else {
		// 	// dalog.Warn("ooopsss an error occurred while fetching asset events, please try again:", err)
		// }
		return nil
	}

	defer response.Body.Close()

	// create a variable of the same type as our model
	var listingsResponse *osmodels.OpenSeaListingsResponse

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		gbl.Log.Errorf("❌ error reading listings response body for %s/%d: %s", contractAddress.Hex(), tokenID, err)

		return nil
	}

	// decode the data
	if !json.Valid(responseBody) {
		gbl.Log.Errorf("❌ error listings response json for %s/%d is invalid: %v", contractAddress.Hex(), tokenID, string(responseBody))

		return nil
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&listingsResponse); err != nil {
		gbl.Log.Errorf("❌ error while decoding listings for %s/%d: %s", contractAddress.Hex(), tokenID, err)

		return nil
	}

	if listingsResponse == nil {
		gbl.Log.Errorf("❌ error while decoding listings response for %s - %d: %s", contractAddress.Hex(), tokenID, err)

		return nil
	}

	if len(listingsResponse.Orders) == 0 {
		gbl.Log.Debugf("🤷‍♀️ no listings found for %s #%d", contractAddress.Hex(), tokenID)

		return nil
	}

	return listingsResponse.Orders
}

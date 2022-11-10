package analytics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
)

type ResponseOwner struct {
	Status   string   `json:"status"`
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
	Cursor   string   `json:"cursor"`
	Result   []*Owner `json:"result"`
}

type Owner struct {
	TokenAddress      string    `json:"token_address"`
	TokenID           string    `json:"token_id"`
	ContractType      string    `json:"contract_type"`
	OwnerOf           string    `json:"owner_of"`
	BlockNumber       string    `json:"block_number"`
	BlockNumberMinted string    `json:"block_number_minted"`
	TokenURI          string    `json:"token_uri"`
	Metadata          string    `json:"metadata"`
	Amount            string    `json:"amount"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	TokenHash         string    `json:"token_hash"`
	LastTokenURISync  time.Time `json:"last_token_uri_sync"`
	LastMetadataSync  time.Time `json:"last_metadata_sync"`
}

// viper.GetString("api_keys.moralis")

func FetchSetOwnersFor(apiToken string, set *CollectionSet) ([]common.Address, error) {
	tokenOwner := make(map[common.Address][]int, 0)

	for _, tokenID := range set.TokenIDs {
		owners, err := FetchOwnersFor(apiToken, set.ContractAddress, tokenID)
		if err != nil {
			gbl.Log.Errorf("❌ error fetching owners: %+v", err.Error())
			fmt.Printf("❌ error fetching owners: %+v\n", err.Error())
			// return setOwners, err
			continue
		}

		for _, owner := range owners {
			ownerAddress := common.HexToAddress(owner.OwnerOf)
			tokenOwner[ownerAddress] = append(tokenOwner[ownerAddress], tokenID)
		}

		// fmt.Printf("token: %d | tokenOwner: %v | setOwners: %v\n", tokenID, len(tokenOwner), len(setOwners))

		time.Sleep(time.Millisecond * 3337)
	}

	// fmt.Printf("tokenOwner: %v\n", len(tokenOwner))

	setOwners := make([]common.Address, 0)

	for ownerAddress, tokenIDs := range tokenOwner {
		if set.Any {
			setOwners = append(setOwners, ownerAddress)
		} else {
			if len(tokenIDs) == len(set.TokenIDs) {
				setOwners = append(setOwners, ownerAddress)
			}
		}
	}

	// fmt.Printf("tokenOwner: %v | setOwners: %v\n", len(tokenOwner), len(setOwners))

	return setOwners, nil
}

func FetchOwnersFor(apiToken string, contractAddress common.Address, tokenID int) (map[string]*Owner, error) {
	uniqueOwner := make(map[string]*Owner, 0)

	cursor := ""

	for {
		ownerResponse, err := getOwnersPage(apiToken, contractAddress, tokenID, cursor)
		if err != nil {
			fmt.Printf("❌ error getOwnersPage: %+v\n", err.Error())
			gbl.Log.Fatal(err)
		}

		// fmt.Println("ownerResponse.Result: ", len(ownerResponse.Result))
		// fmt.Printf("contractAddress: %s | token: %d | %s | %v\n", contractAddress, tokenID, apiToken, ownerResponse.Result)

		for _, owner := range ownerResponse.Result {
			uniqueOwner[owner.OwnerOf] = owner
		}

		if ownerResponse.Cursor == "" {
			break
		}

		cursor = ownerResponse.Cursor

		time.Sleep(time.Millisecond * 3737)
	}

	// fmt.Printf("uniqueOwner: %v\n", len(uniqueOwner))

	return uniqueOwner, nil
}

func getOwnersPage(apiToken string, contractAddress common.Address, tokenID int, cursor string) (*ResponseOwner, error) {
	var responseOwner *ResponseOwner

	var token string
	if tokenID == -1 {
		token = ""
	} else {
		token = "/" + strconv.Itoa(tokenID)
	}

	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2/nft/%s%s/owners?chain=eth&format=decimal&marketplace=opensea", contractAddress.String(), token)

	if cursor != "" {
		url = fmt.Sprintf("%s&cursor=%s", url, cursor)
	}

	// fmt.Printf("contractAddress: %s | token: %d | %s | %v\n", contractAddress, tokenID, apiToken, authHeader(apiToken))

	response, err := utils.HTTP.GetWithHeader(url, authHeader(apiToken))
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching current gas: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ gas oracle error: %+v", err.Error())
		}

		return responseOwner, err
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)

	// validate the data
	if !json.Valid(responseBody) {
		gbl.Log.Errorf("❌ invalid json: %+v\n\n", string(responseBody))
		return responseOwner, errors.New("invalid json")
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&responseOwner); err != nil {
		gbl.Log.Errorf("❌ error decoding json: %+v", err.Error())
		return responseOwner, err
	}

	// fmt.Println("responseOwner: ", responseOwner)

	return responseOwner, nil
}

func authHeader(apiToken string) http.Header {
	header := http.Header{}
	header.Add("X-API-KEY", apiToken)

	return header
}

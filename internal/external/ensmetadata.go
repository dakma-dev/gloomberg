package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
)

type ENSMetadataAttribute struct {
	TraitType   string      `json:"trait_type"`
	DisplayType string      `json:"display_type"`
	Value       interface{} `json:"value"`
}

// ENSMetadata defines model for ENSMetadata.
type ENSMetadata struct {
	Attributes      []ENSMetadataAttribute `json:"attributes"`
	BackgroundImage string                 `json:"background_image"`
	Description     string                 `json:"description"`
	ImageUrl        string                 `json:"image_url"`
	Name            string                 `json:"name"`
	NameLength      int                    `json:"name_length"`
	SegmentLength   int                    `json:"segment_length"`
	Url             string                 `json:"url"`
	Version         int                    `json:"version"`
}

const (
	ensMetadataAPI     = "https://metadata.ens.domains"
	ensContractAddress = "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85"
)

var ENSContract = common.HexToAddress(ensContractAddress)

func GetENSMetadataForTokenID(tokenID *big.Int) (*ENSMetadata, error) {
	if tokenID == nil {
		return nil, errors.New("tokenID is empty")
	}

	// build url
	url := ensMetadataAPI + "/" + "mainnet" + "/" + ensContractAddress + "/" + fmt.Sprint(tokenID)

	client, _ := newClient()
	request, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching ens metadata: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ ens metadata error: %+v", err.Error())
		}

		return nil, err
	}

	gbl.Log.Debugf("ens metadata response status: %s", response.Status)

	defer response.Body.Close()

	return parseENSMetadataResponse(response)
}

func parseENSMetadataResponse(response *http.Response) (*ENSMetadata, error) {
	bodyBytes, err := io.ReadAll(response.Body)

	defer func() { _ = response.Body.Close() }()

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var metadata ENSMetadata

		err = json.Unmarshal(bodyBytes, &metadata)
		if err != nil {
			return nil, err
		}

		return &metadata, nil
	}

	return nil, nil
}

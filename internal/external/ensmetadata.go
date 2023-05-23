package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/utils"
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
	ImageURL        string                 `json:"image_url"`
	Name            string                 `json:"name"`
	NameLength      int                    `json:"name_length"`
	SegmentLength   int                    `json:"segment_length"`
	URL             string                 `json:"url"`
	Version         int                    `json:"version"`
}

const ensMetadataAPI = "https://metadata.ens.domains"

func GetENSMetadataForTokenID(tokenID *big.Int) (*ENSMetadata, error) {
	if tokenID == nil {
		return nil, errors.New("tokenID is empty")
	}

	// build url
	url := ensMetadataAPI + "/" + "mainnet" + "/" + internal.ENSContractAddress.String() + "/" + fmt.Sprint(tokenID)

	response, err := utils.HTTP.Get(context.Background(), url)
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
		return &ENSMetadata{}, err
	}

	if response.StatusCode == http.StatusOK {
		var metadata ENSMetadata

		err = json.Unmarshal(bodyBytes, &metadata)
		if err != nil {
			return nil, err
		}

		return &metadata, nil
	}

	return nil, err
}

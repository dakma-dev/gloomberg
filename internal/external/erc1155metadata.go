package external

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/spf13/viper"
)

type ERC1155MetadataAttribute struct {
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
	DisplayType string `json:"display_type,omitempty"`
	MaxValue    string `json:"max_value,omitempty"`
}

type ERC1155MetadataImageDetails struct {
	Bytes  int    `json:"bytes"`
	Format string `json:"format"`
	Sha256 string `json:"sha256"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ERC1155Metadata struct {
	Name         string                      `json:"name"`
	CreatedBy    string                      `json:"created_by"`
	ExternalURL  string                      `json:"external_url"`
	Description  string                      `json:"description"`
	Attributes   []ERC1155MetadataAttribute  `json:"attributes"`
	ImageDetails ERC1155MetadataImageDetails `json:"image_details"`
	Image        string                      `json:"image"`
	ImageURL     string                      `json:"image_url"`
}

func GetERC1155MetadataForURI(url string, tokenID *big.Int) (*ERC1155Metadata, error) {
	url = strings.Replace(url, "ipfs://", viper.GetString("ipfs.gateway"), 1)
	url = strings.Replace(url, "{id}", tokenID.String(), -1)

	gbl.Log.Debugf("erc1155 metadata url: %+v", url)

	client, _ := newClient()
	request, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Debugf("⌛️ timeout while fetching erc1155 metadata: %+v", err.Error())
		} else {
			gbl.Log.Warnf("❌ erc1155 metadata error: %+v", err.Error())
		}

		return nil, err
	}

	gbl.Log.Debugf("erc1155 response status: %s", response.Status)

	defer response.Body.Close()

	return parseERC1155MetadataResponse(response)
}

func parseERC1155MetadataResponse(response *http.Response) (*ERC1155Metadata, error) {
	bodyBytes, err := io.ReadAll(response.Body)

	defer func() { _ = response.Body.Close() }()

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var metadata ERC1155Metadata

		err = json.Unmarshal(bodyBytes, &metadata)
		if err != nil {
			return nil, err
		}

		gbl.Log.Debugf("erc1155 metadata: %+v\n", metadata)

		return &metadata, nil
	}

	return nil, nil
}

package external

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
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

var ErrMetadataURLNotFound = errors.New("metadata url not found")

func GetERC1155MetadataForURI(ctx context.Context, url string, tokenID *big.Int) (*ERC1155Metadata, error) {
	if url == "" {
		log.Debugf("erc1155 metadata url is empty\n")

		return nil, ErrMetadataURLNotFound
	}

	url = utils.PrepareURL(url)
	url = strings.ReplaceAll(url, "{id}", tokenID.String())

	if url == "" || !strings.Contains(url, "://") {
		log.Debug("erc1155 metadata url is empty")

		return nil, ErrMetadataURLNotFound
	}

	log.Debugf("erc1155 metadata url: %+v", url)

	response, err := utils.HTTP.GetWithTLS12(ctx, url)
	if err != nil {
		if os.IsTimeout(err) {
			log.Debugf("⌛️ timeout while fetching erc1155 metadata: %+v", err.Error())
		} else {
			log.Warnf("❌ erc1155 metadata error | %s: %+v", url, err.Error())
		}

		return nil, err
	}

	log.Debugf("erc1155 response status: %s", response.Status)

	defer response.Body.Close()

	return parseERC1155MetadataResponse(response)
}

func parseERC1155MetadataResponse(response *http.Response) (*ERC1155Metadata, error) {
	bodyBytes, err := io.ReadAll(response.Body)

	defer func() { _ = response.Body.Close() }()

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		var metadata ERC1155Metadata

		err = json.Unmarshal(bodyBytes, &metadata)
		if err != nil {
			return nil, err
		}

		log.Debugf("erc1155 metadata: %+v\n", metadata)

		return &metadata, nil
	}

	return nil, err
}

package external

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"
)

type BlurSlugResponse struct {
	Success     bool `json:"success"`
	Collections []struct {
		ContractAddress string `json:"contractAddress"`
		CollectionSlug  string `json:"collectionSlug"`
		Name            string `json:"name"`
		ImageURL        string `json:"imageUrl"`
		TotalSupply     int    `json:"totalSupply"`
	} `json:"collections"`
}

func GetBlurSlugByName(collectionAddress common.Address) (string, error) {
	// build url
	url := "https://core-api.prod.blur.io/v1/search?query=" + collectionAddress.String()

	response, err := utils.HTTP.Get(url)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching blur slug: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ blur slug error: %+v", err.Error())
		}

		return "", err
	}

	gbl.Log.Infof("blur slug response status: %s", response.Status)

	defer response.Body.Close()

	if collectionList, err := parseBlurSlug(response); err == nil && collectionList != nil && len(collectionList.Collections) > 0 {
		gbl.Log.Infof("blur slug: %s", collectionList.Collections[0].CollectionSlug)
		return collectionList.Collections[0].CollectionSlug, nil
	} else {
		return "", err
	}
}

func parseBlurSlug(response *http.Response) (*BlurSlugResponse, error) {
	bodyBytes, err := io.ReadAll(response.Body)

	defer func() { _ = response.Body.Close() }()

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var blurSlugResponse BlurSlugResponse

		err = json.Unmarshal(bodyBytes, &blurSlugResponse)
		if err != nil {
			return nil, err
		}

		return &blurSlugResponse, nil
	}

	return nil, nil
}

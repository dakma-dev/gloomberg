package doamins

import (
	"encoding/json"
	"io"
	"net/http"
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

func parseResponse(rsp *http.Response) (*ENSMetadata, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)

	defer func() { _ = rsp.Body.Close() }()

	if err != nil {
		return nil, err
	}

	if rsp.StatusCode == 200 {
		var metadata ENSMetadata

		err = json.Unmarshal(bodyBytes, &metadata)
		if err != nil {
			return nil, err
		}

		return &metadata, nil
	}

	return nil, nil
}

package external

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/ethereum/go-ethereum/common"
)

type EventSignatureResponse struct {
	Count    int              `json:"count"`
	Next     interface{}      `json:"next"`
	Previous interface{}      `json:"previous"`
	Results  []EventSignature `json:"results"`
}

type EventSignature struct {
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	TextSignature  string    `json:"text_signature"`
	HexSignature   string    `json:"hex_signature"`
	BytesSignature string    `json:"bytes_signature"`
}

func GetEventSignature(hexSignature common.Hash) (EventSignature, error) {
	// build url
	url := "https://www.4byte.directory/api/v1/event-signatures/?hex_signature=" + hexSignature.String()

	response, err := utils.HTTP.Get(context.Background(), url)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching blur slug: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ blur slug error: %+v", err.Error())
		}

		return EventSignature{}, err
	}

	defer response.Body.Close()

	gbl.Log.Infof("4byte event signatures response status: %s", response.Status)

	// map to EventSignatures
	var eventSignatureRsponse EventSignatureResponse

	// read response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		gbl.Log.Errorf("❌ error reading 4byte event signatures response: %+v", err.Error())

		return EventSignature{}, err
	}

	// unmarshal body
	err = json.Unmarshal(bodyBytes, &eventSignatureRsponse)
	if err != nil {
		gbl.Log.Errorf("❌ error decoding 4byte event signatures response: %+v", err.Error())

		return EventSignature{}, err
	} else if len(eventSignatureRsponse.Results) == 0 {
		gbl.Log.Warnf("❌ no 4byte event signatures found for %s", hexSignature.String())

		return EventSignature{}, err
	}

	return eventSignatureRsponse.Results[0], nil
}

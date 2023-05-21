package oncecmd

import (
	"encoding/base64"
	"encoding/json"
	"math/big"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/abis/lawless"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type lawlessMetadata struct {
	Attributes  []interface{} `json:"attributes,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Image       string        `json:"image,omitempty"`
}

// getLawlessMetadata gets all metadata from the lawless contract
// use following command on the resulting json to get the number of items per type
//
// $ cat lawless_metadata.json| jq -r '.[] | .name' | sed 's/.*-//g' | sort -k3 | uniq -c | sort -n -r
func getLawlessMetadata(client *ethclient.Client) {
	allMetadata := make([]lawlessMetadata, 0)

	lawlessABI, err := lawless.NewLawlessCaller(common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790"), client)
	if err != nil {
		log.Errorf("failed to create lawless contract caller: %s", err)

		return
	}

	totalSupply, err := lawlessABI.TotalSupply(nil)
	if err != nil {
		log.Errorf("failed to get total supply: %s", err)

		return
	}

	for i := big.NewInt(0); i.Cmp(totalSupply) < 0; i.Add(i, big.NewInt(1)) {
		// get base64 encdoded metadata
		metadata, err := lawlessABI.TokenURI(nil, i)
		if err != nil {
			log.Errorf("failed to get token metadata: %s", err)

			return
		}

		// decode base64
		decodedMetadata, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(metadata, "data:application/json;base64,"))
		if err != nil {
			log.Errorf("failed to decode base64 metadata: %s", err)

			return
		}

		// unmarshal json
		var metadataJSON lawlessMetadata
		if err := json.Unmarshal(decodedMetadata, &metadataJSON); err != nil {
			log.Errorf("failed to unmarshal json metadata: %s", err)

			return
		}

		allMetadata = append(allMetadata, metadataJSON)
	}

	log.Print("\n")
	log.Print(len(allMetadata))
	log.Print("\n")

	// write to file
	metadataJSON, err := json.Marshal(allMetadata)
	if err != nil {
		log.Errorf("failed to marshal json metadata: %s", err)
	}

	if err := os.WriteFile("lawless_metadata.json", metadataJSON, 0o644); err != nil {
		log.Errorf("failed to write metadata to file: %s", err)
	}
}

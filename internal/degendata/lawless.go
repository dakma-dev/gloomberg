package degendata

import (
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

type LawlessMetadata struct {
	Attributes  []interface{} `json:"attributes,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Image       string        `json:"image,omitempty"`
}

var lawlessContractAddress = common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790")

func GetLawlessMetadata() []LawlessMetadata {
	if rawMetadata, ok := LoadedMetadata[lawlessContractAddress]; ok {
		log.Printf("metadata for contract %s already loaded", lawlessContractAddress.Hex())

		if md, ok := rawMetadata.(*[]LawlessMetadata); ok {
			return *md
		}
	}

	var metadata []LawlessMetadata
	err := getMetadatForCollection(lawlessContractAddress, &metadata)
	if err != nil {
		log.Errorf("error loading metadata for lawless contract %s: %s", lawlessContractAddress.Hex(), err)

		return nil
	}

	if metadata != nil {
		return metadata
	}

	log.Errorf("error loadincddg metadata for lawless contract %s", lawlessContractAddress.Hex())

	return nil
}

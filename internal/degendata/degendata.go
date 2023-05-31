package degendata

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

type Metadata struct {
	Attributes      []MetadataAttribute `json:"-"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Image           string              `json:"image"`
	TokenID         int64               `json:"token_id"`
	ContractAddress common.Address      `json:"-"`
	Score           Score               `json:"score,omitempty"`
}

type MetadataAttribute struct {
	TraitType string `json:"-"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

type Score struct {
	TokenID       int64   `json:"token_id"`
	Rank          int64   `json:"rank"`
	Score         float64 `json:"score"`
	TokenFeatures struct {
		UniqueAttributeCount int `json:"unique_attribute_count"`
	} `json:"token_features"`
	TokenMetadata map[string]interface{} `json:"token_metadata"`
}

var (
	Metadatas         = make(map[common.Address]map[int64]Metadata, 0)
	ErrNoMetadataFile = errors.New("no metadata file found")

	DegendataDir = "degendata"
	metadataDir  = path.Join(DegendataDir, "metadata")

	LoadedMetadata = make(map[common.Address]interface{}, 0)
)

func LoadMetadatas() error {
	rankAndMetadataFiles, err := os.ReadDir(metadataDir)
	if err != nil {
		log.Errorf("error reading metadata directory %s: %s", DegendataDir, err)
	}

	log.Debugf("Found metadata files: %+v", rankAndMetadataFiles)

	for _, f := range rankAndMetadataFiles {
		log.Debugf("%+v", f)

		if strings.HasPrefix(f.Name(), "0x") && strings.HasSuffix(f.Name(), ".zstd.gob") {
			log.Debugf("name %+v", f.Name())

			// var metadata interface{}
			metadata, err := ReadMetadataFromFile(path.Join(metadataDir, f.Name()))
			if err != nil {
				return err
			}

			contractAddress := common.HexToAddress(f.Name()[:42])

			for _, m := range metadata {
				if Metadatas[contractAddress] == nil {
					Metadatas[contractAddress] = make(map[int64]Metadata, 0)
				}

				Metadatas[contractAddress][m.TokenID] = m
			}

			log.Debugf("Loaded metadata for %s", contractAddress.Hex())
		}
	}

	fmt.Println()
	log.Printf("Metadatas %v", len(Metadatas))
	for k, v := range Metadatas {
		log.Printf("  %s | %d tokens", k.Hex(), len(v))
	}
	fmt.Println()

	return nil
}

// func getMetadatForCollection[T interface{}](contractAddress common.Address, metadata T) error {
// 	metadataFile := findMetadataFile(contractAddress)
// 	if metadataFile == "" {
// 		return ErrNoMetadataFile
// 	}

// 	file, err := os.Open(metadataFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	metadataDecoder := gob.NewDecoder(file)
// 	err = metadataDecoder.Decode(&metadata)
// 	if err != nil {
// 		return err
// 	}

// 	LoadedMetadata[contractAddress] = metadata

// 	return nil
// }

// func getMetadatFromFile(contractAddress common.Address, metadata interface{}) error {
// 	metadataFile := findMetadataFile(contractAddress)
// 	if metadataFile == "" {
// 		return ErrNoMetadataFile
// 	}

// 	file, err := os.Open(metadataFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	metadataDecoder := gob.NewDecoder(file)
// 	err = metadataDecoder.Decode(&metadata)
// 	if err != nil {
// 		return err
// 	}

// 	LoadedMetadata[contractAddress] = metadata

// 	return nil
// }

// func findMetadataFile(contractAddress common.Address) string {
// 	var metadataFile string

// 	files, err := os.ReadDir(metadataDir)
// 	if err != nil {
// 		log.Errorf("error reading metadata directory %s: %s", metadataDir, err)

// 		return metadataFile
// 	}

// 	for _, file := range files {
// 		if strings.HasPrefix(file.Name(), contractAddress.Hex()) && strings.HasSuffix(file.Name(), ".gob") {
// 			metadataFile = path.Join(metadataDir, file.Name())
// 			log.Debugf("found file: %s", metadataFile)
// 		}
// 	}

// 	return metadataFile
// }

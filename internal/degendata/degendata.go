package degendata

import (
	"encoding/gob"
	"errors"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNoMetadataFile = errors.New("no metadata file found")

	degendataDir = "degendata"

	LoadedMetadata = make(map[common.Address]interface{}, 0)
)

func getMetadatForCollection[T interface{}](contractAddress common.Address, metadata T) error {
	metadataFile := findMetadataFile(contractAddress)
	if metadataFile == "" {
		return ErrNoMetadataFile
	}

	file, err := os.Open(metadataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	metadataDecoder := gob.NewDecoder(file)
	err = metadataDecoder.Decode(&metadata)
	if err != nil {
		return err
	}

	LoadedMetadata[contractAddress] = metadata

	return nil
}

func findMetadataFile(contractAddress common.Address) string {
	var metadataFile string

	metadataDir := path.Join(degendataDir, "metadata")

	files, err := os.ReadDir(metadataDir)
	if err != nil {
		log.Errorf("error reading metadata directory %s: %s", metadataDir, err)

		return metadataFile
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), contractAddress.Hex()) && strings.HasSuffix(file.Name(), ".gob") {
			metadataFile = path.Join(metadataDir, file.Name())
			log.Debugf("found file: %s", metadataFile)
		}
	}

	return metadataFile
}

package degendb

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"
	"sort"

	"github.com/charmbracelet/log"
	"github.com/klauspost/compress/zstd"
)

type OpenSeaMetadata []struct {
	TokenIdentifier struct {
		ContractAddress string `json:"contract_address"`
		TokenID         int64  `json:"token_id"`
	} `json:"token_identifier"`
	MetadataDict  map[string]interface{} `json:"metadata_dict"`
	TokenStandard string                 `json:"token_standard"`
	Slug          string                 `json:"slug"`
}

type OpenSeaRanks map[int64]TokenRank

type TokenRank struct {
	Rank  int64   `json:"rank"`
	Score float64 `json:"score"`
}

func (tr TokenRank) GetRankSymbol(totalSupply uint64) string {
	if totalSupply == 0 {
		return ""
	}

	topX := float64(tr.Rank) / float64(totalSupply)

	var rankSymbol string

	switch {
	case topX <= 0.01:
		rankSymbol = "🥇"
	case topX <= 0.1:
		rankSymbol = "🥈"
	case topX <= 0.25:
		rankSymbol = "🥉"
	default:
		rankSymbol = ""
	}

	return rankSymbol
}

//
// helper & utility functions
//

func SortMapByValue(m map[string]int64, reverse bool) []string {
	sorted := make([]string, 0)

	for k := range m {
		sorted = append(sorted, k)
	}

	sort.Slice(sorted, func(i, j int) bool {
		if reverse {
			return m[sorted[i]] < m[sorted[j]]
		}

		return m[sorted[i]] > m[sorted[j]]
	})

	return sorted
}

func WriteDataToFile(data interface{}, filePath string) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(data)
	if err != nil {
		log.Errorf("failed to encode metadata: %s", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Errorf("failed to create file: %s", err)
	}
	defer file.Close()

	err = ZstdCompress(&buf, file)
	if err != nil {
		log.Errorf("failed to compress file: %s", err)
	}
}

// func ReadMetadataFromFile(filePath string) ([]degendb.TokenMetadata, error) {
// 	metadata := make([]degendb.TokenMetadata, 0)

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Errorf("failed to open file: %s", err)

// 		return metadata, err
// 	}
// 	defer file.Close()

// 	log.Debugf("file: %+v", file)

// 	decoder, err := zstd.NewReader(file, zstd.WithDecoderConcurrency(0))
// 	if err != nil {
// 		log.Errorf("failed to create zstd decoder: %s", err)
// 	}
// 	defer decoder.Close()

// 	err = gob.NewDecoder(decoder.IOReadCloser()).Decode(&metadata)
// 	if err != nil {
// 		return metadata, err
// 	}

// 	return metadata, nil
// }

func ZstdCompress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out, zstd.WithEncoderLevel(zstd.SpeedFastest))
	if err != nil {
		return err
	}

	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()

		return err
	}

	return enc.Close()
}

// func ReadOSRawDataFiles(gb *gloomberg.Gloomberg, filePath string) {
// 	rankAndMetadataFiles, err := os.ReadDir(filePath)
// 	if err != nil {
// 		log.Errorf("error reading metadata directory %s: %s", DegendataDir, err)
// 	}

// 	log.Debugf("Found metadata files: %+v", rankAndMetadataFiles)

// 	for _, f := range rankAndMetadataFiles {
// 		log.Debugf("%+v", f)

// 		if strings.HasPrefix(f.Name(), "ranks_") {
// 			slug := strings.TrimSuffix(strings.TrimPrefix(f.Name(), "ranks_"), ".json")

// 			// var metadata interface{}
// 			err := ReadOpenSeaMetadataAndRanks(gb, path.Dir(filePath), slug)
// 			if err != nil {
// 				log.Errorf("failed to read from file: %s", err)

// 				continue
// 			}
// 			// contractAddress := common.HexToAddress(f.Name()[:42])
// 			// log.Debugf("Loaded metadata for %s", contractAddress.Hex())
// 		}
// 	}
// }

// func ReadOpenSeaMetadataAndRanks(gb *gloomberg.Gloomberg, filePath string, slug string) error {
// 	slugAddresses := gb.CollectionDB.OpenseaAddressToSlug()

// 	tokens := make([]Token, 0)
// 	collections := make([]Collection, 0)

// 	osMetadata := make(OpenSeaMetadata, 0)
// 	osRanks := make(OpenSeaRanks, 0)

// 	metadataFileName := fmt.Sprintf("%s_cached_os_trait_data.json", slug)
// 	metadataFile, err := os.Open(filepath.Join(filePath, metadataFileName))
// 	if err != nil {
// 		log.Errorf("failed to open file: %s", err)

// 		return err
// 	}
// 	defer metadataFile.Close()

// 	log.Debugf("metadataFile: %+v", metadataFile)

// 	ranksFileName := fmt.Sprintf("ranks_%s.json", slug)
// 	ranksFile, err := os.Open(filepath.Join(filePath, ranksFileName))
// 	if err != nil {
// 		log.Errorf("failed to open file: %s", err)

// 		return err
// 	}
// 	defer ranksFile.Close()

// 	log.Debugf("ranksFile: %+v", ranksFile)

// 	metadataBytes, err := io.ReadAll(metadataFile)
// 	ranksBytes, err := io.ReadAll(ranksFile)

// 	// parse json to structs
// 	err = json.Unmarshal(metadataBytes, &osMetadata)
// 	if err != nil {
// 		log.Printf("failed to decode metadata json: %s", err)

// 		return err
// 	}

// 	// parse json to structs
// 	err = json.Unmarshal(ranksBytes, &osRanks)
// 	if err != nil {
// 		log.Printf("failed to decode ranks json: %s", err)

// 		return err
// 	}

// 	log.Debugf("osMetadata: %+v | osRanks: %+v", len(osMetadata), len(osRanks))

// 	if len(osMetadata) == 0 || len(osRanks) == 0 { // len(osMetadata) != len(osRanks) {
// 		log.Printf("osMetadata: %+v | osRanks: %+v", len(osMetadata), len(osRanks))

// 		return errors.New("metadata and ranks length mismatch")
// 	}

// 	collection := Collection{
// 		ID:          primitive.NewObjectID(),
// 		Address:     common.HexToAddress(osMetadata[0].TokenIdentifier.ContractAddress),
// 		Slugs:       Slugs{OpenSea: slug},
// 		Name:        slug,
// 		TotalSupply: len(osMetadata),
// 		CreatedAt:   time.Now(),
// 	}

// 	collections = append(collections, collection)

// 	// log.Printf("osMetadata[13]: %#v", osMetadata[:13])
// 	log.Debugf("osRanks[13]: %+v", osRanks[13])

// 	for _, v := range osMetadata {
// 		token := Token{
// 			ID:              primitive.NewObjectID(),
// 			Collection:      collection,
// 			ContractAddress: v.TokenIdentifier.ContractAddress,
// 			TokenID:         v.TokenIdentifier.TokenID,
// 			CollectionSlugs: Slugs{
// 				OpenSea: slugAddresses[common.HexToAddress(v.TokenIdentifier.ContractAddress)],
// 			},
// 			Rank: Rank{
// 				OpenSea: osRanks[v.TokenIdentifier.TokenID].Rank,
// 			},
// 			Score:     osRanks[v.TokenIdentifier.TokenID].Score,
// 			CreatedAt: time.Now(),
// 		}

// 		for k, v := range v.MetadataDict {
// 			var strValue string

// 			if val, ok := v.(string); ok {
// 				strValue = val
// 			} else {
// 				strValue = fmt.Sprintf("%v", v)
// 			}

// 			attribute := MetadataAttribute{
// 				Name:  k,
// 				Value: strValue,
// 			}

// 			if token.Rank.OpenSea == 0 {
// 				gbl.Log.Infof("tmissing rank for token: %+v", token)

// 				continue
// 			}

// 			token.Metadata = append(token.Metadata, attribute)
// 		}

// 		tokens = append(tokens, token)
// 	}

// 	for _, v := range tokens[len(tokens)-1:] {
// 		log.Debugf("%+v\n\n\n", v)
// 	}

// 	// gb.AddCollectionToken(collections, tokens)

// 	Metadatas[collection.Address] = make(map[int64]TokenMetadata)

// 	for _, t := range tokens {
// 		log.Debugf("%+v", t)

// 		// transform token to TokenMetadata
// 		tokenMetadata := TokenMetadata{
// 			Name:            t.Name,
// 			TokenID:         t.TokenID,
// 			ContractAddress: common.HexToAddress(t.ContractAddress),
// 			Attributes:      t.Metadata,
// 		}

// 		Metadatas[collection.Address][t.TokenID] = tokenMetadata
// 	}

// 	log.Printf("Metadata | collections: %+v", len(Metadatas))

// 	return nil
// }

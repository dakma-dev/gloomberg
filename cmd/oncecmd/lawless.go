package oncecmd

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"math/big"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"

// 	"github.com/benleb/gloomberg/internal/abis/lawless"
// 	"github.com/benleb/gloomberg/internal/degendb"
// 	"github.com/benleb/gloomberg/internal/degendb/scorer"
// 	"github.com/benleb/gloomberg/internal/style"
// 	"github.com/charmbracelet/log"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/spf13/viper"
// )

// var (
// 	lawlessContractAddress = common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790")
// 	lawlessOSSlug          = "thelawless"

// 	lawlessMetadataFile = filepath.Join("degendata", "metadata", common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790").Hex()+"_lawless.zstd.gob")

// 	maxTokensPerLine = 8
// 	maxTokenLines    = 10
// )

// func analyzeLawlessTokenNames(client *ethclient.Client) {
// 	lg.Print("\n\n\n  🏴  lawless\n\n\n")

// 	// get metadata
// 	metadata := getLawlessMetadata(client)

// 	preDot := make(map[string]int64, 0)
// 	afterDot := make(map[string]int64, 0)
// 	tokenType := make(map[string]int64, 0)

// 	tokens := make(map[string]degendb.TokenMetadata, 0)

// 	for _, token := range metadata {
// 		// strip id
// 		whiteSplit := strings.Split(token.Name, " ")
// 		if len(whiteSplit) <= 1 {
// 			continue
// 		}

// 		// pre dot
// 		dotSplit := strings.Split(whiteSplit[1], ".")
// 		if len(dotSplit) <= 1 {
// 			continue
// 		}

// 		// after dot / pre hyphen
// 		hyphenSplit := strings.Split(dotSplit[1], "-")
// 		if len(hyphenSplit) <= 1 {
// 			continue
// 		}

// 		tokens[whiteSplit[1]] = token
// 		preDot[dotSplit[0]]++
// 		afterDot[hyphenSplit[0]]++
// 		tokenType[hyphenSplit[1]]++
// 	}

// 	components := make(map[string]map[string]int64, 0)
// 	components["preDot"] = preDot
// 	components["afterDot"] = afterDot
// 	components["tokenType"] = tokenType

// 	for componentName, component := range components {
// 		for _, kind := range degendb.SortMapByValue(component, true)[:int(math.Min(float64(maxTokenLines), float64(len(component))))] {
// 			var fmtKind string

// 			msg := strings.Builder{}

// 			tokenLinks := make([]string, 0)

// 			// if component[kind] <= 5 {
// 			for name, token := range tokens {
// 				var checkFunc func(string, string) bool

// 				var checkKind string

// 				switch componentName {
// 				case "preDot":
// 					checkFunc = strings.HasPrefix
// 					checkKind = kind + "."
// 					fmtKind = style.Bold(kind) + style.GrayStyle.Render(".xx-xx")
// 				case "afterDot":
// 					checkFunc = strings.Contains
// 					checkKind = "." + kind + "-"
// 					fmtKind = style.GrayStyle.Render("xxx.") + style.Bold(kind) + style.GrayStyle.Render("-xx")
// 				case "tokenType":
// 					checkFunc = strings.HasSuffix
// 					checkKind = "-" + kind
// 					fmtKind = style.GrayStyle.Render("xxx.xx-") + style.Bold(kind)
// 				}

// 				if checkFunc(name, checkKind) {
// 					fmtName := strings.ReplaceAll(strings.Split(token.Name, " ")[1], kind, style.Bold(kind))
// 					tokenLinks = append(tokenLinks, style.TerminalLink(fmt.Sprintf("https://pro.opensea.io/collection/%s?view=%s&tokenAddress=%s", lawlessOSSlug, strings.Split(token.Name, " ")[0], lawlessContractAddress.Hex()), fmtName))
// 				}

// 				if len(tokenLinks) >= maxTokensPerLine {
// 					break
// 				}
// 			}

// 			msg.WriteString(fmt.Sprintf("  %s  %s %3d tokens", fmtKind, style.DarkGrayStyle.Render("|"), component[kind]))

// 			msg.WriteString(fmt.Sprintf("  %s  %s", style.DarkGrayStyle.Render("|"), strings.Join(tokenLinks, " ・ ")))

// 			lg.Print(msg.String())
// 		}
// 	}
// }

// // getLawlessMetadata gets all metadata from the lawless contract
// // use following command on the resulting json to get the number of items per type
// //
// // $ cat lawless_metadata.json| jq -r '.[] | .name' | sed 's/.*-//g' | sort -k3 | uniq -c | sort -n -r.
// func getLawlessMetadata(client *ethclient.Client) []degendb.TokenMetadata {
// 	allMetadata := make([]degendb.TokenMetadata, 0)

// 	if !viper.GetBool("fresh") {
// 		if allMetadata, err := degendb.ReadMetadataFromFile(lawlessMetadataFile); err == nil {
// 			log.Print("found metadata file ✓")
// 			return allMetadata
// 		} else if os.IsNotExist(err) {
// 			log.Info("metadata file does not exist, fetching from blockchain")
// 		} else {
// 			log.Errorf("failed to read metadata file: %s", err)
// 		}
// 	}

// 	lawlessABI, err := lawless.NewLawlessCaller(lawlessContractAddress, client)
// 	if err != nil {
// 		log.Errorf("failed to create lawless contract caller: %s", err)

// 		return nil
// 	}

// 	totalSupply, err := lawlessABI.TotalSupply(nil)
// 	if err != nil {
// 		log.Errorf("failed to get total supply: %s", err)

// 		return nil
// 	}

// 	scores := make(map[int64]degendb.Score, 0)
// 	// Let's first read the `config.json` file
// 	if content, err := ioutil.ReadFile(strings.Replace(lawlessMetadataFile, ".zstd.gob", ".json", 1) + "-scores"); err == nil {
// 		log.Debugf("content: %v", content)

// 		rawScores := make([]degendb.Score, 0)

// 		err = json.Unmarshal(content, &rawScores)
// 		if err != nil {
// 			log.Fatal("Error during Unmarshal(): ", err)
// 		}

// 		log.Debugf("rawScores: %v", rawScores)

// 		for _, score := range rawScores {
// 			scores[int64(score.TokenID)] = score
// 		}
// 	} else if os.IsNotExist(err) {
// 		log.Print("scores file does not exist")
// 	} else {
// 		log.Errorf("failed to read scores file: %s", err)
// 	}

// 	for i := big.NewInt(0); i.Cmp(totalSupply) < 0; i.Add(i, big.NewInt(1)) {
// 		// get base64 encdoded metadata
// 		metadata, err := lawlessABI.TokenURI(nil, i)
// 		if err != nil {
// 			log.Errorf("failed to get token metadata: %s", err)

// 			return nil
// 		}

// 		// decode base64
// 		decodedMetadata, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(metadata, "data:application/json;base64,"))
// 		if err != nil {
// 			log.Errorf("failed to decode base64 metadata: %s", err)

// 			return nil
// 		}

// 		// unmarshal json
// 		var metadataJSON degendb.TokenMetadata
// 		if err := json.Unmarshal(decodedMetadata, &metadataJSON); err != nil {
// 			log.Errorf("failed to unmarshal json metadata: %s", err)

// 			return nil
// 		}

// 		// split name to token id and name
// 		splittedName := strings.Split(metadataJSON.Name, " ")
// 		if len(splittedName) < 2 {
// 			continue
// 		}

// 		// parse token id
// 		tokenID, err := strconv.ParseInt(splittedName[0], 10, 64)
// 		if err != nil {
// 			log.Errorf("failed to parse token id: %s", err)

// 			return nil
// 		}

// 		// extra fields added for open-rarity scoring
// 		// (not yet sure if really needed)
// 		metadataJSON.TokenID = tokenID
// 		metadataJSON.ContractAddress = lawlessContractAddress

// 		// artificial metadata until we have the real one
// 		metadataJSON.Attributes = append(metadataJSON.Attributes,
// 			degendb.MetadataAttribute{
// 				Name:  "model",
// 				Value: splittedName[1][:3],
// 			},
// 			degendb.MetadataAttribute{
// 				Name:  "palette",
// 				Value: splittedName[1][4:6],
// 			},
// 			degendb.MetadataAttribute{
// 				Name:  "type",
// 				Value: splittedName[1][7:],
// 			},
// 		)

// 		// add score
// 		if score, ok := scores[tokenID]; ok {
// 			metadataJSON.Score = score
// 		}

// 		allMetadata = append(allMetadata, metadataJSON)
// 	}

// 	// write to json file
// 	marshslled, _ := json.MarshalIndent(allMetadata, "", " ")
// 	_ = ioutil.WriteFile(strings.Replace(lawlessMetadataFile, ".zstd.gob", ".json", 1), marshslled, 0o644)

// 	// gob.Register(degendb.MetadataAttribute{})

// 	lg.Print("\n")
// 	lg.Print(len(allMetadata))
// 	lg.Print("\n")

// 	degendb.WriteDataToFile(allMetadata, lawlessMetadataFile)

// 	return allMetadata
// }

// func TestOwnScorer() {
// 	metadata, err := degendb.ReadMetadataFromFile(lawlessMetadataFile)
// 	if err == nil {
// 		log.Print("found metadata file ✓")
// 	} else if os.IsNotExist(err) {
// 		log.Info("metadata file does not exist, fetching from blockchain")
// 	} else {
// 		log.Errorf("failed to read metadata file: %s", err)
// 	}

// 	// transform metadata to scorer.Token
// 	tokens := make([]*scorer.Token, 0)
// 	for _, m := range metadata {
// 		attributes := make([]scorer.TokenAttribute, 0)

// 		for _, a := range m.Attributes {
// 			attributes = append(attributes, scorer.TokenAttribute{
// 				Name:  a.Name,
// 				Value: a.Value,
// 			})
// 		}

// 		tokens = append(tokens, &scorer.Token{
// 			ContractAddress: m.ContractAddress,
// 			TokenID:         m.TokenID,
// 			Metadata:        attributes,
// 		})
// 	}

// 	collection := scorer.NewCollection("lawless", tokens)

// 	log.Printf("collection tokens: %+v", len(collection.Tokens))
// 	log.Printf("collection token: %+v", collection.Tokens[13])
// 	log.Printf("collection token: %+v", collection.Tokens[1337])

// 	ics := scorer.NewInformationContentScorer()

// 	res, err := ics.ScoreCollection(collection)
// 	if err != nil {
// 		log.Errorf("failed to score collection: %s", err)
// 	}

// 	log.Printf("score res: %+v", res)

// 	// tokenScore, err := ics.GetICScore(collection, tokens[13], make(map[string]*scorer.CollectionAttribute))
// 	// if err != nil {
// 	// 	log.Errorf("failed to get token score: %s", err)
// 	// }

// 	// log.Printf("token score: %+v", tokenScore)

// 	// log.Printf("collection: %+v", collection)
// }

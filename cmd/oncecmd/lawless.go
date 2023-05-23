package oncecmd

import (
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/benleb/gloomberg/internal/abis/lawless"
	"github.com/benleb/gloomberg/internal/style"
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

var (
	lawlessContractAddress = common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790")
	lawlessOSSlug          = "thelawless"

	lawlessMetadataFile = filepath.Join("degendata", "metadata", common.HexToAddress("0xb119ec7ee48928a94789ed0842309faf34f0c790").Hex()+"_lawless.zstd.gob")

	maxTokensPerLine = 8
	maxTokenLines    = 35
)

func analyzeLawlessTokenNames(client *ethclient.Client) {
	lg.Print("\n\n\n  🏴  lawless\n\n\n")

	// get metadata
	metadata := getLawlessMetadata(client)

	preDot := make(map[string]uint64, 0)
	afterDot := make(map[string]uint64, 0)
	tokenType := make(map[string]uint64, 0)

	tokens := make(map[string]lawlessMetadata, 0)

	for _, token := range metadata {
		// strip id
		whiteSplit := strings.Split(token.Name, " ")
		if len(whiteSplit) <= 1 {
			continue
		}

		// pre dot
		dotSplit := strings.Split(whiteSplit[1], ".")
		if len(dotSplit) <= 1 {
			continue
		}

		// after dot / pre hyphen
		hyphenSplit := strings.Split(dotSplit[1], "-")
		if len(hyphenSplit) <= 1 {
			continue
		}

		tokens[whiteSplit[1]] = token
		preDot[dotSplit[0]]++
		afterDot[hyphenSplit[0]]++
		tokenType[hyphenSplit[1]]++
	}

	components := make(map[string]map[string]uint64, 0)
	components["preDot"] = preDot
	components["afterDot"] = afterDot
	components["tokenType"] = tokenType

	for componentName, component := range components {
		for _, kind := range sortMapByValue(component, true)[:int(math.Min(float64(maxTokenLines), float64(len(component))))] {
			var fmtKind string

			msg := strings.Builder{}

			tokenLinks := make([]string, 0)

			// if component[kind] <= 5 {
			for name, token := range tokens {
				var checkFunc func(string, string) bool

				var checkKind string

				switch componentName {
				case "preDot":
					checkFunc = strings.HasPrefix
					checkKind = kind + "."
					fmtKind = style.Bold(kind) + style.GrayStyle.Render(".xx-xx")
				case "afterDot":
					checkFunc = strings.Contains
					checkKind = "." + kind + "-"
					fmtKind = style.GrayStyle.Render("xxx.") + style.Bold(kind) + style.GrayStyle.Render("-xx")
				case "tokenType":
					checkFunc = strings.HasSuffix
					checkKind = "-" + kind
					fmtKind = style.GrayStyle.Render("xxx.xx-") + style.Bold(kind)
				}

				if checkFunc(name, checkKind) {
					fmtName := strings.ReplaceAll(strings.Split(token.Name, " ")[1], kind, style.Bold(kind))
					tokenLinks = append(tokenLinks, style.TerminalLink(fmt.Sprintf("https://pro.opensea.io/collection/%s?view=%s&tokenAddress=%s", lawlessOSSlug, strings.Split(token.Name, " ")[0], lawlessContractAddress.Hex()), fmtName))
				}

				if len(tokenLinks) >= maxTokensPerLine {
					break
				}
			}

			msg.WriteString(fmt.Sprintf("  %s  %s %3d tokens", fmtKind, style.DarkGrayStyle.Render("|"), component[kind]))

			msg.WriteString(fmt.Sprintf("  %s  %s", style.DarkGrayStyle.Render("|"), strings.Join(tokenLinks, " ・ ")))

			lg.Print(msg.String())
		}

		fmt.Println()
	}
}

// getLawlessMetadata gets all metadata from the lawless contract
// use following command on the resulting json to get the number of items per type
//
// $ cat lawless_metadata.json| jq -r '.[] | .name' | sed 's/.*-//g' | sort -k3 | uniq -c | sort -n -r
func getLawlessMetadata(client *ethclient.Client) []lawlessMetadata {
	allMetadata := make([]lawlessMetadata, 0)

	if file, err := os.Open(lawlessMetadataFile); err == nil {
		err := gob.NewDecoder(file).Decode(&allMetadata)
		if err != nil {
			log.Errorf("failed to decode gob: %s", err)
		}

		return allMetadata
	} else if os.IsNotExist(err) {
		log.Info("metadata file does not exist, fetching from blockchain")
	}

	lawlessABI, err := lawless.NewLawlessCaller(lawlessContractAddress, client)
	if err != nil {
		log.Errorf("failed to create lawless contract caller: %s", err)

		return nil
	}

	totalSupply, err := lawlessABI.TotalSupply(nil)
	if err != nil {
		log.Errorf("failed to get total supply: %s", err)

		return nil
	}

	for i := big.NewInt(0); i.Cmp(totalSupply) < 0; i.Add(i, big.NewInt(1)) {
		// get base64 encdoded metadata
		metadata, err := lawlessABI.TokenURI(nil, i)
		if err != nil {
			log.Errorf("failed to get token metadata: %s", err)

			return nil
		}

		// decode base64
		decodedMetadata, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(metadata, "data:application/json;base64,"))
		if err != nil {
			log.Errorf("failed to decode base64 metadata: %s", err)

			return nil
		}

		// unmarshal json
		var metadataJSON lawlessMetadata
		if err := json.Unmarshal(decodedMetadata, &metadataJSON); err != nil {
			log.Errorf("failed to unmarshal json metadata: %s", err)

			return nil
		}

		allMetadata = append(allMetadata, metadataJSON)
	}

	lg.Print("\n")
	lg.Print(len(allMetadata))
	lg.Print("\n")

	writeDataToFile(allMetadata, lawlessMetadataFile)

	return allMetadata
}

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/benleb/gloomberg/internal/models/topic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

func WalletShortAddress(address common.Address) string {
	addressBytes := address.Bytes()

	return fmt.Sprint(
		"0x",
		fmt.Sprintf("%0.2x%0.2x", addressBytes[0], addressBytes[1]),
		"…",
		fmt.Sprintf("%0.2x%0.2x", addressBytes[len(addressBytes)-2], addressBytes[len(addressBytes)-1]),
	)
}

//
//const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
//
//var pattern = regexp.MustCompile(ansi)

//// StripANSI removes ANSI escape sequences from a string. From https://github.com/acarl005/stripansi
//func StripANSI(str string) string {
//	return pattern.ReplaceAllString(str, "")
//}

// ReplaceSchemeWithGateway func replaceSchemeWithGateway(url string, gateway string) string {
func ReplaceSchemeWithGateway(url string) string {
	const schemeIPFS = "ipfs://"

	return strings.Replace(url, schemeIPFS, viper.GetString("ipfs.gateway"), 1)
}

func PrettyString(str []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, str, "", "  "); err != nil {
		return err.Error()
	}

	return prettyJSON.String()
}

func ParseTopics(topics []common.Hash) (topic.Topic, common.Address, common.Address, *big.Int) {
	if len(topics) < 3 {
		//fmt.Printf("Invalid number of topics: %d", len(topics))
		return "", ZeroAddress, ZeroAddress, nil
	}

	logTopic := topic.Topic(topics[0].Hex())

	// parse from/to addresses
	fromAddress := common.HexToAddress(topics[1].Hex())
	toAddress := common.HexToAddress(topics[2].Hex())

	if logTopic == topic.TransferSingle {
		fromAddress = common.HexToAddress(topics[2].Hex())
		toAddress = common.HexToAddress(topics[3].Hex())
	}

	// parse token id
	rawTokenID := big.NewInt(0)
	if len(topics) >= 4 {
		// TODO WRONG FOR Transfer Single
		rawTokenID = topics[3].Big()
	}

	return logTopic, fromAddress, toAddress, rawTokenID
}

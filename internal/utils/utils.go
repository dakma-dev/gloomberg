package utils

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
)

// GetLinks returns the links to etherscan, opensea and blur.
func GetLinks(txHash common.Hash, contractAddress common.Address, tokenID int64) (string, string, string) {
	etherscanURL := GetEtherscanTxURL(txHash.String())
	openseaURL := GetOpenseaLink(contractAddress.String(), tokenID)
	blurURL := getBlurLink(contractAddress.String(), tokenID)

	return etherscanURL, openseaURL, blurURL
}

// etherscan.io.
func GetEtherscanTxURL(txHash string) string {
	return fmt.Sprintf("https://etherscan.io/tx/%s", txHash)
}

// etherscan.io.
func GetEtherscanTokenURL(txHash string) string {
	return fmt.Sprintf("https://etherscan.io/token/%s", txHash)
}

// blur.io.
func getBlurLink(contractAddress string, tokenID int64) string {
	return fmt.Sprintf("https://blur.io/asset/%s/%d", strings.ToLower(contractAddress), tokenID)
}

// opensea.io.
func GetOpenseaLink(contractAddress string, tokenID int64) string {
	return fmt.Sprintf("https://opensea.io/assets/%s/%d", contractAddress, tokenID)
}

func GetDexscreenerLink(contractAddress string) string {
	return fmt.Sprintf("https://dexscreener.com/ethereum/%s", strings.ToLower(contractAddress))
}

func GetTokenSnifferLink(contractAddress string) string {
	return fmt.Sprintf("https://tokensniffer.com/token/eth/%s", strings.ToLower(contractAddress))
}

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
// const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
//
// var pattern = regexp.MustCompile(ansi)

//// StripANSI removes ANSI escape sequences from a string. From https://github.com/acarl005/stripansi
// func StripANSI(str string) string {
//	return pattern.ReplaceAllString(str, "")
//}

// PrepareURL removes not allowed characters and replaces the ipfs:// scheme or "https://ipfs.io" with the configured ipfs gateway.
func PrepareURL(url string) string {
	const schemeIPFS = "ipfs://"

	// regex with characters allowed in a URL
	regexURL := regexp.MustCompile(`[^a-zA-Z0-9-_/:.,?&@=#%]`)
	// strip characters not in regex
	url = string(regexURL.ReplaceAll([]byte(url), []byte("")))

	// replace ipfs scheme/gateway
	url = strings.Replace(url, schemeIPFS, viper.GetString("ipfs.gateway"), 1)
	url = strings.Replace(url, "https://ipfs.io", viper.GetString("ipfs.gateway"), 1)

	return url
}

func ParseFirstTopic(topics []common.Hash) topic.Topic {
	logTopic := topic.Topic(topics[0].Hex())

	return logTopic
}

func ParseTopics(topics []common.Hash) (topic.Topic, common.Address, common.Address, *big.Int) {
	if len(topics) < 3 {
		return "", internal.ZeroAddress, internal.ZeroAddress, nil
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
		// WRONG FOR TransferSingle?
		rawTokenID = topics[3].Big()
	}

	return logTopic, fromAddress, toAddress, rawTokenID
}

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func WeiToGwei(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.GWei))
}

func EtherToWei(ether *big.Float) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.Set(ether), big.NewFloat(params.Wei))
}

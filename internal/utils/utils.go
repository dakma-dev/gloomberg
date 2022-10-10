package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func WalletShortAddress(address common.Address) string {
	addressBytes := address.Bytes()

	return fmt.Sprint(
		"0x",
		fmt.Sprintf("%0.2x%0.2x", addressBytes[0], addressBytes[1]),
		"…",
		fmt.Sprintf("%0.2x%0.2x", addressBytes[len(addressBytes)-2], addressBytes[len(addressBytes)-1]),
	)
}

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var pattern = regexp.MustCompile(ansi)

// StripANSI removes ANSI escape sequences from a string. From https://github.com/acarl005/stripansi
func StripANSI(str string) string {
	return pattern.ReplaceAllString(str, "")
}

// ReplaceSchemeWithGateway func replaceSchemeWithGateway(url string, gateway string) string {
func ReplaceSchemeWithGateway(url string) string {
	const schemeIPFS = "ipfs://"

	return strings.Replace(url, schemeIPFS, viper.GetString("ipfs.gateway"), 1)
}

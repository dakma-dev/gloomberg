package mintcmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/nemo/manifold"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// manifoldCmd represents the manifold command.
var manifoldCmd = &cobra.Command{
	Use:   "manifold",
	Short: "Mint a token from a Manifold mint.",
	Long: fmt.Sprintf(`%s


  Mints the token from the given Manifold URL or Identifier with the configured wallets.`, style.GetSmallHeader(internal.GloombergVersion)),
	Run: mintManifold,
}

func init() { MintCmd.AddCommand(manifoldCmd) }

func mintManifold(_ *cobra.Command, _ []string) {
	url := "https://app.manifold.xyz/c/above-the-noise"

	// rawWallets := &[]string{"0x1234", "0x5678"}

	// // check for valid flashbots signer key
	// if privateKey, err := crypto.HexToECDSA(viper.GetString("flots.signerKey")); err == nil {
	// 	// create the flashbots client
	// 	fbClient = flashbots.MustDial(viper.GetString("flots.relay"), signerKey)

	// 	flots.SignerPublicKey = &signerKey.PublicKey
	// } else {
	// 	log.Fatalf("❌ invalid or missing signer key: %v", err)
	// }

	mintIdentifier, err := GetMintIdentifier("https://app.manifold.xyz/c/above-the-noise")
	if err != nil {
		log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

		return
	}

	mintInfo, err := GetMintInfo(mintIdentifier)
	if err != nil {
		log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

		return
	}

	log.Printf("url: %s", url)
	log.Printf("mintIdentifier: %d", mintIdentifier)
	log.Printf("mintInfo: %#v", mintInfo)
}

func GetMintInfo(mintIdentifier int64) (*manifold.DataResponse, error) {
	url := "https://apps.api.manifoldxyz.dev/public/instance/data?id=" + fmt.Sprint(mintIdentifier)
	response, err := utils.HTTP.GetWithTLS12(context.TODO(), url)
	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("⌛️ GetMintInfo · timeout while fetching: %+v\n", err.Error())
		} else {
			log.Errorf("❌ gGetMintInfo · error: %+v\n", err.Error())
		}

		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Errorf("❌ gGetMintInfo · error: %+v\n", response.Status)

		return nil, err
	}

	// read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("❌ gGetMintInfo · response read error: %+v\n", err.Error())

		return nil, err
	}

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		log.Warnf("getContractMetadata invalid json: %s", err)

		return nil, err
	}

	// fmt.Println(string(responseBody))
	var decoded *manifold.DataResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decoded); err != nil {
		log.Errorf("❌  decode error: %s\n", err.Error())

		return nil, err
	}

	return decoded, nil
}

func GetMintIdentifier(url string) (int64, error) {
	response, err := utils.HTTP.GetWithTLS12(context.TODO(), url)
	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("⌛️ getting mint identifier from manifold failed · timeout while fetching: %+v\n", err.Error())
		} else {
			log.Errorf("❌ getting mint identifier from manifold failed · error: %+v\n", err.Error())
		}

		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Errorf("❌ getting mint identifier from manifold failed · error: %+v\n", response.Status)

		return 0, fmt.Errorf("non-200 status code while getting mint identifier: %+v", response.StatusCode)
	}

	// read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("❌ getting mint identifier from manifold failed · response read error: %+v\n", err.Error())

		return 0, err
	}

	// remove newlines from the response body
	body := strings.ReplaceAll(string(responseBody), "\n", "")

	// create the pattern to extract the identifier
	identifierPattern := regexp.MustCompile(`.*,IDENTIFIER:\"(\d+)\",.*`)

	// extract the identifier
	matches := identifierPattern.FindStringSubmatch(body)
	if len(matches) <= 1 {
		log.Errorf("❌ getting mint identifier from manifold failed · error: %+v\n", err.Error())
	}

	// convert the identifier to int64
	mintIdentifier, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		log.Errorf("❌ getting mint identifier from manifold failed · error: %+v\n", err.Error())

		return 0, err
	}

	return mintIdentifier, nil
}

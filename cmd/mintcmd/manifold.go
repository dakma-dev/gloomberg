// inspired by the manifold-minter of @timfame-codespace, thanks!
// https://github.com/timfame-codespace/manifold-minter/

package mintcmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/benleb/gloomberg/internal"
	manifoldABIs "github.com/benleb/gloomberg/internal/abis/manifold"
	"github.com/benleb/gloomberg/internal/nemo/manifold"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagURL        string
	flagIdentifier int64

	flagPrivateKeys []string
	flagRPCs        []string

	manifoldMagicAppID = 2537426615
)

// manifoldCmd represents the manifold command.
var manifoldCmd = &cobra.Command{
	Use:   "manifold",
	Short: "Mint a token from a Manifold mint.",
	Long: fmt.Sprintf(`%s


  Mints the token from the given Manifold URL %s Identifier with the configured wallets.`, style.Bold("or"), style.GetSmallHeader(internal.GloombergVersion)),
	Run: mintManifold,
}

func init() {
	MintCmd.AddCommand(manifoldCmd)

	// manifold url or identifier
	manifoldCmd.Flags().StringVar(&flagURL, "url", "", "manifold url to mint from (prefer to use identifier if possible)")
	manifoldCmd.Flags().Int64Var(&flagIdentifier, "identifier", -1, "manifold identifier (will be fetched from manifold if not set)")
	manifoldCmd.MarkFlagsMutuallyExclusive("url", "identifier")

	// private keys/wallets to mint with
	manifoldCmd.Flags().StringSliceVarP(&flagPrivateKeys, "keys", "p", make([]string, 0), "private keys/wallets to mint with")
	_ = viper.BindPFlag("mint.keys", manifoldCmd.Flags().Lookup("keys"))
	_ = manifoldCmd.MarkFlagRequired("keys")

	// rpc endpoints to use
	manifoldCmd.Flags().StringSliceVarP(&flagRPCs, "rpcs", "r", make([]string, 0), "rpc endpoints to mint with (randomly chosen)")
	_ = viper.BindPFlag("mint.rpcs", manifoldCmd.Flags().Lookup("rpcs"))
	_ = manifoldCmd.MarkFlagRequired("rpcs")

	manifoldCmd.Flags().Uint16("amount-tx", 1, "number of tokens to mint per transaction")
	_ = viper.BindPFlag("mint.manifold.amount-tx", manifoldCmd.Flags().Lookup("amount-tx"))
	manifoldCmd.Flags().Uint16("amount-wallet", 1, "number of tokens to mint per wallet/key")
	_ = viper.BindPFlag("mint.manifold.amount-wallet", manifoldCmd.Flags().Lookup("amount-wallet"))
}

type MintWallet struct {
	privateKey *ecdsa.PrivateKey
	address    *common.Address
}

func mintManifold(_ *cobra.Command, _ []string) {
	rpcClients := make([]*ethclient.Client, 0)
	mintWallets := make([]*MintWallet, 0)

	// check for valid keys
	for _, privateKey := range flagPrivateKeys {
		mintWallet := &MintWallet{}

		if key, err := crypto.HexToECDSA(privateKey); err == nil {
			mintWallet.privateKey = key
		} else {
			log.Errorf("❌ invalid or missing signer key: %v", err)
		}

		if publicKeyBytes := crypto.FromECDSAPub(&mintWallet.privateKey.PublicKey); publicKeyBytes != nil {
			log.Debugf("public Key: %s", style.AlmostWhiteStyle.Render(hexutil.Encode(publicKeyBytes)))
		}

		if address := crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey); address != (common.Address{}) {
			mintWallet.address = &address
			log.Debugf("address: %s", style.AlmostWhiteStyle.Render(mintWallet.address.Hex()))
		} else {
			log.Errorf("❌ getting address from public key failed | key: %v", mintWallet.privateKey.PublicKey)
		}

		mintWallets = append(mintWallets, mintWallet)
	}

	// connect to rpc endpoints
	for _, rpc := range flagRPCs {
		rpcClient, err := ethclient.Dial(rpc)
		if err != nil {
			log.Fatalf("❌ failed to connect to rpc endpoint %s: %v", rpc, err)
		}

		rpcClients = append(rpcClients, rpcClient)
	}

	mintIdentifier, err := getMintIdentifier("https://app.manifold.xyz/c/above-the-noise")
	if err != nil {
		log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

		return
	}

	mintInfo, err := getMintInfoWithURL(flagURL)
	if err != nil {
		log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

		return
	}

	manifoldInstanceID := *big.NewInt(int64(mintInfo.PublicData.ClaimIndex))

	if mintInfo.PublicData.ExtensionContractAddress != internal.ManifoldLazyClaimERC1155 {
		log.Printf("abi not implemented yet | extension contract address: %s", mintInfo.PublicData.ExtensionContractAddress)

		return
	}

	log.Print("")
	log.Print("")
	log.Printf("  %s  (by %s)", style.TerminalLink(flagURL, style.AlmostWhiteStyle.Render(mintInfo.PublicData.Name)), style.TerminalLink("https://twitter.com/"+mintInfo.Creator.TwitterURL, style.AlmostWhiteStyle.Render(mintInfo.Creator.Name)))
	log.Print("")

	log.Debugf("mint info: %#v", mintInfo)

	log.Print("")
	log.Print("")
	pretty.Println(mintInfo)
	log.Print("")
	log.Print("")

	log.Print("")
	log.Printf("price: %s", style.AlmostWhiteStyle.Render(fmt.Sprintf("%5.4f", mintInfo.MintPrice)))

	fmtWallets := make([]string, 0)
	for _, wallet := range mintWallets {
		fmtWallets = append(fmtWallets, style.AlmostWhiteStyle.Render(wallet.address.Hex()))
	}

	log.Print("")
	log.Printf("wallets: %s", strings.Join(fmtWallets, ", "))

	amountPerTx := viper.GetUint16("mint.manifold.amount-tx")
	amountPerWallet := uint16(math.Max(float64(amountPerTx), float64(viper.GetUint16("mint.manifold.amount-wallet"))))
	txsPerWallet := amountPerWallet / amountPerTx
	totalTxs := txsPerWallet * uint16(len(mintWallets))

	log.Print("")
	log.Printf("amount per tx: %s", style.AlmostWhiteStyle.Render(fmt.Sprint(amountPerTx)))
	log.Printf("amount per wallet: %s", style.AlmostWhiteStyle.Render(fmt.Sprint(amountPerWallet)))

	log.Print("")
	log.Printf("txs per wallet: %s", style.AlmostWhiteStyle.Render(fmt.Sprint(txsPerWallet)))
	log.Printf("total txs: %s", style.AlmostWhiteStyle.Render(fmt.Sprint(totalTxs)))

	log.Print("")
	log.Printf("collection/creator contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&mintInfo.PublicData.CreatorContractAddress), style.AlmostWhiteStyle.Render(mintInfo.PublicData.CreatorContractAddress.Hex())))

	if mintInfo.PublicData.ExtensionContractAddress != internal.ManifoldLazyClaimERC1155 {
		log.Printf("manifold/extension contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&mintInfo.PublicData.ExtensionContractAddress), style.AlmostWhiteStyle.Render(mintInfo.PublicData.ExtensionContractAddress.Hex())))
		log.Printf("manifold lazy claim erc1155 contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&internal.ManifoldLazyClaimERC1155), style.AlmostWhiteStyle.Render(internal.ManifoldLazyClaimERC1155.Hex())))
	}

	log.Printf("mintIdentifier: %d", mintIdentifier)
	log.Printf("manifoldInstanceID: %+v", style.AlmostWhiteStyle.Render(fmt.Sprint(manifoldInstanceID.Int64())))

	// get the mint fee (once)
	lazyClaimERC1155, err := manifoldABIs.NewLazyClaimERC1155(internal.ManifoldLazyClaimERC1155, rpcClients[rand.Intn(len(rpcClients))]) //nolint:gosec
	if err != nil {
		log.Error(err)

		return
	}

	manifoldFee, err := lazyClaimERC1155.MINTFEE(&bind.CallOpts{})
	if err != nil {
		log.Errorf("❌ getting mint fee failed: %s", style.AlmostWhiteStyle.Render(err.Error()))

		return
	}

	log.Printf("manifold fee: %s", style.AlmostWhiteStyle.Render(fmt.Sprintf("%7.5f", price.NewPrice(manifoldFee).Ether())))

	claimInfo, err := lazyClaimERC1155.GetClaim(&bind.CallOpts{}, mintInfo.PublicData.CreatorContractAddress, &manifoldInstanceID)
	if err != nil {
		log.Print("")
		log.Errorf("❌ getClaim(…) failed: %s", style.AlmostWhiteStyle.Render(err.Error()))

		return
	}

	log.Print("")
	log.Print("")
	pretty.Println(claimInfo)
	log.Print("")
	log.Print("")

	totalMints, err := lazyClaimERC1155.GetTotalMints(&bind.CallOpts{}, *mintWallets[0].address, mintInfo.PublicData.CreatorContractAddress, &manifoldInstanceID)
	if err != nil {
		log.Debugf("🤷‍♀️ getting total mints failed: %s", style.AlmostWhiteStyle.Render(err.Error()))
	} else {
		log.Printf("totalMints: %#v", totalMints)
	}

	log.Print("")
	log.Print("")

	for _, mintWallet := range mintWallets {
		// choose random rpc client
		rpcClient := rpcClients[0]
		if len(rpcClients) > 1 {
			rpcClient = rpcClients[rand.Intn(len(rpcClients))] //nolint:gosec
		}

		log.Printf("starting minter with wallet: %s", style.AlmostWhiteStyle.Render(mintWallet.address.String()))
		go mintERC1155(rpcClient, mintWallet, &manifoldInstanceID, mintInfo, claimInfo, manifoldFee)
	}

	select {}
}

func mintERC1155(rpcClient *ethclient.Client, mintWallet *MintWallet, manifoldInstanceID *big.Int, mintInfo *manifold.DataResponse, claimInfo manifoldABIs.IERC1155LazyPayableClaimClaim, manifoldFee *big.Int) {
	log.Printf("  minting with wallet: %s", mintWallet.address.String())

	lazyClaimERC1155, err := manifoldABIs.NewLazyClaimERC1155(internal.ManifoldLazyClaimERC1155, rpcClient)
	if err != nil {
		log.Errorf("❌ binding contract abi failed: %+v", style.AlmostWhiteStyle.Render(err.Error()))

		return
	}

	// get the nonce
	nonce, err := rpcClient.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey))
	if err != nil {
		log.Errorf("❌ getting nonce failed: %+v", style.AlmostWhiteStyle.Render(err.Error()))

		return
	}

	mintCost := utils.EtherToWei(big.NewFloat(mintInfo.MintPrice))
	totalCost := new(big.Int).Add(manifoldFee, mintCost)

	// create the transaction
	var tx *types.Transaction

	if amountPerTx := viper.GetUint16("mint.manifold.amount-tx"); amountPerTx > 1 {
		mintIndices := make([]uint32, 0)
		merkelProofs := make([][][32]byte, 0)

		for i := uint16(0); i < amountPerTx; i++ {
			mintIndices = append(mintIndices, uint32(0))
			merkelProofs = append(merkelProofs, [][32]byte{claimInfo.MerkleRoot})
		}

		tx, err = lazyClaimERC1155.MintBatch(&bind.TransactOpts{
			From:  crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey),
			Nonce: big.NewInt(int64(nonce)),
			Value: totalCost,
		}, mintInfo.PublicData.CreatorContractAddress, manifoldInstanceID, amountPerTx, mintIndices, merkelProofs, *mintWallet.address)
		if err != nil {
			log.Print("")
			log.Errorf("❌ creating transaction failed: %+v", style.AlmostWhiteStyle.Render(err.Error()))

			return
		}
	} else {
		tx, err = lazyClaimERC1155.Mint(&bind.TransactOpts{
			From:  crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey),
			Nonce: big.NewInt(int64(nonce)),
			Value: totalCost,
		}, mintInfo.PublicData.CreatorContractAddress, manifoldInstanceID, 0, [][32]byte{claimInfo.MerkleRoot}, *mintWallet.address)
		if err != nil {
			log.Print("")
			log.Errorf("❌ creating transaction failed: %+v", style.AlmostWhiteStyle.Render(err.Error()))

			return
		}
	}

	log.Printf("tx: %#v", tx)

	// sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), mintWallet.privateKey)
	if err != nil {
		log.Error(err)

		return
	}

	log.Printf("signedTx: %#v", signedTx)

	// send the transaction
	if !viper.GetBool("dev.mode") {
		err = rpcClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Error(err)

			return
		}
	}

	log.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func getMintInfoWithURL(mintURL string) (*manifold.DataResponse, error) {
	// get word after last / in url
	slug := mintURL[strings.LastIndex(mintURL, "/")+1:]

	url := fmt.Sprintf("https://apps.api.manifoldxyz.dev/public/instance/data?appId=%d&instanceSlug=%s", manifoldMagicAppID, slug)
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

func getMintIdentifier(url string) (int64, error) {
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

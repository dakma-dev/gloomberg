// inspired by the manifold-minter of @timfame-codespace, thanks!
// https://github.com/timfame-codespace/manifold-minter/

package mintcmd

import (
	"bytes"
	"context"
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
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	manifoldABIs "github.com/benleb/gloomberg/internal/abis/manifold"
	"github.com/benleb/gloomberg/internal/nemo/manifold"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
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
	flagManifoldInstanceID int64

	manifoldMagicAppID = 2537426615
)

// manifoldCmd represents the manifold command.
var manifoldCmd = &cobra.Command{
	Use:   "manifold",
	Short: "Mint a token from a Manifold mint.",
	// 	Long: fmt.Sprintf(`%s

	//   Mints the token from the given Manifold URL %s Identifier with the configured wallets.`, style.Bold("or"), style.GetSmallHeader(internal.GloombergVersion)),
	Run: mintManifold,
}

func init() {
	MintCmd.AddCommand(manifoldCmd)

	// manifold url or identifier
	manifoldCmd.Flags().StringVar(&flagURL, "url", "", "manifold url to mint from (prefer to use identifier if possible)")
	manifoldCmd.Flags().Int64Var(&flagManifoldInstanceID, "instance-id", -1, "manifold identifier (will be fetched from manifold if not set)")
	manifoldCmd.MarkFlagsMutuallyExclusive("url", "instance-id")

	// do not wait for mint start
	manifoldCmd.Flags().Bool("no-wait", false, "do not wait for mint start")
	_ = viper.BindPFlag("mint.manifold.nowait", manifoldCmd.Flags().Lookup("no-wait"))

	// private keys/wallets to mint with
	manifoldCmd.Flags().StringSliceVarP(&flagPrivateKeys, "keys", "p", make([]string, 0), "private keys/wallets to mint with")
	_ = viper.BindPFlag("mint.keys", manifoldCmd.Flags().Lookup("keys"))
	// _ = manifoldCmd.MarkFlagRequired("keys")

	// rpc endpoints to use
	manifoldCmd.Flags().StringSliceVarP(&flagRPCs, "rpcs", "r", make([]string, 0), "rpc endpoints to mint with (randomly chosen)")
	_ = viper.BindPFlag("mint.rpcs", manifoldCmd.Flags().Lookup("rpcs"))
	// _ = manifoldCmd.MarkFlagRequired("rpcs")

	// gas settings
	manifoldCmd.Flags().Float64Var(&flagGasFeeCapMultiplier, "fee-cap", 1.0, "gas fee cap multiplier")
	_ = viper.BindPFlag("mint.fee_multiplier", manifoldCmd.Flags().Lookup("fee-cap"))
	manifoldCmd.Flags().Float64Var(&flagGasTipCapMultiplier, "tip-cap", 1.0, "gas tip cap multiplier")
	_ = viper.BindPFlag("mint.tip_multiplier", manifoldCmd.Flags().Lookup("tip-cap"))

	// number of wallets to use
	manifoldCmd.Flags().Uint16("num-wallets", 3, "number of wallets to use for minting")
	_ = viper.BindPFlag("mint.manifold.num-wallets", manifoldCmd.Flags().Lookup("num-wallets"))

	manifoldCmd.Flags().Uint16("amount-tx", 1, "number of tokens to mint per transaction")
	_ = viper.BindPFlag("mint.manifold.amount-tx", manifoldCmd.Flags().Lookup("amount-tx"))
	manifoldCmd.Flags().Uint16("amount-wallet", 1, "number of tokens to mint per wallet/key")
	_ = viper.BindPFlag("mint.manifold.amount-wallet", manifoldCmd.Flags().Lookup("amount-wallet"))
}

func mintManifold(_ *cobra.Command, _ []string) {
	// rpcClients := make([]*ethclient.Client, 0)
	// availableWallets := make([]*MintWallet, 0)
	rpcClients := mapset.NewSet[*ethclient.Client]()
	rpcEndpoints := mapset.NewSet[string]()

	availableWallets := make([]*MintWallet, 0)
	mintWallets := mapset.NewSet[*MintWallet]()

	// negate for easier handling...
	waitForStart := !viper.GetBool("mint.manifold.nowait")

	// check for valid keys
	for _, privateKey := range viper.GetStringSlice("mint.keys") {
		mintWallet := &MintWallet{}

		if key, err := crypto.HexToECDSA(privateKey); err == nil {
			mintWallet.privateKey = key
		} else {
			log.Errorf("❌ invalid or missing signer key: %v", err)
		}

		if publicKeyBytes := crypto.FromECDSAPub(&mintWallet.privateKey.PublicKey); publicKeyBytes != nil {
			log.Debugf("public Key: %s", style.BoldAlmostWhite(hexutil.Encode(publicKeyBytes)))
		}

		if address := crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey); address != (common.Address{}) {
			mintWallet.address = &address
			log.Debugf("address: %s", style.BoldAlmostWhite(mintWallet.address.Hex()))
		} else {
			log.Errorf("❌ getting address from public key failed | key: %v", mintWallet.privateKey.PublicKey)
		}

		mintWallet.color = style.GenerateColorWithSeed(mintWallet.address.Hash().Big().Int64())
		mintWallet.tag = lipgloss.NewStyle().Foreground(mintWallet.color).Render(style.ShortenAddress(mintWallet.address))

		availableWallets = append(availableWallets, mintWallet)
	}

	// connect to rpc endpoints
	for _, rpc := range viper.GetStringSlice("mint.rpcs") {
		rpcEndpoints.Add(rpc)

		rpcClient, err := ethclient.Dial(rpc)
		if err != nil {
			log.Fatalf("❌ failed to connect to rpc endpoint %s: %v", rpc, err)
		}

		rpcClients.Add(rpcClient)
	}

	var mintInfo *manifold.DataResponse
	var manifoldInstanceID big.Int

	switch {
	case flagURL != "":
		mintIdentifier, err := getMintIdentifier(flagURL)
		if err != nil {
			log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

			return
		}

		log.Debugf("url mintIdentifier: %d", mintIdentifier)

		mintInfo, err = getMintInfoWithURL(flagURL)
		if err != nil {
			log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

			return
		}
	case flagManifoldInstanceID > 0:
		var err error

		mintInfo, err = getMintInfoWithInstanceID(flagManifoldInstanceID)
		if err != nil {
			log.Fatalf("❌ getting mint identifier from manifold failed: %v", err)

			return
		}
	default:
		log.Fatalf("❌ no url or identifier given")

		return
	}

	manifoldInstanceID = *big.NewInt(int64(mintInfo.PublicData.ClaimIndex))

	if mintInfo.PublicData.ExtensionContractAddress != internal.ManifoldLazyClaimERC1155 {
		log.Printf("abi not implemented yet | extension contract address: %s", mintInfo.PublicData.ExtensionContractAddress)

		return
	}

	mintURL := flagURL

	if mintInfo.Slug != "" {
		mintURL = fmt.Sprintf("https://app.manifold.xyz/c/%s", mintInfo.Slug)
	}

	log.Print("")
	log.Print("")
	log.Printf("  %s  (by %s)", style.TerminalLink(mintURL, style.BoldAlmostWhite(mintInfo.PublicData.Name)), style.TerminalLink("https://twitter.com/"+mintInfo.Creator.TwitterURL, style.BoldAlmostWhite(mintInfo.Creator.Name)))
	log.Print("")

	log.Debugf("mint info: %#v", mintInfo)

	log.Print("")
	log.Print(style.BoldAlmostWhite("configuration"))
	log.Print("")

	fmtWallets := make([]string, 0)
	for _, wallet := range availableWallets {
		fmtWallets = append(fmtWallets, style.Bold(wallet.tag))
	}

	log.Printf("  available wallets: %s", strings.Join(fmtWallets, ", "))

	//
	// randomly choose numWallets from our available wallets

	// (max) number of wallets to use
	numWallets := uint16(math.Min(float64(viper.GetUint16("mint.manifold.num-wallets")), float64(len(availableWallets))))

	// select the first numWallets from our available wallets
	for _, wallet := range availableWallets[:numWallets] {
		mintWallets.Add(wallet)
	}

	fmtWallets = make([]string, 0)
	for _, wallet := range mintWallets.ToSlice() {
		fmtWallets = append(fmtWallets, style.Bold(wallet.tag))
	}

	log.Printf("  mint wallets: %s", strings.Join(fmtWallets, ", "))

	amountPerTx := viper.GetUint16("mint.manifold.amount-tx")
	amountPerWallet := uint16(math.Max(float64(amountPerTx), float64(viper.GetUint16("mint.manifold.amount-wallet"))))
	txsPerWallet := amountPerWallet / amountPerTx
	totalTxs := txsPerWallet * uint16(mintWallets.Cardinality())

	log.Print("")
	log.Printf("  amount per tx: %s", style.BoldAlmostWhite(fmt.Sprint(amountPerTx)))
	log.Printf("  amount per wallet: %s", style.BoldAlmostWhite(fmt.Sprint(amountPerWallet)))

	log.Print("")
	log.Printf("  → txs per wallet: %s", style.BoldAlmostWhite(fmt.Sprint(txsPerWallet)))
	log.Printf("  → total txs: %s", style.BoldAlmostWhite(fmt.Sprint(totalTxs)))

	//
	// manifold api info
	//

	log.Print("")
	log.Print("")
	log.Print(style.BoldAlmostWhite("manifold info") + " (from api)")

	log.Print("")
	log.Printf("  price: %s", style.BoldAlmostWhite(fmt.Sprintf("%5.4f", mintInfo.MintPrice)))

	log.Print("")
	log.Printf("  merkeleTreeID: %s", style.BoldAlmostWhite(fmt.Sprintf("%d", mintInfo.PublicData.MerkleTreeID)))

	log.Print("")
	log.Printf("  collection/creator contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&mintInfo.PublicData.CreatorContractAddress), style.BoldAlmostWhite(mintInfo.PublicData.CreatorContractAddress.Hex())))

	if mintInfo.PublicData.ExtensionContractAddress != internal.ManifoldLazyClaimERC1155 {
		log.Printf("  manifold/extension contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&mintInfo.PublicData.ExtensionContractAddress), style.BoldAlmostWhite(mintInfo.PublicData.ExtensionContractAddress.Hex())))
		log.Printf("  manifold lazy claim erc1155 contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&internal.ManifoldLazyClaimERC1155), style.BoldAlmostWhite(internal.ManifoldLazyClaimERC1155.Hex())))
	}

	// log.Printf("  mintIdentifier: %s", style.BoldAlmostWhite(fmt.Sprint(mintIdentifier)))
	log.Printf("  manifoldInstanceID: %+v", style.BoldAlmostWhite(fmt.Sprint(manifoldInstanceID.Int64())))

	//
	// manifold chain info
	//

	log.Print("")
	log.Print("")
	log.Print(style.BoldAlmostWhite("manifold info") + " (from chain)")

	// get the mint fee (once)
	lazyClaimERC1155, err := manifoldABIs.NewLazyClaimERC1155(internal.ManifoldLazyClaimERC1155, rpcClients.ToSlice()[rand.Intn(len(rpcClients.ToSlice()))]) //nolint:gosec
	if err != nil {
		log.Error(err)

		return
	}

	manifoldFee, err := lazyClaimERC1155.MINTFEE(&bind.CallOpts{})
	if err != nil {
		log.Errorf("❌ getting mint fee failed: %s", style.BoldAlmostWhite(err.Error()))

		return
	}

	log.Print("")
	log.Printf("  fee: %s", style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(manifoldFee).Ether())))

	claimInfo, err := lazyClaimERC1155.GetClaim(&bind.CallOpts{}, mintInfo.PublicData.CreatorContractAddress, &manifoldInstanceID)
	if err != nil {
		log.Errorf("❌ getClaim(…) failed: %s", style.BoldAlmostWhite(err.Error()))

		return
	}

	startDate := time.Unix(claimInfo.StartDate.Int64(), 0)

	log.Print("")
	log.Printf("  mint start: %+v", style.BoldAlmostWhite(fmt.Sprint(startDate.Format("2006-01-02 15:04:05"))))
	log.Printf("            → in %+v", style.BoldAlmostWhite(fmt.Sprint(time.Until(startDate).Truncate(time.Second).String())))
	log.Print("")
	log.Printf("  minted: %+v / %v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.Total)), style.BoldAlmostWhite(fmt.Sprint(claimInfo.TotalMax)))
	log.Printf("  remaining: %+v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.TotalMax-claimInfo.Total)))
	log.Printf("  max/wallet: %+v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.WalletMax)))

	totalMints, err := lazyClaimERC1155.GetTotalMints(&bind.CallOpts{}, *mintWallets.ToSlice()[0].address, mintInfo.PublicData.CreatorContractAddress, &manifoldInstanceID)
	if err != nil {
		log.Debugf("🤷‍♀️ getting total mints failed: %s", style.BoldAlmostWhite(err.Error()))
	} else {
		log.Printf("  totalMints: %d", totalMints)
	}

	log.Print("")

	if startDate.After(time.Now()) && waitForStart {
		log.Print("")
		log.Print("")
		log.Printf(" 💤 💤 💤  waiting for mint start in %s  💤 💤 💤", style.BoldAlmostWhite(fmt.Sprint(time.Until(startDate).Truncate(time.Second).String())))
		log.Print("")
		log.Printf(style.GrayStyle.Render("    (use --no-wait to skip waiting)"))
		log.Print("")
		log.Print("")

		// sleep until 5s before startDate
		time.Sleep(time.Until(startDate.Add(-2 * time.Second)))
	}

	// start the minting jobs

	log.Print("")
	log.Print("")
	log.Print(" 🚀 🚀 🚀  starting minter jobs  🚀 🚀 🚀")
	log.Print("")
	log.Print("")

	wg := sync.WaitGroup{}

	for _, mintWallet := range mintWallets.ToSlice() {
		wg.Add(1)

		go func(mintWallet *MintWallet) {
			defer wg.Done()

			mintERC1155(rpcEndpoints.Clone(), mintWallet, txsPerWallet, &manifoldInstanceID, mintInfo, claimInfo, manifoldFee)
		}(mintWallet)
	}

	wg.Wait()

	log.Print("")
	log.Print("  🍹 all jobs finished! 🍹")
	log.Print("")
}

func mintERC1155(rpcEndpoints mapset.Set[string], mintWallet *MintWallet, txsPerWallet uint16, manifoldInstanceID *big.Int, mintInfo *manifold.DataResponse, claimInfo manifoldABIs.IERC1155LazyPayableClaimClaim, manifoldFee *big.Int) {
	txConfirmed := 0

	prErr := func(mintInfo *manifold.DataResponse, claimInfo manifoldABIs.IERC1155LazyPayableClaimClaim) {
		// log.Print("  ❌ ❌ ❌  error  ❌ ❌ ❌")
		// log.Print("❕ mintInfo 💁‍♀️")
		// pretty.Println(mintInfo)
		// log.Print("")
		// log.Print("❕ claimInfo 💁‍♀️")
		// pretty.Println(claimInfo)
		log.Print("")
	}

	log.Print("\n\n")

	for {
		// choose random rpc endpoint
		rpcIdx := rand.Intn(len(rpcEndpoints.ToSlice())) //nolint:gosec
		rpcEndpoint := rpcEndpoints.ToSlice()[rpcIdx]

		log.Debugf("%s | 📑 rpc endpoints (%d): %s", mintWallet.tag, rpcEndpoints.Cardinality(), style.BoldAlmostWhite(fmt.Sprintf("%+v", rpcEndpoints)))
		log.Debugf("%s | 🎯 selected rpc endpoint: %s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(rpcIdx)))

		// connect to rpc endpoint
		rpcClient, err := ethclient.Dial(rpcEndpoint)
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("❌ failed to connect to rpc endpoint %s: %v", rpcEndpoint, err)

			continue
		}

		// contract binding
		log.Printf("%s | 🧶 create contract binding...", mintWallet.tag)

		lazyClaimERC1155, err := manifoldABIs.NewLazyClaimERC1155(internal.ManifoldLazyClaimERC1155, rpcClient)
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("❌ binding contract abi failed: %+v", style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Debugf("%s | lazyClaimERC1155: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(lazyClaimERC1155)))

		// get the nonce
		nonce, err := rpcClient.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey))
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("%s | ❌ getting nonce failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Printf("%s | nonce: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(nonce)))

		// get the current gas price
		gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("%s | ❌ getting gasPrice failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Printf("%s | ⛽️ gasPrice: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(gasPrice)))

		// get the current gas tip
		gasTip, err := rpcClient.SuggestGasTipCap(context.Background())
		if err != nil {
			log.Errorf("%s | ❌ getting gasTip failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))
			prErr(mintInfo, claimInfo)

			continue
		}

		log.Printf("%s | ⛽️ gasTip: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(gasTip)))

		//
		// apply gas multiplier
		feeCapMultiplier := new(big.Float).SetFloat64(viper.GetFloat64("mint.fee_multiplier"))
		tipCapMultiplier := new(big.Float).SetFloat64(viper.GetFloat64("mint.tip_multiplier"))

		suggestedFee := new(big.Float).SetInt(gasPrice)
		suggestedTip := new(big.Float).SetInt(gasTip)

		gasFeeCap, _ := new(big.Float).Mul(suggestedFee, feeCapMultiplier).Int(nil)
		gasTipCap, _ := new(big.Float).Mul(suggestedTip, tipCapMultiplier).Int(nil)

		// 💸 💸 💸
		mintCost := utils.EtherToWei(big.NewFloat(mintInfo.MintPrice))
		totalCost := new(big.Int).Add(manifoldFee, mintCost)

		log.Print("")
		log.Printf("  totalCost: %s", style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(totalCost).Ether())))

		//
		// create the transaction options
		txOpts, err := bind.NewKeyedTransactorWithChainID(mintWallet.privateKey, big.NewInt(1))
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("%s | ❌ creating transaction options failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Printf("%s | mintCost: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(mintCost).Ether())))
		log.Printf("%s | totalCost: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(totalCost).Ether())))

		txOpts.From = crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey)
		txOpts.Nonce = big.NewInt(int64(nonce))
		txOpts.Value = totalCost
		txOpts.GasFeeCap = gasFeeCap
		txOpts.GasTipCap = gasTipCap

		if viper.GetBool("dev.mode") {
			txOpts.NoSend = true
		}

		log.Printf("%s | txOpts: %#v", mintWallet.tag, txOpts)

		// create the transaction
		var sentTx *types.Transaction

		if amountPerTx := viper.GetUint16("mint.manifold.amount-tx"); amountPerTx > 1 {
			mintIndices := make([]uint32, 0)
			merkelProofs := make([][][32]byte, 0)

			for i := uint16(0); i < amountPerTx; i++ {
				mintIndices = append(mintIndices, uint32(0))
				merkelProofs = append(merkelProofs, [][32]byte{claimInfo.MerkleRoot})
			}

			sentTx, err = lazyClaimERC1155.MintBatch(txOpts, mintInfo.PublicData.CreatorContractAddress, manifoldInstanceID, amountPerTx, mintIndices, merkelProofs, *mintWallet.address)
			if err != nil {
				prErr(mintInfo, claimInfo)
				log.Printf("%s | ❌ creating batch transaction failed: %+v | %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()), err)
			}
		} else {
			sentTx, err = lazyClaimERC1155.Mint(txOpts, mintInfo.PublicData.CreatorContractAddress, manifoldInstanceID, 0, [][32]byte{claimInfo.MerkleRoot}, *mintWallet.address)
			if err != nil {
				prErr(mintInfo, claimInfo)
				log.Printf("%s | ❌ creating transaction failed: %+v | %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()), err)
			}
		}

		if sentTx == nil {
			log.Printf("%s | ❌ executing transaction failed - sentTx is %+v", mintWallet.tag, sentTx)

			continue
		}

		log.Printf("%s | 🙌 tx sent! → %s 🙌", mintWallet.tag, style.TerminalLink(utils.GetEtherscanTxURL(sentTx.Hash().Hex()), style.BoldAlmostWhite(sentTx.Hash().Hex())))

		// wait for the transaction to be mined
		receipt, err := bind.WaitMined(context.Background(), rpcClient, sentTx)
		if err != nil {
			prErr(mintInfo, claimInfo)

			log.Printf("%s | ❌ transaction failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		if receipt != nil {
			log.Printf("%s | 🎉 transaction mined! → %s 🎉", mintWallet.tag, style.TerminalLink(utils.GetEtherscanTxURL(receipt.TxHash.Hex()), style.BoldAlmostWhite(receipt.TxHash.Hex())))
			pretty.Println(receipt)

			txConfirmed++
		}

		if txConfirmed >= int(txsPerWallet) {
			log.Printf("%s | 💥 🙌 💥 all txs confirmed! 🍹 🥃 🥂s", mintWallet.tag)

			return
		}

		time.Sleep(time.Millisecond * 337)
	}
}

func getMintInfoWithInstanceID(identifier int64) (*manifold.DataResponse, error) {
	url := fmt.Sprintf("https://apps.api.manifoldxyz.dev/public/instance/data?id=%d", identifier)

	log.Debugf("Identifier url: %s", url)

	response, err := utils.HTTP.GetWithTLS12(context.TODO(), url)
	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("⌛️ Identifier GetMintInfo · timeout while fetching: %+v\n", err.Error())
		} else {
			log.Errorf("❌ Identifier gGetMintInfo · error: %+v\n", err.Error())
		}

		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Errorf("❌ Identifier gGetMintInfo · error: %+v\n", response.Status)

		return nil, err
	}

	// read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("❌ Identifier gGetMintInfo · response read error: %+v\n", err.Error())

		return nil, err
	}

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		log.Warnf("getContractMetadata invalid json: %s", err)

		return nil, err
	}

	var unmarshalled map[string]interface{}
	_ = json.Unmarshal(responseBody, &unmarshalled)

	var decoded *manifold.DataResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decoded); err != nil {
		log.Errorf("❌  decode error: %s\n", err.Error())

		return nil, err
	}

	return decoded, nil
}

func getMintInfoWithURL(mintURL string) (*manifold.DataResponse, error) {
	// get word after last / in url
	slug := mintURL[strings.LastIndex(mintURL, "/")+1:]

	url := fmt.Sprintf("https://apps.api.manifoldxyz.dev/public/instance/data?appId=%d&instanceSlug=%s", manifoldMagicAppID, slug)

	log.Debugf("URL url: %s", url)

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

	var unmarshalled map[string]interface{}
	_ = json.Unmarshal(responseBody, &unmarshalled)

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

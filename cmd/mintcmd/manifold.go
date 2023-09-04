// inspired by the manifold-minter of @timfame-codespace, thanks!
// https://github.com/timfame-codespace/manifold-minter/

package mintcmd

import (
	"bytes"
	"context"
	"encoding/hex"
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
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagManifoldInstanceID int64

	manifoldMagicAppID = 2537426615

	fmtUnitEther = style.GrayStyle.Render("Ξ")
	fmtUnitGwei  = style.GrayStyle.Render("gwei")
	fmtUnitWei   = style.GrayStyle.Render("wei")
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
	manifoldCmd.Flags().Float64Var(&flagGasFeeCapMultiplier, "fee-cap", 1.5, "gas fee cap multiplier")
	_ = viper.BindPFlag("mint.fee_multiplier", manifoldCmd.Flags().Lookup("fee-cap"))
	manifoldCmd.Flags().Float64Var(&flagGasTipCapMultiplier, "tip-cap", 1.5, "gas tip cap multiplier")
	_ = viper.BindPFlag("mint.tip_multiplier", manifoldCmd.Flags().Lookup("tip-cap"))

	// number of wallets to use
	manifoldCmd.Flags().Uint16("num-wallets", 1, "number of wallets to use for minting")
	_ = viper.BindPFlag("mint.manifold.num-wallets", manifoldCmd.Flags().Lookup("num-wallets"))

	manifoldCmd.Flags().Uint16("amount-tx", 1, "number of tokens to mint per transaction")
	_ = viper.BindPFlag("mint.manifold.amount-tx", manifoldCmd.Flags().Lookup("amount-tx"))
	manifoldCmd.Flags().Uint16("amount-wallet", 1, "number of tokens to mint per wallet/key")
	_ = viper.BindPFlag("mint.manifold.amount-wallet", manifoldCmd.Flags().Lookup("amount-wallet"))

	// ! amount-wallet should be 1 for now !
	manifoldCmd.Flags().StringVar(&mintFor, "mint-for", "", "mint for delegated wallet")
	_ = viper.BindPFlag("mint.manifold.mint-for", manifoldCmd.Flags().Lookup("mint-for"))
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
		mintWallet.tag = lipgloss.NewStyle().Foreground(mintWallet.color).Render(style.ShortenAdressPTR(mintWallet.address))

		if mintFor != "" {
			log.Infof("minting for: %s", style.BoldAlmostWhite(mintFor))
			mintForAddress := common.HexToAddress(mintFor)
			mintWallet.mintFor = &mintForAddress
			// TODO change naming or introduce a new variable to avoid confusion
			mintWallet.address = &mintForAddress
		}

		mintWallet.mintIndices = make([]uint32, 0)
		mintWallet.merkleProofs = make([][][32]byte, 0)

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

	if viper.GetBool("dev.mode") {
		log.Infof("mintInfo:")
		pretty.Println(mintInfo)
	}

	// isPublicMint := mintInfo.PublicData.MerkleTreeID == 0

	// get the mint fee (once)
	lazyClaimERC1155, err := manifoldABIs.NewLazyClaimERC1155(internal.ManifoldLazyClaimERC1155, rpcClients.ToSlice()[rand.Intn(len(rpcClients.ToSlice()))]) //nolint:gosec
	if err != nil {
		log.Error(err)

		return
	}

	manifoldInstanceID = *big.NewInt(int64(mintInfo.PublicData.ClaimIndex))
	//
	// claimInfo
	claimInfo, err := lazyClaimERC1155.GetClaim(&bind.CallOpts{}, mintInfo.PublicData.CreatorContractAddress, &manifoldInstanceID)
	if err != nil {
		log.Errorf("❌ getClaim(…) failed: %s", style.BoldAlmostWhite(err.Error()))

		return
	}

	// MerkleTreeID (API) could be still set if merkleRoot is set to 0x0 via UpdateClaim(…)
	// empty merkleRoot = 0x0000000000000000000000000000000000000000000000000000000000000000
	isMerkleRootEmpty := claimInfo.MerkleRoot == [32]byte{}
	log.Printf("  isMerkleRootEmpty: %s", style.BoldAlmostWhite(fmt.Sprint(isMerkleRootEmpty)))

	isPublicMint := mintInfo.PublicData.MerkleTreeID == 0 || isMerkleRootEmpty

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

	if isPublicMint {
		log.Printf(style.LightGrayStyle.Copy().Italic(true).Render("   🕺💃  Public Mint  🕺💃  "))
	} else {
		log.Printf(style.LightGrayStyle.Copy().Italic(true).Render("   ⭐️  Exclusive Mint  ⭐️  "))
	}

	log.Print("")
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
	log.Printf("  price: %s%s", style.BoldAlmostWhite(fmt.Sprintf("%5.4f", mintInfo.MintPrice)), fmtUnitEther)

	log.Print("")
	log.Printf("  merkleTreeID: %s", style.BoldAlmostWhite(fmt.Sprintf("%d", mintInfo.PublicData.MerkleTreeID)))

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

	//
	// get manifold fee
	var manifoldFee *big.Int
	feeIndicator := ""

	if isPublicMint {
		manifoldFee, err = lazyClaimERC1155.MINTFEE(&bind.CallOpts{})
		if err != nil {
			log.Errorf("❌ getting mint fee failed: %s", style.BoldAlmostWhite(err.Error()))

			return
		}
	} else {
		feeIndicator = " (merkle)"

		manifoldFee, err = lazyClaimERC1155.MINTFEEMERKLE(&bind.CallOpts{})
		if err != nil {
			log.Errorf("❌ getting merkle mint fee failed: %s", style.BoldAlmostWhite(err.Error()))

			return
		}
	}

	log.Print("")
	log.Printf("  fee%s: %s", feeIndicator, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(manifoldFee).Ether())))

	if viper.GetBool("dev.mode") {
		log.Infof("claimInfo:")
		pretty.Println(claimInfo)
	}

	// MerkleTreeID (API) could be still set if merkleRoot is set to 0x0 via UpdateClaim(…)
	// empty merkleRoot = 0x0000000000000000000000000000000000000000000000000000000000000000
	log.Printf("  isMerkleRootEmpty: %s", style.BoldAlmostWhite(fmt.Sprint(isMerkleRootEmpty)))

	isPublicMint = mintInfo.PublicData.MerkleTreeID == 0 || isMerkleRootEmpty
	log.Printf("  isPublicMint: %s", style.BoldAlmostWhite(fmt.Sprint(isPublicMint)))

	startDate := time.Unix(claimInfo.StartDate.Int64(), 0)

	log.Print("")
	log.Printf("  mint start: %+v", style.BoldAlmostWhite(fmt.Sprint(startDate.Format("2006-01-02 15:04:05"))))
	log.Printf("            → in %+v", style.BoldAlmostWhite(fmt.Sprint(time.Until(startDate).Truncate(time.Second).String())))
	log.Print("")
	log.Printf("  minted: %+v / %v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.Total)), style.BoldAlmostWhite(fmt.Sprint(claimInfo.TotalMax)))
	log.Printf("  remaining: %+v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.TotalMax-claimInfo.Total)))
	log.Printf("  max/wallet: %+v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.WalletMax)))
	log.Printf("  merkleRoot: %+v", style.BoldAlmostWhite(fmt.Sprint(claimInfo.MerkleRoot)))

	totalMints, err := lazyClaimERC1155.GetTotalMints(&bind.CallOpts{}, *mintWallets.ToSlice()[0].address, mintInfo.PublicData.CreatorContractAddress, &manifoldInstanceID)
	if err != nil {
		log.Debugf("🤷‍♀️ getting total mints failed: %s", style.BoldAlmostWhite(err.Error()))
	} else {
		log.Printf("  totalMints: %d", totalMints)
	}

	log.Print("")

	//
	// acquire merkle proofs
	for _, mintWallet := range mintWallets.ToSlice() {
		// wg.Add(1)

		if isPublicMint {
			for i := uint16(0); i < amountPerTx; i++ {
				mintWallet.mintIndices = append(mintWallet.mintIndices, uint32(0))
				mintWallet.merkleProofs = append(mintWallet.merkleProofs, [][32]byte{claimInfo.MerkleRoot})
			}
		} else {
			merkleProofData, err := getMerkleProofFromManifoldForAddress(mintInfo.PublicData.MerkleTreeID, *mintWallet.address)
			if err != nil {
				continue
			}

			mintWallet.mintIndices, mintWallet.merkleProofs = getMerkleProofContractParamater(merkleProofData)

			if len(mintWallet.merkleProofs) == 0 {
				log.Errorf("%s | 🤷‍♀️ no merkle proofs for: %s", mintWallet.tag, style.BoldAlmostWhite(mintWallet.address.String()))

				continue
			}
			if len(mintWallet.mintIndices) == 0 {
				log.Errorf("%s | 🤷‍♀️ no mint indices found for: %s", mintWallet.tag, style.BoldAlmostWhite(mintWallet.address.String()))

				continue
			}
		}
	}

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

			if isPublicMint || len(mintWallet.merkleProofs) > 0 {
				mintERC1155(rpcEndpoints.Clone(), mintWallet, txsPerWallet, &manifoldInstanceID, mintInfo, claimInfo, manifoldFee)
			}
		}(mintWallet)
	}

	wg.Wait()

	log.Print("")
	log.Print("  🍹 all jobs finished! 🍹")
	log.Print("")
}

func getMerkleProofContractParamater(merkleProofData []Merkle) ([]uint32, [][][32]byte) {
	mintIndices := make([]uint32, 0)
	merkleProofs := make([][][32]byte, 0)

	for _, proof := range merkleProofData {
		merkleProof := make([][32]byte, 0)
		for _, hash := range proof.MerkleProof {
			// remove 0x from hash
			hash = strings.TrimPrefix(hash, "0x")

			byteSlice, err := hex.DecodeString(hash)
			if err != nil {
				log.Errorf("❌ hex.DecodeString(…) failed: %s", style.BoldAlmostWhite(err.Error()))
				log.Fatal(err)
			}
			var byteArr [32]byte
			copy(byteArr[:], byteSlice)

			merkleProof = append(merkleProof, byteArr)
		}
		merkleProofs = append(merkleProofs, merkleProof)
		mintIndices = append(mintIndices, uint32(proof.Value))
	}

	// log.Printf(" converted mint indices: %v", mintIndices)
	// log.Printf(" converted merkle proofs: %v", merkleProofs)

	return mintIndices, merkleProofs
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

		log.Printf("%s | current nonce: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(nonce)))
		log.Print("")

		// get the current gas price
		gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("%s | ❌ getting gasPrice failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		// get the current gas tip
		gasTip, err := rpcClient.SuggestGasTipCap(context.Background())
		if err != nil {
			log.Errorf("%s | ❌ getting gasTip failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))
			prErr(mintInfo, claimInfo)

			continue
		}

		log.Printf("%s | ⛽️ suggested gasFeeCap: %s%s | %d%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%4.1f", utils.WeiToGwei(gasPrice))), fmtUnitGwei, gasPrice, fmtUnitWei)
		log.Printf("%s | ⛽️ suggested gasTipCap: %s%s | %d%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%4.1f", utils.WeiToGwei(gasTip))), fmtUnitGwei, gasTip, fmtUnitWei)
		log.Print("")

		//
		// apply gas multiplier
		feeCapMultiplier := new(big.Float).SetFloat64(viper.GetFloat64("mint.fee_multiplier"))
		tipCapMultiplier := new(big.Float).SetFloat64(viper.GetFloat64("mint.tip_multiplier"))

		suggestedFee := new(big.Float).SetInt(gasPrice)
		suggestedTip := new(big.Float).SetInt(gasTip)

		gasFeeCapWei, _ := new(big.Float).Mul(suggestedFee, feeCapMultiplier).Int(nil)
		gasTipCapWei, _ := new(big.Float).Mul(suggestedTip, tipCapMultiplier).Int(nil)

		log.Printf("%s | ⛽️ your gasFeeCap: %s%s | %d%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%4.1f", utils.WeiToGwei(gasFeeCapWei))), fmtUnitGwei, gasFeeCapWei, fmtUnitWei)
		log.Printf("%s | ⛽️ your gasTipCap: %s%s | %d%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%4.1f", utils.WeiToGwei(gasTipCapWei))), fmtUnitGwei, gasTipCapWei, fmtUnitWei)
		log.Print("")

		// 💸 💸 💸
		mintPriceWei := ToWei(mintInfo.MintPrice, 18)
		log.Printf("%s | mintPriceWei: %s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%+v", mintPriceWei)))

		// TODO WARNING: this conversion is lossy
		// mintCostFloat := utils.EtherToWeiFloat(big.NewFloat(mintInfo.MintPrice))
		//  log.Printf("%s | mintCostFloat: %s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%+v", mintCostFloat)))
		//  mintCost := utils.EtherToWei(big.NewFloat(mintInfo.MintPrice))

		totalCost := new(big.Int).Add(manifoldFee, mintPriceWei)

		// print big.Int
		log.Printf("%s | totalCost: %s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%+v", totalCost)))

		//
		// create the transaction options
		txOpts, err := bind.NewKeyedTransactorWithChainID(mintWallet.privateKey, big.NewInt(1))
		if err != nil {
			prErr(mintInfo, claimInfo)
			log.Errorf("%s | ❌ creating transaction options failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Printf("%s |  mint cost (1x): %s%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(mintPriceWei).Ether())), fmtUnitEther)
		log.Printf("%s |        fee (1x): %s%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(manifoldFee).Ether())), fmtUnitEther)
		log.Printf("%s |      total (1x): %s%s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(totalCost).Ether())), fmtUnitEther)
		log.Print("")

		txOpts.From = crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey)
		txOpts.Nonce = big.NewInt(int64(nonce))
		txOpts.Value = totalCost
		txOpts.GasFeeCap = gasFeeCapWei
		txOpts.GasTipCap = gasTipCapWei

		if viper.GetBool("dev.mode") {
			txOpts.NoSend = true
			log.Printf("%s | txOpts.NoSend: %t", mintWallet.tag, txOpts.NoSend)
		}

		// create the transactions
		var sentTx *types.Transaction

		if amountPerTx := viper.GetUint16("mint.manifold.amount-tx"); amountPerTx > 1 {
			sumTotalCost := new(big.Int).Mul(totalCost, big.NewInt(int64(amountPerTx)))
			txOpts.Value = sumTotalCost

			log.Printf("%s | txOpts: %#v", mintWallet.tag, txOpts)
			log.Print("")

			sentTx, err = lazyClaimERC1155.MintBatch(txOpts, mintInfo.PublicData.CreatorContractAddress, manifoldInstanceID, amountPerTx, mintWallet.mintIndices, mintWallet.merkleProofs, *mintWallet.address)
			if err != nil {
				prErr(mintInfo, claimInfo)

				log.Print("")
				log.Printf("%s | ❌ creating batch transaction failed: %+v | %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()), err)
				log.Print("")

				continue
			}
		} else {
			var mintIndex uint32
			if len(mintWallet.mintIndices) > 0 {
				mintIndex = mintWallet.mintIndices[0]
			}

			log.Printf("%s | txOpts: %#v", mintWallet.tag, txOpts)
			log.Print("")

			sentTx, err = lazyClaimERC1155.Mint(txOpts, mintInfo.PublicData.CreatorContractAddress, manifoldInstanceID, mintIndex, mintWallet.merkleProofs[0], *mintWallet.address)
			if err != nil {
				prErr(mintInfo, claimInfo)

				log.Print("")
				log.Printf("%s | ❌ creating transaction failed: %+v | %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()), err)
				log.Print("")

				continue
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

type Merkle struct {
	MerkleProof []string `json:"merkleProof"`
	Value       int      `json:"value"`
}

func getMerkleProofFromManifoldForAddress(merkleTreeID int, address common.Address) ([]Merkle, error) {
	// https://apps.api.manifoldxyz.dev/public/merkleTree/1068163731/merkleInfo?address=0x9654F22b9dEBac18396b4815C138A450786a7045
	url := fmt.Sprintf("https://apps.api.manifoldxyz.dev/public/merkleTree/%d/merkleInfo?address=%s", merkleTreeID, address.Hex())

	// log.Printf("Merkle proof url: %s", url)

	response, err := utils.HTTP.GetWithTLS12AndHeader(context.TODO(), url, createCorrectHeadersForManifoldAPI())
	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("⌛️ Merkle proof GetMerkleProofFromManifoldForAddress · timeout while fetching: %+v\n", err.Error())
		} else {
			log.Errorf("❌ Merkle proof GetMerkleProofFromManifoldForAddress · error: %+v\n", err.Error())
		}

		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		log.Errorf("❌ Merkle proof GetMerkleProofFromManifoldForAddress · error: %+v\n", response.Status)

		return nil, err
	}
	defer response.Body.Close()

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

	merkleData := []Merkle{}
	_ = json.Unmarshal(responseBody, &merkleData)

	// log.Printf("Merkle data: %+v", merkleData)

	// var decoded *manifold.Merkle
	// if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decoded); err != nil {
	//	log.Errorf("❌  decode error: %s\n", err.Error())
	//
	//	return nil, err
	//}

	return merkleData, nil
}

func getMintInfoWithInstanceID(identifier int64) (*manifold.DataResponse, error) {
	url := fmt.Sprintf("https://apps.api.manifoldxyz.dev/public/instance/data?id=%d", identifier)

	log.Debugf("Identifier url: %s", url)

	header := createCorrectHeadersForManifoldAPI()

	response, err := utils.HTTP.GetWithTLS12AndHeader(context.TODO(), url, header)
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

func createCorrectHeadersForManifoldAPI() http.Header {
	// convert curl 'https://apps.api.manifoldxyz.dev/public/instance/data?id=72282352' -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/116.0' -H 'Accept: application/json' -H 'Accept-Language: de,en-US;q=0.7,en;q=0.3' -H 'Accept-Encoding: gzip, deflate, br' -H 'Referer: https://app.manifold.xyz/' -H 'Content-Type: application/json' -H 'Origin: https://app.manifold.xyz' -H 'DNT: 1' -H 'Connection: keep-alive' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: cross-site' -H 'Sec-GPC: 1' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' -H 'TE: trailers'
	header := http.Header{}
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/116.0")
	header.Set("Accept", "application/json")
	header.Set("Accept-Language", "de,en-US;q=0.7,en;q=0.3")
	header.Set("Accept-Encoding", "gzip, deflate, br")
	header.Set("Referer", "https://app.manifold.xyz/")
	header.Set("Content-Type", "application/json")
	header.Set("Origin", "https://app.manifold.xyz")
	header.Set("DNT", "1")
	header.Set("Connection", "keep-alive")
	header.Set("Sec-Fetch-Dest", "empty")
	header.Set("Sec-Fetch-Mode", "cors")
	header.Set("Sec-Fetch-Site", "cross-site")
	header.Set("Sec-GPC", "1")
	header.Set("Pragma", "no-cache")
	header.Set("Cache-Control", "no-cache")
	header.Set("TE", "trailers")

	return header
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

// ToWei decimals to wei ,from: https://goethereumbook.org/util-go/
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

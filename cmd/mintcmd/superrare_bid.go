package mintcmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/abis/superrare"
	"github.com/benleb/gloomberg/internal/gbl"
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
	flagOriginContract  string
	flagCurrencyAddress string

	flagAmoutEth float64

	marketAddresses = []string{"0x41a322b28d0ff354040e2cbc676f0320d8c8850d", "0x65b49f7aee40347f5a90b714be4ef086f3fe5e2c", "0x2947f98c42597966a0ec25e92843c09ac17fbaa7", "0x65b49f7aee40347f5a90b714be4ef086f3fe5e2c", "0x2947f98c42597966a0ec25e92843c09ac17fbaa7", "0x6d7c44773c52d396f43c2d511b81aa168e9a7a42", "0x6d7c44773c52d396f43c2d511b81aa168e9a7a42"}
)

// SuperRareCmd represents the superrare command.
var SuperRareCmd = &cobra.Command{
	Use:     "superrare-bid",
	Aliases: []string{"sr-bid"},
	Short:   "Place a bid on a SuperRare auction.",
	// 	Long: fmt.Sprintf(`%s

	//   Mints the token from the given SuperRare URL %s Identifier with the configured wallets.`, style.Bold("or"), style.GetSmallHeader(internal.GloombergVersion)),
	Run: bidSuperRare,
}

func init() {
	// MintCmd.AddCommand(SuperRareCmd)

	// superrare url or identifier
	SuperRareCmd.Flags().StringVar(&flagOriginContract, "origin", "", "origin (collection) contract")
	SuperRareCmd.Flags().StringVar(&flagCurrencyAddress, "currency", "0x0000000000000000000000000000000000000000", "origin (collection) contract")

	SuperRareCmd.Flags().Uint64("token-id", 0, "tokenID to bid on")
	_ = viper.BindPFlag("mint.superrare.token-id", SuperRareCmd.Flags().Lookup("token-id"))

	SuperRareCmd.Flags().Float64Var(&flagAmoutEth, "amount-eth", 0.0, "amount of eth to bid")
	_ = viper.BindPFlag("mint.superrare.amount-eth", SuperRareCmd.Flags().Lookup("amount-eth"))

	// do not wait for auction start
	SuperRareCmd.Flags().Bool("no-wait", false, "do not wait for auction start")
	_ = viper.BindPFlag("mint.superrare.nowait", SuperRareCmd.Flags().Lookup("no-wait"))

	// private keys/wallets to mint with
	SuperRareCmd.Flags().StringSliceVarP(&flagPrivateKeys, "keys", "p", make([]string, 0), "private keys/wallets to mint with")
	_ = viper.BindPFlag("mint.keys", SuperRareCmd.Flags().Lookup("keys"))

	// rpc endpoints to use
	SuperRareCmd.Flags().StringSliceVarP(&flagRPCs, "rpcs", "r", make([]string, 0), "rpc endpoints to mint with (randomly chosen)")
	_ = viper.BindPFlag("mint.rpcs", SuperRareCmd.Flags().Lookup("rpcs"))

	// gas settings
	SuperRareCmd.Flags().Float64Var(&flagGasFeeCapMultiplier, "fee-cap", 1.0, "gas fee cap multiplier")
	_ = viper.BindPFlag("mint.fee_multiplier", SuperRareCmd.Flags().Lookup("fee-cap"))
	SuperRareCmd.Flags().Float64Var(&flagGasTipCapMultiplier, "tip-cap", 1.0, "gas tip cap multiplier")
	_ = viper.BindPFlag("mint.tip_multiplier", SuperRareCmd.Flags().Lookup("tip-cap"))

	// number of wallets to use
	SuperRareCmd.Flags().Uint16("num-wallets", 3, "number of wallets to use for minting")
	_ = viper.BindPFlag("mint.superrare.num-wallets", SuperRareCmd.Flags().Lookup("num-wallets"))
}

func bidSuperRare(_ *cobra.Command, _ []string) {
	// rpcClients := make([]*ethclient.Client, 0)
	// availableWallets := make([]*MintWallet, 0)
	rpcClients := mapset.NewSet[*ethclient.Client]()
	rpcEndpoints := mapset.NewSet[string]()

	availableWallets := make([]*MintWallet, 0)
	mintWallets := mapset.NewSet[*MintWallet]()

	// negate for easier handling...
	waitForStart := !viper.GetBool("mint.superrare.nowait")

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

		mintWallet.color = style.GenerateColorWithSeed(mintWallet.address.Big().Int64())
		mintWallet.tag = lipgloss.NewStyle().Foreground(mintWallet.color).Render(style.ShortenAdressPTR(mintWallet.address))

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

	// log.Printf("  endpoints: %s", strings.Join(rpcEndpoints.ToSlice(), ", "))

	if availableWallets == nil {
		log.Fatalf("❌ no valid signer keys found")

		return
	}

	tokenInfo, err := GetAuctionInfo(common.HexToAddress(flagOriginContract), viper.GetUint64("mint.superrare.token-id"), *availableWallets[0].address) //nolint:gosec
	if err != nil {
		log.Fatalf("❌ getting auction info failed: %v", err)

		return
	}

	gbl.Log.Info(pretty.Sprint(tokenInfo))
	// pretty.Println(tokenInfo)

	srURL := fmt.Sprintf("https://superrare.com/%s/%d", flagOriginContract, viper.GetUint64("mint.superrare.token-id"))

	log.Print("")
	log.Print("")
	log.Printf("  %s  (by %s)", style.TerminalLink(srURL, style.BoldAlmostWhite(tokenInfo.Metadata.Name)), style.BoldAlmostWhite(tokenInfo.Owner.Username))
	log.Print("")

	// 	log.Debugf("mint info: %#v", auctionInfo)

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
	numWallets := uint16(math.Min(float64(viper.GetUint16("mint.superrare.num-wallets")), float64(len(availableWallets))))

	// select the first numWallets from our available wallets
	for _, wallet := range availableWallets[:numWallets] {
		mintWallets.Add(wallet)
	}

	fmtWallets = make([]string, 0)
	for _, wallet := range mintWallets.ToSlice() {
		fmtWallets = append(fmtWallets, style.Bold(wallet.tag))
	}

	log.Printf("  mint wallets: %s", strings.Join(fmtWallets, ", "))

	// 	amountPerTx := viper.GetUint16("mint.superrare.amount-tx")
	// 	amountPerWallet := uint16(math.Max(float64(amountPerTx), float64(viper.GetUint16("mint.superrare.amount-wallet"))))
	// 	txsPerWallet := amountPerWallet / amountPerTx
	// 	totalTxs := txsPerWallet * uint16(mintWallets.Cardinality())

	// 	log.Print("")
	// 	log.Printf("  amount per tx: %s", style.BoldAlmostWhite(fmt.Sprint(amountPerTx)))
	// 	log.Printf("  amount per wallet: %s", style.BoldAlmostWhite(fmt.Sprint(amountPerWallet)))

	// 	log.Print("")
	// 	log.Printf("  → txs per wallet: %s", style.BoldAlmostWhite(fmt.Sprint(txsPerWallet)))
	// 	log.Printf("  → total txs: %s", style.BoldAlmostWhite(fmt.Sprint(totalTxs)))

	//
	// superrare api info
	//

	log.Print("")
	log.Print("")
	log.Print(style.BoldAlmostWhite("superrare info") + " (from api)")

	// 	log.Print("")
	// 	log.Printf("  price: %s", style.BoldAlmostWhite(fmt.Sprintf("%5.4f", auctionInfo.MintPrice)))

	// 	log.Print("")
	// 	log.Printf("  merkeleTreeID: %s", style.BoldAlmostWhite(fmt.Sprintf("%d", auctionInfo.PublicData.MerkleTreeID)))

	creatorContract := common.HexToAddress(tokenInfo.ContractAddress)
	log.Print("")
	log.Printf("  collection/creator contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL((&creatorContract)), style.BoldAlmostWhite(creatorContract.String())))

	// 	if auctionInfo.PublicData.ExtensionContractAddress != internal.SuperRareLazyClaimERC1155 {
	// 		log.Printf("  superrare/extension contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&auctionInfo.PublicData.ExtensionContractAddress), style.BoldAlmostWhite(auctionInfo.PublicData.ExtensionContractAddress.Hex())))
	// 		log.Printf("  superrare lazy claim erc1155 contract: %s", style.TerminalLink(utils.GetEtherscanTokenURL(&internal.SuperRareLazyClaimERC1155), style.BoldAlmostWhite(internal.SuperRareLazyClaimERC1155.Hex())))
	// 	}

	// 	//
	// 	// superrare chain info
	// 	//

	log.Print("")
	log.Printf("  auction start: %+v", style.BoldAlmostWhite(tokenInfo.Auction.StartingTime.Format("15:04:05")))
	log.Printf("            → in %+v", style.BoldAlmostWhite(fmt.Sprint(time.Until(tokenInfo.Auction.StartingTime).Truncate(time.Second).String())))
	log.Print("")

	if tokenInfo.Auction.StartingTime.After(time.Now()) && waitForStart {
		log.Print("")
		log.Print("")
		log.Printf(" 💤 💤 💤  waiting for auction start in %s  💤 💤 💤", style.BoldAlmostWhite(fmt.Sprint(time.Until(tokenInfo.Auction.StartingTime).Truncate(time.Second).String())))
		log.Print("")
		log.Printf(style.GrayStyle.Render("    (use --no-wait to skip waiting)"))

		// sleep until 5s before startDate
		time.Sleep(time.Until(tokenInfo.Auction.StartingTime.Add(-2 * time.Second)))
	} else {
		log.Print("")
		log.Printf(" 💤 🏃 💤  waiting %s  💤 🏃‍♂️ 💤", style.BoldAlmostWhite("skipped!"))
		log.Printf(style.GrayStyle.Render(fmt.Sprintf("       %s flag is set", style.LightGrayStyle.Render("--no-wait"))))
	}

	log.Print("")
	log.Print("")

	// lfg

	log.Print("")
	log.Print("")
	log.Print(" 🚀 🚀 🚀  starting bid jobs  🚀 🚀 🚀")
	log.Print("")
	log.Print("")

	wg := sync.WaitGroup{}

	for _, mintWallet := range mintWallets.ToSlice() {
		wg.Add(1)

		endpoints := rpcEndpoints.Clone().ToSlice()

		go func(mintWallet MintWallet) {
			log.Printf("  %s | 🚀 starting bid job...", mintWallet.tag)

			for _, rpcEndpoint := range endpoints {
				// string to int
				// https://stackoverflow.com/a/44783357/12347616
				colorSeed := binary.BigEndian.Uint64([]byte(rpcEndpoint))
				rpcStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(int64(colorSeed)))

				wallet := mintWallet
				wallet.tag = mintWallet.tag + "|" + rpcStyle.Render("rpc")

				log.Printf("  %s | ...with rpc: %+v", wallet.tag, style.BoldAlmostWhite(rpcEndpoint))

				go func(rpcEndpoint string) {
					defer wg.Done()
					placeBid(rpcEndpoint, &wallet, big.NewInt(0), common.HexToAddress(flagOriginContract), common.HexToAddress(flagCurrencyAddress), &tokenInfo.Auction)
				}(rpcEndpoint)
			}
		}(*mintWallet)
	}

	wg.Wait()

	log.Print("")
	log.Print("  🍹 all jobs finished! 🍹")
	log.Print("")
}

func placeBid(rpcEndpoint string, mintWallet *MintWallet, bidAmount *big.Int, tokenContract common.Address, currencyAddress common.Address, auctionInfo *AuctionInfo) {
	log.Debugf("rpcEndpoint: %+v | mintWallet: %+v bidAmount: %+v | tokenContract: %+v | currencyAddress: %+v | auctionInfo: %+v", rpcEndpoint, mintWallet, bidAmount, tokenContract, currencyAddress, auctionInfo)

	txConfirmed := 0

	log.Print("\n\n")

	for {
		// // choose random rpc endpoint
		// rpcIdx := rand.Intn(len(rpcEndpoints.ToSlice())) //nolint:gosec
		// rpcEndpoint := rpcEndpoints.ToSlice()[rpcIdx]

		// log.Debugf("%s | 📑 rpc endpoints (%d): %s", mintWallet.tag, rpcEndpoints.Cardinality(), style.BoldAlmostWhite(fmt.Sprintf("%+v", rpcEndpoints)))
		// log.Debugf("%s | 🎯 selected rpc endpoint: %s", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(rpcIdx)))

		// connect to rpc endpoint
		rpcClient, err := ethclient.Dial(rpcEndpoint)
		if err != nil {
			log.Errorf("❌ failed to connect to rpc endpoint %s: %v", rpcEndpoint, err)

			continue
		}

		// contract binding
		log.Printf("%s | 🧶 create contract binding...", mintWallet.tag)

		marketplace, err := superrare.NewMarketplace(common.HexToAddress(auctionInfo.AuctionContractAddress), rpcClient)
		if err != nil {
			log.Errorf("❌ binding contract abi failed: %+v", style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Debugf("%s | marketplace: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(marketplace)))

		// get the nonce
		nonce, err := rpcClient.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(mintWallet.privateKey.PublicKey))
		if err != nil {
			log.Errorf("%s | ❌ getting nonce failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Printf("%s | nonce: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(nonce)))

		// get the current gas price
		gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
		if err != nil {
			log.Errorf("%s | ❌ getting gasPrice failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		log.Printf("%s | ⛽️ gasPrice: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprint(gasPrice)))

		// get the current gas tip
		gasTip, err := rpcClient.SuggestGasTipCap(context.Background())
		if err != nil {
			log.Errorf("%s | ❌ getting gasTip failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

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
		// mintCost := utils.EtherToWei(big.NewFloat(mintInfo.MintPrice))
		totalCost := utils.EtherToWei(new(big.Float).SetInt(bidAmount))

		// totalCost = big.NewInt(100000000000000000)

		//
		// create the transaction options
		txOpts, err := bind.NewKeyedTransactorWithChainID(mintWallet.privateKey, big.NewInt(1))
		if err != nil {
			log.Errorf("%s | ❌ creating transaction options failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		// log.Printf("%s | mintCost: %+v", mintWallet.tag, style.BoldAlmostWhite(fmt.Sprintf("%7.5f", price.NewPrice(mintCost).Ether())))
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

		sentTx, err = marketplace.Bid(txOpts, tokenContract, big.NewInt(int64(auctionInfo.TokenID)), currencyAddress, totalCost)
		if err != nil {
			log.Printf("%s | ❌ executing transaction failed: %+v | %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()), err)
		}

		if sentTx == nil {
			log.Printf("%s | ❌ executing transaction failed - sentTx is %+v", mintWallet.tag, sentTx)

			continue
		}

		log.Printf("%s | 🙌 tx sent! → %s 🙌", mintWallet.tag, style.TerminalLink(utils.GetEtherscanTxURL(sentTx.Hash().Hex()), style.BoldAlmostWhite(sentTx.Hash().Hex())))

		// wait for the transaction to be mined
		receipt, err := bind.WaitMined(context.Background(), rpcClient, sentTx)
		if err != nil {
			log.Printf("%s | ❌ transaction failed: %+v", mintWallet.tag, style.BoldAlmostWhite(err.Error()))

			continue
		}

		if receipt != nil {
			log.Printf("%s | 🎉 transaction mined! → %s 🎉", mintWallet.tag, style.TerminalLink(utils.GetEtherscanTxURL(receipt.TxHash.Hex()), style.BoldAlmostWhite(receipt.TxHash.Hex())))
			pretty.Println(receipt)

			txConfirmed++
		}

		// if txConfirmed >= int(txsPerWallet) {
		// 	log.Printf("%s | 💥 🙌 💥 all txs confirmed! 🍹 🥃 🥂s", mintWallet.tag)

		// 	return
		// }

		time.Sleep(time.Millisecond * 337)
	}
}

func GetAuctionInfo(contractAddress common.Address, tokenID uint64, userAddress common.Address) (*TokenInfo, error) {
	url := "https://superrare.com/api/v2/nft/get"

	// header with user agent
	srHeader := http.Header{}
	srHeader.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	srHeader.Add("Accept", "application/json, text/plain, */*")
	srHeader.Add("Cache-Control", "no-cache")
	srHeader.Add("Pragma", "no-cache")
	srHeader.Add("DNT", "1")
	srHeader.Add("Content-Type", "application/json")

	// create the request
	srRequest := TokenInfoRequest{
		ContractAddress: strings.ToLower(contractAddress.String()),
		TokenID:         tokenID,
		UserAddress:     strings.ToLower(userAddress.String()),
		MarketAddresses: marketAddresses,
		Fingerprint:     nil,
	}

	requestBody, err := json.Marshal(srRequest)
	if err != nil {
		return nil, err
	}

	response, err := utils.HTTP.PostWithHeader(context.TODO(), url, srHeader, strings.NewReader(string(requestBody)))
	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("⌛️ auction info · timeout while fetching: %+v\n", err)
		} else {
			log.Errorf("❌ Iauction info · error: %+v\n", err)
		}

		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Errorf("❌ auction info · status error: %d - %+v\n", response.StatusCode, response.Status)

		return nil, err
	}

	// read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("❌ auction info · response read error: %+v\n", err.Error())

		return nil, err
	}

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		log.Warnf("getContractMetadata invalid json: %s", err)

		return nil, err
	}

	var unmarshalled map[string]interface{}
	_ = json.Unmarshal(responseBody, &unmarshalled)

	var decoded *TokenInfoResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decoded); err != nil {
		log.Errorf("❌ auction info decode error: %s\n", err.Error())

		return nil, err
	}

	if decoded == nil || strings.ToLower(decoded.Status) != "success" {
		log.Errorf("❌ auction info · error: %+v\n", decoded.Status)

		return nil, err
	}

	return &decoded.Result, nil
}

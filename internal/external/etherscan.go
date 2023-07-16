package external

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GasOracleResponse struct {
	Response
	Result GasOracle `json:"result"`
}

type GasOracle struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

type TokenBalancesResponse struct {
	Response
	Result string `json:"result"`
}

type AccountBalancesResponse struct {
	Response
	Result []struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
	} `json:"result"`
}

type AccountBalance struct {
	Account         string   `json:"account"`
	BalanceETH      *big.Int `json:"balance"`
	BalanceWETH     *big.Int `json:"balance_weth"`
	BalanceBlurPool *big.Int `json:"balance_blurpool"`
}

type Token string

const apiBaseURL = "https://api.etherscan.io/api"

var ErrInvalidJSON = errors.New("invalid json")

func GetEstimatedGasPrice() *big.Int {
	var estimatedGasPrice *big.Int

	if gasOracle := GetGasOracle(); gasOracle != nil {
		gasPrice, err := strconv.ParseInt(gasOracle.ProposeGasPrice, 10, 64)
		if err != nil {
			gbl.Log.Infof("could not parse proposedGasPrice: %+v | %s", gasPrice, err)

			return nil
		}

		estimatedGasPrice = big.NewInt(gasPrice)
		gbl.Log.Infof("updated proposed gas price to %d gwei", estimatedGasPrice)

		return estimatedGasPrice
	}

	gbl.Log.Info("updated current gas price failed")

	return nil
}

func GetGasOracle() *GasOracle {
	if !viper.IsSet("api_keys.etherscan") {
		log.Fatal("api_keys.etherscan not set")
	}

	url := withAPIKey(fmt.Sprint(apiBaseURL + "?module=gastracker&action=gasoracle"))

	// // client, _ := createEtherscanHTTPClient()
	// client, _ := utils.DefaultHTTPClient()

	// gbl.Log.Debugf("gas oracle url: %s", url)

	// request, _ := http.NewRequest("GET", url, nil)

	// response, err := client.Do(request)
	response, err := utils.HTTP.GetWithTLS12(context.Background(), url)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching current gas: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ gas oracle error: %+v", err.Error())
		}

		return nil
	}

	gbl.Log.Debugf("gas oracle status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var gasOracleResponse *GasOracleResponse

	responseBody, _ := io.ReadAll(response.Body)

	// decode the data
	if !json.Valid(responseBody) {
		gbl.Log.Warnf("gas oracle invalid json: %s", err)

		return nil
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&gasOracleResponse); err != nil {
		gbl.Log.Warnf("gas oracle decode error: %s", err.Error())

		return nil
	}

	gasOracle := gasOracleResponse.Result

	return &gasOracle
}

func GetBalances(wallets *wallet.Wallets) ([]*AccountBalance, error) {
	balances := MultiAccountBalance(wallets)

	for _, balance := range balances {
		wethBalance, err := GetWETHBalance(common.HexToAddress(balance.Account))
		if err != nil || wethBalance == nil {
			gbl.Log.Warnf("could not get weth balance for %s: %s", balance.Account, err.Error())

			continue
		}

		balance.BalanceWETH = wethBalance

		// blur pool
		blurPoolBalance, err := GetBlurPoolBalance(common.HexToAddress(balance.Account))
		if err != nil || blurPoolBalance == nil {
			gbl.Log.Warnf("could not get blur pool balance for %s: %s", balance.Account, err.Error())

			continue
		}

		balance.BalanceBlurPool = blurPoolBalance

		// throttle to avoid hitting the apis reqs/s limit
		time.Sleep(time.Millisecond * 173)
	}

	return balances, nil
}

func MultiAccountBalance(wallets *wallet.Wallets) []*AccountBalance {
	balances := make([]*AccountBalance, 0)

	if !viper.IsSet("api_keys.etherscan") {
		gbl.Log.Warnf("api_keys.etherscan not set")

		return nil
	}

	addressList := strings.Join(wallets.StringAddresses(), ",")
	url := withAPIKey(fmt.Sprint(apiBaseURL+"?module=account&action=balancemulti&tag=latest&address=", addressList))

	gbl.Log.Debugf("multiAccountBalance url: %s", url)

	// client, _ := createEtherscanHTTPClient()
	// client, _ := utils.DefaultHTTPClient()
	// request, _ := http.NewRequest("GET", url, nil)

	// response, err := client.Do(request)
	response, err := utils.HTTP.GetWithTLS12(context.Background(), url)
	if err != nil {
		gbl.Log.Warnf("multiAccountBalance error: %+v", err.Error())

		return nil
	}

	gbl.Log.Debugf("multiAccountBalance status: %s", response.Status)

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	// validate the data
	if err != nil || !json.Valid(responseBody) {
		gbl.Log.Warnf("multiAccountBalance invalid json: %s", err)

		return nil
	}

	// create a variable of the same type as our model
	var accountBalancesResponse *AccountBalancesResponse

	// decodeHooks := mapstructure.ComposeDecodeHookFunc(
	// 	hooks.StringToAddressHookFunc(),
	// 	hooks.StringToBigIntHookFunc(),
	// )

	// decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
	// 	DecodeHook:       decodeHooks,
	// 	Result:           &accountBalancesResponse,
	// 	WeaklyTypedInput: true,
	// })

	// decode the data
	dec := json.NewDecoder(bytes.NewReader(responseBody))

	if err := dec.Decode(&accountBalancesResponse); err != nil {
		gbl.Log.Warnf("multiAccountBalance decode error: %s", err.Error())

		return nil
	}

	for _, accountBalance := range accountBalancesResponse.Result {
		balance := new(big.Int)

		_, err := fmt.Sscan(accountBalance.Balance, balance)
		if err != nil {
			gbl.Log.Warnf("could not parse balance for %s: %s", accountBalance.Account, err.Error())
		} else {
			gbl.Log.Debugf("%s balance: %+v", accountBalance.Account, balance)

			balances = append(balances, &AccountBalance{
				Account:     accountBalance.Account,
				BalanceETH:  balance,
				BalanceWETH: big.NewInt(0),
			})
		}
	}

	return balances
}

// func createEtherscanHTTPClient() (*http.Client, error) {
// 	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

// 	transport := &http.Transport{
// 		TLSClientConfig:     tlsConfig,
// 		MaxIdleConnsPerHost: 20,
// 		IdleConnTimeout:     13 * time.Second,
// 	}

// 	// explicitly use http2
// 	_ = http2.ConfigureTransport(transport)

// 	client := &http.Client{
// 		Timeout:   13 * time.Second,
// 		Transport: transport,
// 	}

// 	return client, nil
// }

func GetWETHBalance(walletAddress common.Address) (*big.Int, error) {
	return GetTokenBalance(walletAddress, internal.WETHContractAddress)
}

func GetBlurPoolBalance(walletAddress common.Address) (*big.Int, error) {
	return GetTokenBalance(walletAddress, internal.BlurPoolTokenContractAddress)
}

func GetTokenBalance(walletAddress common.Address, tokenAddress common.Address) (*big.Int, error) {
	// etherscan api access required
	if !viper.IsSet("api_keys.etherscan") {
		gbl.Log.Fatal("api_keys.etherscan not set")
	}

	url := withAPIKey(fmt.Sprintf(
		apiBaseURL+"?module=account&action=tokenbalance&contractaddress=%s&address=%s&tag=latest",
		tokenAddress, walletAddress,
	))

	// // fetch balance
	// request, _ := http.NewRequest("GET", url, nil)
	// // client, _ := createEtherscanHTTPClient()
	// client, _ := utils.DefaultHTTPClient()

	// response, err := client.Do(request)

	response, err := utils.HTTP.GetWithTLS12(context.Background(), url)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ token balance · timeout while fetching: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ token balance · error: %+v", err.Error())
		}

		return nil, err
	}
	defer response.Body.Close()

	// read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		gbl.Log.Errorf("❌ token balance · response read error: %+v", err.Error())

		return nil, err
	}

	// decode the data
	if !json.Valid(responseBody) {
		gbl.Log.Warnf("token balance · invalid json")

		return nil, ErrInvalidJSON
	}

	// decode the data
	var tokenBalanceResponse *TokenBalancesResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&tokenBalanceResponse); err != nil {
		gbl.Log.Warnf("token balance · decode error: %s", err.Error())

		return nil, err
	}

	balance := new(big.Int)

	_, err = fmt.Sscan(tokenBalanceResponse.Result, balance)
	if err != nil {
		gbl.Log.Warnf("could not parse token balance for %s: %s", walletAddress, err.Error())

		return nil, err
	}

	return balance, nil
}

func withAPIKey(url string) string {
	return url + "&apikey=" + viper.GetString("api_keys.etherscan")
}

func GetFirstTransactionsByContract(numTxs int64, contractAddress common.Address) ([]Transaction, error) {
	if !viper.IsSet("api_keys.etherscan") {
		log.Fatal("api_keys.etherscan not set")
	}

	url := withAPIKey(fmt.Sprintf("%s?module=account&action=tokennfttx&contractaddress=%s&page=1&offset=%d&startblock=0&endblock=99999999&sort=asc", apiBaseURL, contractAddress, numTxs))

	response, err := utils.HTTP.GetWithTLS12(context.Background(), url)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching current gas: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ first 1k txs error: %+v", err.Error())
		}

		return nil, err
	}

	// gbl.Log.Infof("fetched first 1k txs for %s: %s", contractAddress.Hex(), response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var firstTransactions *TransactionsResponse

	responseBody, _ := io.ReadAll(response.Body)

	// decode the data
	if !json.Valid(responseBody) {
		gbl.Log.Warnf("txs response invalid json: %s", err)

		return nil, ErrInvalidJSON
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&firstTransactions); err != nil {
		gbl.Log.Warnf("txs response decode error: %s | %+v", err.Error(), string(responseBody))

		return nil, err
	}

	// log.Printf("first txs: %+v", firstTransactions.Result)
	// log.Printf("first txs: %+v", firstTransactions.Result[0])
	// log.Printf("num first txs: %d", len(firstTransactions.Result))

	if len(firstTransactions.Result) < int(numTxs) {
		gbl.Log.Debugf("only %d txs found for %s (requested %d)", len(firstTransactions.Result), contractAddress.Hex(), numTxs)

		return nil, fmt.Errorf("only %d txs found for %s (requested %d)", len(firstTransactions.Result), contractAddress.Hex(), numTxs)
	}

	return firstTransactions.Result, nil
}

type TransactionsResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

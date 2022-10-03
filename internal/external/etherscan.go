package external

import (
	"bytes"
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

	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
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
	Account     string   `json:"account"`
	BalanceETH  *big.Int `json:"balance"`
	BalanceWETH *big.Int `json:"balance_weth"`
}

type Token string

const (
	WETH Token = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
)

const apiBaseURL = "https://api.etherscan.io/api"

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
	response, err := utils.HTTP.GetWithTLS12(url)
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
	response, err := utils.HTTP.GetWithTLS12(url)
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
		if balance, err := strconv.ParseInt(accountBalance.Balance, 10, 64); err == nil {
			balances = append(balances, &AccountBalance{
				Account:     accountBalance.Account,
				BalanceETH:  big.NewInt(balance),
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
	return GetTokenBalance(walletAddress, common.HexToAddress(string(WETH)))
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

	response, err := utils.HTTP.GetWithTLS12(url)
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

		return nil, errors.New("invalid json")
	}

	// decode the data
	var tokenBalanceResponse *TokenBalancesResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&tokenBalanceResponse); err != nil {
		gbl.Log.Warnf("token balance · decode error: %s", err.Error())

		return nil, err
	}

	// convert the balance to a big int
	balance, err := strconv.ParseInt(tokenBalanceResponse.Result, 10, 64)
	if err != nil {
		gbl.Log.Warnf("token balance · parse error: %s", err.Error())

		return nil, err
	}

	return big.NewInt(balance), nil
}

func withAPIKey(url string) string {
	return url + "&apikey=" + viper.GetString("api_keys.etherscan")
}

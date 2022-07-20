package etherscan

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
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

type AccountBalancesResponse struct {
	Response
	Result []AccountBalance `json:"result"`
}

type AccountBalance struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

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

	client, _ := createEtherscanHTTPClient()

	apiKey := viper.GetString("api_keys.etherscan")
	url := fmt.Sprint("https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=", apiKey)

	gbl.Log.Debugf("gas oracle url: %s", url)

	request, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
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

	responseBody, _ := ioutil.ReadAll(response.Body)

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

func MultiAccountBalance(wallets models.Wallets) *[]AccountBalance {
	if !viper.IsSet("api_keys.etherscan") {
		gbl.Log.Warnf("api_keys.etherscan not set")

		return nil
	}

	client, _ := createEtherscanHTTPClient()

	addressList := strings.Join(wallets.StringAddresses(), ",")

	apiKey := viper.GetString("api_keys.etherscan")
	url := fmt.Sprint("https://api.etherscan.io/api?module=account&action=balancemulti&tag=latest&apikey=", apiKey, "&address=", addressList)

	gbl.Log.Debugf("multiAccountBalance url: %s", url)

	request, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		gbl.Log.Warnf("multiAccountBalance error: %+v", err.Error())

		return nil
	}

	gbl.Log.Debugf("multiAccountBalance status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var accountBalancesResponse *AccountBalancesResponse

	responseBody, err := ioutil.ReadAll(response.Body)

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		gbl.Log.Warnf("multiAccountBalance invalid json: %s", err)

		return nil
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&accountBalancesResponse); err != nil {
		gbl.Log.Warnf("multiAccountBalance decode error: %s", err.Error())

		return nil
	}

	accountBalances := accountBalancesResponse.Result

	return &accountBalances
}

func createEtherscanHTTPClient() (*http.Client, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	transport := &http.Transport{
		TLSClientConfig:     tlsConfig,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     5 * time.Second,
	}

	// explicitly use http2
	_ = http2.ConfigureTransport(transport)

	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: transport,
	}

	return client, nil
}

package flots

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/spf13/viper"
)

type Client struct {
	// mainnet rpc client
	w3Client *w3.Client

	// flashbots rpc client
	fbClient *w3.Client

	PlusBlocks *big.Int

	SignerPublicKey *ecdsa.PublicKey
}

func New() *Client {
	//
	// connection to the flashbots relay node
	var fbClient *w3.Client

	if !viper.IsSet("flots.relay") {
		log.Fatal("❌ invalid or missing flots mainnet provider")
	}

	//
	// create the flots client
	flots := &Client{PlusBlocks: big.NewInt(viper.GetInt64("flots.plusBlocks"))}

	// check for valid flashbots signer key
	if signerKey, err := crypto.HexToECDSA(viper.GetString("flots.signerKey")); err == nil {
		// create the flashbots client
		fbClient = flashbots.MustDial(viper.GetString("flots.relay"), signerKey)

		flots.SignerPublicKey = &signerKey.PublicKey
	} else {
		log.Fatalf("❌ invalid or missing signer key: %v", err)
	}

	//
	// connection to a mainnet node
	var w3Client *w3.Client

	if provider := viper.GetString("flots.provider"); provider != "" {
		w3Client = w3.MustDial(viper.GetString("flots.provider"))
	} else {
		log.Fatal("❌ invalid or missing flots mainnet provider")
	}

	flots.fbClient = fbClient
	flots.w3Client = w3Client

	log.Debugf("flots: %+v\n", flots)

	return flots
}

func (c *Client) UserAddress() common.Address {
	return crypto.PubkeyToAddress(*c.SignerPublicKey)
}

// LatestBlock gets the latest block number.
func (c *Client) LatestBlock() *big.Int {
	var latestBlock big.Int

	if err := c.w3Client.Call(
		eth.BlockNumber().Returns(&latestBlock),
	); err != nil {
		log.Errorf("❌ failed to fetch latest block: %v", err)

		return nil
	}

	return &latestBlock
}

// LatestBlockPlus gets the latest block number plus the configured offset.
func (c *Client) LatestBlockPlus() *big.Int {
	return new(big.Int).Add(c.LatestBlock(), c.PlusBlocks)
}

func (c *Client) GetUserStats() *flashbots.UserStatsV2Response {
	var userStats flashbots.UserStatsV2Response

	if err := c.fbClient.Call(
		flashbots.UserStatsV2(c.LatestBlockPlus()).Returns(&userStats),
	); err != nil {
		log.Errorf("Failed to fetch user stats: %v\n", err)

		return nil
	}

	return &userStats
}

func (c *Client) GetBundleStats(bundleHash common.Hash) *flashbots.BundleStatsV2Response {
	var bundleStats flashbots.BundleStatsV2Response

	if err := c.fbClient.Call(
		flashbots.BundleStatsV2(bundleHash, c.LatestBlockPlus()).Returns(&bundleStats),
	); err != nil {
		log.Fatalf("❌ failed to fetch bundle stats: %v\n", err)

		return nil
	}

	return &bundleStats
}

func (c *Client) CallBundle(rawTxs [][]byte) *flashbots.CallBundleResponse {
	// create request
	callBundleRequest := &flashbots.CallBundleRequest{
		BlockNumber:     c.LatestBlockPlus(),
		RawTransactions: rawTxs,
	}

	log.Debugf("callBundleRequest: %+v\n", callBundleRequest)

	//
	// call bundle
	var callBundle flashbots.CallBundleResponse

	if err := c.fbClient.Call(
		flashbots.CallBundle(callBundleRequest).Returns(&callBundle),
	); err != nil {
		log.Fatalf("❌ failed to call bundle: %v\n", err)

		return nil
	}

	return &callBundle
}

func (c *Client) SendBundleWithRawTxs(rawTxs [][]byte) common.Hash {
	// create request
	sendBundleRequest := &flashbots.SendBundleRequest{
		RawTransactions: rawTxs,
		BlockNumber:     c.LatestBlockPlus(),
	}

	log.Debugf("sendBundleRequest: %+v\n", sendBundleRequest)

	//
	// call bundle
	var bundleHash common.Hash

	if err := c.fbClient.Call(
		flashbots.SendBundle(sendBundleRequest).Returns(&bundleHash),
	); err != nil {
		log.Fatalf("❌ failed to send bundle: %v\n", err)

		return common.Hash{}
	}

	return bundleHash
}

func (c *Client) MevSendBundle(rawTxs [][]byte) common.Hash {
	// create request
	mevSendBundleRequest := &MevSendBundleRequest{
		Version: "v0.1",
		Body:    []MevBundleTx{},
		Inclusion: MevParamInclusion{
			Block:    c.LatestBlock().Text(16),
			MaxBlock: c.LatestBlockPlus().Text(16),
		},
		Privacy: MevParamPrivacy{
			Hints: []string{
				"calldata", "contract_address", "logs", "function_selector", "hash",
			},
		},
	}

	log.Debugf("sendBundleRequest: %+v | %+v", mevSendBundleRequest, rawTxs)

	//
	// call bundle
	var bundleHash common.Hash

	if err := c.fbClient.Call(
		MevSendBundle(mevSendBundleRequest).Returns(&bundleHash),
	); err != nil {
		log.Fatalf("❌ failed to send bundle: %v\n", err)

		return common.Hash{}
	}

	return bundleHash
}

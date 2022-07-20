package gbnode

import (
	"crypto/tls"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/http2"
)

func replaceSchemeWithGateway(url string, gateway string) string {
	const schemeIPFS = "ipfs://"

	return strings.Replace(url, schemeIPFS, gateway, 1)
}

func createMetadataHTTPClient() (*http.Client, error) {
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

func (p ChainNode) GetTokenURI(contractAddress common.Address, tokenID *big.Int) (string, error) {
	gbl.Log.Infof("GetTokenURI || contractAddress: %s | tokenID: %d\n", contractAddress, tokenID)

	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return "", err
	}

	// collection total supply
	tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, tokenID)
	if err != nil {
		gbl.Log.Error(err)

		return "", err
	}

	return tokenURI, nil
}

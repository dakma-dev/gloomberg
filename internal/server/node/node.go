package node

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/topic"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

// Node represents a w3 provider configuration.
type Node struct {
	NodeID             int    `mapstructure:"id"`
	Name               string `mapstructure:"name"`
	Marker             string `mapstructure:"marker"`
	Client             *ethclient.Client
	ClientGeth         *gethclient.Client
	WebsocketsEndpoint string     `mapstructure:"endpoint"`
	ReceivedMessages   uint64     `mapstructure:"received_messages"`
	KillTimer          time.Timer `mapstructure:"kill_timer"`
	Error              error      `mapstructure:"error"`
	NumLogsReceived    uint64
	LastLogReceived    int64
	LocalNode          bool
}

type GasInfo struct {
	LastBlock         uint64   `json:"lastBlock"`
	LastBlockGasLimit uint64   `json:"lastBlockGasLimit"`
	GasPriceWei       *big.Int `json:"gasPrice"`
	GasTipWei         *big.Int `json:"gasTip"`
}

// New returns a new gbnode if connection to the given endpoint succeeds.
func New(nodeID int, name string, marker string, endpoint string, localNode bool) (*Node, error) {
	rpcClient, err := rpc.DialContext(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)
	// if err != nil {
	// 	gbl.Log.Warnf("Failed to connect to the Ethereum node: %s", err)
	// }

	var gethClient *gethclient.Client
	if nodeID == 0 {
		gethClient = gethclient.New(rpcClient)
	}

	return &Node{
		NodeID:             nodeID,
		Name:               name,
		Marker:             marker,
		Client:             ethClient,
		ClientGeth:         gethClient,
		WebsocketsEndpoint: endpoint,
		Error:              err,
		LocalNode:          localNode,
	}, err
}

// GetCurrentGasInfo returns the current gas price and tip
func (p Node) GetCurrentGasInfo() (*GasInfo, error) {
	ctx := context.Background()

	block, err := p.Client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasPrice, err := p.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasTip, err := p.Client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	return &GasInfo{
		LastBlock:         block.NumberU64(),
		LastBlockGasLimit: block.GasLimit(),
		GasPriceWei:       gasPrice,
		GasTipWei:         gasTip,
	}, nil
}

// // gas price
// if gasNode := sNodes.GetRandomGasNode(); gasNode != nil {
// 	if gasPrice, err := gasNode.Client.SuggestGasPrice(context.Background()); err != nil {
// 		gbl.Log.Error(err)
// 	} else {
// 		gasPriceGwei, _ := weiToGwei(gasPrice).Float64()
// 		out.WriteString("  | " + style.GrayStyle.Render(fmt.Sprintf("⛽️ %.1f", gasPriceGwei)))
// 	}

// 	if gasTip, err := gasNode.Client.SuggestGasTipCap(context.Background()); err != nil {
// 		gbl.Log.Error(err)
// 	} else {
// 		gasTipGwei, _ := weiToGwei(gasTip).Uint64()
// 		out.WriteString(style.GrayStyle.Render(fmt.Sprintf(" / %d", gasTipGwei)))
// 	}
// }

func (p Node) SubscribeToPendingTransactions(queueLogs chan<- common.Hash) (ethereum.Subscription, error) {
	return p.ClientGeth.SubscribePendingTransactions(context.Background(), queueLogs)
}

func (p Node) SubscribeToContract(queueLogs chan types.Log, contractAddress common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, nil, []common.Address{contractAddress})
}

func (p Node) SubscribeToTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.Transfer))}}, nil)
}

func (p Node) SubscribeToSingleTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.TransferSingle))}}, nil)
}

func (p Node) SubscribeToTransfersFor(queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.Transfer))}}, contractAddresses)
}

func (p Node) SubscribeToSingleTransfersFor(queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.TransferSingle))}}, contractAddresses)
}

func (p Node) SubscribeToOrdersMatched(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.OrdersMatched))}}, nil)
}

func (p Node) subscribeTo(queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
	ctx := context.Background()

	if topics == nil && contractAddresses == nil {
		return nil, nil
	}

	filterQuery := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics:    topics,
	}

	return p.Client.SubscribeFilterLogs(ctx, filterQuery, queueLogs)
}

func (p Node) GetERC1155TokenName(contractAddress common.Address, tokenID *big.Int) (string, error) {
	//
	// ERC1155
	contractERC1155, err := abis.NewERC1155(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return "", err
	}

	if tokenID == nil {
		tokenID = big.NewInt(1)
	}

	if name, err := contractERC1155.Name(&bind.CallOpts{}, tokenID); err == nil && name != "" {
		gbl.Log.Debugf("found collection name via erc1155 chain call: %s", name)

		return name, nil
	}

	if uri, err := contractERC1155.Uri(&bind.CallOpts{}, tokenID); err == nil {
		gbl.Log.Debugf("found collection uri via erc1155 chain call: %s", uri)

		if metadata, err := external.GetERC1155MetadataForURI(uri, tokenID); err == nil && metadata != nil {
			name := metadata.Name
			if metadata.CreatedBy != "" {
				name = metadata.CreatedBy + " | " + name
			}

			gbl.Log.Debugf("found collection name via erc1155 metadata: %v\n", name)

			return name, nil
		}
	}

	return "", errors.New("could not find collection name")
}

func (p Node) GetCollectionName(contractAddress common.Address, tokenID *big.Int) (string, error) {
	// get the contractERC721 ABI
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return "", err
	}

	if name, err := contractERC721.Name(&bind.CallOpts{}); err == nil {
		gbl.Log.Debugf("found collection name via erc721 chain call: %s", name)

		return name, nil
	}

	return "", nil
}

func (p Node) GetCollectionMetadata(contractAddress common.Address) map[string]interface{} {
	collectionMetadata := make(map[string]interface{}, 0)

	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)
	}

	// attributes := []string{"name", "symbol", "totalSupply", "tokenURI"}

	if value, err := contractERC721.Name(&bind.CallOpts{}); err != nil {
		gbl.Log.Debug(err)
	} else {
		collectionMetadata["name"] = value
	}

	if symbol, err := contractERC721.Symbol(&bind.CallOpts{}); err != nil {
		gbl.Log.Debug(err)
	} else {
		collectionMetadata["symbol"] = symbol
	}

	if totalSupply, err := contractERC721.TotalSupply(&bind.CallOpts{}); err != nil {
		gbl.Log.Debug(err)
	} else {
		collectionMetadata["totalSupply"] = totalSupply.Uint64()
	}

	if tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, big.NewInt(0)); err != nil {
		gbl.Log.Debug(err)
	} else {
		collectionMetadata["tokenURI"] = tokenURI
	}

	return collectionMetadata
}

func (p Node) GetTokenMetadata(tokenURI string) (*models.MetadataERC721, error) {
	gbl.Log.Infof("GetTokenMetadata || tokenURI: %+v\n", tokenURI)

	client, _ := createMetadataHTTPClient()

	// for ipfs use the default "gateway"
	tokenURI = replaceSchemeWithGateway(tokenURI, viper.GetString("ipfs.gateway"))

	request, _ := http.NewRequest("GET", tokenURI, nil)

	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		gbl.Log.Warnf("get token metadata error: %+v | response: %+v\n", err, response)

		return nil, err
	}

	gbl.Log.Infof("get token metadata status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var tokenMetadata *models.MetadataERC721

	responseBody, err := io.ReadAll(response.Body)

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		gbl.Log.Warnf("get token metadata invalid json: %s\n", err)

		return nil, err
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&tokenMetadata); err != nil {
		gbl.Log.Warnf("get token metadata decode error: %s\n", err.Error())

		return nil, err
	}

	gbl.Log.Infof("GetTokenMetadata || tokenMetadata: %+v\n", tokenMetadata)

	return tokenMetadata, nil
}

func (p Node) GetTokenImageURI(contractAddress common.Address, tokenID uint64) (string, error) {
	gbl.Log.Infof("GetTokenImageURI || contractAddress: %s | tokenID: %d\n", contractAddress, tokenID)

	tokenURI, err := p.GetTokenURI(contractAddress, big.NewInt(int64(tokenID)))
	if err != nil {
		gbl.Log.Errorf("get token image uri error: %+v\n", err.Error())

		return "", err
	}

	metadata, err := p.GetTokenMetadata(tokenURI)
	if err != nil || metadata == nil {
		gbl.Log.Errorf("get token image uri error: %+v\n", err.Error())

		return "", err
	}

	gbl.Log.Infof("GetTokenImageURI || metadata: %+v\n", metadata)

	return replaceSchemeWithGateway(metadata.Image, viper.GetString("ipfs.gateway")), nil
}

func (p Node) GetTokenURI(contractAddress common.Address, tokenID *big.Int) (string, error) {
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

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func WeiToGwei(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.GWei))
}

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

package gbnode

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type CollectionMetadata struct {
	Name        string `json:"name"`
	OpenseaSlug string `mapstructure:"slug"`
	Symbol      string `mapstructure:"symbol"`
	TotalSupply uint64 `mapstructure:"total_supply"`
	TokenURI    string `mapstructure:"token_uri"`
}

type MetadataERC721 struct {
	Image string `json:"image"`
}

type NodeCollection []*ChainNode

// GetRandomNode returns a random available ethereum node.
func (nc *NodeCollection) GetRandomNode() *ChainNode {
	if len(*nc) == 0 {
		return nil
	}

	//nolint:gosec
	return (*nc)[rand.Intn(len(*nc))]
}

// GetNodes returns all nodes as slice.
func (nc *NodeCollection) GetNodes() []*ChainNode {
	if len(*nc) == 0 {
		return nil
	}

	nodes := make([]*ChainNode, len(*nc))

	for i, node := range *nc {
		nodes[i] = node
	}

	return nodes
}

func (nc *NodeCollection) SubscribeToAllTransfers(ctx context.Context, queueLogs chan types.Log) {
	for _, node := range nc.GetNodes() {
		// subscribe to all "Transfer" events
		if _, err := node.SubscribeToTransfers(ctx, queueLogs); err != nil {
			gbl.Log.Warnf("Transfers subscribe to %s failed: ", node.WebsocketsEndpoint, err)
		}

		// subscribe to all "SingleTransfer" events
		if _, err := node.SubscribeToSingleTransfers(ctx, queueLogs); err != nil {
			gbl.Log.Warnf("SingleTransfers subscribe to %s failed: ", node.WebsocketsEndpoint, err)
		}

		gbl.Log.Infof("subTransfers & subSingleTransfers for ALL")
	}
}

type Topic string

const (
	OrdersMatched  Topic = "0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9"
	Transfer       Topic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	TransferSingle Topic = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	ApprovalForAll Topic = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
)

func (t Topic) String() string {
	return map[Topic]string{
		OrdersMatched: "OrdersMatched", Transfer: "Transfer", TransferSingle: "TransferSingle", ApprovalForAll: "ApprovalForAll",
	}[t]
}

var ctx = context.Background()

// New returns a new gbnode if connection to the given endpoint succeeds.
func New(endpoint string) (*ChainNode, error) {
	client, err := ethclient.DialContext(ctx, endpoint)
	if err != nil {
		gbl.Log.Errorf("failed to connect to %s: %s", endpoint, err)

		return nil, err
	}

	return &ChainNode{
		Client:             client,
		WebsocketsEndpoint: endpoint,
	}, nil
}

// GetNodesFromConfig reads the websockets endpoints slice from config, connects to them and returns a list with successfully connected nodes.
func GetNodesFromConfig(endpoints []string) *NodeCollection {
	var gbnodeWg sync.WaitGroup

	var nodes NodeCollection

	for _, endpoint := range endpoints {
		gbnodeWg.Add(1)

		endpoint := endpoint

		go func() {
			defer gbnodeWg.Done()

			if provider, err := New(endpoint); err != nil {
				gbl.Log.Errorf("failed to connect to %s: %s", endpoint, err)
			} else {
				nodes = append(nodes, provider)
			}
		}()
	}

	gbnodeWg.Wait()

	return &nodes
}

// ChainNode represents a w3 provider configuration.
type ChainNode struct {
	Name               string `mapstructure:"name"`
	Client             *ethclient.Client
	WebsocketsEndpoint string `mapstructure:"url"`
}

func (p ChainNode) SubscribeToContract(ctx context.Context, queueLogs chan types.Log, contractAddress common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(ctx, queueLogs, nil, []common.Address{contractAddress})
}

func (p ChainNode) SubscribeToTransfers(ctx context.Context, queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(ctx, queueLogs, [][]common.Hash{{common.HexToHash(string(Transfer))}}, nil)
}

func (p ChainNode) SubscribeToSingleTransfers(ctx context.Context, queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(ctx, queueLogs, [][]common.Hash{{common.HexToHash(string(TransferSingle))}}, nil)
}

func (p ChainNode) SubscribeToTransfersFor(ctx context.Context, queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(ctx, queueLogs, [][]common.Hash{{common.HexToHash(string(Transfer))}}, contractAddresses)
}

func (p ChainNode) SubscribeToSingleTransfersFor(ctx context.Context, queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(ctx, queueLogs, [][]common.Hash{{common.HexToHash(string(TransferSingle))}}, contractAddresses)
}

func (p ChainNode) SubscribeToOrdersMatched(ctx context.Context, queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(ctx, queueLogs, [][]common.Hash{{common.HexToHash(string(OrdersMatched))}}, nil)
}

func (p ChainNode) subscribeTo(ctx context.Context, queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
	if topics == nil && contractAddresses == nil {
		return nil, nil
	}

	filterQuery := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics:    topics,
	}

	return p.Client.SubscribeFilterLogs(ctx, filterQuery, queueLogs)
}

func (p ChainNode) GetCollectionName(contractAddress common.Address) string {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)
	}

	// collection name
	collectionName := ""

	// check if the collection is already in the redis cache
	if viper.GetBool("redis.enabled") {
		if rdb := cache.GetRedisClient(); rdb != nil {
			gbl.Log.Debugf("redis | searching for contract address: %s", contractAddress.String())

			if name, err := rdb.Get(ctx, cache.KeyContract(contractAddress)).Result(); err == nil && name != "" {
				gbl.Log.Debugf("redis | using cached contractName: %s", name)

				collectionName = name
			}
		}
	}

	// if not found in redis, we call the contract method to get the name
	if collectionName == "" {
		if name, err := contractERC721.Name(&bind.CallOpts{}); err == nil {
			collectionName = name

			if viper.GetBool("redis.enabled") {
				if rdb := cache.GetRedisClient(); rdb != nil {
					err := rdb.SetEX(ctx, cache.KeyContract(contractAddress), collectionName, time.Hour*48).Err()

					if err != nil {
						gbl.Log.Infof("redis | error while adding: %s", err.Error())
					} else {
						gbl.Log.Debugf("redis | added: %s -> %s", contractAddress.Hex(), collectionName)
					}
				}
			}
		} else {
			collectionName = style.ShortenAddress(&contractAddress)
		}
	}

	return collectionName
}

func (p ChainNode) GetCollectionMetadata(contractAddress common.Address) *CollectionMetadata {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)
	}

	// collection name
	collectionName := ""

	// check if the collection is already in the redis cache
	if viper.GetBool("redis.enabled") {
		if rdb := cache.GetRedisClient(); rdb != nil {
			gbl.Log.Debugf("redis | searching for contract address: %s", contractAddress.String())

			if name, err := rdb.Get(ctx, cache.KeyContract(contractAddress)).Result(); err == nil && name != "" {
				gbl.Log.Debugf("redis | using cached contractName: %s", name)

				collectionName = name
			}
		}
	}

	// if not found in redis, we call the contract method to get the name
	if collectionName == "" {
		if name, err := contractERC721.Name(&bind.CallOpts{}); err == nil {
			collectionName = name

			if viper.GetBool("redis.enabled") {
				if rdb := cache.GetRedisClient(); rdb != nil {
					err := rdb.SetEX(ctx, cache.KeyContract(contractAddress), collectionName, time.Hour*48).Err()

					if err != nil {
						gbl.Log.Infof("redis | error while adding: %s", err.Error())
					} else {
						gbl.Log.Debugf("redis | added: %s -> %s", contractAddress.Hex(), collectionName)
					}
				}
			}
		} else {
			collectionName = style.ShortenAddress(&contractAddress)
		}
	}

	// collection total supply
	collectionTotalSupply := uint64(0)
	if totalSupply, err := contractERC721.TotalSupply(&bind.CallOpts{}); err == nil {
		collectionTotalSupply = totalSupply.Uint64()
	}

	// collection symbol
	collectionSymbol := ""
	if symbol, err := contractERC721.Symbol(&bind.CallOpts{}); err == nil {
		collectionSymbol = symbol
	}

	// collection token uri
	collectionTokenURI := ""
	if tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, big.NewInt(1)); err == nil {
		collectionTokenURI = tokenURI
	}

	return &CollectionMetadata{
		Name:        collectionName,
		Symbol:      collectionSymbol,
		TotalSupply: collectionTotalSupply,
		TokenURI:    collectionTokenURI,
	}
}

func (p ChainNode) GetTokenMetadata(tokenURI string) (*MetadataERC721, error) {
	gbl.Log.Infof("GetTokenMetadata || tokenURI: %+v\n", tokenURI)

	client, _ := createMetadataHTTPClient()

	// for ipfs use the default "gateway"
	tokenURI = replaceSchemeWithGateway(tokenURI, viper.GetString("ipfs.gateway"))

	request, _ := http.NewRequest("GET", tokenURI, nil)

	response, err := client.Do(request)
	if err != nil {
		gbl.Log.Warnf("get token metadata error: %+v\n", err.Error())

		return nil, err
	}

	gbl.Log.Infof("get token metadata status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var tokenMetadata *MetadataERC721

	responseBody, err := ioutil.ReadAll(response.Body)

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

func (p ChainNode) GetTokenImageURI(contractAddress common.Address, tokenID uint64) (string, error) {
	gbl.Log.Infof("GetTokenImageURI || contractAddress: %s | tokenID: %d\n", contractAddress, tokenID)

	tokenURI, err := p.GetTokenURI(contractAddress, big.NewInt(int64(tokenID)))
	if err != nil {
		gbl.Log.Errorf("get token image uri error: %+v\n", err.Error())

		return "", err
	}

	metadata, err := p.GetTokenMetadata(tokenURI)
	if err != nil {
		gbl.Log.Errorf("get token image uri error: %+v\n", err.Error())

		return "", err
	}

	gbl.Log.Infof("GetTokenImageURI || metadata: %+v\n", metadata)

	return replaceSchemeWithGateway(metadata.Image, viper.GetString("ipfs.gateway")), nil
}

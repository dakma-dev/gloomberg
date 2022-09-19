package gbnode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

var ctx = context.Background()

// type CollectionMetadata struct {
// 	ContractName string `json:"contractName"`
// 	Symbol       string `json:"symbol"`
// 	TotalSupply  uint64 `json:"total_supply"`
// 	TokenURI     string `json:"token_uri"`
// }

// type MetadataERC721 struct {
// 	Image string `json:"image"`
// }

type NodeCollection []*ChainNode

// GetRandomNode returns a random available ethereum node.
func (nc *NodeCollection) GetRandomNode() *ChainNode {
	if len(*nc) == 0 {
		return nil
	}

	//nolint:gosec
	return (*nc)[rand.Intn(len(*nc))]
}

// func (nc *NodeCollection) GetRandomGasNode() *ChainNode {
// 	if len(*nc) == 0 {
// 		return nil
// 	}

// 	for _, node := range *nc {
// 		fmt.Println("node: ", node.Name, node.LocalNode)
// 		if node.LocalNode {
// 			return node
// 		}
// 	}

// 	return nil
// }

// GetNodes returns all nodes as slice.
func (nc *NodeCollection) GetNodes() []*ChainNode {
	if len(*nc) == 0 {
		return nil
	}

	nodes := make([]*ChainNode, len(*nc))

	copy(nodes, *nc)

	// for i, node := range *nc {
	// 	nodes[i] = node
	// }

	return nodes
}

// GetNodeByID GetNodes returns all nodes as slice.
func (nc *NodeCollection) GetNodeByID(nodeID int) *ChainNode {
	if len(*nc) == 0 {
		return nil
	}

	for _, node := range *nc {
		if node.NodeID == nodeID {
			return node
		}
	}

	return nil
}

func (nc *NodeCollection) SubscribeToAllTransfers(queueLogs chan types.Log) {
	for _, node := range nc.GetNodes() {
		// subscribe to all "Transfer" events
		if _, err := node.SubscribeToTransfers(queueLogs); err != nil {
			gbl.Log.Warnf("Transfers subscribe to %s failed: %s", node.WebsocketsEndpoint, err)
		}

		// subscribe to all "SingleTransfer" events
		if _, err := node.SubscribeToSingleTransfers(queueLogs); err != nil {
			gbl.Log.Warnf("SingleTransfers subscribe to %s failed: %s", node.WebsocketsEndpoint, err)
		}
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

// ChainNode represents a w3 provider configuration.
type ChainNode struct {
	NodeID             int    `mapstructure:"id"`
	Name               string `mapstructure:"name"`
	Marker             string `mapstructure:"marker"`
	Client             *ethclient.Client
	WebsocketsEndpoint string     `mapstructure:"endpoint"`
	ReceivedMessages   uint64     `mapstructure:"received_messages"`
	KillTimer          time.Timer `mapstructure:"kill_timer"`
	Error              error      `mapstructure:"error"`
	NumLogsReceived    uint64
	LastLogReceived    int64
	LocalNode          bool
}

// // New returns a new gbnode if connection to the given endpoint succeeds.
// func New(nodeID int, name string, marker string, endpoint string, localNode bool) (*ChainNode, error) {
// 	ctx := context.Background()
// 	client, err := ethclient.DialContext(ctx, endpoint)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	return &ChainNode{
// 		NodeID:             nodeID,
// 		Name:               name,
// 		Marker:             marker,
// 		Client:             client,
// 		WebsocketsEndpoint: endpoint,
// 		Error:              err,
// 		LocalNode:          localNode,
// 	}, err
// }

func (p ChainNode) SubscribeToContract(queueLogs chan types.Log, contractAddress common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, nil, []common.Address{contractAddress})
}

func (p ChainNode) SubscribeToTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(Transfer))}}, nil)
}

func (p ChainNode) SubscribeToSingleTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(TransferSingle))}}, nil)
}

func (p ChainNode) SubscribeToTransfersFor(queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(Transfer))}}, contractAddresses)
}

func (p ChainNode) SubscribeToSingleTransfersFor(queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(TransferSingle))}}, contractAddresses)
}

func (p ChainNode) SubscribeToOrdersMatched(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(OrdersMatched))}}, nil)
}

func (p ChainNode) subscribeTo(queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
	if topics == nil && contractAddresses == nil {
		return nil, errors.New("topics and contractAddresses are nil")
	}

	filterQuery := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics:    topics,
	}

	return p.Client.SubscribeFilterLogs(ctx, filterQuery, queueLogs)
}

func (p ChainNode) GetCollectionName(contractAddress common.Address) (string, error) {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return "", err
	}

	if name, err := contractERC721.Name(&bind.CallOpts{}); err == nil {
		gbl.Log.Infof("found collection name via chain call: %s", name)

		return name, nil
	}

	return "", nil
}

func (p ChainNode) GetCollectionMetadata(contractAddress common.Address) *models.CollectionMetadata {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)
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

	return &models.CollectionMetadata{
		Symbol:      collectionSymbol,
		TotalSupply: collectionTotalSupply,
		TokenURI:    collectionTokenURI,
	}
}

func (p ChainNode) GetTokenMetadata(tokenURI string) (*models.MetadataERC721, error) {
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

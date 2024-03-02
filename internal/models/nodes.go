package models

import (
	"context"

	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

// Nodes represents a slice of nodes.
type Nodes []*Node

// Names returns the names of all nodes.
func (n Nodes) Names() []string {
	nodeNames := make([]string, 0)

	for _, node := range n {
		nodeNames = append(nodeNames, node.Name)
	}

	return nodeNames
}

// type NodeConfigurations []NodeConfiguration

// NodeConfiguration represents a node configuration.
type NodeConfiguration struct {
	Name     string `json:"name"     mapstructure:"name"`
	Endpoint string `json:"endpoint" mapstructure:"endpoint"`
}

// Node represents a node in the network/rpc endpoint.
type Node struct {
	NodeConfiguration `json:",inline" mapstructure:",squash"`

	EthClient  *ethclient.Client  `json:"-" mapstructure:"-"`
	GethClient *gethclient.Client `json:"-" mapstructure:"-"`

	// currently unfortunately handled "globally" in chainwatcher
	// LastLogReceivedAt time.Time `json:"-" mapstructure:"-"`
}

func (n *Node) String() string {
	return n.Name
}

// Connect connects to the node.
func (n *Node) reconnectClient() error {
	ethClient, err := ethclient.Dial(n.Endpoint)
	if err != nil {
		return err
	}

	n.EthClient = ethClient
	n.GethClient = gethclient.New(n.EthClient.Client())

	return nil
}

// Reconnect closes the current connection and reconnects to the node.
func (n *Node) Connect() error {
	if n.EthClient != nil {
		n.EthClient.Close()
	}

	n.EthClient = nil
	n.GethClient = nil

	return n.reconnectClient()
}

// SubscribeToNFTTransfers subscribes to all transfer events of erc721/erc1155 tokens (4 topics).
func (n *Node) SubscribeToNFTTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	subscribeTopics := [][]common.Hash{
		{common.HexToHash(string(topic.Transfer)), common.HexToHash(string(topic.TransferSingle))},
		{},
		{},
		{},
	}

	return n.subscribeToLogs(queueLogs, subscribeTopics, nil)
}

// SubscribeToERC20Transfers subscribes to all transfer events of erc20 tokens (3 topics).
func (n *Node) SubscribeToERC20Transfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	subscribeTopics := [][]common.Hash{
		{common.HexToHash(string(topic.Transfer)), common.HexToHash(string(topic.TransferSingle))},
		{},
		{},
	}

	return n.subscribeToLogs(queueLogs, subscribeTopics, nil)
}

// subscribeToLogs subscribes to the given topics and contract addresses.
func (n *Node) subscribeToLogs(queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
	ctx := context.Background()

	if topics == nil && contractAddresses == nil {
		return nil, nil
	}

	filterQuery := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics:    topics,
	}

	return n.EthClient.SubscribeFilterLogs(ctx, filterQuery, queueLogs)
}

// subscribeToLogs subscribes to new block header events.
func (n *Node) SubscribeToHeads(queueHeads chan *types.Header) (ethereum.Subscription, error) {
	ctx := context.Background()

	return n.EthClient.SubscribeNewHead(ctx, queueHeads)
}

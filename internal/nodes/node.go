package nodes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"strconv"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/standard"
	"github.com/benleb/gloomberg/internal/models/topic"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// ERC165 interface identifier
// var iERC165 = [4]byte{0x01, 0xff, 0xc9, 0xa7}

// ERC1155 interface identifier
var iERC1155 = [4]byte{0xd9, 0xb6, 0x7a, 0x26}

// contract.supportsInterface(0x01ffc9a7) as bytes
var supportsInterfaceiERC165 = []byte{0x01, 0xff, 0xc9, 0xa7, 0x01, 0xff, 0xc9, 0xa7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

// contract.supportsInterface(0xd9b67a26) as bytes
var supportsInterfaceiFFFFFF = []byte{0x01, 0xff, 0xc9, 0xa7, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

// Node represents a w3 provider configuration.
type Node struct {
	NodeID             int            `mapstructure:"id"`
	Name               string         `mapstructure:"name"`
	Color              lipgloss.Color `mapstructure:"color"`
	Marker             string         `mapstructure:"marker"`
	WebsocketsEndpoint string         `mapstructure:"endpoint"`
	LocalNode          bool           `mapstructure:"local"`
	Client             *ethclient.Client
	ClientGeth         *gethclient.Client
	NumLogsReceived    uint64
	LastLogReceived    int64
}

type GasInfo struct {
	LastBlock         uint64   `json:"lastBlock"`
	LastBlockGasLimit uint64   `json:"lastBlockGasLimit"`
	GasPriceWei       *big.Int `json:"gasPrice"`
	GasTipWei         *big.Int `json:"gasTip"`
}

// Connect returns a new node if the connection to the given endpoint succeeds.
func (n *Node) Connect() error {
	var err error

	rpcClient, err := rpc.DialContext(context.Background(), n.WebsocketsEndpoint)
	if err != nil {
		gbl.Log.Warnf("Failed to connect to node %s: %s", n.Name, err)
		return err
	}

	ethClient := ethclient.NewClient(rpcClient)
	if ethClient == nil {
		gbl.Log.Warnf("Failed to start eth client for node %s: %s", n.Name, err)
	}

	n.Client = ethClient

	if n.LocalNode {
		n.ClientGeth = gethclient.New(rpcClient)
	}

	return err
}

func (n *Node) GetStyledMarker() string {
	return lipgloss.NewStyle().Foreground(n.Color).Render(n.Marker)
}

// GetCurrentGasInfo returns the current gas price and tip
func (n *Node) GetCurrentGasInfo() (*GasInfo, error) {
	ctx := context.Background()

	block, err := n.Client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasPrice, err := n.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasTip, err := n.Client.SuggestGasTipCap(ctx)
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

func (n *Node) SubscribeToPendingTransactions(queueLogs chan<- common.Hash) (ethereum.Subscription, error) {
	return n.ClientGeth.SubscribePendingTransactions(context.Background(), queueLogs)
}

func (n *Node) SubscribeToContract(queueLogs chan types.Log, contractAddress common.Address) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, nil, []common.Address{contractAddress})
}

func (n *Node) SubscribeToOrderFulfilled(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.OrderFulfilled))}}, nil)
}

func (n *Node) SubscribeToTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.Transfer))}}, nil)
}

func (n *Node) SubscribeToSingleTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.TransferSingle))}}, nil)
}

func (n *Node) SubscribeToTransfersFor(queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.Transfer))}}, contractAddresses)
}

func (n *Node) SubscribeToSingleTransfersFor(queueLogs chan types.Log, contractAddresses []common.Address) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.TransferSingle))}}, contractAddresses)
}

func (n *Node) SubscribeToOrdersMatched(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return n.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.OrdersMatched))}}, nil)
}

func (n *Node) subscribeTo(queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
	ctx := context.Background()

	if topics == nil && contractAddresses == nil {
		return nil, nil
	}

	filterQuery := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics:    topics,
	}

	return n.Client.SubscribeFilterLogs(ctx, filterQuery, queueLogs)
}

func (n *Node) GetERC1155TokenName(contractAddress common.Address, tokenID *big.Int) (string, error) {
	//
	// ERC1155
	contractERC1155, err := abis.NewERC1155(contractAddress, n.Client)
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

			gbl.Log.Debugf("found collection name via erc1155 metadata: %v", name)

			return name, nil
		}
	}

	return "", errors.New("could not find collection name")
}

func (n *Node) GetERC721CollectionName(contractAddress common.Address) (string, error) {
	// get the contractERC721 ABI
	contractERC721, err := abis.NewERC721v3(contractAddress, n.Client)
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

func (n *Node) GetCollectionMetadata(contractAddress common.Address) map[string]interface{} {
	collectionMetadata := make(map[string]interface{}, 0)

	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, n.Client)
	if err != nil {
		gbl.Log.Error(err)
	}

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

func (n *Node) GetTokenMetadata(tokenURI string) (*models.MetadataERC721, error) {
	gbl.Log.Debugf("GetTokenMetadata || tokenURI: %+v", tokenURI)

	response, err := utils.HTTP.GetWithTLS12(tokenURI)
	if err != nil || response.StatusCode != http.StatusOK {
		gbl.Log.Warnf("get token metadata error: %+v | response: %+v", err, response)

		return nil, err
	}

	gbl.Log.Debugf("get token metadata status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var tokenMetadata *models.MetadataERC721

	responseBody, err := io.ReadAll(response.Body)

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		gbl.Log.Warnf("get token metadata invalid json: %s", err)

		return nil, err
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&tokenMetadata); err != nil {
		gbl.Log.Warnf("get token metadata decode error: %s", err.Error())

		return nil, err
	}

	gbl.Log.Debugf("GetTokenMetadata || tokenMetadata: %+v", tokenMetadata)

	return tokenMetadata, nil
}

func (n *Node) GetTokenImageURI(contractAddress common.Address, tokenID *big.Int) (string, error) {
	gbl.Log.Debugf("GetTokenImageURI || contractAddress: %s | tokenID: %d", contractAddress, tokenID)

	tokenURI, err := n.GetTokenURI(contractAddress, tokenID)
	if err != nil {
		gbl.Log.Errorf("get token image uri error: %+v", err.Error())

		return "", err
	}

	gbl.Log.Debugf("GetTokenImageURI || tokenURI: %+v", tokenURI)

	metadata, err := n.GetTokenMetadata(tokenURI)
	if err != nil || metadata == nil {
		gbl.Log.Debugf("get token image uri error: %+v", err)

		return "", err
	}

	gbl.Log.Debugf("GetTokenImageURI || metadata: %+v", metadata)

	return metadata.Image, nil
}

func (n *Node) GetTokenURI(contractAddress common.Address, tokenID *big.Int) (string, error) {
	gbl.Log.Debugf("GetTokenURI || contractAddress: %s | tokenID: %d", contractAddress, tokenID)

	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, n.Client)
	if err != nil {
		gbl.Log.Error(err)
		return "", err
	}

	// collection total supply
	tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, tokenID)
	if err != nil {
		erc1155, _ := abis.NewERC1155(contractAddress, n.Client)
		uri, err2 := erc1155.Uri(&bind.CallOpts{}, tokenID)

		if err2 != nil {
			gbl.Log.Error(err2)
			return "", err2
		}
		return uri, nil

	}

	gbl.Log.Debugf("GetTokenURI || tokenURI: %+v", tokenURI)

	return tokenURI, nil
}

func (n *Node) GetSupportedStandards(contractAddress common.Address) []standard.Standard {
	ctx := context.Background()

	supportedStandards := make([]standard.Standard, 0)

	// check erc165 interface
	queryiERC165Supported := ethereum.CallMsg{
		To:   &contractAddress,
		Data: supportsInterfaceiERC165,
	}

	resultERC165, err := n.Client.CallContract(ctx, queryiERC165Supported, nil)
	if err != nil {
		gbl.Log.Warnf("interface ERC165 check error: %v", err)
		return nil
	}

	queryiERCFFFFFFSupported := ethereum.CallMsg{
		To:   &contractAddress,
		Data: supportsInterfaceiFFFFFF,
	}

	resultFFFFFF, err := n.Client.CallContract(ctx, queryiERCFFFFFFSupported, nil)
	if err != nil {
		gbl.Log.Warnf("interface FFFFFF check error: %v", err)
		return nil
	}

	if (len(resultERC165)-1 > 0 && len(resultFFFFFF)-1 > 0) && bytes.Equal(resultERC165[len(resultERC165)-1:], []byte{0x01}) && bytes.Equal(resultFFFFFF[len(resultFFFFFF)-1:], []byte{0x00}) {
		supportedStandards = append(supportedStandards, standard.ERC165)
	} else {
		gbl.Log.Warnf("interface ERC165 check error: %v", err)
		return nil
	}

	// get the contractERC1155 ABI
	contractERC1155, err := abis.NewERC1155(contractAddress, n.Client)
	if err != nil {
		gbl.Log.Error(err)
	}

	if iERC1155Supported, err := contractERC1155.SupportsInterface(&bind.CallOpts{}, iERC1155); err != nil {
		gbl.Log.Debug(err)
	} else {
		if iERC1155Supported {
			supportedStandards = append(supportedStandards, standard.ERC1155)
		}
	}

	return supportedStandards
}

func (n *Node) ERC1155Supported(contractAddress common.Address) bool {
	// get the contractERC1155 ABI
	contractERC1155, err := abis.NewERC1155(contractAddress, n.Client)
	if err != nil {
		gbl.Log.Error(err)
	}

	if iERC1155Supported, err := contractERC1155.SupportsInterface(&bind.CallOpts{}, iERC1155); err != nil {
		gbl.Log.Debug(err)
	} else {
		if iERC1155Supported {
			return true
		}
	}

	return false
}

func (n *Node) GetERC1155TokenID(data []byte) *big.Int {
	half := len(data) / 2
	tokenID, _ := strconv.ParseInt(common.Bytes2Hex(bytes.Trim(data[:half], "\x00")), 16, 64)
	// value, _ := strconv.ParseInt(string(common.Bytes2Hex(bytes.Trim(data[half:], "\x00"))), 16, 64)

	return big.NewInt(tokenID)
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

func EtherToWei(ether *big.Float) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)

	return f.Quo(fWei.Set(ether), big.NewFloat(params.Wei))
}

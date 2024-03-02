package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/abis/erc20"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wealdtech/go-ens/v3"
)

type Pool []*Provider

func (p *Pool) Node() *Provider {
	if len(*p) == 0 {
		return nil
	}

	return (*p)[0]
}

// Provider represents a rpc-endpoint Provider configuration.
type Provider struct {
	Name      string `json:"name"      mapstructure:"name"`
	Endpoint  string `json:"endpoint"  mapstructure:"endpoint"`
	Preferred bool   `json:"preferred" mapstructure:"preferred"`

	Client     *ethclient.Client  `json:"-" mapstructure:"-"`
	GethClient *gethclient.Client `json:"-" mapstructure:"-"`

	// Rueidi *rueidica.Rueidica

	LastLogReceivedAt time.Time `json:"-" mapstructure:"-"`
}

func (p *Provider) getTokenURI(contractAddress common.Address, tokenID *big.Int) (string, error) {
	// get the contractERC721 ABIs
	contractERC721, err := p.getERC721ABI(contractAddress)
	if err != nil {
		log.Error(err)

		return "", err
	}

	// get the tokens uri
	tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, tokenID)
	if err != nil {
		erc1155, _ := abis.NewERC1155(contractAddress, p.Client)
		uri, err2 := erc1155.Uri(&bind.CallOpts{}, tokenID)

		if err2 != nil {
			log.Error(err2)

			return "", err2
		}

		tokenURI = uri
	}

	return tokenURI, nil
}

func (p *Provider) GetTokenImageURI(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (string, error) {
	log.Debugf("GetTokenImageURI || contractAddress: %s | tokenID: %d", contractAddress, tokenID)

	tokenURI, err := p.getTokenURI(contractAddress, tokenID)
	if err != nil {
		log.Errorf("get token image uri error: %+v", err.Error())

		return "", err
	}

	if strings.HasPrefix(tokenURI, "data:") {
		tokenURI = strings.TrimPrefix(tokenURI, "data:")
		mimeType, data, _ := strings.Cut(tokenURI, ",")

		switch mimeType {
		case "application/json;base64", "data:application/jsonbase64":
			log.Debugf("🧶 base64 json metadata in uri field: %v", data)

			decoded, err := base64.StdEncoding.DecodeString(data)
			if err != nil {
				log.Warn(err)
				log.Warn("")
				log.Warn(data)
				log.Warn("")
			}

			data = string(decoded)

			log.Printf("🧶 base64 json metadata: %+v", data)

			var metadata *nemo.MetadataERC721
			err = json.Unmarshal([]byte(data), &metadata)
			if err != nil {
				log.Print(err)
			}

			log.Printf("🧶 base64 json metadata721: %+v", metadata)

			return metadata.Image, nil
		}
	}

	log.Debugf("GetTokenImageURI || tokenURI: %+v", tokenURI)

	metadata, err := getTokenMetadata(ctx, tokenURI)
	if err != nil || metadata == nil {
		log.Debugf("get token image uri error: %+v", err)

		return "", err
	}

	log.Debugf("GetTokenImageURI || metadata: %+v", metadata)

	return metadata.Image, nil
}

func (p *Provider) GetERC721CollectionName(contractAddress common.Address) (string, error) {
	// get the contractERC721 ABI
	contractERC721, err := p.getERC721ABI(contractAddress)
	if err != nil {
		log.Error(err)

		return "", err
	}

	if name, err := contractERC721.Name(&bind.CallOpts{}); err == nil {
		log.Debugf("found collection name via erc721 chain call: %s", name)

		return name, nil
	}

	return "", nil
}

func (p *Provider) GetERC721CollectionMetadata(contractAddress common.Address) (map[string]interface{}, error) {
	collectionMetadata := make(map[string]interface{})

	// get the contractERC721 ABIs
	contractERC721, err := p.getERC721ABI(contractAddress)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	if value, err := contractERC721.Name(&bind.CallOpts{}); err != nil {
		log.Debug(err)
	} else {
		collectionMetadata["name"] = value
	}

	if symbol, err := contractERC721.Symbol(&bind.CallOpts{}); err != nil {
		log.Debug(err)
	} else {
		collectionMetadata["symbol"] = symbol
	}

	if totalSupply, err := contractERC721.TotalSupply(&bind.CallOpts{}); err != nil {
		log.Debug(err)
	} else {
		collectionMetadata["totalSupply"] = totalSupply.Uint64()
	}

	if tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, big.NewInt(0)); err != nil {
		log.Debug(err)
	} else {
		collectionMetadata["tokenURI"] = tokenURI
	}

	return collectionMetadata, nil
}

func (p *Provider) GetERC1155TotalSupply(address common.Address, tokenID *big.Int) (*big.Int, error) {
	// bind erc1155 abi
	if contractERC1155, err := abis.NewERC1155(address, p.Client); err == nil {
		// call totalSupply
		if totalSupply, err := contractERC1155.TotalSupply(&bind.CallOpts{}, tokenID); err == nil {
			return totalSupply, nil
		}
	}

	return nil, errors.New("could not get total supply")
}

func (p *Provider) GetERC1155TokenName(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (string, error) {
	//
	// ERC1155
	contractERC1155, err := abis.NewERC1155(contractAddress, p.Client)
	if err != nil {
		log.Error(err)

		return "", err
	}

	if tokenID == nil {
		tokenID = big.NewInt(1)
	}

	// unfortunately theres no .Name() method on the ERC1155 abi spec,
	// see https://docs.openzeppelin.com/contracts/3.x/api/token/erc1155
	// the "name" is actually a wild hack by us... 🙄 but it works surprisingly well...^^
	if uri, err := contractERC1155.Uri(&bind.CallOpts{}, tokenID); err == nil {
		log.Debugf("found collection uri via erc1155 chain call: %s", uri)

		if strings.HasPrefix(uri, "data:") {
			uri = strings.TrimPrefix(uri, "data:")
			mimeType, data, _ := strings.Cut(uri, ",")

			switch mimeType {
			case "application/json;base64", "data:application/jsonbase64":
				log.Debugf("🧶 base64 json metadata in uri field: %v", data)

				decoded, err := base64.StdEncoding.DecodeString(data)
				if err != nil {
					log.Warn(err)
					log.Warn("")
					log.Warn(data)
					log.Warn("")
				}

				log.Debugf("🧶 base64 json metadata: %+v", string(decoded))

				data = string(decoded)

				fallthrough

			case "application/json;utf8":
				log.Debugf("🧶 json metadata in uri field: %v", data)

				var metadata map[string]interface{}

				err := json.Unmarshal([]byte(data), &metadata)
				if err != nil {
					log.Warn(err)
					log.Warn("")
					log.Warn(data)
					log.Warn("")
				}

				log.Debugf("🧶 json metadata: %+v", metadata)

				if name, ok := metadata["name"]; ok {
					tokenName, ok := name.(string)
					if !ok {
						log.Warnf("🧶 json metadata name is not a string: %v", name)
					}

					return tokenName, nil
				}

			default:
				log.Infof("🧶 metadata in uri field: %v | %v", mimeType, data)
			}
		}

		if metadata, err := external.GetERC1155MetadataForURI(ctx, uri, tokenID); err == nil && metadata != nil {
			name := strings.TrimRight(metadata.Name, "#0123456789")

			if metadata.CreatedBy != "" {
				name = metadata.CreatedBy + " | " + name
			}

			log.Debugf("found collection name via erc1155 metadata: %v", name)

			return name, nil
		}
	}

	return "", errors.New("could not find collection name")
}

// Connect tries to Connect to the provider and returns an error if it fails.
func (p *Provider) Connect() error {
	var err error

	// reconnect handling
	if p.Client != nil {
		p.Client.Close()

		p.Client = nil
	}

	var rpcClient *rpc.Client

	if strings.HasPrefix(p.Endpoint, "unix://") {
		p.Endpoint = strings.TrimPrefix(p.Endpoint, "unix://")

		rpcClient, err = rpc.DialIPC(context.Background(), p.Endpoint)
		if err != nil {
			log.Debugf("Failed to connect to node %s: %s", p.Name, err)

			return err
		}
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), p.Endpoint)
		if err != nil {
			log.Debugf("Failed to connect to node %s: %s", p.Name, err)

			return err
		}
	}

	ethClient := ethclient.NewClient(rpcClient)
	if ethClient == nil {
		log.Debugf("Failed to start eth client for node %s: %s", p.Name, err)

		return err
	}

	syncing, err := ethClient.SyncProgress(context.Background())
	if err != nil {
		log.Debugf("Failed to get sync progress for node %s: %s", p.Name, err)

		return err
	}

	if syncing != nil {
		log.Debugf("⏳ node %s is still syncing...", p.Name)

		return errors.New("node is still syncing")
	}

	p.Client = ethClient

	if p.Preferred {
		log.Debugf("%s syncing: %v", p.Name, syncing)

		p.GethClient = gethclient.New(rpcClient)
	}

	return err
}

func (p *Provider) SubscribeToAllTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	subscribeTopics := [][]common.Hash{
		{common.HexToHash(string(topic.Transfer)), common.HexToHash(string(topic.TransferSingle)), common.HexToHash(string(topic.BuyPriceSet))},
		{},
		{},
		{},
	}

	return p.subscribeTo(queueLogs, subscribeTopics, nil)
}

func (p *Provider) subscribeTo(queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
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

// IsContract returns true if the given address is a contract address.
// to resource intensive to check this for every address we encounter, so we cache the result.
func (p *Provider) IsContract(address common.Address, rueidi *rueidica.Rueidica) bool {
	// if its a marketplace address, its a contract
	if marketplace.Addresses().Contains(address) {
		return true
	}

	// check if we have a cached the account type already
	accountType, err := rueidi.GetCachedAccountType(context.Background(), address)
	if err == nil {
		return degendb.AccountType(accountType) == degendb.ContractAccount
	}

	log.Debugf("❕ error getting cached account type: %s", err)

	// ok 🙄 seems we really need to check via a node if its a eoa or contract
	codeAt, err := p.codeAt(context.Background(), address)
	if err != nil {
		log.Debugf("❕ failed to get codeAt for %s: %s", address.String(), err)

		return false
	}

	log.Debugf("codeAt(%s): %+v", address.Hex(), codeAt)

	// if there is deployed code at the address, it's a contract
	return len(codeAt) > 0
}

func (p *Provider) getERC721ABI(contractAddress common.Address) (*abis.ERC721v3, error) {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return contractERC721, nil
}

func (p *Provider) GetERC1155ABI(contractAddress common.Address) (*abis.ERC1155, error) {
	// get the contractERC721 ABIs
	contractERC1155, err := abis.NewERC1155(contractAddress, p.Client)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return contractERC1155, nil
}

func (p *Provider) GetWETHABI(contractAddress common.Address) (*abis.WETH, error) {
	// get the contractERC721 ABIs
	contractWETH, err := abis.NewWETH(contractAddress, p.Client)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return contractWETH, nil
}

func (p *Provider) getERC20ABI(contractAddress common.Address) (*erc20.ERC20, error) {
	// get the contractERC721 ABIs
	contractERC20, err := erc20.NewERC20(contractAddress, p.Client)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return contractERC20, nil
}

func getTokenMetadata(ctx context.Context, tokenURI string) (*nemo.MetadataERC721, error) {
	log.Debugf("GetTokenMetadata || tokenURI: %+v", tokenURI)

	tokenURI = utils.PrepareURL(tokenURI)

	response, err := utils.HTTP.GetWithTLS12(ctx, tokenURI)
	if err != nil || response.StatusCode != http.StatusOK {
		status := "unknown"
		if response != nil {
			status = response.Status
		}

		log.Warnf("❌ get token metadata | %s | status: %s | error: %+v", tokenURI, status, err)

		return nil, err
	}

	log.Debugf("get token metadata status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var tokenMetadata *nemo.MetadataERC721

	responseBody, err := io.ReadAll(response.Body)

	// decode the data
	if err != nil || !json.Valid(responseBody) {
		log.Warnf("get token metadata invalid json: %s", err)

		return nil, err
	}

	// decode the data
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&tokenMetadata); err != nil {
		log.Warnf("get token metadata decode error: %s", err.Error())

		return nil, err
	}

	log.Debugf("GetTokenMetadata || tokenMetadata: %+v", tokenMetadata)

	return tokenMetadata, nil
}

//
// ens
//

func (p *Provider) ENSLookup(ensName string) (common.Address, error) {
	var err error

	// do a lookup for the ensName to validate its authenticity
	resolvedAddress, err := ens.Resolve(p.Client, ensName)
	if err != nil {
		log.Debugf("ens resolve error: %s : %s", ensName, err)

		return common.Address{}, err
	}

	return resolvedAddress, nil
}

func (p *Provider) ReverseLookupAndValidate(address common.Address) (string, error) {
	var ensName string

	var err error

	// lookup the ens ensName for an address
	ensName, err = ens.ReverseResolve(p.Client, address)

	if err != nil || common.IsHexAddress(ensName) {
		log.Debugf("ens reverse resolve error: %s -> %s: %s", address, ensName, err)

		return "", err
	}

	// do a lookup for the ensName to validate its authenticity
	resolvedAddress, err := ens.Resolve(p.Client, ensName)
	if err != nil {
		log.Debugf("ens resolve error: %s -> %s: %s", ensName, address, err)

		return "", err
	}

	if resolvedAddress != address {
		log.Debugf("  %s  !=  %s", resolvedAddress.Hex(), address.Hex())

		return "", errors.New("ens forward and reverse resolved addresses do not match")
	}

	return ensName, nil
}

//
// gas
//

// // GetCurrentGasInfo returns the current gas price and tip.
// func (p *Provider) GetGasInfo(ctx context.Context) (*nemo.GasInfo, error) {
// 	// header, err := p.Client.BlockByNumber(ctx, nil)
// 	header, err := p.Client.HeaderByNumber(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	gasPrice, err := p.Client.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	gasTip, err := p.Client.SuggestGasTipCap(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &nemo.GasInfo{
// 		LastBlock:         header.Number.Uint64(),
// 		LastBlockGasLimit: header.GasLimit,
// 		GasPriceWei:       gasPrice,
// 		GasTipWei:         gasTip,
// 	}, nil
// }

//
// bytecode
//

// codeAt returns the current gas price and tip.
func (p *Provider) codeAt(ctx context.Context, address common.Address) ([]byte, error) {
	return p.Client.CodeAt(ctx, address, nil) // nil is latest block
}

//
// nonce
//

// codeAt returns the current gas price and tip.
func (p *Provider) nonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return p.Client.NonceAt(ctx, address, nil)
}

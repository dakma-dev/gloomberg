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

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wealdtech/go-ens/v3"
)

// provider represents a rpc-endpoint provider configuration.
type provider struct {
	Name      string `json:"name" mapstructure:"name"`
	Endpoint  string `json:"endpoint" mapstructure:"endpoint"`
	Preferred bool   `json:"preferred" mapstructure:"preferred"`

	Color lipgloss.Color `json:"color" mapstructure:"color"`
	// Marker string         `json:"marker" mapstructure:"marker"`

	PID common.Hash `json:"pid" mapstructure:"pid"`

	Client     *ethclient.Client  `json:"-" mapstructure:"-"`
	GethClient *gethclient.Client `json:"-" mapstructure:"-"`
}

// // newProvider creates a new provider.
// func newProvider(name, endpoint string, preferred bool) *provider {
// 	return &provider{
// 		Name:      name,
// 		Endpoint:  endpoint,
// 		Preferred: preferred,
// 	}
// }

// ID returns the provider id consisting of the first and last 4 characters of the pid in a human readable format.
func (p *provider) ID() string {
	// remove leading zeroes from the pid
	pid := common.TrimLeftZeroes(p.PID.Bytes())

	return string(pid[0:4]) + "…" + string(pid[28:32])
}

func (p *provider) getTokenURI(contractAddress common.Address, tokenID *big.Int) (string, error) {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return "", err
	}

	// get the tokens uri
	tokenURI, err := contractERC721.TokenURI(&bind.CallOpts{}, tokenID)
	if err != nil {
		erc1155, _ := abis.NewERC1155(contractAddress, p.Client)
		uri, err2 := erc1155.Uri(&bind.CallOpts{}, tokenID)

		if err2 != nil {
			gbl.Log.Error(err2)

			return "", err2
		}

		tokenURI = uri
	}

	return tokenURI, nil
}

func (p *provider) getTokenImageURI(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (string, error) {
	gbl.Log.Debugf("GetTokenImageURI || contractAddress: %s | tokenID: %d", contractAddress, tokenID)

	tokenURI, err := p.getTokenURI(contractAddress, tokenID)
	if err != nil {
		gbl.Log.Errorf("get token image uri error: %+v", err.Error())

		return "", err
	}

	gbl.Log.Debugf("GetTokenImageURI || tokenURI: %+v", tokenURI)

	metadata, err := getTokenMetadata(ctx, tokenURI)
	if err != nil || metadata == nil {
		gbl.Log.Debugf("get token image uri error: %+v", err)

		return "", err
	}

	gbl.Log.Debugf("GetTokenImageURI || metadata: %+v", metadata)

	return metadata.Image, nil
}

func (p *provider) getERC721CollectionName(ctx context.Context, contractAddress common.Address) (string, error) {
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

func (p *provider) getERC721CollectionMetadata(ctx context.Context, contractAddress common.Address) (map[string]interface{}, error) {
	collectionMetadata := make(map[string]interface{}, 0)

	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return nil, err
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

	return collectionMetadata, nil
}

func (p *provider) getERC1155TokenName(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (string, error) {
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

	// unfortunately theres no .Name() method on the ERC1155 abi spec,
	// see https://docs.openzeppelin.com/contracts/3.x/api/token/erc1155
	// the "name" is actually a wild hack by us... 🙄 but it works surprisingly well...^^
	if uri, err := contractERC1155.Uri(&bind.CallOpts{}, tokenID); err == nil {
		gbl.Log.Debugf("found collection uri via erc1155 chain call: %s", uri)

		if strings.HasPrefix(uri, "data:") {
			uri = strings.TrimPrefix(uri, "data:")
			mimeType, data, _ := strings.Cut(uri, ",")

			switch mimeType {
			case "application/json;base64":
				gbl.Log.Debugf("🧶 base64 json metadata in uri field: %v", data)

				decoded, err := base64.StdEncoding.DecodeString(data)
				if err != nil {
					gbl.Log.Warn(err)
					gbl.Log.Warn("")
					gbl.Log.Warn(data)
					gbl.Log.Warn("")
				}

				gbl.Log.Debugf("🧶 base64 json metadata: %+v", string(decoded))

				data = string(decoded)

				fallthrough

			case "application/json;utf8":
				gbl.Log.Debugf("🧶 json metadata in uri field: %v", data)

				var metadata map[string]interface{}
				err := json.Unmarshal([]byte(data), &metadata)
				if err != nil {
					gbl.Log.Warn(err)
					gbl.Log.Warn("")
					gbl.Log.Warn(data)
					gbl.Log.Warn("")
				}

				gbl.Log.Debugf("🧶 json metadata: %+v", metadata)

				if name, ok := metadata["name"]; ok {
					tokenName, ok := name.(string)
					if !ok {
						gbl.Log.Warnf("🧶 json metadata name is not a string: %v", name)
					}

					return tokenName, nil
				}

			default:
				gbl.Log.Infof("🧶 metadata in uri field: %v | %v", mimeType, data)
			}
		}

		if metadata, err := external.GetERC1155MetadataForURI(ctx, uri, tokenID); err == nil && metadata != nil {
			name := strings.TrimRight(metadata.Name, "#0123456789")

			if metadata.CreatedBy != "" {
				name = metadata.CreatedBy + " | " + name
			}

			gbl.Log.Debugf("found collection name via erc1155 metadata: %v", name)

			return name, nil
		}
	}

	return "", errors.New("could not find collection name")
}

// connect tries to connect to the provider and returns an error if it fails.
func (p *provider) connect() error {
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
			gbl.Log.Debugf("Failed to connect to node %s: %s", p.Name, err)

			return err
		}
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), p.Endpoint)
		if err != nil {
			gbl.Log.Debugf("Failed to connect to node %s: %s", p.Name, err)

			return err
		}
	}

	ethClient := ethclient.NewClient(rpcClient)
	if ethClient == nil {
		gbl.Log.Debugf("Failed to start eth client for node %s: %s", p.Name, err)

		return err
	}

	syncing, err := ethClient.SyncProgress(context.Background())
	if err != nil {
		gbl.Log.Debugf("Failed to get sync progress for node %s: %s", p.Name, err)

		return err
	}

	if syncing != nil {
		gbl.Log.Debugf("⏳ node %s is still syncing...", p.Name)

		return errors.New("node is still syncing")
	}

	p.Client = ethClient

	if p.Preferred {
		gbl.Log.Infof("%s syncing: %v", p.Name, syncing)

		p.GethClient = gethclient.New(rpcClient)
	}

	return err
}

func (p *provider) subscribeToAllTransfers(queueLogs chan types.Log) (ethereum.Subscription, error) {
	return p.subscribeTo(queueLogs, [][]common.Hash{{common.HexToHash(string(topic.Transfer)), common.HexToHash(string(topic.TransferSingle))}, {}, {}, {}}, nil)
}

func (p *provider) subscribeTo(queueLogs chan types.Log, topics [][]common.Hash, contractAddresses []common.Address) (ethereum.Subscription, error) {
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

func (p *provider) getERC721ABI(ctx context.Context, contractAddress common.Address) (*abis.ERC721v3, error) {
	// get the contractERC721 ABIs
	contractERC721, err := abis.NewERC721v3(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return nil, err
	}

	return contractERC721, nil
}

func (p *provider) getERC1155ABI(ctx context.Context, contractAddress common.Address) (*abis.ERC1155, error) {
	// get the contractERC721 ABIs
	contractERC1155, err := abis.NewERC1155(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return nil, err
	}

	return contractERC1155, nil
}

func (p *provider) getWETHABI(ctx context.Context, contractAddress common.Address) (*abis.WETH, error) {
	// get the contractERC721 ABIs
	contractWETH, err := abis.NewWETH(contractAddress, p.Client)
	if err != nil {
		gbl.Log.Error(err)

		return nil, err
	}

	return contractWETH, nil
}

func getTokenMetadata(ctx context.Context, tokenURI string) (*nemo.MetadataERC721, error) {
	gbl.Log.Debugf("GetTokenMetadata || tokenURI: %+v", tokenURI)

	tokenURI = utils.PrepareURL(tokenURI)

	response, err := utils.HTTP.GetWithTLS12(ctx, tokenURI)
	if err != nil || response.StatusCode != http.StatusOK {
		status := "unknown"
		if response != nil {
			status = response.Status
		}

		gbl.Log.Warnf("❌ get token metadata | %s | status: %s | error: %+v", tokenURI, status, err)

		return nil, err
	}

	gbl.Log.Debugf("get token metadata status: %s", response.Status)

	defer response.Body.Close()

	// create a variable of the same type as our model
	var tokenMetadata *nemo.MetadataERC721

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

//
// ens
//

// func (p *provider) getENSForAllAddresses(wallets *wallet.Wallets) {
// 	var wgENS sync.WaitGroup

// 	for _, w := range *wallets {
// 		wgENS.Add(1)

// 		go func(w *wallet.Wallet) {
// 			defer wgENS.Done()

// 			name, err := p.getENSForAddress(w.Address)
// 			if err != nil {
// 				gbl.Log.Debugf("❌ failed to resolve ENS name for %s: %s", w.Address.Hex(), err)

// 				return
// 			}

// 			w.ENS = &ens.Name{
// 				Name: name,
// 			}
// 			w.Name = w.ENS.Name
// 		}(w)
// 	}

// 	wgENS.Wait()
// }

func (p *provider) getENSForAddress(ctx context.Context, address common.Address) (string, error) {
	var ensName string

	if cachedName, err := cache.GetENSName(ctx, address); err == nil && cachedName != "" {
		gbl.Log.Debugf("ens ensName for address %s is cached", address.Hex())

		return cachedName, nil
	}

	resolvedName, err := p.reverseLookupAndValidate(address)
	if err != nil {
		gbl.Log.Debugf("ens reverse lookup failed for address %s: %s", address.Hex(), err)

		return "", err
	}

	if resolvedName == "" {
		gbl.Log.Debugf("address %s has no associated ens ensName", address.Hex())
	}

	ensName = resolvedName

	cache.StoreENSName(ctx, address, ensName)

	return ensName, nil
}

func (p *provider) reverseLookupAndValidate(address common.Address) (string, error) {
	var ensName string

	var err error

	// lookup the ens ensName for an address
	ensName, err = ens.ReverseResolve(p.Client, address)

	if err != nil || common.IsHexAddress(ensName) {
		gbl.Log.Debugf("ens reverse resolve error: %s -> %s: %s", address, ensName, err)

		return "", err
	}

	// do a lookup for the ensName to validate its authenticity
	resolvedAddress, err := ens.Resolve(p.Client, ensName)
	if err != nil {
		gbl.Log.Debugf("ens resolve error: %s -> %s: %s", ensName, address, err)

		return "", err
	}

	if resolvedAddress != address {
		gbl.Log.Debugf("  %s  !=  %s", resolvedAddress.Hex(), address.Hex())

		return "", errors.New("ens forward and reverse resolved addresses do not match")
	}

	return ensName, nil
}

//
// gas
//

// GetCurrentGasInfo returns the current gas price and tip.
func (p *provider) getGasInfo(ctx context.Context) (*nemo.GasInfo, error) {
	// header, err := p.Client.BlockByNumber(ctx, nil)
	header, err := p.Client.HeaderByNumber(ctx, nil)
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

	return &nemo.GasInfo{
		LastBlock:         header.Number.Uint64(),
		LastBlockGasLimit: header.GasLimit,
		GasPriceWei:       gasPrice,
		GasTipWei:         gasTip,
	}, nil
}

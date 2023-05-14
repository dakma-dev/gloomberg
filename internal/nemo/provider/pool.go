package provider

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/abis/erc20"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Pool struct {
	LastLogReceivedAt time.Time `json:"-" mapstructure:"-"`

	providers []*Provider

	queueLogs chan types.Log

	Rueidi *rueidica.Rueidica

	// gb *gloomberg.Gloomberg `json:"-" mapstructure:"-"`
}

func (pp *Pool) GetProviders() []*Provider {
	return pp.providers
}

type methodCall string

const (
	BlockNumber        methodCall = "eth_blockNumber"
	TransactionByHash  methodCall = "eth_getTransactionByHash"
	TransactionReceipt methodCall = "eth_getTransactionReceipt"

	TokenImageURI methodCall = "token_image_uri" //nolint:gosec

	ERC721CollectionName     methodCall = "erc721_collection_name"
	ERC721CollectionMetadata methodCall = "erc721_collection_metadata"

	ERC1155TokenName   methodCall = "erc1155_token_name" //nolint:gosec
	ERC1155TotalSupply methodCall = "erc1155_total_supply"

	ReverseResolveENS methodCall = "resolve_ens_address"
	ResolveENS        methodCall = "resolve_ens"

	GasInfo methodCall = "gas_info"
)

type methodCallParams struct {
	TxHash  common.Hash    `json:"hash"`
	Address common.Address `json:"contract_address"`
	TokenID *big.Int       `json:"token_id"`
	EnsName string         `json:"ens_name"`
}

var callMethodCounter uint64

func FromConfig(config interface{}) (*Pool, error) {
	var rawPool []*Provider

	providerPool := &Pool{
		providers: make([]*Provider, 0),
	}

	// spinner
	providerSpinner := style.GetSpinner("setting up the provider connections...")
	_ = providerSpinner.Start()

	config, ok := config.([]interface{})
	if !ok {
		gbl.Log.Warnf("reading provider configuration failed: %+v", config)

		return nil, errors.New("invalid provider configuration")
	}

	//
	// decode the config into a raw node pool
	decodeHooks := mapstructure.ComposeDecodeHookFunc(
		hooks.StringToAddressHookFunc(),
		hooks.StringToLipglossColorHookFunc(),
	)

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: decodeHooks,
		Result:     &rawPool,
	})

	err := decoder.Decode(config)
	if err != nil {
		infoMsg := fmt.Sprintf("❌ ⚙️ reading provider configuration failed: %+v", err)
		gbl.Log.Info(infoMsg)
		fmt.Println("\n" + infoMsg + "\n")

		return nil, err
	}

	//
	// initialize the providers and connect to the endpoints
	for _, provider := range rawPool {
		// hash the endpoint to get a unique id for the provider
		provider.PID = common.BytesToHash([]byte(provider.Endpoint))

		// connect to the endpoint
		if err := provider.connect(); err != nil {
			gbl.Log.Warnf("❔ not adding %s: %s", style.BoldStyle.Render(provider.Name), err)

			continue
		}

		gbl.Log.Infof("✅ added node %s", style.BoldStyle.Render(provider.Name))
		providerPool.providers = append(providerPool.providers, provider)
	}

	// handle reconnects
	go func() {
		// reconnect if no logs received for a while
		maxDelay := internal.BlockTime * 3
		reconnectTicker := time.NewTicker(maxDelay)

		for range reconnectTicker.C {
			// if last log is older than maxDelay, we reconnect
			if !providerPool.LastLogReceivedAt.IsZero() && providerPool.LastLogReceivedAt.Add(maxDelay).Before(time.Now()) {
				infoMsg := fmt.Sprintf(" ❗️ no logs received for %s, reconnecting...", style.Bold(fmt.Sprintf("%.0fsec", maxDelay.Seconds())))
				gbl.Log.Info(infoMsg)
				fmt.Print("\n" + infoMsg + "\n")

				providerPool.ReconnectProviders()

				reconnectTicker.Reset(maxDelay)
			}
		}
	}()

	// get all node names to be shown as a list of connected nodes
	nodeNames := make([]string, 0)
	for _, n := range providerPool.providers {
		nodeNames = append(nodeNames, style.BoldStyle.Render(n.Name))
	}

	// spinner
	providerSpinner.StopMessage(
		fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(providerPool.providers))), " nodes connected: ", strings.Join(nodeNames, ", ")) + "\n",
	)

	_ = providerSpinner.Stop()

	return providerPool, nil
}

// func (pp *Pool) ReconnectProviders(queueLogs *chan types.Log) {.
func (pp *Pool) ReconnectProviders() {
	gbl.Log.Info("🔌 trying to re-connect...")

	// compatibility with old config key
	var providerConfig interface{}
	if cfg := viper.Get("provider"); cfg != nil {
		providerConfig = cfg
	} else {
		providerConfig = viper.Get("nodes")
	}

	if pp.queueLogs == nil {
		gbl.Log.Fatal("❌ queueLogs is nil - can npt re-subscribe, exiting")

		return
	}

	// store the current queueLogs channel
	queueLogs := pp.queueLogs

	// reconnect to the providers
	if pool, err := FromConfig(providerConfig); err != nil {
		gbl.Log.Fatal("❌ running provider failed, exiting")
	} else if pool != nil {
		pp = pool
	}

	// restore the queueLogs channel
	pp.queueLogs = queueLogs

	// re-subscribe
	if _, err := pp.Subscribe(pp.queueLogs); err != nil {
		gbl.Log.Fatalf("❌ subscribing to logs failed: %s", err)

		return
	}
}

func (pp *Pool) PreferredProviderAvailable() bool {
	return len(pp.getPreferredProviders()) > 0
}

func (pp *Pool) GetLogsByBlockNumber(blockNumber int64) []types.Log {
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
	}

	for _, provider := range pp.getProviders() {
		if logs, err := provider.Client.FilterLogs(context.Background(), filterQuery); err == nil {
			return logs
		}
	}

	return nil
}

func (pp *Pool) Subscribe(queueLogs chan types.Log) (uint64, error) {
	if queueLogs == nil {
		return 0, errors.New("queueLogs channel is nil")
	}

	// store channel for later use/reconnects
	pp.queueLogs = queueLogs

	// subscribe
	availableProvider := pp.getProviders()
	if len(pp.getPreferredProviders()) > 0 {
		availableProvider = pp.getPreferredProviders()
	}

	subscribedTo := uint64(0)

	for _, provider := range availableProvider {
		// subscribe to all logs with "Tranfer" or "TransferSingle" as first topic
		if _, err := provider.subscribeToAllTransfers(pp.queueLogs); err != nil {
			gbl.Log.Warnf("subscribe to topic TransferSingle via node %d failed: %s", provider.Name, err)
		} else {
			subscribedTo++
			gbl.Log.Infof("✍️ subscribed to all transfer topics via node %s", style.Bold(provider.Name))
		}
	}

	if subscribedTo == 0 {
		return 0, errors.New("no provider available")
	}

	return subscribedTo, nil
}

func (pp *Pool) SubscribeToEverything(queueLogs chan types.Log) (uint64, error) {
	if queueLogs == nil {
		return 0, errors.New("queueLogs channel is nil")
	}

	// store channel for later use/reconnects
	pp.queueLogs = queueLogs

	// subscribe
	availableProvider := pp.getProviders()
	if len(pp.getPreferredProviders()) > 0 {
		availableProvider = pp.getPreferredProviders()
	}

	subscribedTo := uint64(0)

	for _, provider := range availableProvider {
		// subscribe to all logs with "Tranfer" or "TransferSingle" as first topic
		if _, err := provider.subscribeTo(pp.queueLogs, [][]common.Hash{}, []common.Address{}); err != nil {
			gbl.Log.Warnf("subscribe to everything via node %d failed: %s", provider.Name, err)
		} else {
			subscribedTo++
			gbl.Log.Infof("✍️ subscribed to everything via node %s", style.Bold(provider.Name))
		}
	}

	if subscribedTo == 0 {
		return 0, errors.New("no provider available")
	}

	return subscribedTo, nil
}

func (pp *Pool) SubscribeToAddresses(queueLogs chan types.Log, addresses []common.Address) (uint64, error) {
	if queueLogs == nil {
		return 0, errors.New("queueLogs channel is nil")
	}

	// store channel for later use/reconnects
	pp.queueLogs = queueLogs

	// subscribe
	availableProvider := pp.getProviders()
	if len(pp.getPreferredProviders()) > 0 {
		availableProvider = pp.getPreferredProviders()
	}

	subscribedTo := uint64(0)

	for _, provider := range availableProvider {
		// subscribe to all logs with "Tranfer" or "TransferSingle" as first topic
		if _, err := provider.subscribeTo(pp.queueLogs, [][]common.Hash{}, addresses); err != nil {
			gbl.Log.Warnf("subscribe to addresses failed: %v | %v", addresses, err)
		} else {
			subscribedTo++
			gbl.Log.Infof("✍️ subscribed to address %v", addresses)
		}
	}

	if subscribedTo == 0 {
		return 0, errors.New("no provider available")
	}

	return subscribedTo, nil
}

func (pp *Pool) SubscribeToTopics(queueLogs chan types.Log, topics [][]common.Hash) (uint64, error) {
	if queueLogs == nil {
		return 0, errors.New("queueLogs channel is nil")
	}

	// store channel for later use/reconnects
	pp.queueLogs = queueLogs

	// subscribe
	availableProvider := pp.getProviders()
	if len(pp.getPreferredProviders()) > 0 {
		availableProvider = pp.getPreferredProviders()
	}

	subscribedTo := uint64(0)

	for _, provider := range availableProvider {
		// subscribe to all logs with "Tranfer" or "TransferSingle" as first topic
		if _, err := provider.subscribeTo(pp.queueLogs, topics, nil); err != nil {
			gbl.Log.Warnf("subscribe to topic TransferSingle via node %d failed: %s", provider.Name, err)
		} else {
			subscribedTo++
			gbl.Log.Infof("✍️ subscribed to all transfer topics via node %s", style.Bold(provider.Name))
		}
	}

	if subscribedTo == 0 {
		return 0, errors.New("no provider available")
	}

	return subscribedTo, nil
}

func (pp *Pool) SubscribeToEverythingPending(queuePendingTx chan *types.Transaction) (uint64, error) {
	// subscribe
	availableProvider := pp.getProviders()
	if len(pp.getPreferredProviders()) > 0 {
		availableProvider = pp.getPreferredProviders()
	}

	subscribedTo := uint64(0)

	for _, provider := range availableProvider {
		// subscribe to all logs with "Tranfer" or "TransferSingle" as first topic
		if _, err := provider.GethClient.SubscribeFullPendingTransactions(context.TODO(), queuePendingTx); err != nil {
			gbl.Log.Warnf("subscribe to pending transactions via node %d failed: %s", provider.Name, err)
		} else {
			subscribedTo++
			gbl.Log.Infof("✍️ subscribed to pending transactions via node %s", style.Bold(provider.Name))
		}
	}

	if subscribedTo == 0 {
		return 0, errors.New("no provider available")
	}

	return subscribedTo, nil
}

func (pp *Pool) getPreferredProviders() []*Provider {
	if pp.providers != nil && len(pp.providers) == 0 {
		return nil
	}

	providers := make([]*Provider, 0, len(pp.providers))

	for _, provider := range pp.providers {
		if provider.Preferred {
			providers = append(providers, provider)
		}
	}

	return providers
}

func (pp *Pool) getProviders() []*Provider {
	providers := make([]*Provider, 0)

	// get all provider
	providers = append(providers, pp.providers...)

	// shuffle provider to avoid hitting the same node over and over again
	rand.Shuffle(len(providers), func(i, j int) {
		providers[i], providers[j] = providers[j], providers[i]
	})

	// prefer preferred (formerly 'local') providers if available
	preferredProviders := make([]*Provider, 0)
	if prefProviders := pp.getPreferredProviders(); len(prefProviders) > 0 {
		preferredProviders = append(preferredProviders, prefProviders...)
	}

	return append(preferredProviders, providers...)
}

func (pp *Pool) GetWETHABI(contractAddress common.Address) (*abis.WETH, error) {
	for _, provider := range pp.getProviders() {
		if wethABI, err := provider.getWETHABI(contractAddress); err == nil {
			return wethABI, nil
		}
	}

	return nil, errors.New("no provider available")
}

func (pp *Pool) GeERC20ABI(contractAddress common.Address) (*erc20.ERC20, error) {
	for _, provider := range pp.getProviders() {
		if erc20ABI, err := provider.getERC20ABI(contractAddress); err == nil {
			return erc20ABI, nil
		}
	}

	return nil, errors.New("no provider available")
}

func (pp *Pool) GetERC1155ABI(contractAddress common.Address) (*abis.ERC1155, error) {
	for _, provider := range pp.getProviders() {
		if erc1155ABI, err := provider.getERC1155ABI(contractAddress); err == nil {
			return erc1155ABI, nil
		}
	}

	return nil, errors.New("no provider available")
}

func (pp *Pool) callMethod(ctx context.Context, method methodCall, params methodCallParams) (interface{}, error) {
	var err error

	atomic.AddUint64(&callMethodCounter, 1)

	if callMethodCounter%100 == 0 {
		gbl.Log.Debugf("callMethodCounter: %d", callMethodCounter)
	}

	for _, provider := range pp.getProviders() {
		switch method {
		case TransactionByHash:
			if params.TxHash == (common.Hash{}) {
				return nil, errors.New("invalid transaction hash")
			}

			if tx, _, err := provider.Client.TransactionByHash(ctx, params.TxHash); err == nil {
				return tx, nil
			}

		case BlockNumber:
			if blockNumber, err := provider.Client.BlockNumber(ctx); err == nil {
				return blockNumber, nil
			}

		case TransactionReceipt:
			if params.TxHash == (common.Hash{}) {
				return nil, errors.New("invalid transaction hash")
			}

			if receipt, err := provider.Client.TransactionReceipt(ctx, params.TxHash); err == nil {
				return receipt, nil
			}

		case TokenImageURI:
			if params.Address == (common.Address{}) || params.TokenID == nil {
				return nil, errors.New("invalid contract address or token id")
			}

			if uri, err := provider.getTokenImageURI(ctx, params.Address, params.TokenID); err == nil {
				return uri, nil
			}

		case ERC721CollectionName:
			if params.Address == (common.Address{}) {
				return nil, errors.New("invalid contract address")
			}

			if collectionName, err := provider.getERC721CollectionName(params.Address); err == nil {
				return collectionName, nil
			}

		case ERC721CollectionMetadata:
			if params.Address == (common.Address{}) {
				return nil, errors.New("invalid contract address")
			}

			if metadata, err := provider.getERC721CollectionMetadata(params.Address); err == nil {
				return metadata, nil
			}

		case ERC1155TokenName:
			if params.Address == (common.Address{}) || params.TokenID == nil {
				return nil, errors.New("invalid contract address or token id")
			}

			if tokenName, err := provider.getERC1155TokenName(ctx, params.Address, params.TokenID); err == nil {
				return tokenName, nil
			}

		case ERC1155TotalSupply:
			if params.Address == (common.Address{}) || params.TokenID == nil {
				return nil, errors.New("invalid contract address or token id")
			}

			// bind erc1155 abi
			if contractERC1155, err := abis.NewERC1155(params.Address, provider.Client); err == nil {
				// call totalSupply
				if totalSupply, err := contractERC1155.TotalSupply(&bind.CallOpts{}, params.TokenID); err == nil {
					return totalSupply, nil
				}
			}

		case ReverseResolveENS:
			if params.Address == (common.Address{}) {
				return nil, errors.New("invalid contract address")
			}

			if ensAddress, err := pp.Rueidi.GetCachedENSName(ctx, params.Address); err == nil {
				return ensAddress, nil
			}

			if ensAddress, err := provider.reverseLookupAndValidate(params.Address); err == nil {
				return ensAddress, nil
			}

		case ResolveENS:
			if ensAddress, err := provider.ensLookup(params.EnsName); err == nil {
				return ensAddress, nil
			}

		case GasInfo:
			if gasInfo, err := provider.getGasInfo(ctx); err == nil {
				return gasInfo, nil
			}

		default:
			return nil, errors.New("invalid method")
		}
	}

	return nil, err
}

// BlockNumber returns the most recent block number.
func (pp *Pool) BlockNumber(ctx context.Context) (uint64, error) {
	num, err := pp.callMethod(ctx, BlockNumber, methodCallParams{})
	if blockNum, ok := num.(uint64); err == nil && ok {
		return blockNum, nil
	}

	return 0, err
}

// TransactionByHash returns the transaction for the given hash.
func (pp *Pool) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, error) {
	tx, err := pp.callMethod(ctx, TransactionByHash, methodCallParams{TxHash: txHash})
	if transaction, ok := tx.(*types.Transaction); err == nil && ok {
		return transaction, nil
	}

	return nil, err
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
func (pp *Pool) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	r, err := pp.callMethod(ctx, TransactionReceipt, methodCallParams{TxHash: txHash})
	if receipt, ok := r.(*types.Receipt); err == nil && ok {
		return receipt, nil
	}

	return nil, err
}

func (pp *Pool) GetTokenImageURI(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (string, error) {
	uri, err := pp.callMethod(ctx, TokenImageURI, methodCallParams{Address: contractAddress, TokenID: tokenID})
	if tokenImageURI, ok := uri.(string); err == nil && ok {
		return tokenImageURI, nil
	}

	return "", err
}

func (pp *Pool) ERC721CollectionName(ctx context.Context, contractAddress common.Address) (string, error) {
	name, err := pp.callMethod(ctx, ERC721CollectionName, methodCallParams{Address: contractAddress})
	if tokenName, ok := name.(string); err == nil && ok {
		return tokenName, nil
	}

	return "", err
}

func (pp *Pool) ERC721CollectionMetadata(ctx context.Context, contractAddress common.Address) (map[string]interface{}, error) {
	collectionMetadata, err := pp.callMethod(ctx, ERC721CollectionMetadata, methodCallParams{Address: contractAddress})
	if metadata, ok := collectionMetadata.(map[string]interface{}); err == nil && ok {
		return metadata, nil
	}

	return nil, err
}

func (pp *Pool) ERC1155TokenName(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (string, error) {
	name, err := pp.callMethod(ctx, ERC1155TokenName, methodCallParams{Address: contractAddress, TokenID: tokenID})
	if tokenName, ok := name.(string); err == nil && ok {
		return tokenName, nil
	}

	return "", err
}

// ERC1155TotalSupply returns the (current) total supply of a token.
func (pp *Pool) ERC1155TotalSupply(ctx context.Context, contractAddress common.Address, tokenID *big.Int) (*big.Int, error) {
	if tokenID == nil {
		return nil, errors.New("tokenID is nil")
	}

	supply, err := pp.callMethod(ctx, ERC1155TotalSupply, methodCallParams{Address: contractAddress, TokenID: tokenID})
	if totalSupply, ok := supply.(*big.Int); err == nil && ok {
		return totalSupply, nil
	}

	return nil, err
}

//
// ens related
//

func (pp *Pool) ReverseResolveAddressToENS(ctx context.Context, address common.Address) (string, error) {
	if address == (common.Address{}) {
		return "", errors.New("address is empty")
	}

	if address == internal.ZeroAddress {
		return "", errors.New("address is zero address")
	}

	// if cachedName, err := cache.GetENSName(ctx, address); err == nil && cachedName != "" {
	if cachedName, err := pp.Rueidi.GetCachedENSName(ctx, address); err == nil && cachedName != "" {
		gbl.Log.Debugf("ens ensName for address %s is cached: %s", address.Hex(), cachedName)

		return cachedName, nil
	}

	name, err := pp.callMethod(ctx, ReverseResolveENS, methodCallParams{Address: address})
	gbl.Log.Debugf("pp.callMethod result - ens ensName for address %s is %+v", address.Hex(), name)

	if ensName, ok := name.(string); err == nil && ok && ensName != "" {
		// cache.StoreENSName(ctx, address, ensName)
		pp.Rueidi.StoreENSName(ctx, address, ensName)

		return ensName, nil
	}

	return "", errors.New("ens ensName not found")
}

func (pp *Pool) ResolveENS(ctx context.Context, ensName string) (common.Address, error) {
	if ensName == "" {
		return common.Address{}, errors.New("ensName is empty")
	}

	address, err := pp.callMethod(ctx, ResolveENS, methodCallParams{EnsName: ensName})
	gbl.Log.Debugf("pp.callMethod result - hex address for ensName %s is %+v", ensName, address)

	if err == nil && address != "" {
		if addr, ok := address.(common.Address); ok {
			return addr, nil
		}
		//		cache.StoreENSName(ctx, address, ensName)
		//		return ensName, nil
	}

	if err != nil {
		gbl.Log.Errorf("pp.callMethod error - hex address for ensName %s is %+v", ensName, err)
	}

	return common.Address{}, errors.New("ens address not found")
}

func (pp *Pool) GetCurrentGasInfo() (*nemo.GasInfo, error) {
	// return nc.getNode().GetCurrentGasInfo()s
	gas, err := pp.callMethod(context.Background(), GasInfo, methodCallParams{})
	if gasInfo, ok := gas.(*nemo.GasInfo); err == nil && ok {
		return gasInfo, nil
	}

	return nil, err
}

// // getClients returns a shuffled list of eth clients with local nodes preferred.
// func (pp *Pool) getClients() []*ethclient.Client {
// 	clients := make([]*ethclient.Client, 0)

// 	// get clients from all nodes
// 	for _, node := range *pp {
// 		clients = append(clients, node.Client)
// 	}

// 	// shuffle clients to avoid hitting the same node over and over again
// 	rand.Shuffle(len(clients), func(i, j int) {
// 		clients[i], clients[j] = clients[j], clients[i]
// 	})

// 	// prefer local nodes if available
// 	localNodeclients := make([]*ethclient.Client, 0)
// 	if nodes := pp.getPreferredProviders(); len(nodes) > 0 {
// 		for _, node := range nodes {
// 			localNodeclients = append(localNodeclients, node.Client)
// 		}
// 	}

// 	return append(localNodeclients, clients...)
// }

// // BlockByNumber returns the given full block.
// func (pp *Pool) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
// 	var err error

// 	for _, client := range pp.getClients() {
// 		if block, err := client.BlockByNumber(ctx, number); err == nil {
// 			return block, nil
// 		}
// 	}

// 	return nil, err
// }

//
// token related methods
//

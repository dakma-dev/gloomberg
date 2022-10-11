package nodes

import (
	"errors"
	"math/big"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/models/standard"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/common"

	"github.com/wealdtech/go-ens/v3"
)

type Nodes []*Node

func (nc *Nodes) ConnectAllNodes() *Nodes {
	if len(*nc) == 0 {
		return nil
	}

	var nodesAvailable uint64

	conectedNodes := make(Nodes, 0)

	var wgNodes sync.WaitGroup
	for _, node := range *nc {
		wgNodes.Add(1)

		go func(node *Node) {
			defer wgNodes.Done()

			if err := node.Connect(); err != nil {
				gbl.Log.Warnf("node %d connection failed: %s", node.NodeID, err)
			} else {
				atomic.AddUint64(&nodesAvailable, 1)
				conectedNodes = append(conectedNodes, node)
			}
		}(node)
	}

	wgNodes.Wait()

	*nc = conectedNodes

	if nodesAvailable == 0 {
		gbl.Log.Fatal("no nodes available")
	}

	return &conectedNodes
}

func (nc *Nodes) GetLocalNodes() []*Node {
	if *nc != nil && len(*nc) == 0 {
		return nil
	}

	nodes := make([]*Node, 0, len(*nc))

	for _, node := range *nc {
		if node.LocalNode {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

func (nc *Nodes) GetRandomNode() *Node {
	if *nc != nil && len(*nc) == 0 {
		return nil
	}

	nodes := make([]*Node, 0, len(*nc))

	for _, node := range *nc {
		if node.Client != nil {
			nodes = append(nodes, node)
		}
	}

	//nolint:gosec
	return nodes[rand.Intn(len(nodes))]
}

func (nc *Nodes) GetRandomLocalNode() *Node {
	if *nc != nil && len(*nc) == 0 {
		return nil
	}

	localNodes := nc.GetLocalNodes()

	if len(localNodes) == 0 {
		return nil
	}

	//nolint:gosec
	return localNodes[rand.Intn(len(localNodes))]
}

// GetNodeByID rer
func (nc *Nodes) GetNodeByID(nodeID int) *Node {
	if *nc != nil && len(*nc) == 0 {
		return nil
	}

	for _, node := range *nc {
		if node.NodeID == nodeID {
			return node
		}
	}

	return nil
}

func (nc *Nodes) ReverseResolveAllENS(wallets *wallet.Wallets) {
	var wgENS sync.WaitGroup

	for _, w := range *wallets {
		wgENS.Add(1)

		go func(w *wallet.Wallet) {
			defer wgENS.Done()

			if name, err := nc.GetENSForAddress(w.Address); err != nil {
				gbl.Log.Warnf("❌ failed to resolve ENS name for %s: %s", w.Address.Hex(), err)
				return
			} else {
				if ensName, err := ens.NewName(nc.GetRandomNode().Client, name); err != nil {
					gbl.Log.Warnf("failed to create ENS %s: %s", name, err)
				} else {
					w.ENS = ensName
					w.Name = ensName.Name
				}
			}
		}(w)
	}

	wgENS.Wait()
}

// func (nc *Nodes) ReverseResolveAllENS(wallets *wallet.Wallets) {
// 	var wgENS sync.WaitGroup

// 	for _, w := range *wallets {
// 		wgENS.Add(1)

// 		go func(w *wallet.Wallet) {
// 			ethClient := nc.GetRandomNode().Client
// 			if name := ens.Format(ethClient, w.Address); name != "" {
// 				w.Name = name
// 				if ensName, err := ens.NewName(ethClient, name); err == nil {
// 					w.ENS = ensName
// 				}
// 			}
// 		}(w)
// 	}

// 	wgENS.Wait()
// }

func (nc *Nodes) GetENSForAddress(address common.Address) (string, error) {
	var ensName string

	if cachedName, err := cache.GetENSName(address); err == nil && cachedName != "" {
		gbl.Log.Debugf("ens ensName for address %s is cached", address.Hex())

		return cachedName, nil
	}

	if resolvedName, err := nc.reverseLookupAndValidate(address); err != nil {
		gbl.Log.Debugf("ens reverse lookup failed for address %s: %s", address.Hex(), err)

		return "", err
	} else {
		if resolvedName == "" {
			gbl.Log.Debugf("address %s has no associated ens ensName", address.Hex())
		}

		ensName = resolvedName
	}

	cache.StoreENSName(address, ensName)

	return ensName, nil
}

func (nc *Nodes) reverseLookupAndValidate(address common.Address) (string, error) {
	var ensName string

	var err error

	client := nc.GetRandomNode().Client

	// lookup the ens ensName for an address
	ensName, err = ens.ReverseResolve(client, address)

	if err != nil || common.IsHexAddress(ensName) {
		gbl.Log.Debugf("ens reverse resolve error: %s -> %s: %s", address, ensName, err)

		return "", err
	}

	// do a lookup for the ensName to validate its authenticity
	resolvedAddress, err := ens.Resolve(client, ensName)
	if err != nil {
		gbl.Log.Debugf("ens resolve error: %s -> %s: %s", ensName, address, err)

		return "", err
	}

	if resolvedAddress != address {
		gbl.Log.Warnf("addresses do not match for: %s", style.BoldStyle.Render(ensName))
		gbl.Log.Debugf("  %s  !=  %s", resolvedAddress.Hex(), address.Hex())

		return "", errors.New("ens forward and reverse resolved addresses do not match")
	}

	return ensName, nil
}

func (nc *Nodes) GetSupportedStandards(contractAddress common.Address) []standard.Standard {
	return nc.GetRandomLocalNode().GetSupportedStandards(contractAddress)
}

func (nc *Nodes) GetERC1155TokenID(contractAddress common.Address, data []byte) *big.Int {
	return nc.GetRandomLocalNode().GetERC1155TokenID(contractAddress, data)
}

func (nc *Nodes) ERC1155Supported(contractAddress common.Address) bool {
	if node := nc.GetRandomLocalNode(); node != nil {
		return node.ERC1155Supported(contractAddress)
	}

	return false
}

package node

import (
	"errors"
	"math/rand"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

type Nodes []*Node

func (nc *Nodes) GetLocalNodes() []*Node {
	if len(*nc) == 0 {
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
	if len(*nc) == 0 {
		return nil
	}

	//nolint:gosec
	return (*nc)[rand.Intn(len(*nc))]
}

func (nc *Nodes) GetRandomLocalNode() *Node {
	if len(*nc) == 0 {
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

	cache.CacheENSName(address, ensName)

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
		gbl.Log.Warnf("ens resolve error: %s -> %s: %s", ensName, address, err)

		return "", err
	}

	if resolvedAddress != address {
		gbl.Log.Warnf("addresses do not match: %s != %s", resolvedAddress.Hex(), address.Hex())

		return "", errors.New("ens forward and reverse resolved addresses do not match")
	}

	return ensName, nil
}

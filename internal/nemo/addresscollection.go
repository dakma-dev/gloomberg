package nemo

import "github.com/ethereum/go-ethereum/common"

type AddressCollection []common.Address

// Contains returns true if the given string is in the slice.
func (ac *AddressCollection) Contains(address common.Address) bool {
	for _, collectionAddress := range *ac {
		if address == collectionAddress {
			return true
		}
	}

	return false
}

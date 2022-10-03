package analytics

import (
	"github.com/ethereum/go-ethereum/common"
)

type CollectionSet struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	ContractAddress common.Address `json:"contractAddress"`
	TokenIDs        []int          `json:"token_ids"`
}

type MultiCollectionSet struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	CollectionSets []CollectionSet `json:"ollectionsets"`
}

type Artist struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	CollectionSets      []CollectionSet      `json:"collectionsets"`
	MultiCollectionSets []MultiCollectionSet `json:"multicollectionsets"`
}

var CollectionSets = map[string]*CollectionSet{
	"OSFEditions": {
		ID:              "OSFEditions",
		Name:            "OSF Editions",
		ContractAddress: common.HexToAddress("0xc23a563a26afff06e945ace77173e1568f288ce5"),
		TokenIDs:        []int{1, 3, 4, 5, 6, 7, 8, 11},
	},
	"OSFRLD": {
		ID:              "OSFRLD",
		Name:            "OSF RLD",
		ContractAddress: common.HexToAddress("0x513cd71defc801b9c1aa763db47b5df223da77a2"),
		TokenIDs:        []int{-1},
	},
	"OSF7Sins": {
		ID:              "OSF7Sins",
		Name:            "OSF's 7 Deadly Sins",
		ContractAddress: common.HexToAddress("0x8297d8e55c27aa6ce2d8a65b1fa3debb02410efc"),
		TokenIDs:        []int{2, 3, 4, 5, 6, 7, 8, 9},
	},
}

var MultiCollectionSets = map[string]*MultiCollectionSet{
	"OSFRLDEditions": {
		ID:             "OSFRLDEditions",
		Name:           "OSF RLD Editions",
		CollectionSets: []CollectionSet{*CollectionSets["OSFEditions"], *CollectionSets["OSFRLD"]},
	},
	"OSFRLDEditionsSins": {
		ID:             "OSFRLDEditionsSins",
		Name:           "OSF RLD Editions & 7 Sins",
		CollectionSets: []CollectionSet{*CollectionSets["OSFEditions"], *CollectionSets["OSFRLD"], *CollectionSets["OSF7Sins"]},
	},
}

var Artists = map[string]*Artist{
	"OSF": {
		ID:                  "OSF",
		Name:                "OSF",
		CollectionSets:      []CollectionSet{*CollectionSets["OSFEditions"], *CollectionSets["OSFRLD"], *CollectionSets["OSF7Sins"]},
		MultiCollectionSets: []MultiCollectionSet{*MultiCollectionSets["OSFRLDEditions"], *MultiCollectionSets["OSFRLDEditionsSins"]},
	},
}

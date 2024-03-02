package degendb

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// type AddressType string

// const (
// 	AddressTypeEOA      AddressType = "EOA"
// 	AddressTypeContract AddressType = "contract"
// )

type Address struct {
	// Address is the ethereum address for this address
	ID string `bson:"_id" json:"_id"`

	// // Address is the ethereum Address for this address
	// Address common.Address `json:"address"`

	// Type is the type of this address, e.g. EOA or contract
	Type AccountType `bson:"type,omitempty" json:"type,omitempty"`
	// IsContract bool        `bson:"is_contract,omitempty" json:"is_contract,omitempty"`

	// ENS is the (main/reverseable) domain (e.g. ENS) for this address
	ENS string `bson:"domain,omitempty" json:"domain,omitempty"`

	// Name is a name for this address
	Name string `bson:"name,omitempty" json:"name,omitempty"`

	// Tags is a list of tags associated with this address
	Tags []Tag `bson:"tags,omitempty" json:"tags,omitempty"`

	// First1K collections in which this address was in the first ~1k txs
	First1K []Tag `bson:"first1k,omitempty" json:"first1k,omitempty"`

	// First1K fetched at | in case of a contract, the time we fetched the first 1k txs
	First1KFetchedAt time.Time `bson:"first1kFetchedAt,omitempty" json:"first1k_at,omitempty"`

	// CreatedAt is the time this address was created in the db
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	// UpdatedAt is the time this address was last updated in the db
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`

	// //
	// // Collection data

	// OpenSea is the OpenSea slug of the token
	SlugOpenSea string `bson:"slug_opensea,omitempty" json:"slug_opensea,omitempty"`

	// Slugs are the collection slugs of the collection
	// Slugs Slugs `bson:"slugs,omitempty" json:"slugs,omitempty"`

	// // ContractStandard is the standard of the contract
	// ContractStandard string `bson:"contract_standard,omitempty" json:"contract_standard,omitempty"`

	// // Description is the description of the collection
	// Description string `bson:"description,omitempty" json:"description,omitempty"`

	// // TotalSupply is the total supply of the collection
	// TotalSupply int `bson:"total_supply,omitempty" json:"total_supply,omitempty"`

	// // ImageURL is the URL of the image of the collection
	// ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty"`

	// // ExternalURL is the URL of the collection
	// ExternalURL string `bson:"external_url,omitempty" json:"external_url,omitempty"`

	// creates to many writes probably
	// // LastSeen is the time this wallet was last seen
	// LastSeen int64 `bson:"last_seen,omitempty" json:"last_seen"`
}

func (a *Address) String() string {
	return a.AsAddress().Hex()
}

func (a *Address) AsAddress() common.Address {
	return common.HexToAddress(a.ID)
}

// func (a *Address) SaveWithFirst1KSlugs(ddb *DegenDB, address common.Address, first1kSlugs []Tag) {
// 	addressesColl := ddb.mongo.Database(mongoDatabase).Collection(collAddresses)

// 	// addr := ddb.NewAddress(address)
// 	addr := &Address{HexAddress: address.Hex(), Address: address}

// 	addrFilter := bson.D{{"HexAddress", address.Hex()}}

// 	upd := bson.D{
// 		{"$set", addr},
// 		{"$addToSet", bson.D{{"first1K", bson.D{{"$each", first1kSlugs}}}}}, // bson.D{{"First1K", bson.D{{"$each", first1kSlugs}}
// 	}

// 	reallyTrue := true

// 	addressResult := addressesColl.FindOneAndUpdate(context.TODO(), addrFilter, upd, &options.FindOneAndUpdateOptions{
// 		Upsert: &reallyTrue,
// 	})

// 	if addressResult.Err() != nil {
// 		log.Error(addressResult.Err())
// 	}

// 	if addressResult == nil {
// 		log.Printf("addressResult.InsertedIDs is nil")

// 		return
// 	}

// 	log.Printf("addressResult: %+v", addressResult)
// }

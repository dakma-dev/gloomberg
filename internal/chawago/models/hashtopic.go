package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type HashTopics map[int]HashTopic

type HashTopic struct {
	common.Hash
	topicName string
}

func HashTopicsFromLog(rawTopics []common.Hash) HashTopics {
	topics := make(HashTopics, len(rawTopics))

	for i, topic := range rawTopics {
		topics[i] = NewHashTopic(topic)
	}

	return topics
}

func NewHashTopic(topicHash common.Hash) HashTopic {
	// get topic name from a db we have to create
	// topicName := "moepel"

	return HashTopic{topicHash, ""}
}

func (t *HashTopic) String() string {
	return t.topicName
}

var (
	Transfer       = HashTopic{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), "Transfer"}
	TransferSingle = HashTopic{common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"), "TransferSingle"}
	ApprovalForAll = HashTopic{common.HexToHash("0x17307eabf5b53c6c7f7c9b5b9d9e9f0b3b2b2b7b0b4b2b2b2b2b2b2b2b2b2b2b2"), "ApprovalForAll"}
	// // opensea.
	// OrderFulfilled HashTopic = HashTopic{"OrderFulfilled", common.HexToHash("0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31")}
	// // blur.
	// OrdersMatched HashTopic = HashTopic{"OrdersMatched", common.HexToHash("0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64")}
	// // manifold.
	// ClaimMint      HashTopic = HashTopic{"ClaimMint", common.HexToHash("0x5d404f369772cfab2b65717fca9bc2077efeab89a0dbec036bf0c13783154eb1")}
	// ClaimMintBatch HashTopic = HashTopic{"ClaimMintBatch", common.HexToHash("0x74f5d3254dfa39a7b1217a27d5d9b3e061eafe11720eca1cf499da2dc1eb1259")}
	// // foundation.
	// BuyPriceSet HashTopic = HashTopic{"BuyPriceSet", common.HexToHash("0xfcc77ea8bdcce862f43b7fb00fe6b0eb90d6aeead27d3800d9257cf7a05f9d96")}.
)

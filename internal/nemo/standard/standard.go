package standard

import (
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	Standard  int64
	Standards []Standard
)

const (
	UNKNOWN Standard = iota
	ERC20
	ERC165
	ERC721
	ERC1155
)

func (s Standard) String() string {
	return map[Standard]string{
		UNKNOWN: "UNKNOWN", ERC20: "ERC20", ERC165: "ERC165", ERC721: "ERC721", ERC1155: "ERC1155",
	}[s]
}

func (st Standards) Contains(standard Standard) bool {
	for _, s := range st {
		if s == standard {
			return true
		}
	}

	return false
}

func (s Standard) IsERC721orERC1155() bool {
	return s == ERC721 || s == ERC1155
}

func (s Standard) IsERC20() bool {
	return s == ERC20
}

func GetStandard(txLog *types.Log) Standard {
	logStandard := UNKNOWN

	topic0 := topic.Topic(txLog.Topics[0].String())

	switch {
	// erc20
	case topic0 == topic.Transfer && len(txLog.Topics) <= 3:
		logStandard = ERC20

	// erc721
	case topic0 == topic.Transfer && len(txLog.Topics) >= 4:
		logStandard = ERC721

	// erc1155
	case topic0 == topic.TransferSingle && len(txLog.Topics) >= 4:
		logStandard = ERC1155

	default:
		gbl.Log.Debugf("unknown log standard | len(log.Topics): %d | topic0: %s", len(txLog.Topics), topic0)
	}

	return logStandard
}

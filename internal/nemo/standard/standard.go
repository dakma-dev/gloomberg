package standard

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

package standard

type (
	Standard  int64
	Standards []Standard
)

const (
	ERC20 Standard = iota
	ERC165
	ERC721
	ERC1155
)

func (s Standard) String() string {
	return map[Standard]string{
		ERC20: "ERC20", ERC165: "ERC165", ERC721: "ERC721", ERC1155: "ERC1155",
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

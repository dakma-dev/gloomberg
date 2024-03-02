package eip6551

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	TokenboundERC6551Registry     = common.HexToAddress("0x02101dfb77fde026414827fdc604ddaf224f0921")
	TokenboundERC6551AccountProxy = common.HexToAddress("0x2D25602551487C3f3354dD80D76D54383A243358")

	// ethereum mainnet.
	chainID = common.LeftPadBytes(big.NewInt(1).Bytes(), 32)

	// optional, use 0 for now.
	salt = common.LeftPadBytes(big.NewInt(0).Bytes(), 32)

	// erc-1167 header & footer bytes.
	erc1167Header = []byte{0x3d, 0x60, 0xad, 0x80, 0x60, 0x0a, 0x3d, 0x39, 0x81, 0xf3, 0x36, 0x3d, 0x3d, 0x37, 0x3d, 0x3d, 0x3d, 0x36, 0x3d, 0x73}
	erc1167Footer = []byte{0x5a, 0xf4, 0x3d, 0x82, 0x80, 0x3e, 0x90, 0x3d, 0x91, 0x60, 0x2b, 0x57, 0xfd, 0x5b, 0xf3}
)

func GetTokenboundTokenAddress(tokenContract *common.Address, tokenID *big.Int) common.Address {
	creationCode := getCreationCode(TokenboundERC6551AccountProxy, tokenContract, tokenID)

	return computeAddress(crypto.Keccak256(creationCode), [32]byte(salt))
}

func getCreationCode(implementation common.Address, tokenContract *common.Address, tokenID *big.Int) []byte {
	// use default implementation (tokenbound.org) if none provided
	if implementation == (common.Address{}) {
		implementation = TokenboundERC6551AccountProxy
	}

	// concat everything together
	creationCode := make([]byte, 0)
	creationCode = append(creationCode, erc1167Header...)
	creationCode = append(creationCode, implementation.Bytes()...)
	creationCode = append(creationCode, erc1167Footer...)
	creationCode = append(creationCode, salt...)
	creationCode = append(creationCode, chainID...)
	creationCode = append(creationCode, common.LeftPadBytes(tokenContract.Bytes(), 32)...)
	creationCode = append(creationCode, common.LeftPadBytes(tokenID.Bytes(), 32)...)

	return creationCode
}

func computeAddress(creationCodeHash []byte, salt [32]byte) common.Address {
	return crypto.CreateAddress2(TokenboundERC6551Registry, salt, creationCodeHash)
}

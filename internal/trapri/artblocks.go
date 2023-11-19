// nolint: golint
package trapri

import (
	"math/big"

	"github.com/benleb/gloomberg/internal/abis/artblocks/GenArt721"
	"github.com/benleb/gloomberg/internal/abis/artblocks/GenArt721Core"
	"github.com/benleb/gloomberg/internal/abis/artblocks/GenArt721CoreV3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ArtblocksContracts = []string{artblockContract1, artblockContract2, artblocksContract}

func IsAddressArtblocksContract(address common.Address) bool {
	for _, contract := range ArtblocksContracts {
		if address == common.HexToAddress(contract) {
			return true
		}
	}

	return false
}

// GenArt721CoreV3.go.
const artblocksContract = "0x99a9b7c1116f9ceeb1652de04d5969cce509b069"

// GenArt721Core Old BLOCKS Token.
const artblockContract1 = "0xa7d8d9ef8d8ce8992df33d8b8cf4aebabd5bd270"

// GenArt721 Old BLOCKS Token.
const artblockContract2 = "0x059EDD72Cd353dF5106D2B9cC5ab83a52287aC3a"

func getProjectNameByContract(tokenID *big.Int, contract common.Address, ethClient *ethclient.Client) (string, *big.Int) {
	projectName := ""
	projectID := big.NewInt(0)

	if contract == common.HexToAddress(artblockContract1) {
		caller1, err := GenArt721Core.NewArtblocksCaller(contract, ethClient)
		if err != nil {
			return "", nil
		}

		projectID, err = caller1.TokenIdToProjectId(&bind.CallOpts{}, tokenID)
		if err != nil {
			return "", nil
		}

		details, err := caller1.ProjectDetails(&bind.CallOpts{}, projectID)
		if err != nil {
			return "", nil
		}
		projectName = details.ProjectName
	}

	if contract == common.HexToAddress(artblockContract2) {
		caller2, err := GenArt721.NewGenArt721(contract, ethClient)
		if err != nil {
			return "", nil
		}
		projectID, err = caller2.TokenIdToProjectId(&bind.CallOpts{}, tokenID)
		if err != nil {
			return "", nil
		}
		details, err := caller2.ProjectDetails(&bind.CallOpts{}, projectID)
		if err != nil {
			return "", nil
		}
		projectName = details.ProjectName
	}

	if contract == common.HexToAddress(artblocksContract) {
		caller, err := GenArt721CoreV3.NewArtblocksCaller(contract, ethClient)
		if err != nil {
			return "", nil
		}
		projectID, err = caller.TokenIdToProjectId(&bind.CallOpts{}, tokenID)
		if err != nil {
			return "", nil
		}
		details, err := caller.ProjectDetails(&bind.CallOpts{}, projectID)
		if err != nil {
			return "", nil
		}
		projectName = details.ProjectName
	}

	return projectName, projectID
}

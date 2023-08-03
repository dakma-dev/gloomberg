package gloomberg

import (
	"context"
	"fmt"
	"math/big"
	"reflect"

	"github.com/benleb/gloomberg/internal/eip6551"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	eip6551TokenCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_eip6551_tokens_checked_count_total",
		Help: "No of tokens checked for an active EIP-6551 account.",
	},
	)
	eip6551TokenDiscoveredCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_eip6551_tokens_discovered_count_total",
		Help: "No of tokens with an active EIP-6551 account.",
	},
	)
)

func JobCheckEIP6551TokenAccount(params ...any) {
	if len(params) != 3 {
		gbl.Log.Warnf("🤦‍♀️ wrong number of params: %d", len(params))

		return
	}

	gb, ok := params[0].(*Gloomberg)
	if !ok {
		gbl.Log.Warnf("🙄 wrong type of param: %s", reflect.TypeOf(params[0]))

		return
	}

	tokenContract, ok := params[1].(*common.Address)
	if !ok {
		gbl.Log.Warnf("🤷‍♀️ wrong type of param: %s", reflect.TypeOf(params[1]))

		return
	}

	tokenID, ok := params[2].(*big.Int)
	if !ok {
		gbl.Log.Warnf("🤷 wrong type of param: %s", reflect.TypeOf(params[2]))

		return
	}

	eip6551TokenCounter.Inc()

	_, _, err := CheckEIP6551TokenAccount(gb, tokenContract, tokenID)
	if err != nil {
		gbl.Log.Warnf("🤷‍♀️ failed to check token account: %s", err)
	}
}

func CheckEIP6551TokenAccount(gb *Gloomberg, tokenContract *common.Address, tokenID *big.Int) (bool, *common.Address, error) {
	nftShortID := fmt.Sprintf("%s/%s", tokenContract.String(), style.BoldAlmostWhite(tokenID.String()))

	// check if token is/has tokenbound account
	eip6551AccountAddress := eip6551.GetTokenboundTokenAddress(tokenContract, tokenID)
	if eip6551AccountAddress == (common.Address{}) {
		log.Errorf("EIP6551 account for %s does not exist??!", nftShortID)

		return false, nil, nil
	}

	log.Debugf("EIP6551 %s → %+v 🧮", nftShortID, eip6551AccountAddress)

	// check if account exists
	codeAt, err := gb.ProviderPool.GetCodeAt(context.Background(), eip6551AccountAddress)
	if err != nil {
		log.Debugf("failed to get nonce for %s: %s", eip6551AccountAddress.String(), err)
		gbl.Log.Errorf("failed to get nonce for %s: %s", eip6551AccountAddress.String(), err)

		return false, nil, fmt.Errorf("failed to get nonce for %s: %w", eip6551AccountAddress.String(), err)
	}

	log.Debugf("EIP6551 codeAt: %+v", codeAt)

	// if there is deployed code at the address, it's a contract
	if isContract := len(codeAt) > 0; !isContract {
		log.Debugf("EIP6551 address %s is not a contract 🙄", eip6551AccountAddress.String())

		return false, nil, nil
	}

	gbl.Log.Infof("‼️ EIP6551 account %s is a contract!? 🧐", eip6551AccountAddress.String())

	// check if account has nonce
	nonceAt, err := gb.ProviderPool.GetNonceAt(context.Background(), eip6551AccountAddress)
	if err != nil {
		log.Debugf("failed to get nonce for %s: %s", eip6551AccountAddress.String(), err)

		return false, nil, fmt.Errorf("failed to get nonce for %s: %w", eip6551AccountAddress.String(), err)
	}

	gbl.Log.Infof("EIP6551 nonceAt: %d", nonceAt)

	if nonceAt < 2 {
		log.Debugf("EIP6551 account %s has no transactions besides contract creation 🤷‍♀️", eip6551AccountAddress.String())

		return false, nil, fmt.Errorf("EIP6551 account %s has no transactions besides contract creation 🤷‍♀️", eip6551AccountAddress.String())
	}

	PrModf(
		"e6551", "💥  %s  →  %+v (nonce: %d)",
		style.TerminalLink(utils.GetEtherscanAddressURL(tokenContract), style.ShortenAddressStyled(tokenContract, lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(tokenContract.Hash().Big().Int64())))),
		style.TerminalLink(utils.GetEtherscanAddressURL(&eip6551AccountAddress), eip6551AccountAddress.String()),
		nonceAt,
	)

	eip6551TokenDiscoveredCounter.Inc()

	return true, &eip6551AccountAddress, nil
}

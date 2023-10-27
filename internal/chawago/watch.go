package chawago

import (
	"fmt"
	"strings"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type WalletWatcher struct {
	gb *gloomberg.Gloomberg

	Wallets        map[common.Address]*wallet.Wallet
	watchAddresses mapset.Set[common.Address]
}

// NewWalletWatcher returns a new walletWatcher.
func NewWalletWatcher(gb *gloomberg.Gloomberg) *WalletWatcher {
	return &WalletWatcher{
		gb: gb,

		Wallets:        make(map[common.Address]*wallet.Wallet),
		watchAddresses: mapset.NewSet[common.Address](),
	}
}

// Pr prints messages from seawatcher to the terminal.
func (ww *WalletWatcher) Pr(message string) {
	gloomberg.PrWithKeywordAndIcon("👀", lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0099")).Faint(true).Render("wawa"), message)
}

func (ww *WalletWatcher) PrDf(format string, a ...interface{}) {
	lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0099")).Faint(true).Render("wawa")

	ww.Pr(fmt.Sprintf(format, a...))
}

// Prf formats and prints messages from seawatcher to the terminal.
func (ww *WalletWatcher) Prf(format string, a ...interface{}) {
	ww.Pr(fmt.Sprintf(format, a...))
}

func (ww *WalletWatcher) FormattedWallets() []string {
	names := make([]string, 0)
	for _, w := range ww.Wallets {
		names = append(names, lipgloss.NewStyle().Foreground(w.Color).Render(w.Name))
	}

	return names
}

func (ww *WalletWatcher) Watch() {
	// demo/testing wallets
	ww.Wallets[common.HexToAddress("0x34d3119a6b983af8eec6bcebd79bc5a235671b5b")] = &wallet.Wallet{
		Name:    "scammy",
		Address: common.HexToAddress("0x34d3119a6b983af8eec6bcebd79bc5a235671b5b"),
		Color:   lipgloss.Color("#ff0099"),
	}
	ww.Wallets[common.HexToAddress("0xdcae87821fa6caea05dbc2811126f4bc7ff73bd1")] = &wallet.Wallet{
		Name:    "OSF",
		Address: common.HexToAddress("0xdcae87821fa6caea05dbc2811126f4bc7ff73bd1"),
		Color:   lipgloss.Color("#031099"),
	}

	// collect wallet addresses in a set
	ww.watchAddresses = mapset.NewSetFromMapKeys[common.Address](ww.Wallets)

	ww.Prf("watching %d addresses: %+v", ww.watchAddresses.Cardinality(), strings.Join(ww.FormattedWallets(), ", "))

	// watch for new transactions
	txsWithLogs := ww.gb.SubscribeTxWithLogs()

	for tx := range txsWithLogs {
		addressesInTx := mapset.NewSet[common.Address]()

		if tx.Sender() != nil {
			addressesInTx.Add(*tx.Sender())
		}

		if tx.To() != nil {
			addressesInTx.Add(*tx.To())
		}

		for _, txLog := range tx.Logs {
			for _, txTopic := range txLog.Topics {
				if address := common.HexToAddress(txTopic.Hex()); address != (common.Address{}) {
					addressesInTx.Add(address)
				}
			}
		}

		// check if any of the addresses in the tx are in the watchedWallets
		watchedAddressesInTx := addressesInTx.Intersect(ww.watchAddresses)

		if watchedAddressesInTx.Cardinality() == 0 {
			gbl.Log.Debugf("no watched wallet (%s) in tx %s", ww.watchAddresses, tx.Hash().Hex())

			continue
		}

		fmtStalkees := make([]string, 0)
		for _, address := range watchedAddressesInTx.ToSlice() {
			fmtStalkees = append(fmtStalkees, lipgloss.NewStyle().Foreground(ww.Wallets[address].Color).Render(ww.Wallets[address].Name))
		}

		notifyMsg := fmt.Sprintf(
			"  👀  %s  👀  → %s",
			strings.Join(fmtStalkees, ", "),
			style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.ShortenHashStyled(tx.TxHash)),
		)

		if viper.GetInt64("notifications.telegram.my_chat_id") != 0 {
			fmtStalkees := make([]string, 0)
			for _, address := range watchedAddressesInTx.ToSlice() {
				fmtStalkees = append(fmtStalkees, ww.Wallets[address].Name)
			}

			notifyMsg := fmt.Sprintf(
				"  👀  %s  👀  → %s",
				strings.Join(fmtStalkees, ", "),
				"https://etherscan.io/tx/"+tx.Hash().Hex(),
			)

			go notify.SendMessageViaTelegram(notifyMsg, viper.GetInt64("notifications.telegram.my_chat_id"), "", 0, nil)
		}

		// do something with the transaction
		txSentFromWatchedAddress := tx.Sender() != nil && ww.watchAddresses.Contains(*tx.Sender())

		if txSentFromWatchedAddress {
			// do something
			log.Printf("")
			log.Printf(notifyMsg)
			log.Printf("")
		}

		for _, txLog := range tx.Logs {
			if len(txLog.Topics) > 0 && ww.watchAddresses.Contains(common.HexToAddress(txLog.Topics[0].Hex())) {
				// do something
				log.Debugf("wallet %s in topic0: %s", ww.Wallets[common.HexToAddress(txLog.Topics[0].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
			if len(txLog.Topics) > 1 && ww.watchAddresses.Contains(common.HexToAddress(txLog.Topics[1].Hex())) {
				// do something
				log.Debugf("wallet %s in topic1: %s", ww.Wallets[common.HexToAddress(txLog.Topics[1].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
			if len(txLog.Topics) > 2 && ww.watchAddresses.Contains(common.HexToAddress(txLog.Topics[2].Hex())) {
				// do something
				log.Debugf("wallet %s in topic2: %s:", ww.Wallets[common.HexToAddress(txLog.Topics[2].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
			if len(txLog.Topics) > 3 && ww.watchAddresses.Contains(common.HexToAddress(txLog.Topics[3].Hex())) {
				// do something
				log.Debugf("wallet %s in topic3: %s", ww.Wallets[common.HexToAddress(txLog.Topics[3].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
		}
	}
}

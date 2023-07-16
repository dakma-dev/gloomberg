package chawago

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
)

type WalletWatcher struct {
	gb *gloomberg.Gloomberg

	Wallets        map[common.Address]*wallet.Wallet
	watchedWallets mapset.Set[common.Address]
}

// NewWalletWatcher returns a new walletWatcher.
func NewWalletWatcher(gb *gloomberg.Gloomberg) *WalletWatcher {
	return &WalletWatcher{
		gb: gb,

		Wallets:        make(map[common.Address]*wallet.Wallet),
		watchedWallets: mapset.NewSet[common.Address](),
	}
}

// Pr prints messages from seawatcher to the terminal.
func (ww *WalletWatcher) Pr(message string) {
	ww.gb.PrWithKeywordAndIcon("👀", lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0099")).Faint(true).Render("wawa"), message)
}

func (ww *WalletWatcher) PrDf(format string, a ...interface{}) {
	lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0099")).Faint(true).Render("wawa")

	ww.Pr(fmt.Sprintf(format, a...))
}

// Prf formats and prints messages from seawatcher to the terminal.
func (ww *WalletWatcher) Prf(format string, a ...interface{}) {
	ww.Pr(fmt.Sprintf(format, a...))
}

func (ww *WalletWatcher) Watch() {
	ww.Wallets[common.HexToAddress("0x1329d762e5e34e53a6c5e24bd15fa032482eb0f9")] = &wallet.Wallet{
		Name:    "unknown",
		Address: common.HexToAddress("0x1329d762e5e34e53a6c5e24bd15fa032482eb0f9"),
	}

	ww.watchedWallets = mapset.NewSetFromMapKeys[common.Address](ww.Wallets)

	ww.Prf("watching %d wallets: %+v", ww.watchedWallets.Cardinality(), ww.watchedWallets)

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
		if addressesInTx.Intersect(ww.watchedWallets).Cardinality() == 0 {
			gbl.Log.Debugf("no watched wallet in tx %s", tx.Hash().Hex())

			continue
		}

		// go notify.SendMessageViaTelegram(fmt.Sprintf("👀  %s", "https://etherscan.io/tx/"+tx.Hash().Hex()), viper.GetInt64("notifications.telegram.chat_id"), "", 0, nil)

		log.Printf(" | %s", style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))

		// do something with the transaction
		if tx.Sender() != nil && ww.watchedWallets.Contains(*tx.Sender()) {
			// do something
			log.Printf("Transaction from wallet %s: %s", ww.Wallets[*tx.Sender()].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
		}

		for _, txLog := range tx.Logs {
			if len(txLog.Topics) > 0 && ww.watchedWallets.Contains(common.HexToAddress(txLog.Topics[0].Hex())) {
				// do something
				log.Printf("wallet %s in topic0: %s", ww.Wallets[common.HexToAddress(txLog.Topics[0].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
			if len(txLog.Topics) > 1 && ww.watchedWallets.Contains(common.HexToAddress(txLog.Topics[1].Hex())) {
				// do something
				log.Printf("wallet %s in topic1: %s", ww.Wallets[common.HexToAddress(txLog.Topics[1].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
			if len(txLog.Topics) > 2 && ww.watchedWallets.Contains(common.HexToAddress(txLog.Topics[2].Hex())) {
				// do something
				log.Printf("wallet %s in topic2: %s:", ww.Wallets[common.HexToAddress(txLog.Topics[2].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
			if len(txLog.Topics) > 3 && ww.watchedWallets.Contains(common.HexToAddress(txLog.Topics[3].Hex())) {
				// do something
				log.Printf("wallet %s in topic3: %s", ww.Wallets[common.HexToAddress(txLog.Topics[3].Hex())].Name, style.TerminalLink("https://etherscan.io/tx/"+tx.Hash().Hex(), style.AlmostWhiteStyle.Render(tx.Hash().Hex())))
			}
		}
	}
}

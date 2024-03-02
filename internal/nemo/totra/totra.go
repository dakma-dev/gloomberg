package totra

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TokenTransaction struct {
	// the original tx
	Tx *types.Transaction `json:"tx"`

	// the tx Hash
	TxHash common.Hash `json:"hash"`

	// the receipt containing the logs
	TxReceipt *types.Receipt `json:"tx_receipt"`

	// action performed by the tx
	Action degendb.EventType `json:"action"`

	// the sender of the tx
	From common.Address `json:"from"`

	// signature of the called contract function
	MethodID [4]byte `json:"function_signature"`

	// the amount of eth/weth transferred in the tx
	AmountPaid *big.Int `json:"amount_paid"`

	Marketplace *marketplace.MarketPlace `json:"marketplace"`

	// token transfers parsed from the tx logs
	Transfers []*TokenTransfer `json:"transfers"`

	TotalTokens int64 `json:"total_tokens"`

	ReceivedAt time.Time `json:"received_at"`

	logsByStandard map[standard.Standard][]*types.Log

	sentMoney map[common.Address]*big.Int
	sentToken map[common.Address][]*token.Token

	DoNotPrint bool `json:"do_not_print"`
	Highlight  bool `json:"highlight"`
}

// var methodSignaturesTransfers = map[[4]byte]string{
// 	{0x23, 0xb8, 0x72, 0xdd}: "transferFrom(address,address,uint256)",
// 	{0x42, 0x84, 0x2e, 0x0e}: "safeTransferFrom(address,address,uint256)",
// 	{0xb8, 0x8d, 0x4f, 0xde}: "safeTransferFrom(address,address,uint256,bytes)",
// 	{0xf2, 0x42, 0x43, 0x2a}: "safeTransferFrom(address,address,uint256,uint256,bytes)",
// 	{0x32, 0x38, 0x9b, 0x71}: "bulkTransfer(((uint8,address,uint256,uint256)[],address,bool)[],bytes32)",
// }

func (ttx *TokenTransaction) GetEtherscanTxURL() string {
	return fmt.Sprintf("https://etherscan.io/tx/%s", ttx.TxHash)
}

func (ttx *TokenTransaction) GetTransferredTokenContractAdresses() mapset.Set[common.Address] {
	return mapset.NewSetFromMapKeys[common.Address](ttx.GetTransfersByContract())
}

func (ttx *TokenTransaction) GetTransfersByContract() map[common.Address][]*TokenTransfer {
	transfersByContract := make(map[common.Address][]*TokenTransfer)

	for _, transfer := range ttx.Transfers {
		transfersByContract[transfer.Token.Address] = append(transfersByContract[transfer.Token.Address], transfer)
	}

	return transfersByContract
}

func (ttx *TokenTransaction) GetPrice() *price.Price {
	if ttx.AmountPaid == nil || ttx.AmountPaid.Cmp(big.NewInt(0)) == 0 {
		return price.NewPrice(big.NewInt(0))
	}

	return price.NewPrice(ttx.AmountPaid)
}

func (ttx *TokenTransaction) GetNFTReceivers() map[common.Address][]*TokenTransfer {
	nftReceivers := make(map[common.Address][]*TokenTransfer)

	for _, transfer := range ttx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			nftReceivers[transfer.To] = append(nftReceivers[transfer.To], transfer)
		}
	}

	return nftReceivers
}

func (ttx *TokenTransaction) GetNFTReceiverAddresses() mapset.Set[common.Address] {
	return mapset.NewSetFromMapKeys[common.Address](ttx.GetNFTReceivers())
}

func (ttx *TokenTransaction) GetNFTSenders() map[common.Address][]*TokenTransfer {
	nftSenders := make(map[common.Address][]*TokenTransfer)

	for _, transfer := range ttx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			nftSenders[transfer.From] = append(nftSenders[transfer.From], transfer)
		}
	}

	return nftSenders
}

func (ttx *TokenTransaction) GetNonZeroNFTSenders() map[common.Address][]*TokenTransfer {
	nftSenders := ttx.GetNFTSenders()

	nonZeroSenders := make(map[common.Address][]*TokenTransfer)

	for addr, sender := range nftSenders {
		if addr == internal.ZeroAddress {
			continue
		}

		nonZeroSenders[addr] = sender
	}

	return nonZeroSenders
}

func (ttx *TokenTransaction) GetNFTSenderAddresses() mapset.Set[common.Address] {
	return mapset.NewSetFromMapKeys[common.Address](ttx.GetNFTSenders())
}

func (ttx *TokenTransaction) GetNFTSenderAndReceiverAddresses() mapset.Set[common.Address] {
	return ttx.GetNFTSenderAddresses().Union(ttx.GetNFTReceiverAddresses())
}

func (ttx *TokenTransaction) parseTransfersFromReceipt() {
	// assuming every nft is just sold once per tx
	uniqueTransfers := make(map[string][]*TokenTransfer)

	for logStandard, txLogs := range ttx.logsByStandard {
		log.Debugf("  🧱 ttx logs to parse: %+v", len(txLogs))

		for _, txLog := range txLogs {
			// parse Transfer & TransferSingle logs
			var transfer *TokenTransfer

			switch logStandard {
			case standard.ERC20:
				transfer = parseERC20TransferLog()

			case standard.ERC721:
				transfer = parseERC721TransferLog(txLog)

			case standard.ERC1155:
				transfer = parseERC1155TransferLog()
				// if transfer != nil {
				// 	out := strings.Builder{}
				// 	out.WriteString(fmt.Sprintf("from: %+v | to: %+v | amount: %+v | tokenid: %+v | addr: %+v\n", style.ShortenAddress(&transfer.OldFrom), style.ShortenAddress(transfer.To()), transfer.AmountTokens(), transfer.Token.ID, transfer.Token.Address))
				// 	parsedLog := lopas.ParseTransferLog(txLog, ethNode)
				// 	out.WriteString(fmt.Sprintf("from: %+v | to: %+v | amount: %+v | tokenid: %+v | addr: %+v\n\n", style.ShortenAddress(parsedLog.From()), style.ShortenAddress(parsedLog.To()), parsedLog.AmountTokens(), parsedLog.TokenID(), parsedLog.Token.Address))
				// }
			}

			if transfer == nil {
				continue
			}

			nftID := transfer.Token.NftID()

			if len(uniqueTransfers[nftID]) == 0 || transfer.Standard.IsERC20() {
				ttx.Transfers = append(ttx.Transfers, transfer)
			}

			uniqueTransfers[nftID] = append(uniqueTransfers[nftID], transfer)
		}
	}
}

// func (ttx *TokenTransaction) parseTransfersFromReceipt(ethNode *nodes.Node) {
// 	// assuming every nft is just sold once per tx
// 	uniqueTransfers := make(map[string][]*TokenTransfer, 0)

// 	for _, txLog := range ttx.TxReceipt.Logs {
// 		log.Debugf("  🧱 blockParser | ttx log: %+v", txLog)

// 		// parse Transfer & TransferSingle logs
// 		var transfer *TokenTransfer

// 		switch logStandard := getTransferLogStandard(txLog); logStandard {
// 		case standard.ERC20:
// 			transfer = parseERC20TransferLog(txLog, ethNode)
// 		case standard.ERC721:
// 			transfer = parseERC721TransferLog(txLog)
// 		case standard.ERC1155:
// 			transfer = parseERC1155TransferLog(txLog, ethNode)
// 		}

// 		if transfer != nil {
// 			nftID := utils.GetNFTID(transfer.Token.Address, transfer.Token.ID.Uint64())

// 			if transfer != nil && len(uniqueTransfers[nftID]) == 0 {
// 				ttx.Transfers = append(ttx.Transfers, transfer)
// 			}

// 			uniqueTransfers[nftID] = append(uniqueTransfers[nftID], transfer)
// 		}
// 	}

// 	ttx.parseERC20Transfers()
// }

func (ttx *TokenTransaction) parseERC20Transfers() {
	ttx.sentMoney = make(map[common.Address]*big.Int)
	ttx.sentToken = make(map[common.Address][]*token.Token)

	amountPaidERC20 := big.NewInt(0)

	for _, transfer := range ttx.Transfers {
		// reactivate me
		// if transfer.Standard == standard.ERC20 {
		// 	log.Debugf("providerPool.IsContract(%s): %+v ", transfer.From.Hex(), providerPool.IsContract(transfer.From, rueidi))

		// 	if providerPool.IsContract(transfer.From, rueidi) || marketplace.Blur.ContractAddresses.Contains(transfer.From) {
		// 		continue
		// 	}

		// 	if _, ok := ttx.sentMoney[transfer.From]; !ok {
		// 		ttx.sentMoney[transfer.From] = big.NewInt(0)
		// 	}

		// 	ttx.sentMoney[transfer.From].Add(ttx.sentMoney[transfer.From], transfer.AmountTokens)

		// 	amountPaidERC20.Add(amountPaidERC20, transfer.AmountTokens)
		// }

		if transfer.Standard == standard.ERC721 || transfer.Standard == standard.ERC1155 {
			if _, ok := ttx.sentToken[transfer.From]; !ok {
				ttx.sentToken[transfer.From] = make([]*token.Token, 0)
			}

			ttx.sentToken[transfer.From] = append(ttx.sentToken[transfer.From], transfer.Token)
		}
	}

	ttx.AmountPaid.Add(ttx.AmountPaid, amountPaidERC20)
}

// discoverItemPrices tries to find the price of single nfts in a transaction
// it does so by looking at the amount of money someone who sent nfts received in return.
func (ttx *TokenTransaction) discoverItemPrices() {
	for _, tokenTransfer := range ttx.Transfers {
		if tokenTransfer.Standard.IsERC721orERC1155() {
			for _, moneyTransfer := range ttx.Transfers {
				if moneyTransfer.Standard == standard.ERC20 {
					if tokenTransfer.From == moneyTransfer.To && moneyTransfer.AmountTokens.Cmp(big.NewInt(0)) > 0 {
						tokenTransfer.AmountEtherReturned = moneyTransfer.AmountTokens

						// remove money transfer from list
						moneyTransfer.AmountTokens = big.NewInt(0)

						break
					}
				}
			}
		}
	}
}

//
// PAOI = Purchase or Accepted Offer Indicator
//

// GetPAOI returns a string indicating if the tx is a purchase or someone dumped into bids.
func (ttx *TokenTransaction) GetPAOI() string {
	indicatorString := "・"

	return ttx.getPAOIStyle().Render(indicatorString)
}

// getPAOIStyle returns a lipgloss style for the "purchase or accepted offer indicator".
func (ttx *TokenTransaction) getPAOIStyle() lipgloss.Style {
	var purchaseOrBidStyle lipgloss.Style

	switch {
	case ttx.IsAcceptedOffer():
		purchaseOrBidStyle = style.TrendRedStyle

	case ttx.IsListing():
		purchaseOrBidStyle = style.OpenSea

	case ttx.IsCollectionOffer():
		purchaseOrBidStyle = style.PurplePower

	case ttx.IsTransfer():
		purchaseOrBidStyle = style.DarkGrayStyle

	default:
		purchaseOrBidStyle = style.TrendLightGreenStyle
	}

	return purchaseOrBidStyle
}

func (ttx *TokenTransaction) Is() map[string]bool {
	isFunctions := map[string]bool{
		"IsAcceptedOffer":   ttx.IsAcceptedOffer(),
		"IsAirdrop":         ttx.IsAirdrop(),
		"IsBurn":            ttx.IsBurn(),
		"IsCollectionOffer": ttx.IsCollectionOffer(),
		"IsItemBid":         ttx.IsItemBid(),
		"IsListing":         ttx.IsListing(),
		"IsLoan":            ttx.IsLoan(),
		"IsLoanPayback":     ttx.IsLoanPayback(),
		"IsMint":            ttx.IsMint(),
		"IsMovingNFTs":      ttx.IsMovingNFTs(),
		"IsReBurn":          ttx.IsReBurn(),
		"IsTransfer":        ttx.IsTransfer(),
	}

	return isFunctions
}

func (ttx *TokenTransaction) FormattedIs() []string {
	fmtIsFunctions := make([]string, 0)

	for k, v := range ttx.Is() {
		var fmtVal string

		if v {
			fmtVal = style.TrendGreenStyle.Render(strconv.FormatBool(v))
		} else {
			fmtVal = style.TrendRedStyle.Render(strconv.FormatBool(v))
		}

		fmtIsFunctions = append(fmtIsFunctions, fmt.Sprintf("%s(): %s", k, fmtVal))
	}

	return fmtIsFunctions
}

func (ttx *TokenTransaction) IsMovingNFTs() bool {
	return len(ttx.logsByStandard[standard.ERC721]) > 0 || len(ttx.logsByStandard[standard.ERC1155]) > 0
}

func (ttx *TokenTransaction) IsListing() bool {
	return ttx.Action == degendb.Listing
}

func (ttx *TokenTransaction) IsItemBid() bool {
	return ttx.Action == degendb.Bid
}

func (ttx *TokenTransaction) IsAcceptedOffer() bool {
	return ttx.GetNFTSenderAddresses().Cardinality() > 0 && ttx.GetNFTSenderAddresses().Contains(ttx.From)
}

func (ttx *TokenTransaction) IsCollectionOffer() bool {
	return ttx.Action == degendb.CollectionOffer
}

func (ttx *TokenTransaction) IsMint() bool {
	// if no nfts are moved, this is not a mint
	if !ttx.IsMovingNFTs() {
		return false
	}

	senders := ttx.GetNFTSenders()
	receivers := ttx.GetNFTReceivers()

	// if there are multiple senders or receivers, this is not a (typical) mint
	if len(senders) != 1 || len(receivers) != 1 {
		return false
	}

	// mints comes always from the zero address and never go to it
	if senders[internal.ZeroAddress] == nil {
		return false
	}

	if receivers[internal.ZeroAddress] != nil {
		return false
	}

	return true
}

func (ttx *TokenTransaction) IsAirdrop() bool {
	// if no nfts are moved, this is not a mint
	if !ttx.IsMovingNFTs() {
		return false
	}

	// airdrops are always free
	if ttx.AmountPaid.Cmp(big.NewInt(0)) > 0 {
		return false
	}

	senders := ttx.GetNFTSenderAddresses()
	receivers := ttx.GetNFTReceiverAddresses()

	// airdrops come from the zeroAddress
	if senders.Cardinality() != 1 || senders.Contains(internal.ZeroAddress) {
		return false
	}

	// airdrops are sent to multiple addresses
	if receivers.Cardinality() < 2 {
		return false
	}

	return true
}

func (ttx *TokenTransaction) IsReBurn() bool {
	// we define a "re-burn", a return-burn 😁😂, as a tx that moves one or more nfts to the zero address
	// and gets one or more nfts from the zero address back
	// -> the currently typical burn events
	senders := ttx.GetNFTSenders()
	receivers := ttx.GetNFTReceivers()

	// there must be exactly two sender and receiver
	// the zero address and one to send/receive the nfts
	if len(senders) != 2 || len(receivers) != 2 {
		return false
	}

	for sender := range senders {
		if receivers[sender] == nil {
			return false
		}
	}

	if receivers[internal.ZeroAddress] == nil || senders[internal.ZeroAddress] == nil {
		return false
	}

	return true
}

func (ttx *TokenTransaction) IsBurn() bool {
	// a burn is a costless transfer/tx that moves one or more nfts to the zero address
	if ttx.AmountPaid.Cmp(big.NewInt(0)) != 0 {
		return false
	}

	receivers := ttx.GetNFTReceivers()

	// there must be exactly one receiver and it must be the zero address
	if len(receivers) != 1 || receivers[internal.ZeroAddress] == nil {
		return false
	}

	return true
}

func (ttx *TokenTransaction) IsLoan() bool {
	if !ttx.IsMovingNFTs() {
		return false
	}

	txContainsLoanNFT := ttx.GetTransferredTokenContractAdresses().Intersect(marketplace.TokenAddresses())
	if txContainsLoanNFT.Cardinality() > 0 {
		log.Printf("")
		log.Printf("  🐄 txContainsLoanNFT: %+v", txContainsLoanNFT)

		// txSenderAndLoanContractAreNFTSender := ttx.GetNFTSenderAddresses().Contains(ttx.From) && ttx.GetNFTSenderAddresses().ContainsAny(marketplace.LoanContracts.ToSlice()...)
		// log.Printf("  🐄 txSenderAndLoanContractAreNFTSender: %+v", txSenderAndLoanContractAreNFTSender)

		txSenderAndZeroAreNFTSender := ttx.GetNFTSenderAddresses().Contains(ttx.From) && ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress)
		log.Printf("  🐄 txSenderAndZeroAreNFTSender: %+v", txSenderAndZeroAreNFTSender)

		loanContractIsNFTReceiver := ttx.GetNFTReceiverAddresses().ContainsAny(marketplace.LoanContracts.ToSlice()...)
		log.Printf("  🐄 loanContractIsNFTReceiver: %+v", loanContractIsNFTReceiver)

		// txSenderAndZeroAreNFTReceiver := ttx.GetNFTReceiverAddresses().Contains(ttx.From) && ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress)
		// log.Printf("  🐄 txSenderAndZeroAreNFTReceiver: %+v", txSenderAndZeroAreNFTReceiver)

		log.Printf("  🐄 → %+v %+v %+v | is loan: %+v", txContainsLoanNFT, txSenderAndZeroAreNFTSender, loanContractIsNFTReceiver, txContainsLoanNFT.Cardinality() > 0 && txSenderAndZeroAreNFTSender && loanContractIsNFTReceiver)
		log.Printf("")
	}

	// if ttx.GetTransferredTokenContractAdresses().ContainsAny(marketplace.LoanContracts.ToSlice()...) {
	// 	zeroAddressInOne := ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress) || ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress)
	// 	zeroAddressInBoth := ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress) && ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress)

	// 	ttyTy := "in none"

	// 	if zeroAddressInOne {
	// 		ttyTy = "←LOAN"
	// 	} else if zeroAddressInBoth {
	// 		ttyTy = "→REPAY"
	// 	}

	// 	log.Infof("")
	// 	// log.Infof(pretty.Sprint(ttx.Is()))
	// 	log.Infof("isLoan | NFTSenderAddresses: %+v", ttx.GetNFTSenderAddresses())
	// 	log.Infof("isLoan | NFTReceiverAddresses: %+v", ttx.GetNFTReceiverAddresses())
	// 	log.Infof("isLoan | zeroAddressInOne: %+v | zeroAddressInBoth: %+v", zeroAddressInOne, zeroAddressInBoth)
	// 	if zeroAddressInOne {
	// 		log.Infof("isLoan | → %+v", ttyTy)
	// 	}
	// 	log.Infof("")

	// 	log.Infof("")
	// 	log.Infof("isLoan | NFTSenderAddresses: %+v", ttx.GetNFTSenderAddresses())
	// 	log.Infof("isLoan | NFTReceiverAddresses: %+v", ttx.GetNFTReceiverAddresses())
	// 	log.Infof("isLoan | zeroAddressInOne: %+v | zeroAddressInBoth: %+v", zeroAddressInOne, zeroAddressInBoth)
	// 	if zeroAddressInOne {
	// 		log.Infof("isLoan | → %+v", ttyTy)
	// 	}
	// 	log.Infof("")

	// 	// log.Printf("")
	// 	// log.Printf("")
	// 	// log.Printf(" LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??")

	// 	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress), ttx.GetNFTSenderAddresses())
	// 	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTSenderAddresses().ContainsAny(internal.ZeroAddress), ttx.GetNFTSenderAddresses())
	// 	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTSenderAddresses().Contains(common.Address{}), ttx.GetNFTSenderAddresses())
	// 	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTSenderAddresses().ContainsAny(common.Address{}), ttx.GetNFTSenderAddresses())
	// 	// log.Printf("  🧱 → %v  |   %+v", ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress), ttx.GetNFTSenderAddresses())

	// 	if ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress) {
	// 		return true
	// 	}

	// 	log.Printf(" LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??  LOAN??")
	// 	log.Printf("")
	// 	log.Printf("")

	// 	// return true
	// }

	for tokenAddress := range ttx.GetTransfersByContract() {
		if marketplace.LoanContracts.Contains(tokenAddress) {
			log.Printf("  🧱 ← %v | %+v  in  %+v", ttx.GetNFTReceiverAddresses().Contains(tokenAddress), tokenAddress, ttx.GetNFTReceiverAddresses())
			log.Printf("  🧱 → %v | %+v  in  %+v", ttx.GetNFTSenderAddresses().Contains(tokenAddress), tokenAddress, ttx.GetNFTSenderAddresses())
			log.Printf("")
			log.Printf("")

			return true
		}
	}

	return false
}

func (ttx *TokenTransaction) IsLoanPayback() bool {
	if !ttx.IsMovingNFTs() {
		return false
	}

	// log.Printf("")
	// log.Printf("")
	// log.Printf(" REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??")

	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress), ttx.GetNFTReceiverAddresses())
	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTReceiverAddresses().ContainsAny(internal.ZeroAddress), ttx.GetNFTReceiverAddresses())
	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTReceiverAddresses().Contains(common.Address{}), ttx.GetNFTReceiverAddresses())
	// log.Printf("  🧱 ← %v  |   %+v", ttx.GetNFTReceiverAddresses().ContainsAny(common.Address{}), ttx.GetNFTReceiverAddresses())
	// log.Printf("  🧱 → %v  |   %+v", ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress), ttx.GetNFTReceiverAddresses())

	// zeroAddressInOne := ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress) || ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress)
	// zeroAddressInBoth := ttx.GetNFTSenderAddresses().Contains(internal.ZeroAddress) && ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress)

	// ttyTy := "in none"

	// if zeroAddressInOne {
	// 	ttyTy = "←LOAN"
	// } else if zeroAddressInBoth {
	// 	ttyTy = "→REPAY"
	// }

	// log.Infof("")
	// // log.Infof(pretty.Sprint(ttx.Is()))
	// log.Infof("isPayback | NFTSenderAddresses: %+v", ttx.GetNFTSenderAddresses())
	// log.Infof("isPayback | NFTReceiverAddresses: %+v", ttx.GetNFTReceiverAddresses())
	// log.Infof("isPayback | zeroAddressInOne: %+v | zeroAddressInBoth: %+v", zeroAddressInOne, zeroAddressInBoth)
	// if zeroAddressInOne {
	// 	log.Infof("isPayback | → %+v", ttyTy)
	// }
	// log.Infof("")

	// log.Infof("")
	// log.Infof("isPayback | NFTSenderAddresses: %+v", ttx.GetNFTSenderAddresses())
	// log.Infof("isPayback | NFTReceiverAddresses: %+v", ttx.GetNFTReceiverAddresses())
	// log.Infof("isPayback | zeroAddressInOne: %+v | zeroAddressInBoth: %+v", zeroAddressInOne, zeroAddressInBoth)
	// if zeroAddressInOne {
	// 	log.Infof("isPayback | → %+v", ttyTy)
	// }
	// log.Infof("")

	// if ttx.GetNFTReceiverAddresses().Contains(internal.ZeroAddress) {
	// 	return true
	// }

	// log.Printf(" REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??  REPAY??")
	// log.Printf("")
	// log.Printf("")

	for tokenAddress := range ttx.GetTransfersByContract() {
		if marketplace.LoanContracts.Contains(tokenAddress) {
			log.Printf("  🧱 ← %v | %+v  in  %+v", ttx.GetNFTReceiverAddresses().Contains(tokenAddress), tokenAddress, ttx.GetNFTReceiverAddresses())
			log.Printf("  🧱 → %v | %+v  in  %+v", ttx.GetNFTSenderAddresses().Contains(tokenAddress), tokenAddress, ttx.GetNFTSenderAddresses())
			log.Printf("")
			log.Printf("")

			return true
		}
	}

	return false
}

func (ttx *TokenTransaction) IsTransfer() bool {
	// opensea transfer helper contract
	if ttx.Tx == nil || ttx.Tx.To() == nil || (*ttx.Tx.To() == common.HexToAddress("0x0000000000c2d145a2526bd8c716263bfebe1a72")) {
		return true
	}

	if len(ttx.Tx.Data()) < 4 {
		return false
	}

	if ttx.AmountPaid.Cmp(big.NewInt(0)) != 0 {
		return false
	}

	if ttx.Marketplace != nil && ttx.Marketplace != &marketplace.Unknown {
		return false
	}

	for _, transfer := range ttx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			if transfer.To == internal.ZeroAddress || transfer.From == internal.ZeroAddress {
				return false
			}
		}
	}

	// methodSignature := ttx.Tx.Data()[0:4]

	// if methodSignaturesTransfers[[4]byte(methodSignature)] == "" {
	// 	log.Debugf("wrong method signature: %x", methodSignature)

	// 	return false
	// }

	return true
}

func (ttx *TokenTransaction) getAction() *degendb.GBEventType {
	if !ttx.IsMovingNFTs() {
		return degendb.Unknown
	}

	switch {
	case ttx.IsMint():
		return degendb.Mint
	case ttx.IsLoan():
		return degendb.Loan
	case ttx.Marketplace != nil && ttx.Marketplace != &marketplace.Unknown:
		return degendb.Sale
	// case ttx.IsLoanPayback():
	// 	return degendb.RepayLoan
	case ttx.IsAirdrop():
		return degendb.Airdrop
	case ttx.IsReBurn():
		return degendb.BurnRedeem
	case ttx.IsBurn():
		return degendb.Burn
	case ttx.AmountPaid.Cmp(big.NewInt(0)) > 0:
		return degendb.Sale
	case ttx.IsTransfer():
		return degendb.Transfer
	default:
		return degendb.Unknown
	}
}

func parseERC721TransferLog(txLog *types.Log) *TokenTransfer {
	return &TokenTransfer{
		From:                common.HexToAddress(txLog.Topics[1].String()),
		To:                  common.HexToAddress(txLog.Topics[2].String()),
		AmountTokens:        big.NewInt(1),
		AmountEtherReturned: big.NewInt(0),
		Standard:            standard.ERC721,
		Token: &token.Token{
			Address: txLog.Address,
			ID:      txLog.Topics[3].Big(),
		},
	}
}

func parseERC1155TransferLog() *TokenTransfer {
	// reactivate me
	// // abiERC1155, err := abis.NewERC1155(txLog.Address, ethNode.Client)
	// abiERC1155, err := providerPool.GetERC1155ABI(txLog.Address)
	// if err != nil {
	// 	log.Errorf("❗️ error binding erc1155 contract abi: %s", err)

	// 	return nil
	// }

	return nil

	// transferLog, err := abiERC1155.ParseTransferSingle(*txLog)
	// if err != nil {
	// 	log.Errorf("❗️ error parsing TransferSingle log: %s", err)

	// 	return nil
	// }

	// return &TokenTransfer{
	// 	From:                transferLog.From,
	// 	To:                  transferLog.To,
	// 	AmountTokens:        transferLog.Value,
	// 	AmountEtherReturned: big.NewInt(0),
	// 	Standard:            standard.ERC1155,
	// 	Token: &token.Token{
	// 		Address: transferLog.Raw.Address,
	// 		ID:      transferLog.Id,
	// 	},
	// }
}

func parseERC20TransferLog() *TokenTransfer {
	// reactivate me
	// abiWETH, err := providerPool.GetWETHABI(txLog.Address)
	// if err != nil {
	// 	log.Errorf("❗️ error binding erc721 contract abi: %s", err)

	// 	return nil
	// }

	// transferLog, err := abiWETH.ParseTransfer(*txLog)
	// if err != nil {
	// 	log.Infof("❗️ error parsing transfer log (%d topics): %s", len(txLog.Topics), err)

	// 	return nil
	// }

	// // we only care about certain tokens like WETH and Blur Pool Token
	// if transferLog.Raw.Address != internal.WETHContractAddress && transferLog.Raw.Address != internal.BlurPoolTokenContractAddress {
	// 	log.Debugf("❗️ non-WETH ERC20 token, ignoring: %s", transferLog.Raw.Address.String())

	// 	return nil
	// }

	// // handle blur pool txs
	// amount := transferLog.Wad

	// if transferLog.Raw.Address == internal.BlurPoolTokenContractAddress && len(txLog.Topics) == 3 {
	// 	amount = new(big.Int).SetBytes(transferLog.Raw.Data)
	// }

	// tokenTransfer := &TokenTransfer{
	// 	From:                transferLog.Src,
	// 	To:                  transferLog.Dst,
	// 	AmountTokens:        amount,
	// 	AmountEtherReturned: big.NewInt(0),
	// 	Standard:            standard.ERC20,
	// 	Token: &token.Token{
	// 		Address: transferLog.Raw.Address,

	// 		// set a random ID for ERC20 tokens
	// 		ID: big.NewInt(-1),
	// 		// ID: big.NewInt(0).Add(big.NewInt(rand.Int63n(1337)), amount), //nolint:gosec
	// 	},
	// }

	// return tokenTransfer

	return nil
}

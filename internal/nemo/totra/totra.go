package totra

import (
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/topic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type TokenTransaction struct {
	// the original tx
	Tx *types.Transaction `json:"tx"`

	// the receipt containing the logs
	TxReceipt *types.Receipt `json:"tx_receipt"`

	// action performed by the tx
	Action TxType `json:"action"`

	// the sender of the tx
	From common.Address `json:"from"`

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
}

// var methodSignaturesTransfers = map[[4]byte]string{
// 	{0x23, 0xb8, 0x72, 0xdd}: "transferFrom(address,address,uint256)",
// 	{0x42, 0x84, 0x2e, 0x0e}: "safeTransferFrom(address,address,uint256)",
// 	{0xb8, 0x8d, 0x4f, 0xde}: "safeTransferFrom(address,address,uint256,bytes)",
// 	{0xf2, 0x42, 0x43, 0x2a}: "safeTransferFrom(address,address,uint256,uint256,bytes)",
// 	{0x32, 0x38, 0x9b, 0x71}: "bulkTransfer(((uint8,address,uint256,uint256)[],address,bool)[],bytes32)",
// }

var LoanContracts = map[common.Address]string{
	common.HexToAddress("0x5660E206496808F7b5cDB8C56A696a96AE5E9b23"): "NFTfi",
	common.HexToAddress("0x0E258c84Df0f8728ae4A6426EA5FD163Eb6b9D1B"): "NFT Loan Ticket V2",
	common.HexToAddress("0xbD85BF4C970b91984e6A2b8Ba9C577A58A8C20f9"): "Borrower Note Ticket",
}

func NewTokenTransaction(tx *types.Transaction, receipt *types.Receipt, providerPool *provider.Pool) *TokenTransaction {
	tfLogsByStandard := make(map[standard.Standard][]*types.Log, 0)

	for _, txLog := range receipt.Logs {
		if len(txLog.Topics) == 0 {
			continue
		}

		logStandard := getTransferLogStandard(txLog)

		tfLogsByStandard[logStandard] = append(tfLogsByStandard[logStandard], txLog)
	}

	msg, err := core.TransactionToMessage(tx, types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		gbl.Log.Warnf("could not get message for tx %s: %s", tx.Hash().Hex(), err)
	}

	ttx := &TokenTransaction{
		Tx:             tx,
		TxReceipt:      receipt,
		From:           msg.From,
		logsByStandard: tfLogsByStandard,
		Transfers:      make([]*TokenTransfer, 0),
		AmountPaid:     tx.Value(),

		ReceivedAt: time.Now(),

		// print all tx by default
		DoNotPrint: false,
	}

	// marketplace
	switch {
	case marketplace.OpenSea.ContractAddresses[*tx.To()]:
		ttx.Marketplace = &marketplace.OpenSea
	case marketplace.Blur.ContractAddresses[*tx.To()]:
		ttx.Marketplace = &marketplace.Blur
	case marketplace.X2Y2.ContractAddresses[*tx.To()]:
		ttx.Marketplace = &marketplace.X2Y2
	default:
		ttx.Marketplace = &marketplace.Unknown
	}

	// parse transfers from logs to get the amount paid and other data
	ttx.parseTransfersFromReceipt(providerPool)

	// erc20
	ttx.parseERC20Transfers()

	// connect nft transfers and erc20 transfers
	ttx.discoverItemPrices()

	// action performed by the tx
	ttx.Action = ttx.getAction()

	if len(ttx.Transfers) == 0 {
		gbl.Log.Debugf("  🧱 no transfers found for ttx: %+v", ttx)

		return nil
	}

	return ttx
}

func (ttx *TokenTransaction) GetTransfersByContract() map[common.Address][]*TokenTransfer {
	transfersByContract := make(map[common.Address][]*TokenTransfer, 0)

	for _, transfer := range ttx.Transfers {
		transfersByContract[transfer.Token.Address] = append(transfersByContract[transfer.Token.Address], transfer)
	}

	return transfersByContract
}

func (ttx *TokenTransaction) GetPrice() *price.Price {
	return price.NewPrice(ttx.AmountPaid)
}

func (ttx *TokenTransaction) GetNFTReceivers() map[common.Address][]*TokenTransfer {
	nftReceivers := make(map[common.Address][]*TokenTransfer, 0)

	for _, transfer := range ttx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			nftReceivers[transfer.To] = append(nftReceivers[transfer.To], transfer)
		}
	}

	return nftReceivers
}

func (ttx *TokenTransaction) GetNFTReceiverAddresses() []common.Address {
	nftReceivers := ttx.GetNFTReceivers()

	receivers := make([]common.Address, 0)

	for receiver := range nftReceivers {
		receivers = append(receivers, receiver)
	}

	return receivers
}

func (ttx *TokenTransaction) GetNFTSenders() map[common.Address][]*TokenTransfer {
	nftSenders := make(map[common.Address][]*TokenTransfer, 0)

	for _, transfer := range ttx.Transfers {
		if transfer.Standard.IsERC721orERC1155() {
			nftSenders[transfer.From] = append(nftSenders[transfer.From], transfer)
		}
	}

	return nftSenders
}

func (ttx *TokenTransaction) GetNonZeroNFTSenders() map[common.Address][]*TokenTransfer {
	nftSenders := ttx.GetNFTSenders()

	nonZeroSenders := make(map[common.Address][]*TokenTransfer, 0)

	for addr, sender := range nftSenders {
		if addr == internal.ZeroAddress {
			continue
		}

		nonZeroSenders[addr] = sender
	}

	return nonZeroSenders
}

func (ttx *TokenTransaction) GetNFTSenderAddresses() []common.Address {
	nftSenders := ttx.GetNFTSenders()

	senders := make([]common.Address, 0)

	for sender := range nftSenders {
		senders = append(senders, sender)
	}

	return senders
}

func (ttx *TokenTransaction) GetNFTSenderAndReceiverAddresses() []common.Address {
	addresses := make([]common.Address, 0)
	addresses = append(addresses, ttx.GetNFTSenderAddresses()...)
	addresses = append(addresses, ttx.GetNFTReceiverAddresses()...)

	return addresses
}

func (ttx *TokenTransaction) parseTransfersFromReceipt(providerPool *provider.Pool) {
	// assuming every nft is just sold once per tx
	uniqueTransfers := make(map[string][]*TokenTransfer, 0)

	for logStandard, txLogs := range ttx.logsByStandard {
		gbl.Log.Debugf("  🧱 ttx logs to parse: %+v", len(txLogs))

		for _, txLog := range txLogs {
			// parse Transfer & TransferSingle logs
			var transfer *TokenTransfer

			switch logStandard {
			case standard.ERC20:
				transfer = parseERC20TransferLog(txLog, providerPool)

			case standard.ERC721:
				transfer = parseERC721TransferLog(txLog)

			case standard.ERC1155:
				transfer = parseERC1155TransferLog(txLog, providerPool)
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
// 		gbl.Log.Debugf("  🧱 blockParser | ttx log: %+v", txLog)

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
		if transfer.Standard == standard.ERC20 {
			if _, ok := ttx.sentMoney[transfer.From]; !ok {
				ttx.sentMoney[transfer.From] = big.NewInt(0)
			}

			ttx.sentMoney[transfer.From].Add(ttx.sentMoney[transfer.From], transfer.AmountTokens)

			amountPaidERC20.Add(amountPaidERC20, transfer.AmountTokens)
		}

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

func (ttx *TokenTransaction) IsMovingNFTs() bool {
	return len(ttx.logsByStandard[standard.ERC721]) > 0 || len(ttx.logsByStandard[standard.ERC1155]) > 0
}

func (ttx *TokenTransaction) IsListing() bool {
	return ttx.Action == Listing
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
	if len(senders) != 1 || senders[0] != internal.ZeroAddress {
		return false
	}

	// airdrops are sent to multiple addresses
	if len(receivers) < 2 {
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
	// if no nfts are moved, this is not a mint
	if !ttx.IsMovingNFTs() {
		return false
	}

	for tokenAddress := range ttx.GetTransfersByContract() {
		if LoanContracts[tokenAddress] != "" {
			return true
		}
	}

	return false
}

func (ttx *TokenTransaction) IsTransfer() bool {
	// opensea transfer helper contract
	if *ttx.Tx.To() == common.HexToAddress("0x0000000000c2d145a2526bd8c716263bfebe1a72") {
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
	// 	gbl.Log.Debugf("wrong method signature: %x", methodSignature)

	// 	return false
	// }

	return true
}

func (ttx *TokenTransaction) getAction() TxType {
	if !ttx.IsMovingNFTs() {
		return Unknown
	}

	switch {
	case ttx.IsMint():
		return Mint
	case ttx.IsLoan():
		return Loan
	case ttx.Marketplace != nil && ttx.Marketplace != &marketplace.Unknown:
		return Sale
	case ttx.IsAirdrop():
		return Airdrop
	case ttx.IsReBurn():
		return ReBurn
	case ttx.IsBurn():
		return Burn
	case ttx.AmountPaid.Cmp(big.NewInt(0)) > 0:
		return Sale
	case ttx.IsTransfer():
		return Transfer
	default:
		return Unknown
	}
}

func (ttx *TokenTransaction) AsHistoryTokenTransaction(collection *collections.Collection, fmtTokensTransferred []string) *HistoryTokenTransaction {
	return &HistoryTokenTransaction{
		ReceivedAt: ttx.ReceivedAt,
		AmountPaid: ttx.AmountPaid,
		// TxType:               ttx.GetTxType(),
		FmtTokensTransferred: fmtTokensTransferred,
		Collection:           collection,

		TokenTransaction: ttx,
	}
}

func getTransferLogStandard(log *types.Log) standard.Standard {
	logStandard := standard.UNKNOWN

	topic0 := topic.Topic(log.Topics[0].String())

	switch {
	// erc20
	case topic0 == topic.Transfer && len(log.Topics) <= 3:
		logStandard = standard.ERC20

	// erc721
	case topic0 == topic.Transfer && len(log.Topics) >= 4:
		logStandard = standard.ERC721

	// erc1155
	case topic0 == topic.TransferSingle && len(log.Topics) >= 4:
		logStandard = standard.ERC1155

	default:
		gbl.Log.Debugf("unknown log standard | len(log.Topics): %d | topic0: %s", len(log.Topics), topic0)
	}

	return logStandard
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

func parseERC1155TransferLog(txLog *types.Log, providerPool *provider.Pool) *TokenTransfer {
	// abiERC1155, err := abis.NewERC1155(txLog.Address, ethNode.Client)
	abiERC1155, err := providerPool.GetERC1155ABI(txLog.Address)
	if err != nil {
		gbl.Log.Errorf("❗️ error binding erc1155 contract abi: %s", err)

		return nil
	}

	transferLog, err := abiERC1155.ParseTransferSingle(*txLog)
	if err != nil {
		gbl.Log.Errorf("❗️ error parsing TransferSingle log: %s", err)

		return nil
	}

	return &TokenTransfer{
		From:                transferLog.From,
		To:                  transferLog.To,
		AmountTokens:        transferLog.Value,
		AmountEtherReturned: big.NewInt(0),
		Standard:            standard.ERC1155,
		Token: &token.Token{
			Address: transferLog.Raw.Address,
			ID:      transferLog.Id,
		},
	}
}

func parseERC20TransferLog(txLog *types.Log, providerPool *provider.Pool) *TokenTransfer {
	abiWETH, err := providerPool.GetWETHABI(txLog.Address)
	if err != nil {
		gbl.Log.Errorf("❗️ error binding erc721 contract abi: %s", err)

		return nil
	}

	transferLog, err := abiWETH.ParseTransfer(*txLog)
	if err != nil {
		gbl.Log.Infof("❗️ error parsing transfer log (%d topics): %s", len(txLog.Topics), err)

		return nil
	}

	// we only care about certain tokens like WETH and Blur Pool Token
	if transferLog.Raw.Address != internal.WETHContractAddress && transferLog.Raw.Address != internal.BlurPoolTokenContractAddress {
		gbl.Log.Debugf("❗️ non-WETH ERC20 token, ignoring: %s", transferLog.Raw.Address.String())

		return nil
	}

	// handle blur pool txs
	amount := transferLog.Wad

	if transferLog.Raw.Address == internal.BlurPoolTokenContractAddress && len(txLog.Topics) == 3 {
		amount = new(big.Int).SetBytes(transferLog.Raw.Data)
	}

	tokenTransfer := &TokenTransfer{
		From:                transferLog.Src,
		To:                  transferLog.Dst,
		AmountTokens:        amount,
		AmountEtherReturned: big.NewInt(0),
		Standard:            standard.ERC20,
		Token: &token.Token{
			Address: transferLog.Raw.Address,

			// set a random ID for ERC20 tokens
			ID: big.NewInt(-1),
			// ID: big.NewInt(0).Add(big.NewInt(rand.Int63n(1337)), amount), //nolint:gosec
		},
	}

	return tokenTransfer
}

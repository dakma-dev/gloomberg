package seaport

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/gloomberg"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Seaport11Address = common.HexToAddress("0x00000000006c3852cbef3e08e8df289169ede581")
	SeaportAddress   = Seaport11Address
)

func getKeyAndAddressFromString(key string) (*ecdsa.PrivateKey, common.Address, error) {
	// parse private key
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		gbl.Log.Error("❌ error while parsing private key: ", err)
		return nil, utils.ZeroAddress, err
	}

	// derive public key
	publicKey := privateKey.Public()

	// get ecdsa public key
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		gbl.Log.Error("❌ error casting public key to ECDSA")
		return nil, utils.ZeroAddress, err
	}

	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicAddress, nil
}

func FulfillBasicOrder(gb *gloomberg.Gloomberg, order *models.SeaportOrder, privKey string) (*types.Transaction, error) {
	// get key and address
	privateKey, publicAddress, err := getKeyAndAddressFromString(privKey)
	if err != nil {
		gbl.Log.Error("❌ error while getting key and address from string: ", err)
		return nil, err
	}

	// eth client
	client := gb.Nodes.GetRandomLocalNode().Client

	// seaport binding
	abiSeaport, err := abis.NewSeaport(SeaportAddress, gb.Nodes.GetRandomLocalNode().Client)
	if err != nil {
		gbl.Log.Error(err)
	}

	////////////////////////////////////////////////////////////////////////////////////////

	params := order.ProtocolData.Parameters

	// decode the salt
	salt, err := hexutil.DecodeBig(params.Salt)
	if err != nil {
		gbl.Log.Error(err)
	}

	// build the recipients list
	additionalRecipients := make([]abis.AdditionalRecipient, 0)
	for _, recipient := range params.Consideration {
		additionalRecipients = append(additionalRecipients, abis.AdditionalRecipient{
			Recipient: recipient.Recipient,
			Amount:    big.NewInt(recipient.StartAmount),
		})
	}

	orderParameters := abis.BasicOrderParameters{
		ConsiderationToken:      utils.ZeroAddress,
		ConsiderationAmount:     big.NewInt(params.Consideration[0].StartAmount),
		ConsiderationIdentifier: big.NewInt(params.Consideration[0].IdentifierOrCriteria),

		Offerer: params.Offerer,

		OfferToken:      params.Offer[0].Token,
		OfferIdentifier: big.NewInt(params.Offer[0].IdentifierOrCriteria),
		OfferAmount:     big.NewInt(params.Offer[0].StartAmount),

		BasicOrderType: uint8(params.OrderType),

		OffererConduitKey:   params.ConduitKey,
		FulfillerConduitKey: params.ConduitKey,

		TotalOriginalAdditionalRecipients: big.NewInt(int64(params.TotalOriginalConsiderationItems - 1)),

		AdditionalRecipients: additionalRecipients,

		StartTime: big.NewInt(params.StartTime),
		EndTime:   big.NewInt(params.EndTime),

		Salt: salt,

		Zone:     params.Zone,
		ZoneHash: params.ZoneHash,

		Signature: []byte(order.ProtocolData.Signature),
	}
	gbl.Log.Debugf("orderParameters: %+v", orderParameters)

	basicOrderParameters := abis.BasicOrderParameters{
		ConsiderationToken:      utils.ZeroAddress,
		ConsiderationAmount:     big.NewInt(270000000000000),
		ConsiderationIdentifier: big.NewInt(0),

		OfferToken:      common.HexToAddress("0x799CcCb6D0C57345c4DAB303c81A5CDffbCDEa76"),
		OfferIdentifier: big.NewInt(5334),
		OfferAmount:     big.NewInt(1),

		BasicOrderType: uint8(2),

		OffererConduitKey:   common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000"),
		FulfillerConduitKey: common.HexToHash("0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000"),

		TotalOriginalAdditionalRecipients: big.NewInt(2),

		AdditionalRecipients: []abis.AdditionalRecipient{
			{
				Amount:    big.NewInt(7500000000000),
				Recipient: common.HexToAddress("0x0000a26b00c1F0DF003000390027140000fAa719"),
			},
			{
				Amount:    big.NewInt(22500000000000),
				Recipient: common.HexToAddress("0xeAFaA00e16E7D47508569B561f87C8ED4F1A559b"),
			},
		},

		Offerer: common.HexToAddress("0x0000a83deaa073245cfbde660fd8daf09d78de00"),

		StartTime: big.NewInt(1667821610),
		EndTime:   big.NewInt(1670413610),

		Salt: salt,

		Zone:     common.HexToAddress("0x004C00500000aD104D7DBd00e3ae0A5C00560C00"),
		ZoneHash: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),

		Signature: []byte{0x94, 0x51, 0x26, 0x7c, 0xf9, 0x91, 0x3e, 0x7a, 0x99, 0xe6, 0xeb, 0x9f, 0xc4, 0xc1, 0x7d, 0xdf, 0x30, 0x03, 0xcf, 0x0c, 0x12, 0xba, 0xc6, 0xf3, 0xb4, 0x3b, 0x05, 0x2b, 0xf7, 0xf4, 0xd6, 0x20, 0x06, 0x07, 0x5b, 0xd4, 0x09, 0x4a, 0x87, 0xf7, 0x53, 0xfb, 0x3c, 0xc7, 0xe4, 0x1a, 0xd4, 0x78, 0x1b, 0xcb, 0xc0, 0xb3, 0xbe, 0x27, 0xd1, 0xa8, 0x10, 0xb3, 0xac, 0x45, 0xa3, 0xe5, 0x1e, 0x16, 0x1b},
	}
	gbl.Log.Debugf("basicOrderParameters: %+v", basicOrderParameters)

	//
	// create the tx

	// get next nonce
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		gbl.Log.Error(err)
	}

	// get suggested gas settings
	gasFeeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		gbl.Log.Error(err)
	}

	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		gbl.Log.Error(err)
	}

	// eth_estimateGas todo!
	gasLimit := uint64(160000)

	// create & set tx options
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	if err != nil {
		gbl.Log.Errorf("❌ NewKeyedTransactorWithChainID error: %v", err.Error())
	}

	txOpts.From = publicAddress
	txOpts.GasFeeCap = gasFeeCap
	txOpts.GasTipCap = gasTipCap
	txOpts.GasLimit = uint64(gasLimit)            // in units
	txOpts.Value = big.NewInt(order.CurrentPrice) // in wei
	txOpts.Nonce = big.NewInt(int64(nonce))

	// disable sending the tx to a node (for testing/debugging)
	txOpts.NoSend = false
	gbl.Log.Debugf("txOpts: %+v", txOpts)

	// send the tx
	tx, err := abiSeaport.FulfillBasicOrder(txOpts, orderParameters)
	if err != nil {
		gbl.Log.Errorf("❌ FulfillBasicOrder error: %v", err.Error())
	}

	gbl.Log.Debugf("tx.Data: %+v", tx)
	gbl.Log.Debugf("tx.Data: %x", tx.Data())
	gbl.Log.Debugf("tx.Hash: %s", tx.Hash().String())

	// wait for tx to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		gbl.Log.Errorf("❌ WaitMined error: %v", err.Error())
	}

	gbl.Log.Debugf("receipt: %+v", receipt)

	return tx, nil
}

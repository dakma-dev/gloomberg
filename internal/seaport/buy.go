package seaport

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"strconv"

	"github.com/benleb/gloomberg/internal/abis"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/gloomberg"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gookit/goutil/dump"
)

var OpenSeaSeaportContract = common.HexToAddress("0x00000000006c3852cbef3e08e8df289169ede581")

func sToI64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		gbl.Log.Errorf("❌ error parsing int64: %s", err.Error())
		return 0
	}

	return val
}

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
	abiSeaport, err := abis.NewSeaport(OpenSeaSeaportContract, gb.Nodes.GetRandomLocalNode().Client)
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
	for _, recipient := range params.Consideration[1:] {
		additionalRecipients = append(additionalRecipients, abis.AdditionalRecipient{
			Recipient: recipient.Recipient,
			Amount:    big.NewInt(sToI64(recipient.StartAmount)),
		})
	}

	signature, err := hexutil.Decode(order.ProtocolData.Signature)
	if err != nil {
		gbl.Log.Error(err)
	}

	orderParameters := abis.BasicOrderParameters{
		ConsiderationToken:      utils.ZeroAddress,
		ConsiderationAmount:     big.NewInt(sToI64(params.Consideration[0].StartAmount)),
		ConsiderationIdentifier: big.NewInt(sToI64(params.Consideration[0].IdentifierOrCriteria)),

		Offerer: params.Offerer,

		OfferToken:      params.Offer[0].Token,
		OfferIdentifier: big.NewInt(sToI64(params.Offer[0].IdentifierOrCriteria)),
		OfferAmount:     big.NewInt(sToI64(params.Offer[0].StartAmount)),

		BasicOrderType: uint8(params.OrderType),

		OffererConduitKey:   params.ConduitKey,
		FulfillerConduitKey: params.ConduitKey,

		TotalOriginalAdditionalRecipients: big.NewInt(int64(params.TotalOriginalConsiderationItems - 1)),

		AdditionalRecipients: additionalRecipients,

		StartTime: big.NewInt(sToI64(params.StartTime)),
		EndTime:   big.NewInt(sToI64(params.EndTime)),

		Salt: salt,

		Zone:     params.Zone,
		ZoneHash: params.ZoneHash,

		Signature: signature,
	}
	gbl.Log.Debugf("orderParameters: %+v", orderParameters)

	//
	// create the tx

	// get next nonce
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		gbl.Log.Errorf("❌ failed to get nonce: %v", err)
	}

	// get suggested gas settings
	gasFeeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		gbl.Log.Errorf("❌ failed to get gas fee cap: %v", err)
		gbl.Log.Error(err)
	}

	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		gbl.Log.Errorf("❌ failed to get gas tip cap: %v", err)
	}

	// gas limit
	seaportABIFile, _ := os.Open("internal/abis/seaport11.json")
	seaport, _ := abi.JSON(seaportABIFile)

	orderData, err := seaport.Pack("fulfillBasicOrder", orderParameters)
	if err != nil {
		gbl.Log.Errorf("❌ failed to pack order data: %v", err)
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &OpenSeaSeaportContract,
		Data: orderData,
	})
	if err != nil {
		gbl.Log.Errorf("❌ failed to get gas limit: %v", err)

		gasLimit = uint64(160000)
	}

	//
	// create & set tx options
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	if err != nil {
		gbl.Log.Errorf("❌ NewKeyedTransactorWithChainID error: %v", err.Error())
	}

	txOpts.From = publicAddress
	txOpts.GasFeeCap = gasFeeCap
	txOpts.GasTipCap = gasTipCap
	txOpts.GasLimit = gasLimit                            // in units
	txOpts.Value = big.NewInt(sToI64(order.CurrentPrice)) // in wei
	txOpts.Nonce = big.NewInt(int64(nonce))

	// disable sending the tx to a node (for testing/debugging)
	txOpts.NoSend = false

	gbl.Log.Debugf("txOpts: %+v", txOpts)
	gbl.Log.Debugf("orderParameters: %+v", orderParameters)

	// send the tx
	tx, err := abiSeaport.FulfillBasicOrder(txOpts, orderParameters)
	if err != nil {
		gbl.Log.Errorf("❌ FulfillBasicOrder error: %v", err.Error())
		return nil, err
	}

	gbl.Log.Infof("tx %s: %+v", tx.Hash().String(), tx)

	if data := tx.Data(); len(data) > 0 {
		gbl.Log.Debugf("tx.Data() (%d):\n%x", len(data), data)
		gbl.Log.Debugf("orderData (%d):\n%x", len(orderData), orderData)
		dump.P(data)
	}

	// wait for tx to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		gbl.Log.Errorf("❌ WaitMined error: %v", err.Error())
	}

	gbl.Log.Debugf("receipt: %+v", receipt)

	return tx, nil
}

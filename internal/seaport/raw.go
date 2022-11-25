package seaport

import (
	"context"
	"fmt"
	"math/big"

	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gookit/goutil/dump"
)

func SendRawTx(ethNodes *nodes.Nodes, to common.Address, value float64, data []byte, privKey string) (*types.Transaction, error) {
	// get key and address
	privateKey, publicAddress, err := getKeyAndAddressFromString(privKey)
	if err != nil {
		gbl.Log.Error("❌ error while getting key and address from string: ", err)
		return nil, err
	}

	// eth client
	client := ethNodes.GetRandomLocalNode().Client

	//
	// create the tx

	// retrieve the chainid (needed for signer)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		gbl.Log.Errorf("❌ failed to get chainID: %v", err)
	}

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

	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &OpenSeaSeaportContract,
	// 	Data: orderData,
	// })
	// if err != nil {
	// 	gbl.Log.Errorf("❌ failed to get gas limit: %v", err)

	// 	gasLimit = uint64(160000)
	// }
	gasLimit := uint64(200000)

	valueEther := big.NewFloat(value)
	valueWei, _ := nodes.EtherToWei(valueEther).Int(big.NewInt(0))

	fmt.Printf("value: %+v | valueEther: %+v | valueWei: %+v\n", value, valueEther, valueWei)

	// Create a new transaction
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     valueWei,
			Data:      nil,
		})

	gbl.Log.Infof("tx %s: %+v", tx.Hash().String(), tx)
	fmt.Printf("tx %s: %+v\n", tx.Hash().String(), tx)

	// Sign the transaction using our keys
	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)

	// send the tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		gbl.Log.Warnf("❌ SendTransaction error: %v", err.Error())
		return nil, err
	}

	if data := tx.Data(); len(data) > 0 {
		gbl.Log.Debugf("tx.Data() (%d):\n%x", len(data), data)
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

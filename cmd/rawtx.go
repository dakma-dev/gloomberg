package cmd

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"sync"

// 	"github.com/benleb/gloomberg/internal/config"
// 	"github.com/benleb/gloomberg/internal/nodes"
// 	"github.com/benleb/gloomberg/internal/utils"
// 	"github.com/benleb/gloomberg/internal/utils/gbl"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/gookit/goutil/dump"
// 	"github.com/spf13/cobra"
// )

// // rawtxCmd represents the rawtx command.
// var rawtxCmd = &cobra.Command{
// 	Use:   "rawtx",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	// Run: func(cmd *cobra.Command, args []string) {
// 	Run: rawtx,
// }

// func init() { rootCmd.AddCommand(rawtxCmd) }

// func rawtx(_ *cobra.Command, _ []string) {
// 	fmt.Println("rawtx called")

// 	gbl.GetSugaredLogger()

// 	to := common.HexToAddress("0x522640f5413325d652237737a6bd8b8ee11fe561")

// 	// ethNodes := config.GetNodesFromConfig()
// 	// if ethNodes.ConnectAllNodes() == nil {
// 	// 	fmt.Println("no ethNodes")

// 	// 	return
// 	// }

// 	var ethNodes *nodes.Nodes
// 	// read ethNodes from config & establish connections to the ethNodes
// 	if ethNodes := config.GetNodesFromConfig(); ethNodes.ConnectAllNodes() != nil {
// 	}

// 	client := ethNodes.GetRandomLocalNode().Client

// 	wMap := map[string]map[string]interface{}{
// 		"clevil": {
// 			"private_key": "2053232cc9e6f7dc4fb3f26d7f7ab8cd3a2286d20457a2a65f059fbd66261f3c",
// 			"nonce":       8170539426785050,
// 			"sig":         "481d7cd0a9fe6edfff5d942765229a761d9687326fd00e7379a516b9cda83c9d422caa96a3e507eb2f82c187ec08ff8958fc08fcc4606d7141613af1885d65641c",
// 		},
// 		"dakma": {
// 			"private_key": "26310fcb15bdf48c70f502cb5652ca364159feba74717669276c002ee3fa5c91",
// 			"nonce":       8351503829570598,
// 			"sig":         "3ecfd9f08d732d38018cbb3c059fe3a9b55d7e98821c13df683e7319b197655062676d4e06682ad8567fb67b542b44af0f95f0af43eb6260ade602d362cccdc31b",
// 		},
// 	}

// 	var wgBuy sync.WaitGroup

// 	for _, walletData := range wMap {
// 		wgBuy.Add(1)

// 		go func(walletData map[string]interface{}) {
// 			defer wgBuy.Done()

// 			privKey := walletData["private_key"].(string)
// 			if privKey == "" {
// 				gbl.Log.Error("❌ private key is empty")

// 				return
// 			}

// 			madBinding, err := mad.NewMad(to, client)
// 			if err != nil {
// 				gbl.Log.Error("❌ error sending raw tx: %s", err)
// 			}

// 			// get key and address
// 			privateKey, publicAddress, err := getKeyAndAddressFromString(privKey)
// 			if err != nil || privateKey == nil {
// 				gbl.Log.Error("❌ error while getting key and address from string: ", err)

// 				return
// 			}

// 			// // retrieve the chainid (needed for signer)
// 			// chainID, err := client.ChainID(context.Background())
// 			// if err != nil {
// 			// 	gbl.Log.Errorf("❌ failed to get chainID: %v", err)
// 			// }

// 			//
// 			// create & set tx options
// 			txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
// 			if err != nil {
// 				gbl.Log.Errorf("❌ NewKeyedTransactorWithChainID error: %v", err.Error())
// 			}

// 			// get next nonce
// 			nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
// 			if err != nil {
// 				gbl.Log.Errorf("❌ failed to get nonce: %v", err)
// 			}

// 			// get suggested gas settings
// 			gasFeeCap, err := client.SuggestGasPrice(context.Background())
// 			if err != nil {
// 				gbl.Log.Errorf("❌ failed to get gas fee cap: %v", err)
// 				gbl.Log.Error(err)
// 			}

// 			gasTipCap, err := client.SuggestGasTipCap(context.Background())
// 			if err != nil {
// 				gbl.Log.Errorf("❌ failed to get gas tip cap: %v", err)
// 			}

// 			// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 			// 	To:   &OpenSeaSeaportContract,
// 			// 	Data: orderData,
// 			// })
// 			// if err != nil {
// 			// 	gbl.Log.Errorf("❌ failed to get gas limit: %v", err)

// 			// 	gasLimit = uint64(160000)
// 			// }
// 			gasLimit := uint64(200000)

// 			txOpts.From = publicAddress
// 			txOpts.GasFeeCap = big.NewInt(0).Mul(gasFeeCap, big.NewInt(2))
// 			txOpts.GasTipCap = big.NewInt(0).Mul(gasTipCap, big.NewInt(2))
// 			txOpts.GasLimit = gasLimit   // in units
// 			txOpts.Value = big.NewInt(0) // in wei
// 			txOpts.Nonce = big.NewInt(int64(nonce))

// 			// disable sending the tx to a node (for testing/debugging)
// 			txOpts.NoSend = false

// 			gbl.Log.Infof("txOpts: %+v", txOpts)
// 			// gbl.Log.Debugf("orderParameters: %+v", orderParameters)

// 			data, err := hex.DecodeString(walletData["sig"].(string))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Printf("data: %x\n", data)

// 			tx, err := madBinding.MintPresale(txOpts, big.NewInt(2), big.NewInt(2), big.NewInt(int64(walletData["nonce"].(int))), data)
// 			if err != nil {
// 				gbl.Log.Errorf("❌ MintPresale error: %v", err.Error())
// 			}

// 			fmt.Printf("tx: %+v\n", tx)
// 			fmt.Printf("privateKey: %+v\n", privateKey)

// 			// Sign the transaction using our keys
// 			signedTx, _ := types.SignTx(tx, types.NewLondonSigner(big.NewInt(1)), privateKey)

// 			// send the tx
// 			err = client.SendTransaction(context.Background(), signedTx)
// 			if err != nil {
// 				gbl.Log.Warnf("❌ SendTransaction error: %v", err.Error())

// 				return
// 			}

// 			if data := tx.Data(); len(data) > 0 {
// 				gbl.Log.Debugf("tx.Data() (%d):\n%x", len(data), data)
// 				dump.P(data)
// 			}

// 			// wait for tx to be mined
// 			receipt, err := bind.WaitMined(context.Background(), client, tx)
// 			if err != nil {
// 				gbl.Log.Errorf("❌ WaitMined error: %v", err.Error())
// 			}

// 			gbl.Log.Debugf("receipt: %+v", receipt)

// 			// tx, err := seaport.SendRawTx(ethNodes, to, 0.003, privateKey)
// 			// if err != nil {
// 			// 	gbl.Log.Error("❌ error sending raw tx: %s", err)
// 			// }

// 			gbl.Log.Info("✅ tx sent: https://etherscan.io/tx/%s", tx.Hash().Hex())
// 		}(walletData)

// 		wgBuy.Wait()
// 	}
// }

// func getKeyAndAddressFromString(key string) (*ecdsa.PrivateKey, common.Address, error) {
// 	// parse private key
// 	privateKey, err := crypto.HexToECDSA(key)
// 	if err != nil {
// 		gbl.Log.Error("❌ error while parsing private key: ", err)

// 		return nil, utils.ZeroAddress, err
// 	}

// 	// derive public key
// 	publicKey := privateKey.Public()

// 	// get ecdsa public key
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		gbl.Log.Error("❌ error casting public key to ECDSA")

// 		return nil, utils.ZeroAddress, err
// 	}

// 	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	return privateKey, publicAddress, nil
// }

// func sendRawTx(ethNodes *nodes.Nodes, to common.Address, value float64, privKey string) (*types.Transaction, error) {
// 	// get key and address
// 	privateKey, publicAddress, err := getKeyAndAddressFromString(privKey)
// 	if err != nil {
// 		gbl.Log.Error("❌ error while getting key and address from string: ", err)

// 		return nil, err
// 	}

// 	// eth client
// 	client := ethNodes.GetRandomLocalNode().Client

// 	//
// 	// create the tx

// 	// retrieve the chainid (needed for signer)
// 	chainID, err := client.ChainID(context.Background())
// 	if err != nil {
// 		gbl.Log.Errorf("❌ failed to get chainID: %v", err)
// 	}

// 	// get next nonce
// 	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
// 	if err != nil {
// 		gbl.Log.Errorf("❌ failed to get nonce: %v", err)
// 	}

// 	// get suggested gas settings
// 	gasFeeCap, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		gbl.Log.Errorf("❌ failed to get gas fee cap: %v", err)
// 		gbl.Log.Error(err)
// 	}

// 	gasTipCap, err := client.SuggestGasTipCap(context.Background())
// 	if err != nil {
// 		gbl.Log.Errorf("❌ failed to get gas tip cap: %v", err)
// 	}

// 	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 	// 	To:   &OpenSeaSeaportContract,
// 	// 	Data: orderData,
// 	// })
// 	// if err != nil {
// 	// 	gbl.Log.Errorf("❌ failed to get gas limit: %v", err)

// 	// 	gasLimit = uint64(160000)
// 	// }
// 	gasLimit := uint64(200000)

// 	valueEther := big.NewFloat(value)
// 	valueWei, _ := nodes.EtherToWei(valueEther).Int(big.NewInt(0))

// 	fmt.Printf("value: %+v | valueEther: %+v | valueWei: %+v\n", value, valueEther, valueWei)

// 	// Create a new transaction
// 	tx := types.NewTx(
// 		&types.DynamicFeeTx{
// 			ChainID:   chainID,
// 			Nonce:     nonce,
// 			GasTipCap: gasTipCap,
// 			GasFeeCap: gasFeeCap,
// 			Gas:       gasLimit,
// 			To:        &to,
// 			Value:     valueWei,
// 			Data:      nil,
// 		})

// 	gbl.Log.Infof("tx %s: %+v", tx.Hash().String(), tx)
// 	fmt.Printf("tx %s: %+v\n", tx.Hash().String(), tx)

// 	// Sign the transaction using our keys
// 	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)

// 	// send the tx
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		gbl.Log.Warnf("❌ SendTransaction error: %v", err.Error())

// 		return nil, err
// 	}

// 	if data := tx.Data(); len(data) > 0 {
// 		gbl.Log.Debugf("tx.Data() (%d):\n%x", len(data), data)
// 		dump.P(data)
// 	}

// 	// wait for tx to be mined
// 	receipt, err := bind.WaitMined(context.Background(), client, tx)
// 	if err != nil {
// 		gbl.Log.Errorf("❌ WaitMined error: %v", err.Error())
// 	}

// 	gbl.Log.Debugf("receipt: %+v", receipt)

// 	return tx, nil
// }

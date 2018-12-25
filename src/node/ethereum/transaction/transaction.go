package transaction

import (
	"context"
	"crypto/ecdsa"
	"github.com/BANKEX/plasma-research/src/node/ethereum/etherUtils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func SendTransactionInWei(client *ethclient.Client, privateKey string, value int64, toAddress string, ) (string, error) {
	rawPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := rawPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	err = etherUtils.IsValidAddress(fromAddress)
	if err != nil {
		return "", err
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	err = etherUtils.IsValidAddress(toAddress)
	if err != nil {
		return "", err
	}

	rawToAddress := common.HexToAddress(toAddress)
	rawValue := big.NewInt(value)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	var data []byte
	tx := types.NewTransaction(nonce, rawToAddress, rawValue, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), rawPrivateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), err
}


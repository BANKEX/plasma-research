package deposit

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/BANKEX/plasma-research/src/node/ethereum/etherUtils"
	"github.com/BANKEX/plasma-research/src/node/ethereum/plasmacontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Deposit(client *ethclient.Client, privateKey string, contractAddress string, value int64) (string, error) {
	rawPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", errors.New("error casting string private key to raw private key")
		log.Println(err)
	}

	publicKey := rawPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
		log.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	err = etherUtils.IsValidAddress(fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(rawPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	address := common.HexToAddress(contractAddress)
	err = etherUtils.IsValidAddress(address)
	if err != nil {
		return "", err
	}

	instance, err := store.NewStore(address, client)
	if err != nil {
		return "", err
	}

	tx, err := instance.Deposit(auth)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

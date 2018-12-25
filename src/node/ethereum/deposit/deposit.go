package deposit

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/BANKEX/plasma-research/src/node/ethereum/plasmacontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func Deposit(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress common.Address, value int64) (string, error) {
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		return "", err
	}

	tx, err := instance.Deposit(auth)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

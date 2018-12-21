package transaction

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"regexp"
)

func SendTransactionInWei(client *ethclient.Client, value int, toAddress string, privateKey string) (string, error) {
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
	err = IsValidAddress(fromAddress)
	if err != nil {
		return "", err
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	err = IsValidAddress(toAddress)
	if err != nil {
		err := errors.New("failed to validate address")
		return "", err
	}

	rawToAddress := common.HexToAddress(toAddress)

	rawValue := big.NewInt(int64(value))

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

func IsValidAddress(iaddress interface{}) (error) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	ok := false
	switch v := iaddress.(type) {
	case string:
		ok = re.MatchString(v)
		if !ok {
			err := errors.New("failed to validate address")
			return err
		}
		return nil
	case common.Address:
		ok = re.MatchString(v.Hex())
		if !ok {
			err := errors.New("failed to validate address")
			return err
		}
		return nil
	default:
		err := errors.New("failed to validate address")
		return err
	}
}
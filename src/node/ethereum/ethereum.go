package ethereum

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strconv"

	"github.com/BANKEX/plasma-research/src/node/config"
	"github.com/BANKEX/plasma-research/src/node/ethereum/plasmacontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(address string) string {
	client, err := ethclient.Dial(config.GetVerifier().GethHost)
	if err != nil {
		log.Println(err)
	}

	account := common.HexToAddress(address)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Println(err)
	}
	return pendingBalance.String()
}

func PushHashBlock(blockNumber uint32, hash []byte) {
	client, err := ethclient.Dial(config.GetVerifier().GethHost)
	if err != nil {
		log.Println(err)
	}

	privateKey, err := crypto.HexToECDSA(config.GetVerifier().VerifierPrivateKey)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(config.GetVerifier().PlasmaContractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Println(err)
	}

	// _, err = instance.SubmitBlocks(blockNumber, hash) // TODO: uncomment after regenerating abi
	_, err = instance.SubmitBlocks(auth, nil, nil) // TODO: normal params
	if err != nil {
		log.Println(err)
	}
}

func GetLastBlockNumber() string {
	client, err := ethclient.Dial(config.GetVerifier().GethHost)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(config.GetVerifier().PlasmaContractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	blockLength, err := instance.BlocksLength(nil)
	if err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa((int)(blockLength.Uint64() - 1))
}

func Exit() {
	client, err := ethclient.Dial(config.GetVerifier().GethHost)
	if err != nil {
		log.Println(err)
	}

	privateKey, err := crypto.HexToECDSA(config.GetVerifier().VerifierPrivateKey)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(config.GetVerifier().PlasmaContractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Println(err)
	}

	_, err = instance.WithdrawalBegin(auth, nil)
	if err != nil {
		log.Println(err)
	}

}

package main

import (
	"./plasmacontract"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)
type Config struct {
	Port     int    `json:port`
	Key      string `json:key`
	Operator string `json:operator`
	Node     string `json:node`
	Smart    string `json:smart`
}
func OpenConfig(file string) Config {
	// Open  json config File
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var config Config

	// write byte-info in conf
	json.Unmarshal(byteValue, &config)

	return config
}


var Nonce = 0
func main () {

	fmt.Println("THIS IS TEST TX FOR EVENT")

	config := flag.String("c", "config.json", "config file for verifier")
	flag.Parse()
	conf := OpenConfig(*config)


	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	for i:=0; i < 100000; i++ {
		do(client, conf)
		Nonce++
		//sendTx(client, conf)
		//Nonce++
		//time.Sleep(time.Millisecond)
	}
}

func sendTx(client *ethclient.Client, conf Config) {
	privateKey, err := crypto.HexToECDSA("240D6AD83930067D82E0803696996F743ACD78D8FA6A5F6E4F148FD9DEF37C55")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	_, _ = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(100000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(conf.Smart)
	var data []byte
	tx := types.NewTransaction(uint64(Nonce), toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}


}

func do(client *ethclient.Client, conf Config) {


	privateKey, err := crypto.HexToECDSA("240D6AD83930067D82E0803696996F743ACD78D8FA6A5F6E4F148FD9DEF37C55")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	_, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(Nonce))
	auth.Value = big.NewInt(10000000)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice


	address := common.HexToAddress(conf.Smart)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}


	_, _ = instance.Deposit(auth)


	//fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
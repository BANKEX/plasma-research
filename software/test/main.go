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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
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



func main () {

	fmt.Println("THIS IS TEST TX FOR EVENT")

	config := flag.String("c", "config.json", "config file for verifier")
	flag.Parse()
	conf := OpenConfig(*config)


	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	for i:=0; i < 1000; i++ {
		do(client, conf)
		time.Sleep(time.Second*2)
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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
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
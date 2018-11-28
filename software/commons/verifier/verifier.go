package main

import (
	"../config"
	"../dispatchers"
	"../listeners"
	"../listeners/balance"
	"../listeners/ethClient"
	"../listeners/event"
	"../listeners/storage"
	"./handlers"
	"./portscanner"
	"./server"
	"./cli"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

func main() {

	defaultConfigPath, _ := filepath.Abs("../config/config.verifier.json")

	configFileName := flag.String("c", defaultConfigPath, "config file for verifier")
	flag.Parse()

	_, conf, err := config.ReadConfig(*configFileName, "v")
	if err != nil {
		log.Fatal(err)
	}
	if conf.Verifier_port == 0 {
		fmt.Println("Unmarshalling error!!!")
		return
	}

	fmt.Println("\n\n")
	fmt.Println("PORT: " + strconv.Itoa(conf.Verifier_port))
	fmt.Println("KEY: " + conf.Main_account_private_key)
	fmt.Println("Smart Contract address: " + conf.Plasma_operator_address)
	//fmt.Println("Operator IP: " + conf.Plasma_operator_address)
	fmt.Println("Node: " + conf.Geth_account)
	fmt.Println("\n\n")

	ethClient.InitClient(conf.Geth_account)
	dispatchers.CreateGenesisBlock()
	handlers.OperatorAddress = conf.Plasma_operator_address

	go listeners.Checker()
	go balance.UpdateBalance(&storage.Balance, conf.Plasma_operator_address)
	go event.Start(storage.Client, conf.Plasma_operator_address, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)
	go server.GinServer(conf)
	go portscanner.RunScanner()
	cli.CLI()

}

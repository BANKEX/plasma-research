package main

import (
	"../config"
	"../dispatchers"
	"../listeners"
	"../listeners/balance"
	"../listeners/ethClient"
	"../listeners/event"
	"../listeners/storage"
	"./cli"
	"./portscanner"
	"./server"
)

func main() {

	conf := config.InitConfig()

	ethClient.InitClient(conf.Geth_host)
	dispatchers.CreateGenesisBlock()

	go listeners.Checker()
	go balance.UpdateBalance(&storage.Balance, conf.Plasma_contract_address)
	go event.Start(storage.Client, conf.Plasma_contract_address, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)
	go server.GinServer(conf)
	go portscanner.RunScanner()

	println("Verifier started")

	cli.CLI()

}

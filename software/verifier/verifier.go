package main

import (
	"../commons/config"
	"../commons/dispatchers"
	"../commons/listeners"
	"../commons/listeners/balance"
	"../commons/listeners/ethClient"
	"../commons/listeners/event"
	"../commons/listeners/storage"
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

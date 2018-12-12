package main

import (
	"github.com/BANKEX/plasma-research/src/node/verifier/server"
)

func main() {

	// go listeners.Checker()
	// go balance.UpdateBalance(&storage.Balance, conf.Plasma_contract_address)
	// go event.Start(storage.Client, conf.Plasma_contract_address, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)
	go server.GinServer()

	println("Verifier started")

	//cli.CLI()

}

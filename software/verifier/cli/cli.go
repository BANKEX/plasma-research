package cli

import (
	"../../commons/config"
	"../../commons/db"
	"../../commons/listeners/storage"
	"fmt"
	"github.com/c-bata/go-prompt"
)

func commandsInfo(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "smartContractAddress", Description: "Get Smart Contract address"},
		{Text: "plasmaBalance", Description: "Get balance of my account in Plasma"},
		{Text: "smartContractBalance", Description: "Get balance of Plasma smart contract"},
		{Text: "events", Description: "Get all events"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func commandsListener(input string) {
	switch input {
	case "smartContractAddress":
		fmt.Println(config.SmartContractAddress)
	case "plasmaBalance":
		fmt.Println("Not working yet")
	case "smartContractBalance":
		fmt.Println("Smart Contract Balance:" + storage.Balance)
	case "events":
		events, err := db.Event("database").GetAll()
		if err != nil {
			println("Mistake at cli.go 34 line with DB")
		}
		fmt.Println(events)
	}
	return
}

func CLI() {
	prompter := prompt.New(
		commandsListener,
		commandsInfo,
	)
	prompter.Run()
}

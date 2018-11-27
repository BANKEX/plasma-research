package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	 "../../db"
	"../handlers"
	"../../listeners/storage"
	"../../listeners/event"


)

// For CLI
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "plasmaBalance", Description: "Get plasma balance"},
		{Text: "smartBalance", Description: "Get Smart Contract balance"},
		{Text: "eventMap", Description: "Get all events map"},
		{Text: "dbEvents", Description: "Get all events from dbEvents"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(comm string) {
	if comm == "plasmaBalance" {
		plasmaBalance := handlers.GetReqBalance(handlers.OperatorAddress, "/pbalance")
		fmt.Println("Plasma balance:" + plasmaBalance)
	} else if comm == "smartBalance" {
		fmt.Println("Smart Contract Balance:" + storage.Balance)
	} else if comm == "eventMap" {
		for i, j := range event.EventMap {
			fmt.Println(i, j)
		}
	} else if comm == "dbEvents" {
		events, err := db.Event("database").GetAll()
		if err != nil {
			println("Mistake DB")
		}
		fmt.Println(events)
	}

	return
}

func CLI() {
	fmt.Println("-------------Plasma Verifier-------------")
	p := prompt.New(
		executor,
		completer,
	)
	p.Run()
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"../db"
	"../dispatchers"
	"../listeners"
	"../listeners/balance"
	"../listeners/ethClient"
	"../listeners/event"
	"../listeners/storage"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"../config"
	"./handlers"
	"github.com/c-bata/go-prompt"
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
		// fmt.Println(fmt.Println(event.EventMap))
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

func GinServer(conf config.VerifierConfig) {

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	r.GET("/scgetbalance", handlers.SCGetBalance)
	r.GET("/plasmabalance", handlers.GetMyPlasmaBalance)

	r.GET("/conf", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"main_account_public_key": conf.Main_account_public_key,
			"plasma_operator_address": conf.Plasma_operator_address,
			"geth_account":            conf.Geth_account,
		})
	})

	r.Run(":" + strconv.Itoa(conf.Verifier_port))
}

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
	fmt.Println("Smart Contract address: " + conf.Main_account_public_key)
	fmt.Println("Operator IP: " + conf.Plasma_operator_address)
	fmt.Println("Node: " + conf.Geth_account)
	fmt.Println("\n\n")

	ethClient.InitClient(conf.Geth_account)

	dispatchers.CreateGenesisBlock()

	go listeners.Checker()
	go balance.UpdateBalance(&storage.Balance, conf.Main_account_public_key)
	go event.Start(storage.Client, conf.Main_account_public_key, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)

	handlers.OperatorAddress = conf.Plasma_operator_address

	// Uncomment for start CLI
	// CLI()

	// Uncomment for start ginServer
	GinServer(conf)

}

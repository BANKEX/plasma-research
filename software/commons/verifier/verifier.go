package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	"./handlers"
	"github.com/c-bata/go-prompt"
)

type Config struct {
	Verifier_port            int    `json:verifier_port`
	Main_account_private_key string `json:main_account_private_key`
	Plasma_operator_address  string `json:plasma_operator_address`
	Geth_account             string `json:geth_account`
	Main_account_public_key  string `json:main_account_public_key`
}

func ReadConfig(fileName string) (Config, error) {

	var config Config

	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(config)

	return config, nil
}

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

func GinServer(conf Config) {
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

	defaultConfigPath, _ := filepath.Abs("../config.json")

	configFileName := flag.String("c", defaultConfigPath, "config file for verifier")
	flag.Parse()

	conf, err := ReadConfig(*configFileName)
	if err != nil {
		log.Fatal(err)
	}
	if conf.Verifier_port == 0 {
		fmt.Println("Can't read config.json!!!")
		return
	}

	fmt.Println("\n\n")
	fmt.Println("PORT: " + strconv.Itoa(conf.Verifier_port))
	fmt.Println("KEY: " + conf.Main_account_private_key)
	fmt.Println("Operator IP: " + conf.Plasma_operator_address)
	fmt.Println("Node: " + conf.Geth_account)
	fmt.Println("Smart Contract address: " + conf.Main_account_public_key)
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

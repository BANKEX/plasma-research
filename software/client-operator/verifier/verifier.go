package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	Port     int    `json:port`
	Key      string `json:key`
	Operator string `json:operator`
	EthNode  string `json:ethNode`
	Smart    string `json:smart`
}

func ReadConfig(fileName string) (Config, error) {
	var config Config

	f, err := os.Open(fileName)
	if err != nil {
		return config, err
	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return config, err
	}

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
			"Smart":    conf.Smart,
			"Operator": conf.Operator,
			"Node":     conf.Node,
		})
	})

	r.Run(":" + strconv.Itoa(conf.Port))
}

func main() {

	configFileName := flag.String("c", "config.json", "config file for verifier")
	flag.Parse()

	conf, err := ReadConfig(*configFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n")
	fmt.Println("PORT: " + strconv.Itoa(conf.Port))
	fmt.Println("KEY: " + conf.Key)
	fmt.Println("Operator IP: " + conf.Operator)
	fmt.Println("Node: " + conf.EthNode)
	fmt.Println("Smart Contract address: " + conf.Smart)
	fmt.Println("\n\n")

	ethClient.InitClient(conf.EthNode)

	dispatchers.CreateGenesisBlock()

	go listeners.Checker()
	go balance.UpdateBalance(&storage.Balance, conf.Smart)
	go event.Start(storage.Client, conf.Smart, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)

	handlers.OperatorAddress = conf.Operator

	// Uncomment for start CLI
	CLI()

	// Uncomment for start ginServer
	//GinServer(conf)
}

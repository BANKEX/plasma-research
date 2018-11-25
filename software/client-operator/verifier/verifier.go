package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"../listeners/balance"
	"../listeners/ethClient"
	"../listeners/event"
	"../listeners/storage"

	"./handlers"

	"github.com/c-bata/go-prompt"
)

type Config struct {
	Port     int    `json:port`
	Key      string `json:key`
	Operator string `json:operator`
	Node     string `json:node`
	Smart    string `json:smart`
}

var info = ""

// For open config file
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

func Checker() {
	for {
		if storage.StateForEvent == 1 {
			fmt.Println("\n\n\n")

			event.EventCount++

			event.EventMap[event.EventCount] =
					"Who" + "\n" + storage.Who.String() + "\n" +
					"Amount" + "\n" + storage.Amount.String() + "\n" +
					"BlockHash" + "\n" + storage.EventBlockHash + "\n" +
					"BlockNumber" + "\n" + strconv.Itoa(int(storage.EventBlockNumber))

			storage.StateForEvent = 0

			fmt.Println(event.EventMap[event.EventCount])
		}
		time.Sleep(time.Second * 0)
	}
}

// For CLI
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "plasmaBalance", Description: "Get plasma balance"},
		{Text: "smartBalance", Description: "Get Smart Contract balance"},
		{Text: "eventMap", Description: "Get all events map"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(comm string) {
	if comm == "plasmaBalance" {
		balance := handlers.GetReqBalance(handlers.OperatorAddress, "/pbalance")
		fmt.Println("Plasma balance:" + balance)
	} else if comm == "smartBalance" {
		fmt.Println("Smart Contract Balance:" + storage.Balance)
	} else if comm == "eventMap" {
		// fmt.Println(fmt.Println(event.EventMap))
		for i, j := range event.EventMap {
			fmt.Println(i, j)
		}
	}
	return
}

func CLI() {
	fmt.Println("------------Plasma Verifier----------")

	p := prompt.New(
		executor,
		completer,
	)
	p.Run()
}

func main() {

	config := flag.String("c", "config.json", "config file for verifier")

	flag.Parse()

	conf := OpenConfig(*config)

	fmt.Println("\n\n")

	fmt.Println("PORT: " + strconv.Itoa(conf.Port))
	fmt.Println("KEY: " + conf.Key)
	fmt.Println("Operator IP: " + conf.Operator)
	fmt.Println("Node: " + conf.Node)
	fmt.Println("Smart Contract address: " + conf.Smart)

	fmt.Println("\n\n")

	ethClient.InitClient(conf.Node)
	go Checker()
	go balance.UpdateBalance(&storage.Balance, conf.Smart)
	go event.Start(storage.Client, conf.Smart, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)

	handlers.OperatorAddress = conf.Operator

	//CLI()

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

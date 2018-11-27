package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"../config"
	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	defaultConfigPath, _ := filepath.Abs("../config/config.operator.json")

	configFileName := flag.String("c", defaultConfigPath, "config file for verifier")
	flag.Parse()
	conf, _, err := config.ReadConfig(*configFileName, "o")
	if err != nil {
		log.Fatal(err)
	}
	if conf.Operator_port == 0 {
		fmt.Println("Can't read config.json!!!")
		return
	}
	fmt.Println("\n\n")
	fmt.Println("PORT: " + strconv.Itoa(conf.Operator_port))
	fmt.Println("KEY: " + conf.Main_account_private_key)
	fmt.Println("Smart Contract address: " + conf.Main_account_public_key)
	fmt.Println("Node: " + conf.Geth_account)
	fmt.Println("\n\n")



	r := gin.Default()

	r.POST("/settx/:tx", handlers.SetTx)
	r.GET("/gettx/:tx", handlers.GetTx)
	r.GET("/getall", handlers.GetAllTx)
	r.GET("/publishblock", handlers.PublishBlock)
	r.GET("/pbalance", handlers.PBalance)
	r.Run(":" + strconv.Itoa(conf.Operator_port))

}

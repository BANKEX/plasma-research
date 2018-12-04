package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"../commons/config"
	"./handlers"
	"github.com/gin-gonic/gin"
)

func assembleBlocks(d time.Duration) {
	for range time.Tick(d) {
		assembleBlock()
	}
}

func main() {

	defaultConfigPath, _ := filepath.Abs("../commons/config/config.operator.json")

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

	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	// Assemble block ~ each second
	go assembleBlocks(time.Second * 1)

	// Handler that accept new transaction from verifier
	r.POST("/settx/:tx", handlers.SetTx)
	r.GET("/gettx/:tx", handlers.GetTx)
	r.GET("/getall", handlers.GetAllTx)
	// Handler for debug
	r.GET("/publishblock", handlers.PublishBlock)
	r.GET("/pbalance", handlers.PBalance)
	r.Run(":" + strconv.Itoa(conf.Operator_port))

	println("Operator started")

}

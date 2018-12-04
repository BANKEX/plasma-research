package main

import (
	b "../commons/blockchain"
	"../commons/config"
	"./handlers"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

// Pool of pending transactions
var transactionsPool = make(b.TransactionsPool)
var utxoPool = make(b.UtxoPool)

func assembleBlocks(d time.Duration, privateKey string) {

	privateKeyBytes, _ := hexutil.Decode(privateKey)

	for range time.Tick(d) {
		if transactionsPool.IsEmpty() {
			fmt.Print(".")
			continue
		}

		utxoPoolCopy := utxoPool.GetCopy()
		pendingTransactions := transactionsPool.GetTransactions()
		block, newUtxoPool := b.AssembleBlock(*utxoPoolCopy, pendingTransactions, privateKeyBytes)

		// TODO: atomic update of utxoPool and pending transactions
		{
			for _, t := range block.Transactions {
				transactionsPool.Remove(t)
			}
			utxoPool = newUtxoPool
		}
	}
}

func AddTxToThePool() gin.HandlerFunc {
	return func(c *gin.Context) {
		handlers.SetTx(transactionsPool, c)
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
	go assembleBlocks(time.Second*1, conf.Main_account_private_key)

	// Handler that accept new transaction from verifier
	r.POST("/settx/:tx", AddTxToThePool())
	r.GET("/gettx/:tx", handlers.GetTx)
	r.GET("/getall", handlers.GetAllTx)

	// Handler for debug
	r.GET("/publishblock", handlers.PublishBlock)
	r.GET("/pbalance", handlers.PBalance)
	r.Run(":" + strconv.Itoa(conf.Operator_port))

	println("Operator started")

}

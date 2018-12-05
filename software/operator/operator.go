package main

import (
	a "../commons/alias"
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

// Intermediate Index that maps outputs to transactions
// Map string that looks like "BlockNumber:TransactionNumber:OutputNumber" to transaction bytes
var txIndex = make(b.TxIndex)

// Pool of pending transactions
var transactionsPool = make(b.TransactionsPool)

// Map UTXO to transactions
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
		block, newUtxoPool := b.AssembleBlock(*utxoPoolCopy, pendingTransactions, txIndex, privateKeyBytes)

		// TODO: atomic update of utxoPool, pending transactions, and TxIndex
		{
			for transactionNumber, t := range block.Transactions {
				// Remove from pending transactions
				transactionsPool.Remove(t)

				// Update Tx Index (block:transaction:output to txHash)
				for outputNumber, _ := range t.UnsignedTransaction.Outputs {
					key := fmt.Sprintf("%d:%d:%d", block.BlockNumber, transactionNumber, outputNumber)
					txIndex[key] = a.ToTxHashBytes(t.GetHash())
				}
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

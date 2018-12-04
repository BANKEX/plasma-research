package handlers

import (
	"../../commons/blockchain"
	"../../commons/db"
	"../../commons/ether"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// for test only
func PBalance(c *gin.Context) {
	balance := blockchain.Balance["balance"]
	c.JSON(http.StatusOK, gin.H{
		"balance": strconv.Itoa(balance),
	})
}

func PublishBlock(c *gin.Context) {
	blockchain.Balance["balance"] += 1
	c.JSON(http.StatusOK, gin.H{
		"resp": "ok",
	})
}

func GetTx(c *gin.Context) {
	data, err := db.Tx("database").Get([]byte(c.Param("tx")))
	if err != nil {
		println("Mistake DB")
	}
	c.JSON(http.StatusOK, gin.H{
		"Tx": string(data),
	})
}

func SetTx(pool blockchain.TransactionsPool, c *gin.Context) {
	rawTransaction := []byte(c.Param("tx"))

	//TODO: Check that this unmarshaling actually works
	var t blockchain.Transaction

	err := json.Unmarshal(rawTransaction, &t)
	if err != nil {
		// TODO: use different log level for production and development
		log.Println(err)
		return
	}

	//TODO: Implement contend of validation function transaction
	//err = t.ValidateSlices()
	//if err != nil {
	//	log.Println("wrong ranges")
	//	return
	//}
	//
	err = t.ValidateSignatures()
	if err != nil {
		// TODO: use different log level for production and development
		log.Println("wrong signatures")
		return
	}

	txHash := ether.GetTxHash(rawTransaction)

	// 3) Put real one to pool
	pool.Add(t)

	err = db.Tx("database").Put(txHash, rawTransaction)
	fmt.Print(err)
	if err != nil {
		println("Mistake DB")
		c.JSON(http.StatusOK, gin.H{
			"Status": "Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": "Ok",
	})

}

func GetAllTx(c *gin.Context) {
	tx, err := db.Tx("database").GetAll()
	if err != nil {
		println("Mistake DB")
	}
	c.JSON(http.StatusOK, gin.H{
		"Txs": tx,
	})
}

package handlers

import (
	"../../commons/blockchain"
	"../../commons/db"
	"../../commons/ether"
	tp "../pool"
	"fmt"
	"github.com/gin-gonic/gin"
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

func SetTx(pool *tp.TransactionsPool, c *gin.Context) {
	rawTransaction := []byte(c.Param("tx"))
	txHash := ether.GetTxHash(rawTransaction)

	pool.Add(txHash)

	err := db.Tx("database").Put(txHash, rawTransaction)
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

package handlers

import (
	"../../blockchain"
	"../../plasmautils/slice"
	"../../transactionManager"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// todo refactor this
var (
	Manager *transactionManager.TransactionManager
)

func PostTransaction(c *gin.Context) {
	var t blockchain.Transaction
	err := c.BindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Manager.SubmitTransaction(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// returns a list of utxos for an address
func GetUtxos(c *gin.Context) {
	addr := c.Param("address")
	utxos, err := Manager.GetUtxosForAddress(addr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, utxos)
}

// returns last plasma block number etc.
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"lastBlock": Manager.GetLastBlockNumber(),
	})
}

// returns contract address and abi
func GetConfig(c *gin.Context) {
	// todo not implemented
	c.JSON(http.StatusOK, nil)
}

// ==== debug handlers =====

// returns a list of utxos for an address
func FundAddress(c *gin.Context) {
	addr, _ := hex.DecodeString(c.Param("address")[2:])
	out := blockchain.Output{
		Owner: addr,
		Slice: slice.Slice{Begin: 10, End: 20},
	}
	_, err := Manager.AssembleDepositBlock(out)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "OK")
}

func Transact(c *gin.Context) {
	addr, _ := hex.DecodeString(c.Param("address")[2:])
	key, _ := hex.DecodeString(c.Param("key")[2:])
	blockN, _ := strconv.Atoi(c.Param("block"))
	txN, _ := strconv.Atoi(c.Param("tx"))
	outN, _ := strconv.Atoi(c.Param("out"))

	in := Manager.GetUtxo(uint32(blockN), uint32(txN), uint32(outN))

	if in == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such utxo"})
		return
	}

	tx := blockchain.Transaction{
		UnsignedTransaction: blockchain.UnsignedTransaction{
			Inputs: []blockchain.Input{*in},
			Outputs: []blockchain.Output{
				{Owner: addr, Slice: in.Slice},
			},
		},
	}
	err := tx.Sign(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Manager.SubmitTransaction(&tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "OK")
}

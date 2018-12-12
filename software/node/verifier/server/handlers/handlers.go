package handlers

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BANKEX/plasma-research/software/node/blockchain"
	"github.com/BANKEX/plasma-research/software/node/config"
	"github.com/BANKEX/plasma-research/software/node/ethereum"
	"github.com/BANKEX/plasma-research/software/node/plasmautils/slice"
	"github.com/gin-gonic/gin"
)

var sumRes uint

func EthereumBalance(c *gin.Context) {
	response := ethereum.GetBalance(config.GetVerifier().VerifierPublicKey)
	c.JSON(http.StatusOK, gin.H{
		"balance": response,
	})
}

func PlasmaBalance(c *gin.Context) {

	st := make([]blockchain.Input, 0)

	resp, err := http.Get("http://localhost:3001/utxo/" + config.GetVerifier().VerifierPublicKey)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &st)
	if err != nil {
		log.Println(err)
	}

	// for _, tx := range st {
	// 	sumRes = (tx.Slice.End - tx.Slice.Begin) * blockchain.WeiPerCoin
	// }

	c.JSON(http.StatusOK, st)
}

func PlasmaContractAddress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"address": config.GetVerifier().PlasmaContractAddress,
	})
}

func DepositHandler(c *gin.Context) {
	result := ethereum.Deposit(c.Param("sum"))
	c.JSON(http.StatusOK, gin.H{
		"txHash": result,
	})
}

func TransferHandler(c *gin.Context) {
	address, _ := hex.DecodeString(c.Param("address")[2:])
	usum, _ := strconv.Atoi(c.Param("sum"))
	sum := uint32(usum)
	in := new(blockchain.Input)
	c.BindJSON(&in)

	uTx := blockchain.Transaction{
		UnsignedTransaction: blockchain.UnsignedTransaction{
			Inputs: []blockchain.Input{*in},
			Outputs: []blockchain.Output{
				{
					Owner: address,
					Slice: slice.Slice{
						Begin: in.Slice.Begin,
						End:   in.Slice.Begin + sum,
					},
				},
			},
		},
	}

	if in.Slice.End-in.Slice.Begin > sum {
		uTx.Outputs = append(uTx.Outputs, blockchain.Output{
			Owner: in.Owner,
			Slice: slice.Slice{
				Begin: in.Slice.Begin + sum + 1,
				End:   in.Slice.End,
			},
		})
	}

	key, _ := hex.DecodeString(config.GetVerifier().VerifierPrivateKey)

	fmt.Println("Tx")
	fmt.Println(len(uTx.Outputs[0].Owner))

	uTx.Sign(key)

	fmt.Println("Signature")
	fmt.Println(len(uTx.Signatures))

	port := strconv.Itoa(config.GetOperator().OperatorPort)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(uTx)
	res, _ := http.Post("http://localhost:"+port+"/tx", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)

	c.JSON(http.StatusOK, gin.H{
		"status": res.Body,
	})
}

func makeUnsignedTx() {

}

func ExitHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

var id = 0

type Resp struct {
	LatestBlock string `json:"lastBlock"`
}

func LatestBlockHandler(c *gin.Context) {
	st := Resp{}
	resp, err := http.Get("http://localhost:3001/status")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &st)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, string(body))
}

func VerifiersAmountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"verifiers_amount": "2",
	})
}

func TotalBalanceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"balance": 1677721600000000000,
	})
}

func HistoryAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "0",
	})
}

func HistoryTxHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"verifiers_amount": "0",
	})
}

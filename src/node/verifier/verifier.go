package verifier

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/ethereum"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

type Verifier struct {
	cfg *Config
}

func NewVerifier(cfg *Config) (*Verifier, error) {
	return &Verifier{}, nil
}

func (v *Verifier) Serve(ctx context.Context) error {
	// go listeners.Checker()
	// go balance.UpdateBalance(&storage.Balance, conf.Plasma_contract_address)
	// go event.Start(storage.Client, conf.Plasma_contract_address, &storage.Who, &storage.Amount, &storage.EventBlockHash, &storage.EventBlockNumber)

	r := gin.Default()
	// r := gin.New()
	// r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	r.GET("/etherBalance", v.EthereumBalance)
	r.GET("/verifiersAmount", v.VerifiersAmountHandler)
	r.GET("/totalBalance", v.TotalBalanceHandler)
	r.GET("/contractAddress", v.PlasmaContractAddress)
	r.GET("/deposit/:sum", v.DepositHandler)
	r.POST("/transfer/:address/:sum", v.TransferHandler)
	r.GET("/plasmaBalance", v.PlasmaBalance)
	r.GET("/exit", v.ExitHandler)
	r.GET("/latestBlock", v.LatestBlockHandler)

	r.Static("/frontend", "frontend")

	r.Run(":8080")

	return nil
}

func (v *Verifier) EthereumBalance(c *gin.Context) {
	response := ethereum.GetBalance(v.cfg.VerifierPublicKey)
	c.JSON(http.StatusOK, gin.H{
		"balance": response,
	})
}

func (v *Verifier) PlasmaBalance(c *gin.Context) {

	st := make([]blockchain.Input, 0)

	resp, err := http.Get("http://localhost:3001/utxo/" + v.cfg.VerifierPublicKey)
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

	c.JSON(http.StatusOK, st)
}

func (v *Verifier) PlasmaContractAddress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"address": v.cfg.PlasmaContractAddress,
	})
}

func (v *Verifier) DepositHandler(c *gin.Context) {
	result := ethereum.Deposit(c.Param("sum"))
	c.JSON(http.StatusOK, gin.H{
		"txHash": result,
	})
}

func (v *Verifier) TransferHandler(c *gin.Context) {
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

	key, _ := hex.DecodeString(v.cfg.VerifierPrivateKey)

	fmt.Println("Tx")
	fmt.Println(len(uTx.Outputs[0].Owner))

	uTx.Sign(key)

	fmt.Println("Signature")
	fmt.Println(len(uTx.Signatures))

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(uTx)
	res, _ := http.Post(v.cfg.OperatorHost+"/tx", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)

	c.JSON(http.StatusOK, gin.H{
		"status": res.Body,
	})
}

func (v *Verifier) ExitHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (v *Verifier) LatestBlockHandler(c *gin.Context) {
	st := struct {
		LatestBlock string `json:"lastBlock"`
	}{}

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

func (v *Verifier) VerifiersAmountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"verifiers_amount": "2",
	})
}

func (v *Verifier) TotalBalanceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"balance": 1677721600000000000,
	})
}

func (v *Verifier) HistoryAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "0",
	})
}

func (v *Verifier) HistoryTxHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"verifiers_amount": "0",
	})
}

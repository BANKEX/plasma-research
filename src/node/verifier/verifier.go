package verifier

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/BANKEX/plasma-research/src/node/ethereum/deposit"
	"github.com/BANKEX/plasma-research/src/node/ethereum/etherUtils"
	"github.com/ethereum/go-ethereum/common"

	"github.com/BANKEX/plasma-research/src/node/ethereum/transaction"
	"github.com/BANKEX/plasma-research/src/node/verifier/cli/completer"
	"github.com/c-bata/go-prompt"

	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/ethereum"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

type Verifier struct {
	serverInstance *gin.Engine
	cfg            *Config
	client         *ethclient.Client
}

func NewVerifier(cfg *Config) (*Verifier, error) {
	client, err := ethclient.Dial(cfg.GethHost)
	if err != nil {
		return nil, err
	}
	serverInstance, err := ServerInit()
	if err != nil {
		return nil, err
	}
	return &Verifier{
		cfg:            cfg,
		serverInstance: serverInstance,
		client:         client,
	}, nil
}

func (v *Verifier) Serve(ctx context.Context) error {
	go v.ServerStart(v.serverInstance)
	v.CLIToolStart()
	return nil
}

func (v *Verifier) CLIToolStart() {
	fmt.Println("------------Plasma Verifier----------")
	p := prompt.New(
		v.CLIToolExecutor,
		completer.Completer,
		prompt.OptionPrefix("--> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}

func ServerInit() (*gin.Engine, error) {

	//r := gin.Default()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	return r, nil
}

func (v *Verifier) CLIToolExecutor(userText string) {
	if userText == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	args := strings.Split(userText, " ")
	if len(args) > 2 {
		switch args[0] {
		case "plasma":
			switch args[1] {
			case "smartContractAddress", "sca":
				if len(args) == 3 {
					fmt.Println("Smart contract address: " + args[2])
				}
			case "plasmaBalance", "pb":
				if len(args) == 3 {
					fmt.Println("Plasma balance:" + args[2])
				}
			case "smartContractBalance", "scb":
				if len(args) == 3 {
					fmt.Println("Smart contract balance:" + args[2])
				}
			case "events", "e":
				if len(args) == 3 {
					fmt.Println("Events ...")
				}
			}
		case "eth":
			switch args[1] {
			case "smartContractAddress", "sca":
				if len(args) == 3 {
					fmt.Println("Smart contract address: " + args[2])
				}
			case "smartContractBalance", "scb":
				if len(args) == 3 {
					fmt.Println("Smart contract balance:" + args[2])
				}
			case "transfer", "tr":
				if len(args) == 4 {
					amountStr := args[2]
					recipientStr := args[3]

					amountInt64, err := strconv.ParseInt(amountStr, 10, 64)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(amountStr, recipientStr)
					transaction.SendTransactionInWei(v.client, v.cfg.VerifierPrivateKey, amountInt64, recipientStr)
				}
			case "accBalance", "ab":
				if len(args) == 3 {
					accStr := args[2]
					fmt.Println(accStr)

					//utils.AccountBalance(accStr)
				}
			}
		case "main":
			switch args[1] {
			case "deposit", "dep":
				if len(args) == 3 {
					value, ok := big.NewInt(0).SetString(args[2], 10)
					if !ok {
						log.Fatal(fmt.Errorf("given value not integer"))
					}
					rawPublicKey, err := etherUtils.ConvertStringPrivateKeyToRaw(v.cfg.VerifierPrivateKey)
					if err != nil {
						log.Fatal(err)
					}
					rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
					err = etherUtils.IsValidAddress(rawContractAddress)
					if err != nil {
						log.Fatal(err)
					}
					res, err := deposit.Deposit(context.TODO(), rawContractAddress, v.client, rawPublicKey, value)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(res.Hash().String())
				}
			case "exit", "ex":
				if len(args) == 3 {
					fmt.Println("Exit func")
				}
			}
		}
	} else {
		fmt.Println("Bad args!")
	}
}

func (v *Verifier) ServerStart(r *gin.Engine) error {
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

	err := r.Run(fmt.Sprintf(":%d", v.cfg.VerifierPort))
	if err != nil {
		log.Fatal(err)
		return err
	}
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
	value, ok := big.NewInt(0).SetString(c.Param("sum"), 10)
	if !ok {
		log.Fatal(fmt.Errorf("given value not integer"))
	}

	rawPublicKey, err := etherUtils.ConvertStringPrivateKeyToRaw(v.cfg.VerifierPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
	err = etherUtils.IsValidAddress(rawContractAddress)
	if err != nil {
		log.Fatal(err)
	}
	result, err := deposit.Deposit(context.TODO(), rawContractAddress, v.client, rawPublicKey, value)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
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

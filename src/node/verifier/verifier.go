package verifier

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/BANKEX/plasma-research/src/node/types"
	"github.com/gin-gonic/contrib/static"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/BANKEX/plasma-research/src/node/ethereum/deposit"
	"github.com/BANKEX/plasma-research/src/node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/BANKEX/plasma-research/src/node/ethereum/transaction"
	"github.com/BANKEX/plasma-research/src/node/verifier/cli/completer"
	"github.com/c-bata/go-prompt"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/ethereum"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/gin-gonic/gin"
)

type Verifier struct {
	rpcServer *gin.Engine
	cfg       *Config
	client    *ethclient.Client
	key       *ecdsa.PrivateKey
}

func NewVerifier(cfg *Config) (*Verifier, error) {
	client, err := ethclient.Dial(cfg.GethHost)
	if err != nil {
		return nil, err
	}

	key, err := crypto.HexToECDSA(cfg.VerifierPrivateKey[2:])
	if err != nil {
		return nil, err
	}

	return &Verifier{
		cfg:       cfg,
		rpcServer: utils.NewGinServer(),
		client:    client,
		key:       key,
	}, nil
}

func (v *Verifier) Serve(ctx context.Context) error {
	go v.ServerStart(v.rpcServer)
	v.CLIToolStart()
	return nil
}

func (v *Verifier) CLIToolStart() {
	log.Println("------------Plasma Verifier----------")

	if len(os.Args) > 1 {
		v.CLIToolExecutor(strings.Join(os.Args[1:], " "))
		os.Exit(0)
	} else {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Failed to create interactive console: ", r)
				fmt.Println("Running in non-interactive mode")
			}
		}()
		p := prompt.New(
			v.CLIToolExecutor,
			completer.Completer,
			prompt.OptionPrefix("--> "),
			prompt.OptionInputTextColor(prompt.Yellow),
		)
		p.Run()
	}
}

func (v *Verifier) CLIToolExecutor(userText string) {
	if userText == "quit" {
		log.Println("Bye!")
		os.Exit(0)
		return
	}
	args := strings.Split(userText, " ")
	if len(args) >= 2 {
		switch args[0] {
		case "eth":
			switch args[1] {
			case "transfer", "tr":
				if len(args) == 4 {
					value, ok := big.NewInt(0).SetString(args[2], 10)
					if !ok {
						log.Fatalf("")
					}

					if !common.IsHexAddress(args[3]) {
						log.Fatal(fmt.Errorf("given to address %s is not valid ethereum address", args[3]))
					}
					to := common.HexToAddress(args[3])

					tx, err := transaction.SendTransactionInWei(context.TODO(), v.client, v.key, value, to)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("transaction sended: %s", tx.Hash().String())
				}
			}
		case "plasma":
			switch args[1] {
			case "deposit", "dep":
				if len(args) == 3 {
					value, ok := big.NewInt(0).SetString(args[2], 10)
					if !ok {
						log.Fatal(fmt.Errorf("given value not integer"))
					}
					rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
					res, err := deposit.Deposit(context.TODO(), rawContractAddress, v.client, v.key, value)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(res.Hash().String())
				}
			case "transfer", "tr":
				if len(args) == 7 {
					block, err := strconv.ParseUint(args[2], 10, 32)
					if err != nil {
						log.Fatalf("")
					}
					txN, err := strconv.ParseUint(args[3], 10, 32)
					if err != nil {
						log.Fatalf("")
					}
					out, err := strconv.ParseUint(args[4], 10, 8)
					if err != nil {
						log.Fatalf("")
					}

					value, err := strconv.ParseUint(args[5], 10, 32)
					if err != nil {
						log.Fatalf("")
					}
					if !common.IsHexAddress(args[6]) {
						log.Fatal(fmt.Errorf("given to address %s is not valid ethereum address", args[6]))
					}
					to, err := hex.DecodeString(args[6][2:])
					if err != nil {
						log.Fatalf("")
					}

					txs, err := v.getTransactionHistory(v.cfg.VerifierPublicKey)
					if err != nil {
						log.Fatal("error ", err)
					}

					in := findTransaction(txs, uint32(block), uint32(txN), byte(out))
					if in == nil {
						log.Fatal("no such output")
					}

					_, err = v.sendToOperatorPlasmaTx(in, uint32(value), to)
					if err != nil {
						log.Fatal("error ", err)
					}
				}
			case "utxo":
				if len(args) == 2 {
					txs, err := v.getTransactionHistory(v.cfg.VerifierPublicKey)
					if err != nil {
						log.Fatal("error ", err)
					}
					log.Printf("Utxo list for %s:", v.cfg.VerifierPublicKey)
					for _, tx := range txs {
						log.Printf("%d:%d:%d -> %d coins", tx.BlockIndex, tx.TxIndex, tx.OutputIndex, tx.Slice.End-tx.Slice.Begin)
					}
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathToStatic := dir + "/src/node/verifier/frontend"

	r.Use(static.Serve("/", static.LocalFile(pathToStatic, true)))

	err = r.Run(fmt.Sprintf(":%d", v.cfg.VerifierPort))
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

	resp, err := http.Get(v.cfg.OperatorHost + "/utxo/" + v.cfg.VerifierPublicKey)
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

	rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
	result, err := deposit.Deposit(context.TODO(), rawContractAddress, v.client, v.key, value)
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

	key, err := hex.DecodeString(v.cfg.VerifierPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
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
	st := types.LastBlock{}
	resp, err := http.Get(v.cfg.OperatorHost + "/status")
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

	c.JSON(http.StatusOK, st.LastBlock)
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

func (v *Verifier) getTransactionHistory(address string) ([]blockchain.Input, error) {
	res, err := http.Get(v.cfg.OperatorHost + "/utxo/" + address)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := make([]blockchain.Input, 0)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *Verifier) getBalance(address string) (sum uint32, err error) {
	txs, err := v.getTransactionHistory(address)
	if err != nil {
		return 0, err
	}
	for _, tx := range txs {
		sum += tx.Slice.End - tx.Slice.Begin
	}
	return sum, nil
}

func (v *Verifier) sendToOperatorPlasmaTx(in *blockchain.Input, sum uint32, address []byte) (*http.Response, error) {

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

	key, err := hex.DecodeString(v.cfg.VerifierPrivateKey[2:])
	if err != nil {
		return nil, err
	}
	fmt.Println("Tx")
	fmt.Println(len(uTx.Outputs[0].Owner))

	err = uTx.Sign(key)
	if err != nil {
		return nil, err
	}

	fmt.Println("Signature")
	fmt.Println(len(uTx.Signatures))

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(uTx)
	if err != nil {
		return nil, err
	}
	res, _ := http.Post(v.cfg.OperatorHost+"/tx", "application/json; charset=utf-8", b)
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("success")
	return res, nil
}

func findTransaction(slice []blockchain.Input, block, tx uint32, out byte) *blockchain.Input {
	for _, item := range slice {
		if item.BlockIndex == block && item.TxIndex == tx && item.OutputIndex == out {
			return &item
		}
	}
	return nil
}

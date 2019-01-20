package verifier

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/BANKEX/plasma-research/src/node/types"
	"github.com/BANKEX/plasma-research/src/node/verifier/history"
	"github.com/gin-gonic/contrib/static"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	"github.com/imroc/req"
	"math"
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
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	var arguments []string

	argsWithSpace := strings.Split(userText, " ")

	for _, j := range argsWithSpace {
		if len(j) > 0 {
			arguments = append(arguments, j)
		}
	}

	if len(arguments) > 2 {
		firstWorld := arguments[0]
		switch firstWorld {
		case "eth":
			secondWorld := arguments[1]
			switch secondWorld {
			case "transfer", "tr":
				if len(arguments) == 4 {

					value, ok := big.NewInt(0).SetString(arguments[2], 10)

					if !ok {
						fmt.Println("Bad int!")
					}

					if !common.IsHexAddress(arguments[3]) {
						fmt.Println(fmt.Errorf("given to address %s is not valid ethereum address", arguments[3]))
					}

					to := common.HexToAddress(arguments[3])

					tx, err := transaction.SendTransactionInWei(context.TODO(), v.client, v.key, value, to)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("transaction sended: %s", tx.Hash().String())
					}
				} else if len(arguments) < 4 {
					fmt.Println("Not anough arguments!")
				} else {
					fmt.Println("Bad arguments!")
				}
			case "balance", "bal":
				if len(arguments) == 3 {
					balanceFloat, err := GetEtherAccountBalance(arguments[2])
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("Balance of account %s : %f\n", arguments[2], balanceFloat)
					}
				} else {
					fmt.Println("Bad arguments!")
				}
				// TODO:check this method
				// now not work correctly
			case "ownerBalance", "ob":
				if len(arguments) == 2 {
					fmt.Println(v.cfg.VerifierEthereumAddress)
					balanceFloat, err := GetEtherAccountBalance(v.cfg.VerifierEthereumAddress)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("Balance of account %s : %f\n", v.cfg.VerifierEthereumAddress, balanceFloat)
					}
				} else {
					fmt.Println("Bad arguments!")
				}
			default:
				fmt.Println("Bad arguments!")
			}
		case "plasma":
			secondWorld := arguments[1]
			switch secondWorld {
			case "deposit", "dep":
				if len(arguments) == 3 {
					value, ok := big.NewInt(0).SetString(arguments[2], 10)
					switch ok {
					case true:
						rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
						res, err := deposit.Deposit(context.TODO(), rawContractAddress, v.client, v.key, value)
						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(res.Hash().String())
						}
					case false:
						fmt.Println("Error!")
					}
				} else if len(arguments) < 3 {
					fmt.Println("Not enough arguments!")
				} else {
					fmt.Println("Bad arguments!")
				}
			case "transfer", "tr":
				if len(arguments) == 7 {
					block, err := strconv.ParseUint(arguments[2], 10, 32)
					if err != nil {
						fmt.Println(err)
					}
					txN, err := strconv.ParseUint(arguments[3], 10, 32)
					if err != nil {
						fmt.Println(err)
					}
					out, err := strconv.ParseUint(arguments[4], 10, 8)
					if err != nil {
						fmt.Println(err)
					}

					value, err := strconv.ParseUint(arguments[5], 10, 32)
					if err != nil {
						fmt.Println(err)
					}
					if !common.IsHexAddress(arguments[6]) {
						fmt.Println(fmt.Errorf("given to address %s is not valid ethereum address", arguments[6]))
					}
					to, err := hex.DecodeString(arguments[6][2:])
					if err != nil {
						fmt.Println(err)
					}

					txs, err := v.getTransactionHistory(v.cfg.VerifierEthereumAddress)
					if err != nil {
						fmt.Println(err)
					}

					in := findTransaction(txs, uint32(block), uint32(txN), byte(out))
					if in == nil {
						fmt.Println("no such output")
					}

					_, err = v.sendToOperatorPlasmaTx(in, uint32(value), to)
					if err != nil {
						fmt.Println(err)
					}
				} else if len(arguments) < 7 {
					fmt.Println("Not anought arguments!")
				} else {
					fmt.Println("Bad arguments!")
				}
			case "utxo":
				if len(arguments) == 2 {
					txs, err := v.getTransactionHistory(v.cfg.VerifierEthereumAddress)
					if err != nil {
						fmt.Println("error ", err)
					}

					fmt.Printf("Utxo list for %s:", v.cfg.VerifierEthereumAddress)

					for _, tx := range txs {
						fmt.Printf("%d:%d:%d -> %d coins", tx.BlockIndex, tx.TxIndex, tx.OutputIndex, tx.Slice.End-tx.Slice.Begin)
					}

				} else {
					fmt.Println("Not anought arguments!")
				}

				// TODO:check this method
				// now not work correctly
			case "balance", "bal":
				if len(arguments) == 2 {
					st := make([]blockchain.Input, 0)
					resp, err := req.Get(v.cfg.OperatorHost + "/utxo/" + v.cfg.VerifierEthereumAddress)
					if err != nil {
						fmt.Println(err)
					}
					resp.ToJSON(&st)
					fmt.Println(st)

				} else {
					fmt.Println("Not anought arguments!")
				}
			case "exit", "ex":
				if len(arguments) == 3 {
					fmt.Println("Exit func")
				}
			default:
				fmt.Println("Bad arguments!")
			}
		}
	}
}

func (v *Verifier) ServerStart(r *gin.Engine) error {

	r.POST("/deposit", v.DepositHandler)
	r.POST("/transfer", v.TransferHandler)
	r.POST("/exit", v.ExitHandler)
	r.GET("/common", v.CommonInfoHandler)
	r.GET("/history", v.HistoryAllHandler)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathToStatic := dir + "/src/node/verifier/frontend"

	r.Use(static.Serve("/", static.LocalFile(pathToStatic, true)))

	err = r.Run(fmt.Sprintf(":%d", v.cfg.VerifierPort))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (v *Verifier) DepositHandler(c *gin.Context) {

	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)

	e := new(history.Event)

	err := json.Unmarshal([]byte(x), e)
	if err != nil {
		fmt.Println(err)
	}

	value, ok := big.NewInt(0).SetString(e.Sum, 10)
	if !ok {
		fmt.Println(fmt.Errorf("given value not integer"))
	}

	rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
	result, err := deposit.Deposit(context.TODO(), rawContractAddress, v.client, v.key, value)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	t := time.Now().Unix()

	history.Log(value, t, "", "Deposit")

	c.JSON(http.StatusOK, gin.H{
		"txHash": result,
	})

	c.Next()
	return
}

func (v *Verifier) TransferHandler(c *gin.Context) {

	body := c.Request.Body
	x, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	e := new(history.Event)

	err = json.Unmarshal([]byte(x), e)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	value, ok := big.NewInt(0).SetString(e.Sum, 10)
	if !ok {
		log.Println(fmt.Errorf("given value not integer"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("given value not integer"),
		})
	}
	t := time.Now().Unix()

	history.Log(value, t, e.Who, "Transfer")

	address, err := hex.DecodeString(e.Who)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	usum, err := strconv.Atoi(e.Sum)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	sum := uint32(usum)

	in := new(blockchain.Input)
	err = c.BindJSON(&in)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

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
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	fmt.Println("Tx")
	fmt.Println(len(uTx.Outputs[0].Owner))

	err = uTx.Sign(key)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	fmt.Println("Signature")
	fmt.Println(len(uTx.Signatures))

	b := new(bytes.Buffer)

	err = json.NewEncoder(b).Encode(uTx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	res, err := http.Post(v.cfg.OperatorHost+"/tx", "application/json; charset=utf-8", b)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": res.Body,
	})

	c.Next()
	return
}

func (v *Verifier) ExitHandler(c *gin.Context) {
	e := new(history.Event)

	value, ok := big.NewInt(0).SetString(e.Sum, 10)
	if !ok {
		log.Println(fmt.Errorf("given value not integer"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("given value not integer"),
		})
	}
	t := time.Now().Unix()

	history.Log(value, t, e.Who, "Exit")

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
func (v *Verifier) CommonInfoHandler(c *gin.Context) {

	st := make([]blockchain.Input, 0)

	resp, err := http.Get(v.cfg.OperatorHost + "/utxo/" + v.cfg.VerifierEthereumAddress)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	err = json.Unmarshal(body, &st)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	respInputs := make([]types.InputResponse, 0)

	for i := range st {
		in := types.InputResponse{}
		in.OutputResponse.Owner = hex.EncodeToString(st[i].Owner)
		in.OutputResponse.Slice.Begin = st[i].Slice.Begin
		in.OutputResponse.Slice.End = st[i].Slice.End
		in.TxIndex = st[i].TxIndex
		in.BlockIndex = st[i].BlockIndex
		in.OutputIndex = st[i].OutputIndex

		respInputs = append(respInputs, in)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	contractBalance, err := GetEtherAccountBalance(v.cfg.PlasmaContractAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	verifierEtherBalance, err := GetEtherAccountBalance(v.cfg.VerifierEthereumAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	verifierPlasmaBalance, err := v.getBalance(v.cfg.VerifierEthereumAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	lb := types.LastBlock{}
	respTwo, err := http.Get(v.cfg.OperatorHost + "/status")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	bodyTwo, err := ioutil.ReadAll(respTwo.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	err = json.Unmarshal(bodyTwo, &lb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"contract_address":        v.cfg.PlasmaContractAddress,
		"contract_balance":        contractBalance,
		"verifier_ether_balance":  verifierEtherBalance,
		"verifier_plasma_balance": verifierPlasmaBalance,
		"latest_block":            lb.LastBlock,
		"verifier_inputs":         respInputs,
	})
}

func (v *Verifier) HistoryAllHandler(c *gin.Context) {
	o := history.GetAllOperations()

	events := make([]history.Event, 0)

	for _, a := range o {
		event := history.Event{}
		event.Who = a.Who
		event.Sum = a.Sum.String()
		event.Date = a.Date
		event.OperationType = a.OperationType
		events = append(events, event)
	}

	c.JSON(http.StatusOK, gin.H{
		"Events": events,
	})
	c.Next()
	return
}

func (v *Verifier) getTransactionHistory(address string) ([]blockchain.Input, error) {
	result := make([]blockchain.Input, 0)

	resp, err := req.Get(v.cfg.OperatorHost + "/utxo/" + address)
	if err != nil {
		fmt.Println(err)
	}

	resp.ToJSON(&result)

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
	res, err := http.Post(v.cfg.OperatorHost+"/tx", "application/json; charset=utf-8", b)
	if err != nil {
		return nil, err
	}
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

func GetEtherAccountBalance(address string) (float64, error) {

	b, err := strconv.ParseFloat(ethereum.GetBalance(address), 64)
	if err != nil {
		return 0, err
	}

	return b / math.Pow(10, 18), err
}

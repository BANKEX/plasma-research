package verifier

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/BANKEX/plasma-research/src/node/types"
	"github.com/gin-gonic/contrib/static"

	"github.com/BANKEX/plasma-research/src/node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/BANKEX/plasma-research/src/node/verifier/cli/completer"
	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/ethereum"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/BANKEX/plasma-research/src/node/verifier/cli/options"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

type Verifier struct {
	rpcServer *gin.Engine
	cfg       *Config
	key       *ecdsa.PrivateKey
	eth       *ethereum.Ethereum
}

func NewVerifier(cfg *Config) (*Verifier, error) {
	plasmaContractAddress := common.HexToAddress(cfg.PlasmaContractAddress)

	eth, err := ethereum.NewEthereum(cfg.GethHost, plasmaContractAddress)
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
		key:       key,
		eth:       eth,
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

func splitAndTrimArgs(userText string) []string {
	var result []string
	// NOTE: can got problems with tabs, '\n', and so on
	for _, j := range strings.Split(userText, " ") {
		if len(j) > 0 {
			result = append(result, j)
		}
	}
	return result
}

func parseBigInt(s string) (*big.Int, bool) {
	return big.NewInt(0).SetString(s, 10)
}

func (v *Verifier) CLIToolExecutor(userText string) {
	arguments := splitAndTrimArgs(userText)

	if len(arguments) < 1 {
		return
	}

	command := arguments[0]
	if command == "quit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	switch command {
	case "eth":
		secondWorld := arguments[1]
		switch secondWorld {

		case "transfer":
			if len(arguments) != 4 {
				fmt.Println(options.Eth["transfer"])
				return
			}

			amountInWeiArg, toAddressArg := arguments[2], arguments[3]
			amountInWei, ok := parseBigInt(amountInWeiArg)
			if !ok {
				fmt.Println("Can't parse amount of ether in wei")
				return
			}

			if !common.IsHexAddress(toAddressArg) {
				fmt.Printf("Given address '%s' is not valid ethereum address\n", toAddressArg)
				return
			}
			to := common.HexToAddress(toAddressArg)

			tx, err := v.eth.SendTransactionInWei(context.TODO(), v.key, amountInWei, to)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("transaction sended: %s", tx.Hash().String())

		case "balance":
			if len(arguments) != 3 {
				fmt.Println(options.Eth["balance"])
				return
			}
			balanceFloat, err := GetETHAccountBalance(arguments[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Balance of account %s : %f\n", arguments[2], balanceFloat)

		case "ownerBalance":
			if len(arguments) != 2 {
				fmt.Println(options.Eth["ownerBalance"])
				return
			}
			balanceFloat, err := GetETHAccountBalance(v.cfg.VerifierEthereumAddress)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Balance of account %s : %f\n", v.cfg.VerifierEthereumAddress, balanceFloat)

		default:
			fmt.Println("Bad arguments!")
			return
		}

	case "plasma":
		secondWorld := arguments[1]
		switch secondWorld {
		case "deposit":
			if len(arguments) != 3 {
				fmt.Println(options.Plasma["deposit"])
				return
			}

			value, ok := big.NewInt(0).SetString(arguments[2], 10)
			switch ok {
			case true:
				rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
				res, err := v.eth.Deposit(context.TODO(), rawContractAddress, v.key, value)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(res.Hash().String())
			case false:
				fmt.Println(options.Plasma["deposit"])
				return
			}
		case "transfer":
			if len(arguments) != 7 {
				fmt.Println(options.Plasma["transfer"])
				return
			}

			block, err := strconv.ParseUint(arguments[2], 10, 32)
			if err != nil {
				fmt.Println(err)
				return
			}
			txN, err := strconv.ParseUint(arguments[3], 10, 32)
			if err != nil {
				fmt.Println(err)
				return
			}
			out, err := strconv.ParseUint(arguments[4], 10, 8)
			if err != nil {
				fmt.Println(err)
				return
			}

			value, err := strconv.ParseUint(arguments[5], 10, 32)
			if err != nil {
				fmt.Println(err)
				return
			}
			if !common.IsHexAddress(arguments[6]) {
				fmt.Println(fmt.Errorf("given to address %s is not valid ethereum address", arguments[6]))
				return
			}
			to, err := hex.DecodeString(arguments[6][2:])
			if err != nil {
				fmt.Println(err)
				return
			}

			txs, err := v.getTransactionHistory(v.cfg.VerifierEthereumAddress)
			if err != nil {
				fmt.Println(err)
				return
			}

			in := findTransaction(txs, uint32(block), uint32(txN), byte(out))
			if in == nil {
				fmt.Println("no such output")
				return
			}

			_, err = v.sendToOperatorPlasmaTx(in, uint32(value), to)
			if err != nil {
				fmt.Println(err)
				return
			}

		case "utxo":
			if len(arguments) != 2 {
				fmt.Println(options.Plasma["utxo"])
				return
			}
			txs, err := v.getTransactionHistory(v.cfg.VerifierEthereumAddress)
			if err != nil {
				fmt.Println("error ", err)
				return
			}
			fmt.Printf("Utxo list for %s:", v.cfg.VerifierEthereumAddress)
			for _, tx := range txs {
				fmt.Printf("%d:%d:%d -> %d coins", tx.BlockIndex, tx.TxIndex, tx.OutputIndex, tx.Slice.End-tx.Slice.Begin)
			}

		case "balance":
			if len(arguments) != 2 {
				fmt.Println(options.Plasma["balance"])
				return
			}

			st := make([]blockchain.Input, 0)
			resp, err := req.Get(v.cfg.OperatorHost + "/utxo/" + v.cfg.VerifierEthereumAddress)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp.ToJSON(&st)
			fmt.Println(st)

		case "exit":
			if len(arguments) != 2 {
				fmt.Println(options.Plasma["exit"])
				return
			}
			fmt.Println("Exit func ...")

		default:
			fmt.Println("Bad arguments!")
			return
		}
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
		return err
	}
	pathToStatic := dir + "/src/node/verifier/frontend"

	r.Use(static.Serve("/", static.LocalFile(pathToStatic, true)))

	err = r.Run(fmt.Sprintf(":%d", v.cfg.VerifierPort))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (v *Verifier) EthereumBalance(c *gin.Context) {
	addr := common.HexToAddress(v.cfg.VerifierEthereumAddress)
	response, err := v.eth.GetBalance(context.TODO(), addr)
	if err != nil {
		log.Printf("failed to get verifier eth balance: %s", err)
		// TODO: handle error
	}
	c.JSON(http.StatusOK, gin.H{
		"balance": response,
	})
}

func (v *Verifier) PlasmaBalance(c *gin.Context) {

	st := make([]blockchain.Input, 0)

	resp, err := req.Get(v.cfg.OperatorHost + "/utxo/" + v.cfg.VerifierEthereumAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	resp.ToJSON(&st)

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
		log.Println("given value not integer")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "given value not integer"})
		return
	}

	rawContractAddress := common.HexToAddress(v.cfg.PlasmaContractAddress)
	result, err := v.eth.Deposit(context.TODO(), rawContractAddress, v.key, value)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
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
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
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

	resp, err := req.Get(v.cfg.OperatorHost + "/status")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	resp.ToJSON(&st)

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
	result := make([]blockchain.Input, 0)

	resp, err := req.Get(v.cfg.OperatorHost + "/utxo/" + address)
	if err != nil {
		log.Println(err)
		return nil, err
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

func GetETHAccountBalance(address string) (float64, error) {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		return 0, err
	}

	ctx := context.Background()

	account := common.HexToAddress(address)

	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		return 0, err
	}
	ethBalance, _ := strconv.ParseFloat(balance.String(), 64)

	balanceFloat := ethBalance / math.Pow(10, 18)

	return balanceFloat, err
}

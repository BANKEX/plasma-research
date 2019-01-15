package operator

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BANKEX/plasma-research/src/node/ethereum/plasmacontract"
	"github.com/BANKEX/plasma-research/src/node/types"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/ethereum/events"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/BANKEX/plasma-research/src/node/transactionManager"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

type Operator struct {
	cfg *Config

	txManager *transactionManager.TransactionManager
	publisher *transactionManager.BlockPublisher
	monitor   *transactionManager.EventMonitor
}

func NewOperator(cfg *Config) (*Operator, error) {
	manager := transactionManager.NewTransactionManager()
	publisher, err := transactionManager.NewBlockPublisher(manager)
	if err != nil {
		return nil, err
	}

	eventMonitor, err := transactionManager.NewEventMonitor(manager, publisher)
	if err != nil {
		return nil, err
	}

	// TODO: refactor this place
	go events.EventListener(manager)

	return &Operator{
		cfg:       cfg,
		txManager: manager,
		publisher: publisher,
		monitor:   eventMonitor,
	}, nil
}

func (o *Operator) Serve(ctx context.Context) error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	r.POST("/tx", o.PostTransaction)
	r.GET("/config", o.GetConfig)
	r.GET("/status", o.GetStatus)
	r.GET("/utxo/:address", o.GetUtxos)

	// debug handlers
	r.GET("/test/fund/:address", o.FundAddress)
	r.GET("/test/transfer/:block/:tx/:out/:address/:key", o.Transact)

	err := r.Run(fmt.Sprintf(":%d", o.cfg.OperatorPort))
	if err != nil {
		return err
	}

	println("Operator started")

	return nil
}

func (o *Operator) PostTransaction(c *gin.Context) {
	var t blockchain.Transaction
	err := c.BindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = o.txManager.SubmitTransaction(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// returns a list of utxos for an address
func (o *Operator) GetUtxos(c *gin.Context) {
	addr := c.Param("address")
	utxos, err := o.txManager.GetUtxosForAddress(addr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, utxos)
}

// returns last plasma block number etc.
func (o *Operator) GetStatus(c *gin.Context) {
	st := types.LastBlock{}
	st.LastBlock = strconv.Itoa(int(o.txManager.GetLastBlockNumber()))
	c.JSON(http.StatusOK, st)
}

// returns contract address and abi
func (o *Operator) GetConfig(c *gin.Context) {
	info := types.OperatorInfo{}
	info.Config = o.cfg
	info.ABI = store.StoreABI
	c.JSON(http.StatusOK, info)
}

// ==== debug handlers =====

// returns a list of utxos for an address
func (o *Operator) FundAddress(c *gin.Context) {
	addr, _ := hex.DecodeString(c.Param("address")[2:])
	out := blockchain.Output{
		Owner: addr,
		Slice: slice.Slice{Begin: 10, End: 20},
	}
	_, err := o.txManager.AssembleDepositBlock(out)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "OK")
}

func (o *Operator) Transact(c *gin.Context) {
	addr, _ := hex.DecodeString(c.Param("address")[2:])
	key, _ := hex.DecodeString(c.Param("key")[2:])
	blockN, _ := strconv.Atoi(c.Param("block"))
	txN, _ := strconv.Atoi(c.Param("tx"))
	outN, _ := strconv.Atoi(c.Param("out"))

	in := o.txManager.GetUtxo(uint32(blockN), uint32(txN), uint32(outN))

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

	err = o.txManager.SubmitTransaction(&tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "OK")
}

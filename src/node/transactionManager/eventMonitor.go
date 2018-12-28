package transactionManager

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/BANKEX/plasma-research/src/node/config"
	"github.com/BANKEX/plasma-research/src/node/ethereum/plasmacontract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventMonitor struct {
	transactionManager *TransactionManager
	blockPublisher     *BlockPublisher
	client             *ethclient.Client
	currentBlock       uint64
}

func NewEventMonitor(m *TransactionManager, p *BlockPublisher) (*EventMonitor, error) {
	c, err := ethclient.Dial(config.GetOperator().GethHost)
	if err != nil {
		return nil, err
	}

	result := EventMonitor{
		transactionManager: m,
		client:             c,
		currentBlock:       config.GetOperator().StartingBlock,
	}

	manager = m
	publisher = p

	go result.loop()

	return &result, nil
}

//// todo monitor deposit events, forward to transaction manager
//// todo monitor withdraw events, forward to transaction manager
//// todo if we need to send some challenges from the operator, this is the place to do it

func (m *EventMonitor) loop() {
	for {
		latest, err := m.client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Println("Error getting latest block number")
			time.Sleep(time.Second * 10)
			continue
		}
		if latest.Number.Uint64() <= m.currentBlock {
			time.Sleep(time.Second)
			continue
		}

		err = m.processBlock(m.currentBlock)
		if err == nil {
			log.Printf("processed block %d for events", m.currentBlock)
			m.currentBlock++
		} else {
			log.Printf("error processing block %d for events: %s", m.currentBlock, err.Error())
			time.Sleep(time.Second * 10)
		}
	}
}

func (m *EventMonitor) processBlock(blockNumber uint64) error {
	contractAddress := common.HexToAddress(config.GetOperator().PlasmaContractAddress[2:])
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := m.client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}
	err = m.processLogs(logs)
	if err != nil {
		return err
	}

	return nil
}

func (m *EventMonitor) processLogs(logs []types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		return err
	}

	for _, item := range logs {
		for _, h := range Handlers {
			if crypto.Keccak256Hash([]byte(h.Signature)).Hex() == item.Topics[0].Hex() {
				log.Println("Received event", strings.Split(h.Signature, "(")[0])
				h.Handler(item, contractAbi)
			}
		}
	}

	return nil
}

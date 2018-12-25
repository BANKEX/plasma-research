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

const (
	statusSuccess     = iota
	statusError       = iota
	statusNoSuchBlock = iota
)

type EventMonitor struct {
	transactionManager *TransactionManager
	client             *ethclient.Client
	currentBlock       uint64
}

func NewEventMonitor(m *TransactionManager) (*EventMonitor, error) {
	c, err := ethclient.Dial(config.GetOperator().GethHost)
	if err != nil {
		return nil, err
	}

	result := EventMonitor{
		transactionManager: m,
		client:             c,
		currentBlock:       config.GetOperator().StartingBlock,
	}

	go result.loop()

	return &result, nil
}

//// todo monitor deposit events, forward to transaction manager
//// todo monitor withdraw events, forward to transaction manager
//// todo if we need to send some challenges from the operator, this is the place to do it

func (m *EventMonitor) loop() {
	for {
		result, err := m.processBlock(m.currentBlock)
		switch result {
		case statusSuccess:
			log.Printf("processed block %d for events", m.currentBlock)
			m.currentBlock++
			break
		case statusError:
			log.Printf("error processing block %d for events: %s", m.currentBlock, err.Error())
			time.Sleep(time.Second * 10)
			break
		case statusNoSuchBlock:
			time.Sleep(time.Second)
		}
	}
}

func (m *EventMonitor) processBlock(blockNumber uint64) (int, error) {
	contractAddress := common.HexToAddress(config.GetOperator().MainAccountPublicKey[2:])
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := m.client.FilterLogs(context.Background(), query)
	if err != nil {
		return statusError, err
	}
	err = m.processLogs(logs)
	if err != nil {
		return statusError, err
	}

	return statusSuccess, nil
}

func (m *EventMonitor) processLogs(logs []types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		return err
	}

	for _, item := range logs {
		for _, h := range Handlers {
			if crypto.Keccak256Hash([]byte(h.Signature)).Hex() == item.Topics[0].Hex() {
				for range item.Topics {
					h.Handler(item, contractAbi)
				}

			}
		}
	}

	return nil
}

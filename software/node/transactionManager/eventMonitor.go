package transactionManager

import (
	"../config"
	"../ethereum/plasmacontract"
	"./eventHandlers"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
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
	//
}

func (m *EventMonitor) processBlock(blockNumber int64) (int, error) {
	contractAddress := common.HexToAddress(config.GetOperator().MainAccountPublicKey[2:])
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
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

	for _, vLog := range logs {
		for _, h := range eventHandlers.Handlers {
			if crypto.Keccak256Hash([]byte(h.Signature)).Hex() == vLog.Topics[0].Hex() {
				for range vLog.Topics {
					h.Handler(vLog, contractAbi)
				}

			}
		}
	}

	return nil
}

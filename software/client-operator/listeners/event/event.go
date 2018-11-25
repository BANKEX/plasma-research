package event

import (
	"context"
	"log"
	"math/big"
	"strings"

	"../storage"

	"../plasmacontract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var EventCount int
var EventMap = map[int]map[string]string{}

func SubscribeToEvent(e *ethclient.Client, addr string) (abi.ABI, ethereum.Subscription, chan types.Log, bool) {
	client := e
	contractAddress := common.HexToAddress(addr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		return abi.ABI{}, nil, nil, false
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return abi.ABI{}, nil, nil, false
	}
	return contractAbi, sub, logs, true
}

func Start(client *ethclient.Client, addr string, who *common.Address, amount **big.Int, eventBlockHash *string, blockNumber *uint64) {
	abi, sub, logs, errbool := SubscribeToEvent(client, addr)
	if errbool == false {
		println(errbool)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			Sig := []byte("CoinDeposited(address,uint256)")
			SigHash := crypto.Keccak256Hash(Sig)
			switch vLog.Topics[0].Hex() {
			case SigHash.Hex():
				event := struct {
					Who    common.Address
					Amount *big.Int
				}{}
				err := abi.Unpack(&event, "CoinDeposited", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println("------------------------------")
				// fmt.Println("Triggered event: CoinDeposited")

				// fmt.Print("Who: ")
				// fmt.Println(common.HexToAddress(vLog.Topics[1].Hex()).String())

				// fmt.Print("Sum in wei: ")
				// fmt.Println(event.Amount.String())
				// fmt.Println("------------------------------")

				*who = common.HexToAddress(vLog.Topics[1].Hex())
				*amount = event.Amount
				*eventBlockHash = vLog.BlockHash.Hex()
				*blockNumber = vLog.BlockNumber
				storage.StateForEvent = 1

			}
		}
	}
}

type EventCoinDeposited struct {
	Who    common.Address
	Amount *big.Int
}

type EventGroup struct {
	EventCoinDeposited EventCoinDeposited
}

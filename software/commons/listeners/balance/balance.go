package balance

import (
	"context"
	"log"

	"../storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(e *ethclient.Client, addr string) string {
	bal, err := e.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		log.Println(err)
	}
	return bal.String()
}

func UpdateBalance(balance *string, connection string) {
	for {
		*balance = GetBalance(storage.Client, connection)
	}
}

package ethClient

import (
	"log"

	"../storage"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InitClient(rawUrl string) {
	client, err := ethclient.Dial(rawUrl)
	if err != nil {
		log.Println(err)
	}
	storage.Client = client
}

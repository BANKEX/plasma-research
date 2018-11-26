package dispatchers

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"../blocks"
	"../db"
	"github.com/ethereum/go-ethereum/common"
)

type TxTest struct {
	From int
	To   *common.Address
	Sum  *big.Int
}

var BlockCounter int

func CreateGenesisBlock() {
	var genBlock blocks.Block

	hash, err := genBlock.GetHash()
	if err != nil {
		log.Println(err)
	}

	err = db.Block("database").Put([]byte(strconv.Itoa(BlockCounter)), hash)
	if err != nil {
		log.Println(err)
	}

	BlockCounter++

	fmt.Println("Genesis Block was created!!!")
	fmt.Println("\n\n")

}

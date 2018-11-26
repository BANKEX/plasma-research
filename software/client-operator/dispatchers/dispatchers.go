package dispatchers

import (
	"../db"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strconv"
)

//type Hashes interface {
//	GetHash() []byte
//}

var BlockCounter = 0

// Test tx struct
type TxTest struct {
	From *common.Address
	To   *common.Address
	Sum  *big.Int
}
// for tx hash
func (tx *TxTest) GetHash() []byte{
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(*tx)
	return reqBodyBytes.Bytes()
}

// Test block struct
type BlockTest struct {
	LastBlockHash []byte
	TxHashes [][]byte // 1 tx for test
	ThisBlockHash []byte // hash of this struct
}
// for block hash
func (bt *BlockTest) GetHash() []byte{
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(*bt)
	return reqBodyBytes.Bytes()
}


// Hash for test Tx
func TxToBytesArr(s TxTest) []byte{
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	return reqBodyBytes.Bytes()

}

// Hash for test block
func BlockHash(s BlockTest) []byte{
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	return reqBodyBytes.Bytes()

}

func CreateGenesisBlock() {

	var genBlock BlockTest

	hash := BlockHash(genBlock)

	err := db.Block("database").Put([]byte(strconv.Itoa(BlockCounter)), hash)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Genesis Block was created!!!")


	fmt.Println("\n\n")

	check, err := db.Block("database").GetAll()
	if err != nil {
		println("Mistake DB")
	}
	fmt.Println(check)


}

func GetTxHashTest(who *common.Address, amount *big.Int)[]byte{
	tx := TxTest{To:who, Sum:amount,From:who} // from == who for test
	hash := crypto.Keccak256(TxToBytesArr(tx))

	return hash
}

func PushTestBlock(tx [][]byte){

	block := BlockTest{TxHashes:tx}
	block.ThisBlockHash = BlockHash(block)
	lbh, err := db.Block("database").Get([]byte(strconv.Itoa(BlockCounter)))
	if err != nil {
		println("Mistake DB")
	}
	block.LastBlockHash = lbh

	// push to leveldb
	BlockCounter++
	key := []byte(strconv.Itoa(BlockCounter))
	value := BlockHash(block)

	err = db.Block("database").Put(key,value)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Add new Block! BlockCounter: " + strconv.Itoa(BlockCounter))

}
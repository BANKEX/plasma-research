package main

import (
	"../blockExplorer/blocks"
	// "../blockExplorer/ethereum"
	// "../blockExplorer/rlp"
	// "encoding/binary"
	"fmt"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	// "strconv"
	"encoding/json"
)

type tes struct {
	val big.Int
}

func main() {
	blockString, _ := json.MarshalIndent(blocks.CreateBlock(), "", "  ")
	fmt.Println("Works")
	fmt.Printf("%s", blockString)
}

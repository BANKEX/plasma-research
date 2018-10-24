package main

import (
	"../blockExplorer/blocks"
	// "../blockExplorer/ethereum"
	// "../blockExplorer/rlp"
	// "encoding/binary"
	"fmt"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "math/big"
	// "strconv"
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/rlp"
)

type Foo struct {
	Data uint
}

type Bar struct {
	Foo
	Sign uint
}

func main() {

	bar := Bar{}
	b := new(bytes.Buffer)
	err := rlp.Encode(b, bar)
	result := b.Bytes()
	fmt.Println(result, err)

	blockString, _ := json.MarshalIndent(blocks.CreateBlock(), "", "  ")
	//fmt.Println("Works")
	fmt.Printf("%s", blockString)
}

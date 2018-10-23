package main

import (
	"../blockExplorer/blocks"
	"../blockExplorer/ethereum"
	// "../blockExplorer/rlp"
	// "encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	// "math/big"
	// "strconv"
)

var secretKey, _ = hexutil.Decode("0x3d345c1036f325e046c8a013707def71a8854f563f7b17f7c61d81975c7de479")
var wallet, _ = ethereum.Wallet(secretKey)

func main() {
	block := blocks.Block{}
	fmt.Println("Works", secretKey, block)
}

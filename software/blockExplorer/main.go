package main

import (
	"../blockExplorer/types"
	"../blockExplorer/rlp"
)

func main() {
	block := types.BlockInit()

	val := rlp.EncodeBlock(&block)

	println(val)

}

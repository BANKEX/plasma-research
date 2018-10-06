package main

import (
	"../blockExplorer/types"
	"../blockExplorer/rlp"
)

func main() {
	block := types.BlockInit()
	val := rlp.EncodeBlock(&block)

	t:= rlp.DecodeBlock(val)

	println(t.BlockHeader.BlockNumber)
}

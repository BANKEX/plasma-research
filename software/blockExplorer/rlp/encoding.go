package rlp

import (
	"../types"
	"github.com/ethereum/go-ethereum/rlp"
)

//EncodeBlock - will return an encoded block in bytes
func EncodeBlock(block *types.Block) []byte {
	val, _ := rlp.EncodeToBytes(block)
	return val
}

//DecodeBlocl - will decode to Block structure from bytes
func DecodeBlock(v []byte) types.Block {
	b := types.Block{}
	rlp.DecodeBytes(v, b)
	return b
}

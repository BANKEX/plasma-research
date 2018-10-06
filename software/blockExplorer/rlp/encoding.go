package rlp

import (
	"../types"
	"github.com/ethereum/go-ethereum/rlp"
)

// EncodeBlock - will return an encoded block in bytes
func EncodeBlock(block *types.Block) []byte {
	val, _ := rlp.EncodeToBytes(block)
	return val
}

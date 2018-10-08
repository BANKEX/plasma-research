package rlp

import (
	"../types"
	"github.com/ethereum/go-ethereum/rlp"
	"testing"
)

// EncodeBlock - will return an encoded block in bytes
func EncodeBlock(block *types.Block) []byte {
	val, _ := rlp.EncodeToBytes(block)
	return val
}

// DecodeBlock - will decode to Block structure from bytes
func DecodeBlock(v []byte) types.Block {
	b := types.Block{}
	rlp.DecodeBytes(v, b)
	return b
}


func TestEncodeBlock(t *testing.T) {
	// TODO: Step1. Create block, assign value to the fields and save to RLP
	// TODO: Step2. Compare bytes with sample predefined bytes[]
}

func TestDecodeBlock(t *testing.T) {
	// TODO: Get hardcoded RLP values and check that we can successfully decode it
}
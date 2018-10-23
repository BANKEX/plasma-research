package blocks

import (
	"math/big"
)

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type BlockHeader struct {
	BlockNumber       uint32   `json:"blockNumber"`
	PreviousBlockHash big.Int  `json:"previousBlockHash"`
	MerkleRoot        big.Int  `json:"merkleRoot"`
	Signature         [65]byte `json:"signature"`
}

func (b *Block) Sign(key []byte) error {
	// b.Header.Signature =
	return nil
}

func (b *Block) CalculateMerkleRoot() error {
	// b.Header.MerkleRoot =
	return nil
}

// === validation ===
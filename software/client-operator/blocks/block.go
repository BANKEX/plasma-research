package blocks

import (
	//"math/big"
	"../utils"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type UnsignedBlockHeader struct {
	BlockNumber       uint32   `json:"blockNumber"`
	PreviousBlockHash [32]byte `json:"previousBlockHash"`
	MerkleRoot        [32]byte `json:"merkleRoot"`
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature [65]byte `json:"signature"`
}

func (b *Block) GetHash() ([]byte, error) {
	data, err := utils.EncodeToRLP(b.Header.UnsignedBlockHeader)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(data)
	if len(hash) != 32 {
		return nil, fmt.Errorf("wrong hash length %n, expected length: %n", len(hash), 32)
	}
	return hash, nil
}

func (b *Block) Sign(key []byte) error {
	hash, err := b.GetHash()
	if err != nil {
		return err
	}
	signature, err := utils.Sign(hash, key)
	if err != nil {
		return err
	}
	if len(signature) != 65 {
		return fmt.Errorf("wrong signature length %n, expected length: %n", len(signature), 65)
	}
	copy(b.Header.Signature[:], signature)
	return nil
}

func (b *Block) CalculateMerkleRoot() error {
	// b.Header.MerkleRoot =
	return nil
}

// === validation ===

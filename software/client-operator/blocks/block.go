package blocks

import (
	"../utils"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)
type Uint160 [20]byte
type Uint256 [32]byte
// RSA Accumulator
type Uint2048 [256]byte
// Signature 65 bytes long ECDSA signature encoded in RSV format
// R(32) bytes S(32) bytes  V(1) byte
type Signature [65]byte

type Block struct {
	Header       BlockHeader     `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type UnsignedBlockHeader struct {
	BlockNumber       uint32    `json:"blockNumber"`
	PreviousBlockHash Uint256 `json:"previousBlockHash"`
	MerkleRoot        Uint256 `json:"merkleRoot"`
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature Signature `json:"signature"`
}

type RSAInclusionProof struct {
	B Uint2048
	R Uint256
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
	return nil
}

// === validation ===

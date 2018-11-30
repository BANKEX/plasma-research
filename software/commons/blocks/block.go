package blocks

import (
	"../utils"
	a "./alias"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

// For test only
var Balance = map[string]int{"balance": 0}

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type UnsignedBlockHeader struct {
	BlockNumber      uint32            `json:"blockNumber"`
	TransactionsRoot SumMerkleNode     `json:"transactionsRoot"`
	RSAAccumulator   a.Uint2048        `json:"sumMerkleNode"`
	RSAChainProof    RSAInclusionProof `json:"RSAInclusionProof"`
}

const PlasmaRangeSpace = 2 ^ 24 - 1

type SumMerkleNode struct {
	// We use 24 bit
	Length uint32
	Hash   a.Uint160
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature a.Signature `json:"signature"`
}

type RSAInclusionProof struct {
	B a.Uint2048
	R a.Uint256
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

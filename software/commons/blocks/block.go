package blocks

import (
	"../utils"
	a "./alias"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"

	"../../plasmautils/plasmacrypto"
	"../../plasmautils/slice"
)

// For test only
var Balance = map[string]int{"balance": 0}

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type UnsignedBlockHeader struct {
	BlockNumber       uint32    `json:"blockNumber"`
	PreviousBlockHash a.Uint256 `json:"previousBlockHash"`
	MerkleRoot        a.Uint256 `json:"merkleRoot"`
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

// === utils ===

// UpdateRSAAccumulator adds input ranges for all submitted transactions to the RSA accumulator
// Algorithm complexity is O(N*logN) for N transactions
func UpdateRSAAccumulator(previous *big.Int, transactions []Transaction) *big.Int {
	acc := plasmacrypto.Accumulator{}.SetInt(previous)
	for _, t := range transactions {
		for _, i := range t.Inputs {
			s := slice.Slice{Begin: i.Amount.Begin, End: i.Amount.End}
			for _, p := range s.GetAlignedSlices() {
				acc.Accumulate(p)
			}
		}
	}
	return acc.Value()
}

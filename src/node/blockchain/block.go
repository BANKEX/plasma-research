package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"

	. "github.com/BANKEX/plasma-research/src/node/alias"
	"github.com/BANKEX/plasma-research/src/node/config"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/plasmacrypto"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/primeset"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/BANKEX/plasma-research/src/node/utils"
	"github.com/ethereum/go-ethereum/crypto"
)

// todo merkle proofs for transaction inclusion for client
// todo rsa proofs for transaction inclusion and exclusion

var WeiPerCoin uint = 1e11

// UnsignedBlockHeader is a structure that signature is calculated for.
type UnsignedBlockHeader struct {
	BlockNumber    uint32         `json:"blockNumber"`
	PreviousHash   Uint256        `json:"previousHash"`
	MerkleRoot     SumTreeRoot    `json:"merkleRoot"`
	RSAAccumulator Uint2048       `json:"rsaAccumulator"`
	hash           Uint256        // private variable because it should not be serialized
	merkleTree     *SumMerkleTree // private variable because it should not be serialized
}

// BlockHeader is a structure that gets sent to a smart contract.
type BlockHeader struct {
	UnsignedBlockHeader
	Signature Signature `json:"signature"`
}

// Block is a complete block that gets uploaded to public storage.
type Block struct {
	BlockHeader  `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

// NewBlock creates a block from previous block metadata and an array of transaction.
// This function will calculate merkle root, RSA accumulator, and sign the block.
func NewBlock(blockNumber uint32, previousHash Uint256, previousRSAAccumulator Uint2048, transactions []Transaction) (*Block, error) {
	block := Block{
		BlockHeader: BlockHeader{
			UnsignedBlockHeader: UnsignedBlockHeader{
				BlockNumber:  blockNumber,
				PreviousHash: previousHash,
			},
		},
		Transactions: transactions,
	}

	// todo enable RSA accumulator
	// block.UpdateRSAAccumulator(previousRSAAccumulator)

	err := block.CalculateMerkleRoot()
	if err != nil {
		return nil, err
	}

	err = block.calculateHash()
	if err != nil {
		return nil, err
	}

	key, err := hex.DecodeString(config.GetOperator().MainAccountPrivateKey[2:])
	if err != nil {
		return nil, err
	}

	err = block.Sign(key)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

// GetHash gets the hash of block header.
func (b *Block) GetHash() Uint256 {
	return b.hash
}

// calculateHash calculates the hash of block header.
func (b *Block) calculateHash() error {
	data, err := utils.EncodeToRLP(b.UnsignedBlockHeader)
	if err != nil {
		return err
	}
	hash := crypto.Keccak256(data)
	if len(hash) != 32 {
		return fmt.Errorf("wrong hash length %d, expected length: %d", len(hash), 32)
	}
	b.hash = hash
	return nil
}

// Sign signs the block with the specified private key.
func (b *Block) Sign(key []byte) error {
	signature, err := utils.Sign(b.GetHash(), key)
	if err != nil {
		return err
	}
	if len(signature) != 65 {
		return fmt.Errorf("wrong signature length %d, expected length: %d", len(signature), 65)
	}
	b.Signature = signature
	return nil
}

// CalculateMerkleRoot calculates merkle root for transactions in the block.
func (b *Block) CalculateMerkleRoot() error {
	leaves, err := PrepareLeaves(b.Transactions)
	if err != nil {
		return err
	}
	tree := NewSumMerkleTree(leaves, utils.Keccak160)
	b.merkleTree = tree
	b.MerkleRoot = b.merkleTree.GetRoot()
	return nil
}

// UpdateRSAAccumulator adds input ranges for all submitted transactions to the RSA accumulator.
// Algorithm complexity is O(N*logN) for N transactions.
// This function accepts previous accumulator as argument instead of mutating the block to avoid double invocation.
func (b *Block) UpdateRSAAccumulator(previous Uint2048) {
	acc := new(plasmacrypto.Accumulator).SetInt(new(big.Int).SetBytes(previous))
	for _, t := range b.Transactions {
		for _, i := range t.Inputs {
			s := slice.Slice{Begin: i.Slice.Begin, End: i.Slice.End}
			for _, p := range s.GetAlignedSlices() {
				acc.Accumulate(primeset.PrimeN(int(p)))
			}
		}
	}
	b.RSAAccumulator = acc.Value().Bytes()
}

func (b *Block) SerializeHeader() []byte {
	buf := new(bytes.Buffer)
	s := io.Writer(buf)
	_ = binary.Write(s, binary.LittleEndian, b.BlockNumber)
	buf.Write(b.PreviousHash)
	_ = binary.Write(s, binary.LittleEndian, b.MerkleRoot.Length)
	buf.Write(b.MerkleRoot.Hash)
	buf.Write(b.RSAAccumulator)
	return buf.Bytes()

	// result := make([]byte, 0, 4 + 32 + 4 + 20 + 256)
	// binary.LittleEndian.PutUint32(result, b.BlockNumber)
	// result = append(result, b.PreviousHash...)
	// binary.LittleEndian.PutUint32(result, b.MerkleRoot.Length)
	// result = append(result, b.MerkleRoot.Hash...)
	// result = append(result, b.RSAAccumulator...)
	// return result
}

func (b *Block) Serialize() ([]byte, error) {
	// todo
	return utils.EncodeToRLP(b)
}

func Deserialize(data []byte) *Block {
	// todo
	return nil
}

package blockchain

import (
	. "../alias"
	"../utils"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"

	"../../plasmautils/plasmacrypto"
	"../../plasmautils/primeset"
	"../../plasmautils/slice"
)

const PlasmaRangeSpace = 2 ^ 24 - 1

// For test only
var Balance = map[string]int{"balance": 0}
var PrivateKey []byte // todo move and init in config

// UnsignedBlockHeader is a structure that signature is calculated for.
type UnsignedBlockHeader struct {
	BlockNumber    uint32      `json:"blockNumber"`
	PreviousHash   Uint256     `json:"previousHash"`
	MerkleRoot     SumTreeNode `json:"merkleRoot"`
	RSAAccumulator Uint2048    `json:"rsaAccumulator"`
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

/*
// currently not used anywhere

type RSAInclusionProof struct {
	B Uint2048
	R Uint256
}
*/

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

	block.UpdateRSAAccumulator(previousRSAAccumulator)

	err := block.CalculateMerkleRoot()
	if err != nil {
		return nil, err
	}

	err = block.Sign(PrivateKey)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

// GetHash gets the hash of block header for signing.
func (b *Block) GetHash() ([]byte, error) {
	data, err := utils.EncodeToRLP(b.UnsignedBlockHeader)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(data)
	if len(hash) != 32 {
		return nil, fmt.Errorf("wrong hash length %n, expected length: %n", len(hash), 32)
	}
	return hash, nil
}

// Sign signs the block with the specified private key.
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
	b.Signature = signature
	return nil
}

// CalculateMerkleRoot calculates merkle root for transactions in the block.
func (b *Block) CalculateMerkleRoot() error {
	// todo
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

func (b *Block) Serialize() []byte {
	// todo
	return nil
}

func Deserialize(data []byte) *Block {
	// todo
	return nil
}

// todo merkle proofs for transaction inclusion for client
// todo rsa proofs for transaction inclusion and exclusion

func AssembleBlock(utxoPool UtxoPool, pendingTransactions []Transaction, txIndex TxIndex, privateKey []byte) (Block, UtxoPool) {
	block := Block{}

	for _, transaction := range HandleTxs(utxoPool, pendingTransactions, txIndex) {
		block.Transactions = append(block.Transactions, transaction)
	}

	leaves := PrepareLeaves(block.Transactions)
	tree := NewSumMerkleTree(leaves)

	RSAAccumulator := Uint2048{}
	// RSAInclusionProof := b.GetRSAInclusionProof(block.Transactions)

	headerContent := UnsignedBlockHeader{
		BlockNumber:    1, // TODO: current + 1
		PreviousHash:   []byte{0x0},
		MerkleRoot:     *tree.GetRoot(),
		RSAAccumulator: RSAAccumulator,
		// RSAChainProof:    RSAInclusionProof,
	}

	//block.BlockHeader
	/*
		signature := SingHeader(headerContent, privateKey)
	*/

	block.BlockHeader = BlockHeader{
		UnsignedBlockHeader: headerContent,
		// Signature:           signature,
	}

	block.Sign(privateKey)

	return block, utxoPool
}

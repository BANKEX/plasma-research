package blockchain

import (
	. "github.com/BANKEX/plasma-research/src/node/alias"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/BANKEX/plasma-research/src/node/utils"

	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	MaxInputs     = 3
	MaxOutputs    = 3
	MaxSignatures = 2
)

type UnsignedTransaction struct {
	Inputs   []Input  `json:"inputs"`
	Outputs  []Output `json:"outputs"`
	Metadata Metadata `json:"metadata"`
}

// Signatures may only contain one or two signatures
type Transaction struct {
	UnsignedTransaction
	// Signatures 65 bytes long ECDSA signature encoded in RSV format
	// R(32) bytes S(32) bytes  V(1) byte
	Signatures []Signature `json:"signatures"`
}

type Metadata struct {
	// MaxBlockNumber is a block number before the transaction should be included,
	// otherwise the transaction is considered invalid
	MaxBlockNumber uint32 `json:"maxBlockNumber"`
}

// Input represents transaction input in terms of UTXO model
// Input should refers to output of some previous transaction
// BlockIndex, TxIndex and OutputIndex helps to find out where that input are
type Input struct {
	// index of the block that contains corresponding output
	BlockIndex uint32 `json:"blockNumber"`
	// index of the transaction within the block
	TxIndex uint32 `json:"txNumber"`
	// index of the output within transaction
	OutputIndex uint8 `json:"outputNumber"`
	Output
}

// Output represents transaction output in terms of UTXO model
type Output struct {
	// Ethereum address of the owner
	Owner Uint160     `json:"owner"`
	Slice slice.Slice `json:"slice"`
}

// GetKey gets unique input hash consisting of block number, transaction number, and output number.
func (i *Input) GetKey() string {
	return fmt.Sprintf("%d:%d:%d", i.BlockIndex, i.TxIndex, i.OutputIndex)
}

// GetInputOwners gets array of unique input owner addresses.
func (t *UnsignedTransaction) GetInputOwners() [][]byte {
	// someone can do more optimal unique extraction
	result := make([][]byte, 0, len(t.Inputs))
	for _, in := range t.Inputs {
		if !utils.Contains(result, in.Owner) {
			result = append(result, in.Owner)
		}
	}
	return result
}

// GetMerkleRoot gets the root of merklized transaction inputs, outputs, and metadata.
func (t *UnsignedTransaction) GetMerkleRoot() Uint160 {
	var leaves []Item

	for _, data := range t.Inputs {
		rlpEncoded, _ := utils.EncodeToRLP(data)
		leaves = append(leaves, rlpEncoded)
	}

	for _, data := range t.Outputs {
		rlpEncoded, _ := utils.EncodeToRLP(data)
		leaves = append(leaves, rlpEncoded)
	}

	var rlpMetadata, _ = utils.EncodeToRLP(t.Metadata)
	leaves = append(leaves, rlpMetadata)

	tree := NewMerkleTree(leaves, 3, utils.Keccak160)

	return []byte(tree.GetRoot())
}

// Pads 20 byte hash to 32 bytes with zeros
func PadHash(in Uint160) Uint256 {
	return append(make([]byte, 12), in...)
}

// GetSignaturesHash returns a hash of concatenated signatures.
func (t *Transaction) GetSignaturesHash() Uint160 {
	result := make([]byte, 0, 65*len(t.Signatures))
	for _, s := range t.Signatures {
		result = append(result, s...)
	}
	return Uint160(utils.Keccak160(result))
}

// GetHash returns a full hash of signed transaction.
func (t *Transaction) GetHash() Uint160 {
	// TODO: We call GetHash many times - think refactor of transaction to kind of freezed object with hash that calculated once
	var result []byte
	result = append(result, t.GetMerkleRoot()...)
	result = append(result, t.GetSignaturesHash()...)
	return Uint160(utils.Keccak160(result))
}

// Signs a transaction with a specified private key.
// This function will append the generated signature to transactions' Signatures array.
func (t *Transaction) Sign(key []byte) error {
	privateKey, err := crypto.ToECDSA(key)
	if err != nil {
		return err
	}
	signature, err := crypto.Sign(PadHash(t.GetMerkleRoot()), privateKey)
	if err != nil {
		return err
	}
	if len(signature) != 65 {
		return fmt.Errorf("wrong signature length %d, expected length: %d", len(signature), 65)
	}
	t.Signatures = append(t.Signatures, signature)
	return nil
}

// todo serialization and deserialization

// todo maybe validation should be in a separate file
// === validation ===

// ValidateSoftLimits checks that soft array size limits are not exceeded.
func (t *Transaction) ValidateSoftLimits() error {
	if t.Inputs == nil || len(t.Inputs) > MaxInputs {
		return errors.New("wrong input count")
	}
	if t.Outputs == nil || len(t.Outputs) > MaxOutputs {
		return errors.New("wrong output count")
	}
	if t.Signatures == nil || len(t.Signatures) > MaxSignatures {
		return errors.New("wrong signature count")
	}
	return nil
}

func (t *Transaction) ValidateSlices() error {
	// TODO: check overlapping of slices e.g. double spend
	// TODO: check that slices are ordered in the correct way - e.g. we have same sequence after sorting
	// todo validate slices non-intersection
	// todo validate that input slices fully cover output slices
	return nil
}

// ValidateSignatures checks that the transaction is properly signed.
func (t *Transaction) ValidateSignatures() error {
	for _, owner := range t.GetInputOwners() {
		signed := false
		for _, sig := range t.Signatures {
			pubkey, err := crypto.Ecrecover(PadHash(t.GetMerkleRoot()), sig)
			if err != nil {
				return err
			}
			key, err := crypto.UnmarshalPubkey(pubkey)
			if err != nil {
				return err
			}
			addr := crypto.PubkeyToAddress(*key).Bytes()
			if bytes.Compare(owner, addr) == 0 {
				signed = true
				break
			}
		}
		if !signed {
			return fmt.Errorf("input belonging to owner %x is not properly signed", owner)
		}
	}
	return nil
}

// Validate checks that transaction is well formed
func (t *Transaction) Validate() error {
	// maybe could be written shorter
	if err := t.ValidateSoftLimits(); err != nil {
		return err
	}
	if err := t.ValidateSlices(); err != nil {
		return err
	}
	if err := t.ValidateSignatures(); err != nil {
		return err
	}
	return nil
}

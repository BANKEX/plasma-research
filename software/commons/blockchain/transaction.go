package blockchain

import (
	"../../plasmautils/slice"
	. "../alias"
	"../rlp"
	"../utils"
	"errors"
	"fmt"
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

type Transaction struct {
	UnsignedTransaction
	Signatures []Signature `json:"signatures"`
}

type Metadata struct {
	// MaxBlockNumber is a block number before the transaction should be included,
	// otherwise the transaction is considered invalid
	MaxBlockNumber uint32 `json:"maxBlockNumber"`
}

type Input struct {
	BlockIndex  uint32 `json:"blockNumber"`
	TxIndex     uint32 `json:"txNumber"`
	OutputIndex uint8  `json:"outputNumber"`
	Output
}

type Output struct {
	Owner Uint160     `json:"owner"`
	Slice slice.Slice `json:"slice"`
}

// GetMerkleRoot gets the root of merklized transaction inputs, outputs, and metadata.
func (ut *UnsignedTransaction) GetMerkleRoot() Uint160 {
	var leafs []utils.Item

	for _, data := range ut.Inputs {
		rlpEncoded, _ := rlp.EncodeToRLP(data)
		leafs = append(leafs, rlpEncoded)
	}

	for _, data := range ut.Outputs {
		rlpEncoded, _ := rlp.EncodeToRLP(data)
		leafs = append(leafs, rlpEncoded)
	}

	var rlpMetadata, _ = rlp.EncodeToRLP(ut.Metadata)
	leafs = append(leafs, rlpMetadata)

	tree := utils.NewMerkleTree(leafs, 3, utils.Keccak160)

	return []byte(tree.GetRoot())
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
	var result []byte
	result = append(result, t.GetMerkleRoot()...)
	result = append(result, t.GetSignaturesHash()...)
	return Uint160(utils.Keccak160(result))
}

// Signs a transaction with a specified private key.
// This function will append the generated signature to transactions' Signatures array
func (t *Transaction) Sign(key []byte) error {
	hash := t.GetHash()
	signature, err := utils.Sign(hash, key)
	if err != nil {
		return err
	}
	if len(signature) != 65 {
		return fmt.Errorf("wrong signature length %n, expected length: %n", len(signature), 65)
	}
	t.Signatures = append(t.Signatures, signature)
	return nil
}

// todo serialization and deserialization

// === validation ===

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
	return nil
}

func (t *Transaction) ValidateOutputSum() error {
	return nil
}

func (t *Transaction) ValidateSignatures() error {
	return nil
}

// todo this should receive fee arguments from the outside, needed only for operator
func (t *Transaction) ValidateFee() error {
	return nil
}

// todo validate slices non-intersection
func (t *Transaction) Validate() error {
	return t.ValidateSoftLimits() // || tr.ValidateOutputSum() || tr.ValidateSignatures() || tr.ValidateFee() || nil
}

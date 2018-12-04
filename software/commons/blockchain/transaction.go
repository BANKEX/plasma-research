package blockchain

import (
	"../rlp"
	"../utils"
	a "./alias"
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

type Transaction struct {
	UnsignedTransaction
	Signatures []a.Signature `json:"signatures"`
}

type Metadata struct {
	MaxBlockNumber uint32 `json:"maxBlockNumber"`
}

type Input struct {
	Owner       a.Uint160 `json:"owner"`
	BlockIndex  uint32    `json:"blockNumber"`
	TxIndex     uint32    `json:"txNumber"`
	OutputIndex uint8     `json:"outputNumber"`
	//AssetID     uint256  `json:"assetId"`
	Amount Segment `json:"amount"`

	// TODO: optimize probably we don't need it
	// TxHash of previous transaction
	PrevTxHash a.Uint160
}

type Output struct {
	Owner a.Uint160 `json:"owner"`
	//AssetID uint256 `json:"assetId"`
	Amount Segment `json:"amount"`
}

type Segment struct {
	Begin uint32 // Index of one of 2^24 coins - starts from 0
	End   uint32
}

func (ut *UnsignedTransaction) GetMerkleRoot() a.Uint160 {

	var leafs []utils.Item

	for _, data := range ut.Outputs {
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

	var result [20]byte
	copy(result[:], tree.GetRoot())

	return result
}

func concat(values ...[]byte) []byte {
	var buffer bytes.Buffer
	for _, s := range values {
		buffer.Write(s)
	}
	return buffer.Bytes()
}

func getSignaturesHash(t *Transaction) []byte {
	if len(t.Signatures) == 1 {
		b := (t.Signatures[0])[:]
		return utils.Keccak160(b)
	}

	b1 := (t.Signatures[0])[:]
	b2 := (t.Signatures[0])[:]
	return utils.Keccak160(concat(b1, b2))
}

func (t *Transaction) GetWTFHash() (a.Uint160, error) {

	contentRoot := t.UnsignedTransaction.GetMerkleRoot()
	rootData := concat(contentRoot[:], getSignaturesHash(t))
	rootHash := utils.Keccak160(rootData)

	result := a.Uint160{}
	copy(result[:], rootHash)

	return result, nil
}

func (t *Transaction) GetHash() ([]byte, error) {
	data, err := utils.EncodeToRLP(t.UnsignedTransaction)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(data)
	if len(hash) != 32 {
		return nil, fmt.Errorf("wrong hash length %n, expected length: %n", len(hash), 32)
	}
	return hash, nil
}

func (t *Transaction) Sign(key []byte) error {
	hash, err := t.GetHash()
	if err != nil {
		return err
	}
	signature, err := utils.Sign(hash[:], key)
	if err != nil {
		return err
	}
	if len(signature) != 65 {
		return fmt.Errorf("wrong signature length %n, expected length: %n", len(signature), 65)
	}
	copy(t.Signatures[0][:], signature)
	return nil
}

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

func (t *Transaction) ValidateOutputSum() error {
	return nil
}

func (t *Transaction) ValidateSignatures() error {
	return nil
}

// todo this should receive fee arguments from outside, needed only for operator
func (t *Transaction) ValidateFee() error {
	return nil
}

func (t *Transaction) Validate() error {
	return t.ValidateSoftLimits() // || tr.ValidateOutputSum() || tr.ValidateSignatures() || tr.ValidateFee() || nil
}

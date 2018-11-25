package blocks

import (
	// "math/big"
	"../utils"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type UnsignedTransaction struct {
	Inputs   []Input  `json:"inputs"`
	Outputs  []Output `json:"outputs"`
	Metadata Metadata `json:"metadata"`
}

type Transaction struct {
	UnsignedTransaction
	Signatures [][65]byte `json:"signatures"`
}

type Metadata struct {
	MaxBlockNumber uint32 `json:"maxBlockNumber"`
}

func (tr *Transaction) GetHash() ([]byte, error) {
	data, err := utils.EncodeToRLP(tr.UnsignedTransaction)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256(data)
	if len(hash) != 32 {
		return nil, fmt.Errorf("wrong hash length %n, expected length: %n", len(hash), 32)
	}
	return hash, nil
}

func (tr *Transaction) Sign(key []byte) error {
	hash, err := tr.GetHash()
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
	copy(tr.Signatures[0][:], signature)
	return nil
}

// === validation ===

func (tr *Transaction) ValidateSoftLimits() error {
	if tr.Inputs == nil || len(tr.Inputs) > MaxInputs {
		return errors.New("wrong input count")
	}
	if tr.Outputs == nil || len(tr.Outputs) > MaxOutputs {
		return errors.New("wrong output count")
	}
	if tr.Signatures == nil || len(tr.Signatures) > MaxSignatures {
		return errors.New("wrong signature count")
	}
	return nil
}

func (tr *Transaction) ValidateOutputSum() error {
	return nil
}

func (tr *Transaction) ValidateSignatures() error {
	return nil
}

// todo this should receive fee arguments from outside, needed only for operator
func (tr *Transaction) ValidateFee() error {
	return nil
}

func (tr *Transaction) Validate() error {
	return tr.ValidateSoftLimits() // || tr.ValidateOutputSum() || tr.ValidateSignatures() || tr.ValidateFee() || nil
}

package blocks

import (
// "math/big"
	"errors"
)

type Transaction struct {
	Inputs     []Input    `json:"inputs"`  // maybe we need input1, input2 ...
	Outputs    []Output   `json:"outputs"` // maybe we need output1, output2 ...
	Metadata   Metadata   `json:metadata`
	Signatures [][64]byte `json:"signatures"` // maybe we need signature1, ...
}

type Metadata struct {
	MaxBlockNumber uint32 `json:"maxBlockNumber"`
}

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

func (tr *Transaction) Sign(key []byte) error {
	return nil
}

// === validation ===

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
	return tr.ValidateSoftLimits()// || tr.ValidateOutputSum() || tr.ValidateSignatures() || tr.ValidateFee() || nil
}
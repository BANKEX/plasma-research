// Copyright 2018 The plasma-research Authors
// This file is part of the plasma-research library.
// It's unstable. Currently is  not production ready. Structure will be changed potentially.

package types

import (
	"math/big"
	"crypto/rand"
)

type Block struct {
	BlockHeader  *BlockHeader   `json:"blockHeader"`
	Transactions *[]Transaction `json:"transactions"`
}

type BlockHeader struct {
	BlockNumber       uint32         `json:"blockNumber"`
	PreviousBlockHash big.Int        `json:"previousBlockHash"`
	MerkleRoot        big.Int        `json:"merkleRoot"`
	Signature         *Signature     `json:"signature"`
	Transactions      *[]Transaction `json:"transactions"`
}

type Transaction struct {
	Input      *[]Input     `json:"inputs"`  // maybe we need input1, input2 ...
	Output     *[]Output    `json:"outputs"` // maybe we need output1, output2 ...
	Metadata   *Metadata    `json:"metadata"`
	Signatures *[]Signature `json:"signatures"` // maybe we need signature1, ...
}

type Input struct {
	Owner       big.Int `json:"owner"`
	BlockIndex  uint32  `json:"blockIndex"`
	TxIndex     uint32  `json:"txIndex"`
	OutputIndex uint8   `json:"outputIndex"`
	AssetId     big.Int `json:"assetId"`
	Amount      big.Int `json:"amount"`
}

type Output struct {
	Owner   big.Int `json:"owner"`
	AssetId big.Int `json:"assetId"`
	Amount  big.Int `json:"amount"`
}

type Signature struct {
	Signature [64]byte `json:"signature"`
}

type Metadata struct {
	Max_block_id uint32 `json:"max_block_id"`
}

func BlockInit() Block {

	metadata := Metadata{}
	metadata.Max_block_id = 32

	signature := Signature{}
	slice := []byte("abcdefgh")
	var arr [64]byte
	copy(arr[:], slice[:4])
	signature.Signature = arr

	signatures := make([]Signature, 1)
	signatures = append(signatures, signature)

	output := Output{}
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		//error handling
	}
	output.Owner = *n
	output.AssetId = *n
	output.Amount = *n

	outputs := make([]Output, 1)
	outputs = append(outputs, output)
	outputs = append(outputs, output)

	input := Input{}
	input.Amount = *n
	input.AssetId = *n
	input.Owner = *n
	input.OutputIndex = 0
	input.TxIndex = 0
	input.BlockIndex = 0

	inputs := make([]Input, 1)
	inputs = append(inputs, input)
	inputs = append(inputs, input)


	transaction := Transaction{}
	transaction.Input = &inputs
	transaction.Output = &outputs
	transaction.Signatures = &signatures
	transaction.Metadata = &metadata

	transactions := make([]Transaction, 1)
	transactions = append(transactions, transaction)
	transactions = append(transactions, transaction)

	header := BlockHeader{}
	header.Signature = &signature
	header.Transactions = &transactions
	header.PreviousBlockHash = *n
	header.MerkleRoot = *n
	header.BlockNumber = 0

	block := Block{}
	block.Transactions = &transactions
	block.BlockHeader = &header

	return block
}

// Copyright 2018 The plasma-research Authors
// This file is part of the plasma-research library.
// It's unstable. Currently is  not production ready. Structure will be changed potentially.
package types


import (
	"crypto/rand"
	"math/big"
	// secp256k1 "github.com/matterinc/PlasmaCommons/crypto/secp256k1"
)

func NewBlock() Block {

	metadata := Metadata{}
	metadata.Max_block_id = 0

	signature := Signature{}
	slice := []byte("abcdefgh")
	var arr [65]byte
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

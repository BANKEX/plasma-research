// Copyright 2018 The plasma-research Authors
// This file is part of the plasma-research library.
// It's unstable. Currently is  not production ready. Structure will be changed potentially.

package types

import (
	"math/big"
)

type Block struct {
	BlockHeader  *BlockHeader   `json:"blockHeader"`
	Transactions *[]Transaction `json:"transactions"`
}

type BlockHeader struct {
	BlockNumber       uint32         `json:"blockNumber"`
	PreviousBlockHash big.Int        `json:"previousBlockHash"`
	MerkleRoot        big.Int        `json:"merkleRoot"`
	Signature         *signatue      `json:"signature"`
	Transactions      *[]Transaction `json:"transactions"`
}

type Transaction struct {
	Input      *[5]Input    `json:"inputs"`     // maybe we need input1, input2 ...
	Output     *[5]Output   `json:"outputs"`    // maybe we need output1, output2 ...
	Signatures *[1]signatue `json:"signatures"` // maybe we need signature1, ...
}

type Input struct {
	Owner       [159]big.Int `json:"owner"`
	BlockIndex  uint32       `json:"blockIndex"`
	TxIndex     uint32       `json:"txIndex"`
	OutputIndex uint8        `json:"outputIndex"`
	AssetId     [159]big.Int `json:"assetId"`
	Amount      big.Int      `json:"amount"`
}

type Output struct {
	Owner   [159]big.Int `json:"owner"`
	AssetId [159]big.Int `json:"assetId"`
	Amount  big.Int      `json:"amount"`
}

type signatue struct {
	Signature [64]byte `json:"signature"`
}

type Metadata struct {
	Max_block_id uint32 `json:"max_block_id"`
}


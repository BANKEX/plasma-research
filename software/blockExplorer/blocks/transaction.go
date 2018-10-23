package blocks

import (
// "math/big"
)

type Transaction struct {
	Input      []Input    `json:"inputs"`  // maybe we need input1, input2 ...
	Output     []Output   `json:"outputs"` // maybe we need output1, output2 ...
	Metadata   Metadata   `json:metadata`
	Signatures [][64]byte `json:"signatures"` // maybe we need signature1, ...
}

type Metadata struct {
	MaxBlockNum uint32 `json:"maxBlockNum"`
}

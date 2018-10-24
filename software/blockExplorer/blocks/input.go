package blocks

import (
	"math/big"
)

type Input struct {
	BlockNumber  uint32   `json:"blockNumber"`
	TxNumber     uint32   `json:"txNumber"`
	OutputNumber uint8    `json:"outputNumber"`
	Owner        [20]byte `json:"owner"`
	AssetID      [32]byte `json:"assetId"`
	Amount       big.Int  `json:"amount"`
}

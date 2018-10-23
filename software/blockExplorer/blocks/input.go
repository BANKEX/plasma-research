package blocks

import (
	"math/big"
)

type Input struct {
	BlockNumber  uint32  `json:"blockNumber"`
	TxNumber     uint32  `json:"txNumber"`
	OutputNumber uint8   `json:"outputNumber"`
	Owner        big.Int `json:"owner"`
	AssetID      big.Int `json:"assetId"`
	Amount       big.Int `json:"amount"`
}

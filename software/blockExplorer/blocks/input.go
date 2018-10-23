package blocks

import (
	"math/big"
)

type Input struct {
	BlockIndex  uint32  `json:"blockIndex"`
	TxIndex     uint32  `json:"txIndex"`
	OutputIndex uint8   `json:"outputIndex"`
	Owner       big.Int `json:"owner"`
	AssetId     big.Int `json:"assetId"`
	Amount      big.Int `json:"amount"`
}

package blocks

import (
	"math/big"
)

type Output struct {
	Owner   [20]byte `json:"owner"`
	AssetID [32]byte `json:"assetId"`
	Amount  big.Int  `json:"amount"`
}

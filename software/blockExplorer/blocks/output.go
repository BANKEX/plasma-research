package blocks

import (
	"math/big"
)

type Output struct {
	Owner   big.Int `json:"owner"`
	AssetId big.Int `json:"assetId"`
	Amount  big.Int `json:"amount"`
}

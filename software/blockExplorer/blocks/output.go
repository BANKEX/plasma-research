package blocks

import (
	"math/big"
)

type Output struct {
	Owner   [159]big.Int `json:"owner"`
	AssetId [159]big.Int `json:"assetId"`
	Amount  big.Int      `json:"amount"`
}

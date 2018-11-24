package blocks

type Output struct {
	Owner uint160 `json:"owner"`
	//AssetID uint256 `json:"assetId"`
	Amount Segment `json:"amount"`
}

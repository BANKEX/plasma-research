package blocks

type Input struct {
	Owner       uint160 `json:"owner"`
	BlockIndex  uint32  `json:"blockNumber"`
	TxIndex     uint32  `json:"txNumber"`
	OutputIndex uint8   `json:"outputNumber"`
	//AssetID     uint256  `json:"assetId"`
	Amount Segment `json:"amount"`
}

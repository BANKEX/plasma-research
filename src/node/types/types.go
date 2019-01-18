package types

type LastBlock struct {
	LastBlock string `json:"lastBlock"`
}

type OperatorInfo struct {
	Config interface{}
	ABI    string
}

type InputResponse struct {
	// index of the block that contains corresponding output
	BlockIndex uint32 `json:"blockNumber"`
	// index of the transaction within the block
	TxIndex uint32 `json:"txNumber"`
	// index of the output within transaction
	OutputIndex    uint8 `json:"outputNumber"`
	OutputResponse `json:"output"`
}

// Output represents transaction output in terms of UTXO model
type OutputResponse struct {
	// Ethereum address of the owner
	Owner string         `json:"owner"`
	Slice SliceRespopnse `json:"slice"`
}

type SliceRespopnse struct {
	Begin uint32 `json:"begin"`
	End   uint32 `json:"end"`
}

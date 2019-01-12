package types

type LastBlock struct {
	LastBlock string `json:"lastBlock"`
}

type OperatorInfo struct {
	Config interface{}
	ABI    string
}

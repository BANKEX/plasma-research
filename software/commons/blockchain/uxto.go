package blockchain

import a "../alias"

type UTXO struct {
	TxHash      a.TxHashBytes
	OutputIndex uint8
}

type UtxoPool map[UTXO]Output

// Shallow copy of UTXO pool
func (src *UtxoPool) GetCopy() *UtxoPool {
	var poolCopy UtxoPool
	for k, v := range *src {
		poolCopy[k] = v
	}
	return &poolCopy
}

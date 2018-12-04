package pool

import "../../commons/blockchain"
import a "../../commons/alias"

type UTXO struct {
	TxHash      a.Uint160
	OutputIndex uint8
}

type UtxoPool map[UTXO]blockchain.Block

// Shallow copy of UTXO pool
func (src *UtxoPool) GetCopy() *UtxoPool {
	var poolCopy UtxoPool
	for k, v := range *src {
		poolCopy[k] = v
	}
	return &poolCopy
}

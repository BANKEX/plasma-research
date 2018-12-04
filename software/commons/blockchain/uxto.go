package blockchain

import a "./alias"

type UtxoPool map[UTXO]Output

type UTXO struct {
	TxHash      a.Uint160
	OutputIndex uint8
}

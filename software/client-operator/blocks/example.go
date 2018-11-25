package blocks

import (
	"math/big"
)

func CreateOutput() Output {
	return Output{
		Owner:   [20]byte{},
		AssetID: [32]byte{},
		Amount:  *big.NewInt(1337),
	}
}

func CreateInput() Input {
	return Input{
		BlockNumber:  1,
		TxNumber:     2,
		OutputNumber: 3,
		Owner:        [20]byte{},
		AssetID:      [32]byte{},
		Amount:       *big.NewInt(1337),
	}
}

func CreateTransaction() Transaction {
	return Transaction{
		UnsignedTransaction: UnsignedTransaction{
			Inputs: []Input{
				CreateInput(),
				CreateInput(),
			},
			Outputs: []Output{
				CreateOutput(),
			},
			Metadata: Metadata{
				MaxBlockNumber: 10,
			},
		},
	}
}

func CreateBlock() Block {
	return Block{
		Header: BlockHeader{
			UnsignedBlockHeader: UnsignedBlockHeader{
				BlockNumber:       1,
				PreviousBlockHash: [32]byte{},
				MerkleRoot:        [32]byte{},
			},
		},
		Transactions: []Transaction{
			CreateTransaction(),
		},
	}
}

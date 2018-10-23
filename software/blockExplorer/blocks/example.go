package blocks

import (
	"math/big"
)

func CreateOutput() Output {
	return Output{
		Owner:   *big.NewInt(101),
		AssetID: *big.NewInt(42),
		Amount:  *big.NewInt(1337),
	}
}

func CreateInput() Input {
	return Input{
		BlockNumber:  1,
		TxNumber:     2,
		OutputNumber: 3,
		Owner:        *big.NewInt(101),
		AssetID:      *big.NewInt(42),
		Amount:       *big.NewInt(1337),
	}
}

func CreateTransaction() Transaction {
	return Transaction{
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
	}
}

func CreateBlock() Block {
	return Block{
		Header: BlockHeader{
			BlockNumber:       1,
			PreviousBlockHash: *big.NewInt(101),
			MerkleRoot:        *big.NewInt(101),
		},
		Transactions: []Transaction{
			CreateTransaction(),
		},
	}
}

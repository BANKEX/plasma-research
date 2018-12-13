package transactionManager

import (
	"log"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Event struct {
	Signature string
	Handler   func(data types.Log, abi abi.ABI)
}

var Handlers = []Event{
	{"AssetDeposited(address,address,uint64,uint64,uint64)", HandleDeposit},
}

var Manager *TransactionManager

type EventAssetDeposited struct {
	Who         common.Address
	IntervalId  uint64
	Begin       uint64
	End         uint64
	BlockNumber uint64
}

func HandleDeposit(data types.Log, abi abi.ABI) {
	var depositEvent EventAssetDeposited
	err := abi.Unpack(&depositEvent, "AssetDeposited", data.Data)
	if err != nil {
		log.Fatal(err)
	}
	who := data.Topics[2].Bytes()

	out := blockchain.Output{
		Owner: who,
		Slice: slice.Slice{
			Begin: uint32(depositEvent.Begin),
			End:   uint32(depositEvent.End),
		},
	}
	_, err = Manager.AssembleDepositBlock(out)
	if err != nil {
		log.Fatal(err)
	}
}

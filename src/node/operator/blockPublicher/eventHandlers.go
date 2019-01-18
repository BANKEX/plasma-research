package blockPublicher

import (
	"log"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/slice"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

var manager *TransactionManager
var publisher *BlockPublisher

type Event struct {
	Signature string
	Handler   func(data types.Log, abi abi.ABI)
}

var Handlers = []Event{
	{"AssetDeposited(address,address,uint64,uint64,uint64)", HandleDeposit},
	// {"WithdrawalBegin(address,uint32,uint32,uint8,address,uint64,uint64)", HandleDeposit},
}

func HandleDeposit(data types.Log, abi abi.ABI) {
	var depositEvent struct {
		IntervalId  uint64
		Begin       uint64
		End         uint64
		BlockNumber uint64
	}
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
	block, err := manager.AssembleDepositBlock(out)
	if err != nil {
		log.Fatal(err)
	}
	err = publisher.PublishBlock(block)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Handled deposit of %d coins to 0x%x", out.Slice.End-out.Slice.Begin, who)
}

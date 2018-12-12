package eventHandlers

import (
	"../../blockchain"
	"../../plasmautils/slice"
	"../../transactionManager"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

type Event struct {
	Signature string
	Handler   func(data types.Log, abi abi.ABI)
}

var Handlers = []Event{
	{"AssetDeposited(address,address,uint64,uint64,uint64)", HandleDeposit},
}

var Manager *transactionManager.TransactionManager

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
	depositEvent.Who = common.HexToAddress(data.Topics[2].Hex())
	depositEvent.BlockNumber = data.BlockNumber

	out := blockchain.Output{
		Owner: depositEvent.Who.Bytes(),
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

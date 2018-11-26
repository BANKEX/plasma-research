package storage

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client *ethclient.Client

	Addr string

	Balance string

	Who common.Address

	Amount *big.Int

	EventBlockHash string

	EventBlockNumber uint64

	StateForEvent = 0
)

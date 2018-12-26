package deposit

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/BANKEX/plasma-research/src/node/ethereum/plasmacontract"
	"github.com/BANKEX/plasma-research/src/node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const gasLimitForDeposit = uint64(300000)

func Deposit(ctx context.Context, contractAddress common.Address, client *ethclient.Client, key *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error) {
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts := utils.GetTxOpts(ctx, key, gasLimitForDeposit, gasPrice)
	opts.Value = value

	tx, err := instance.Deposit(opts)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

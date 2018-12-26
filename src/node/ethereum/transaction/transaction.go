package transaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const defaultGasLimitForTransfer = uint64(21000)

func SendTransactionInWei(ctx context.Context, client *ethclient.Client, key *ecdsa.PrivateKey, value *big.Int, to common.Address) (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, value, defaultGasLimitForTransfer, gasPrice, data)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, err
}

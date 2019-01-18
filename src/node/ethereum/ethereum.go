package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/BANKEX/plasma-research/src/contracts/api"
	"github.com/BANKEX/plasma-research/src/node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	defaultGasLimitForTransfer = uint64(21000)
	defaultGasLimit            = uint64(300000)
)

type Ethereum struct {
	client                *ethclient.Client
	plasmaContractAddress common.Address
	plasmaContract        *api.BankexPlasma
}

func NewEthereum(endpoint string, plasmaContractAddress common.Address) (*Ethereum, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	contract, err := api.NewBankexPlasma(plasmaContractAddress, client)
	if err != nil {
		return nil, err
	}

	return &Ethereum{
		client:                client,
		plasmaContractAddress: plasmaContractAddress,
		plasmaContract:        contract,
	}, nil
}

func (e *Ethereum) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return e.client.PendingBalanceAt(ctx, address)
}

func (e *Ethereum) PushHashBlock(ctx context.Context, key *ecdsa.PrivateKey, blockNumber uint32, hash []byte) (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts := utils.GetTxOpts(ctx, key, defaultGasLimit, gasPrice)
	opts.Value = big.NewInt(0)
	opts.Nonce = big.NewInt(0).SetUint64(nonce)

	// _, err = instance.SubmitBlocks(blockNumber, hash) // TODO: uncomment after regenerating abi
	tx, err := e.plasmaContract.SubmitBlocks(opts, nil, nil) // TODO: normal params
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (e *Ethereum) GetLastBlockNumber() (*big.Int, error) {
	return e.plasmaContract.BlocksLength(nil)
}

func (e *Ethereum) Exit(ctx context.Context, key *ecdsa.PrivateKey) (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts := utils.GetTxOpts(ctx, key, defaultGasLimit, gasPrice)
	opts.Value = big.NewInt(0)
	opts.Nonce = big.NewInt(0).SetUint64(nonce)

	tx, err := e.plasmaContract.WithdrawalBegin(opts, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (e *Ethereum) Deposit(ctx context.Context, key *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error) {
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts := utils.GetTxOpts(ctx, key, defaultGasLimit, gasPrice)
	opts.Value = value

	return e.plasmaContract.Deposit(opts)
}

func (e *Ethereum) DepositERC20(ctx context.Context, key *ecdsa.PrivateKey, tokenAddress common.Address, value *big.Int) (*types.Transaction, error) {
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts := utils.GetTxOpts(ctx, key, defaultGasLimit, gasPrice)
	return e.plasmaContract.DepositERC20(opts, tokenAddress, value)
}

func (e *Ethereum) SendTransactionInWei(ctx context.Context, key *ecdsa.PrivateKey, value *big.Int, to common.Address) (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := e.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, value, defaultGasLimitForTransfer, gasPrice, data)

	chainID, err := e.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	if err != nil {
		return nil, err
	}

	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, err
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SumMerkleProofABI is the input ABI used to generate the binding from.
const SumMerkleProofABI = "[]"

// SumMerkleProofBin is the compiled bytecode used for deploying new contracts.
const SumMerkleProofBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820b8c9e4535fa84ea61bc0f3c343a524b323f8c12caa37330b2b4028252d86f7d90029`

// DeploySumMerkleProof deploys a new Ethereum contract, binding an instance of SumMerkleProof to it.
func DeploySumMerkleProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SumMerkleProof, error) {
	parsed, err := abi.JSON(strings.NewReader(SumMerkleProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SumMerkleProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SumMerkleProof{SumMerkleProofCaller: SumMerkleProofCaller{contract: contract}, SumMerkleProofTransactor: SumMerkleProofTransactor{contract: contract}, SumMerkleProofFilterer: SumMerkleProofFilterer{contract: contract}}, nil
}

// SumMerkleProof is an auto generated Go binding around an Ethereum contract.
type SumMerkleProof struct {
	SumMerkleProofCaller     // Read-only binding to the contract
	SumMerkleProofTransactor // Write-only binding to the contract
	SumMerkleProofFilterer   // Log filterer for contract events
}

// SumMerkleProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type SumMerkleProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SumMerkleProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SumMerkleProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SumMerkleProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SumMerkleProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SumMerkleProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SumMerkleProofSession struct {
	Contract     *SumMerkleProof   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SumMerkleProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SumMerkleProofCallerSession struct {
	Contract *SumMerkleProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SumMerkleProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SumMerkleProofTransactorSession struct {
	Contract     *SumMerkleProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SumMerkleProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type SumMerkleProofRaw struct {
	Contract *SumMerkleProof // Generic contract binding to access the raw methods on
}

// SumMerkleProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SumMerkleProofCallerRaw struct {
	Contract *SumMerkleProofCaller // Generic read-only contract binding to access the raw methods on
}

// SumMerkleProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SumMerkleProofTransactorRaw struct {
	Contract *SumMerkleProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSumMerkleProof creates a new instance of SumMerkleProof, bound to a specific deployed contract.
func NewSumMerkleProof(address common.Address, backend bind.ContractBackend) (*SumMerkleProof, error) {
	contract, err := bindSumMerkleProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProof{SumMerkleProofCaller: SumMerkleProofCaller{contract: contract}, SumMerkleProofTransactor: SumMerkleProofTransactor{contract: contract}, SumMerkleProofFilterer: SumMerkleProofFilterer{contract: contract}}, nil
}

// NewSumMerkleProofCaller creates a new read-only instance of SumMerkleProof, bound to a specific deployed contract.
func NewSumMerkleProofCaller(address common.Address, caller bind.ContractCaller) (*SumMerkleProofCaller, error) {
	contract, err := bindSumMerkleProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofCaller{contract: contract}, nil
}

// NewSumMerkleProofTransactor creates a new write-only instance of SumMerkleProof, bound to a specific deployed contract.
func NewSumMerkleProofTransactor(address common.Address, transactor bind.ContractTransactor) (*SumMerkleProofTransactor, error) {
	contract, err := bindSumMerkleProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofTransactor{contract: contract}, nil
}

// NewSumMerkleProofFilterer creates a new log filterer instance of SumMerkleProof, bound to a specific deployed contract.
func NewSumMerkleProofFilterer(address common.Address, filterer bind.ContractFilterer) (*SumMerkleProofFilterer, error) {
	contract, err := bindSumMerkleProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofFilterer{contract: contract}, nil
}

// bindSumMerkleProof binds a generic wrapper to an already deployed contract.
func bindSumMerkleProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SumMerkleProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SumMerkleProof *SumMerkleProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SumMerkleProof.Contract.SumMerkleProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SumMerkleProof *SumMerkleProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SumMerkleProof.Contract.SumMerkleProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SumMerkleProof *SumMerkleProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SumMerkleProof.Contract.SumMerkleProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SumMerkleProof *SumMerkleProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SumMerkleProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SumMerkleProof *SumMerkleProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SumMerkleProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SumMerkleProof *SumMerkleProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SumMerkleProof.Contract.contract.Transact(opts, method, params...)
}

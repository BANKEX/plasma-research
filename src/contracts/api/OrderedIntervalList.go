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

// OrderedIntervalListABI is the input ABI used to generate the binding from.
const OrderedIntervalListABI = "[]"

// OrderedIntervalListBin is the compiled bytecode used for deploying new contracts.
const OrderedIntervalListBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582058cf379029e79c0a38b562ee3d91f7e8076d6f560b0d72cec5302fd79daae8f90029`

// DeployOrderedIntervalList deploys a new Ethereum contract, binding an instance of OrderedIntervalList to it.
func DeployOrderedIntervalList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrderedIntervalList, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderedIntervalListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderedIntervalListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderedIntervalList{OrderedIntervalListCaller: OrderedIntervalListCaller{contract: contract}, OrderedIntervalListTransactor: OrderedIntervalListTransactor{contract: contract}, OrderedIntervalListFilterer: OrderedIntervalListFilterer{contract: contract}}, nil
}

// OrderedIntervalList is an auto generated Go binding around an Ethereum contract.
type OrderedIntervalList struct {
	OrderedIntervalListCaller     // Read-only binding to the contract
	OrderedIntervalListTransactor // Write-only binding to the contract
	OrderedIntervalListFilterer   // Log filterer for contract events
}

// OrderedIntervalListCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderedIntervalListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderedIntervalListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderedIntervalListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderedIntervalListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderedIntervalListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderedIntervalListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderedIntervalListSession struct {
	Contract     *OrderedIntervalList // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderedIntervalListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderedIntervalListCallerSession struct {
	Contract *OrderedIntervalListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OrderedIntervalListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderedIntervalListTransactorSession struct {
	Contract     *OrderedIntervalListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OrderedIntervalListRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderedIntervalListRaw struct {
	Contract *OrderedIntervalList // Generic contract binding to access the raw methods on
}

// OrderedIntervalListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderedIntervalListCallerRaw struct {
	Contract *OrderedIntervalListCaller // Generic read-only contract binding to access the raw methods on
}

// OrderedIntervalListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderedIntervalListTransactorRaw struct {
	Contract *OrderedIntervalListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderedIntervalList creates a new instance of OrderedIntervalList, bound to a specific deployed contract.
func NewOrderedIntervalList(address common.Address, backend bind.ContractBackend) (*OrderedIntervalList, error) {
	contract, err := bindOrderedIntervalList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalList{OrderedIntervalListCaller: OrderedIntervalListCaller{contract: contract}, OrderedIntervalListTransactor: OrderedIntervalListTransactor{contract: contract}, OrderedIntervalListFilterer: OrderedIntervalListFilterer{contract: contract}}, nil
}

// NewOrderedIntervalListCaller creates a new read-only instance of OrderedIntervalList, bound to a specific deployed contract.
func NewOrderedIntervalListCaller(address common.Address, caller bind.ContractCaller) (*OrderedIntervalListCaller, error) {
	contract, err := bindOrderedIntervalList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListCaller{contract: contract}, nil
}

// NewOrderedIntervalListTransactor creates a new write-only instance of OrderedIntervalList, bound to a specific deployed contract.
func NewOrderedIntervalListTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderedIntervalListTransactor, error) {
	contract, err := bindOrderedIntervalList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListTransactor{contract: contract}, nil
}

// NewOrderedIntervalListFilterer creates a new log filterer instance of OrderedIntervalList, bound to a specific deployed contract.
func NewOrderedIntervalListFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderedIntervalListFilterer, error) {
	contract, err := bindOrderedIntervalList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListFilterer{contract: contract}, nil
}

// bindOrderedIntervalList binds a generic wrapper to an already deployed contract.
func bindOrderedIntervalList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderedIntervalListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderedIntervalList *OrderedIntervalListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrderedIntervalList.Contract.OrderedIntervalListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderedIntervalList *OrderedIntervalListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderedIntervalList.Contract.OrderedIntervalListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderedIntervalList *OrderedIntervalListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderedIntervalList.Contract.OrderedIntervalListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderedIntervalList *OrderedIntervalListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrderedIntervalList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderedIntervalList *OrderedIntervalListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderedIntervalList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderedIntervalList *OrderedIntervalListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderedIntervalList.Contract.contract.Transact(opts, method, params...)
}

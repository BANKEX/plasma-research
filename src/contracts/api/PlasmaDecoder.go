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

// PlasmaDecoderABI is the input ABI used to generate the binding from.
const PlasmaDecoderABI = "[]"

// PlasmaDecoderBin is the compiled bytecode used for deploying new contracts.
const PlasmaDecoderBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058200c9e3676d7131dc4e286bafc66ec94fede64c5746407efda5f8b68216bc738760029`

// DeployPlasmaDecoder deploys a new Ethereum contract, binding an instance of PlasmaDecoder to it.
func DeployPlasmaDecoder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlasmaDecoder, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaDecoderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlasmaDecoderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlasmaDecoder{PlasmaDecoderCaller: PlasmaDecoderCaller{contract: contract}, PlasmaDecoderTransactor: PlasmaDecoderTransactor{contract: contract}, PlasmaDecoderFilterer: PlasmaDecoderFilterer{contract: contract}}, nil
}

// PlasmaDecoder is an auto generated Go binding around an Ethereum contract.
type PlasmaDecoder struct {
	PlasmaDecoderCaller     // Read-only binding to the contract
	PlasmaDecoderTransactor // Write-only binding to the contract
	PlasmaDecoderFilterer   // Log filterer for contract events
}

// PlasmaDecoderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaDecoderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaDecoderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaDecoderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaDecoderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaDecoderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaDecoderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaDecoderSession struct {
	Contract     *PlasmaDecoder    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaDecoderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaDecoderCallerSession struct {
	Contract *PlasmaDecoderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PlasmaDecoderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaDecoderTransactorSession struct {
	Contract     *PlasmaDecoderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PlasmaDecoderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaDecoderRaw struct {
	Contract *PlasmaDecoder // Generic contract binding to access the raw methods on
}

// PlasmaDecoderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaDecoderCallerRaw struct {
	Contract *PlasmaDecoderCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaDecoderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaDecoderTransactorRaw struct {
	Contract *PlasmaDecoderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaDecoder creates a new instance of PlasmaDecoder, bound to a specific deployed contract.
func NewPlasmaDecoder(address common.Address, backend bind.ContractBackend) (*PlasmaDecoder, error) {
	contract, err := bindPlasmaDecoder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaDecoder{PlasmaDecoderCaller: PlasmaDecoderCaller{contract: contract}, PlasmaDecoderTransactor: PlasmaDecoderTransactor{contract: contract}, PlasmaDecoderFilterer: PlasmaDecoderFilterer{contract: contract}}, nil
}

// NewPlasmaDecoderCaller creates a new read-only instance of PlasmaDecoder, bound to a specific deployed contract.
func NewPlasmaDecoderCaller(address common.Address, caller bind.ContractCaller) (*PlasmaDecoderCaller, error) {
	contract, err := bindPlasmaDecoder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaDecoderCaller{contract: contract}, nil
}

// NewPlasmaDecoderTransactor creates a new write-only instance of PlasmaDecoder, bound to a specific deployed contract.
func NewPlasmaDecoderTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaDecoderTransactor, error) {
	contract, err := bindPlasmaDecoder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaDecoderTransactor{contract: contract}, nil
}

// NewPlasmaDecoderFilterer creates a new log filterer instance of PlasmaDecoder, bound to a specific deployed contract.
func NewPlasmaDecoderFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaDecoderFilterer, error) {
	contract, err := bindPlasmaDecoder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaDecoderFilterer{contract: contract}, nil
}

// bindPlasmaDecoder binds a generic wrapper to an already deployed contract.
func bindPlasmaDecoder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaDecoderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaDecoder *PlasmaDecoderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaDecoder.Contract.PlasmaDecoderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaDecoder *PlasmaDecoderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaDecoder.Contract.PlasmaDecoderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaDecoder *PlasmaDecoderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaDecoder.Contract.PlasmaDecoderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaDecoder *PlasmaDecoderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaDecoder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaDecoder *PlasmaDecoderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaDecoder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaDecoder *PlasmaDecoderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaDecoder.Contract.contract.Transact(opts, method, params...)
}

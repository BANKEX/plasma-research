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

// PlasmaBlocksABI is the input ABI used to generate the binding from.
const PlasmaBlocksABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"length\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BlocksSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"blocksLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"bytes\"}],\"name\":\"submitBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"bytes\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"submitBlocksSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PlasmaBlocksBin is the compiled bytecode used for deploying new contracts.
const PlasmaBlocksBin = `0x6080604081905260008054600160a060020a0319163317808255600160a060020a0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3610926806100576000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c0100000000000000000000000000000000000000000000000000000000900480638f32d59b116100785780638f32d59b14610226578063bada116414610242578063f25b3f99146102ef578063f2fde38b1461030c576100a5565b80631f10e5da146100aa578063715018a6146101f05780638ce0b5a2146101fa5780638da5cb5b14610202575b600080fd5b6101de600480360360608110156100c057600080fd5b813591908101906040810160208201356401000000008111156100e257600080fd5b8201836020820111156100f457600080fd5b8035906020019184600183028401116401000000008311171561011657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561016957600080fd5b82018360208201111561017b57600080fd5b8035906020019184600183028401116401000000008311171561019d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610332945050505050565b60408051918252519081900360200190f35b6101f861045e565b005b6101de6104c8565b61020a6104cf565b60408051600160a060020a039092168252519081900360200190f35b61022e6104de565b604080519115158252519081900360200190f35b6101de6004803603604081101561025857600080fd5b8135919081019060408101602082013564010000000081111561027a57600080fd5b82018360208201111561028c57600080fd5b803590602001918460018302840111640100000000831117156102ae57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104ef945050505050565b61020a6004803603602081101561030557600080fd5b5035610515565b6101f86004803603602081101561032257600080fd5b5035600160a060020a0316610541565b60008084846040516020018083815260200182805190602001908083835b6020831061036f5780518252601f199092019160209182019101610350565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405280519060200120905060006103b882610560565b90506103c481856105b1565b600160a060020a03166103d56104cf565b600160a060020a03161461044a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c6964207369676e6174757265000000000000000000000000000000604482015290519081900360640190fd5b6104548686610689565b9695505050505050565b6104666104de565b151561047157600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6001545b90565b600054600160a060020a031690565b600054600160a060020a0316331490565b60006104f96104de565b151561050457600080fd5b61050e8383610689565b5092915050565b600060018281548110151561052657fe5b600091825260209091200154600160a060020a031692915050565b6105496104de565b151561055457600080fd5b61055d81610808565b50565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600080600080845160411415156105ce5760009350505050610683565b50505060208201516040830151606084015160001a601b60ff821610156105f357601b015b8060ff16601b1415801561060b57508060ff16601c14155b1561061c5760009350505050610683565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015610673573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b6000806014835181151561069957fe5b6001549190049150841461070e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c69642066726f6d496e646578000000000000000000000000000000604482015290519081900360640190fd5b600154600090610724908663ffffffff61088516565b9050610736858363ffffffff61089a16565b6107416001826108b3565b50805b828110156107bf576000808260140260200190506c0100000000000000000000000081880151049150816001848a0181548110151561077f57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555050600101610744565b50818110156107ff576001546040805142815290517ff32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb999181900360200190a25b90039392505050565b600160a060020a038116151561081d57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561089457600080fd5b50900390565b6000828201838110156108ac57600080fd5b9392505050565b8154818355818111156108d7576000838152602090206108d79181019083016108dc565b505050565b6104cc91905b808211156108f657600081556001016108e2565b509056fea165627a7a72305820e4f58f85709cf1f33bf3e077fd3a272bcb87c3271b7bc21357eb9c08c2e32a700029`

// DeployPlasmaBlocks deploys a new Ethereum contract, binding an instance of PlasmaBlocks to it.
func DeployPlasmaBlocks(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlasmaBlocks, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaBlocksABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlasmaBlocksBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlasmaBlocks{PlasmaBlocksCaller: PlasmaBlocksCaller{contract: contract}, PlasmaBlocksTransactor: PlasmaBlocksTransactor{contract: contract}, PlasmaBlocksFilterer: PlasmaBlocksFilterer{contract: contract}}, nil
}

// PlasmaBlocks is an auto generated Go binding around an Ethereum contract.
type PlasmaBlocks struct {
	PlasmaBlocksCaller     // Read-only binding to the contract
	PlasmaBlocksTransactor // Write-only binding to the contract
	PlasmaBlocksFilterer   // Log filterer for contract events
}

// PlasmaBlocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaBlocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaBlocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaBlocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaBlocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaBlocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaBlocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaBlocksSession struct {
	Contract     *PlasmaBlocks     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaBlocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaBlocksCallerSession struct {
	Contract *PlasmaBlocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PlasmaBlocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaBlocksTransactorSession struct {
	Contract     *PlasmaBlocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PlasmaBlocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaBlocksRaw struct {
	Contract *PlasmaBlocks // Generic contract binding to access the raw methods on
}

// PlasmaBlocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaBlocksCallerRaw struct {
	Contract *PlasmaBlocksCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaBlocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaBlocksTransactorRaw struct {
	Contract *PlasmaBlocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaBlocks creates a new instance of PlasmaBlocks, bound to a specific deployed contract.
func NewPlasmaBlocks(address common.Address, backend bind.ContractBackend) (*PlasmaBlocks, error) {
	contract, err := bindPlasmaBlocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlocks{PlasmaBlocksCaller: PlasmaBlocksCaller{contract: contract}, PlasmaBlocksTransactor: PlasmaBlocksTransactor{contract: contract}, PlasmaBlocksFilterer: PlasmaBlocksFilterer{contract: contract}}, nil
}

// NewPlasmaBlocksCaller creates a new read-only instance of PlasmaBlocks, bound to a specific deployed contract.
func NewPlasmaBlocksCaller(address common.Address, caller bind.ContractCaller) (*PlasmaBlocksCaller, error) {
	contract, err := bindPlasmaBlocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlocksCaller{contract: contract}, nil
}

// NewPlasmaBlocksTransactor creates a new write-only instance of PlasmaBlocks, bound to a specific deployed contract.
func NewPlasmaBlocksTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaBlocksTransactor, error) {
	contract, err := bindPlasmaBlocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlocksTransactor{contract: contract}, nil
}

// NewPlasmaBlocksFilterer creates a new log filterer instance of PlasmaBlocks, bound to a specific deployed contract.
func NewPlasmaBlocksFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaBlocksFilterer, error) {
	contract, err := bindPlasmaBlocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlocksFilterer{contract: contract}, nil
}

// bindPlasmaBlocks binds a generic wrapper to an already deployed contract.
func bindPlasmaBlocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaBlocksABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaBlocks *PlasmaBlocksRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaBlocks.Contract.PlasmaBlocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaBlocks *PlasmaBlocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.PlasmaBlocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaBlocks *PlasmaBlocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.PlasmaBlocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaBlocks *PlasmaBlocksCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaBlocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaBlocks *PlasmaBlocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaBlocks *PlasmaBlocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_PlasmaBlocks *PlasmaBlocksCaller) Blocks(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaBlocks.contract.Call(opts, out, "blocks", i)
	return *ret0, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_PlasmaBlocks *PlasmaBlocksSession) Blocks(i *big.Int) (common.Address, error) {
	return _PlasmaBlocks.Contract.Blocks(&_PlasmaBlocks.CallOpts, i)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_PlasmaBlocks *PlasmaBlocksCallerSession) Blocks(i *big.Int) (common.Address, error) {
	return _PlasmaBlocks.Contract.Blocks(&_PlasmaBlocks.CallOpts, i)
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksCaller) BlocksLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaBlocks.contract.Call(opts, out, "blocksLength")
	return *ret0, err
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksSession) BlocksLength() (*big.Int, error) {
	return _PlasmaBlocks.Contract.BlocksLength(&_PlasmaBlocks.CallOpts)
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksCallerSession) BlocksLength() (*big.Int, error) {
	return _PlasmaBlocks.Contract.BlocksLength(&_PlasmaBlocks.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PlasmaBlocks *PlasmaBlocksCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaBlocks.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PlasmaBlocks *PlasmaBlocksSession) IsOwner() (bool, error) {
	return _PlasmaBlocks.Contract.IsOwner(&_PlasmaBlocks.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PlasmaBlocks *PlasmaBlocksCallerSession) IsOwner() (bool, error) {
	return _PlasmaBlocks.Contract.IsOwner(&_PlasmaBlocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaBlocks *PlasmaBlocksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaBlocks.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaBlocks *PlasmaBlocksSession) Owner() (common.Address, error) {
	return _PlasmaBlocks.Contract.Owner(&_PlasmaBlocks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaBlocks *PlasmaBlocksCallerSession) Owner() (common.Address, error) {
	return _PlasmaBlocks.Contract.Owner(&_PlasmaBlocks.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlasmaBlocks *PlasmaBlocksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaBlocks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlasmaBlocks *PlasmaBlocksSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.RenounceOwnership(&_PlasmaBlocks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlasmaBlocks *PlasmaBlocksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.RenounceOwnership(&_PlasmaBlocks.TransactOpts)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0xbada1164.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes) returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksTransactor) SubmitBlocks(opts *bind.TransactOpts, fromIndex *big.Int, newBlocks []byte) (*types.Transaction, error) {
	return _PlasmaBlocks.contract.Transact(opts, "submitBlocks", fromIndex, newBlocks)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0xbada1164.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes) returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksSession) SubmitBlocks(fromIndex *big.Int, newBlocks []byte) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.SubmitBlocks(&_PlasmaBlocks.TransactOpts, fromIndex, newBlocks)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0xbada1164.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes) returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksTransactorSession) SubmitBlocks(fromIndex *big.Int, newBlocks []byte) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.SubmitBlocks(&_PlasmaBlocks.TransactOpts, fromIndex, newBlocks)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x1f10e5da.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, rsv bytes) returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksTransactor) SubmitBlocksSigned(opts *bind.TransactOpts, fromIndex *big.Int, newBlocks []byte, rsv []byte) (*types.Transaction, error) {
	return _PlasmaBlocks.contract.Transact(opts, "submitBlocksSigned", fromIndex, newBlocks, rsv)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x1f10e5da.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, rsv bytes) returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksSession) SubmitBlocksSigned(fromIndex *big.Int, newBlocks []byte, rsv []byte) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.SubmitBlocksSigned(&_PlasmaBlocks.TransactOpts, fromIndex, newBlocks, rsv)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x1f10e5da.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, rsv bytes) returns(uint256)
func (_PlasmaBlocks *PlasmaBlocksTransactorSession) SubmitBlocksSigned(fromIndex *big.Int, newBlocks []byte, rsv []byte) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.SubmitBlocksSigned(&_PlasmaBlocks.TransactOpts, fromIndex, newBlocks, rsv)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PlasmaBlocks *PlasmaBlocksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlasmaBlocks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PlasmaBlocks *PlasmaBlocksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.TransferOwnership(&_PlasmaBlocks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PlasmaBlocks *PlasmaBlocksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlasmaBlocks.Contract.TransferOwnership(&_PlasmaBlocks.TransactOpts, newOwner)
}

// PlasmaBlocksBlocksSubmittedIterator is returned from FilterBlocksSubmitted and is used to iterate over the raw logs and unpacked data for BlocksSubmitted events raised by the PlasmaBlocks contract.
type PlasmaBlocksBlocksSubmittedIterator struct {
	Event *PlasmaBlocksBlocksSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaBlocksBlocksSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaBlocksBlocksSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaBlocksBlocksSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaBlocksBlocksSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaBlocksBlocksSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaBlocksBlocksSubmitted represents a BlocksSubmitted event raised by the PlasmaBlocks contract.
type PlasmaBlocksBlocksSubmitted struct {
	Length *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlocksSubmitted is a free log retrieval operation binding the contract event 0xf32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99.
//
// Solidity: e BlocksSubmitted(length indexed uint256, time uint256)
func (_PlasmaBlocks *PlasmaBlocksFilterer) FilterBlocksSubmitted(opts *bind.FilterOpts, length []*big.Int) (*PlasmaBlocksBlocksSubmittedIterator, error) {

	var lengthRule []interface{}
	for _, lengthItem := range length {
		lengthRule = append(lengthRule, lengthItem)
	}

	logs, sub, err := _PlasmaBlocks.contract.FilterLogs(opts, "BlocksSubmitted", lengthRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlocksBlocksSubmittedIterator{contract: _PlasmaBlocks.contract, event: "BlocksSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlocksSubmitted is a free log subscription operation binding the contract event 0xf32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99.
//
// Solidity: e BlocksSubmitted(length indexed uint256, time uint256)
func (_PlasmaBlocks *PlasmaBlocksFilterer) WatchBlocksSubmitted(opts *bind.WatchOpts, sink chan<- *PlasmaBlocksBlocksSubmitted, length []*big.Int) (event.Subscription, error) {

	var lengthRule []interface{}
	for _, lengthItem := range length {
		lengthRule = append(lengthRule, lengthItem)
	}

	logs, sub, err := _PlasmaBlocks.contract.WatchLogs(opts, "BlocksSubmitted", lengthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaBlocksBlocksSubmitted)
				if err := _PlasmaBlocks.contract.UnpackLog(event, "BlocksSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaBlocksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PlasmaBlocks contract.
type PlasmaBlocksOwnershipTransferredIterator struct {
	Event *PlasmaBlocksOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaBlocksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaBlocksOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaBlocksOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaBlocksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaBlocksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaBlocksOwnershipTransferred represents a OwnershipTransferred event raised by the PlasmaBlocks contract.
type PlasmaBlocksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PlasmaBlocks *PlasmaBlocksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlasmaBlocksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PlasmaBlocks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlocksOwnershipTransferredIterator{contract: _PlasmaBlocks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PlasmaBlocks *PlasmaBlocksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlasmaBlocksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PlasmaBlocks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaBlocksOwnershipTransferred)
				if err := _PlasmaBlocks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

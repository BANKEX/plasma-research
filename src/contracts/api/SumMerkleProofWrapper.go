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

// SumMerkleProofWrapperABI is the input ABI used to generate the binding from.
const SumMerkleProofWrapperABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"PLASMA_ASSETS_TOTAL_SIZE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint32\"},{\"name\":\"begin\",\"type\":\"uint32\"},{\"name\":\"end\",\"type\":\"uint32\"},{\"name\":\"item\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"root\",\"type\":\"uint256\"},{\"name\":\"rootLength\",\"type\":\"uint32\"}],\"name\":\"sumMerkleProofTest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"root\",\"type\":\"uint256\"},{\"name\":\"txProofBytes\",\"type\":\"bytes\"}],\"name\":\"sumMerkleProofFromBytesTest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SumMerkleProofWrapperBin is the compiled bytecode used for deploying new contracts.
const SumMerkleProofWrapperBin = `0x610a83610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004e5760e060020a60003504634775d43f8114610053578063a1e1005c14610114578063f93a936f146101ea575b600080fd5b6101006004803603604081101561006957600080fd5b8135919081019060408101602082013564010000000081111561008b57600080fd5b82018360208201111561009d57600080fd5b803590602001918460018302840111640100000000831117156100bf57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061020b945050505050565b604080519115158252519081900360200190f35b610100600480360360e081101561012a57600080fd5b63ffffffff823581169260208101358216926040820135909216916060820135919081019060a08101608082013564010000000081111561016a57600080fd5b82018360208201111561017c57600080fd5b8035906020019184600183028401116401000000008311171561019e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013563ffffffff1661023c565b6101f26102c8565b6040805163ffffffff9092168252519081900360200190f35b6000610215610a11565b61021e836102cf565b9050610234818562ffffff63ffffffff6102f816565b949350505050565b6000610246610a11565b6080604051908101604052808a63ffffffff16815260200160408051908101604052808b63ffffffff1681526020018a63ffffffff1681525081526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681525090506102bb8484836102f89092919063ffffffff16565b9998505050505050505050565b62ffffff81565b6102d7610a11565b6102f06102eb6102e68461045e565b610483565b610553565b90505b919050565b600080601885606001515181151561030c57fe5b049050600061033a86602001516000015187602001516020015163ffffffff1661062690919063ffffffff16565b60408701516020880151518851929350909160005b858160ff1610156103fb5760008061036e8c606001518460ff16610644565b9092509050600180851614156103bd5761038a82888389610670565b95506103a063ffffffff80871690849061062616565b94506103b663ffffffff8089169084906106e816565b96506103e2565b6103c987838884610670565b95506103df63ffffffff8089169084906106e816565b96505b5050600263ffffffff909216919091049060010161034f565b5063ffffffff821615801561041b57508663ffffffff168463ffffffff16145b80156102bb57508773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16149998505050505050505050565b610466610a40565b506040805180820190915281518152602082810190820152919050565b606061048e8261070a565b151561049957600080fd5b60006104a483610749565b9050806040519080825280602002602001820160405280156104e057816020015b6104cd610a40565b8152602001906001900390816104c55790505b50915060006104f284602001516107a7565b60208501510190506000805b8381101561054a5761050f83610810565b9150604080519081016040528083815260200184815250858281518110151561053457fe5b60209081029091010152918101916001016104fe565b50505050919050565b61055b610a11565b60806040519081016040528061058884600081518110151561057957fe5b906020019060200201516108a0565b63ffffffff1681526020016105bc6105b78560018151811015156105a857fe5b90602001906020020151610483565b6108dd565b81526020016105e28460028151811015156105d357fe5b9060200190602002015161092d565b73ffffffffffffffffffffffffffffffffffffffff16815260200161061e84600381518110151561060f57fe5b90602001906020020151610948565b905292915050565b600063ffffffff808416908316111561063e57600080fd5b50900390565b60180201602081015160249091015160e060020a909104916c0100000000000000000000000090910490565b6040805160e060020a63ffffffff96871681026020808401919091529590961690950260248601526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff938416810260288701529190921602603c8401528051603081850301815260509093019052815191012090565b600082820163ffffffff808516908216101561070357600080fd5b9392505050565b8051600090151561071d575060006102f3565b6020820151805160001a9060c060ff8316101561073f576000925050506102f3565b5060019392505050565b8051600090151561075c575060006102f3565b6000809050600061077084602001516107a7565b602085015185519181019250015b8082101561079e5761078f82610810565b6001909301929091019061077e565b50909392505050565b8051600090811a60808110156107c15760009150506102f3565b60b88110806107dc575060c081108015906107dc575060f881105b156107eb5760019150506102f3565b60c08110156107ff5760b5190190506102f3565b60f5190190506102f3565b50919050565b8051600090811a608081101561082a5760019150506102f3565b60b881101561083e57607e190190506102f3565b60c081101561086b5760b78103600184019350806020036101000a8451046001820181019350505061080a565b60f881101561087f5760be190190506102f3565b60019290920151602083900360f7016101000a900490910160f51901919050565b805160009081106108b057600080fd5b60006108bf83602001516107a7565b83516020948501518201519190039093036101000a90920492915050565b6108e5610a40565b604080519081016040528061090284600081518110151561057957fe5b63ffffffff16815260200161091f84600181518110151561057957fe5b63ffffffff16905292915050565b80516000906015101561093f57600080fd5b6102f0826108a0565b805160609060001061095957600080fd5b600061096883602001516107a7565b83516040805191839003808352601f19601f820116830160200190915291925060609082801561099f576020820181803883390190505b50905060008160200190506109bb8487602001510182856109c4565b50949350505050565b8015156109d057610a0c565b5b602081106109f0578251825260209283019290910190601f19016109d1565b8251825160208390036101000a60001901801990921691161782525b505050565b6040805160a081019091526000815260208101610a2c610a40565b815260006020820152606060409091015290565b60408051808201909152600080825260208201529056fea165627a7a72305820be6a21ed9c17ca75626f44751b5fc91f612af7e3d0fc3e06b9fdd300f8e5245d0029`

// DeploySumMerkleProofWrapper deploys a new Ethereum contract, binding an instance of SumMerkleProofWrapper to it.
func DeploySumMerkleProofWrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SumMerkleProofWrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(SumMerkleProofWrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SumMerkleProofWrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SumMerkleProofWrapper{SumMerkleProofWrapperCaller: SumMerkleProofWrapperCaller{contract: contract}, SumMerkleProofWrapperTransactor: SumMerkleProofWrapperTransactor{contract: contract}, SumMerkleProofWrapperFilterer: SumMerkleProofWrapperFilterer{contract: contract}}, nil
}

// SumMerkleProofWrapper is an auto generated Go binding around an Ethereum contract.
type SumMerkleProofWrapper struct {
	SumMerkleProofWrapperCaller     // Read-only binding to the contract
	SumMerkleProofWrapperTransactor // Write-only binding to the contract
	SumMerkleProofWrapperFilterer   // Log filterer for contract events
}

// SumMerkleProofWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SumMerkleProofWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SumMerkleProofWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SumMerkleProofWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SumMerkleProofWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SumMerkleProofWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SumMerkleProofWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SumMerkleProofWrapperSession struct {
	Contract     *SumMerkleProofWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SumMerkleProofWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SumMerkleProofWrapperCallerSession struct {
	Contract *SumMerkleProofWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SumMerkleProofWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SumMerkleProofWrapperTransactorSession struct {
	Contract     *SumMerkleProofWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SumMerkleProofWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SumMerkleProofWrapperRaw struct {
	Contract *SumMerkleProofWrapper // Generic contract binding to access the raw methods on
}

// SumMerkleProofWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SumMerkleProofWrapperCallerRaw struct {
	Contract *SumMerkleProofWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// SumMerkleProofWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SumMerkleProofWrapperTransactorRaw struct {
	Contract *SumMerkleProofWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSumMerkleProofWrapper creates a new instance of SumMerkleProofWrapper, bound to a specific deployed contract.
func NewSumMerkleProofWrapper(address common.Address, backend bind.ContractBackend) (*SumMerkleProofWrapper, error) {
	contract, err := bindSumMerkleProofWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofWrapper{SumMerkleProofWrapperCaller: SumMerkleProofWrapperCaller{contract: contract}, SumMerkleProofWrapperTransactor: SumMerkleProofWrapperTransactor{contract: contract}, SumMerkleProofWrapperFilterer: SumMerkleProofWrapperFilterer{contract: contract}}, nil
}

// NewSumMerkleProofWrapperCaller creates a new read-only instance of SumMerkleProofWrapper, bound to a specific deployed contract.
func NewSumMerkleProofWrapperCaller(address common.Address, caller bind.ContractCaller) (*SumMerkleProofWrapperCaller, error) {
	contract, err := bindSumMerkleProofWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofWrapperCaller{contract: contract}, nil
}

// NewSumMerkleProofWrapperTransactor creates a new write-only instance of SumMerkleProofWrapper, bound to a specific deployed contract.
func NewSumMerkleProofWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*SumMerkleProofWrapperTransactor, error) {
	contract, err := bindSumMerkleProofWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofWrapperTransactor{contract: contract}, nil
}

// NewSumMerkleProofWrapperFilterer creates a new log filterer instance of SumMerkleProofWrapper, bound to a specific deployed contract.
func NewSumMerkleProofWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*SumMerkleProofWrapperFilterer, error) {
	contract, err := bindSumMerkleProofWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SumMerkleProofWrapperFilterer{contract: contract}, nil
}

// bindSumMerkleProofWrapper binds a generic wrapper to an already deployed contract.
func bindSumMerkleProofWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SumMerkleProofWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SumMerkleProofWrapper *SumMerkleProofWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SumMerkleProofWrapper *SumMerkleProofWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SumMerkleProofWrapper *SumMerkleProofWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SumMerkleProofWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SumMerkleProofWrapper *SumMerkleProofWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SumMerkleProofWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SumMerkleProofWrapper *SumMerkleProofWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SumMerkleProofWrapper.Contract.contract.Transact(opts, method, params...)
}

// PLASMAASSETSTOTALSIZE is a free data retrieval call binding the contract method 0xf93a936f.
//
// Solidity: function PLASMA_ASSETS_TOTAL_SIZE() constant returns(uint32)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCaller) PLASMAASSETSTOTALSIZE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _SumMerkleProofWrapper.contract.Call(opts, out, "PLASMA_ASSETS_TOTAL_SIZE")
	return *ret0, err
}

// PLASMAASSETSTOTALSIZE is a free data retrieval call binding the contract method 0xf93a936f.
//
// Solidity: function PLASMA_ASSETS_TOTAL_SIZE() constant returns(uint32)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperSession) PLASMAASSETSTOTALSIZE() (uint32, error) {
	return _SumMerkleProofWrapper.Contract.PLASMAASSETSTOTALSIZE(&_SumMerkleProofWrapper.CallOpts)
}

// PLASMAASSETSTOTALSIZE is a free data retrieval call binding the contract method 0xf93a936f.
//
// Solidity: function PLASMA_ASSETS_TOTAL_SIZE() constant returns(uint32)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCallerSession) PLASMAASSETSTOTALSIZE() (uint32, error) {
	return _SumMerkleProofWrapper.Contract.PLASMAASSETSTOTALSIZE(&_SumMerkleProofWrapper.CallOpts)
}

// SumMerkleProofFromBytesTest is a free data retrieval call binding the contract method 0x4775d43f.
//
// Solidity: function sumMerkleProofFromBytesTest(root uint256, txProofBytes bytes) constant returns(bool)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCaller) SumMerkleProofFromBytesTest(opts *bind.CallOpts, root *big.Int, txProofBytes []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SumMerkleProofWrapper.contract.Call(opts, out, "sumMerkleProofFromBytesTest", root, txProofBytes)
	return *ret0, err
}

// SumMerkleProofFromBytesTest is a free data retrieval call binding the contract method 0x4775d43f.
//
// Solidity: function sumMerkleProofFromBytesTest(root uint256, txProofBytes bytes) constant returns(bool)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperSession) SumMerkleProofFromBytesTest(root *big.Int, txProofBytes []byte) (bool, error) {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofFromBytesTest(&_SumMerkleProofWrapper.CallOpts, root, txProofBytes)
}

// SumMerkleProofFromBytesTest is a free data retrieval call binding the contract method 0x4775d43f.
//
// Solidity: function sumMerkleProofFromBytesTest(root uint256, txProofBytes bytes) constant returns(bool)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCallerSession) SumMerkleProofFromBytesTest(root *big.Int, txProofBytes []byte) (bool, error) {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofFromBytesTest(&_SumMerkleProofWrapper.CallOpts, root, txProofBytes)
}

// SumMerkleProofTest is a free data retrieval call binding the contract method 0xa1e1005c.
//
// Solidity: function sumMerkleProofTest(index uint32, begin uint32, end uint32, item uint256, data bytes, root uint256, rootLength uint32) constant returns(bool)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCaller) SumMerkleProofTest(opts *bind.CallOpts, index uint32, begin uint32, end uint32, item *big.Int, data []byte, root *big.Int, rootLength uint32) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SumMerkleProofWrapper.contract.Call(opts, out, "sumMerkleProofTest", index, begin, end, item, data, root, rootLength)
	return *ret0, err
}

// SumMerkleProofTest is a free data retrieval call binding the contract method 0xa1e1005c.
//
// Solidity: function sumMerkleProofTest(index uint32, begin uint32, end uint32, item uint256, data bytes, root uint256, rootLength uint32) constant returns(bool)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperSession) SumMerkleProofTest(index uint32, begin uint32, end uint32, item *big.Int, data []byte, root *big.Int, rootLength uint32) (bool, error) {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofTest(&_SumMerkleProofWrapper.CallOpts, index, begin, end, item, data, root, rootLength)
}

// SumMerkleProofTest is a free data retrieval call binding the contract method 0xa1e1005c.
//
// Solidity: function sumMerkleProofTest(index uint32, begin uint32, end uint32, item uint256, data bytes, root uint256, rootLength uint32) constant returns(bool)
func (_SumMerkleProofWrapper *SumMerkleProofWrapperCallerSession) SumMerkleProofTest(index uint32, begin uint32, end uint32, item *big.Int, data []byte, root *big.Int, rootLength uint32) (bool, error) {
	return _SumMerkleProofWrapper.Contract.SumMerkleProofTest(&_SumMerkleProofWrapper.CallOpts, index, begin, end, item, data, root, rootLength)
}

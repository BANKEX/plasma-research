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

// OrderedIntervalListWrapperABI is the input ABI used to generate the binding from.
const OrderedIntervalListWrapperABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastInserted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"begin\",\"type\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prev\",\"type\":\"uint64\"},{\"name\":\"next\",\"type\":\"uint64\"},{\"name\":\"begin\",\"type\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\"},{\"name\":\"begin\",\"type\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrderedIntervalListWrapperBin is the compiled bytecode used for deploying new contracts.
const OrderedIntervalListWrapperBin = `0x608060405234801561001057600080fd5b5061002860006401000000006110a961002d82021704565b610159565b805415610085576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b81526020018061150a602b913960400191505060405180910390fd5b604080516080810182526000808252602080830182815293830182815260608401838152865460018101885596845291909220925192909401805493519151945167ffffffffffffffff1990941667ffffffffffffffff93841617604060020a608060020a031916680100000000000000009284169290920291909117608060020a60c060020a0319167001000000000000000000000000000000009483169490940293909317600160c060020a031678010000000000000000000000000000000000000000000000009190921602179055565b6113a2806101686000396000f3fe608060405234801561001057600080fd5b50600436106100b0576000357c0100000000000000000000000000000000000000000000000000000000900480638e008d8d116100835780638e008d8d14610145578063ada867981461014d578063d8b7cee514610199578063f3f6f0d7146101d3578063fadb10f1146101db576100b0565b8063089f5bda146100b55780632ca869bf146100f75780636b4f2c0f146100ff57806387bc028a14610107575b600080fd5b6100db600480360360208110156100cb57600080fd5b50356001604060020a0316610201565b604080516001604060020a039092168252519081900360200190f35b6100db61023c565b6100db610246565b6100db6004803603608081101561011d57600080fd5b506001604060020a038135811691602081013582169160408201358116916060013516610257565b6100db610291565b6101736004803603602081101561016357600080fd5b50356001604060020a03166102a0565b604080516001604060020a03938416815291909216602082015281519081900390910190f35b6101d1600480360360608110156101af57600080fd5b506001604060020a0381358116916020810135821691604090910135166102d2565b005b6100db6102eb565b6100db600480360360208110156101f157600080fd5b50356001604060020a03166102f7565b6000805481906001604060020a03841690811061021a57fe5b60009182526020909120015460c060020a90046001604060020a031692915050565b6000546000190190565b60006102526000610332565b905090565b600061026c818686868663ffffffff61034216565b6002805467ffffffffffffffff19166001604060020a03831617905595945050505050565b6002546001604060020a031690565b600080806102b4818563ffffffff61035d16565b546001604060020a0380821696604060020a90920416945092505050565b6102e5600084848463ffffffff6103d416565b50505050565b60006102526000610863565b6000805481906001604060020a03841690811061031057fe5b600091825260209091200154608060020a90046001604060020a031692915050565b600101546001604060020a031690565b60006103538686868686600061087a565b9695505050505050565b81546000906001604060020a038316106103ab5760405160e560020a62461bcd02815260040180806020018281038252602a815260200180611219602a913960400191505060405180910390fd5b825483906001604060020a0384169081106103c257fe5b90600052602060002001905092915050565b60006001604060020a0380831690841610610439576040805160e560020a62461bcd02815260206004820181905260248201527f726967687420626f756e64206c657373207468616e206c65667420626f756e64604482015290519081900360640190fd5b84546001604060020a0385161061049a576040805160e560020a62461bcd02815260206004820152601260248201527f76616c696420696e64657820626f756e64730000000000000000000000000000604482015290519081900360640190fd5b845460009086906001604060020a0387169081106104b457fe5b6000918252602082200180548854919350889160c060020a9091046001604060020a03169081106104e157fe5b60009182526020822084548a5491909201935089916001604060020a03608060020a9091041690811061051057fe5b60009182526020909120845491019150604060020a90046001604060020a03161515610586576040805160e560020a62461bcd02815260206004820152601f60248201527f72656d6f76656420696e74657276616c20646f65736e27742065786973747300604482015290519081900360640190fd5b82546001604060020a038088169116118015906105b8575082546001604060020a03604060020a909104811690861611155b151561060e576040805160e560020a62461bcd02815260206004820152601e60248201527f696e636f72726563742072656d6f7665642072616e676520626f756e64730000604482015290519081900360640190fd5b82546001604060020a038082168882161491604060020a90048116908716148180156106375750805b15610776578454600060c060020a9091046001604060020a03161115610693578454845477ffffffffffffffff000000000000000000000000000000001916608060020a918290046001604060020a03169091021784556106c0565b845460018b018054608060020a9092046001604060020a031667ffffffffffffffff199092169190911790555b84546000608060020a9091046001604060020a03161115610716578454835477ffffffffffffffffffffffffffffffffffffffffffffffff1660c060020a918290046001604060020a031690910217835561074f565b845460018b0180546fffffffffffffffff0000000000000000191660c060020a9092046001604060020a0316604060020a029190911790555b89548a906001604060020a038b1690811061076657fe5b6000918252602082200155610856565b811561079b57845467ffffffffffffffff19166001604060020a038816178555610856565b80156107ce5784546fffffffffffffffff00000000000000001916604060020a6001604060020a038a1602178555610856565b84546001604060020a03898116604060020a9081026fffffffffffffffff000000000000000019841617808955920481169161081a918d918d91608060020a909104168b85600161087a565b865477ffffffffffffffff000000000000000000000000000000001916608060020a6001604060020a0392831681029190911780895504169650505b5050505050949350505050565b60010154604060020a90046001604060020a031690565b60006001604060020a03808416908516106108c95760405160e560020a62461bcd0281526004018080602001828103825260278152602001806112436027913960400191505060405180910390fd5b600187015460006001604060020a03918216119087161515806108f457506001604060020a03861615155b1515146109355760405160e560020a62461bcd02815260040180806020018281038252602c815260200180611291602c913960400191505060405180910390fd5b61093e876110a2565b151561094d5761094d876110a9565b865460009088906001604060020a03891690811061096757fe5b6000918252602082208a549101925089906001604060020a03891690811061098b57fe5b906000526020600020019050876001604060020a0316600014806109c4575081546001604060020a03604060020a909104811690871610155b1515610a045760405160e560020a62461bcd02815260040180806020018281038252602781526020018061126a6027913960400191505060405180910390fd5b6001604060020a0387161580610a28575080546001604060020a0390811690861611155b1515610a685760405160e560020a62461bcd0281526004018080602001828103825260258152602001806113006025913960400191505060405180910390fd5b6000876001604060020a03161115156000896001604060020a03161115151415610b075781546001604060020a03888116608060020a90920416148015610ac2575080546001604060020a0389811660c060020a90920416145b1515610b025760405160e560020a62461bcd0281526004018080602001828103825260378152602001806111e26037913960400191505060405180910390fd5b610c8f565b6000876001604060020a03161115610b875760018901546001604060020a038881169116148015610b475750805460c060020a90046001604060020a0316155b1515610b025760405160e560020a62461bcd0281526004018080602001828103825260278152602001806113256027913960400191505060405180910390fd5b6000886001604060020a03161115610c8f5760018901546001604060020a03898116604060020a90920416148015610bce57508154608060020a90046001604060020a0316155b1515610c0e5760405160e560020a62461bcd0281526004018080602001828103825260268152602001806111bc6026913960400191505060405180910390fd5b8380610c31575060018901546001604060020a03898116604060020a9092041614155b80610c4f575081546001604060020a03878116604060020a90920416145b1515610c8f5760405160e560020a62461bcd0281526004018080602001828103825260438152602001806112bd6043913960600191505060405180910390fd5b600080896001604060020a0316118015610cbc575082546001604060020a03888116604060020a90920416145b9050600080896001604060020a0316118015610ce4575082546001604060020a038881169116145b905081158015610cf2575080155b15610f07578a6000018054905094508a6000016080604051908101604052808a6001604060020a03168152602001896001604060020a031681526020018b6001604060020a031681526020018c6001604060020a031681525090806001815401808255809150509060018203906000526020600020016000909192909190915060008201518160000160006101000a8154816001604060020a0302191690836001604060020a0316021790555060208201518160000160086101000a8154816001604060020a0302191690836001604060020a0316021790555060408201518160000160106101000a8154816001604060020a0302191690836001604060020a0316021790555060608201518160000160186101000a8154816001604060020a0302191690836001604060020a031602179055505050506000896001604060020a03161115610e6f57825477ffffffffffffffffffffffffffffffffffffffffffffffff1660c060020a6001604060020a03871602178355610e9c565b60018b0180546fffffffffffffffff00000000000000001916604060020a6001604060020a038816021790555b60008a6001604060020a03161115610ee357835477ffffffffffffffff000000000000000000000000000000001916608060020a6001604060020a03871602178455610f02565b60018b01805467ffffffffffffffff19166001604060020a0387161790555b611094565b818015610f115750805b1561103a57825484546001604060020a03604060020a92839004811683026fffffffffffffffff000000000000000019909216919091178087558554608060020a9081900483160277ffffffffffffffff000000000000000000000000000000001990911617865560018d01548c97508b8216929004161415610fbf5760018b0180546fffffffffffffffff00000000000000001916604060020a6001604060020a03881602179055611013565b82548b5486918d91608060020a9091046001604060020a0316908110610fe157fe5b9060005260206000200160000160186101000a8154816001604060020a0302191690836001604060020a031602179055505b8a548b906001604060020a038b1690811061102a57fe5b6000918252602082200155611094565b81156110705783546fffffffffffffffff00000000000000001916604060020a6001604060020a03891602178455899450611094565b801561109457825467ffffffffffffffff19166001604060020a0389161783558894505b505050509695505050505050565b5460001090565b8054156110ea5760405160e560020a62461bcd02815260040180806020018281038252602b81526020018061134c602b913960400191505060405180910390fd5b604080516080810182526000808252602080830182815293830182815260608401838152865460018101885596845291909220925192909401805493519151945167ffffffffffffffff199094166001604060020a03938416176fffffffffffffffff00000000000000001916604060020a928416929092029190911777ffffffffffffffff000000000000000000000000000000001916608060020a948316949094029390931777ffffffffffffffffffffffffffffffffffffffffffffffff1660c060020a919092160217905556fe707265762073686f756c6420726566657220746f20746865206c61737420696e74657276616c7072657620616e64206e6578742073686f756c6420726566657220746f20746865206e65696768626f72696e6720696e74657276616c73696e74657276616c20696420646f65736e27742065786973747320696e20696e74657276616c20736574726967687420626f756e64206c657373206f7220657175616c20746f206c65667420626f756e64626567696e20636f756c64206e6f7420696e74657273656374207072657620696e74657276616c7072657620616e64206e65787420636f756c64206265207a65726f20696666206e6f20696e74657276616c7373686f756c6420626567696e2066726f6d2074686520656e64206f66206c617465737420696e74657276616c207768656e20616464696e6720746f2074686520656e64656e6420636f756c64206e6f7420696e74657273656374206e65787420696e74657276616c6e6578742073686f756c6420726566657220746f2074686520666972737420696e74657276616c4f726465726564496e74657276616c4c6973742077617320616c726561647920696e697469616c697a6564a165627a7a723058200517041ee47925de32b469523c49804ab60535f788e9bfb6df64743abf7aed9600294f726465726564496e74657276616c4c6973742077617320616c726561647920696e697469616c697a6564`

// DeployOrderedIntervalListWrapper deploys a new Ethereum contract, binding an instance of OrderedIntervalListWrapper to it.
func DeployOrderedIntervalListWrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrderedIntervalListWrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderedIntervalListWrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderedIntervalListWrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderedIntervalListWrapper{OrderedIntervalListWrapperCaller: OrderedIntervalListWrapperCaller{contract: contract}, OrderedIntervalListWrapperTransactor: OrderedIntervalListWrapperTransactor{contract: contract}, OrderedIntervalListWrapperFilterer: OrderedIntervalListWrapperFilterer{contract: contract}}, nil
}

// OrderedIntervalListWrapper is an auto generated Go binding around an Ethereum contract.
type OrderedIntervalListWrapper struct {
	OrderedIntervalListWrapperCaller     // Read-only binding to the contract
	OrderedIntervalListWrapperTransactor // Write-only binding to the contract
	OrderedIntervalListWrapperFilterer   // Log filterer for contract events
}

// OrderedIntervalListWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderedIntervalListWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderedIntervalListWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderedIntervalListWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderedIntervalListWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderedIntervalListWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderedIntervalListWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderedIntervalListWrapperSession struct {
	Contract     *OrderedIntervalListWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OrderedIntervalListWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderedIntervalListWrapperCallerSession struct {
	Contract *OrderedIntervalListWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// OrderedIntervalListWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderedIntervalListWrapperTransactorSession struct {
	Contract     *OrderedIntervalListWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// OrderedIntervalListWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderedIntervalListWrapperRaw struct {
	Contract *OrderedIntervalListWrapper // Generic contract binding to access the raw methods on
}

// OrderedIntervalListWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderedIntervalListWrapperCallerRaw struct {
	Contract *OrderedIntervalListWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// OrderedIntervalListWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderedIntervalListWrapperTransactorRaw struct {
	Contract *OrderedIntervalListWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderedIntervalListWrapper creates a new instance of OrderedIntervalListWrapper, bound to a specific deployed contract.
func NewOrderedIntervalListWrapper(address common.Address, backend bind.ContractBackend) (*OrderedIntervalListWrapper, error) {
	contract, err := bindOrderedIntervalListWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListWrapper{OrderedIntervalListWrapperCaller: OrderedIntervalListWrapperCaller{contract: contract}, OrderedIntervalListWrapperTransactor: OrderedIntervalListWrapperTransactor{contract: contract}, OrderedIntervalListWrapperFilterer: OrderedIntervalListWrapperFilterer{contract: contract}}, nil
}

// NewOrderedIntervalListWrapperCaller creates a new read-only instance of OrderedIntervalListWrapper, bound to a specific deployed contract.
func NewOrderedIntervalListWrapperCaller(address common.Address, caller bind.ContractCaller) (*OrderedIntervalListWrapperCaller, error) {
	contract, err := bindOrderedIntervalListWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListWrapperCaller{contract: contract}, nil
}

// NewOrderedIntervalListWrapperTransactor creates a new write-only instance of OrderedIntervalListWrapper, bound to a specific deployed contract.
func NewOrderedIntervalListWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderedIntervalListWrapperTransactor, error) {
	contract, err := bindOrderedIntervalListWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListWrapperTransactor{contract: contract}, nil
}

// NewOrderedIntervalListWrapperFilterer creates a new log filterer instance of OrderedIntervalListWrapper, bound to a specific deployed contract.
func NewOrderedIntervalListWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderedIntervalListWrapperFilterer, error) {
	contract, err := bindOrderedIntervalListWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderedIntervalListWrapperFilterer{contract: contract}, nil
}

// bindOrderedIntervalListWrapper binds a generic wrapper to an already deployed contract.
func bindOrderedIntervalListWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderedIntervalListWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrderedIntervalListWrapper.Contract.OrderedIntervalListWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.OrderedIntervalListWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.OrderedIntervalListWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrderedIntervalListWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.contract.Transact(opts, method, params...)
}

// FirstIndex is a free data retrieval call binding the contract method 0x6b4f2c0f.
//
// Solidity: function firstIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) FirstIndex(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "firstIndex")
	return *ret0, err
}

// FirstIndex is a free data retrieval call binding the contract method 0x6b4f2c0f.
//
// Solidity: function firstIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) FirstIndex() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.FirstIndex(&_OrderedIntervalListWrapper.CallOpts)
}

// FirstIndex is a free data retrieval call binding the contract method 0x6b4f2c0f.
//
// Solidity: function firstIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) FirstIndex() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.FirstIndex(&_OrderedIntervalListWrapper.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(index uint64) constant returns(begin uint64, end uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) Get(opts *bind.CallOpts, index uint64) (struct {
	Begin uint64
	End   uint64
}, error) {
	ret := new(struct {
		Begin uint64
		End   uint64
	})
	out := ret
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "get", index)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(index uint64) constant returns(begin uint64, end uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) Get(index uint64) (struct {
	Begin uint64
	End   uint64
}, error) {
	return _OrderedIntervalListWrapper.Contract.Get(&_OrderedIntervalListWrapper.CallOpts, index)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(index uint64) constant returns(begin uint64, end uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) Get(index uint64) (struct {
	Begin uint64
	End   uint64
}, error) {
	return _OrderedIntervalListWrapper.Contract.Get(&_OrderedIntervalListWrapper.CallOpts, index)
}

// GetNext is a free data retrieval call binding the contract method 0xfadb10f1.
//
// Solidity: function getNext(index uint64) constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) GetNext(opts *bind.CallOpts, index uint64) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "getNext", index)
	return *ret0, err
}

// GetNext is a free data retrieval call binding the contract method 0xfadb10f1.
//
// Solidity: function getNext(index uint64) constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) GetNext(index uint64) (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.GetNext(&_OrderedIntervalListWrapper.CallOpts, index)
}

// GetNext is a free data retrieval call binding the contract method 0xfadb10f1.
//
// Solidity: function getNext(index uint64) constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) GetNext(index uint64) (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.GetNext(&_OrderedIntervalListWrapper.CallOpts, index)
}

// GetPrev is a free data retrieval call binding the contract method 0x089f5bda.
//
// Solidity: function getPrev(index uint64) constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) GetPrev(opts *bind.CallOpts, index uint64) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "getPrev", index)
	return *ret0, err
}

// GetPrev is a free data retrieval call binding the contract method 0x089f5bda.
//
// Solidity: function getPrev(index uint64) constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) GetPrev(index uint64) (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.GetPrev(&_OrderedIntervalListWrapper.CallOpts, index)
}

// GetPrev is a free data retrieval call binding the contract method 0x089f5bda.
//
// Solidity: function getPrev(index uint64) constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) GetPrev(index uint64) (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.GetPrev(&_OrderedIntervalListWrapper.CallOpts, index)
}

// LastIndex is a free data retrieval call binding the contract method 0xf3f6f0d7.
//
// Solidity: function lastIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) LastIndex(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "lastIndex")
	return *ret0, err
}

// LastIndex is a free data retrieval call binding the contract method 0xf3f6f0d7.
//
// Solidity: function lastIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) LastIndex() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.LastIndex(&_OrderedIntervalListWrapper.CallOpts)
}

// LastIndex is a free data retrieval call binding the contract method 0xf3f6f0d7.
//
// Solidity: function lastIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) LastIndex() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.LastIndex(&_OrderedIntervalListWrapper.CallOpts)
}

// LastInserted is a free data retrieval call binding the contract method 0x8e008d8d.
//
// Solidity: function lastInserted() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) LastInserted(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "lastInserted")
	return *ret0, err
}

// LastInserted is a free data retrieval call binding the contract method 0x8e008d8d.
//
// Solidity: function lastInserted() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) LastInserted() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.LastInserted(&_OrderedIntervalListWrapper.CallOpts)
}

// LastInserted is a free data retrieval call binding the contract method 0x8e008d8d.
//
// Solidity: function lastInserted() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) LastInserted() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.LastInserted(&_OrderedIntervalListWrapper.CallOpts)
}

// MaxIndex is a free data retrieval call binding the contract method 0x2ca869bf.
//
// Solidity: function maxIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCaller) MaxIndex(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _OrderedIntervalListWrapper.contract.Call(opts, out, "maxIndex")
	return *ret0, err
}

// MaxIndex is a free data retrieval call binding the contract method 0x2ca869bf.
//
// Solidity: function maxIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) MaxIndex() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.MaxIndex(&_OrderedIntervalListWrapper.CallOpts)
}

// MaxIndex is a free data retrieval call binding the contract method 0x2ca869bf.
//
// Solidity: function maxIndex() constant returns(uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperCallerSession) MaxIndex() (uint64, error) {
	return _OrderedIntervalListWrapper.Contract.MaxIndex(&_OrderedIntervalListWrapper.CallOpts)
}

// Insert is a paid mutator transaction binding the contract method 0x87bc028a.
//
// Solidity: function insert(prev uint64, next uint64, begin uint64, end uint64) returns(id uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperTransactor) Insert(opts *bind.TransactOpts, prev uint64, next uint64, begin uint64, end uint64) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.contract.Transact(opts, "insert", prev, next, begin, end)
}

// Insert is a paid mutator transaction binding the contract method 0x87bc028a.
//
// Solidity: function insert(prev uint64, next uint64, begin uint64, end uint64) returns(id uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) Insert(prev uint64, next uint64, begin uint64, end uint64) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.Insert(&_OrderedIntervalListWrapper.TransactOpts, prev, next, begin, end)
}

// Insert is a paid mutator transaction binding the contract method 0x87bc028a.
//
// Solidity: function insert(prev uint64, next uint64, begin uint64, end uint64) returns(id uint64)
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperTransactorSession) Insert(prev uint64, next uint64, begin uint64, end uint64) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.Insert(&_OrderedIntervalListWrapper.TransactOpts, prev, next, begin, end)
}

// Remove is a paid mutator transaction binding the contract method 0xd8b7cee5.
//
// Solidity: function remove(index uint64, begin uint64, end uint64) returns()
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperTransactor) Remove(opts *bind.TransactOpts, index uint64, begin uint64, end uint64) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.contract.Transact(opts, "remove", index, begin, end)
}

// Remove is a paid mutator transaction binding the contract method 0xd8b7cee5.
//
// Solidity: function remove(index uint64, begin uint64, end uint64) returns()
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperSession) Remove(index uint64, begin uint64, end uint64) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.Remove(&_OrderedIntervalListWrapper.TransactOpts, index, begin, end)
}

// Remove is a paid mutator transaction binding the contract method 0xd8b7cee5.
//
// Solidity: function remove(index uint64, begin uint64, end uint64) returns()
func (_OrderedIntervalListWrapper *OrderedIntervalListWrapperTransactorSession) Remove(index uint64, begin uint64, end uint64) (*types.Transaction, error) {
	return _OrderedIntervalListWrapper.Contract.Remove(&_OrderedIntervalListWrapper.TransactOpts, index, begin, end)
}

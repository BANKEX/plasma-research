// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier

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

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"

// PairingBin is the compiled bytecode used for deploying new contracts.
const PairingBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208670e4248a54e7b3c474b64e54b64b7d65e9a6e719fdbad32a29460f5d1a29a30029`

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256[2]\"},{\"name\":\"a_p\",\"type\":\"uint256[2]\"},{\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"name\":\"b_p\",\"type\":\"uint256[2]\"},{\"name\":\"c\",\"type\":\"uint256[2]\"},{\"name\":\"c_p\",\"type\":\"uint256[2]\"},{\"name\":\"h\",\"type\":\"uint256[2]\"},{\"name\":\"k\",\"type\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256[3]\"}],\"name\":\"verifyTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"s\",\"type\":\"string\"}],\"name\":\"Verified\",\"type\":\"event\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
const VerifierBin = `0x608060405234801561001057600080fd5b506113f3806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c834d03d8114610045575b600080fd5b34801561005157600080fd5b506040805180820182526101d69136916004916044919083906002908390839080828437505060408051808201825294979695818101959450925060029150839083908082843750506040805180820190915293969594608081019493509150600290506000835b828210156100f25760408051808201825290808402860190600290839083908082843750505091835250506001909101906020016100b9565b5050604080518082018252939695948181019493509150600290839083908082843750506040805180820182529497969581810195945092506002915083908390808284375050604080518082018252949796958181019594509250600291508390839080828437505060408051808201825294979695818101959450925060029150839083908082843750506040805180820182529497969581810195945092506002915083908390808284375050604080516060818101909252949796958181019594509250600391508390839080828437509396506101ea95505050505050565b604080519115158252519081900360200190f35b60006101f4611232565b6040805180820182528c5181526020808e015181830152908352815180830183528c5181528c82015181830152838201528151608080820184528c51518285019081528d51840151606080850191909152908352845180860186528e8501805151825251850151818601528385015285850192909252835180850185528c5181528c8401518185015282860152835180850185528b5181528b8401518185015281860152835180850185528a5181528a8401518185015260a08601528351808501855289518152898401518185015260e08601528351808501855288518152888401518185015260c086015283516003808252918101909452909260009290919082018480388339019050509150600090505b60038110156103415784816003811061031c57fe5b6020020151828281518110151561032f57fe5b60209081029091010152600101610307565b61034b82846103f5565b15156103e0576040805160208082526022908201527f5472616e73616374696f6e207375636365737366756c6c792076657269666965818301527f642e000000000000000000000000000000000000000000000000000000000000606082015290517f3f3cfdb26fb5f9f1786ab4f1a1f9cd4c0b5e726cbdfc26e495261731aad44e399181900360800190a1600193506103e5565b600093505b5050509998505050505050505050565b60006103ff6112a7565b610407611317565b6000610411610600565b92508260e0015151865160010114151561042a57600080fd5b5050604080518082019091526000808252602082018190525b855181101561049e576104948261048f8560e001518460010181518110151561046857fe5b90602001906020020151898581518110151561048057fe5b90602001906020020151610c27565b610c82565b9150600101610443565b6104c4828460e0015160008151811015156104b557fe5b90602001906020020151610c82565b8551845160208801519294506104ea926104dd90610cd1565b6104e5610d5e565b610e1e565b15156104f957600193506105f7565b610513836020015186604001516104dd8860600151610cd1565b151561052257600293506105f7565b61053c856080015184604001516104dd8860a00151610cd1565b151561054b57600393506105f7565b6105928560c0015184606001516105766105718661048f8b600001518c60800151610c82565b610cd1565b8660a001516105888860800151610cd1565b8a60400151610f09565b15156105a157600493506105f7565b6105e36105b2838760000151610c82565b86604001516105c48860e00151610cd1565b8660c001516105d68a60800151610cd1565b6105de610d5e565b610f09565b15156105f257600593506105f7565b600093505b50505092915050565b6106086112a7565b60408051608080820183527f079f9d407f001c74133eafa3d60b1d8587ca4ffe84425b6766ad170642d2c36c8284019081527f032d491df3ba3c44e580db83d9bdc6e573a68c560ad9604559e01b8f1c3d1061606080850191909152908352835180850185527f1335b32bbefdf416908cca78265304b75fe7de5b5d826ba2f0c9b7d328f2429081527f25c71faaee4521b72a64c4e34760cd80da33f3d5a64f209c9cf653e247e51f8d60208281019190915280850191909152928552835180850185527f1ec6b704ff873c188e0fd8822d0d974fbbe34ba6cfc4f962dca84fd5ea8aadf481527f0bb6ed351010e05f1b7d8e096d97d1ea272b175acc6d4f356ca4f5bbff4423f38185015285840152835180830185527e246b33a1b58bf42401609c73e15ab68cd6e80a34f989155d89ffa9a8248ca48186019081527f2f819e180c22967ad3dce966a20f66e5935769e4d634535a43df0f560ea21626828401528152845180860186527f03c2acf22d7f5e3291949ad039fd444c8e9115ffe24936b881749f07cbbb6df481527f0cfccf3c6a8e48f59764f8e039a7eae079d4170a1dbb8432d0c9d61210335d6c818601528185015285850152835180830185527f26375034e3259890fca1b84d5f37e086f4c4424099743ed1a8cbb6c7bff221348186019081527f07312b6989154b86af327a4247111b06107f9f14a30045e88c504e0efc5bce8c828401528152845180860186527f219905bce1adbfa8a0e1cda2f91bfca2df5b9425cc5329e62b319418e6841c1781527f0c1d9376b8a039332a23a10b2cceb69a88db54b63469fa47c35e90764077f788818601528185015285820152835180850185527f2c715ab1d691fafacf8449f53d705aaebc4658c5d0b4982127ee324eab59bf1281527f2a98866742bb386adcf05843a66420a3bbfb9218bfa0b34ba80a26adb53fd42a8185015285830152835180830185527f1bbde84e30d10848b498c2a747b150cf2c395145a9f120d09a1be0664104ee0b8186019081527f0307f39b5036faf71a1f2b2ee90424838370fdbddd07a8ff16200db6331de3fa828401528152845180860186527f2fe8c1df1f3d813c026a72c8ca3425270890434d88b3e22dcbd7502e2d11420181527f3030a22cca28d8457bfd078cd0a0abe08e6d5626ef9481219063f409f7ea9398818601528185015260a080870191909152845192830185527f0e242be564a2f8e998f8ca0adf9fd9bd03ae2f6ea3b94aeb170ea626a6d66ee58386019081527f2a802b0f73d8699d631f9e6fb499ce1a330644cd7e09ef31627dbfbcdffde76192840192909252908252835180850185527f2d012db8b943274a3e074d3964675c1b860109ea3651c77d1f77fd6180dfb2cd81527e63d7734c58792b72fdb16742c875d9c46765a871eb162effa1b0fa51e39b76818501529282019290925260c084015281516004808252918101909252816020015b610a40611317565b815260200190600190039081610a3857505060e08201908152604080518082019091527f278711fe12e218c6c09e0890a97099c97437e126490b55c50632d8bb33eb811081527f2da72a36dc613036b6c06c1ed9f8664bb7a1138501343ed09bfb22f48e9601c46020820152905180516000908110610abb57fe5b9060200190602002018190525060408051908101604052807f1d24d21101dcf4dd4b14ddb7ad96afcbe0a7bcb6f811ec60013cb99cd72168c981526020017f2feca8ed372c5b6f8f1bf6f5a11084bc33ecb19e9580890430221a7d6dbb089b8152508160e001516001815181101515610b3057fe5b9060200190602002018190525060408051908101604052807f2e5de941e2e7597876927ac222346ff429328cb5469324b4d894ef195183dcb081526020017f13da8a6394f4053005de0f03d04ef9b881a88b711263dd511288356d4407443d8152508160e001516002815181101515610ba557fe5b9060200190602002018190525060408051908101604052807f126efd1002db6c73dcc66eee305bef0110cd2c8960b50ff9ea59e0b2f98779dc81526020017f091108cb60af3abebfc0e42afd3c0fa9b33606b24ff09ac5c0ed28d8e25712ef8152508160e001516003815181101515610c1a57fe5b6020908102909101015290565b610c2f611317565b610c3761132e565b83518152602080850151908201526040810183905260006060836080848460076107d05a03f19050808015610c6b57610c6d565bfe5b50801515610c7a57600080fd5b505092915050565b610c8a611317565b610c9261134d565b8351815260208085015181830152835160408301528301516060808301919091526000908360c0848460066107d05a03f19050808015610c6b57610c6d565b610cd9611317565b81517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790158015610d0c57506020830151155b15610d2c5760408051808201909152600080825260208201529150610d58565b604080519081016040528084600001518152602001828560200151811515610d5057fe5b068303905291505b50919050565b610d6661136c565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b82527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60208381019190915281019190915290565b6040805160028082526060828101909352600092918291816020015b610e42611317565b815260200190600190039081610e3a57505060408051600280825260608201909252919350602082015b610e7461136c565b815260200190600190039081610e6c57905050905086826000815181101515610e9957fe5b602090810290910101528151859083906001908110610eb457fe5b602090810290910101528051869082906000908110610ecf57fe5b602090810290910101528051849082906001908110610eea57fe5b60209081029091010152610efe828261102d565b979650505050505050565b604080516003808252608082019092526000916060918291816020015b610f2e611317565b815260200190600190039081610f2657505060408051600380825260808201909252919350602082015b610f6061136c565b815260200190600190039081610f5857905050905088826000815181101515610f8557fe5b602090810290910101528151879083906001908110610fa057fe5b602090810290910101528151859083906002908110610fbb57fe5b602090810290910101528051889082906000908110610fd657fe5b602090810290910101528051869082906001908110610ff157fe5b60209081029091010152805184908290600290811061100c57fe5b60209081029091010152611020828261102d565b9998505050505050505050565b60008060006060600061103e61138d565b865188516000911461104f57600080fd5b8851955085600602945084604051908082528060200260200182016040528015611083578160200160208202803883390190505b509350600092505b858310156111f85788838151811015156110a157fe5b602090810290910101515184518590600686029081106110bd57fe5b6020908102909101015288518990849081106110d557fe5b906020019060200201516020015184846006026001018151811015156110f757fe5b60209081029091010152875188908490811061110f57fe5b60209081029190910101515151845185906002600687020190811061113057fe5b60209081029091010152875188908490811061114857fe5b602090810291909101810151510151845185906003600687020190811061116b57fe5b60209081029091010152875188908490811061118357fe5b60209081029190910181015101515184518590600460068702019081106111a657fe5b6020908102909101015287518890849081106111be57fe5b602090810291909101810151810151015184518590600560068702019081106111e357fe5b6020908102909101015260019092019161108b565b6020826020870260208701600060086107d05a03f19050808015610c6b575080151561122357600080fd5b50511515979650505050505050565b61024060405190810160405280611247611317565b8152602001611254611317565b815260200161126161136c565b815260200161126e611317565b815260200161127b611317565b8152602001611288611317565b8152602001611295611317565b81526020016112a2611317565b905290565b610320604051908101604052806112bc61136c565b81526020016112c9611317565b81526020016112d661136c565b81526020016112e361136c565b81526020016112f0611317565b81526020016112fd61136c565b815260200161130a61136c565b8152602001606081525090565b604080518082019091526000808252602082015290565b6060604051908101604052806003906020820280388339509192915050565b6080604051908101604052806004906020820280388339509192915050565b6080604051908101604052806113806113ac565b81526020016112a26113ac565b6020604051908101604052806001906020820280388339509192915050565b604080518082018252906002908290803883395091929150505600a165627a7a7230582009a9820f061e7a2b6acf34fae71349ba6e543a4bf991748c38273724a515a9150029`

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyTx is a free data retrieval call binding the contract method 0xc834d03d.
//
// Solidity: function verifyTx(a uint256[2], a_p uint256[2], b uint256[2][2], b_p uint256[2], c uint256[2], c_p uint256[2], h uint256[2], k uint256[2], input uint256[3]) constant returns(bool)
func (_Verifier *VerifierCaller) VerifyTx(opts *bind.CallOpts, a [2]*big.Int, a_p [2]*big.Int, b [2][2]*big.Int, b_p [2]*big.Int, c [2]*big.Int, c_p [2]*big.Int, h [2]*big.Int, k [2]*big.Int, input [3]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Verifier.contract.Call(opts, out, "verifyTx", a, a_p, b, b_p, c, c_p, h, k, input)
	return *ret0, err
}

// VerifyTx is a free data retrieval call binding the contract method 0xc834d03d.
//
// Solidity: function verifyTx(a uint256[2], a_p uint256[2], b uint256[2][2], b_p uint256[2], c uint256[2], c_p uint256[2], h uint256[2], k uint256[2], input uint256[3]) constant returns(bool)
func (_Verifier *VerifierSession) VerifyTx(a [2]*big.Int, a_p [2]*big.Int, b [2][2]*big.Int, b_p [2]*big.Int, c [2]*big.Int, c_p [2]*big.Int, h [2]*big.Int, k [2]*big.Int, input [3]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyTx(&_Verifier.CallOpts, a, a_p, b, b_p, c, c_p, h, k, input)
}

// VerifyTx is a free data retrieval call binding the contract method 0xc834d03d.
//
// Solidity: function verifyTx(a uint256[2], a_p uint256[2], b uint256[2][2], b_p uint256[2], c uint256[2], c_p uint256[2], h uint256[2], k uint256[2], input uint256[3]) constant returns(bool)
func (_Verifier *VerifierCallerSession) VerifyTx(a [2]*big.Int, a_p [2]*big.Int, b [2][2]*big.Int, b_p [2]*big.Int, c [2]*big.Int, c_p [2]*big.Int, h [2]*big.Int, k [2]*big.Int, input [3]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyTx(&_Verifier.CallOpts, a, a_p, b, b_p, c, c_p, h, k, input)
}

// VerifierVerifiedIterator is returned from FilterVerified and is used to iterate over the raw logs and unpacked data for Verified events raised by the Verifier contract.
type VerifierVerifiedIterator struct {
	Event *VerifierVerified // Event containing the contract specifics and raw log

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
func (it *VerifierVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierVerified)
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
		it.Event = new(VerifierVerified)
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
func (it *VerifierVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierVerified represents a Verified event raised by the Verifier contract.
type VerifierVerified struct {
	S   string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVerified is a free log retrieval operation binding the contract event 0x3f3cfdb26fb5f9f1786ab4f1a1f9cd4c0b5e726cbdfc26e495261731aad44e39.
//
// Solidity: e Verified(s string)
func (_Verifier *VerifierFilterer) FilterVerified(opts *bind.FilterOpts) (*VerifierVerifiedIterator, error) {

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return &VerifierVerifiedIterator{contract: _Verifier.contract, event: "Verified", logs: logs, sub: sub}, nil
}

// WatchVerified is a free log subscription operation binding the contract event 0x3f3cfdb26fb5f9f1786ab4f1a1f9cd4c0b5e726cbdfc26e495261731aad44e39.
//
// Solidity: e Verified(s string)
func (_Verifier *VerifierFilterer) WatchVerified(opts *bind.WatchOpts, sink chan<- *VerifierVerified) (event.Subscription, error) {

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierVerified)
				if err := _Verifier.contract.UnpackLog(event, "Verified", log); err != nil {
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

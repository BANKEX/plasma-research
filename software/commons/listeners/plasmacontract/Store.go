// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"bytes\"},{\"name\":\"protectedBlockNumber\",\"type\":\"uint256\"},{\"name\":\"protectedBlockHash\",\"type\":\"address\"}],\"name\":\"submitBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"bytes\"},{\"name\":\"protectedBlockNumber\",\"type\":\"uint256\"},{\"name\":\"protectedBlockHash\",\"type\":\"address\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"submitBlocksSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blocksLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"calculateAssetId\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAIN_COIN_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AssetDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CoinDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"length\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BlocksSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3611f3a806100cf6000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063150b7a02146100d557806347097aea146101da5780636680c5481461028b578063715018a6146103825780638ce0b5a2146103995780638da5cb5b146103c45780638f32d59b1461041b57806397feb9261461044a578063d0e30db014610497578063d29a4bf6146104a1578063dce1e772146104ee578063f25b3f991461057b578063f2fde38b146105e8578063f3b6eb911461062b575b600080fd5b3480156100e157600080fd5b50610186600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610682565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b3480156101e657600080fd5b5061027560048036038101908080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061086e565b6040518082815260200191505060405180910390f35b34801561029757600080fd5b5061036c60048036038101908080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610898565b6040518082815260200191505060405180910390f35b34801561038e57600080fd5b50610397610a98565b005b3480156103a557600080fd5b506103ae610b6a565b6040518082815260200191505060405180910390f35b3480156103d057600080fd5b506103d9610b77565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561042757600080fd5b50610430610ba0565b604051808215151515815260200191505060405180910390f35b34801561045657600080fd5b50610495600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bf7565b005b61049f611035565b005b3480156104ad57600080fd5b506104ec600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611237565b005b3480156104fa57600080fd5b50610539600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116be565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561058757600080fd5b506105a660048036038101908080359060200190929190505050611791565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105f457600080fd5b50610629600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117d4565b005b34801561063757600080fd5b506106406117f3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60003073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614151561074d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f4f6e6c79207468697320636f6e74726163742073686f756c64206465706f736981526020017f742045524337323120746f6b656e73000000000000000000000000000000000081525060400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661079033856116be565b73ffffffffffffffffffffffffffffffffffffffff1614151561081b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f45524337323120746f6b656e20776173206e6f7420657870656374656400000081525060200191505060405180910390fd5b600260006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905563150b7a027c0100000000000000000000000000000000000000000000000000000000029050949350505050565b6000610878610ba0565b151561088357600080fd5b61088f8585858561180b565b50949350505050565b6000806000878787876040516020018085815260200184805190602001908083835b6020831015156108df57805182526020820191506020810190506020830392506108ba565b6001836020036101000a0380198251168184511680821785525050505050509050018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014019450505050506040516020818303038152906040526040518082805190602001908083835b602083101515610992578051825260208201915060208101905060208303925061096d565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506109ca82611aab565b90506109d68185611b66565b73ffffffffffffffffffffffffffffffffffffffff166109f4610b77565b73ffffffffffffffffffffffffffffffffffffffff16141515610a7f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c6964207369676e617475726500000000000000000000000000000081525060200191505060405180910390fd5b610a8b8888888861180b565b9250505095945050505050565b610aa0610ba0565b1515610aab57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600180549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60008060008473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610c9757600080fd5b505af1158015610cab573d6000803e3d6000fd5b505050506040513d6020811015610cc157600080fd5b81019080805190602001909291905050509250610d013330868873ffffffffffffffffffffffffffffffffffffffff16611c5e909392919063ffffffff16565b610de9838673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610da057600080fd5b505af1158015610db4573d6000803e3d6000fd5b505050506040513d6020811015610dca57600080fd5b8101908080519060200190929190505050611d8190919063ffffffff16565b91503373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fe33e9822e3317b004d587136bab2627ea1ecfbba4eb79abddd0a56cfdd09c0e1846040518082815260200191505060405180910390a33373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f01527be533d184d44e3111afa7800fa60ced6e1b44bd025f8b457deb8ce0ce35846040518082815260200191505060405180910390a3843383604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140182815260200193505050506040516020818303038152906040526040518082805190602001908083835b602083101515610f905780518252602082019150602081019050602083039250610f6b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff167f7d6babeeae6799e032644c4c2d100c2ab47a967aec6115cf3ec5c09b818a62b6346040518082815260200191505060405180910390a23373ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff167f01527be533d184d44e3111afa7800fa60ced6e1b44bd025f8b457deb8ce0ce35346040518082815260200191505060405180910390a33334604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831015156111965780518252602082019150602081019050602083039250611171565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555050565b60008061124484846116be565b915081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508373ffffffffffffffffffffffffffffffffffffffff166342842e0e3330866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561135e57600080fd5b505af1158015611372573d6000803e3d6000fd5b50505050600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561143c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f45524337323120746f6b656e206e6f742072656365697665640000000000000081525060200191505060405180910390fd5b823373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f20d9d42fbcdd65fce5c3986b701b04ffb8e09852d04a93422dd4be124ae10a8e60405160405180910390a43373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f01527be533d184d44e3111afa7800fa60ced6e1b44bd025f8b457deb8ce0ce35856040518082815260200191505060405180910390a381843385604051602001808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018281526020019450505050506040516020818303038152906040526040518082805190602001908083835b60208310151561161a57805182526020820191506020810190506020830392506115f5565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555050505050565b60008282604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831015156117585780518252602082019150602081019050602083039250611733565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060019004905092915050565b60006001828154811015156117a257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6117dc610ba0565b15156117e757600080fd5b6117f081611da2565b50565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b6000806000806000806014895181151561182157fe5b0494506001805490508a1415156118a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c69642066726f6d496e64657800000000000000000000000000000081525060200191505060405180910390fd5b60008a148061191357508673ffffffffffffffffffffffffffffffffffffffff166001898154811015156118d057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b1515611987576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f57726f6e672070726f74656374656420626c6f636b206e756d6265720000000081525060200191505060405180910390fd5b61199f8a600180549050611d8190919063ffffffff16565b93506119b4858b611e9c90919063ffffffff16565b6001816119c19190611ebd565b508392505b84831015611a53576014830260200190506c01000000000000000000000000818a0151049150816001848c018154811015156119fe57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082806001019350506119c6565b84841015611a99576001805490507ff32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99426040518082815260200191505060405180910390a25b83850395505050505050949350505050565b60008160405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c0182600019166000191681526020019150506040516020818303038152906040526040518082805190602001908083835b602083101515611b325780518252602082019150602081019050602083039250611b0d565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050919050565b60008060008060418551141515611b805760009350611c55565b6020850151925060408501519150606085015160001a9050601b8160ff161015611bab57601b810190505b601b8160ff1614158015611bc35750601c8160ff1614155b15611bd15760009350611c55565b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015611c48573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd8484846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611d3557600080fd5b505af1158015611d49573d6000803e3d6000fd5b505050506040513d6020811015611d5f57600080fd5b81019080805190602001909291905050501515611d7b57600080fd5b50505050565b600080838311151515611d9357600080fd5b82840390508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611dde57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000808284019050838110151515611eb357600080fd5b8091505092915050565b815481835581811115611ee457818360005260206000209182019101611ee39190611ee9565b5b505050565b611f0b91905b80821115611f07576000816000905550600101611eef565b5090565b905600a165627a7a7230582048d27981ab1514041e14d43f4df8bd65ad910aa92d979c5483d58b11dc598cf20029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// MAINCOINADDRESS is a free data retrieval call binding the contract method 0xf3b6eb91.
//
// Solidity: function MAIN_COIN_ADDRESS() constant returns(address)
func (_Store *StoreCaller) MAINCOINADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "MAIN_COIN_ADDRESS")
	return *ret0, err
}

// MAINCOINADDRESS is a free data retrieval call binding the contract method 0xf3b6eb91.
//
// Solidity: function MAIN_COIN_ADDRESS() constant returns(address)
func (_Store *StoreSession) MAINCOINADDRESS() (common.Address, error) {
	return _Store.Contract.MAINCOINADDRESS(&_Store.CallOpts)
}

// MAINCOINADDRESS is a free data retrieval call binding the contract method 0xf3b6eb91.
//
// Solidity: function MAIN_COIN_ADDRESS() constant returns(address)
func (_Store *StoreCallerSession) MAINCOINADDRESS() (common.Address, error) {
	return _Store.Contract.MAINCOINADDRESS(&_Store.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_Store *StoreCaller) Blocks(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "blocks", i)
	return *ret0, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_Store *StoreSession) Blocks(i *big.Int) (common.Address, error) {
	return _Store.Contract.Blocks(&_Store.CallOpts, i)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_Store *StoreCallerSession) Blocks(i *big.Int) (common.Address, error) {
	return _Store.Contract.Blocks(&_Store.CallOpts, i)
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_Store *StoreCaller) BlocksLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "blocksLength")
	return *ret0, err
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_Store *StoreSession) BlocksLength() (*big.Int, error) {
	return _Store.Contract.BlocksLength(&_Store.CallOpts)
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_Store *StoreCallerSession) BlocksLength() (*big.Int, error) {
	return _Store.Contract.BlocksLength(&_Store.CallOpts)
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xdce1e772.
//
// Solidity: function calculateAssetId(token address, tokenId uint256) constant returns(address)
func (_Store *StoreCaller) CalculateAssetId(opts *bind.CallOpts, token common.Address, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "calculateAssetId", token, tokenId)
	return *ret0, err
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xdce1e772.
//
// Solidity: function calculateAssetId(token address, tokenId uint256) constant returns(address)
func (_Store *StoreSession) CalculateAssetId(token common.Address, tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.CalculateAssetId(&_Store.CallOpts, token, tokenId)
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xdce1e772.
//
// Solidity: function calculateAssetId(token address, tokenId uint256) constant returns(address)
func (_Store *StoreCallerSession) CalculateAssetId(token common.Address, tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.CalculateAssetId(&_Store.CallOpts, token, tokenId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Store *StoreCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Store *StoreSession) IsOwner() (bool, error) {
	return _Store.Contract.IsOwner(&_Store.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Store *StoreCallerSession) IsOwner() (bool, error) {
	return _Store.Contract.IsOwner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Store *StoreTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Store *StoreSession) Deposit() (*types.Transaction, error) {
	return _Store.Contract.Deposit(&_Store.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Store *StoreTransactorSession) Deposit() (*types.Transaction, error) {
	return _Store.Contract.Deposit(&_Store.TransactOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(token address, amount uint256) returns()
func (_Store *StoreTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "depositERC20", token, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(token address, amount uint256) returns()
func (_Store *StoreSession) DepositERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DepositERC20(&_Store.TransactOpts, token, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(token address, amount uint256) returns()
func (_Store *StoreTransactorSession) DepositERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DepositERC20(&_Store.TransactOpts, token, amount)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(token address, tokenId uint256) returns()
func (_Store *StoreTransactor) DepositERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "depositERC721", token, tokenId)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(token address, tokenId uint256) returns()
func (_Store *StoreSession) DepositERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DepositERC721(&_Store.TransactOpts, token, tokenId)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(token address, tokenId uint256) returns()
func (_Store *StoreTransactorSession) DepositERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DepositERC721(&_Store.TransactOpts, token, tokenId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(operator address,  address, tokenId uint256,  bytes) returns(bytes4)
func (_Store *StoreTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "onERC721Received", operator, arg1, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(operator address,  address, tokenId uint256,  bytes) returns(bytes4)
func (_Store *StoreSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Store.Contract.OnERC721Received(&_Store.TransactOpts, operator, arg1, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(operator address,  address, tokenId uint256,  bytes) returns(bytes4)
func (_Store *StoreTransactorSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Store.Contract.OnERC721Received(&_Store.TransactOpts, operator, arg1, tokenId, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x47097aea.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address) returns(uint256)
func (_Store *StoreTransactor) SubmitBlocks(opts *bind.TransactOpts, fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "submitBlocks", fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x47097aea.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address) returns(uint256)
func (_Store *StoreSession) SubmitBlocks(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocks(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x47097aea.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address) returns(uint256)
func (_Store *StoreTransactorSession) SubmitBlocks(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocks(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x6680c548.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address, rsv bytes) returns(uint256)
func (_Store *StoreTransactor) SubmitBlocksSigned(opts *bind.TransactOpts, fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address, rsv []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "submitBlocksSigned", fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash, rsv)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x6680c548.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address, rsv bytes) returns(uint256)
func (_Store *StoreSession) SubmitBlocksSigned(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address, rsv []byte) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocksSigned(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash, rsv)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x6680c548.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address, rsv bytes) returns(uint256)
func (_Store *StoreTransactorSession) SubmitBlocksSigned(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address, rsv []byte) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocksSigned(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash, rsv)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// StoreAssetDepositedIterator is returned from FilterAssetDeposited and is used to iterate over the raw logs and unpacked data for AssetDeposited events raised by the Store contract.
type StoreAssetDepositedIterator struct {
	Event *StoreAssetDeposited // Event containing the contract specifics and raw log

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
func (it *StoreAssetDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAssetDeposited)
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
		it.Event = new(StoreAssetDeposited)
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
func (it *StoreAssetDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAssetDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAssetDeposited represents a AssetDeposited event raised by the Store contract.
type StoreAssetDeposited struct {
	Token  common.Address
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetDeposited is a free log retrieval operation binding the contract event 0x01527be533d184d44e3111afa7800fa60ced6e1b44bd025f8b457deb8ce0ce35.
//
// Solidity: e AssetDeposited(token indexed address, who indexed address, amount uint256)
func (_Store *StoreFilterer) FilterAssetDeposited(opts *bind.FilterOpts, token []common.Address, who []common.Address) (*StoreAssetDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "AssetDeposited", tokenRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &StoreAssetDepositedIterator{contract: _Store.contract, event: "AssetDeposited", logs: logs, sub: sub}, nil
}

// WatchAssetDeposited is a free log subscription operation binding the contract event 0x01527be533d184d44e3111afa7800fa60ced6e1b44bd025f8b457deb8ce0ce35.
//
// Solidity: e AssetDeposited(token indexed address, who indexed address, amount uint256)
func (_Store *StoreFilterer) WatchAssetDeposited(opts *bind.WatchOpts, sink chan<- *StoreAssetDeposited, token []common.Address, who []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "AssetDeposited", tokenRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAssetDeposited)
				if err := _Store.contract.UnpackLog(event, "AssetDeposited", log); err != nil {
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

// StoreBlocksSubmittedIterator is returned from FilterBlocksSubmitted and is used to iterate over the raw logs and unpacked data for BlocksSubmitted events raised by the Store contract.
type StoreBlocksSubmittedIterator struct {
	Event *StoreBlocksSubmitted // Event containing the contract specifics and raw log

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
func (it *StoreBlocksSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBlocksSubmitted)
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
		it.Event = new(StoreBlocksSubmitted)
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
func (it *StoreBlocksSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBlocksSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBlocksSubmitted represents a BlocksSubmitted event raised by the Store contract.
type StoreBlocksSubmitted struct {
	Length *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlocksSubmitted is a free log retrieval operation binding the contract event 0xf32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99.
//
// Solidity: e BlocksSubmitted(length indexed uint256, time uint256)
func (_Store *StoreFilterer) FilterBlocksSubmitted(opts *bind.FilterOpts, length []*big.Int) (*StoreBlocksSubmittedIterator, error) {

	var lengthRule []interface{}
	for _, lengthItem := range length {
		lengthRule = append(lengthRule, lengthItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "BlocksSubmitted", lengthRule)
	if err != nil {
		return nil, err
	}
	return &StoreBlocksSubmittedIterator{contract: _Store.contract, event: "BlocksSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlocksSubmitted is a free log subscription operation binding the contract event 0xf32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99.
//
// Solidity: e BlocksSubmitted(length indexed uint256, time uint256)
func (_Store *StoreFilterer) WatchBlocksSubmitted(opts *bind.WatchOpts, sink chan<- *StoreBlocksSubmitted, length []*big.Int) (event.Subscription, error) {

	var lengthRule []interface{}
	for _, lengthItem := range length {
		lengthRule = append(lengthRule, lengthItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "BlocksSubmitted", lengthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBlocksSubmitted)
				if err := _Store.contract.UnpackLog(event, "BlocksSubmitted", log); err != nil {
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

// StoreCoinDepositedIterator is returned from FilterCoinDeposited and is used to iterate over the raw logs and unpacked data for CoinDeposited events raised by the Store contract.
type StoreCoinDepositedIterator struct {
	Event *StoreCoinDeposited // Event containing the contract specifics and raw log

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
func (it *StoreCoinDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCoinDeposited)
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
		it.Event = new(StoreCoinDeposited)
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
func (it *StoreCoinDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCoinDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCoinDeposited represents a CoinDeposited event raised by the Store contract.
type StoreCoinDeposited struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCoinDeposited is a free log retrieval operation binding the contract event 0x7d6babeeae6799e032644c4c2d100c2ab47a967aec6115cf3ec5c09b818a62b6.
//
// Solidity: e CoinDeposited(who indexed address, amount uint256)
func (_Store *StoreFilterer) FilterCoinDeposited(opts *bind.FilterOpts, who []common.Address) (*StoreCoinDepositedIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "CoinDeposited", whoRule)
	if err != nil {
		return nil, err
	}
	return &StoreCoinDepositedIterator{contract: _Store.contract, event: "CoinDeposited", logs: logs, sub: sub}, nil
}

// WatchCoinDeposited is a free log subscription operation binding the contract event 0x7d6babeeae6799e032644c4c2d100c2ab47a967aec6115cf3ec5c09b818a62b6.
//
// Solidity: e CoinDeposited(who indexed address, amount uint256)
func (_Store *StoreFilterer) WatchCoinDeposited(opts *bind.WatchOpts, sink chan<- *StoreCoinDeposited, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "CoinDeposited", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCoinDeposited)
				if err := _Store.contract.UnpackLog(event, "CoinDeposited", log); err != nil {
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

// StoreERC20DepositedIterator is returned from FilterERC20Deposited and is used to iterate over the raw logs and unpacked data for ERC20Deposited events raised by the Store contract.
type StoreERC20DepositedIterator struct {
	Event *StoreERC20Deposited // Event containing the contract specifics and raw log

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
func (it *StoreERC20DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreERC20Deposited)
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
		it.Event = new(StoreERC20Deposited)
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
func (it *StoreERC20DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreERC20DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreERC20Deposited represents a ERC20Deposited event raised by the Store contract.
type StoreERC20Deposited struct {
	Token  common.Address
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20Deposited is a free log retrieval operation binding the contract event 0xe33e9822e3317b004d587136bab2627ea1ecfbba4eb79abddd0a56cfdd09c0e1.
//
// Solidity: e ERC20Deposited(token indexed address, who indexed address, amount uint256)
func (_Store *StoreFilterer) FilterERC20Deposited(opts *bind.FilterOpts, token []common.Address, who []common.Address) (*StoreERC20DepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ERC20Deposited", tokenRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &StoreERC20DepositedIterator{contract: _Store.contract, event: "ERC20Deposited", logs: logs, sub: sub}, nil
}

// WatchERC20Deposited is a free log subscription operation binding the contract event 0xe33e9822e3317b004d587136bab2627ea1ecfbba4eb79abddd0a56cfdd09c0e1.
//
// Solidity: e ERC20Deposited(token indexed address, who indexed address, amount uint256)
func (_Store *StoreFilterer) WatchERC20Deposited(opts *bind.WatchOpts, sink chan<- *StoreERC20Deposited, token []common.Address, who []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ERC20Deposited", tokenRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreERC20Deposited)
				if err := _Store.contract.UnpackLog(event, "ERC20Deposited", log); err != nil {
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

// StoreERC721DepositedIterator is returned from FilterERC721Deposited and is used to iterate over the raw logs and unpacked data for ERC721Deposited events raised by the Store contract.
type StoreERC721DepositedIterator struct {
	Event *StoreERC721Deposited // Event containing the contract specifics and raw log

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
func (it *StoreERC721DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreERC721Deposited)
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
		it.Event = new(StoreERC721Deposited)
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
func (it *StoreERC721DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreERC721DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreERC721Deposited represents a ERC721Deposited event raised by the Store contract.
type StoreERC721Deposited struct {
	Token   common.Address
	Who     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC721Deposited is a free log retrieval operation binding the contract event 0x20d9d42fbcdd65fce5c3986b701b04ffb8e09852d04a93422dd4be124ae10a8e.
//
// Solidity: e ERC721Deposited(token indexed address, who indexed address, tokenId indexed uint256)
func (_Store *StoreFilterer) FilterERC721Deposited(opts *bind.FilterOpts, token []common.Address, who []common.Address, tokenId []*big.Int) (*StoreERC721DepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ERC721Deposited", tokenRule, whoRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreERC721DepositedIterator{contract: _Store.contract, event: "ERC721Deposited", logs: logs, sub: sub}, nil
}

// WatchERC721Deposited is a free log subscription operation binding the contract event 0x20d9d42fbcdd65fce5c3986b701b04ffb8e09852d04a93422dd4be124ae10a8e.
//
// Solidity: e ERC721Deposited(token indexed address, who indexed address, tokenId indexed uint256)
func (_Store *StoreFilterer) WatchERC721Deposited(opts *bind.WatchOpts, sink chan<- *StoreERC721Deposited, token []common.Address, who []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ERC721Deposited", tokenRule, whoRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreERC721Deposited)
				if err := _Store.contract.UnpackLog(event, "ERC721Deposited", log); err != nil {
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

// StoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Store contract.
type StoreOwnershipTransferredIterator struct {
	Event *StoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOwnershipTransferred)
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
		it.Event = new(StoreOwnershipTransferred)
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
func (it *StoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOwnershipTransferred represents a OwnershipTransferred event raised by the Store contract.
type StoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Store *StoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOwnershipTransferredIterator{contract: _Store.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Store *StoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOwnershipTransferred)
				if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

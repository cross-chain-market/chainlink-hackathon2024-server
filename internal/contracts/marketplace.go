// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"CollectionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ListingBought\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"addCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"buyListing\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600260015534801561001557600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611001806100666000396000f3fe60806040526004361061008a5760003560e01c8063ba07feaa11610059578063ba07feaa1461013e578063bec3fa171461017b578063cf52ffd2146101a4578063ddca3f43146101c0578063f2fde38b146101eb57610091565b806312065fe01461009657806319f101fa146100c15780638da5cb5b146100ea5780639012c4a81461011557610091565b3661009157005b600080fd5b3480156100a257600080fd5b506100ab610214565b6040516100b8919061094d565b60405180910390f35b3480156100cd57600080fd5b506100e860048036038101906100e391906109f7565b61021c565b005b3480156100f657600080fd5b506100ff6102b5565b60405161010c9190610a46565b60405180910390f35b34801561012157600080fd5b5061013c60048036038101906101379190610a61565b6102db565b005b34801561014a57600080fd5b5061016560048036038101906101609190610a61565b6103fc565b6040516101729190610a46565b60405180910390f35b34801561018757600080fd5b506101a2600480360381019061019d91906109f7565b61042f565b005b6101be60048036038101906101b99190610a8e565b610517565b005b3480156101cc57600080fd5b506101d561085a565b6040516101e2919061094d565b60405180910390f35b3480156101f757600080fd5b50610212600480360381019061020d9190610af5565b610860565b005b600047905090565b8160008083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff167fafefbf76b73bee0197e5579b75960a45d7b20800ba84b19adfd5e4ae6b8d0b9060405160405180910390a35050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461036b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036290610b7f565b60405180910390fd5b60008111801561037c575060648111155b6103bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b290610beb565b60405180910390fd5b806001819055507f8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76816040516103f1919061094d565b60405180910390a150565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b690610b7f565b60405180910390fd5b478111156104cc57600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610512573d6000803e3d6000fd5b505050565b6000341161055a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055190610c7d565b60405180910390fd5b6000821161059d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059490610d0f565b60405180910390fd5b600080600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610643576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063a90610d7b565b60405180910390fd5b60008190508173ffffffffffffffffffffffffffffffffffffffff16635c3465bc8487876040518463ffffffff1660e01b815260040161068593929190610d9b565b600060405180830381600087803b15801561069f57600080fd5b505af11580156106b3573d6000803e3d6000fd5b5050505060006064600154346106c99190610e01565b6106d39190610e72565b9050600081346106e39190610ea3565b905060008373ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610732573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107569190610eec565b73ffffffffffffffffffffffffffffffffffffffff168260405161077990610f4a565b60006040518083038185875af1925050503d80600081146107b6576040519150601f19603f3d011682016040523d82523d6000602084013e6107bb565b606091505b50509050806107ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f690610fab565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff16888a7f67e902d908bccbae91b916b3c0c92ba6d02d4a12724ba9770dd0481cfd550f1a8a604051610847919061094d565b60405180910390a4505050505050505050565b60015481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e790610b7f565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000819050919050565b61094781610934565b82525050565b6000602082019050610962600083018461093e565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109988261096d565b9050919050565b6109a88161098d565b81146109b357600080fd5b50565b6000813590506109c58161099f565b92915050565b6109d481610934565b81146109df57600080fd5b50565b6000813590506109f1816109cb565b92915050565b60008060408385031215610a0e57610a0d610968565b5b6000610a1c858286016109b6565b9250506020610a2d858286016109e2565b9150509250929050565b610a408161098d565b82525050565b6000602082019050610a5b6000830184610a37565b92915050565b600060208284031215610a7757610a76610968565b5b6000610a85848285016109e2565b91505092915050565b60008060008060808587031215610aa857610aa7610968565b5b6000610ab6878288016109e2565b9450506020610ac7878288016109e2565b9350506040610ad8878288016109e2565b9250506060610ae9878288016109b6565b91505092959194509250565b600060208284031215610b0b57610b0a610968565b5b6000610b19848285016109b6565b91505092915050565b600082825260208201905092915050565b7f6e6f74206f776e6572206f6620636f6e74726163740000000000000000000000600082015250565b6000610b69601583610b22565b9150610b7482610b33565b602082019050919050565b60006020820190508181036000830152610b9881610b5c565b9050919050565b7f666565206973206e6f7420696e2072616e676500000000000000000000000000600082015250565b6000610bd5601383610b22565b9150610be082610b9f565b602082019050919050565b60006020820190508181036000830152610c0481610bc8565b9050919050565b7f7061796d656e742073686f756c642062652067726561746572207468656e207a60008201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b6000610c67602383610b22565b9150610c7282610c0b565b604082019050919050565b60006020820190508181036000830152610c9681610c5a565b9050919050565b7f616d6f756e74206f66206c697374696e67732073686f756c642062652067726560008201527f61746572207468656e207a65726f000000000000000000000000000000000000602082015250565b6000610cf9602e83610b22565b9150610d0482610c9d565b604082019050919050565b60006020820190508181036000830152610d2881610cec565b9050919050565b7f436f6c6c656374696f6e20646f6573206e6f7420657869737400000000000000600082015250565b6000610d65601983610b22565b9150610d7082610d2f565b602082019050919050565b60006020820190508181036000830152610d9481610d58565b9050919050565b6000606082019050610db06000830186610a37565b610dbd602083018561093e565b610dca604083018461093e565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e0c82610934565b9150610e1783610934565b9250828202610e2581610934565b91508282048414831517610e3c57610e3b610dd2565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610e7d82610934565b9150610e8883610934565b925082610e9857610e97610e43565b5b828204905092915050565b6000610eae82610934565b9150610eb983610934565b9250828203905081811115610ed157610ed0610dd2565b5b92915050565b600081519050610ee68161099f565b92915050565b600060208284031215610f0257610f01610968565b5b6000610f1084828501610ed7565b91505092915050565b600081905092915050565b50565b6000610f34600083610f19565b9150610f3f82610f24565b600082019050919050565b6000610f5582610f27565b9150819050919050565b7f5061796d656e7420746f206f776e6572206661696c6564000000000000000000600082015250565b6000610f95601783610b22565b9150610fa082610f5f565b602082019050919050565b60006020820190508181036000830152610fc481610f88565b905091905056fea2646970667358221220cbd4a71ad4e9c139f9ab36be78414d2db469fa7f68a8a4d84a817eb184e71ad864736f6c63430008140033",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// MarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketplaceMetaData.Bin instead.
var MarketplaceBin = MarketplaceMetaData.Bin

// DeployMarketplace deploys a new Ethereum contract, binding an instance of Marketplace to it.
func DeployMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Marketplace, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketplaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Marketplace *MarketplaceCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Marketplace *MarketplaceSession) Fee() (*big.Int, error) {
	return _Marketplace.Contract.Fee(&_Marketplace.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) Fee() (*big.Int, error) {
	return _Marketplace.Contract.Fee(&_Marketplace.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetBalance() (*big.Int, error) {
	return _Marketplace.Contract.GetBalance(&_Marketplace.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetBalance() (*big.Int, error) {
	return _Marketplace.Contract.GetBalance(&_Marketplace.CallOpts)
}

// MarketCollection is a free data retrieval call binding the contract method 0xba07feaa.
//
// Solidity: function marketCollection(uint256 ) view returns(address)
func (_Marketplace *MarketplaceCaller) MarketCollection(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "marketCollection", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketCollection is a free data retrieval call binding the contract method 0xba07feaa.
//
// Solidity: function marketCollection(uint256 ) view returns(address)
func (_Marketplace *MarketplaceSession) MarketCollection(arg0 *big.Int) (common.Address, error) {
	return _Marketplace.Contract.MarketCollection(&_Marketplace.CallOpts, arg0)
}

// MarketCollection is a free data retrieval call binding the contract method 0xba07feaa.
//
// Solidity: function marketCollection(uint256 ) view returns(address)
func (_Marketplace *MarketplaceCallerSession) MarketCollection(arg0 *big.Int) (common.Address, error) {
	return _Marketplace.Contract.MarketCollection(&_Marketplace.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// AddCollection is a paid mutator transaction binding the contract method 0x19f101fa.
//
// Solidity: function addCollection(address collectionAddress, uint256 collectionId) returns()
func (_Marketplace *MarketplaceTransactor) AddCollection(opts *bind.TransactOpts, collectionAddress common.Address, collectionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "addCollection", collectionAddress, collectionId)
}

// AddCollection is a paid mutator transaction binding the contract method 0x19f101fa.
//
// Solidity: function addCollection(address collectionAddress, uint256 collectionId) returns()
func (_Marketplace *MarketplaceSession) AddCollection(collectionAddress common.Address, collectionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.AddCollection(&_Marketplace.TransactOpts, collectionAddress, collectionId)
}

// AddCollection is a paid mutator transaction binding the contract method 0x19f101fa.
//
// Solidity: function addCollection(address collectionAddress, uint256 collectionId) returns()
func (_Marketplace *MarketplaceTransactorSession) AddCollection(collectionAddress common.Address, collectionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.AddCollection(&_Marketplace.TransactOpts, collectionAddress, collectionId)
}

// BuyListing is a paid mutator transaction binding the contract method 0xcf52ffd2.
//
// Solidity: function buyListing(uint256 collectionId, uint256 listingId, uint256 amount, address to) payable returns()
func (_Marketplace *MarketplaceTransactor) BuyListing(opts *bind.TransactOpts, collectionId *big.Int, listingId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyListing", collectionId, listingId, amount, to)
}

// BuyListing is a paid mutator transaction binding the contract method 0xcf52ffd2.
//
// Solidity: function buyListing(uint256 collectionId, uint256 listingId, uint256 amount, address to) payable returns()
func (_Marketplace *MarketplaceSession) BuyListing(collectionId *big.Int, listingId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyListing(&_Marketplace.TransactOpts, collectionId, listingId, amount, to)
}

// BuyListing is a paid mutator transaction binding the contract method 0xcf52ffd2.
//
// Solidity: function buyListing(uint256 collectionId, uint256 listingId, uint256 amount, address to) payable returns()
func (_Marketplace *MarketplaceTransactorSession) BuyListing(collectionId *big.Int, listingId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyListing(&_Marketplace.TransactOpts, collectionId, listingId, amount, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address account) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address account) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address account) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, account)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(address to, uint256 amount) returns()
func (_Marketplace *MarketplaceTransactor) TransferTokens(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferTokens", to, amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(address to, uint256 amount) returns()
func (_Marketplace *MarketplaceSession) TransferTokens(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferTokens(&_Marketplace.TransactOpts, to, amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(address to, uint256 amount) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferTokens(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferTokens(&_Marketplace.TransactOpts, to, amount)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 _fee) returns()
func (_Marketplace *MarketplaceTransactor) UpdateFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateFee", _fee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 _fee) returns()
func (_Marketplace *MarketplaceSession) UpdateFee(_fee *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateFee(&_Marketplace.TransactOpts, _fee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 _fee) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateFee(_fee *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateFee(&_Marketplace.TransactOpts, _fee)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// MarketplaceCollectionAddedIterator is returned from FilterCollectionAdded and is used to iterate over the raw logs and unpacked data for CollectionAdded events raised by the Marketplace contract.
type MarketplaceCollectionAddedIterator struct {
	Event *MarketplaceCollectionAdded // Event containing the contract specifics and raw log

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
func (it *MarketplaceCollectionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCollectionAdded)
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
		it.Event = new(MarketplaceCollectionAdded)
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
func (it *MarketplaceCollectionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCollectionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCollectionAdded represents a CollectionAdded event raised by the Marketplace contract.
type MarketplaceCollectionAdded struct {
	CollectionAddress common.Address
	CollectionId      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollectionAdded is a free log retrieval operation binding the contract event 0xafefbf76b73bee0197e5579b75960a45d7b20800ba84b19adfd5e4ae6b8d0b90.
//
// Solidity: event CollectionAdded(address indexed collectionAddress, uint256 indexed collectionId)
func (_Marketplace *MarketplaceFilterer) FilterCollectionAdded(opts *bind.FilterOpts, collectionAddress []common.Address, collectionId []*big.Int) (*MarketplaceCollectionAddedIterator, error) {

	var collectionAddressRule []interface{}
	for _, collectionAddressItem := range collectionAddress {
		collectionAddressRule = append(collectionAddressRule, collectionAddressItem)
	}
	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CollectionAdded", collectionAddressRule, collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCollectionAddedIterator{contract: _Marketplace.contract, event: "CollectionAdded", logs: logs, sub: sub}, nil
}

// WatchCollectionAdded is a free log subscription operation binding the contract event 0xafefbf76b73bee0197e5579b75960a45d7b20800ba84b19adfd5e4ae6b8d0b90.
//
// Solidity: event CollectionAdded(address indexed collectionAddress, uint256 indexed collectionId)
func (_Marketplace *MarketplaceFilterer) WatchCollectionAdded(opts *bind.WatchOpts, sink chan<- *MarketplaceCollectionAdded, collectionAddress []common.Address, collectionId []*big.Int) (event.Subscription, error) {

	var collectionAddressRule []interface{}
	for _, collectionAddressItem := range collectionAddress {
		collectionAddressRule = append(collectionAddressRule, collectionAddressItem)
	}
	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CollectionAdded", collectionAddressRule, collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCollectionAdded)
				if err := _Marketplace.contract.UnpackLog(event, "CollectionAdded", log); err != nil {
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

// ParseCollectionAdded is a log parse operation binding the contract event 0xafefbf76b73bee0197e5579b75960a45d7b20800ba84b19adfd5e4ae6b8d0b90.
//
// Solidity: event CollectionAdded(address indexed collectionAddress, uint256 indexed collectionId)
func (_Marketplace *MarketplaceFilterer) ParseCollectionAdded(log types.Log) (*MarketplaceCollectionAdded, error) {
	event := new(MarketplaceCollectionAdded)
	if err := _Marketplace.contract.UnpackLog(event, "CollectionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the Marketplace contract.
type MarketplaceFeeUpdatedIterator struct {
	Event *MarketplaceFeeUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceFeeUpdated)
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
		it.Event = new(MarketplaceFeeUpdated)
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
func (it *MarketplaceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceFeeUpdated represents a FeeUpdated event raised by the Marketplace contract.
type MarketplaceFeeUpdated struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 _fee)
func (_Marketplace *MarketplaceFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*MarketplaceFeeUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceFeeUpdatedIterator{contract: _Marketplace.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 _fee)
func (_Marketplace *MarketplaceFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceFeeUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 _fee)
func (_Marketplace *MarketplaceFilterer) ParseFeeUpdated(log types.Log) (*MarketplaceFeeUpdated, error) {
	event := new(MarketplaceFeeUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingBoughtIterator is returned from FilterListingBought and is used to iterate over the raw logs and unpacked data for ListingBought events raised by the Marketplace contract.
type MarketplaceListingBoughtIterator struct {
	Event *MarketplaceListingBought // Event containing the contract specifics and raw log

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
func (it *MarketplaceListingBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingBought)
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
		it.Event = new(MarketplaceListingBought)
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
func (it *MarketplaceListingBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingBought represents a ListingBought event raised by the Marketplace contract.
type MarketplaceListingBought struct {
	CollectionId *big.Int
	ListingId    *big.Int
	Amount       *big.Int
	To           common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterListingBought is a free log retrieval operation binding the contract event 0x67e902d908bccbae91b916b3c0c92ba6d02d4a12724ba9770dd0481cfd550f1a.
//
// Solidity: event ListingBought(uint256 indexed collectionId, uint256 indexed listingId, uint256 amount, address indexed to)
func (_Marketplace *MarketplaceFilterer) FilterListingBought(opts *bind.FilterOpts, collectionId []*big.Int, listingId []*big.Int, to []common.Address) (*MarketplaceListingBoughtIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingBought", collectionIdRule, listingIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingBoughtIterator{contract: _Marketplace.contract, event: "ListingBought", logs: logs, sub: sub}, nil
}

// WatchListingBought is a free log subscription operation binding the contract event 0x67e902d908bccbae91b916b3c0c92ba6d02d4a12724ba9770dd0481cfd550f1a.
//
// Solidity: event ListingBought(uint256 indexed collectionId, uint256 indexed listingId, uint256 amount, address indexed to)
func (_Marketplace *MarketplaceFilterer) WatchListingBought(opts *bind.WatchOpts, sink chan<- *MarketplaceListingBought, collectionId []*big.Int, listingId []*big.Int, to []common.Address) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingBought", collectionIdRule, listingIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingBought)
				if err := _Marketplace.contract.UnpackLog(event, "ListingBought", log); err != nil {
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

// ParseListingBought is a log parse operation binding the contract event 0x67e902d908bccbae91b916b3c0c92ba6d02d4a12724ba9770dd0481cfd550f1a.
//
// Solidity: event ListingBought(uint256 indexed collectionId, uint256 indexed listingId, uint256 amount, address indexed to)
func (_Marketplace *MarketplaceFilterer) ParseListingBought(log types.Log) (*MarketplaceListingBought, error) {
	event := new(MarketplaceListingBought)
	if err := _Marketplace.contract.UnpackLog(event, "ListingBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

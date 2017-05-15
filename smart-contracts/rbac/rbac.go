// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package rbac

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RBACContractABI is the input ABI used to generate the binding from.
const RBACContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"inspectorAddress\",\"type\":\"address\"}],\"name\":\"isInspector\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"inspectorAddress\",\"type\":\"address\"}],\"name\":\"removeInspector\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"inspectorAddress\",\"type\":\"address\"}],\"name\":\"addInspector\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"farmerAddress\",\"type\":\"address\"}],\"name\":\"addFarmer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"farmerAddress\",\"type\":\"address\"}],\"name\":\"isFarmer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"farmerAddress\",\"type\":\"address\"}],\"name\":\"removeFarmer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// RBACContractBin is the compiled bytecode used for deploying new contracts.
const RBACContractBin = `6060604052341561000c57fe5b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b6108338061005f6000396000f30060606040523615610097576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631785f53c1461009957806324d7806c146100cf5780632cdad41c1461011d578063704802751461016b5780637c70e791146101a15780637e458492146101d757806380c3f96d1461020d578063b942906914610243578063e6bf3fdc14610291575bfe5b34156100a157fe5b6100cd600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506102c7565b005b34156100d757fe5b610103600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610377565b604051808215151515815260200191505060405180910390f35b341561012557fe5b610151600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103ce565b604051808215151515815260200191505060405180910390f35b341561017357fe5b61019f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610425565b005b34156101a957fe5b6101d5600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104de565b005b34156101df57fe5b61020b600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061058e565b005b341561021557fe5b610241600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610647565b005b341561024b57fe5b610277600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610700565b604051808215151515815260200191505060405180910390f35b341561029957fe5b6102c5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610757565b005b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156103245760006000fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b50565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156104825760006000fd5b6001600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614151561053b5760006000fd5b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156105eb5760006000fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156106a45760006000fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156107b45760006000fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b505600a165627a7a7230582026eba01b4a897f95c063ae5b1f46517b1d702d9651b00ef69fdc289c0d6e6b2b0029`

// DeployRBACContract deploys a new Ethereum contract, binding an instance of RBACContract to it.
func DeployRBACContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RBACContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RBACContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RBACContract{RBACContractCaller: RBACContractCaller{contract: contract}, RBACContractTransactor: RBACContractTransactor{contract: contract}}, nil
}

// RBACContract is an auto generated Go binding around an Ethereum contract.
type RBACContract struct {
	RBACContractCaller     // Read-only binding to the contract
	RBACContractTransactor // Write-only binding to the contract
}

// RBACContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RBACContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RBACContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RBACContractSession struct {
	Contract     *RBACContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RBACContractCallerSession struct {
	Contract *RBACContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RBACContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RBACContractTransactorSession struct {
	Contract     *RBACContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RBACContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RBACContractRaw struct {
	Contract *RBACContract // Generic contract binding to access the raw methods on
}

// RBACContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RBACContractCallerRaw struct {
	Contract *RBACContractCaller // Generic read-only contract binding to access the raw methods on
}

// RBACContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RBACContractTransactorRaw struct {
	Contract *RBACContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRBACContract creates a new instance of RBACContract, bound to a specific deployed contract.
func NewRBACContract(address common.Address, backend bind.ContractBackend) (*RBACContract, error) {
	contract, err := bindRBACContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RBACContract{RBACContractCaller: RBACContractCaller{contract: contract}, RBACContractTransactor: RBACContractTransactor{contract: contract}}, nil
}

// NewRBACContractCaller creates a new read-only instance of RBACContract, bound to a specific deployed contract.
func NewRBACContractCaller(address common.Address, caller bind.ContractCaller) (*RBACContractCaller, error) {
	contract, err := bindRBACContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RBACContractCaller{contract: contract}, nil
}

// NewRBACContractTransactor creates a new write-only instance of RBACContract, bound to a specific deployed contract.
func NewRBACContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RBACContractTransactor, error) {
	contract, err := bindRBACContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RBACContractTransactor{contract: contract}, nil
}

// bindRBACContract binds a generic wrapper to an already deployed contract.
func bindRBACContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBACContract *RBACContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBACContract.Contract.RBACContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBACContract *RBACContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBACContract.Contract.RBACContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBACContract *RBACContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBACContract.Contract.RBACContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBACContract *RBACContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBACContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBACContract *RBACContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBACContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBACContract *RBACContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBACContract.Contract.contract.Transact(opts, method, params...)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(adminAddress address) constant returns(bool)
func (_RBACContract *RBACContractCaller) IsAdmin(opts *bind.CallOpts, adminAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBACContract.contract.Call(opts, out, "isAdmin", adminAddress)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(adminAddress address) constant returns(bool)
func (_RBACContract *RBACContractSession) IsAdmin(adminAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsAdmin(&_RBACContract.CallOpts, adminAddress)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(adminAddress address) constant returns(bool)
func (_RBACContract *RBACContractCallerSession) IsAdmin(adminAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsAdmin(&_RBACContract.CallOpts, adminAddress)
}

// IsFarmer is a free data retrieval call binding the contract method 0xb9429069.
//
// Solidity: function isFarmer(farmerAddress address) constant returns(bool)
func (_RBACContract *RBACContractCaller) IsFarmer(opts *bind.CallOpts, farmerAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBACContract.contract.Call(opts, out, "isFarmer", farmerAddress)
	return *ret0, err
}

// IsFarmer is a free data retrieval call binding the contract method 0xb9429069.
//
// Solidity: function isFarmer(farmerAddress address) constant returns(bool)
func (_RBACContract *RBACContractSession) IsFarmer(farmerAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsFarmer(&_RBACContract.CallOpts, farmerAddress)
}

// IsFarmer is a free data retrieval call binding the contract method 0xb9429069.
//
// Solidity: function isFarmer(farmerAddress address) constant returns(bool)
func (_RBACContract *RBACContractCallerSession) IsFarmer(farmerAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsFarmer(&_RBACContract.CallOpts, farmerAddress)
}

// IsInspector is a free data retrieval call binding the contract method 0x2cdad41c.
//
// Solidity: function isInspector(inspectorAddress address) constant returns(bool)
func (_RBACContract *RBACContractCaller) IsInspector(opts *bind.CallOpts, inspectorAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBACContract.contract.Call(opts, out, "isInspector", inspectorAddress)
	return *ret0, err
}

// IsInspector is a free data retrieval call binding the contract method 0x2cdad41c.
//
// Solidity: function isInspector(inspectorAddress address) constant returns(bool)
func (_RBACContract *RBACContractSession) IsInspector(inspectorAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsInspector(&_RBACContract.CallOpts, inspectorAddress)
}

// IsInspector is a free data retrieval call binding the contract method 0x2cdad41c.
//
// Solidity: function isInspector(inspectorAddress address) constant returns(bool)
func (_RBACContract *RBACContractCallerSession) IsInspector(inspectorAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsInspector(&_RBACContract.CallOpts, inspectorAddress)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(adminAddress address) returns()
func (_RBACContract *RBACContractTransactor) AddAdmin(opts *bind.TransactOpts, adminAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "addAdmin", adminAddress)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(adminAddress address) returns()
func (_RBACContract *RBACContractSession) AddAdmin(adminAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddAdmin(&_RBACContract.TransactOpts, adminAddress)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(adminAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) AddAdmin(adminAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddAdmin(&_RBACContract.TransactOpts, adminAddress)
}

// AddFarmer is a paid mutator transaction binding the contract method 0x80c3f96d.
//
// Solidity: function addFarmer(farmerAddress address) returns()
func (_RBACContract *RBACContractTransactor) AddFarmer(opts *bind.TransactOpts, farmerAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "addFarmer", farmerAddress)
}

// AddFarmer is a paid mutator transaction binding the contract method 0x80c3f96d.
//
// Solidity: function addFarmer(farmerAddress address) returns()
func (_RBACContract *RBACContractSession) AddFarmer(farmerAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddFarmer(&_RBACContract.TransactOpts, farmerAddress)
}

// AddFarmer is a paid mutator transaction binding the contract method 0x80c3f96d.
//
// Solidity: function addFarmer(farmerAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) AddFarmer(farmerAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddFarmer(&_RBACContract.TransactOpts, farmerAddress)
}

// AddInspector is a paid mutator transaction binding the contract method 0x7e458492.
//
// Solidity: function addInspector(inspectorAddress address) returns()
func (_RBACContract *RBACContractTransactor) AddInspector(opts *bind.TransactOpts, inspectorAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "addInspector", inspectorAddress)
}

// AddInspector is a paid mutator transaction binding the contract method 0x7e458492.
//
// Solidity: function addInspector(inspectorAddress address) returns()
func (_RBACContract *RBACContractSession) AddInspector(inspectorAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddInspector(&_RBACContract.TransactOpts, inspectorAddress)
}

// AddInspector is a paid mutator transaction binding the contract method 0x7e458492.
//
// Solidity: function addInspector(inspectorAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) AddInspector(inspectorAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddInspector(&_RBACContract.TransactOpts, inspectorAddress)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(adminAddress address) returns()
func (_RBACContract *RBACContractTransactor) RemoveAdmin(opts *bind.TransactOpts, adminAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "removeAdmin", adminAddress)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(adminAddress address) returns()
func (_RBACContract *RBACContractSession) RemoveAdmin(adminAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemoveAdmin(&_RBACContract.TransactOpts, adminAddress)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(adminAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) RemoveAdmin(adminAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemoveAdmin(&_RBACContract.TransactOpts, adminAddress)
}

// RemoveFarmer is a paid mutator transaction binding the contract method 0xe6bf3fdc.
//
// Solidity: function removeFarmer(farmerAddress address) returns()
func (_RBACContract *RBACContractTransactor) RemoveFarmer(opts *bind.TransactOpts, farmerAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "removeFarmer", farmerAddress)
}

// RemoveFarmer is a paid mutator transaction binding the contract method 0xe6bf3fdc.
//
// Solidity: function removeFarmer(farmerAddress address) returns()
func (_RBACContract *RBACContractSession) RemoveFarmer(farmerAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemoveFarmer(&_RBACContract.TransactOpts, farmerAddress)
}

// RemoveFarmer is a paid mutator transaction binding the contract method 0xe6bf3fdc.
//
// Solidity: function removeFarmer(farmerAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) RemoveFarmer(farmerAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemoveFarmer(&_RBACContract.TransactOpts, farmerAddress)
}

// RemoveInspector is a paid mutator transaction binding the contract method 0x7c70e791.
//
// Solidity: function removeInspector(inspectorAddress address) returns()
func (_RBACContract *RBACContractTransactor) RemoveInspector(opts *bind.TransactOpts, inspectorAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "removeInspector", inspectorAddress)
}

// RemoveInspector is a paid mutator transaction binding the contract method 0x7c70e791.
//
// Solidity: function removeInspector(inspectorAddress address) returns()
func (_RBACContract *RBACContractSession) RemoveInspector(inspectorAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemoveInspector(&_RBACContract.TransactOpts, inspectorAddress)
}

// RemoveInspector is a paid mutator transaction binding the contract method 0x7c70e791.
//
// Solidity: function removeInspector(inspectorAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) RemoveInspector(inspectorAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemoveInspector(&_RBACContract.TransactOpts, inspectorAddress)
}

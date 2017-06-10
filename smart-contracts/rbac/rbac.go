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
const RBACContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"inspectorAddress\",\"type\":\"address\"}],\"name\":\"isInspector\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cantonEmployeeAddress\",\"type\":\"address\"}],\"name\":\"isCantonEmployee\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"inspectorAddress\",\"type\":\"address\"}],\"name\":\"removeInspector\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"inspectorAddress\",\"type\":\"address\"}],\"name\":\"addInspector\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"farmerAddress\",\"type\":\"address\"}],\"name\":\"addFarmer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"farmerAddress\",\"type\":\"address\"}],\"name\":\"isFarmer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"farmerAddress\",\"type\":\"address\"}],\"name\":\"removeFarmer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cantonEmployeeAddress\",\"type\":\"address\"}],\"name\":\"addCantonEmployee\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cantonEmployeeAddress\",\"type\":\"address\"}],\"name\":\"removecantonEmployee\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// RBACContractBin is the compiled bytecode used for deploying new contracts.
const RBACContractBin = `6060604052341561000c57fe5b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b610ace8061005f6000396000f300606060405236156100b8576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631785f53c146100ba57806324d7806c146100f05780632cdad41c1461013e5780636bb164c91461018c57806370480275146101da5780637c70e791146102105780637e4584921461024657806380c3f96d1461027c578063b9429069146102b2578063e6bf3fdc14610300578063eaec3fc814610336578063f6bf0edd1461036c575bfe5b34156100c257fe5b6100ee600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103a2565b005b34156100f857fe5b610124600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610452565b604051808215151515815260200191505060405180910390f35b341561014657fe5b610172600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104a9565b604051808215151515815260200191505060405180910390f35b341561019457fe5b6101c0600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610500565b604051808215151515815260200191505060405180910390f35b34156101e257fe5b61020e600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610557565b005b341561021857fe5b610244600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610610565b005b341561024e57fe5b61027a600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506106c0565b005b341561028457fe5b6102b0600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610779565b005b34156102ba57fe5b6102e6600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610832565b604051808215151515815260200191505060405180910390f35b341561030857fe5b610334600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610889565b005b341561033e57fe5b61036a600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610939565b005b341561037457fe5b6103a0600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109f2565b005b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156103ff5760006000fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b50565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156105b45760006000fd5b6001600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614151561066d5760006000fd5b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614151561071d5760006000fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156107d65760006000fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156108e65760006000fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161415156109965760006000fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16141515610a4f5760006000fd5b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555b505600a165627a7a7230582029cb798615b55ae8d4dad374814bd03038eb6534ee706fc6a6e7d21ff3fc14eb0029`

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

// IsCantonEmployee is a free data retrieval call binding the contract method 0x6bb164c9.
//
// Solidity: function isCantonEmployee(cantonEmployeeAddress address) constant returns(bool)
func (_RBACContract *RBACContractCaller) IsCantonEmployee(opts *bind.CallOpts, cantonEmployeeAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBACContract.contract.Call(opts, out, "isCantonEmployee", cantonEmployeeAddress)
	return *ret0, err
}

// IsCantonEmployee is a free data retrieval call binding the contract method 0x6bb164c9.
//
// Solidity: function isCantonEmployee(cantonEmployeeAddress address) constant returns(bool)
func (_RBACContract *RBACContractSession) IsCantonEmployee(cantonEmployeeAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsCantonEmployee(&_RBACContract.CallOpts, cantonEmployeeAddress)
}

// IsCantonEmployee is a free data retrieval call binding the contract method 0x6bb164c9.
//
// Solidity: function isCantonEmployee(cantonEmployeeAddress address) constant returns(bool)
func (_RBACContract *RBACContractCallerSession) IsCantonEmployee(cantonEmployeeAddress common.Address) (bool, error) {
	return _RBACContract.Contract.IsCantonEmployee(&_RBACContract.CallOpts, cantonEmployeeAddress)
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

// AddCantonEmployee is a paid mutator transaction binding the contract method 0xeaec3fc8.
//
// Solidity: function addCantonEmployee(cantonEmployeeAddress address) returns()
func (_RBACContract *RBACContractTransactor) AddCantonEmployee(opts *bind.TransactOpts, cantonEmployeeAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "addCantonEmployee", cantonEmployeeAddress)
}

// AddCantonEmployee is a paid mutator transaction binding the contract method 0xeaec3fc8.
//
// Solidity: function addCantonEmployee(cantonEmployeeAddress address) returns()
func (_RBACContract *RBACContractSession) AddCantonEmployee(cantonEmployeeAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddCantonEmployee(&_RBACContract.TransactOpts, cantonEmployeeAddress)
}

// AddCantonEmployee is a paid mutator transaction binding the contract method 0xeaec3fc8.
//
// Solidity: function addCantonEmployee(cantonEmployeeAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) AddCantonEmployee(cantonEmployeeAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.AddCantonEmployee(&_RBACContract.TransactOpts, cantonEmployeeAddress)
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

// RemovecantonEmployee is a paid mutator transaction binding the contract method 0xf6bf0edd.
//
// Solidity: function removecantonEmployee(cantonEmployeeAddress address) returns()
func (_RBACContract *RBACContractTransactor) RemovecantonEmployee(opts *bind.TransactOpts, cantonEmployeeAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.contract.Transact(opts, "removecantonEmployee", cantonEmployeeAddress)
}

// RemovecantonEmployee is a paid mutator transaction binding the contract method 0xf6bf0edd.
//
// Solidity: function removecantonEmployee(cantonEmployeeAddress address) returns()
func (_RBACContract *RBACContractSession) RemovecantonEmployee(cantonEmployeeAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemovecantonEmployee(&_RBACContract.TransactOpts, cantonEmployeeAddress)
}

// RemovecantonEmployee is a paid mutator transaction binding the contract method 0xf6bf0edd.
//
// Solidity: function removecantonEmployee(cantonEmployeeAddress address) returns()
func (_RBACContract *RBACContractTransactorSession) RemovecantonEmployee(cantonEmployeeAddress common.Address) (*types.Transaction, error) {
	return _RBACContract.Contract.RemovecantonEmployee(&_RBACContract.TransactOpts, cantonEmployeeAddress)
}

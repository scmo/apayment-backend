// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package apaymenttoken

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// APaymentTokenContractABI is the input ABI used to generate the binding from.
const APaymentTokenContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"transferWithMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"requestAdr\",\"type\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"transferWithMessageAndRequestAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// APaymentTokenContractBin is the compiled bytecode used for deploying new contracts.
const APaymentTokenContractBin = `6060604052341561000c57fe5b604051602080610bfa833981016040528080519060200190919050505b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806000819055505b505b610b74806100866000396000f3006060604052361561008c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806307560f131461008e578063095ea7b31461012857806318160ddd1461017f57806323b872dd146101a557806370a082311461021b578063a9059cbb14610265578063d3088b52146102bc578063dd62ed3e14610375575bfe5b341561009657fe5b61010e600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506103de565b604051808215151515815260200191505060405180910390f35b341561013057fe5b610165600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506103f4565b604051808215151515815260200191505060405180910390f35b341561018757fe5b61018f6104e7565b6040518082815260200191505060405180910390f35b34156101ad57fe5b610201600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506104f2565b604051808215151515815260200191505060405180910390f35b341561022357fe5b61024f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610856565b6040518082815260200191505060405180910390f35b341561026d57fe5b6102a2600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506108a0565b604051808215151515815260200191505060405180910390f35b34156102c457fe5b61035b600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610a75565b604051808215151515815260200191505060405180910390f35b341561037d57fe5b6103c8600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610a8c565b6040518082815260200191505060405180910390f35b60006103ea84846108a0565b90505b9392505050565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a3600190505b92915050565b600060005490505b90565b600081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015151561053f57fe5b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156105c757fe5b61064d600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610b14565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610716600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610b14565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506107a2600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610b2e565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b9392505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b919050565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156108ed57fe5b610936600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610b14565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506109c2600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610b2e565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b92915050565b6000610a8185856108a0565b90505b949350505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b92915050565b60008282840391508111151515610b2757fe5b5b92915050565b60008282840191508110151515610b4157fe5b5b929150505600a165627a7a72305820745a41f4e7f8cfe35e1f24035d95acfec294c09169dca59019258d17d5cee93c0029`

// DeployAPaymentTokenContract deploys a new Ethereum contract, binding an instance of APaymentTokenContract to it.
func DeployAPaymentTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend, supply *big.Int) (common.Address, *types.Transaction, *APaymentTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APaymentTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(APaymentTokenContractBin), backend, supply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APaymentTokenContract{APaymentTokenContractCaller: APaymentTokenContractCaller{contract: contract}, APaymentTokenContractTransactor: APaymentTokenContractTransactor{contract: contract}}, nil
}

// APaymentTokenContract is an auto generated Go binding around an Ethereum contract.
type APaymentTokenContract struct {
	APaymentTokenContractCaller     // Read-only binding to the contract
	APaymentTokenContractTransactor // Write-only binding to the contract
}

// APaymentTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type APaymentTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APaymentTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APaymentTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APaymentTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APaymentTokenContractSession struct {
	Contract     *APaymentTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// APaymentTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APaymentTokenContractCallerSession struct {
	Contract *APaymentTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// APaymentTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APaymentTokenContractTransactorSession struct {
	Contract     *APaymentTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// APaymentTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type APaymentTokenContractRaw struct {
	Contract *APaymentTokenContract // Generic contract binding to access the raw methods on
}

// APaymentTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APaymentTokenContractCallerRaw struct {
	Contract *APaymentTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// APaymentTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APaymentTokenContractTransactorRaw struct {
	Contract *APaymentTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPaymentTokenContract creates a new instance of APaymentTokenContract, bound to a specific deployed contract.
func NewAPaymentTokenContract(address common.Address, backend bind.ContractBackend) (*APaymentTokenContract, error) {
	contract, err := bindAPaymentTokenContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APaymentTokenContract{APaymentTokenContractCaller: APaymentTokenContractCaller{contract: contract}, APaymentTokenContractTransactor: APaymentTokenContractTransactor{contract: contract}}, nil
}

// NewAPaymentTokenContractCaller creates a new read-only instance of APaymentTokenContract, bound to a specific deployed contract.
func NewAPaymentTokenContractCaller(address common.Address, caller bind.ContractCaller) (*APaymentTokenContractCaller, error) {
	contract, err := bindAPaymentTokenContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &APaymentTokenContractCaller{contract: contract}, nil
}

// NewAPaymentTokenContractTransactor creates a new write-only instance of APaymentTokenContract, bound to a specific deployed contract.
func NewAPaymentTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*APaymentTokenContractTransactor, error) {
	contract, err := bindAPaymentTokenContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &APaymentTokenContractTransactor{contract: contract}, nil
}

// bindAPaymentTokenContract binds a generic wrapper to an already deployed contract.
func bindAPaymentTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APaymentTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APaymentTokenContract *APaymentTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _APaymentTokenContract.Contract.APaymentTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APaymentTokenContract *APaymentTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.APaymentTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APaymentTokenContract *APaymentTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.APaymentTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APaymentTokenContract *APaymentTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _APaymentTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APaymentTokenContract *APaymentTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APaymentTokenContract *APaymentTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _APaymentTokenContract.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _APaymentTokenContract.Contract.Allowance(&_APaymentTokenContract.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _APaymentTokenContract.Contract.Allowance(&_APaymentTokenContract.CallOpts, src, guy)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _APaymentTokenContract.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _APaymentTokenContract.Contract.BalanceOf(&_APaymentTokenContract.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _APaymentTokenContract.Contract.BalanceOf(&_APaymentTokenContract.CallOpts, src)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _APaymentTokenContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractSession) TotalSupply() (*big.Int, error) {
	return _APaymentTokenContract.Contract.TotalSupply(&_APaymentTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_APaymentTokenContract *APaymentTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _APaymentTokenContract.Contract.TotalSupply(&_APaymentTokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactor) Approve(opts *bind.TransactOpts, guy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.contract.Transact(opts, "approve", guy, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractSession) Approve(guy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.Approve(&_APaymentTokenContract.TransactOpts, guy, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactorSession) Approve(guy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.Approve(&_APaymentTokenContract.TransactOpts, guy, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.Transfer(&_APaymentTokenContract.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.Transfer(&_APaymentTokenContract.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.TransferFrom(&_APaymentTokenContract.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, amount uint256) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.TransferFrom(&_APaymentTokenContract.TransactOpts, src, dst, amount)
}

// TransferWithMessage is a paid mutator transaction binding the contract method 0x07560f13.
//
// Solidity: function transferWithMessage(dst address, amount uint256, message bytes) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactor) TransferWithMessage(opts *bind.TransactOpts, dst common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _APaymentTokenContract.contract.Transact(opts, "transferWithMessage", dst, amount, message)
}

// TransferWithMessage is a paid mutator transaction binding the contract method 0x07560f13.
//
// Solidity: function transferWithMessage(dst address, amount uint256, message bytes) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractSession) TransferWithMessage(dst common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.TransferWithMessage(&_APaymentTokenContract.TransactOpts, dst, amount, message)
}

// TransferWithMessage is a paid mutator transaction binding the contract method 0x07560f13.
//
// Solidity: function transferWithMessage(dst address, amount uint256, message bytes) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactorSession) TransferWithMessage(dst common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.TransferWithMessage(&_APaymentTokenContract.TransactOpts, dst, amount, message)
}

// TransferWithMessageAndRequestAddress is a paid mutator transaction binding the contract method 0xd3088b52.
//
// Solidity: function transferWithMessageAndRequestAddress(dst address, amount uint256, requestAdr address, message bytes) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactor) TransferWithMessageAndRequestAddress(opts *bind.TransactOpts, dst common.Address, amount *big.Int, requestAdr common.Address, message []byte) (*types.Transaction, error) {
	return _APaymentTokenContract.contract.Transact(opts, "transferWithMessageAndRequestAddress", dst, amount, requestAdr, message)
}

// TransferWithMessageAndRequestAddress is a paid mutator transaction binding the contract method 0xd3088b52.
//
// Solidity: function transferWithMessageAndRequestAddress(dst address, amount uint256, requestAdr address, message bytes) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractSession) TransferWithMessageAndRequestAddress(dst common.Address, amount *big.Int, requestAdr common.Address, message []byte) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.TransferWithMessageAndRequestAddress(&_APaymentTokenContract.TransactOpts, dst, amount, requestAdr, message)
}

// TransferWithMessageAndRequestAddress is a paid mutator transaction binding the contract method 0xd3088b52.
//
// Solidity: function transferWithMessageAndRequestAddress(dst address, amount uint256, requestAdr address, message bytes) returns(bool)
func (_APaymentTokenContract *APaymentTokenContractTransactorSession) TransferWithMessageAndRequestAddress(dst common.Address, amount *big.Int, requestAdr common.Address, message []byte) (*types.Transaction, error) {
	return _APaymentTokenContract.Contract.TransferWithMessageAndRequestAddress(&_APaymentTokenContract.TransactOpts, dst, amount, requestAdr, message)
}

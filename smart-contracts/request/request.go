// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package smartcontracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RequestContractABI is the input ABI used to generate the binding from.
const RequestContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"string\"},{\"name\":\"pointGroupId\",\"type\":\"string\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numLacks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userId\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_inspectorId\",\"type\":\"int64\"}],\"name\":\"setInspectorId\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspectorId\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributionCodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCode\",\"type\":\"uint16\"},{\"name\":\"_controlCategoryId\",\"type\":\"string\"},{\"name\":\"_pointGroupId\",\"type\":\"string\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"}],\"name\":\"addLack\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_userId\",\"type\":\"int64\"},{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_remark\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"}]"

// RequestContractBin is the compiled bytecode used for deploying new contracts.
const RequestContractBin = `6060604052341561000c57fe5b604051610b9c380380610b9c833981016040528080519060200190919080518201919060200180518201919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b82600060146101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555081600290805190602001906100c09291906100e1565b5080600390805190602001906100d792919061018b565b505b505050610261565b82805482825590600052602060002090600f0160109004810192821561017a5791602002820160005b8382111561014a57835183826101000a81548161ffff021916908361ffff160217905550926020019260020160208160010104928301926001030261010a565b80156101785782816101000a81549061ffff021916905560020160208160010104928301926001030261014a565b505b509050610187919061020b565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101cc57805160ff19168380011785556101fa565b828001600101855582156101fa579182015b828111156101f95782518255916020019190600101906101de565b5b509050610207919061023c565b5090565b61023991905b8082111561023557600081816101000a81549061ffff021916905550600101610211565b5090565b90565b61025e91905b8082111561025a576000816000905550600101610242565b5090565b90565b61092c806102706000396000f300606060405236156100a2576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630478e042146100a4578063103097811461028857806341c0e1b5146102ae57806355bcbf57146102c057806358975919146102e6578063727af9c61461031257806379a517fe146103355780639b5d25d914610361578063a47d1f5e1461039d578063ad0672fd14610436575bfe5b34156100ac57fe5b6100c2600480803590602001909190505061052f565b604051808661ffff1661ffff1681526020018060200180602001806020018560070b60070b815260200184810384528881815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561016d5780601f106101425761010080835404028352916020019161016d565b820191906000526020600020905b81548152906001019060200180831161015057829003601f168201915b50508481038352878181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156101f05780601f106101c5576101008083540402835291602001916101f0565b820191906000526020600020905b8154815290600101906020018083116101d357829003601f168201915b50508481038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102735780601f1061024857610100808354040283529160200191610273565b820191906000526020600020905b81548152906001019060200180831161025657829003601f168201915b50509850505050505050505060405180910390f35b341561029057fe5b61029861057d565b6040518082815260200191505060405180910390f35b34156102b657fe5b6102be610586565b005b34156102c857fe5b6102d061061a565b6040518082815260200191505060405180910390f35b34156102ee57fe5b6102f6610620565b604051808260070b60070b815260200191505060405180910390f35b341561031a57fe5b610333600480803560070b906020019091905050610633565b005b341561033d57fe5b610345610663565b604051808260070b60070b815260200191505060405180910390f35b341561036957fe5b61037f6004808035906020019091905050610676565b604051808261ffff1661ffff16815260200191505060405180910390f35b34156103a557fe5b6103ad6106ae565b60405180806020018281038252838181518152602001915080519060200190808383600083146103fc575b8051825260208311156103fc576020820191506020810190506020830392506103d8565b505050905090810190601f1680156104285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561043e57fe5b61052d600480803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803560070b90602001909190505061074c565b005b60056020528060005260406000206000915090508060000160009054906101000a900461ffff16908060010190806002019080600301908060040160009054906101000a900460070b905085565b60004290505b90565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561061757600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60045481565b600060149054906101000a900460070b81565b80600160006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055505b50565b600160009054906101000a900460070b81565b60028181548110151561068557fe5b90600052602060002090601091828204019190066002025b915054906101000a900461ffff1681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107445780601f1061071957610100808354040283529160200191610744565b820191906000526020600020905b81548152906001019060200180831161072757829003601f168201915b505050505081565b60006004600081548092919060010191905055905060a0604051908101604052808761ffff1681526020018681526020018581526020018481526020018360070b8152506005600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160010190805190602001906107e292919061085b565b5060408201518160020190805190602001906107ff92919061085b565b50606082015181600301908051906020019061081c92919061085b565b5060808201518160040160006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055509050505b505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061089c57805160ff19168380011785556108ca565b828001600101855582156108ca579182015b828111156108c95782518255916020019190600101906108ae565b5b5090506108d791906108db565b5090565b6108fd91905b808211156108f95760008160009055506001016108e1565b5090565b905600a165627a7a72305820ab9a862c291dabd41459b7ad61267ddd577656a67ee79212a752b1133c50122b0029`

// DeployRequestContract deploys a new Ethereum contract, binding an instance of RequestContract to it.
func DeployRequestContract(auth *bind.TransactOpts, backend bind.ContractBackend, _userId int64, _contributionCodes []uint16, _remark string) (common.Address, *types.Transaction, *RequestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestContractBin), backend, _userId, _contributionCodes, _remark)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestContract{RequestContractCaller: RequestContractCaller{contract: contract}, RequestContractTransactor: RequestContractTransactor{contract: contract}}, nil
}

// RequestContract is an auto generated Go binding around an Ethereum contract.
type RequestContract struct {
	RequestContractCaller     // Read-only binding to the contract
	RequestContractTransactor // Write-only binding to the contract
}

// RequestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestContractSession struct {
	Contract     *RequestContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestContractCallerSession struct {
	Contract *RequestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RequestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestContractTransactorSession struct {
	Contract     *RequestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RequestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestContractRaw struct {
	Contract *RequestContract // Generic contract binding to access the raw methods on
}

// RequestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestContractCallerRaw struct {
	Contract *RequestContractCaller // Generic read-only contract binding to access the raw methods on
}

// RequestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestContractTransactorRaw struct {
	Contract *RequestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestContract creates a new instance of RequestContract, bound to a specific deployed contract.
func NewRequestContract(address common.Address, backend bind.ContractBackend) (*RequestContract, error) {
	contract, err := bindRequestContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestContract{RequestContractCaller: RequestContractCaller{contract: contract}, RequestContractTransactor: RequestContractTransactor{contract: contract}}, nil
}

// NewRequestContractCaller creates a new read-only instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractCaller(address common.Address, caller bind.ContractCaller) (*RequestContractCaller, error) {
	contract, err := bindRequestContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RequestContractCaller{contract: contract}, nil
}

// NewRequestContractTransactor creates a new write-only instance of RequestContract, bound to a specific deployed contract.
func NewRequestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestContractTransactor, error) {
	contract, err := bindRequestContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RequestContractTransactor{contract: contract}, nil
}

// bindRequestContract binds a generic wrapper to an already deployed contract.
func bindRequestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestContract *RequestContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestContract.Contract.RequestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestContract *RequestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.Contract.RequestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestContract *RequestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestContract.Contract.RequestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestContract *RequestContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestContract *RequestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestContract *RequestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestContract.Contract.contract.Transact(opts, method, params...)
}

// ContributionCodes is a free data retrieval call binding the contract method 0x9b5d25d9.
//
// Solidity: function contributionCodes( uint256) constant returns(uint16)
func (_RequestContract *RequestContractCaller) ContributionCodes(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "contributionCodes", arg0)
	return *ret0, err
}

// ContributionCodes is a free data retrieval call binding the contract method 0x9b5d25d9.
//
// Solidity: function contributionCodes( uint256) constant returns(uint16)
func (_RequestContract *RequestContractSession) ContributionCodes(arg0 *big.Int) (uint16, error) {
	return _RequestContract.Contract.ContributionCodes(&_RequestContract.CallOpts, arg0)
}

// ContributionCodes is a free data retrieval call binding the contract method 0x9b5d25d9.
//
// Solidity: function contributionCodes( uint256) constant returns(uint16)
func (_RequestContract *RequestContractCallerSession) ContributionCodes(arg0 *big.Int) (uint16, error) {
	return _RequestContract.Contract.ContributionCodes(&_RequestContract.CallOpts, arg0)
}

// CreatedTimestamp is a free data retrieval call binding the contract method 0x10309781.
//
// Solidity: function createdTimestamp() constant returns(uint256)
func (_RequestContract *RequestContractCaller) CreatedTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "createdTimestamp")
	return *ret0, err
}

// CreatedTimestamp is a free data retrieval call binding the contract method 0x10309781.
//
// Solidity: function createdTimestamp() constant returns(uint256)
func (_RequestContract *RequestContractSession) CreatedTimestamp() (*big.Int, error) {
	return _RequestContract.Contract.CreatedTimestamp(&_RequestContract.CallOpts)
}

// CreatedTimestamp is a free data retrieval call binding the contract method 0x10309781.
//
// Solidity: function createdTimestamp() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) CreatedTimestamp() (*big.Int, error) {
	return _RequestContract.Contract.CreatedTimestamp(&_RequestContract.CallOpts)
}

// InspectorId is a free data retrieval call binding the contract method 0x79a517fe.
//
// Solidity: function inspectorId() constant returns(int64)
func (_RequestContract *RequestContractCaller) InspectorId(opts *bind.CallOpts) (int64, error) {
	var (
		ret0 = new(int64)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "inspectorId")
	return *ret0, err
}

// InspectorId is a free data retrieval call binding the contract method 0x79a517fe.
//
// Solidity: function inspectorId() constant returns(int64)
func (_RequestContract *RequestContractSession) InspectorId() (int64, error) {
	return _RequestContract.Contract.InspectorId(&_RequestContract.CallOpts)
}

// InspectorId is a free data retrieval call binding the contract method 0x79a517fe.
//
// Solidity: function inspectorId() constant returns(int64)
func (_RequestContract *RequestContractCallerSession) InspectorId() (int64, error) {
	return _RequestContract.Contract.InspectorId(&_RequestContract.CallOpts)
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId string, pointGroupId string, controlPointId string, lackId int64)
func (_RequestContract *RequestContractCaller) Lacks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId string
	PointGroupId      string
	ControlPointId    string
	LackId            int64
}, error) {
	ret := new(struct {
		ContributionCode  uint16
		ControlCategoryId string
		PointGroupId      string
		ControlPointId    string
		LackId            int64
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "lacks", arg0)
	return *ret, err
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId string, pointGroupId string, controlPointId string, lackId int64)
func (_RequestContract *RequestContractSession) Lacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId string
	PointGroupId      string
	ControlPointId    string
	LackId            int64
}, error) {
	return _RequestContract.Contract.Lacks(&_RequestContract.CallOpts, arg0)
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId string, pointGroupId string, controlPointId string, lackId int64)
func (_RequestContract *RequestContractCallerSession) Lacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId string
	PointGroupId      string
	ControlPointId    string
	LackId            int64
}, error) {
	return _RequestContract.Contract.Lacks(&_RequestContract.CallOpts, arg0)
}

// NumLacks is a free data retrieval call binding the contract method 0x55bcbf57.
//
// Solidity: function numLacks() constant returns(uint256)
func (_RequestContract *RequestContractCaller) NumLacks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "numLacks")
	return *ret0, err
}

// NumLacks is a free data retrieval call binding the contract method 0x55bcbf57.
//
// Solidity: function numLacks() constant returns(uint256)
func (_RequestContract *RequestContractSession) NumLacks() (*big.Int, error) {
	return _RequestContract.Contract.NumLacks(&_RequestContract.CallOpts)
}

// NumLacks is a free data retrieval call binding the contract method 0x55bcbf57.
//
// Solidity: function numLacks() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) NumLacks() (*big.Int, error) {
	return _RequestContract.Contract.NumLacks(&_RequestContract.CallOpts)
}

// Remark is a free data retrieval call binding the contract method 0xa47d1f5e.
//
// Solidity: function remark() constant returns(string)
func (_RequestContract *RequestContractCaller) Remark(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "remark")
	return *ret0, err
}

// Remark is a free data retrieval call binding the contract method 0xa47d1f5e.
//
// Solidity: function remark() constant returns(string)
func (_RequestContract *RequestContractSession) Remark() (string, error) {
	return _RequestContract.Contract.Remark(&_RequestContract.CallOpts)
}

// Remark is a free data retrieval call binding the contract method 0xa47d1f5e.
//
// Solidity: function remark() constant returns(string)
func (_RequestContract *RequestContractCallerSession) Remark() (string, error) {
	return _RequestContract.Contract.Remark(&_RequestContract.CallOpts)
}

// UserId is a free data retrieval call binding the contract method 0x58975919.
//
// Solidity: function userId() constant returns(int64)
func (_RequestContract *RequestContractCaller) UserId(opts *bind.CallOpts) (int64, error) {
	var (
		ret0 = new(int64)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "userId")
	return *ret0, err
}

// UserId is a free data retrieval call binding the contract method 0x58975919.
//
// Solidity: function userId() constant returns(int64)
func (_RequestContract *RequestContractSession) UserId() (int64, error) {
	return _RequestContract.Contract.UserId(&_RequestContract.CallOpts)
}

// UserId is a free data retrieval call binding the contract method 0x58975919.
//
// Solidity: function userId() constant returns(int64)
func (_RequestContract *RequestContractCallerSession) UserId() (int64, error) {
	return _RequestContract.Contract.UserId(&_RequestContract.CallOpts)
}

// AddLack is a paid mutator transaction binding the contract method 0xad0672fd.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupId string, controlPointId string, lackId int64) returns()
func (_RequestContract *RequestContractTransactor) AddLack(opts *bind.TransactOpts, _contributionCode uint16, _controlCategoryId string, _pointGroupId string, controlPointId string, lackId int64) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "addLack", _contributionCode, _controlCategoryId, _pointGroupId, controlPointId, lackId)
}

// AddLack is a paid mutator transaction binding the contract method 0xad0672fd.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupId string, controlPointId string, lackId int64) returns()
func (_RequestContract *RequestContractSession) AddLack(_contributionCode uint16, _controlCategoryId string, _pointGroupId string, controlPointId string, lackId int64) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLack(&_RequestContract.TransactOpts, _contributionCode, _controlCategoryId, _pointGroupId, controlPointId, lackId)
}

// AddLack is a paid mutator transaction binding the contract method 0xad0672fd.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupId string, controlPointId string, lackId int64) returns()
func (_RequestContract *RequestContractTransactorSession) AddLack(_contributionCode uint16, _controlCategoryId string, _pointGroupId string, controlPointId string, lackId int64) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLack(&_RequestContract.TransactOpts, _contributionCode, _controlCategoryId, _pointGroupId, controlPointId, lackId)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RequestContract *RequestContractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RequestContract *RequestContractSession) Kill() (*types.Transaction, error) {
	return _RequestContract.Contract.Kill(&_RequestContract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RequestContract *RequestContractTransactorSession) Kill() (*types.Transaction, error) {
	return _RequestContract.Contract.Kill(&_RequestContract.TransactOpts)
}

// SetInspectorId is a paid mutator transaction binding the contract method 0x727af9c6.
//
// Solidity: function setInspectorId(_inspectorId int64) returns()
func (_RequestContract *RequestContractTransactor) SetInspectorId(opts *bind.TransactOpts, _inspectorId int64) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "setInspectorId", _inspectorId)
}

// SetInspectorId is a paid mutator transaction binding the contract method 0x727af9c6.
//
// Solidity: function setInspectorId(_inspectorId int64) returns()
func (_RequestContract *RequestContractSession) SetInspectorId(_inspectorId int64) (*types.Transaction, error) {
	return _RequestContract.Contract.SetInspectorId(&_RequestContract.TransactOpts, _inspectorId)
}

// SetInspectorId is a paid mutator transaction binding the contract method 0x727af9c6.
//
// Solidity: function setInspectorId(_inspectorId int64) returns()
func (_RequestContract *RequestContractTransactorSession) SetInspectorId(_inspectorId int64) (*types.Transaction, error) {
	return _RequestContract.Contract.SetInspectorId(&_RequestContract.TransactOpts, _inspectorId)
}

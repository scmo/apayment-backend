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
const RequestContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_inspectorAddress\",\"type\":\"address\"}],\"name\":\"setInspectorId\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"string\"},{\"name\":\"pointGroupId\",\"type\":\"string\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modified\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numLacks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userId\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspectorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setCreated\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributionCodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCode\",\"type\":\"uint16\"},{\"name\":\"_controlCategoryId\",\"type\":\"string\"},{\"name\":\"_pointGroupId\",\"type\":\"string\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"}],\"name\":\"addLack\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setModified\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_userId\",\"type\":\"int64\"},{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_remark\",\"type\":\"string\"},{\"name\":\"rbacAddress\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// RequestContractBin is the compiled bytecode used for deploying new contracts.
const RequestContractBin = `606060405234156200000d57fe5b60405162000d7f38038062000d7f833981016040528080519060200190919080518201919060200180518201919060200180519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055508260049080519060200190620001129291906200015f565b5081600590805190602001906200012b92919062000210565b506200014a62000155640100000000026200074d176401000000009004565b5b50505050620002f3565b426001819055505b565b82805482825590600052602060002090600f01601090048101928215620001fd5791602002820160005b83821115620001cb57835183826101000a81548161ffff021916908361ffff160217905550926020019260020160208160010104928301926001030262000189565b8015620001fb5782816101000a81549061ffff0219169055600201602081600101049283019260010302620001cb565b505b5090506200020c919062000297565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025357805160ff191683800117855562000284565b8280016001018555821562000284579182015b828111156200028357825182559160200191906001019062000266565b5b509050620002939190620002cb565b5090565b620002c891905b80821115620002c457600081816101000a81549061ffff0219169055506001016200029e565b5090565b90565b620002f091905b80821115620002ec576000816000905550600101620002d2565b5090565b90565b610a7c80620003036000396000f300606060405236156100c3576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063021c7bd7146100c55780630478e042146100fb578063325a19f1146102df57806341c0e1b51461030557806345ecd3d71461031757806355bcbf571461033d5780635897591914610363578063696f8c571461038f57806390e10250146103e15780639b5d25d9146103f3578063a47d1f5e1461042f578063ad0672fd146104c8578063af7fdd76146105c1575bfe5b34156100cd57fe5b6100f9600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506105d3565b005b341561010357fe5b6101196004808035906020019091905050610620565b604051808661ffff1661ffff1681526020018060200180602001806020018560070b60070b81526020018481038452888181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156101c45780601f10610199576101008083540402835291602001916101c4565b820191906000526020600020905b8154815290600101906020018083116101a757829003601f168201915b50508481038352878181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102475780601f1061021c57610100808354040283529160200191610247565b820191906000526020600020905b81548152906001019060200180831161022a57829003601f168201915b50508481038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102ca5780601f1061029f576101008083540402835291602001916102ca565b820191906000526020600020905b8154815290600101906020018083116102ad57829003601f168201915b50509850505050505050505060405180910390f35b34156102e757fe5b6102ef61066e565b6040518082815260200191505060405180910390f35b341561030d57fe5b610315610674565b005b341561031f57fe5b610327610708565b6040518082815260200191505060405180910390f35b341561034557fe5b61034d61070e565b6040518082815260200191505060405180910390f35b341561036b57fe5b610373610714565b604051808260070b60070b815260200191505060405180910390f35b341561039757fe5b61039f610727565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103e957fe5b6103f161074d565b005b34156103fb57fe5b6104116004808035906020019091905050610757565b604051808261ffff1661ffff16815260200191505060405180910390f35b341561043757fe5b61043f61078f565b604051808060200182810382528381815181526020019150805190602001908083836000831461048e575b80518252602083111561048e5760208201915060208101905060208303925061046a565b505050905090810190601f1680156104ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156104d057fe5b6105bf600480803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803560070b90602001909190505061082d565b005b34156105c957fe5b6105d16109a1565b005b80600360086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061061c6109a1565b5b50565b60086020528060005260406000206000915090508060000160009054906101000a900461ffff16908060010190806002019080600301908060040160009054906101000a900460070b905085565b60015481565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561070557600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60025481565b60075481565b600360009054906101000a900460070b81565b600360089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b426001819055505b565b60048181548110151561076657fe5b90600052602060002090601091828204019190066002025b915054906101000a900461ffff1681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108255780601f106107fa57610100808354040283529160200191610825565b820191906000526020600020905b81548152906001019060200180831161080857829003601f168201915b505050505081565b6000600360089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561088c5760006000fd5b6007600081548092919060010191905055905060a0604051908101604052808761ffff1681526020018681526020018581526020018481526020018360070b8152506008600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160010190805190602001906109209291906109ab565b50604082015181600201908051906020019061093d9291906109ab565b50606082015181600301908051906020019061095a9291906109ab565b5060808201518160040160006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055509050506109986109a1565b5b505050505050565b426002819055505b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109ec57805160ff1916838001178555610a1a565b82800160010185558215610a1a579182015b82811115610a195782518255916020019190600101906109fe565b5b509050610a279190610a2b565b5090565b610a4d91905b80821115610a49576000816000905550600101610a31565b5090565b905600a165627a7a7230582020452c69d48901b4bb34f2031c1fbf1e641be434eed28caeb9872b552acefa7c0029`

// DeployRequestContract deploys a new Ethereum contract, binding an instance of RequestContract to it.
func DeployRequestContract(auth *bind.TransactOpts, backend bind.ContractBackend, _userId int64, _contributionCodes []uint16, _remark string, rbacAddress common.Address) (common.Address, *types.Transaction, *RequestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestContractBin), backend, _userId, _contributionCodes, _remark, rbacAddress)
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

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_RequestContract *RequestContractCaller) Created(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "created")
	return *ret0, err
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_RequestContract *RequestContractSession) Created() (*big.Int, error) {
	return _RequestContract.Contract.Created(&_RequestContract.CallOpts)
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) Created() (*big.Int, error) {
	return _RequestContract.Contract.Created(&_RequestContract.CallOpts)
}

// InspectorAddress is a free data retrieval call binding the contract method 0x696f8c57.
//
// Solidity: function inspectorAddress() constant returns(address)
func (_RequestContract *RequestContractCaller) InspectorAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "inspectorAddress")
	return *ret0, err
}

// InspectorAddress is a free data retrieval call binding the contract method 0x696f8c57.
//
// Solidity: function inspectorAddress() constant returns(address)
func (_RequestContract *RequestContractSession) InspectorAddress() (common.Address, error) {
	return _RequestContract.Contract.InspectorAddress(&_RequestContract.CallOpts)
}

// InspectorAddress is a free data retrieval call binding the contract method 0x696f8c57.
//
// Solidity: function inspectorAddress() constant returns(address)
func (_RequestContract *RequestContractCallerSession) InspectorAddress() (common.Address, error) {
	return _RequestContract.Contract.InspectorAddress(&_RequestContract.CallOpts)
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

// Modified is a free data retrieval call binding the contract method 0x45ecd3d7.
//
// Solidity: function modified() constant returns(uint256)
func (_RequestContract *RequestContractCaller) Modified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "modified")
	return *ret0, err
}

// Modified is a free data retrieval call binding the contract method 0x45ecd3d7.
//
// Solidity: function modified() constant returns(uint256)
func (_RequestContract *RequestContractSession) Modified() (*big.Int, error) {
	return _RequestContract.Contract.Modified(&_RequestContract.CallOpts)
}

// Modified is a free data retrieval call binding the contract method 0x45ecd3d7.
//
// Solidity: function modified() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) Modified() (*big.Int, error) {
	return _RequestContract.Contract.Modified(&_RequestContract.CallOpts)
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

// SetCreated is a paid mutator transaction binding the contract method 0x90e10250.
//
// Solidity: function setCreated() returns()
func (_RequestContract *RequestContractTransactor) SetCreated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "setCreated")
}

// SetCreated is a paid mutator transaction binding the contract method 0x90e10250.
//
// Solidity: function setCreated() returns()
func (_RequestContract *RequestContractSession) SetCreated() (*types.Transaction, error) {
	return _RequestContract.Contract.SetCreated(&_RequestContract.TransactOpts)
}

// SetCreated is a paid mutator transaction binding the contract method 0x90e10250.
//
// Solidity: function setCreated() returns()
func (_RequestContract *RequestContractTransactorSession) SetCreated() (*types.Transaction, error) {
	return _RequestContract.Contract.SetCreated(&_RequestContract.TransactOpts)
}

// SetInspectorId is a paid mutator transaction binding the contract method 0x021c7bd7.
//
// Solidity: function setInspectorId(_inspectorAddress address) returns()
func (_RequestContract *RequestContractTransactor) SetInspectorId(opts *bind.TransactOpts, _inspectorAddress common.Address) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "setInspectorId", _inspectorAddress)
}

// SetInspectorId is a paid mutator transaction binding the contract method 0x021c7bd7.
//
// Solidity: function setInspectorId(_inspectorAddress address) returns()
func (_RequestContract *RequestContractSession) SetInspectorId(_inspectorAddress common.Address) (*types.Transaction, error) {
	return _RequestContract.Contract.SetInspectorId(&_RequestContract.TransactOpts, _inspectorAddress)
}

// SetInspectorId is a paid mutator transaction binding the contract method 0x021c7bd7.
//
// Solidity: function setInspectorId(_inspectorAddress address) returns()
func (_RequestContract *RequestContractTransactorSession) SetInspectorId(_inspectorAddress common.Address) (*types.Transaction, error) {
	return _RequestContract.Contract.SetInspectorId(&_RequestContract.TransactOpts, _inspectorAddress)
}

// SetModified is a paid mutator transaction binding the contract method 0xaf7fdd76.
//
// Solidity: function setModified() returns()
func (_RequestContract *RequestContractTransactor) SetModified(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "setModified")
}

// SetModified is a paid mutator transaction binding the contract method 0xaf7fdd76.
//
// Solidity: function setModified() returns()
func (_RequestContract *RequestContractSession) SetModified() (*types.Transaction, error) {
	return _RequestContract.Contract.SetModified(&_RequestContract.TransactOpts)
}

// SetModified is a paid mutator transaction binding the contract method 0xaf7fdd76.
//
// Solidity: function setModified() returns()
func (_RequestContract *RequestContractTransactorSession) SetModified() (*types.Transaction, error) {
	return _RequestContract.Contract.SetModified(&_RequestContract.TransactOpts)
}

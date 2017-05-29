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
const RequestContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_inspectorAddress\",\"type\":\"address\"}],\"name\":\"setInspectorId\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"string\"},{\"name\":\"pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"},{\"name\":\"points\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modified\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCode\",\"type\":\"uint16\"},{\"name\":\"_controlCategoryId\",\"type\":\"string\"},{\"name\":\"_pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"},{\"name\":\"points\",\"type\":\"uint8\"}],\"name\":\"addLack\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numLacks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userId\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspectorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calculateBTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setCreated\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributionCodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setModified\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_userId\",\"type\":\"int64\"},{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_remark\",\"type\":\"string\"},{\"name\":\"rbacAddress\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// RequestContractBin is the compiled bytecode used for deploying new contracts.
const RequestContractBin = `60606040526101406040519081016040528061045661ffff16815260200161047e61ffff16815260200161046361ffff16815260200161046861ffff16815260200161047561ffff16815260200161047661ffff16815260200161046461ffff16815260200161046961ffff16815260200161047761ffff16815260200161047861ffff16815250600a90600a62000099929190620001f5565b503415620000a357fe5b6040516200132738038062001327833981016040528080519060200190919080518201919060200180518201919060200180519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055508260049080519060200190620001a892919062000299565b508160059080519060200190620001c19291906200034a565b50620001e0620001eb6401000000000262000d2f176401000000009004565b5b505050506200042d565b426001819055505b565b82600a600f01601090048101928215620002865791602002820160005b838211156200025457835183826101000a81548161ffff021916908361ffff160217905550926020019260020160208160010104928301926001030262000212565b8015620002845782816101000a81549061ffff021916905560020160208160010104928301926001030262000254565b505b509050620002959190620003d1565b5090565b82805482825590600052602060002090600f01601090048101928215620003375791602002820160005b838211156200030557835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302620002c3565b8015620003355782816101000a81549061ffff021916905560020160208160010104928301926001030262000305565b505b509050620003469190620003d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200038d57805160ff1916838001178555620003be565b82800160010185558215620003be579182015b82811115620003bd578251825591602001919060010190620003a0565b5b509050620003cd919062000405565b5090565b6200040291905b80821115620003fe57600081816101000a81549061ffff021916905550600101620003d8565b5090565b90565b6200042a91905b80821115620004265760008160009055506001016200040c565b5090565b90565b610eea806200043d6000396000f300606060405236156100ce576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063021c7bd7146100d05780630478e04214610106578063325a19f11461027d57806341c0e1b5146102a357806345ecd3d7146102b557806348fa1156146102db57806355bcbf57146103aa57806358975919146103d0578063696f8c57146103fc5780637145644c1461044e57806390e102501461047c5780639b5d25d91461048e578063a47d1f5e146104ca578063af7fdd7614610563575bfe5b34156100d857fe5b610104600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610575565b005b341561010e57fe5b610124600480803590602001909190505061078e565b604051808761ffff1661ffff168152602001806020018661ffff1661ffff168152602001806020018560070b60070b81526020018460ff1660ff1681526020018381038352888181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156101e55780601f106101ba576101008083540402835291602001916101e5565b820191906000526020600020905b8154815290600101906020018083116101c857829003601f168201915b50508381038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102685780601f1061023d57610100808354040283529160200191610268565b820191906000526020600020905b81548152906001019060200180831161024b57829003601f168201915b50509850505050505050505060405180910390f35b341561028557fe5b61028d6107fe565b6040518082815260200191505060405180910390f35b34156102ab57fe5b6102b3610804565b005b34156102bd57fe5b6102c5610898565b6040518082815260200191505060405180910390f35b34156102e357fe5b6103a8600480803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803560070b90602001909190803560ff1690602001909190505061089e565b005b34156103b257fe5b6103ba610a47565b6040518082815260200191505060405180910390f35b34156103d857fe5b6103e0610a4d565b604051808260070b60070b815260200191505060405180910390f35b341561040457fe5b61040c610a60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561045657fe5b61045e610a86565b604051808261ffff1661ffff16815260200191505060405180910390f35b341561048457fe5b61048c610d2f565b005b341561049657fe5b6104ac6004808035906020019091905050610d39565b604051808261ffff1661ffff16815260200191505060405180910390f35b34156104d257fe5b6104da610d71565b6040518080602001828103825283818151815260200191508051906020019080838360008314610529575b80518252602083111561052957602082019150602081019050602083039250610505565b505050905090810190601f1680156105555780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561056b57fe5b610573610e0f565b005b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d7806c336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561063757fe5b6102c65a03f1151561064557fe5b50505060405180519050151561065b5760006000fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632cdad41c826000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561071d57fe5b6102c65a03f1151561072b57fe5b5050506040518051905015156107415760006000fd5b80600360086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061078a610e0f565b5b50565b60086020528060005260406000206000915090508060000160009054906101000a900461ffff169080600101908060020160009054906101000a900461ffff169080600301908060040160009054906101000a900460070b908060040160089054906101000a900460ff16905086565b60015481565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561089557600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60025481565b6000600360089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108fd5760006000fd5b6007600081548092919060010191905055905060c0604051908101604052808861ffff1681526020018781526020018661ffff1681526020018581526020018460070b81526020018360ff168152506008600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff160217905550602082015181600101908051906020019061099e929190610e19565b5060408201518160020160006101000a81548161ffff021916908361ffff16021790555060608201518160030190805190602001906109de929190610e19565b5060808201518160040160006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060a08201518160040160086101000a81548160ff021916908360ff160217905550905050610a3d610e0f565b5b50505050505050565b60075481565b600360009054906101000a900460070b81565b600360089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600060006000600060009350600092505b600754831015610cae576008600084815260200190815260200160002091506115288260000160009054906101000a900461ffff1661ffff161415610ca057600960008360020160009054906101000a900461ffff1661ffff1661ffff168152602001908152602001600020905060001515600960008460020160009054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002060000160049054906101000a900460ff1615151415610c5757606060405190810160405280600061ffff1681526020018360040160089054906101000a900460ff1660ff1661ffff16815260200160011515815250600960008460020160009054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a81548161ffff021916908361ffff16021790555060408201518160000160046101000a81548160ff021916908315150217905550905050600960008360020160009054906101000a900461ffff1661ffff1661ffff1681526020019081526020016000209050610c9f565b8160040160089054906101000a900460ff1660ff168160000160029054906101000a900461ffff16018160000160026101000a81548161ffff021916908361ffff1602179055505b5b5b8280600101935050610a99565b600092505b600a831015610d245760096000600a85600a81101515610ccf57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002060000160029054906101000a900461ffff16840193505b8280600101935050610cb3565b8394505b5050505090565b426001819055505b565b600481815481101515610d4857fe5b90600052602060002090601091828204019190066002025b915054906101000a900461ffff1681565b60058054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e075780601f10610ddc57610100808354040283529160200191610e07565b820191906000526020600020905b815481529060010190602001808311610dea57829003601f168201915b505050505081565b426002819055505b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e5a57805160ff1916838001178555610e88565b82800160010185558215610e88579182015b82811115610e87578251825591602001919060010190610e6c565b5b509050610e959190610e99565b5090565b610ebb91905b80821115610eb7576000816000905550600101610e9f565b5090565b905600a165627a7a7230582059928c2446ff11a664c7479abd335ee8fb88c9821be3c21d5950c55cdabceda10029`

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

// CalculateBTS is a free data retrieval call binding the contract method 0x7145644c.
//
// Solidity: function calculateBTS() constant returns(uint16)
func (_RequestContract *RequestContractCaller) CalculateBTS(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "calculateBTS")
	return *ret0, err
}

// CalculateBTS is a free data retrieval call binding the contract method 0x7145644c.
//
// Solidity: function calculateBTS() constant returns(uint16)
func (_RequestContract *RequestContractSession) CalculateBTS() (uint16, error) {
	return _RequestContract.Contract.CalculateBTS(&_RequestContract.CallOpts)
}

// CalculateBTS is a free data retrieval call binding the contract method 0x7145644c.
//
// Solidity: function calculateBTS() constant returns(uint16)
func (_RequestContract *RequestContractCallerSession) CalculateBTS() (uint16, error) {
	return _RequestContract.Contract.CalculateBTS(&_RequestContract.CallOpts)
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
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId string, pointGroupCode uint16, controlPointId string, lackId int64, points uint8)
func (_RequestContract *RequestContractCaller) Lacks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId string
	PointGroupCode    uint16
	ControlPointId    string
	LackId            int64
	Points            uint8
}, error) {
	ret := new(struct {
		ContributionCode  uint16
		ControlCategoryId string
		PointGroupCode    uint16
		ControlPointId    string
		LackId            int64
		Points            uint8
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "lacks", arg0)
	return *ret, err
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId string, pointGroupCode uint16, controlPointId string, lackId int64, points uint8)
func (_RequestContract *RequestContractSession) Lacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId string
	PointGroupCode    uint16
	ControlPointId    string
	LackId            int64
	Points            uint8
}, error) {
	return _RequestContract.Contract.Lacks(&_RequestContract.CallOpts, arg0)
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId string, pointGroupCode uint16, controlPointId string, lackId int64, points uint8)
func (_RequestContract *RequestContractCallerSession) Lacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId string
	PointGroupCode    uint16
	ControlPointId    string
	LackId            int64
	Points            uint8
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

// AddLack is a paid mutator transaction binding the contract method 0x48fa1156.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, controlPointId string, lackId int64, points uint8) returns()
func (_RequestContract *RequestContractTransactor) AddLack(opts *bind.TransactOpts, _contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, controlPointId string, lackId int64, points uint8) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "addLack", _contributionCode, _controlCategoryId, _pointGroupCode, controlPointId, lackId, points)
}

// AddLack is a paid mutator transaction binding the contract method 0x48fa1156.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, controlPointId string, lackId int64, points uint8) returns()
func (_RequestContract *RequestContractSession) AddLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, controlPointId string, lackId int64, points uint8) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLack(&_RequestContract.TransactOpts, _contributionCode, _controlCategoryId, _pointGroupCode, controlPointId, lackId, points)
}

// AddLack is a paid mutator transaction binding the contract method 0x48fa1156.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, controlPointId string, lackId int64, points uint8) returns()
func (_RequestContract *RequestContractTransactorSession) AddLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, controlPointId string, lackId int64, points uint8) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLack(&_RequestContract.TransactOpts, _contributionCode, _controlCategoryId, _pointGroupCode, controlPointId, lackId, points)
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

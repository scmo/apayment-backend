// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package directpaymentrequest

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RequestContractABI is the input ABI used to generate the binding from.
const RequestContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_inspectorAddress\",\"type\":\"address\"}],\"name\":\"setInspectorId\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gve1110\",\"type\":\"uint16\"},{\"name\":\"_gve1150\",\"type\":\"uint16\"},{\"name\":\"_gve1128\",\"type\":\"uint16\"},{\"name\":\"_gve1141\",\"type\":\"uint16\"},{\"name\":\"_gve1142\",\"type\":\"uint16\"},{\"name\":\"_gve1124\",\"type\":\"uint16\"},{\"name\":\"_gve1129\",\"type\":\"uint16\"},{\"name\":\"_gve1143\",\"type\":\"uint16\"},{\"name\":\"_gve1144\",\"type\":\"uint16\"}],\"name\":\"setGVE\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"string\"},{\"name\":\"pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"},{\"name\":\"points\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AddPayment\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modified\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCode\",\"type\":\"uint16\"},{\"name\":\"_controlCategoryId\",\"type\":\"string\"},{\"name\":\"_pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"_controlPointId\",\"type\":\"string\"},{\"name\":\"_lackId\",\"type\":\"int64\"},{\"name\":\"_points\",\"type\":\"uint8\"}],\"name\":\"addLack\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numLacks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userId\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspectorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"calculateBTS\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pointGroups\",\"outputs\":[{\"name\":\"gve\",\"type\":\"uint16\"},{\"name\":\"points\",\"type\":\"uint16\"},{\"name\":\"btsPoints\",\"type\":\"uint16\"},{\"name\":\"btsTotal\",\"type\":\"uint32\"},{\"name\":\"btsDeduction\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"_points\",\"type\":\"uint16\"}],\"name\":\"updateBtsPoint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setCreated\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributionCodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setModified\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentList\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_userId\",\"type\":\"int64\"},{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_remark\",\"type\":\"string\"},{\"name\":\"rbacAddress\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// RequestContractBin is the compiled bytecode used for deploying new contracts.
const RequestContractBin = `60606040526101206040519081016040528061045661ffff16815260200161047e61ffff16815260200161046861ffff16815260200161047561ffff16815260200161047661ffff16815260200161046461ffff16815260200161046961ffff16815260200161047761ffff16815260200161047861ffff1681525060099060096200008d929190620001e9565b5034156200009757fe5b6040516200197238038062001972833981016040528080519060200190919080518201919060200180518201919060200180519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555082600490805190602001906200019c9291906200028d565b508160059080519060200190620001b59291906200033e565b50620001d4620001df6401000000000262001335176401000000009004565b5b5050505062000421565b426001819055505b565b826009600f016010900481019282156200027a5791602002820160005b838211156200024857835183826101000a81548161ffff021916908361ffff160217905550926020019260020160208160010104928301926001030262000206565b8015620002785782816101000a81549061ffff021916905560020160208160010104928301926001030262000248565b505b509050620002899190620003c5565b5090565b82805482825590600052602060002090600f016010900481019282156200032b5791602002820160005b83821115620002f957835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302620002b7565b8015620003295782816101000a81549061ffff0219169055600201602081600101049283019260010302620002f9565b505b5090506200033a9190620003c5565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200038157805160ff1916838001178555620003b2565b82800160010185558215620003b2579182015b82811115620003b157825182559160200191906001019062000394565b5b509050620003c19190620003f9565b5090565b620003f691905b80821115620003f257600081816101000a81549061ffff021916905550600101620003cc565b5090565b90565b6200041e91905b808211156200041a57600081600090555060010162000400565b5090565b90565b61154180620004316000396000f30060606040523615610110576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063021c7bd7146101125780630284bd8f146101485780630478e042146101d4578063325a19f11461034b57806333bdd1a31461037157806341c0e1b5146103b057806345ecd3d7146103c257806348fa1156146103e857806355bcbf57146104b757806358975919146104dd578063696f8c57146105095780637145644c1461055b5780637184f7251461056d57806387d81789146105f15780638a0011261461065f57806390e10250146106905780639b5d25d9146106a2578063a47d1f5e146106de578063af7fdd7614610777578063ccc7905814610789575bfe5b341561011a57fe5b610146600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506107bd565b005b341561015057fe5b6101d2600480803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff16906020019091905050610ab7565b005b34156101dc57fe5b6101f26004808035906020019091905050610cd0565b604051808761ffff1661ffff168152602001806020018661ffff1661ffff168152602001806020018560070b60070b81526020018460ff1660ff1681526020018381038352888181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102b35780601f10610288576101008083540402835291602001916102b3565b820191906000526020600020905b81548152906001019060200180831161029657829003601f168201915b50508381038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156103365780601f1061030b57610100808354040283529160200191610336565b820191906000526020600020905b81548152906001019060200180831161031957829003601f168201915b50509850505050505050505060405180910390f35b341561035357fe5b61035b610d40565b6040518082815260200191505060405180910390f35b341561037957fe5b6103ae600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610d46565b005b34156103b857fe5b6103c0610deb565b005b34156103ca57fe5b6103d2610e7f565b6040518082815260200191505060405180910390f35b34156103f057fe5b6104b5600480803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803560070b90602001909190803560ff16906020019091905050610e85565b005b34156104bf57fe5b6104c761104a565b6040518082815260200191505060405180910390f35b34156104e557fe5b6104ed611050565b604051808260070b60070b815260200191505060405180910390f35b341561051157fe5b610519611063565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561056357fe5b61056b611089565b005b341561057557fe5b61058f600480803561ffff16906020019091905050611213565b604051808661ffff1661ffff1681526020018561ffff1661ffff1681526020018461ffff1661ffff1681526020018363ffffffff1663ffffffff1681526020018263ffffffff1663ffffffff1681526020019550505050505060405180910390f35b34156105f957fe5b61060f6004808035906020019091905050611293565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390f35b341561066757fe5b61068e600480803561ffff1690602001909190803561ffff169060200190919050506112dd565b005b341561069857fe5b6106a0611335565b005b34156106aa57fe5b6106c0600480803590602001909190505061133f565b604051808261ffff1661ffff16815260200191505060405180910390f35b34156106e657fe5b6106ee611377565b604051808060200182810382528381815181526020019150805190602001908083836000831461073d575b80518252602083111561073d57602082019150602081019050602083039250610719565b505050905090810190601f1680156107695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561077f57fe5b610787611415565b005b341561079157fe5b6107a7600480803590602001909190505061141f565b6040518082815260200191505060405180910390f35b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d7806c336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561087f57fe5b6102c65a03f1151561088d57fe5b50505060405180519050806109785750600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636bb164c9336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561095f57fe5b6102c65a03f1151561096d57fe5b505050604051805190505b15156109845760006000fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632cdad41c826000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515610a4657fe5b6102c65a03f11515610a5457fe5b505050604051805190501515610a6a5760006000fd5b80600360086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ab3611415565b5b50565b6000600a600061045661ffff1681526020019081526020016000209050898160000160006101000a81548161ffff021916908361ffff160217905550600a600061047e61ffff1681526020019081526020016000209050888160000160006101000a81548161ffff021916908361ffff160217905550600a600061046861ffff1681526020019081526020016000209050878160000160006101000a81548161ffff021916908361ffff160217905550600a600061047561ffff1681526020019081526020016000209050868160000160006101000a81548161ffff021916908361ffff160217905550600a600061047661ffff1681526020019081526020016000209050858160000160006101000a81548161ffff021916908361ffff160217905550600a600061046461ffff1681526020019081526020016000209050848160000160006101000a81548161ffff021916908361ffff160217905550600a600061046961ffff1681526020019081526020016000209050838160000160006101000a81548161ffff021916908361ffff160217905550600a600061047761ffff1681526020019081526020016000209050828160000160006101000a81548161ffff021916908361ffff160217905550600a600061047861ffff1681526020019081526020016000209050818160000160006101000a81548161ffff021916908361ffff1602179055505b50505050505050505050565b60086020528060005260406000206000915090508060000160009054906101000a900461ffff169080600101908060020160009054906101000a900461ffff169080600301908060040160009054906101000a900460070b908060040160089054906101000a900460ff16905086565b60015481565b600042905082600b600083815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600b600083815260200190815260200160002060010181905550600c8054806001018281610dcf9190611444565b916000526020600020900160005b83909190915055505b505050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610e7c57600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60025481565b6000600360089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ee45760006000fd5b6007600081548092919060010191905055905060c0604051908101604052808861ffff1681526020018781526020018661ffff1681526020018581526020018460070b81526020018360ff168152506008600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001019080519060200190610f85929190611470565b5060408201518160020160006101000a81548161ffff021916908361ffff1602179055506060820151816003019080519060200190610fc5929190611470565b5060808201518160040160006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060a08201518160040160086101000a81548160ff021916908360ff160217905550905050611024611415565b6115288761ffff1614156110405761103f858360ff166112dd565b5b5b50505050505050565b60075481565b600360009054906101000a900460070b81565b600360089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006000600091505b60098261ffff16101561120e57600a600060098461ffff166009811015156110b657fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff168152602001908152602001600020905061047660098361ffff166009811015156110ff57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1614158061115e575061047860098361ffff1660098110151561113c57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1614155b1561120057606e8160000160049054906101000a900461ffff1661ffff1610156111ff578060000160009054906101000a900461ffff166123280261ffff168160000160066101000a81548163ffffffff021916908363ffffffff160217905550605a600a8260000160049054906101000a900461ffff16030261ffff1681600001600a6101000a81548163ffffffff021916908363ffffffff1602179055505b5b5b8180600101925050611092565b5b5050565b600a6020528060005260406000206000915090508060000160009054906101000a900461ffff16908060000160029054906101000a900461ffff16908060000160049054906101000a900461ffff16908060000160069054906101000a900463ffffffff169080600001600a9054906101000a900463ffffffff16905085565b600b6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b6000600a60008461ffff1661ffff1681526020019081526020016000209050818160000160049054906101000a900461ffff16018160000160046101000a81548161ffff021916908361ffff1602179055505b505050565b426001819055505b565b60048181548110151561134e57fe5b90600052602060002090601091828204019190066002025b915054906101000a900461ffff1681565b60058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561140d5780601f106113e25761010080835404028352916020019161140d565b820191906000526020600020905b8154815290600101906020018083116113f057829003601f168201915b505050505081565b426002819055505b565b600c8181548110151561142e57fe5b906000526020600020900160005b915090505481565b81548183558181151161146b5781836000526020600020918201910161146a91906114f0565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114b157805160ff19168380011785556114df565b828001600101855582156114df579182015b828111156114de5782518255916020019190600101906114c3565b5b5090506114ec91906114f0565b5090565b61151291905b8082111561150e5760008160009055506001016114f6565b5090565b905600a165627a7a72305820b2ebc4936c42fde76994fa73c8d67c0264b9e99dbb582041e7e36f70dd827ef40029`

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

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList( uint256) constant returns(uint256)
func (_RequestContract *RequestContractCaller) PaymentList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "paymentList", arg0)
	return *ret0, err
}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList( uint256) constant returns(uint256)
func (_RequestContract *RequestContractSession) PaymentList(arg0 *big.Int) (*big.Int, error) {
	return _RequestContract.Contract.PaymentList(&_RequestContract.CallOpts, arg0)
}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList( uint256) constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) PaymentList(arg0 *big.Int) (*big.Int, error) {
	return _RequestContract.Contract.PaymentList(&_RequestContract.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments( uint256) constant returns(from address, amount uint256, timestamp uint256)
func (_RequestContract *RequestContractCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		From      common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "payments", arg0)
	return *ret, err
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments( uint256) constant returns(from address, amount uint256, timestamp uint256)
func (_RequestContract *RequestContractSession) Payments(arg0 *big.Int) (struct {
	From      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _RequestContract.Contract.Payments(&_RequestContract.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments( uint256) constant returns(from address, amount uint256, timestamp uint256)
func (_RequestContract *RequestContractCallerSession) Payments(arg0 *big.Int) (struct {
	From      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _RequestContract.Contract.Payments(&_RequestContract.CallOpts, arg0)
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint16, points uint16, btsPoints uint16, btsTotal uint32, btsDeduction uint32)
func (_RequestContract *RequestContractCaller) PointGroups(opts *bind.CallOpts, arg0 uint16) (struct {
	Gve          uint16
	Points       uint16
	BtsPoints    uint16
	BtsTotal     uint32
	BtsDeduction uint32
}, error) {
	ret := new(struct {
		Gve          uint16
		Points       uint16
		BtsPoints    uint16
		BtsTotal     uint32
		BtsDeduction uint32
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "pointGroups", arg0)
	return *ret, err
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint16, points uint16, btsPoints uint16, btsTotal uint32, btsDeduction uint32)
func (_RequestContract *RequestContractSession) PointGroups(arg0 uint16) (struct {
	Gve          uint16
	Points       uint16
	BtsPoints    uint16
	BtsTotal     uint32
	BtsDeduction uint32
}, error) {
	return _RequestContract.Contract.PointGroups(&_RequestContract.CallOpts, arg0)
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint16, points uint16, btsPoints uint16, btsTotal uint32, btsDeduction uint32)
func (_RequestContract *RequestContractCallerSession) PointGroups(arg0 uint16) (struct {
	Gve          uint16
	Points       uint16
	BtsPoints    uint16
	BtsTotal     uint32
	BtsDeduction uint32
}, error) {
	return _RequestContract.Contract.PointGroups(&_RequestContract.CallOpts, arg0)
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

// AddPayment is a paid mutator transaction binding the contract method 0x33bdd1a3.
//
// Solidity: function AddPayment(_from address, _amount uint256) returns()
func (_RequestContract *RequestContractTransactor) AddPayment(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "AddPayment", _from, _amount)
}

// AddPayment is a paid mutator transaction binding the contract method 0x33bdd1a3.
//
// Solidity: function AddPayment(_from address, _amount uint256) returns()
func (_RequestContract *RequestContractSession) AddPayment(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RequestContract.Contract.AddPayment(&_RequestContract.TransactOpts, _from, _amount)
}

// AddPayment is a paid mutator transaction binding the contract method 0x33bdd1a3.
//
// Solidity: function AddPayment(_from address, _amount uint256) returns()
func (_RequestContract *RequestContractTransactorSession) AddPayment(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RequestContract.Contract.AddPayment(&_RequestContract.TransactOpts, _from, _amount)
}

// AddLack is a paid mutator transaction binding the contract method 0x48fa1156.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, _controlPointId string, _lackId int64, _points uint8) returns()
func (_RequestContract *RequestContractTransactor) AddLack(opts *bind.TransactOpts, _contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, _controlPointId string, _lackId int64, _points uint8) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "addLack", _contributionCode, _controlCategoryId, _pointGroupCode, _controlPointId, _lackId, _points)
}

// AddLack is a paid mutator transaction binding the contract method 0x48fa1156.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, _controlPointId string, _lackId int64, _points uint8) returns()
func (_RequestContract *RequestContractSession) AddLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, _controlPointId string, _lackId int64, _points uint8) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLack(&_RequestContract.TransactOpts, _contributionCode, _controlCategoryId, _pointGroupCode, _controlPointId, _lackId, _points)
}

// AddLack is a paid mutator transaction binding the contract method 0x48fa1156.
//
// Solidity: function addLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, _controlPointId string, _lackId int64, _points uint8) returns()
func (_RequestContract *RequestContractTransactorSession) AddLack(_contributionCode uint16, _controlCategoryId string, _pointGroupCode uint16, _controlPointId string, _lackId int64, _points uint8) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLack(&_RequestContract.TransactOpts, _contributionCode, _controlCategoryId, _pointGroupCode, _controlPointId, _lackId, _points)
}

// CalculateBTS is a paid mutator transaction binding the contract method 0x7145644c.
//
// Solidity: function calculateBTS() returns()
func (_RequestContract *RequestContractTransactor) CalculateBTS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "calculateBTS")
}

// CalculateBTS is a paid mutator transaction binding the contract method 0x7145644c.
//
// Solidity: function calculateBTS() returns()
func (_RequestContract *RequestContractSession) CalculateBTS() (*types.Transaction, error) {
	return _RequestContract.Contract.CalculateBTS(&_RequestContract.TransactOpts)
}

// CalculateBTS is a paid mutator transaction binding the contract method 0x7145644c.
//
// Solidity: function calculateBTS() returns()
func (_RequestContract *RequestContractTransactorSession) CalculateBTS() (*types.Transaction, error) {
	return _RequestContract.Contract.CalculateBTS(&_RequestContract.TransactOpts)
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

// SetGVE is a paid mutator transaction binding the contract method 0x0284bd8f.
//
// Solidity: function setGVE(_gve1110 uint16, _gve1150 uint16, _gve1128 uint16, _gve1141 uint16, _gve1142 uint16, _gve1124 uint16, _gve1129 uint16, _gve1143 uint16, _gve1144 uint16) returns()
func (_RequestContract *RequestContractTransactor) SetGVE(opts *bind.TransactOpts, _gve1110 uint16, _gve1150 uint16, _gve1128 uint16, _gve1141 uint16, _gve1142 uint16, _gve1124 uint16, _gve1129 uint16, _gve1143 uint16, _gve1144 uint16) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "setGVE", _gve1110, _gve1150, _gve1128, _gve1141, _gve1142, _gve1124, _gve1129, _gve1143, _gve1144)
}

// SetGVE is a paid mutator transaction binding the contract method 0x0284bd8f.
//
// Solidity: function setGVE(_gve1110 uint16, _gve1150 uint16, _gve1128 uint16, _gve1141 uint16, _gve1142 uint16, _gve1124 uint16, _gve1129 uint16, _gve1143 uint16, _gve1144 uint16) returns()
func (_RequestContract *RequestContractSession) SetGVE(_gve1110 uint16, _gve1150 uint16, _gve1128 uint16, _gve1141 uint16, _gve1142 uint16, _gve1124 uint16, _gve1129 uint16, _gve1143 uint16, _gve1144 uint16) (*types.Transaction, error) {
	return _RequestContract.Contract.SetGVE(&_RequestContract.TransactOpts, _gve1110, _gve1150, _gve1128, _gve1141, _gve1142, _gve1124, _gve1129, _gve1143, _gve1144)
}

// SetGVE is a paid mutator transaction binding the contract method 0x0284bd8f.
//
// Solidity: function setGVE(_gve1110 uint16, _gve1150 uint16, _gve1128 uint16, _gve1141 uint16, _gve1142 uint16, _gve1124 uint16, _gve1129 uint16, _gve1143 uint16, _gve1144 uint16) returns()
func (_RequestContract *RequestContractTransactorSession) SetGVE(_gve1110 uint16, _gve1150 uint16, _gve1128 uint16, _gve1141 uint16, _gve1142 uint16, _gve1124 uint16, _gve1129 uint16, _gve1143 uint16, _gve1144 uint16) (*types.Transaction, error) {
	return _RequestContract.Contract.SetGVE(&_RequestContract.TransactOpts, _gve1110, _gve1150, _gve1128, _gve1141, _gve1142, _gve1124, _gve1129, _gve1143, _gve1144)
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

// UpdateBtsPoint is a paid mutator transaction binding the contract method 0x8a001126.
//
// Solidity: function updateBtsPoint(_pointGroupCode uint16, _points uint16) returns()
func (_RequestContract *RequestContractTransactor) UpdateBtsPoint(opts *bind.TransactOpts, _pointGroupCode uint16, _points uint16) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "updateBtsPoint", _pointGroupCode, _points)
}

// UpdateBtsPoint is a paid mutator transaction binding the contract method 0x8a001126.
//
// Solidity: function updateBtsPoint(_pointGroupCode uint16, _points uint16) returns()
func (_RequestContract *RequestContractSession) UpdateBtsPoint(_pointGroupCode uint16, _points uint16) (*types.Transaction, error) {
	return _RequestContract.Contract.UpdateBtsPoint(&_RequestContract.TransactOpts, _pointGroupCode, _points)
}

// UpdateBtsPoint is a paid mutator transaction binding the contract method 0x8a001126.
//
// Solidity: function updateBtsPoint(_pointGroupCode uint16, _points uint16) returns()
func (_RequestContract *RequestContractTransactorSession) UpdateBtsPoint(_pointGroupCode uint16, _points uint16) (*types.Transaction, error) {
	return _RequestContract.Contract.UpdateBtsPoint(&_RequestContract.TransactOpts, _pointGroupCode, _points)
}

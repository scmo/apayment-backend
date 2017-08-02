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
const RequestContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_inspectorAddress\",\"type\":\"address\"}],\"name\":\"setInspectorId\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"int64\"},{\"name\":\"pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"controlPointId\",\"type\":\"int64\"},{\"name\":\"lackId\",\"type\":\"int64\"},{\"name\":\"points\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_controlCategoryIds\",\"type\":\"int64[]\"},{\"name\":\"_pointGroupCodes\",\"type\":\"uint16[]\"},{\"name\":\"_controlPointIds\",\"type\":\"int64[]\"},{\"name\":\"_lackIds\",\"type\":\"int64[]\"},{\"name\":\"_points\",\"type\":\"uint8[]\"}],\"name\":\"addLacks\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modified\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFinalPaymentAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numLacks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspectorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pointGroups\",\"outputs\":[{\"name\":\"gve\",\"type\":\"uint32\"},{\"name\":\"btsPoints\",\"type\":\"uint16\"},{\"name\":\"btsTotal\",\"type\":\"uint256\"},{\"name\":\"btsDeduction\",\"type\":\"uint256\"},{\"name\":\"rausPoints\",\"type\":\"uint16\"},{\"name\":\"rausTotal\",\"type\":\"uint256\"},{\"name\":\"rausDeduction\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributionCodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountPreviousYear\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFirstPaymentAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_remark\",\"type\":\"string\"},{\"name\":\"_rbacAddress\",\"type\":\"address\"},{\"name\":\"_gves\",\"type\":\"uint32[]\"},{\"name\":\"_amountPreviousYear\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"}]"

// RequestContractBin is the compiled bytecode used for deploying new contracts.
const RequestContractBin = `60606040526101206040519081016040528061045661ffff16815260200161047e61ffff16815260200161046861ffff16815260200161047561ffff16815260200161047661ffff16815260200161046461ffff16815260200161046961ffff16815260200161047761ffff16815260200161047861ffff1681525060049060096200008d929190620008d6565b5034156200009a57600080fd5b604051620021be380380620021be833981016040528080518201919060200180518201919060200180519060200190919080518201919060200180519060200190919050505b5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600a90805190602001906200017b9291906200097a565b5083600b90805190602001906200019492919062000a2b565b506200029d826000815181101515620001a957fe5b90602001906020020151836001815181101515620001c357fe5b90602001906020020151846002815181101515620001dd57fe5b90602001906020020151856003815181101515620001f757fe5b906020019060200201518660048151811015156200021157fe5b906020019060200201518760058151811015156200022b57fe5b906020019060200201518860068151811015156200024557fe5b906020019060200201518960078151811015156200025f57fe5b906020019060200201518a60088151811015156200027957fe5b90602001906020020151620002ce6401000000000262001415176401000000009004565b80600981905550620002c262000565640100000000026200166a176401000000009004565b5b505050505062000b0e565b60006007600061045661ffff1681526020019081526020016000209050898160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047e61ffff1681526020019081526020016000209050888160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061046861ffff1681526020019081526020016000209050878160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047561ffff1681526020019081526020016000209050868160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047661ffff1681526020019081526020016000209050858160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061046461ffff1681526020019081526020016000209050848160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061046961ffff1681526020019081526020016000209050838160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047761ffff1681526020019081526020016000209050828160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047861ffff1681526020019081526020016000209050818160000160006101000a81548163ffffffff021916908363ffffffff1602179055506200051c6200056f64010000000002620010af176401000000009004565b6200053a620007116401000000000262001246176401000000009004565b62000558620008cc6401000000000262000ff5176401000000009004565b5b50505050505050505050565b426001819055505b565b600080600091505b60098261ffff1610156200070c576007600060048461ffff166009811015156200059d57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002090506000151561047660048461ffff16600981101515620005eb57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1614806200064a575061047860048461ffff166009811015156200062957fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16145b15151415620006fd576123288160000160009054906101000a900463ffffffff1663ffffffff1602816001018190555060008160000160049054906101000a900461ffff1661ffff161415620006a057620006fe565b606e8160000160049054906101000a900461ffff1661ffff161115620006d35780600101548160020181905550620006fe565b612710605a600a8360000160049054906101000a900461ffff160361ffff16020281600201819055505b5b818060010192505062000577565b5b5050565b60008060008092505b60098361ffff161015620008c6576007600060048561ffff166009811015156200074057fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002091506000905061047660048461ffff166009811015156200078e57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff161480620007ed575061047860048461ffff16600981101515620007cc57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16145b15620007fe57619088905062000804565b614a3890505b808260000160009054906101000a900463ffffffff1663ffffffff1602826004018190555060008260030160009054906101000a900461ffff1661ffff1614156200084f57620008b8565b606e8260030160009054906101000a900461ffff1661ffff161115620008825781600401548260050181905550620008b8565b6127106064828115156200089257fe5b04600a8460030160009054906101000a900461ffff160361ffff16020282600501819055505b82806001019350506200071a565b5b505050565b426002819055505b565b826009600f01601090048101928215620009675791602002820160005b838211156200093557835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302620008f3565b8015620009655782816101000a81549061ffff021916905560020160208160010104928301926001030262000935565b505b50905062000976919062000ab2565b5090565b82805482825590600052602060002090600f0160109004810192821562000a185791602002820160005b83821115620009e657835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302620009a4565b801562000a165782816101000a81549061ffff0219169055600201602081600101049283019260010302620009e6565b505b50905062000a27919062000ab2565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1062000a6e57805160ff191683800117855562000a9f565b8280016001018555821562000a9f579182015b8281111562000a9e57825182559160200191906001019062000a81565b5b50905062000aae919062000ae6565b5090565b62000ae391905b8082111562000adf57600081816101000a81549061ffff02191690555060010162000ab9565b5090565b90565b62000b0b91905b8082111562000b0757600081600090555060010162000aed565b5090565b90565b6116a08062000b1e6000396000f300606060405236156100ce576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063021c7bd7146100d35780630478e0421461010c578063325a19f11461018e57806338d1cc6a146101b757806341c0e1b51461035157806345ecd3d7146103665780634f3ded8a1461038f57806355bcbf57146103b8578063696f8c57146103e15780637184f725146104365780639b5d25d9146104b7578063a47d1f5e146104f6578063ce30ce4b14610585578063e0f8c670146105ae575b600080fd5b34156100de57600080fd5b61010a600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506105d7565b005b341561011757600080fd5b61012d60048080359060200190919050506108e1565b604051808761ffff1661ffff1681526020018660070b60070b81526020018561ffff1661ffff1681526020018460070b60070b81526020018360070b60070b81526020018260ff1660ff168152602001965050505050505060405180910390f35b341561019957600080fd5b6101a161096d565b6040518082815260200191505060405180910390f35b34156101c257600080fd5b61034f6004808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091905050610973565b005b341561035c57600080fd5b610364610cf9565b005b341561037157600080fd5b610379610d8b565b6040518082815260200191505060405180910390f35b341561039a57600080fd5b6103a2610d91565b6040518082815260200191505060405180910390f35b34156103c357600080fd5b6103cb610e53565b6040518082815260200191505060405180910390f35b34156103ec57600080fd5b6103f4610e59565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561044157600080fd5b61045b600480803561ffff16906020019091905050610e7f565b604051808863ffffffff1663ffffffff1681526020018761ffff1661ffff1681526020018681526020018581526020018461ffff1661ffff16815260200183815260200182815260200197505050505050505060405180910390f35b34156104c257600080fd5b6104d86004808035906020019091905050610eed565b604051808261ffff1661ffff16815260200191505060405180910390f35b341561050157600080fd5b610509610f25565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561054a5780820151818401525b60208101905061052e565b50505050905090810190601f1680156105775780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561059057600080fd5b610598610fc3565b6040518082815260200191505060405180910390f35b34156105b957600080fd5b6105c1610fc9565b6040518082815260200191505060405180910390f35b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d7806c336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561069c57600080fd5b6102c65a03f115156106ad57600080fd5b505050604051805190508061079e5750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636bb164c9336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561078257600080fd5b6102c65a03f1151561079357600080fd5b505050604051805190505b15156107a957600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632cdad41c826000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561086e57600080fd5b6102c65a03f1151561087f57600080fd5b50505060405180519050151561089457600080fd5b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506108dd610ff5565b5b50565b60066020528060005260406000206000915090508060000160009054906101000a900461ffff16908060000160029054906101000a900460070b9080600001600a9054906101000a900461ffff169080600001600c9054906101000a900460070b908060000160149054906101000a900460070b9080600001601c9054906101000a900460ff16905086565b60015481565b600080600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109d257600080fd5b600091505b87518261ffff161015610cd6576005600081548092919060010191905055905060c060405190810160405280898461ffff16815181101515610a1557fe5b9060200190602002015161ffff168152602001888461ffff16815181101515610a3a57fe5b9060200190602002015160070b8152602001878461ffff16815181101515610a5e57fe5b9060200190602002015161ffff168152602001868461ffff16815181101515610a8357fe5b9060200190602002015160070b8152602001858461ffff16815181101515610aa757fe5b9060200190602002015160070b8152602001848461ffff16815181101515610acb57fe5b9060200190602002015160ff168152506006600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff160217905550604082015181600001600a6101000a81548161ffff021916908361ffff160217905550606082015181600001600c6101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060808201518160000160146101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060a082015181600001601c6101000a81548160ff021916908360ff160217905550905050611528888361ffff16815181101515610c0357fe5b9060200190602002015161ffff161415610c5b57610c5a868361ffff16815181101515610c2c57fe5b90602001906020020151848461ffff16815181101515610c4857fe5b9060200190602002015160ff16610fff565b5b611529888361ffff16815181101515610c7057fe5b9060200190602002015161ffff161415610cc857610cc7868361ffff16815181101515610c9957fe5b90602001906020020151848461ffff16815181101515610cb557fe5b9060200190602002015160ff16611057565b5b5b81806001019250506109d7565b610cde6110af565b610ce6611246565b610cee610ff5565b5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610d88576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60025481565b60008060008060009250600091505b60098261ffff161015610e33576007600060048461ffff16600981101515610dc457fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff1681526020019081526020016000209050610e0a81600201548260010154036113f4565b83019250610e2181600501548260040154036113f4565b830192505b8180600101925050610da0565b610e3b610fc9565b61271084811515610e4857fe5b040393505b50505090565b60055481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60076020528060005260406000206000915090508060000160009054906101000a900463ffffffff16908060000160049054906101000a900461ffff16908060010154908060020154908060030160009054906101000a900461ffff16908060040154908060050154905087565b600a81815481101515610efc57fe5b90600052602060002090601091828204019190066002025b915054906101000a900461ffff1681565b600b8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fbb5780601f10610f9057610100808354040283529160200191610fbb565b820191906000526020600020905b815481529060010190602001808311610f9e57829003601f168201915b505050505081565b60095481565b6000806000905060006009541115610fed576002600954811515610fe957fe5b0490505b8091505b5090565b426002819055505b565b6000600760008461ffff1661ffff1681526020019081526020016000209050818160000160049054906101000a900461ffff16018160000160046101000a81548161ffff021916908361ffff1602179055505b505050565b6000600760008461ffff1661ffff1681526020019081526020016000209050818160030160009054906101000a900461ffff16018160030160006101000a81548161ffff021916908361ffff1602179055505b505050565b600080600091505b60098261ffff161015611241576007600060048461ffff166009811015156110db57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002090506000151561047660048461ffff1660098110151561112857fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff161480611185575061047860048461ffff1660098110151561116457fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16145b15151415611233576123288160000160009054906101000a900463ffffffff1663ffffffff1602816001018190555060008160000160049054906101000a900461ffff1661ffff1614156111d857611234565b606e8160000160049054906101000a900461ffff1661ffff1611156112095780600101548160020181905550611234565b612710605a600a8360000160049054906101000a900461ffff160361ffff16020281600201819055505b5b81806001019250506110b7565b5b5050565b60008060008092505b60098361ffff1610156113ee576007600060048561ffff1660098110151561127357fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002091506000905061047660048461ffff166009811015156112c057fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16148061131d575061047860048461ffff166009811015156112fc57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16145b1561132c576190889050611332565b614a3890505b808260000160009054906101000a900463ffffffff1663ffffffff1602826004018190555060008260030160009054906101000a900461ffff1661ffff16141561137b576113e1565b606e8260030160009054906101000a900461ffff1661ffff1611156113ac57816004015482600501819055506113e1565b6127106064828115156113bb57fe5b04600a8460030160009054906101000a900461ffff160361ffff16020282600501819055505b828060010193505061124f565b5b505050565b6000806113888381151561140457fe5b049050611388810291505b50919050565b60006007600061045661ffff1681526020019081526020016000209050898160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047e61ffff1681526020019081526020016000209050888160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061046861ffff1681526020019081526020016000209050878160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047561ffff1681526020019081526020016000209050868160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047661ffff1681526020019081526020016000209050858160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061046461ffff1681526020019081526020016000209050848160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061046961ffff1681526020019081526020016000209050838160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047761ffff1681526020019081526020016000209050828160000160006101000a81548163ffffffff021916908363ffffffff1602179055506007600061047861ffff1681526020019081526020016000209050818160000160006101000a81548163ffffffff021916908363ffffffff16021790555061164d6110af565b611655611246565b61165d610ff5565b5b50505050505050505050565b426001819055505b5600a165627a7a72305820a19f5479f195e4856a603eb661063197e3697e2a8138efef9581d561f7b23ef20029`

// DeployRequestContract deploys a new Ethereum contract, binding an instance of RequestContract to it.
func DeployRequestContract(auth *bind.TransactOpts, backend bind.ContractBackend, _contributionCodes []uint16, _remark string, _rbacAddress common.Address, _gves []uint32, _amountPreviousYear *big.Int) (common.Address, *types.Transaction, *RequestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestContractBin), backend, _contributionCodes, _remark, _rbacAddress, _gves, _amountPreviousYear)
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

// AmountPreviousYear is a free data retrieval call binding the contract method 0xce30ce4b.
//
// Solidity: function amountPreviousYear() constant returns(uint256)
func (_RequestContract *RequestContractCaller) AmountPreviousYear(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "amountPreviousYear")
	return *ret0, err
}

// AmountPreviousYear is a free data retrieval call binding the contract method 0xce30ce4b.
//
// Solidity: function amountPreviousYear() constant returns(uint256)
func (_RequestContract *RequestContractSession) AmountPreviousYear() (*big.Int, error) {
	return _RequestContract.Contract.AmountPreviousYear(&_RequestContract.CallOpts)
}

// AmountPreviousYear is a free data retrieval call binding the contract method 0xce30ce4b.
//
// Solidity: function amountPreviousYear() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) AmountPreviousYear() (*big.Int, error) {
	return _RequestContract.Contract.AmountPreviousYear(&_RequestContract.CallOpts)
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

// GetFinalPaymentAmount is a free data retrieval call binding the contract method 0x4f3ded8a.
//
// Solidity: function getFinalPaymentAmount() constant returns(uint256)
func (_RequestContract *RequestContractCaller) GetFinalPaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "getFinalPaymentAmount")
	return *ret0, err
}

// GetFinalPaymentAmount is a free data retrieval call binding the contract method 0x4f3ded8a.
//
// Solidity: function getFinalPaymentAmount() constant returns(uint256)
func (_RequestContract *RequestContractSession) GetFinalPaymentAmount() (*big.Int, error) {
	return _RequestContract.Contract.GetFinalPaymentAmount(&_RequestContract.CallOpts)
}

// GetFinalPaymentAmount is a free data retrieval call binding the contract method 0x4f3ded8a.
//
// Solidity: function getFinalPaymentAmount() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) GetFinalPaymentAmount() (*big.Int, error) {
	return _RequestContract.Contract.GetFinalPaymentAmount(&_RequestContract.CallOpts)
}

// GetFirstPaymentAmount is a free data retrieval call binding the contract method 0xe0f8c670.
//
// Solidity: function getFirstPaymentAmount() constant returns(uint256)
func (_RequestContract *RequestContractCaller) GetFirstPaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestContract.contract.Call(opts, out, "getFirstPaymentAmount")
	return *ret0, err
}

// GetFirstPaymentAmount is a free data retrieval call binding the contract method 0xe0f8c670.
//
// Solidity: function getFirstPaymentAmount() constant returns(uint256)
func (_RequestContract *RequestContractSession) GetFirstPaymentAmount() (*big.Int, error) {
	return _RequestContract.Contract.GetFirstPaymentAmount(&_RequestContract.CallOpts)
}

// GetFirstPaymentAmount is a free data retrieval call binding the contract method 0xe0f8c670.
//
// Solidity: function getFirstPaymentAmount() constant returns(uint256)
func (_RequestContract *RequestContractCallerSession) GetFirstPaymentAmount() (*big.Int, error) {
	return _RequestContract.Contract.GetFirstPaymentAmount(&_RequestContract.CallOpts)
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
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId int64, pointGroupCode uint16, controlPointId int64, lackId int64, points uint8)
func (_RequestContract *RequestContractCaller) Lacks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId int64
	PointGroupCode    uint16
	ControlPointId    int64
	LackId            int64
	Points            uint8
}, error) {
	ret := new(struct {
		ContributionCode  uint16
		ControlCategoryId int64
		PointGroupCode    uint16
		ControlPointId    int64
		LackId            int64
		Points            uint8
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "lacks", arg0)
	return *ret, err
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId int64, pointGroupCode uint16, controlPointId int64, lackId int64, points uint8)
func (_RequestContract *RequestContractSession) Lacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId int64
	PointGroupCode    uint16
	ControlPointId    int64
	LackId            int64
	Points            uint8
}, error) {
	return _RequestContract.Contract.Lacks(&_RequestContract.CallOpts, arg0)
}

// Lacks is a free data retrieval call binding the contract method 0x0478e042.
//
// Solidity: function lacks( uint256) constant returns(contributionCode uint16, controlCategoryId int64, pointGroupCode uint16, controlPointId int64, lackId int64, points uint8)
func (_RequestContract *RequestContractCallerSession) Lacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId int64
	PointGroupCode    uint16
	ControlPointId    int64
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

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint32, btsPoints uint16, btsTotal uint256, btsDeduction uint256, rausPoints uint16, rausTotal uint256, rausDeduction uint256)
func (_RequestContract *RequestContractCaller) PointGroups(opts *bind.CallOpts, arg0 uint16) (struct {
	Gve           uint32
	BtsPoints     uint16
	BtsTotal      *big.Int
	BtsDeduction  *big.Int
	RausPoints    uint16
	RausTotal     *big.Int
	RausDeduction *big.Int
}, error) {
	ret := new(struct {
		Gve           uint32
		BtsPoints     uint16
		BtsTotal      *big.Int
		BtsDeduction  *big.Int
		RausPoints    uint16
		RausTotal     *big.Int
		RausDeduction *big.Int
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "pointGroups", arg0)
	return *ret, err
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint32, btsPoints uint16, btsTotal uint256, btsDeduction uint256, rausPoints uint16, rausTotal uint256, rausDeduction uint256)
func (_RequestContract *RequestContractSession) PointGroups(arg0 uint16) (struct {
	Gve           uint32
	BtsPoints     uint16
	BtsTotal      *big.Int
	BtsDeduction  *big.Int
	RausPoints    uint16
	RausTotal     *big.Int
	RausDeduction *big.Int
}, error) {
	return _RequestContract.Contract.PointGroups(&_RequestContract.CallOpts, arg0)
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint32, btsPoints uint16, btsTotal uint256, btsDeduction uint256, rausPoints uint16, rausTotal uint256, rausDeduction uint256)
func (_RequestContract *RequestContractCallerSession) PointGroups(arg0 uint16) (struct {
	Gve           uint32
	BtsPoints     uint16
	BtsTotal      *big.Int
	BtsDeduction  *big.Int
	RausPoints    uint16
	RausTotal     *big.Int
	RausDeduction *big.Int
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

// AddLacks is a paid mutator transaction binding the contract method 0x38d1cc6a.
//
// Solidity: function addLacks(_contributionCodes uint16[], _controlCategoryIds int64[], _pointGroupCodes uint16[], _controlPointIds int64[], _lackIds int64[], _points uint8[]) returns()
func (_RequestContract *RequestContractTransactor) AddLacks(opts *bind.TransactOpts, _contributionCodes []uint16, _controlCategoryIds []int64, _pointGroupCodes []uint16, _controlPointIds []int64, _lackIds []int64, _points []uint8) (*types.Transaction, error) {
	return _RequestContract.contract.Transact(opts, "addLacks", _contributionCodes, _controlCategoryIds, _pointGroupCodes, _controlPointIds, _lackIds, _points)
}

// AddLacks is a paid mutator transaction binding the contract method 0x38d1cc6a.
//
// Solidity: function addLacks(_contributionCodes uint16[], _controlCategoryIds int64[], _pointGroupCodes uint16[], _controlPointIds int64[], _lackIds int64[], _points uint8[]) returns()
func (_RequestContract *RequestContractSession) AddLacks(_contributionCodes []uint16, _controlCategoryIds []int64, _pointGroupCodes []uint16, _controlPointIds []int64, _lackIds []int64, _points []uint8) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLacks(&_RequestContract.TransactOpts, _contributionCodes, _controlCategoryIds, _pointGroupCodes, _controlPointIds, _lackIds, _points)
}

// AddLacks is a paid mutator transaction binding the contract method 0x38d1cc6a.
//
// Solidity: function addLacks(_contributionCodes uint16[], _controlCategoryIds int64[], _pointGroupCodes uint16[], _controlPointIds int64[], _lackIds int64[], _points uint8[]) returns()
func (_RequestContract *RequestContractTransactorSession) AddLacks(_contributionCodes []uint16, _controlCategoryIds []int64, _pointGroupCodes []uint16, _controlPointIds []int64, _lackIds []int64, _points []uint8) (*types.Transaction, error) {
	return _RequestContract.Contract.AddLacks(&_RequestContract.TransactOpts, _contributionCodes, _controlCategoryIds, _pointGroupCodes, _controlPointIds, _lackIds, _points)
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

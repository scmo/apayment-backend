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
const RequestContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_inspectorAddress\",\"type\":\"address\"}],\"name\":\"setInspectorId\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gve1110\",\"type\":\"uint16\"},{\"name\":\"_gve1150\",\"type\":\"uint16\"},{\"name\":\"_gve1128\",\"type\":\"uint16\"},{\"name\":\"_gve1141\",\"type\":\"uint16\"},{\"name\":\"_gve1142\",\"type\":\"uint16\"},{\"name\":\"_gve1124\",\"type\":\"uint16\"},{\"name\":\"_gve1129\",\"type\":\"uint16\"},{\"name\":\"_gve1143\",\"type\":\"uint16\"},{\"name\":\"_gve1144\",\"type\":\"uint16\"}],\"name\":\"setGVE\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"string\"},{\"name\":\"pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"controlPointId\",\"type\":\"string\"},{\"name\":\"lackId\",\"type\":\"int64\"},{\"name\":\"points\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_controlCategoryIds\",\"type\":\"int64[]\"},{\"name\":\"_pointGroupCodes\",\"type\":\"uint16[]\"},{\"name\":\"_controlPointIds\",\"type\":\"int64[]\"},{\"name\":\"_lackIds\",\"type\":\"int64[]\"},{\"name\":\"_points\",\"type\":\"uint8[]\"}],\"name\":\"addLacks\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modified\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionCode\",\"type\":\"uint16\"},{\"name\":\"_controlCategoryId\",\"type\":\"string\"},{\"name\":\"_pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"_controlPointId\",\"type\":\"string\"},{\"name\":\"_lackId\",\"type\":\"int64\"},{\"name\":\"_points\",\"type\":\"uint8\"}],\"name\":\"addLack\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numLacks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspectorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"calculateBTS\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pointGroups\",\"outputs\":[{\"name\":\"gve\",\"type\":\"uint16\"},{\"name\":\"points\",\"type\":\"uint16\"},{\"name\":\"btsPoints\",\"type\":\"uint16\"},{\"name\":\"btsTotal\",\"type\":\"uint256\"},{\"name\":\"btsDeduction\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"_points\",\"type\":\"uint16\"}],\"name\":\"updateBtsPoint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setCreated\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newLacks\",\"outputs\":[{\"name\":\"contributionCode\",\"type\":\"uint16\"},{\"name\":\"controlCategoryId\",\"type\":\"int64\"},{\"name\":\"pointGroupCode\",\"type\":\"uint16\"},{\"name\":\"controlPointId\",\"type\":\"int64\"},{\"name\":\"lackId\",\"type\":\"int64\"},{\"name\":\"points\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributionCodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"remark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setModified\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFirstPaymentAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_contributionCodes\",\"type\":\"uint16[]\"},{\"name\":\"_remark\",\"type\":\"string\"},{\"name\":\"gves\",\"type\":\"uint16[9]\"}],\"payable\":false,\"type\":\"constructor\"}]"

// RequestContractBin is the compiled bytecode used for deploying new contracts.
const RequestContractBin = `60606040526101206040519081016040528061045661ffff16815260200161047e61ffff16815260200161046861ffff16815260200161047561ffff16815260200161047661ffff16815260200161046461ffff16815260200161046961ffff16815260200161047761ffff16815260200161047861ffff16815250600a9060096200008d929190620005e6565b5034156200009757fe5b6040516200200a3803806200200a833981016040528080518201919060200180518201919060200190505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b730a3c03f7de1cefc9284358503225a0f269a4d868600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260049080519060200190620001729291906200068a565b5081600590805190602001906200018b9291906200073b565b5062000267816000600981101515620001a057fe5b6020020151826001600981101515620001b557fe5b6020020151836002600981101515620001ca57fe5b6020020151846003600981101515620001df57fe5b6020020151856004600981101515620001f457fe5b60200201518660056009811015156200020957fe5b60200201518760066009811015156200021e57fe5b60200201518860076009811015156200023357fe5b60200201518960086009811015156200024857fe5b6020020151620002716401000000000262000bc3176401000000009004565b5b5050506200081e565b6000600b600061045661ffff1681526020019081526020016000209050898160000160006101000a81548161ffff021916908361ffff160217905550600b600061047e61ffff1681526020019081526020016000209050888160000160006101000a81548161ffff021916908361ffff160217905550600b600061046861ffff1681526020019081526020016000209050878160000160006101000a81548161ffff021916908361ffff160217905550600b600061047561ffff1681526020019081526020016000209050868160000160006101000a81548161ffff021916908361ffff160217905550600b600061047661ffff1681526020019081526020016000209050858160000160006101000a81548161ffff021916908361ffff160217905550600b600061046461ffff1681526020019081526020016000209050848160000160006101000a81548161ffff021916908361ffff160217905550600b600061046961ffff1681526020019081526020016000209050838160000160006101000a81548161ffff021916908361ffff160217905550600b600061047761ffff1681526020019081526020016000209050828160000160006101000a81548161ffff021916908361ffff160217905550600b600061047861ffff1681526020019081526020016000209050818160000160006101000a81548161ffff021916908361ffff1602179055506200049b620004a8640100000000026200131e176401000000009004565b5b50505050505050505050565b60006000600091505b60098261ffff161015620005e157600b6000600a8461ffff16600981101515620004d757fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff168152602001908152602001600020905060001515610476600a8461ffff166009811015156200052557fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff161480620005845750610478600a8461ffff166009811015156200056357fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16145b15151415620005d257606e8160000160049054906101000a900461ffff1661ffff161015620005d1576123288160000160009054906101000a900461ffff1661ffff160281600101819055505b5b5b8180600101925050620004b1565b5b5050565b826009600f01601090048101928215620006775791602002820160005b838211156200064557835183826101000a81548161ffff021916908361ffff160217905550926020019260020160208160010104928301926001030262000603565b8015620006755782816101000a81549061ffff021916905560020160208160010104928301926001030262000645565b505b509050620006869190620007c2565b5090565b82805482825590600052602060002090600f01601090048101928215620007285791602002820160005b83821115620006f657835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302620006b4565b8015620007265782816101000a81549061ffff0219169055600201602081600101049283019260010302620006f6565b505b509050620007379190620007c2565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200077e57805160ff1916838001178555620007af565b82800160010185558215620007af579182015b82811115620007ae57825182559160200191906001019062000791565b5b509050620007be9190620007f6565b5090565b620007f391905b80821115620007ef57600081816101000a81549061ffff021916905550600101620007c9565b5090565b90565b6200081b91905b8082111562000817576000816000905550600101620007fd565b5090565b90565b6117dc806200082e6000396000f30060606040523615610105576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063021c7bd7146101075780630284bd8f1461013d5780630478e042146101c9578063325a19f11461034057806338d1cc6a1461036657806341c0e1b5146104fd57806345ecd3d71461050f57806348fa11561461053557806355bcbf5714610604578063696f8c571461062a5780637145644c1461067c5780637184f7251461068e5780638a001126146106fa57806390e102501461072b57806395be78691461073d5780639b5d25d9146107bc578063a47d1f5e146107f8578063af7fdd7614610891578063e0f8c670146108a3575bfe5b341561010f57fe5b61013b600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506108c9565b005b341561014557fe5b6101c7600480803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff1690602001909190803561ffff16906020019091905050610bc3565b005b34156101d157fe5b6101e76004808035906020019091905050610de4565b604051808761ffff1661ffff168152602001806020018661ffff1661ffff168152602001806020018560070b60070b81526020018460ff1660ff1681526020018381038352888181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102a85780601f1061027d576101008083540402835291602001916102a8565b820191906000526020600020905b81548152906001019060200180831161028b57829003601f168201915b505083810382528681815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561032b5780601f106103005761010080835404028352916020019161032b565b820191906000526020600020905b81548152906001019060200180831161030e57829003601f168201915b50509850505050505050505060405180910390f35b341561034857fe5b610350610e54565b6040518082815260200191505060405180910390f35b341561036e57fe5b6104fb6004808035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091905050610e5a565b005b341561050557fe5b61050d611093565b005b341561051757fe5b61051f611127565b6040518082815260200191505060405180910390f35b341561053d57fe5b610602600480803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803561ffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803560070b90602001909190803560ff1690602001909190505061112d565b005b341561060c57fe5b6106146112f2565b6040518082815260200191505060405180910390f35b341561063257fe5b61063a6112f8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561068457fe5b61068c61131e565b005b341561069657fe5b6106b0600480803561ffff16906020019091905050611454565b604051808661ffff1661ffff1681526020018561ffff1661ffff1681526020018461ffff1661ffff1681526020018381526020018281526020019550505050505060405180910390f35b341561070257fe5b610729600480803561ffff1690602001909190803561ffff169060200190919050506114b4565b005b341561073357fe5b61073b61150c565b005b341561074557fe5b61075b6004808035906020019091905050611516565b604051808761ffff1661ffff1681526020018660070b60070b81526020018561ffff1661ffff1681526020018460070b60070b81526020018360070b60070b81526020018260ff1660ff168152602001965050505050505060405180910390f35b34156107c457fe5b6107da60048080359060200190919050506115a2565b604051808261ffff1661ffff16815260200191505060405180910390f35b341561080057fe5b6108086115da565b6040518080602001828103825283818151815260200191508051906020019080838360008314610857575b80518252602083111561085757602082019150602081019050602083039250610833565b505050905090810190601f1680156108835780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561089957fe5b6108a1611678565b005b34156108ab57fe5b6108b3611682565b6040518082815260200191505060405180910390f35b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d7806c336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561098b57fe5b6102c65a03f1151561099957fe5b5050506040518051905080610a845750600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636bb164c9336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515610a6b57fe5b6102c65a03f11515610a7957fe5b505050604051805190505b1515610a905760006000fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632cdad41c826000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515610b5257fe5b6102c65a03f11515610b6057fe5b505050604051805190501515610b765760006000fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610bbf611678565b5b50565b6000600b600061045661ffff1681526020019081526020016000209050898160000160006101000a81548161ffff021916908361ffff160217905550600b600061047e61ffff1681526020019081526020016000209050888160000160006101000a81548161ffff021916908361ffff160217905550600b600061046861ffff1681526020019081526020016000209050878160000160006101000a81548161ffff021916908361ffff160217905550600b600061047561ffff1681526020019081526020016000209050868160000160006101000a81548161ffff021916908361ffff160217905550600b600061047661ffff1681526020019081526020016000209050858160000160006101000a81548161ffff021916908361ffff160217905550600b600061046461ffff1681526020019081526020016000209050848160000160006101000a81548161ffff021916908361ffff160217905550600b600061046961ffff1681526020019081526020016000209050838160000160006101000a81548161ffff021916908361ffff160217905550600b600061047761ffff1681526020019081526020016000209050828160000160006101000a81548161ffff021916908361ffff160217905550600b600061047861ffff1681526020019081526020016000209050818160000160006101000a81548161ffff021916908361ffff160217905550610dd761131e565b5b50505050505050505050565b60086020528060005260406000206000915090508060000160009054906101000a900461ffff169080600101908060020160009054906101000a900461ffff169080600301908060040160009054906101000a900460070b908060040160089054906101000a900460ff16905086565b60015481565b60006000600091505b87518261ffff161015611088576007600081548092919060010191905055905060c060405190810160405280898461ffff16815181101515610ea157fe5b9060200190602002015161ffff168152602001888461ffff16815181101515610ec657fe5b9060200190602002015160070b8152602001878461ffff16815181101515610eea57fe5b9060200190602002015161ffff168152602001868461ffff16815181101515610f0f57fe5b9060200190602002015160070b8152602001858461ffff16815181101515610f3357fe5b9060200190602002015160070b8152602001848461ffff16815181101515610f5757fe5b9060200190602002015160ff168152506009600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff160217905550604082015181600001600a6101000a81548161ffff021916908361ffff160217905550606082015181600001600c6101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060808201518160000160146101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060a082015181600001601c6101000a81548160ff021916908360ff1602179055509050505b8180600101925050610e63565b5b5050505050505050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561112457600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60025481565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561118c5760006000fd5b6007600081548092919060010191905055905060c0604051908101604052808861ffff1681526020018781526020018661ffff1681526020018581526020018460070b81526020018360ff168152506008600083815260200190815260200160002060008201518160000160006101000a81548161ffff021916908361ffff160217905550602082015181600101908051906020019061122d92919061170b565b5060408201518160020160006101000a81548161ffff021916908361ffff160217905550606082015181600301908051906020019061126d92919061170b565b5060808201518160040160006101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060a08201518160040160086101000a81548160ff021916908360ff1602179055509050506112cc611678565b6115288761ffff1614156112e8576112e7858360ff166114b4565b5b5b50505050505050565b60075481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006000600091505b60098261ffff16101561144f57600b6000600a8461ffff1660098110151561134b57fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff168152602001908152602001600020905060001515610476600a8461ffff1660098110151561139857fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1614806113f55750610478600a8461ffff166009811015156113d457fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff16145b1515141561144157606e8160000160049054906101000a900461ffff1661ffff161015611440576123288160000160009054906101000a900461ffff1661ffff160281600101819055505b5b5b8180600101925050611327565b5b5050565b600b6020528060005260406000206000915090508060000160009054906101000a900461ffff16908060000160029054906101000a900461ffff16908060000160049054906101000a900461ffff16908060010154908060020154905085565b6000600b60008461ffff1661ffff1681526020019081526020016000209050818160000160049054906101000a900461ffff16018160000160046101000a81548161ffff021916908361ffff1602179055505b505050565b426001819055505b565b60096020528060005260406000206000915090508060000160009054906101000a900461ffff16908060000160029054906101000a900460070b9080600001600a9054906101000a900461ffff169080600001600c9054906101000a900460070b908060000160149054906101000a900460070b9080600001601c9054906101000a900460ff16905086565b6004818154811015156115b157fe5b90600052602060002090601091828204019190066002025b915054906101000a900461ffff1681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116705780601f1061164557610100808354040283529160200191611670565b820191906000526020600020905b81548152906001019060200180831161165357829003601f168201915b505050505081565b426002819055505b565b600060006000600060009250600091505b60098261ffff16101561170157600b6000600a8461ffff166009811015156116b757fe5b601091828204019190066002025b9054906101000a900461ffff1661ffff1661ffff16815260200190815260200160002090508060010154830192505b8180600101925050611693565b8293505b50505090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061174c57805160ff191683800117855561177a565b8280016001018555821561177a579182015b8281111561177957825182559160200191906001019061175e565b5b509050611787919061178b565b5090565b6117ad91905b808211156117a9576000816000905550600101611791565b5090565b905600a165627a7a72305820ad3becb62a691ad1e78f10040de7c7c685a31826478844d12953c2807b4f8b960029`

// DeployRequestContract deploys a new Ethereum contract, binding an instance of RequestContract to it.
func DeployRequestContract(auth *bind.TransactOpts, backend bind.ContractBackend, _contributionCodes []uint16, _remark string, gves [9]uint16) (common.Address, *types.Transaction, *RequestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestContractBin), backend, _contributionCodes, _remark, gves)
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

// NewLacks is a free data retrieval call binding the contract method 0x95be7869.
//
// Solidity: function newLacks( uint256) constant returns(contributionCode uint16, controlCategoryId int64, pointGroupCode uint16, controlPointId int64, lackId int64, points uint8)
func (_RequestContract *RequestContractCaller) NewLacks(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RequestContract.contract.Call(opts, out, "newLacks", arg0)
	return *ret, err
}

// NewLacks is a free data retrieval call binding the contract method 0x95be7869.
//
// Solidity: function newLacks( uint256) constant returns(contributionCode uint16, controlCategoryId int64, pointGroupCode uint16, controlPointId int64, lackId int64, points uint8)
func (_RequestContract *RequestContractSession) NewLacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId int64
	PointGroupCode    uint16
	ControlPointId    int64
	LackId            int64
	Points            uint8
}, error) {
	return _RequestContract.Contract.NewLacks(&_RequestContract.CallOpts, arg0)
}

// NewLacks is a free data retrieval call binding the contract method 0x95be7869.
//
// Solidity: function newLacks( uint256) constant returns(contributionCode uint16, controlCategoryId int64, pointGroupCode uint16, controlPointId int64, lackId int64, points uint8)
func (_RequestContract *RequestContractCallerSession) NewLacks(arg0 *big.Int) (struct {
	ContributionCode  uint16
	ControlCategoryId int64
	PointGroupCode    uint16
	ControlPointId    int64
	LackId            int64
	Points            uint8
}, error) {
	return _RequestContract.Contract.NewLacks(&_RequestContract.CallOpts, arg0)
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
// Solidity: function pointGroups( uint16) constant returns(gve uint16, points uint16, btsPoints uint16, btsTotal uint256, btsDeduction uint256)
func (_RequestContract *RequestContractCaller) PointGroups(opts *bind.CallOpts, arg0 uint16) (struct {
	Gve          uint16
	Points       uint16
	BtsPoints    uint16
	BtsTotal     *big.Int
	BtsDeduction *big.Int
}, error) {
	ret := new(struct {
		Gve          uint16
		Points       uint16
		BtsPoints    uint16
		BtsTotal     *big.Int
		BtsDeduction *big.Int
	})
	out := ret
	err := _RequestContract.contract.Call(opts, out, "pointGroups", arg0)
	return *ret, err
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint16, points uint16, btsPoints uint16, btsTotal uint256, btsDeduction uint256)
func (_RequestContract *RequestContractSession) PointGroups(arg0 uint16) (struct {
	Gve          uint16
	Points       uint16
	BtsPoints    uint16
	BtsTotal     *big.Int
	BtsDeduction *big.Int
}, error) {
	return _RequestContract.Contract.PointGroups(&_RequestContract.CallOpts, arg0)
}

// PointGroups is a free data retrieval call binding the contract method 0x7184f725.
//
// Solidity: function pointGroups( uint16) constant returns(gve uint16, points uint16, btsPoints uint16, btsTotal uint256, btsDeduction uint256)
func (_RequestContract *RequestContractCallerSession) PointGroups(arg0 uint16) (struct {
	Gve          uint16
	Points       uint16
	BtsPoints    uint16
	BtsTotal     *big.Int
	BtsDeduction *big.Int
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

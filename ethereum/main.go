package ethereum

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{ "address": "9d1b510c234184566671e944fd6b1913766b3336", "crypto": { "cipher": "aes-128-ctr", "ciphertext": "a26f2810194c9f199d9d7184912db82dd1406ed0014791d9e6eac537323142be", "cipherparams": { "iv": "156e047a6ebaa4b6f8bfcd73b18d57c5" }, "kdf": "scrypt", "kdfparams": { "dklen": 32, "n": 262144, "p": 1, "r": 8, "salt": "805ac0b1e2b3be77d39bf9c36e5e09c39edb0b4abb8cd5f7cc0388a196c5fbc1" }, "mac": "1ef8da158d550f47d27b531d56146722e891ceb36bae81b1c3b332fe53a1e8e3" }, "id": "9b9fde22-bcfc-4de0-b7f8-49a48336972c", "version": 3 }`

type EthereumController struct {
	Auth   *bind.TransactOpts
	Client *ethclient.Client
}

var ethereumController EthereumController

func Init() {
	pathToIPCEndpoint := beego.AppConfig.String("pathToIPCEndpoint")
	if beego.BConfig.RunMode == "dev" {
		pathToIPCEndpoint = pathToIPCEndpoint + "testnet/"
	}
	pathToIPCEndpoint = pathToIPCEndpoint + "geth.ipc"

	// Create an IPC based RPC connection to a remote node
	client, err := ethclient.Dial(pathToIPCEndpoint)

	if err != nil {
		beego.Critical("Failed to connect to the Ethereum client: ", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(beego.AppConfig.String("systemAccount")), beego.AppConfig.String("systemAccountPassword"))
	if err != nil {
		beego.Critical("Failed to create authorized transactor: ", err)
	}

	ethereumController = EthereumController{Auth:auth, Client:client}
}

func GetEthereumController() EthereumController{
	return ethereumController
}


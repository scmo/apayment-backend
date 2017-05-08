package ethereum

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"io/ioutil"
	"bytes"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"context"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum"
)

type EthereumController struct {
	Auth     *bind.TransactOpts
	Client   *ethclient.Client
	Keystore *keystore.KeyStore
}

var ethereumController EthereumController

func Init() {
	pathToEthereum := beego.AppConfig.String("pathToEthereum")
	if beego.BConfig.RunMode == "dev" {
		pathToEthereum = pathToEthereum + "testnet/"
	}
	pathToIPCEndpoint := pathToEthereum + "geth.ipc"

	// Create an IPC based RPC connection to a remote node
	client, err := ethclient.Dial(pathToIPCEndpoint)

	if err != nil {
		beego.Critical("Failed to connect to the Ethereum client: ", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(beego.AppConfig.String("systemAccount")), beego.AppConfig.String("systemAccountPassword"))
	if err != nil {
		beego.Critical("Failed to create authorized transactor: ", err)
	}

	// Keystore to administrate accounts
	ks := keystore.NewKeyStore(pathToEthereum + "keystore", keystore.LightScryptN, keystore.LightScryptP)

	ethereumController = EthereumController{Auth:auth, Client:client, Keystore: ks}

	// Assuption: SyncProgress returns nil if it is not syncing
	//sy, err := client.SyncProgress(ctx)
	//if err != nil {
	//	beego.Critical("Failed to get SyncProgress: ", err)
	//}
	//_ = sy


}

func SendEther(from string, to string, amountEther int64) {
	ctx := context.Background()
	fromAccount, err := ethereumController.Keystore.Find(accounts.Account{Address: common.HexToAddress(from)})
	nonce, err := ethereumController.Client.PendingNonceAt(ctx, fromAccount.Address)
	if err != nil {
		beego.Critical("Failed to get nounce: ", err)
	}
	amount := new(big.Int).Mul(big.NewInt(amountEther), big.NewInt(params.Ether))
	//estimateGas, err := ethereumController.Client.EstimateGas(ctx, ethereum.CallMsg{From: fromAccount.Address, To: common.StringToAddress(to), Value:amount, Data: nil})
	estimateGas, err := ethereumController.Client.EstimateGas(ctx, ethereum.CallMsg{From: fromAccount.Address, Value:amount, Data: nil})
	if err != nil {
		beego.Critical("Failed to estimate gas: ", err)
	}
	tx := types.NewTransaction(nonce, common.StringToAddress(to), amount, estimateGas, big.NewInt(50000000000), nil)

	tx, err = ethereumController.Keystore.SignTxWithPassphrase(fromAccount, beego.AppConfig.String("systemAccountPassword"), tx, big.NewInt(3))
	if err != nil {
		beego.Critical("Failed to Sign Transaction: ", err)
	}
	err = ethereumController.Client.SendTransaction(ctx, tx)
	if err != nil {
		beego.Critical("Failed to Send Transaction: ", err)
	}

}

func GetAuth(address string) {
	account, err := ethereumController.Keystore.Find(accounts.Account{Address: common.HexToAddress(address)})
	if ( err != nil ) {
		beego.Error("Error find: ", err)
	}
	dat, err := ioutil.ReadFile(account.URL.Path)
	auth, err := bind.NewTransactor(bytes.NewReader(dat), beego.AppConfig.String("userAccountPassword"))
	if err != nil {
		beego.Critical("Failed  to create authorized transactor: ", err)
	}
	ethereumController.Auth = auth
}

func GetEthereumController() EthereumController {
	return ethereumController
}


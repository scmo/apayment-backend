package services

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/scmo/apayment-backend/smart-contracts/apayment-token"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/ethereum"
)

func Transfer(sender common.Address, dst common.Address, amount *big.Int) (error) {
	ethereumController := ethereum.GetEthereumController()
	token, err := apaymenttoken.NewAPaymentTokenContract(common.HexToAddress(beego.AppConfig.String("apaymentTokenContract")), ethereumController.Client)
	if (err != nil ) {
		beego.Critical("Failed to instantiate a APaymentTokenContract contract:", err)
		return err

	}
	tx, err := token.Transfer(ethereum.GetAuth(sender.String()), dst, amount)
	if err != nil {
		beego.Error("Failed to deploy new token contract: ", err)
		return err
	}
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	return err
}

func GetBalanceOf(address common.Address) (*big.Int, error) {
	ethereumController := ethereum.GetEthereumController()
	token, err := apaymenttoken.NewAPaymentTokenContract(common.HexToAddress(beego.AppConfig.String("apaymentTokenContract")), ethereumController.Client)
	if (err != nil ) {
		beego.Critical("Failed to instantiate a APaymentTokenContract contract:", err)
		return nil, err

	}
	balance, err := token.BalanceOf(nil, address)
	if ( err != nil ) {
		beego.Error("Error while getting aPayment Token Balance")
	}
	return balance, err
}
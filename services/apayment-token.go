package services

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/scmo/apayment-backend/smart-contracts/apayment-token"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/ethereum"
	"github.com/scmo/apayment-backend/models"
	"net/http"
	"encoding/json"
	"encoding/hex"
	"strconv"
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

func GetTransactions() ([]*models.APaymentTokenTransaction, error) {

	transactions := make([]*models.APaymentTokenTransaction, 0)
	etherScanResult, err := fetchTransaction()
	if (err != nil ) {
		beego.Error("Error while fetch Transcations from Etherscan", err)
		return transactions, err
	}

	for _, tx := range etherScanResult.Result {
		isError, err := strconv.ParseBool(tx.IsError)
		if ( err != nil) {
			beego.Debug("Error while converting isError (string) to int. ", err)
			return transactions, err
		}
		if (isError) {
			continue
		}
		dst, amount, err := parseInput(tx.Input)
		if (err != nil) {
			beego.Error("Error while parsing transaction input. ", err)
			return transactions, err
		}

		from := &models.User{Firstname: "aPayment", Lastname:"System"}
		if (tx.From != beego.AppConfig.String("systemAccountAddress")) {
			from, err = GetUserByAddress(tx.From)
			if (err != nil) {
				beego.Error("Error while getting User by Address. ", err)
				continue
			}
		}

		to, err := GetUserByAddress(dst)
		if (err != nil) {
			beego.Error("Error while getting User by Address. ", err)
			continue
		}
		transactions = append(transactions, &models.APaymentTokenTransaction{From: from, To: to, Amount: &amount, Timestamp:tx.Timestamp})
	}
	return transactions, err
}

func fetchTransaction() (models.EtherScanTransactionResult, error) {
	var result models.EtherScanTransactionResult

	url := beego.AppConfig.String("etherScanApiUrl")
	url = url + "?module=account&action=txlist&address=" + beego.AppConfig.String("apaymentTokenContract") + "&apikey=" + beego.AppConfig.String("etherScanApiToken")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if ( err != nil ) {
		beego.Error("Error while creating request to fetch transactions. ", err)
		return result, err
	}

	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		beego.Error("Error while sending request to fetch transactions. ", err)
		return result, err
	}
	defer resp.Body.Close()


	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		beego.Error("Error while reading result from Etherscan. ", err)
		return result, err
	}
	return result, err
}

func parseInput(input string) (string, big.Int, error) {
	//methodID := input[0:10]
	paramsInput := input[10:]
	dstInput := paramsInput[0:64]
	amountInput := paramsInput[64:]

	//Get destination
	dst := "0x" + dstInput[len(dstInput) - 40:]
	beego.Debug(dst)
	// Get Amount
	amountBytes, err := hex.DecodeString(amountInput[len(amountInput) - 16:])
	if ( err != nil ) {
		beego.Error("Failed to decode hex to bytes. ", err)
		return dst, big.Int{}, err
	}
	amount := big.Int{}
	amount.SetBytes(amountBytes)
	return dst, amount, err
}
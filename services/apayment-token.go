package services

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/ethereum"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services/tvd"
	"github.com/scmo/apayment-backend/smart-contracts/apayment-token"
	"net/http"
	"strconv"
)

func Transfer(aPaymentTokenTransfer *models.APaymentTokenTransfer, requestAddress string) error {
	// check if sender has enough fund
	balance, err := GetBalanceOf(common.HexToAddress(aPaymentTokenTransfer.From))
	if err != nil {
		beego.Error("Error while reading balance. ", err)
	}
	if aPaymentTokenTransfer.Amount.Cmp(balance) == 1 {
		// +1 if x >  y
		beego.Info("aPaymetToken balance is too small.")
		return errors.New("aPaymetToken balance is too small.")
	}

	ethereumController := ethereum.GetEthereumController()
	token, err := apaymenttoken.NewAPaymentTokenContract(common.HexToAddress(beego.AppConfig.String("apaymentTokenContract")), ethereumController.Client)
	if err != nil {
		beego.Critical("Failed to instantiate a APaymentTokenContract contract:", err)
		return err

	}
	if len(requestAddress) == 0 {
		tx, err := token.TransferWithMessage(ethereum.GetAuth(aPaymentTokenTransfer.From), common.HexToAddress(aPaymentTokenTransfer.To), aPaymentTokenTransfer.Amount, []byte(aPaymentTokenTransfer.Message))
		if err != nil {
			beego.Error("Failed to send new transaction: ", err)
			return err
		}
		beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	} else {
		tx, err := token.TransferWithMessageAndRequestAddress(ethereum.GetAuth(aPaymentTokenTransfer.From), common.HexToAddress(aPaymentTokenTransfer.To), aPaymentTokenTransfer.Amount, common.HexToAddress(requestAddress), []byte(aPaymentTokenTransfer.Message))
		if err != nil {
			beego.Error("Failed to send new transaction: ", err)
			return err
		}
		beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	}
	return err
}

func GetBalanceOf(address common.Address) (*big.Int, error) {
	ethereumController := ethereum.GetEthereumController()
	token, err := apaymenttoken.NewAPaymentTokenContract(common.HexToAddress(beego.AppConfig.String("apaymentTokenContract")), ethereumController.Client)
	if err != nil {
		beego.Critical("Failed to instantiate a APaymentTokenContract contract:", err)
		return nil, err

	}
	balance, err := token.BalanceOf(nil, address)
	if err != nil {
		beego.Error("Error while getting aPayment Token Balance")
	}
	return balance, err
}
func GetTransactionForRequest(requestAddress string) []*models.APaymentTokenTransaction {
	transactions, err := GetTransactions()
	requestTransaction := make([]*models.APaymentTokenTransaction, 0)
	if err != nil {
		beego.Error("Error while fetch transactions. ", err)
	}
	for _, tx := range transactions {
		if tx.Request != nil && tx.Request.Address == requestAddress {
			requestTransaction = append(requestTransaction, tx)
		}
	}
	return requestTransaction
}

func GetTransactions() ([]*models.APaymentTokenTransaction, error) {
	transactions := make([]*models.APaymentTokenTransaction, 0)
	etherScanResult, err := fetchTransaction()
	if err != nil {
		beego.Error("Error while fetch Transcations from Etherscan", err)
		return transactions, err
	}
	for i, tx := range etherScanResult.Result {
		// skip contract creation transaction
		if i == 0 {
			continue
		}
		isError, err := strconv.ParseBool(tx.IsError)
		if err != nil {
			beego.Error("Error while converting isError (string) to int. ", err)
			return transactions, err
		}
		if isError {
			continue
		}
		dst, amount, msg, requestAddress, err := parseInput(tx.Input)
		if err != nil {
			beego.Error("Error while parsing transaction input. ", err)
			return transactions, err
		}
		//from := &models.User{Firstname: "aPayment", Lastname:"System"}
		from := &models.User{AnimalHusbandryDetailResult: &tvd.GetAnimalHusbandryDetailResult{PostData: &tvd.HusbandryResult{Name: "aPayment System"}}}
		if tx.From != beego.AppConfig.String("systemAccountAddress") {
			from, err = GetUserByAddress(tx.From)
			if err != nil {
				beego.Error("Error while getting User by Address. ", err)
				continue
			}
		}
		to, err := GetUserByAddress(dst)
		if err != nil {
			beego.Error("Error while getting User by Address. ", err)
			continue
		}
		if requestAddress == "" {
			transactions = append(transactions, &models.APaymentTokenTransaction{From: from, To: to, Amount: &amount, Timestamp: tx.Timestamp, Message: msg})
		} else {
			requestId := GetRequestIdByAddress(requestAddress)
			request := GetRequestById(requestId, false)
			transactions = append(transactions, &models.APaymentTokenTransaction{From: from, To: to, Amount: &amount, Timestamp: tx.Timestamp, Message: msg, Request: request})
		}
	}
	return transactions, err
}

func fetchTransaction() (models.EtherScanTransactionResult, error) {
	var result models.EtherScanTransactionResult

	url := beego.AppConfig.String("etherScanApiUrl")
	url = url + "?module=account&action=txlist&address=" + beego.AppConfig.String("apaymentTokenContract") + "&apikey=" + beego.AppConfig.String("etherScanApiToken")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
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

func parseInput(input string) (string, big.Int, string, string, error) {
	methodID := input[0:10]
	paramsInput := input[10:]
	dstInput := paramsInput[0:64]
	amountInput := paramsInput[64:128]

	requestAddress := ""
	msgStart := 128
	if methodID == "0xd3088b52" {
		requestAddressInput := paramsInput[128:192]
		requestAddress = "0x" + requestAddressInput[len(requestAddressInput) - 40:]
		msgStart = msgStart + 64
	}
	msgInput := paramsInput[msgStart:]
	//requestAddressInput := paramsInput[192:]
	//Get destination
	dst := "0x" + dstInput[len(dstInput) - 40:]
	// Get Amount
	amountBytes, err := hex.DecodeString(amountInput[len(amountInput) - 16:])
	if err != nil {
		beego.Error("Failed to decode hex to bytes. ", err)
		return dst, big.Int{}, "", requestAddress, err
	}
	amount := big.Int{}
	amount.SetBytes(amountBytes)

	// Get Message
	//msgLengthHex := msgInput[0:64]
	//msgLengthBytes, err := hex.DecodeString(msgLengthHex[len(msgLengthHex) - 16:])
	//if ( err != nil ) {
	//	beego.Error("Failed to decode hex to bytes. ", err)
	//	return dst, big.Int{}, err
	//}
	//msgLength := big.Int{}
	//msgLength.SetBytes(msgLengthBytes)

	msgInput = msgInput[128:]
	msgBytes, err := hex.DecodeString(msgInput)
	if err != nil {
		beego.Error("Failed to decode hex to string. ", err)
		return dst, big.Int{}, "", requestAddress, err
	}
	msg := string(msgBytes)

	//beego.Info(dst)
	//beego.Info(amount)
	//beego.Info(msg)
	//beego.Info(requestAddress)
	return dst, amount, msg, requestAddress, err
}

package models

import "math/big"

type APaymentTokenTransfer struct {
	From    string          `json:"from"`
	To      string          `json:"to"`
	Amount  *big.Int        `json:"amount"`
	Message string          `json:"message"`
}

type APaymentTokenTransaction struct {
	Timestamp string        `json:"timestamp"`
	From      *User         `json:"from"`
	To        *User         `json:"to"`
	Amount    *big.Int      `json:"amount"`
	Message   string        `json:"message"`
	Request   *Request      `json:"request"`
}

type EtherScanTransactionResult struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []*EtherScanTransaction `json:"result"`
}

type EtherScanTransaction struct {
	BlockNumber     string `json:"blockNumber"`
	Timestamp       string `json:"timeStamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	Input           string `json:"input"`
	ContractAddress string `json:"contractAddress"`
	IsError         string `json:"isError"`
}

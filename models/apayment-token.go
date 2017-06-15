package models

import "math/big"

type APaymentTokenTransfer struct {
	From   string                        `json:"from"`
	To     string                `json:"to"`
	Amount *big.Int                `json:"amount"`
}

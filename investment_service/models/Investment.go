package models

import "time"

type Investment struct {
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Qty           int    `json:"qty"`
	UpdatedAt     time.Time
	CreatedAt     time.Time
}

func NewInvestment(accNumber string, amount, qty int) Investment {
	return Investment{
		AccountNumber: accNumber,
		Amount:        amount,
		Qty:           qty,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
	}

}

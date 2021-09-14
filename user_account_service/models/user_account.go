package models

import "time"

type UserAccount struct {
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	UpdatedAt     time.Time
	CreatedAt     time.Time
}

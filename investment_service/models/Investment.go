package models

import "time"

type Investment struct {
	ProjectId     string `json:"project_id"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Qty           int64  `json:"qty"`
	UpdatedAt     time.Time
	CreatedAt     time.Time
}

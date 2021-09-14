package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type InvestmentStatus string

const (
	INVEST_REQUEST InvestmentStatus = "INVEST_REQUEST"
	INVEST_SUCCESS InvestmentStatus = "INVEST_SUCCESS"
	INVEST_FAIL    InvestmentStatus = "INVEST_FAIL"
)

type Investment struct {
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Qty           int    `json:"qty"`
	Status        string `json:"status"`
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

type InvestmentDtoBind struct {
	Amount        int    `json:"amount" binding:"required"`
	Qty           int    `json:"qty" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
}

func main() {

	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/investments", func(c *gin.Context) {
		req := &InvestmentDtoBind{}
		err := c.Bind(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.Status(http.StatusOK)
	})

	r.Run(":8081")
}

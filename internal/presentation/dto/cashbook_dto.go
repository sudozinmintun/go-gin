package dto

import "github.com/shopspring/decimal"

type CashbookRequest struct {
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"`
}

type CashbookResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name" `
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"`
}

package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Cashbook struct {
	ID          uint
	Name        string
	Description string
	Amount      decimal.Decimal
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

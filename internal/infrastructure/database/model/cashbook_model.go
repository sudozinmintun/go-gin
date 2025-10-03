package model

import (
	"example.com/accounting/internal/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CashbookModel struct {
	gorm.Model
	ID          uint            `gorm:"primaryKey"`
	Name        string          `gorm:"size:255;null"`
	Description string          `gorm:"type:text"`
	Amount      decimal.Decimal `gorm:"type:numeric(18, 2);not null"`
}

func (CashbookModel) TableName() string {
	return "cashbooks"
}

func ToCashbookDomain(model *CashbookModel) *domain.Cashbook {
	return &domain.Cashbook{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Amount:      model.Amount,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func ToCashbookModel(entity *domain.Cashbook) *CashbookModel {
	return &CashbookModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Amount:      entity.Amount,
	}
}

package database

import (
	"context"
	"errors"

	"example.com/accounting/internal/domain"
	"example.com/accounting/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type CashbookRepositoryImpl struct {
	DB *gorm.DB
}

func NewCashbookRepositoryImpl(db *gorm.DB) domain.CashbookRepository {
	return &CashbookRepositoryImpl{DB: db}
}

func (r CashbookRepositoryImpl) FindAll(ctx context.Context) ([]*domain.Cashbook, error) {
	var cashbookModels []model.CashbookModel
	if err := r.DB.WithContext(ctx).Find(&cashbookModels).Error; err != nil {
		return nil, err
	}

	cashbooks := make([]*domain.Cashbook, len(cashbookModels))
	for i, m := range cashbookModels {
		cashbooks[i] = model.ToCashbookDomain(&m)
	}
	return cashbooks, nil
}

func (r CashbookRepositoryImpl) Save(ctx context.Context, cashbook *domain.Cashbook) (*domain.Cashbook, error) {
	cashbookModel := model.ToCashbookModel(cashbook)

	if err := r.DB.WithContext(ctx).Create(cashbookModel).Error; err != nil {
		return nil, err
	}

	return model.ToCashbookDomain(cashbookModel), nil
}

func (r CashbookRepositoryImpl) FindByID(ctx context.Context, id uint) (*domain.Cashbook, error) {
	var cashbookModel model.CashbookModel
	if err := r.DB.WithContext(ctx).First(&cashbookModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return model.ToCashbookDomain(&cashbookModel), nil
}

func (r CashbookRepositoryImpl) Delete(ctx context.Context, id uint) error {
	result := r.DB.WithContext(ctx).Unscoped().Delete(&model.CashbookModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r CashbookRepositoryImpl) Update(ctx context.Context, cashbook *domain.Cashbook) (*domain.Cashbook, error) {
	result := r.DB.WithContext(ctx).
		Model(&model.CashbookModel{}).
		Where("id = ?", cashbook.ID).
		Updates(model.CashbookModel{
			Name:        cashbook.Name,
			Description: cashbook.Description,
			Amount:      cashbook.Amount,
		})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNotFound
	}

	return cashbook, nil
}

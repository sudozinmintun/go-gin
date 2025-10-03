package application

import (
	"context"
	"errors"

	"example.com/accounting/internal/domain"
	"github.com/shopspring/decimal"
)

var (
	ErrNegativeCashbookAmount = errors.New("cashbook amount must be at least 0")
)

type CashbookService struct {
	repo domain.CashbookRepository
}

func NewCashbookService(repo domain.CashbookRepository) *CashbookService {
	return &CashbookService{repo: repo}
}

func (s CashbookService) GetCashbooks(ctx context.Context) ([]*domain.Cashbook, error) {
	return s.repo.FindAll(ctx)
}

func (s CashbookService) CreateCashbook(ctx context.Context, name, description string, amount decimal.Decimal) (*domain.Cashbook, error) {

	//Example: (amount >= 0) if -10
	if amount.IsNegative() {
		return nil, ErrNegativeCashbookAmount
	}

	newCashbook := &domain.Cashbook{
		Name:        name,
		Description: description,
		Amount:      amount,
	}
	return s.repo.Save(ctx, newCashbook)
}

func (s CashbookService) GetCashbookById(ctx context.Context, id uint) (*domain.Cashbook, error) {
	return s.repo.FindByID(ctx, id)
}

func (s CashbookService) DeleteCashbook(ctx context.Context, id uint) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s CashbookService) UpdateCashbook(ctx context.Context, name, description string, amount decimal.Decimal, id uint) (*domain.Cashbook, error) {
	if amount.IsNegative() {
		return nil, ErrNegativeCashbookAmount
	}

	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	updateCashbook := &domain.Cashbook{
		ID:          id,
		Name:        name,
		Description: description,
		Amount:      amount,
	}
	return s.repo.Update(ctx, updateCashbook)
}

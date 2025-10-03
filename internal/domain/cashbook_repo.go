package domain

import (
	"context"
)

type CashbookRepository interface {
	FindAll(ctx context.Context) ([]*Cashbook, error)
	Save(ctx context.Context, cashbook *Cashbook) (*Cashbook, error)
	Update(ctx context.Context, cashbook *Cashbook) (*Cashbook, error)
	FindByID(ctx context.Context, id uint) (*Cashbook, error)
	Delete(ctx context.Context, id uint) error
}

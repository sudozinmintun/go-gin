package domain

import "context"

type CategoryRepository interface {
	Save(ctx context.Context, category *Category) (*Category, error)
	FindByID(ctx context.Context, id uint) (*Category, error)
	FindAll(ctx context.Context) ([]*Category, error)
	FindByName(ctx context.Context, name string) (*Category, error)
	Delete(ctx context.Context, id uint) error
}

package application

import (
	"context"
	"errors"

	"example.com/accounting/internal/domain"
)

var (
	ErrInvalidCategoryName   = errors.New("category name must be at least 3 characters")
	ErrCategoryAlreadyExists = errors.New("category with this name already exists")
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, name, description, unit string) (*domain.Category, error) {
	if len(name) < 3 {
		return nil, ErrInvalidCategoryName
	}

	existingCategory, err := s.repo.FindByName(ctx, name)

	if err != nil {
		return nil, err
	}

	if existingCategory != nil {
		return nil, ErrCategoryAlreadyExists
	}

	newCategory := &domain.Category{
		Name:        name,
		Description: description,
		Unit:        unit,
	}

	return s.repo.Save(ctx, newCategory)
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]*domain.Category, error) {
	return s.repo.FindAll(ctx)
}

func (s *CategoryService) GetCategoryByID(ctx context.Context, id uint) (*domain.Category, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id uint) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, name, description, unit string, id uint) (*domain.Category, error) {
	if len(name) < 3 {
		return nil, ErrInvalidCategoryName
	}

	existingCategory, err := s.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if existingCategory != nil {
		if existingCategory.ID != id {
			return nil, ErrCategoryAlreadyExists
		}
	}

	_, err = s.repo.FindByID(ctx, id)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	newCategory := &domain.Category{
		ID:          id,
		Name:        name,
		Description: description,
		Unit:        unit,
	}

	return s.repo.Update(ctx, newCategory)
}

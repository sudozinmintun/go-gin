package database

import (
	"context"
	"errors"

	"example.com/accounting/internal/domain"
	"example.com/accounting/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) domain.CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	catModel := model.ToCategoryModel(category)
	if err := r.DB.WithContext(ctx).Create(catModel).Error; err != nil {
		return nil, err
	}
	return model.ToCategoryDomain(catModel), nil
}

func (r *CategoryRepositoryImpl) FindByID(ctx context.Context, id uint) (*domain.Category, error) {
	var catModel model.CategoryModel

	if err := r.DB.WithContext(ctx).First(&catModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return model.ToCategoryDomain(&catModel), nil
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]*domain.Category, error) {
	var catModels []model.CategoryModel
	if err := r.DB.WithContext(ctx).Find(&catModels).Error; err != nil {
		return nil, err
	}

	categories := make([]*domain.Category, len(catModels))
	for i, m := range catModels {
		categories[i] = model.ToCategoryDomain(&m)
	}
	return categories, nil
}

func (r *CategoryRepositoryImpl) FindByName(ctx context.Context, name string) (*domain.Category, error) {
	var catModel model.CategoryModel
	err := r.DB.WithContext(ctx).Where("name = ?", name).First(&catModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return model.ToCategoryDomain(&catModel), nil
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, id uint) error {

	result := r.DB.WithContext(ctx).
		Unscoped().
		Delete(&model.CategoryModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

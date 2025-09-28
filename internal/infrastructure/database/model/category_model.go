package model

import (
	"example.com/accounting/internal/domain"
	"gorm.io/gorm"
)

type CategoryModel struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null;uniqueIndex"`
	Description string `gorm:"type:text"`
	Unit        string `gorm:"size:50"`
}

func (CategoryModel) TableName() string {
	return "categories"
}

func ToCategoryDomain(model *CategoryModel) *domain.Category {
	return &domain.Category{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Unit:        model.Unit,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func ToCategoryModel(entity *domain.Category) *CategoryModel {
	return &CategoryModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Unit:        entity.Unit,
	}
}

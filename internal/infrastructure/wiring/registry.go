package wiring

import (
	"example.com/accounting/internal/application"
	"example.com/accounting/internal/infrastructure/database"
	"example.com/accounting/internal/presentation/http"

	"gorm.io/gorm"
)

// Registry holds all initialized application components
type Registry struct {
	CategoryHandler *http.CategoryHandler
}

// NewRegistry initializes all dependencies and wires them together.
func NewRegistry(db *gorm.DB) *Registry {
	// 1. Infrastructure Layer: Repositories
	categoryRepo := database.NewCategoryRepositoryImpl(db)

	// 2. Application Layer: Services
	categoryService := application.NewCategoryService(categoryRepo)

	// 3. Presentation Layer: Handlers
	categoryHandler := http.NewCategoryHandler(categoryService)

	return &Registry{
		CategoryHandler: categoryHandler,
	}
}

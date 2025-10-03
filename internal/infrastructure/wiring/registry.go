package wiring

import (
	"example.com/accounting/internal/application"
	"example.com/accounting/internal/infrastructure/database"
	"example.com/accounting/internal/presentation/http"
	"example.com/accounting/internal/presentation/views"
	"gorm.io/gorm"
)

// Registry holds all initialized application components
type Registry struct {
	CategoryHandler *http.CategoryHandler
	CashbookHandler *http.CashbookHandler
	ViewHandler     *views.ViewHandler
}

// NewRegistry initializes all dependencies and wires them together.
func NewRegistry(db *gorm.DB) *Registry {
	// 1. Infrastructure Layer: Repositories
	categoryRepo := database.NewCategoryRepositoryImpl(db)
	cashbookRepo := database.NewCashbookRepositoryImpl(db)

	// 2. Application Layer: Services
	categoryService := application.NewCategoryService(categoryRepo)
	cashbookService := application.NewCashbookService(cashbookRepo)

	// 3. Presentation Layer: Handlers
	categoryHandler := http.NewCategoryHandler(categoryService)
	cashbookHandler := http.NewCashbookHandler(cashbookService)
	viewHandler := views.NewViewHandler()

	return &Registry{
		CategoryHandler: categoryHandler,
		CashbookHandler: cashbookHandler,
		ViewHandler:     viewHandler,
	}
}

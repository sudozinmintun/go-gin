package database

import "example.com/accounting/internal/infrastructure/database/model"

// AllModels lists all GORM structs that need to be migrated.
var AllModels = []interface{}{
	&model.CategoryModel{},
	&model.CashbookModel{},
}

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB sets up the database connection and runs auto-migrations.
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate all models
	err = db.AutoMigrate(AllModels...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

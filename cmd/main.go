package main

import (
	"log"

	"example.com/accounting/internal/infrastructure/database"
	"example.com/accounting/internal/infrastructure/wiring"
)

func main() {
	// --- Configuration ---
	// TODO: Move DSN to environment variables or config file
	dsn := "root:@tcp(127.0.0.1:3306)/go_accounting?charset=utf8mb4&parseTime=True&loc=Local"

	// 1. Initialize Database (Connect and AutoMigrate all models)
	db, err := database.InitDB(dsn)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// 2. Dependency Wiring/Registry - Initializes ALL components
	registry := wiring.NewRegistry(db)

	// 3. Router Setup - Registers ALL endpoints
	r := wiring.SetupRouter(registry)

	// 4. Start Server
	log.Println("Server running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

package internal

import (
	"expense-tracker/internal/database"
	"fmt"

	"github.com/joho/godotenv"
)

func Initialize() {
	// Load configuration
	err := godotenv.Load("../../internal/config/.env")

	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	// Check database connection
	if err := database.CheckConnection(); err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return
	}

	database.Connect()
}

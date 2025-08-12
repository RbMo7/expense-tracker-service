package main

import (
	"expense-tracker/internal/database"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("Welcome to the Expense Tracker Service!\n")

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

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	http.ListenAndServe(":"+os.Getenv("BASE_PORT"), nil)
	fmt.Printf("Server is running on port %s\n", os.Getenv("BASE_PORT"))
}

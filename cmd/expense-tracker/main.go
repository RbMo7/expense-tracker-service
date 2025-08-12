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

	database.Connect()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})



	fmt.Printf("Configuration loaded successfully!\n")
	fmt.Printf("DB_HOST: %s\n", os.Getenv("DB_HOST"))
	fmt.Printf("DB_PORT: %s\n", os.Getenv("DB_PORT"))
	fmt.Printf("DB_USER: %s\n", os.Getenv("DB_USER"))
	fmt.Printf("DB_PASSWORD: %s\n", os.Getenv("DB_PASSWORD"))
	fmt.Printf("DB_NAME: %s\n", os.Getenv("DB_NAME"))
	fmt.Printf("BASE_URL: %s\n", os.Getenv("BASE_URL"))

	http.ListenAndServe(":"+os.Getenv("BASE_PORT"), nil)
}

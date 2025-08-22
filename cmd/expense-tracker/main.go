package main

import (
	"expense-tracker/internal"
	"expense-tracker/internal/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Printf("Welcome to the Expense Tracker Service!\n")

	// Dotenv and database initialization
	internal.Initialize()

	// Route handling
	r := chi.NewRouter()

	// Register all routers
	handlers.RegisterRoutes(r)

	http.ListenAndServe(":"+os.Getenv("BASE_PORT"), r)
	fmt.Printf("Server is running on port %s\n", os.Getenv("BASE_PORT"))
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	// Load env variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Create connection string
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	DB = pool
	fmt.Println("Database connection established successfully!")
}

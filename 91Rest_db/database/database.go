package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func Connect() (*sqlx.DB, error) {
// 	dsn := os.Getenv("DATABASE_URL")
// 	if dsn == "" {
// 		return nil, fmt.Errorf("DATABASE_URL is not set")
// 	}

// 	db, err := sqlx.Connect("postgres", dsn)

// 	if err != nil {
// 		return nil, fmt.Errorf("database connection failed: %w", err)
// 	}

// 	return db, nil
// }

// using the gorm package to connect to the database
func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

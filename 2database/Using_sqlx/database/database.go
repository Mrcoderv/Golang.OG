package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// loadEnv reads KEY=VALUE pairs from a .env file into the process
// environment. Variables that are already set are left untouched.
func loadEnv(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("read %s: %w", path, err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return fmt.Errorf("invalid entry in %s: %q", path, line)
		}

		key = strings.TrimSpace(key)
		if key == "" {
			return fmt.Errorf("invalid empty key in %s", path)
		}
		value = strings.Trim(strings.TrimSpace(value), `"'`)

		if _, set := os.LookupEnv(key); !set {
			if err := os.Setenv(key, value); err != nil {
				return fmt.Errorf("set %s: %w", key, err)
			}
		}
	}
	return nil
}

func Connect() (*sqlx.DB, error) {
	if err := loadEnv(".env"); err != nil {
		return nil, fmt.Errorf("load environment: %w", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set (export it or put it in .env)")
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db, nil
}

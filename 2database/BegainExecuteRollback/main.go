package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func transferMoney(
	ctx context.Context,
	db *sql.DB,
	fromID, toID int,
	amount int,
) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Roll back if the transaction hasn't been committed.
	defer tx.Rollback()

	// Deduct money from the sender.
	result, err := tx.ExecContext(ctx,
		`UPDATE accounts
		 SET balance = balance - $1
		 WHERE id = $2 AND balance >= $1`,
		amount, fromID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("sender not found or insufficient balance")
	}

	Add money to the receiver.
	result, err = tx.ExecContext(ctx,
		`UPDATE accounts
		 SET balance = balance + $1
		 WHERE id = $2`,
		amount, toID,
	)
// error cases
// 	result, err = tx.ExecContext(ctx,
//     `UPDATE accounts
//      SET balance = balance + $1
//      WHERE id = $2 AND 1 = 0`,
//     amount, toID,
// )
	if err != nil {
		return err
	}

	rows, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("receiver not found")
	}

	// Commit both changes.
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	fmt.Println("Transfer successful")
	return nil
}

func main() {
	ctx := context.Background()

	db, err := sql.Open(
		"postgres",
		"host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = transferMoney(ctx, db, 1, 2, 1000)
	if err != nil {
		log.Println("Transfer failed:", err)
		return
	}

	log.Println("Transfer completed")
}

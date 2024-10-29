package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute DB queries and transaction
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transaction
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Start a new transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx) // Create a new Queries instance with the transaction
	err = fn(q)  // Execute the function with the Queries instance
	if err != nil {
		// Rollback the transaction if an error occurred
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	// Commit the transaction if no error occurred
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// Execute the transfer transaction
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create from account entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create to account entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// Lock accounts in a consistent order to prevent deadlock
		if arg.FromAccountID < arg.ToAccountID {
			// Lock and update from account first
			result.FromAccount, result.ToAccount, err = updateBalances(ctx, q, arg.FromAccountID, arg.ToAccountID, -arg.Amount, arg.Amount)
		} else {
			// Lock and update to account first
			result.ToAccount, result.FromAccount, err = updateBalances(ctx, q, arg.ToAccountID, arg.FromAccountID, arg.Amount, -arg.Amount)
		}

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// Helper function to update account balances in a consistent order
func updateBalances(
	ctx context.Context,
	q *Queries,
	firstAccountID, secondAccountID int64,
	firstAccountAmount, secondAccountAmount int64,
) (firstAccount, secondAccount Account, err error) {
	// Update the first account
	firstAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     firstAccountID,
		Amount: firstAccountAmount,
	})
	if err != nil {
		return
	}

	firstAccount, err = q.GetAccount(ctx, firstAccountID)
	if err != nil {
		return
	}

	// Update the second account
	secondAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     secondAccountID,
		Amount: secondAccountAmount,
	})
	if err != nil {
		return
	}

	secondAccount, err = q.GetAccount(ctx, secondAccountID)
	return
}

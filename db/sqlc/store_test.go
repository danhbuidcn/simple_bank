package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	// Initialize a new store with a test database connection
	store := NewStore(testDB)

	// Create two random accounts for the transfer
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// Number of concurrent transfer transactions
	n := 5
	// Amount to be transferred in each transaction
	amount := int64(10)

	// Create channels to receive errors and results from goroutines
	errs := make(chan error)
	results := make(chan TransferTxResult)

	// Execute n concurrent transfer transactions
	for i := 0; i < n; i++ {
		go func() {
			// Call TransferTx to perform the transfer within a transaction
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
				Currency:      "USD",
			})

			// Send error and result into their respective channels
			errs <- err
			results <- result
		}()
	}

	// Map to track the number of transfers that have been performed
	existed := make(map[int]bool)

	// Receive results from the transactions
	for i := 0; i < n; i++ {
		// Receive and check for errors
		err := <-errs
		require.NoError(t, err)

		// Receive and check the transfer result
		result := <-results
		require.NotEmpty(t, result)

		// Check transfer details
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID) // Verify the sender's account ID
		require.Equal(t, account2.ID, transfer.ToAccountID)   // Verify the receiver's account ID
		require.Equal(t, amount, transfer.Amount)             // Check the transferred amount
		require.NotZero(t, transfer.ID)                       // Ensure transfer ID is valid
		require.NotZero(t, transfer.CreatedAt)                // Ensure transfer creation timestamp is set

		// Verify the transfer record in the database
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check the 'from' account entry (for the sender)
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID) // Verify sender's account ID
		require.Equal(t, -amount, fromEntry.Amount)        // Ensure the transfer amount is negative for the sender
		require.NotZero(t, fromEntry.ID)                   // Ensure entry ID is valid
		require.NotZero(t, fromEntry.CreatedAt)            // Ensure entry creation timestamp is set

		// Verify the 'from' account entry in the database
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// Check the 'to' account entry (for the receiver)
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID) // Verify receiver's account ID
		require.Equal(t, amount, toEntry.Amount)         // Ensure the transfer amount is positive for the receiver
		require.NotZero(t, toEntry.ID)                   // Ensure entry ID is valid
		require.NotZero(t, toEntry.CreatedAt)            // Ensure entry creation timestamp is set

		// Verify the 'to' account entry in the database
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Check the account balances
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID) // Ensure it's the correct sender's account

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID) // Ensure it's the correct receiver's account

		// Check the balances after the transaction
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

		// Calculate the difference in balance for the 'from' account (sender)
		diff1 := account1.Balance - fromAccount.Balance
		// Calculate the difference in balance for the 'to' account (receiver)
		diff2 := toAccount.Balance - account2.Balance
		// Ensure the differences are equal, meaning the total money transferred is consistent
		require.Equal(t, diff1, diff2)
		// Ensure the difference is positive (i.e., some money was transferred)
		require.True(t, diff1 > 0)
		// Ensure the difference is a multiple of the transfer amount
		require.True(t, diff1%amount == 0)

		// Calculate how many transfers have occurred (diff1/amount)
		k := int(diff1 / amount)
		// Ensure k is between 1 and n, meaning the correct number of transfers occurred
		require.True(t, k >= 1 && k <= n)
		// Ensure this transfer has not been processed before
		require.NotContains(t, existed, k)
		// Mark this transfer as processed by adding it to the map
		existed[k] = true
	}

	// After all transactions, check the final updated balances

	// Get the updated balance for account1 (sender)
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	// Get the updated balance for account2 (receiver)
	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	// Print the final balances after all transactions
	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	// Ensure the final balance for account1 is correct (original balance minus the total transferred amount)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	// Ensure the final balance for account2 is correct (original balance plus the total received amount)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)
	n := 10
	amount := int64(10)
	errs := make(chan error)
	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID
		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
				Currency:      "USD",
			})
			errs <- err
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}

func TestValidateAccount(t *testing.T) {
	// Create a random account with a specific currency
	account := createRandomAccount(t)
	account.Currency = "USD"
	account.Balance = 100

	tests := []struct {
		name          string
		accountID     int64
		currency      string
		amount        int64
		checkBalance  bool
		expectedError string
	}{
		{
			name:          "ValidAccount",
			accountID:     account.ID,
			currency:      "USD",
			amount:        50,
			checkBalance:  true,
			expectedError: "",
		},
		{
			name:          "AccountNotFound",
			accountID:     -1,
			currency:      "USD",
			amount:        50,
			checkBalance:  true,
			expectedError: "account [-1] not found",
		},
		{
			name:          "CurrencyMismatch",
			accountID:     account.ID,
			currency:      "EUR",
			amount:        50,
			checkBalance:  true,
			expectedError: fmt.Sprintf("account [%d] currency mismatch: USD vs EUR", account.ID),
		},
		{
			name:          "InsufficientBalance",
			accountID:     account.ID,
			currency:      "USD",
			amount:        150,
			checkBalance:  true,
			expectedError: fmt.Sprintf("account [%d] has insufficient balance", account.ID),
		},
		{
			name:          "NoBalanceCheck",
			accountID:     account.ID,
			currency:      "USD",
			amount:        150,
			checkBalance:  false,
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAccount(context.Background(), testQueries, tt.accountID, tt.currency, tt.amount, tt.checkBalance)
			if tt.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.expectedError)
			}
		})
	}
}

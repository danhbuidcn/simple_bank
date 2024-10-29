package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create a random account
func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// createRandomAccountWithArgs creates a random account with the specified arguments
func createRandomAccountWithArgs(t *testing.T, arg CreateAccountParams) Account {
	user := createRandomUser(t)

	// add username into arg
	arg.Owner = user.Username

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// Test creating an account
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// Test creating an account with invalid data
func TestCreateAccountInvalidCurrency(t *testing.T) {
	// Create an account with an invalid currency code
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: "INVALID",
	}

	_, err := testQueries.CreateAccount(context.Background(), arg)
	require.Error(t, err)
}

// Test retrieving account information
func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	expected_account, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expected_account)

	require.Equal(t, account.ID, expected_account.ID)
	require.Equal(t, account.Owner, expected_account.Owner)
	require.Equal(t, account.Balance, expected_account.Balance)
	require.Equal(t, account.Currency, expected_account.Currency)
	require.Equal(t, account.CreatedAt, expected_account.CreatedAt)
}

// Test retrieving a non-existing account
func TestGetAccountNotFound(t *testing.T) {
	_, err := testQueries.GetAccount(context.Background(), 999999)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error()) // Change to sql.ErrNoRows
}

// Test updating an account
func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: 200,
	}

	expected_account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expected_account)

	require.Equal(t, account.ID, expected_account.ID)
	require.Equal(t, account.Owner, expected_account.Owner)
	require.Equal(t, arg.Balance, expected_account.Balance)
	require.Equal(t, account.Currency, expected_account.Currency)
	require.Equal(t, account.CreatedAt, expected_account.CreatedAt)
}

// Test updating a non-existing account
func TestUpdateAccountNotFound(t *testing.T) {
	arg := UpdateAccountParams{
		ID:      999999,
		Balance: 200,
	}

	_, err := testQueries.UpdateAccount(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error()) // Change to sql.ErrNoRows
}

// Test deleting an account
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	expected_account, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error()) // Change to sql.ErrNoRows
	require.Empty(t, expected_account)
}

// Test deleting a non-existing account
func TestDeleteAccountNotFound(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 999999)
	require.NoError(t, err)
}

// Test listing accounts
func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

// Test listing accounts when there are no results
func TestListAccountsNoResults(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  2,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 0)
}

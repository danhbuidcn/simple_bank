package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/jackc/pgx/v5"
)

// tạo một tài khoản ngẫu nhiên
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
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

// kiểm tra việc tạo một tài khoản
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// kiểm tra tạo tài khoản với dữ liệu không hợp lệ
// func TestCreateAccountInvalidCurrency(t *testing.T) {
// 	// Tạo tài khoản với mã tiền tệ không hợp lệ
// 	arg := CreateAccountParams{
// 		Owner:    "tom",
// 		Balance:  100,
// 		Currency: "INVALID",
// 	}

// 	_, err := testQueries.CreateAccount(context.Background(), arg)
// 	require.Error(t, err)
// }

// kiểm tra việc lấy thông tin tài khoản
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

// kiểm tra việc lấy thông tin tài khoản không tồn tại
func TestGetAccountNotFound(t *testing.T) {
	_, err := testQueries.GetAccount(context.Background(), 999999)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

// TestUpdateAccount kiểm tra việc cập nhật tài khoản.
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

// kiểm tra cập nhật tài khoản không tồn tại
func TestUpdateAccountNotFound(t *testing.T) {
	arg := UpdateAccountParams{
		ID:      999999,
		Balance: 200,
	}

	_, err := testQueries.UpdateAccount(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

// kiểm tra việc xóa tài khoản
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	expected_account, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, expected_account)
}

// kiểm tra xóa tài khoản không tồn tại
func TestDeleteAccountNotFound(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 999999)
	require.NoError(t, err)
}

// kiểm tra việc liệt kê danh sách tài khoản
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

// kiểm tra danh sách tài khoản khi không có kết quả
func TestListAccountsNoResults(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 99999,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 0)
}

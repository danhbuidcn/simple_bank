package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Create a random user with random data
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

// Test creating an user
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

// Test retrieving user information
func TestGetUsert(t *testing.T) {
	user := createRandomUser(t)
	expected_user, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, expected_user)

	require.Equal(t, user.Username, expected_user.Username)
	require.Equal(t, user.HashedPassword, expected_user.HashedPassword)
	require.Equal(t, user.FullName, expected_user.FullName)
	require.Equal(t, user.Email, expected_user.Email)

	// Check if the time is within a second
	require.WithinDuration(t, user.PasswordChangedAt, expected_user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, expected_user.CreatedAt, time.Second)
}

// Test retrieving a non-existing user
func TestGetUserNotFound(t *testing.T) {
	_, err := testQueries.GetUser(context.Background(), util.RandomOwner())
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error()) // Change to sql.ErrNoRows
}

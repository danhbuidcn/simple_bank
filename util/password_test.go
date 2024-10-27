package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// TestPassword tests the password hashing and checking functions
func TestPassword(t *testing.T) {
	password := RandomString(10)

	// Test hashing and checking password
	hash1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)

	// Test checking password
	err = CheckPassword(password, hash1)
	require.NoError(t, err)

	// Test checking wrong password
	wrongPassword := RandomString(10)
	err = CheckPassword(wrongPassword, hash1)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// Test hashing a different password
	hash2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash2)
	require.NotEqual(t, hash1, hash2)
}

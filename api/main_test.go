package api

import (
	"os"
	db "simple_bank/db/sqlc"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// newTestServer creates a new test server
func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

// TestMain is the entry point for the test
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode
	os.Exit(m.Run())          // Exit with code 0: success, 1: fail
}

package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode
	os.Exit(m.Run())          // Exit with code 0: success, 1: fail
}

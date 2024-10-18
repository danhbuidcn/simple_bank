package db

import (
	"database/sql"
	"log"
	"os"
	"simple_bank/config"
	"simple_bank/database"
	"testing"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// Global variable for dbSource
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	var err error
	testDB, err = database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Ensure that the connection is successful
	if err := testDB.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	testQueries = New(testDB)

	// Run the tests
	code := m.Run() // run tests with Test prefix
	if err := testDB.Close(); err != nil {
		log.Fatal("cannot close db connection:", err)
	}
	os.Exit(code) // Exit with code 0: success, 1: fail
}

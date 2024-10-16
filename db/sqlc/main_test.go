package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// Global variable for dbSource
var dbSource string
var testQueries *Queries

func init() {
	// Get connection information from environment variables and initialize dbSource
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Initialize dbSource value
	dbSource = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func TestMain(m *testing.M) {
	// Connect to the database using database/sql
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Ensure that the connection is successful
	if err := conn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	testQueries = New(conn)

	// Run the tests
	code := m.Run() // run tests with Test prefix
	if err := conn.Close(); err != nil {
		log.Fatal("cannot close db connection:", err)
	}
	os.Exit(code) // Exit with code 0: success, 1: fail
}

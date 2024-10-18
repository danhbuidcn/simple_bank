package database

import (
	"database/sql"
	"fmt"
	"simple_bank/config"

	_ "github.com/lib/pq"
)

// ConnectDB connects to the database and returns a database object
func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

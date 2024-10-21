package main

import (
	"log"
	"simple_bank/api"
	"simple_bank/config"
	"simple_bank/database"

	db "simple_bank/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig() // or config.LoadConfig(".")
	log.Println("=======================> Configuration loaded successfully")

	// Connect to database
	conn, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer conn.Close()
	log.Println("=======================> Connected to database successfully")

	// Create database store
	store := db.NewStore(conn)
	log.Println("=======================> Database store created successfully")

	// Create API server
	server := api.NewServer(store)
	log.Println("=======================> API server created successfully")

	// Start server
	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
	log.Println("=======================> Server started successfully on", cfg.ServerAddress)
}

package main

import (
	"log"
	"simple_bank/api"
	"simple_bank/config"
	"simple_bank/database"
	db "simple_bank/db/sqlc"
	"simple_bank/util"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()          // load config from /config/config.go by os.Getenv
	config, err := util.LoadConfig(".") // load config from /util/config.go by viper
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}
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
	server, error := api.NewServer(config, store)
	if error != nil {
		log.Fatalf("Could not create server: %v", error)
	}
	log.Println("=======================> API server created successfully")

	// Start server
	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
	log.Println("=======================> Server started successfully")
}

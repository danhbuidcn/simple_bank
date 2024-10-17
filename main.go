package main

import (
	"log"
	"simple_bank/app"
	"simple_bank/config"
	"simple_bank/database"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// // Perform migration
	// if err := database.RunMigrations(db); err != nil {
	//     log.Fatalf("Migration failed: %v", err)
	// }

	// Start the application
	router := app.SetupRouter()
	log.Println("Application started.")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

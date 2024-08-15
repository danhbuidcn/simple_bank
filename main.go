package main

import (
    "log"
    "simple_bank/config"
    "simple_bank/database"
    "simple_bank/app"
)

func main() {
    // Tải cấu hình
    cfg := config.LoadConfig()

    // Kết nối cơ sở dữ liệu
    db, err := database.ConnectDB(cfg)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

    // // Thực hiện migration
    // if err := database.RunMigrations(db); err != nil {
    //     log.Fatalf("Migration failed: %v", err)
    // }

    // Khởi động ứng dụng
    router := app.SetupRouter()
    log.Println("Application started.")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}

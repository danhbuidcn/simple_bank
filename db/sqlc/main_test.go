package db

import (
	"context"
	"log"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool" // cung cấp quản lý pool kết nối giúp quản lý nhiều kết nối đến cơ sở dữ liệu hiệu quả hơn.
)

// Biến toàn cục cho dbSource
var dbSource string
var testQueries *Queries

func init() {
	// Lấy thông tin kết nối từ biến môi trường và khởi tạo dbSource
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Khởi tạo giá trị dbSource
	dbSource = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func TestMain(m *testing.M) {
	// Kết nối cơ sở dữ liệu bằng pgxpool
  conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	// Chạy các test
	os.Exit(m.Run()) // chạy các test có tiền tố Test và thoát với mã thoát 0:success, 1:fail
}

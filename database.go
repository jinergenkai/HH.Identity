package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Không tìm thấy file .env, dùng biến môi trường hệ thống")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("❌ DB_URL chưa được cấu hình!")
	}

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Không thể kết nối database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("❌ Database không phản hồi:", err)
	}
	fmt.Println("✅ Kết nối database thành công!")
}

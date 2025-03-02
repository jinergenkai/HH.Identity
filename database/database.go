package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// InitDB khởi tạo kết nối đến MongoDB
func InitDB() {
	// mongoURI := os.Getenv("MONGO_URI")
	mongoURI := "mongodb://admin:password@localhost:27017"

	clientOptions := options.Client().ApplyURI(mongoURI)

	// print mongoURI
	fmt.Println("🔗 URI MongoDB:", mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("❌ Lỗi tạo client MongoDB:", err)
	}

	// Kiểm tra kết nối
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Không thể kết nối MongoDB:", err)
	}

	// Gán database sau khi chắc chắn đã kết nối
	DB = client.Database("land_db")
	fmt.Println("✅ Kết nối MongoDB thành công!")
}

// GetCollection trả về một collection cụ thể
func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ Database chưa được khởi tạo!")
	}
	return DB.Collection(name)
}

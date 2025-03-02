package main

import (
	"fmt"
	"log"

	"land_service/database"
	_ "land_service/docs"
	"land_service/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Land Service API
// @version 1.0
// @description API sử dụng Gin & Swagger
// @host localhost:9000
// @BasePath /api

func main() {
	// Khởi tạo database
	database.InitDB()

	// Tạo router với Gin
	r := gin.Default()

	// Định nghĩa route
	r.POST("/api/lands", handlers.CreateLand)
	r.GET("/api/lands", handlers.GetLands)

	// Route Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Chạy server
	fmt.Println("🚀 Server chạy trên cổng 9000...")
	log.Fatal(r.Run(":9000"))
}

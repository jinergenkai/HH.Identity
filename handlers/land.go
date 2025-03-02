package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"land_service/database"
	"land_service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Thêm vùng đất mới
func CreateLand(c *gin.Context) {
	var newLand models.Land

	// Bind JSON vào struct
	if err := c.ShouldBindJSON(&newLand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy collection trong từng request (tránh nil pointer)
	landCollection := database.GetCollection("lands")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := landCollection.InsertOne(ctx, newLand)
	if err != nil {
		log.Println("❌ Lỗi khi thêm vùng đất:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm vùng đất"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thêm vùng đất thành công!", "id": result.InsertedID})
}

// @Summary Lấy danh sách vùng đất
// @Description Trả về danh sách tất cả các vùng đất
// @Tags Lands
// @Produce json
// @Success 200 {array} models.Land
// @Router /api/lands [get]
func GetLands(c *gin.Context) {
	var lands []models.Land

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	landCollection := database.GetCollection("lands")

	cursor, err := landCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách vùng đất"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var land models.Land
		if err := cursor.Decode(&land); err != nil {
			log.Println("❌ Lỗi khi decode vùng đất:", err)
			continue
		}
		lands = append(lands, land)
	}

	c.JSON(http.StatusOK, lands)
}

package main

import (
	"golang_test_day_07/database"
	"golang_test_day_07/models"
	"golang_test_day_07/routes"
	"log"
)

func main() {
	// Kết nối cơ sở dữ liệu
	database.Connect()

	// Tự động migrate mô hình Book
	database.DB.AutoMigrate(&models.Book{})

	// Cài đặt router
	router := routes.SetupRouter()

	// Chạy server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

package main

import (
	"golang-ecomm/config"
	_ "golang-ecomm/docs" // This is required for Swagger docs
	"golang-ecomm/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin-GORM-Redis API
// @version 1.0
// @description This is a sample Gin server with GORM and Redis

func main() {
	router := gin.Default()

	err := config.LoadEnv()
	if err != nil {
		panic("Failed to load .env file")
	}

	mysqlDb, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database!")
	}

	redisCli, err := config.ConnectRedis()
	if err != nil {
		panic("Failed to connect to Redis!")
	}

	routes.SetupRoutes(router, mysqlDb, redisCli)

	// Swagger setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8083")
}

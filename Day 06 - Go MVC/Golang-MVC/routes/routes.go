package routes

import (
	"golang-ecomm/controllers"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis/v8"

	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, redisCli *redis.Client) {

	userController := controllers.NewUserController(db, redisCli)
	router.Group("/api")
	{
		router.GET("/users", userController.GetUsers)
		router.POST("/users", userController.CreateUser)
		router.GET("/users/:id", userController.GetUserByID)
		router.PUT("/users/:id", userController.UpdateUser)
		router.DELETE("/users/:id", userController.DeleteUser)
	}
}

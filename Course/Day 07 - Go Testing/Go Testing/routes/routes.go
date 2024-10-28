package routes

import (
	"github.com/gin-gonic/gin"
	"golang_test_day_07/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Định nghĩa các route cho sách
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}

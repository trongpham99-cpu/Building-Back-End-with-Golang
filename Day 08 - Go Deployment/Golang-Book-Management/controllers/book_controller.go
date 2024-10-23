package controllers

import (
	"golang-book-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{DB: db}
}

// Get all books
func (ctrl *BookController) GetBooks(c *gin.Context) {
	var books []models.Book
	if err := ctrl.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Create a new book
func (ctrl *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// Get a book by ID
func (ctrl *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Update a book by ID
func (ctrl *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Delete a book by ID
func (ctrl *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := ctrl.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

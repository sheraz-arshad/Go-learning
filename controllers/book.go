package controllers

import (
	"learning/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	var UpdateBookInput struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	if err := c.ShouldBindJSON(&UpdateBookInput); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	db.Model(&book).Updates(UpdateBookInput)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	book := models.Book{
		Author: input.Author,
		Title:  input.Title,
	}

	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

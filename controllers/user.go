package controllers

import (
	"fmt"
	"learning/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetUserById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	user := models.User{
		Name:     input.Name,
		Password: input.Password,
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data is bad",
		})
		return
	}

	var user models.User
	if err := db.Where("name = ?", input.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "User not found",
			"success": false,
		})
		return
	}
	fmt.Println(user)
	session := sessions.Default(c)
	session.Set("signed_in", true)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

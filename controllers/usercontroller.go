package controllers

import (
	"DzMart/dtos"
	"DzMart/initializers"
	"DzMart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Createuser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Credit:   0,
		Sub:      false,
		Password: body.Password,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func Getusers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

func Finduser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Updateuser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	var input dtos.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Name = *input.Email
	}
	if input.Password != nil {
		user.Name = *input.Password
	}

	updateresult := initializers.DB.Save(&user)
	if updateresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not updated"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func Deleteuser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	deletionresult := initializers.DB.Delete(&user)
	if deletionresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting User"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "user deleted"})
}

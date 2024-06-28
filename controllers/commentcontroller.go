package controllers

import (
	"DzMart/initializers"
	"DzMart/models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context) {
	var body models.Comment
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customErrs []string
	if body.Review == 0 {
		customErrs = append(customErrs, "Review field is required")
	}
	if body.ProductID == 0 {
		customErrs = append(customErrs, "Product id field is required")
	}
	if body.UserID == 0 {
		customErrs = append(customErrs, "User id field is required")
	}
	if len(customErrs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": customErrs})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, body.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var product models.Product
	if err := initializers.DB.First(&product, body.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var existingComment models.Comment
	if err := initializers.DB.Where("user_id = ? AND product_id = ?", body.UserID, body.ProductID).First(&existingComment).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User has already commented on this product"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing comments"})
		return
	}

	comment := models.Comment{
		UserID:    body.UserID,
		ProductID: body.ProductID,
		Content:   body.Content,
		Review:    body.Review,
		CreatedAt: time.Now(),
	}

	result := initializers.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": comment,
	})
}

func Getcomments(c *gin.Context) {
	var comments []models.Comment
	initializers.DB.Find(&comments)
	c.JSON(200, gin.H{
		"comments": comments,
	})
}

func Findcomment(c *gin.Context) {
	ID := c.Param("id")
	var comment models.Comment
	result := initializers.DB.Where("ID = ?", ID).First(&comment)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func Deletecomment(c *gin.Context) {
	ID := c.Param("id")
	var Comment models.Comment
	result := initializers.DB.Where("ID = ?", ID).First(&Comment)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	deletionresult := initializers.DB.Delete(&Comment)
	if deletionresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting Comment"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Comment deleted"})
}

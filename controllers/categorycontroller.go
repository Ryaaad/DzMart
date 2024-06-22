package controllers

import (
	"DzMart/dtos"
	"DzMart/initializers"
	"DzMart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Createcategory(c *gin.Context) {
	var body models.Category
	if err := c.ShouldBindJSON(&body); err != nil {
		var customErr string

		if body.CatName == "" {
			customErr = "Category Name field is required"
		}

		if customErr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": customErr})
		}
		return
	}

	category := models.Category{
		CatName: body.CatName,
	}

	result := initializers.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"category": category,
	})
}

func Getcategories(c *gin.Context) {
	var categories []models.Category
	initializers.DB.Find(&categories)
	c.JSON(200, gin.H{
		"categories": categories,
	})
}

func Findcategory(c *gin.Context) {
	CatName := c.Param("name")
	var category models.Category
	result := initializers.DB.Where("CatName = ?", CatName).First(&category)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func Updatecategory(c *gin.Context) {
	CatName := c.Param("name")
	var category models.Category
	result := initializers.DB.Where("CatName = ?", CatName).First(&category)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	var input dtos.UpdateCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.CatName = *input.CatName

	updateresult := initializers.DB.Save(&category)
	if updateresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Category not updated"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"category": category})
}

func Deletecategory(c *gin.Context) {
	CatName := c.Param("name")
	var category models.Category
	result := initializers.DB.Where("CatName = ?", CatName).First(&category)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	deletionresult := initializers.DB.Delete(&category)
	if deletionresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting category"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "category deleted"})
}

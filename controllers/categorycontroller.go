package controllers

import (
	"DzMart/models"
	"DzMart/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	category.CatName = c.PostForm("CatName")
	if category.CatName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CatName field is required"})
		return
	}
	img, err := c.FormFile("Category_image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving img: %v", err))
		return
	}
	fileReader, err := img.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening img: %v", err))
		return
	}
	defer fileReader.Close()

	CreateErr := services.CreateCategory(&category, img, c)
	if CreateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": CreateErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"category": category,
	})
}

func GetCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func FindCategory(c *gin.Context) {
	CatName := c.Param("name")
	category, err := services.GetCategoryByID(CatName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func UpdateCategory(c *gin.Context) {
	CatName := c.Param("name")
	img, err := c.FormFile("Category_image")
	if err != nil {
		if err != http.ErrMissingFile {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving img: %v", err))
			return
		}
	}
	NewName := c.PostForm("CatName")
	if NewName == "" {
		NewName = CatName
	}
	if NewName == "" && img == nil {
		return
	}
	if img != nil {
		fileReader, err := img.Open()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening img: %v", err))
			return
		}
		defer fileReader.Close()
	}

	category, updateresult := services.UpdateCategory(CatName, NewName, img, c)
	if updateresult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint("Category not updated : ", updateresult)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"category": category})
}

func DeleteCategory(c *gin.Context) {
	CatName := c.Param("name")
	err := services.DeleteCategory(CatName, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint("err deleting category : ", err)})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "category deleted"})
}

func GetCategoriesImage(c *gin.Context) {
	categoriesimage, err := services.GetAllImageCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categoriesimage": categoriesimage,
	})
}

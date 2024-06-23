package controllers

import (
	"DzMart/dtos"
	"DzMart/initializers"
	"DzMart/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Createuser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		var customErr string

		if body.Email == "" {
			customErr = "Email field is required"
		}

		if body.Name == "" {
			customErr = "Name field is required"
		}
		if body.Email == "" && body.Name == "" {
			customErr = "Name & Email fields are required"
		}
		if customErr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": customErr})
		}
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
		user.Email = *input.Email
	}
	if input.Password != nil {
		user.Password = *input.Password
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

func AddFavorite(c *gin.Context) {
	var favProduct dtos.Favorite
	userID := c.Param("id") // Assuming "id" is the user ID parameter in the URL
	fmt.Println("User ID:", userID)
	if err := c.ShouldBindJSON(&favProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var product models.Product
	fmt.Println("Favorite Product ID:", *favProduct.IDProduct)
	if err := initializers.DB.First(&product, *favProduct.IDProduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	fmt.Println(" Product :", product)

	association := initializers.DB.Model(&user).Association("Fav").Append(&product)
	fmt.Println(" association :", association)

	c.JSON(http.StatusAccepted, gin.H{"msg": "Product added to favorites"})
}

func GetFavorites(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := initializers.DB.Preload("Fav").Find(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{
		"Favorites": user.Fav,
	})
}

func DeleteFavorite(c *gin.Context) {
	userid := c.Param("id")
	productid := c.Param("productid")

	var user models.User
	if err := initializers.DB.Find(&user, userid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var product models.User

	if err := initializers.DB.Find(&product, productid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	association := initializers.DB.Model(&user).Association("Fav").Delete(product)
	fmt.Println(" association :", association)

	c.JSON(http.StatusOK, gin.H{"msg": "Product removed from favorites"})
}

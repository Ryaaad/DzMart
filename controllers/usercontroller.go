package controllers

import (
	"DzMart/dtos"
	"DzMart/models"
	"DzMart/services"
	"DzMart/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
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

	result := services.CreateUser(&body)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": body,
	})
}

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}

func FindUser(c *gin.Context) {
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}
	user, result := services.GetUserById(id)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}
	user, result := services.GetUserById(id)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error()})
		return
	}
	var input dtos.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, updateresult := services.UpdateUser(input, user)
	if updateresult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateresult.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func DeleteUser(c *gin.Context) {
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}
	user, result := services.GetUserById(id)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error()})
		return
	}
	Deleteresult := services.DeleteUser(user)
	if Deleteresult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Deleteresult.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "user deleted"})
}

func AddFavorite(c *gin.Context) {
	var favProduct dtos.Favorite
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}
	if err := c.ShouldBindJSON(&favProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, Geterr := services.GetUserById(id)
	if Geterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("user ", Geterr.Error())})
		return
	}

	product, Geterr := services.GetProductById(*favProduct.IDProduct)
	if Geterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("product ", Geterr.Error())})
		return
	}

	user, Errassociation := services.AddFavorite(user, product)
	if Errassociation != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Errassociation.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user": user})
}

func DeleteFavorite(c *gin.Context) {
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}
	productIDStr := c.Param("productid")

	productID, err := utils.ToUint(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	user, geterr := services.GetUserById(id)
	if geterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("user ", geterr.Error())})
		return
	}
	product, geterr := services.GetProductById(productID)
	if geterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("product ", geterr.Error())})
		return
	}

	DeleteErr := services.DeleteFavorite(*user, *product)
	if DeleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": DeleteErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Product removed from favorites"})
}

func DeleteAllFavorite(c *gin.Context) {
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}
	user, geterr := services.GetUserById(id)
	if geterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("user ", geterr.Error())})
		return
	}

	DeleteErr := services.DeleteAllFavorite(*user)
	if DeleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": DeleteErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Removed All product from favorites"})
}

func GetUserTransactions(c *gin.Context) {
	userid := c.Param("id")
	id, convertErr := utils.ToUint(userid)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid user id"})
		return
	}

	user, GetuserErr := services.GetUserById(id)
	if GetuserErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": GetuserErr.Error()})
		return
	}
	transactions, GetErr := services.GetUserTransactions(user)
	if GetErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": GetErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Transactions": transactions,
	})
}

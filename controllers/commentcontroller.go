package controllers

import (
	"DzMart/models"
	"DzMart/services"
	"DzMart/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	_, Usererr := services.GetUserById(body.UserID)
	if Usererr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("user ", Usererr.Error())})
		return
	}

	_, Producterr := services.GetProductById(body.ProductID)
	if Producterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint("product ", Producterr.Error())})
		return
	}

	CreateResult := services.CreateComment(body)

	if CreateResult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": CreateResult.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": body,
	})
}

func GetAllComments(c *gin.Context) {
	comments, err := services.GetAllComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"comments": comments,
	})
}

func GetComment(c *gin.Context) {
	ID := c.Param("id")
	id, convertErr := utils.ToUint(ID)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid comment id"})
		return
	}
	comment, result := services.GetCommentbyId(id)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{"comment": comment})
}

func Deletecomment(c *gin.Context) {
	ID := c.Param("id")
	id, convertErr := utils.ToUint(ID)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid comment id"})
		return
	}
	comment, Geterr := services.GetCommentbyId(id)
	if Geterr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": Geterr.Error()})
		return
	}
	result := services.Deletecomment(comment)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Comment deleted"})
}

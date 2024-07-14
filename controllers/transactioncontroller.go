package controllers

import (
	"DzMart/dtos"
	"DzMart/services"
	"DzMart/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var body dtos.CreateTransactionInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, GetProductErr := services.GetProductById(*body.ProductID)
	if GetProductErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": GetProductErr.Error()})
		return
	}
	_, GetUserErr := services.GetUserById(*body.UserID)
	if GetUserErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": GetUserErr.Error()})
		return
	}
	transaction, createErr := services.CreateTransaction(body, *product)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction": transaction,
	})
}

func GetAllTransaction(c *gin.Context) {
	transactions, err := services.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{
		"transactions": transactions,
	})
}

func GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	Tid, Converterr := utils.ToUint(id)
	if Converterr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid transaction id"})
		return
	}
	transaction, err := services.GetTransactionById(Tid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"transaction": transaction,
	})
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	Tid, Converterr := utils.ToUint(id)
	if Converterr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid transaction id"})
		return
	}
	transaction, getErr := services.GetTransactionById(Tid)
	if getErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": getErr.Error()})
		return
	}
	var input dtos.UpdateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, updateErr := services.UpdateTransaction(input, transaction)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"transaction": transaction})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	Tid, Converterr := utils.ToUint(id)
	if Converterr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unvalid transaction id"})
		return
	}
	transaction, result := services.GetTransactionById(Tid)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error()})
		return
	}
	deletionresult := services.DeleteTransaction(transaction)
	if deletionresult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deletionresult.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "transaction deleted"})
}

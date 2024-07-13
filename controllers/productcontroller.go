package controllers

import (
	"DzMart/dtos"
	"DzMart/models"
	"DzMart/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addproduct(c *gin.Context) {
	var body models.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		var customErr string
		if body.Category == "" {
			customErr = "Category field is required"
		}
		if body.Description == "" {
			customErr = "Description field is required"
		}
		if body.Name == "" {
			customErr = "Product Name is required"
		}
		if fmt.Sprintf("%d", body.Price) == "" {
			customErr = "Price field is required"
		}
		if fmt.Sprintf("%d", body.Sold) == "" {
			body.Sold = 0
		}
		if fmt.Sprintf("%d", body.Qte) == "" {
			customErr = "Quantity field is required"
		}

		if customErr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": customErr})
		}
		return
	}
	_, GetcategoryErr := services.GetCategoryByID(body.Category)
	if GetcategoryErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": GetcategoryErr.Error()})
		return
	}
	createErr := services.CreateProduct(&body)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"product": body,
	})
}

func Getproducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{
		"products": products,
	})
}

func Findproduct(c *gin.Context) {
	Name := c.Param("name")
	product, err := services.GetProductByName(Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}

func Updateproduct(c *gin.Context) {
	Name := c.Param("name")
	product, getErr := services.GetProductByName(Name)
	if getErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": getErr.Error()})
		return
	}
	var input dtos.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, updateErr := services.UpdateProduct(input, product)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func Deleteproduct(c *gin.Context) {
	Name := c.Param("name")
	product, result := services.GetProductByName(Name)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error()})
		return
	}
	deletionresult := services.DeleteProduct(product, c)
	if deletionresult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deletionresult.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "product deleted"})
}

func AddProductImage(c *gin.Context) {
	ProductName := c.Param("name")
	mainImg, err := c.FormFile("Product_img")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving product img : %v", err))
		return
	}
	var product *models.Product
	product, GetErr := services.GetProductByName(ProductName)
	if GetErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": GetErr.Error()})
		return
	}

	fileReader, err := mainImg.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening file: %v", err))
		return
	}
	defer fileReader.Close()

	product, resultCreate := services.AddProductImage(product, c, mainImg)

	if resultCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Image Added": product.Images,
	})
}

func DeleteProductImage(c *gin.Context) {
	Name := c.Param("name")
	PublicId := c.Param("id")
	resultFind := services.DeleteProductImage(Name, PublicId, c)
	if resultFind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultFind.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "img deleted"})
}

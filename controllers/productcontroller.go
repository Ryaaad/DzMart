package controllers

import (
	"DzMart/dtos"
	"DzMart/initializers"
	interfaces "DzMart/interface"
	"DzMart/models"
	"DzMart/utils"
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
	var category models.Category
	if err := initializers.DB.First(&category, "CatName = ?", body.Category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found " + err.Error()})
		return
	}
	product := models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Sold:        body.Sold,
		Rating:      0,
		Qte:         body.Qte,
		Category:    body.Category,
	}

	result := initializers.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}

func Getproducts(c *gin.Context) {
	var products []models.Product
	initializers.DB.Preload("Comments").Find(&products)
	c.JSON(200, gin.H{
		"products": products,
	})
}

func Findproduct(c *gin.Context) {
	Name := c.Param("name")
	var product []models.Product
	result := initializers.DB.Where("name = ?", Name).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}

func GetproductCategory(c *gin.Context) {
	Category := c.Param("name")
	var product []models.Product
	result := initializers.DB.Where("category = ?", Category).Find(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no product with such category"})
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}
func Updateproduct(c *gin.Context) {
	Name := c.Param("name")
	var product models.Product
	result := initializers.DB.Where("name = ?", Name).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	var input dtos.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Category != nil {
		product.Category = *input.Category
	}

	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Qte != nil {
		product.Qte = *input.Qte
	}
	if input.Sold != nil {
		product.Sold = *input.Sold
	}
	updateresult := initializers.DB.Save(&product)
	if updateresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "product not updated"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func Deleteproduct(c *gin.Context) {
	Name := c.Param("name")
	var product models.Product
	result := initializers.DB.Where("name = ?", Name).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	var imgs []models.ProductImage
	resultFind := initializers.DB.Where("Product = ?", Name).Find(&imgs)
	if resultFind.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultFind.Error.Error()})
		return
	}

	for _, img := range imgs {
		_, err := utils.DestroyImg(c, img.PublicID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		deletionResult := initializers.DB.Delete(&img)
		if deletionResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting img"})
			return
		}
	}

	deletionresult := initializers.DB.Delete(&product)
	if deletionresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting product"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "product deleted"})
}

func AddProductImage(c *gin.Context) {
	ProductName := c.Param("name")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving file: %v", err))
		return
	}
	var product []models.Product
	result := initializers.DB.Where("name = ?", ProductName).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	fmt.Println("product: ", product)
	count := initializers.DB.Model(&product).Association("Images").Count()
	fmt.Println("img count: ", count)

	// Open the uploaded file
	fileReader, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening file: %v", err))
		return
	}
	defer fileReader.Close()
	PublicID := interfaces.BaseImage{
		PublicID: fmt.Sprint(ProductName, ":", count),
	}
	Imageproduct := models.ProductImage{
		Product:   ProductName,
		BaseImage: PublicID,
	}

	resultCreate := initializers.DB.Create(&Imageproduct)
	fmt.Println("new img : ", Imageproduct)

	if resultCreate.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultCreate.Error.Error()})
		return
	}
	utils.UploadImage(c, file, Imageproduct.PublicID)
	c.JSON(http.StatusOK, gin.H{
		"PublicID": Imageproduct.PublicID,
	})
}

func GetProductImages(c *gin.Context) {
	Name := c.Param("name")
	var imgs []models.ProductImage
	resultFind := initializers.DB.Where("Product = ?", Name).Find(&imgs)
	if resultFind.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultFind.Error.Error()})
		return
	}

	var imageUrls []string
	for _, img := range imgs {
		// Use utils function to get asset info, assuming it returns URLs
		pic, err := utils.GetAssetInfo(c, img.PublicID)
		fmt.Println("pics", pic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		imageUrls = append(imageUrls, pic.URL) // Assuming pic has URL field
	}

	// Return the list of image URLs as JSON response
	c.JSON(http.StatusOK, gin.H{"images": imageUrls})
}

func DeleteProductImage(c *gin.Context) {
	Name := c.Param("name")
	PublicId := c.Param("id")
	var img []models.ProductImage

	resultFind := initializers.DB.Where("Product = ?", Name).Where("public_id = ?", PublicId).Find(&img)
	if resultFind.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultFind.Error.Error()})
		return
	}

	deletionresult := initializers.DB.Delete(&img)
	if deletionresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting img"})
		return
	}
	_, err := utils.DestroyImg(c, PublicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "err deleting img"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "img deleted"})

}

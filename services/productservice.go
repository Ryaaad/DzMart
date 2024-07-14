package services

import (
	"DzMart/dtos"
	"DzMart/initializers"
	interfaces "DzMart/interface"
	"DzMart/models"
	"DzMart/utils/cloudinary"
	"context"
	"fmt"
)

func GetAllProducts() ([]models.Product, error) {
	var Products []models.Product
	result := initializers.DB.Find(&Products)
	return Products, result.Error
}

func GetProductByName(Name string) (*models.Product, error) {
	var product models.Product
	result := initializers.DB.Where("Name=?", Name).Preload("Images").Preload("Comments").First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	result := initializers.DB.Preload("Images").Preload("Comments").First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func CreateProduct(product *models.Product) error {
	Newproduct := models.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Sold:        product.Sold,
		Rating:      0,
		Qte:         product.Qte,
		Category:    product.Category,
	}

	resultCreate := initializers.DB.Create(&Newproduct)
	if resultCreate.Error != nil {
		return resultCreate.Error
	}
	return nil
}

func UpdateProduct(input dtos.UpdateProductInput, product *models.Product) (*models.Product, error) {
	OldName := product.Name
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
	if input.Name != nil {
		product.Name = *input.Name
		SaveImgErr := initializers.DB.Model(&models.ProductImage{}).Where("Product = ?", OldName).Update("Product", product.Name)
		if SaveImgErr.Error != nil {
			return nil, SaveImgErr.Error
		}
	}

	SaveErr := initializers.DB.Model(&models.Product{}).Where("Name = ?", OldName).Updates(product)

	if SaveErr.Error != nil {
		return nil, SaveErr.Error
	}

	return product, nil
}

func AddProductImage(product *models.Product, c context.Context, img interface{}) (*models.Product, error) {
	count := initializers.DB.Model(&product).Association("Images").Count()
	fmt.Println("img count: ", count)
	PublicID := fmt.Sprint(product.Name, ":", count)

	resp, Uploaderr := cloudinary.UploadImage(c, img, PublicID)
	if Uploaderr != nil {
		return nil, Uploaderr
	}
	BaseInfoo := interfaces.BaseImage{
		PublicID: PublicID,
		Url:      resp.SecureURL,
	}
	Imageproduct := models.ProductImage{
		Product:   product.Name,
		BaseImage: BaseInfoo,
	}

	resultCreate := initializers.DB.Create(&Imageproduct)

	if resultCreate.Error != nil {
		return nil, resultCreate.Error
	}
	return product, nil
}

func DeleteProduct(product *models.Product, c context.Context) error {
	var imgs []models.ProductImage
	resultFind := initializers.DB.Where("Product = ?", product.Name).Find(&imgs)
	if resultFind.Error != nil {
		return resultFind.Error
	}

	for _, img := range imgs {
		_, err := cloudinary.DestroyImg(c, img.PublicID)
		if err != nil {
			return err
		}
		deletionResult := initializers.DB.Delete(&img)
		if deletionResult.Error != nil {
			return deletionResult.Error
		}
	}

	deletionresult := initializers.DB.Delete(&product)
	if deletionresult.Error != nil {
		return deletionresult.Error
	}
	return nil
}

func DeleteProductImage(Name string, PublicId string, c context.Context) error {
	var img models.ProductImage

	resultFind := initializers.DB.Where("Product = ?", Name).Where("public_id = ?", PublicId).First(&img)
	if resultFind.Error != nil {
		return resultFind.Error
	}
	fmt.Println(img)
	deletionresult := initializers.DB.Delete(&img)
	if deletionresult.Error != nil {
		return deletionresult.Error
	}
	_, err := cloudinary.DestroyImg(c, PublicId)
	if err != nil {
		return err
	}
	return nil
}

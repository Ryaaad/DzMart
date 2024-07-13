package services

import (
	"DzMart/initializers"
	interfaces "DzMart/interface"
	"DzMart/models"
	"DzMart/utils/cloudinary"
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := initializers.DB.Preload("CategoryImage").Find(&categories)
	return categories, result.Error
}

func GetCategoryByID(CatName string) (*models.Category, error) {
	var category models.Category
	result := initializers.DB.Where("CatName=?", CatName).Preload("CategoryImage").First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("category not found")
		} else {
			return nil, result.Error
		}
	}
	return &category, nil
}

func GetCategoryProducts(CatName string) (*models.Category, error) {
	var category models.Category
	result := initializers.DB.Where("CatName=?", CatName).Preload("Items").First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("category not found")
		} else {
			return nil, result.Error
		}
	}
	return &category, nil
}

func CreateCategory(category *models.Category, img interface{}, c context.Context) error {
	PublicID := fmt.Sprint(category.CatName, "img")
	resp, Uploaderr := cloudinary.UploadImage(c, img, PublicID)
	if Uploaderr != nil {
		return Uploaderr
	}
	Baseinfo := interfaces.BaseImage{
		PublicID: PublicID,
		Url:      resp.SecureURL,
	}
	ImageCategory := models.CategoryImage{
		CatName:   category.CatName,
		BaseImage: Baseinfo,
	}
	category.CategoryImage = ImageCategory
	resultCreate := initializers.DB.Create(&ImageCategory)
	if resultCreate.Error != nil {
		return resultCreate.Error
	}
	result := initializers.DB.Create(&category)
	if result.Error != nil {
		return resultCreate.Error
	}
	return nil
}

func UpdateCategory(CatName string, NewName string, img *multipart.FileHeader, c context.Context) (*models.Category, error) {
	category, GetCategoryErr := GetCategoryByID(CatName)
	if GetCategoryErr != nil {
		return nil, GetCategoryErr
	}
	OldPublicID := category.CategoryImage.PublicID
	var imagecategory models.CategoryImage
	resultFind := initializers.DB.Where("public_id = ?", category.CategoryImage.PublicID).First(&imagecategory)
	if resultFind.Error != nil {
		return nil, resultFind.Error
	}
	PublicID := fmt.Sprint(NewName, "img")
	var Baseinfo interfaces.BaseImage
	if img != nil {
		_, err := cloudinary.DestroyImg(c, OldPublicID)
		if err != nil {
			return nil, err
		}
		resp, Uploaderr := cloudinary.UploadImage(c, img, PublicID)
		if Uploaderr != nil {
			return nil, Uploaderr
		}
		Baseinfo = interfaces.BaseImage{
			PublicID: PublicID,
			Url:      resp.SecureURL,
		}
	} else {
		resp, RenameErr := cloudinary.RenameImage(c, OldPublicID, PublicID)
		if RenameErr != nil {
			return nil, RenameErr
		}
		Baseinfo = interfaces.BaseImage{
			PublicID: PublicID,
			Url:      resp.SecureURL,
		}
	}
	imagecategory = models.CategoryImage{
		CatName:   NewName,
		BaseImage: Baseinfo,
	}
	categoryimgSaveErr := initializers.DB.Model(&models.CategoryImage{}).Where("public_id = ?", OldPublicID).Updates(map[string]interface{}{
		"CatName":  NewName,
		"PublicID": imagecategory.PublicID,
		"Url":      imagecategory.Url,
	})

	if categoryimgSaveErr.Error != nil {
		return nil, categoryimgSaveErr.Error
	}
	categorySaveresult := initializers.DB.Model(&models.Category{}).Where("CatName = ?", CatName).Updates(map[string]interface{}{
		"CatName":       NewName,
		"CategoryImage": imagecategory,
	})

	if categorySaveresult.Error != nil {
		return nil, categorySaveresult.Error
	}
	productSaveresult := initializers.DB.Model(&models.Product{}).Where("Category= ?", CatName).Update("Category", NewName)
	if productSaveresult.Error != nil {
		return nil, productSaveresult.Error
	}
	Newcategory, CategoryErr := GetCategoryByID(NewName)
	if CategoryErr != nil {
		return nil, CategoryErr
	}
	return Newcategory, nil
}

func DeleteCategory(CatName string, c context.Context) error {
	var category *models.Category
	category, GetCategoryerr := GetCategoryByID(CatName)
	if GetCategoryerr != nil {
		return GetCategoryerr
	}

	var CategoryImage models.CategoryImage
	resultFind := initializers.DB.Where("public_id = ?", category.CategoryImage.PublicID).First(&CategoryImage)
	if resultFind.Error != nil {
		return resultFind.Error
	}
	_, err := cloudinary.DestroyImg(c, category.CategoryImage.PublicID)
	if err != nil {
		return err
	}
	categoryimageresult := initializers.DB.Delete(&CategoryImage)
	if categoryimageresult.Error != nil {
		return categoryimageresult.Error
	}

	categoryresult := initializers.DB.Delete(&category)
	if categoryresult.Error != nil {
		return categoryresult.Error
	}

	return nil
}

func GetAllImageCategories() ([]models.CategoryImage, error) {
	var imagecategories []models.CategoryImage
	result := initializers.DB.Find(&imagecategories)
	return imagecategories, result.Error
}

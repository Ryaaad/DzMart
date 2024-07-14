package services

import (
	"DzMart/initializers"
	"DzMart/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateComment(body models.Comment) error {
	var existingComment models.Comment
	err := initializers.DB.Where("user_id = ? AND product_id = ?", body.UserID, body.ProductID).First(&existingComment).Error
	if err == nil {
		return fmt.Errorf("user has already commented on this product")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	comment := models.Comment{
		UserID:    body.UserID,
		ProductID: body.ProductID,
		Content:   body.Content,
		Review:    body.Review,
		CreatedAt: time.Now(),
	}

	result := initializers.DB.Create(&comment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	err := initializers.DB.Find(&comments)
	if err.Error != nil {
		return nil, err.Error
	}
	return comments, nil
}

func GetCommentbyId(id uint) (*models.Comment, error) {
	var comment models.Comment
	result := initializers.DB.Where("ID = ?", id).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func Deletecomment(comment *models.Comment) error {
	Deleteresult := initializers.DB.Delete(&comment)
	if Deleteresult.Error != nil {
		return Deleteresult.Error
	}
	return nil
}

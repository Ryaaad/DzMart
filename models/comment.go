package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"CommentID"`
	UserID    uint      `gorm:"not null" json:"UserID"`
	ProductID uint      `gorm:"not null" json:"ProductID"`
	Content   string    `gorm:"not null" json:"content" binding:"required,min=1"`
	Review    int       `gorm:"not null" json:"Review" binding:"required,gte=0,lte=5"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (comment *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var product Product
	if err := tx.Find(&product, comment.ProductID).Error; err != nil {
		return errors.New("product not found")
	}
	var comments []Comment
	if err := tx.Model(&Comment{}).Where("product_id = ?", comment.ProductID).Find(&comments).Error; err != nil {
		return errors.New("failed to count product comments")
	}
	product.Rating = MeanScore(comments)

	if err := tx.Save(&product).Error; err != nil {
		return errors.New("product not updated")
	}
	return nil
}
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var product Product
	if err := tx.Find(&product, comment.ProductID).Error; err != nil {
		return errors.New("product not found")
	}
	var comments []Comment
	if err := tx.Model(&Comment{}).Where("product_id = ?", comment.ProductID).Find(&comments).Error; err != nil {
		return errors.New("failed to count product comments")
	}
	product.Rating = MeanScore(comments)

	if err := tx.Save(&product).Error; err != nil {
		return errors.New("product not updated")
	}
	return nil
}

func MeanScore(Reviews []Comment) float64 {
	if len(Reviews) < 1 {
		return 0
	}
	if len(Reviews) == 1 {
		return float64(Reviews[0].Review)
	}
	var score float64
	score = 0
	for _, comment := range Reviews {
		score = score + float64(comment.Review)
	}
	score = score / float64(len(Reviews))
	return float64(score)
}

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
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("product_id = ?", comment.ProductID).Count(&commentCount).Error; err != nil {
		return errors.New("failed to count product comments")
	}

	if commentCount > 1 {
		product.Rating = (product.Rating + float64(comment.Review)) / 2
	} else {
		product.Rating = float64(comment.Review)
	}

	if err := tx.Save(&product).Error; err != nil {
		return errors.New("product not updated")
	}
	return nil

}

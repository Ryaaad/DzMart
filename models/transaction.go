package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"TransactionID"`
	UserID    uint      `gorm:"not null" json:"UserID"`
	ProductID uint      `gorm:"not null" json:"ProductID"`
	Qte       int       `gorm:"not null" json:"Qte" binding:"required,gte=0"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Total     int       `gorm:"not null" json:"Total_price" binding:"required,gte=0"`
}

func (transaction *Transaction) AfterCreate(tx *gorm.DB) (err error) {
	var product Product
	if err := tx.Find(&product, transaction.ProductID).Error; err != nil {
		return errors.New("product not found")
	}
	product.Qte = product.Qte - transaction.Qte
	if err := tx.Save(&product).Error; err != nil {
		return errors.New("product not updated")
	}
	return nil
}

func (transaction *Transaction) AfterDelete(tx *gorm.DB) (err error) {
	var product Product
	if err := tx.Find(&product, transaction.ProductID).Error; err != nil {
		return errors.New("product not found")
	}
	product.Qte = product.Qte + transaction.Qte
	if err := tx.Save(&product).Error; err != nil {
		return errors.New("product not updated")
	}
	return nil
}

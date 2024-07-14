package services

import (
	"DzMart/dtos"
	"DzMart/initializers"
	"DzMart/models"
	"fmt"
	"time"
)

func GetAllTransactions() ([]models.Transaction, error) {
	var Transactions []models.Transaction
	result := initializers.DB.Find(&Transactions)
	return Transactions, result.Error
}

func GetTransactionById(id uint) (*models.Transaction, error) {
	var Transaction models.Transaction
	result := initializers.DB.First(&Transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Transaction, nil
}

func CreateTransaction(body dtos.CreateTransactionInput, product models.Product) (*models.Transaction, error) {
	if *body.Qte > product.Qte {
		return nil, fmt.Errorf("quantity selected is more than there is in stock")
	}
	RealPrice := product.Price - product.Price*product.Sold
	TotalPrice := RealPrice * *body.Qte
	transaction := models.Transaction{
		UserID:    *body.UserID,
		ProductID: *body.ProductID,
		Qte:       *body.Qte,
		CreatedAt: time.Now(),
		Total:     TotalPrice,
	}

	result := initializers.DB.Create(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func UpdateTransaction(input dtos.UpdateTransactionInput, transaction *models.Transaction) (*models.Transaction, error) {
	product, err := GetProductById(transaction.ProductID)
	if err != nil {
		return nil, err
	}
	RealPrice := product.Price - product.Price*product.Sold
	NewTotal := RealPrice * *input.Qte
	SaveErr := initializers.DB.Model(&models.Transaction{}).Where("id = ?", transaction.ID).Updates(map[string]interface{}{
		"Qte":   *input.Qte,
		"Total": NewTotal,
	})
	if SaveErr.Error != nil {
		return nil, SaveErr.Error
	}
	NewQte := transaction.Qte - *input.Qte
	NewQte = product.Qte + NewQte
	SaveProductErr := initializers.DB.Model(&models.Product{}).Where("Name = ?", product.Name).Updates(map[string]interface{}{
		"Qte": NewQte,
	})
	if SaveProductErr.Error != nil {
		return nil, SaveProductErr.Error
	}
	transaction, finderr := GetTransactionById(transaction.ID)
	if finderr != nil {
		return nil, finderr
	}
	return transaction, nil
}

func DeleteTransaction(transaction *models.Transaction) error {
	Deleteresult := initializers.DB.Delete(&transaction)
	if Deleteresult.Error != nil {
		return Deleteresult.Error
	}
	return nil
}

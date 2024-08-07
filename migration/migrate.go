package migration

import (
	"DzMart/initializers"
	"DzMart/models"
	"log"
)

func DropTables() {
	initializers.DB.Migrator().DropTable(&models.User{}, &models.Category{}, &models.Product{}, &models.Comment{}, &models.ProductImage{}, &models.CategoryImage{}, &models.Transaction{})
}

func MigrateAllTables() error {
	// DropTables() //  remove this in production!
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Comment{}, &models.ProductImage{}, &models.CategoryImage{}, &models.Transaction{})
	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
		return err
	}
	return nil
}

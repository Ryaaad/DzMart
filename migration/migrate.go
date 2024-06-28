package migration

import (
	"DzMart/initializers"
	"DzMart/models"
	"log"
)

func DropTables() {
	initializers.DB.Migrator().DropTable(&models.User{}, &models.Category{}, &models.Product{}, &models.Comment{})
}

func MigrateAllTables() error {
	// DropTables() //  remove this in production!
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
		return err
	}
	return nil
}

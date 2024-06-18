package migration

import (
	"DzMart/initializers"
	"DzMart/models"
	"log"
)

func DropTables() {
	initializers.DB.Migrator().DropTable(&models.User{}, &models.Category{})
}

func MigrateAllTables() error {
	// DropTables() // Only for development, remove this in production!
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Category{})
	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
		return err
	}
	return nil
}

package migration

import (
	"DzMart/initializers"
	"DzMart/models"
)

func MigrateAllTables() error {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

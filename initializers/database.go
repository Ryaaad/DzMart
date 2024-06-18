package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("DzMart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		fmt.Println("Connected")
	}
	return err
}

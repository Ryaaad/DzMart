package main

import (
	"DzMart/controllers"
	"DzMart/initializers"
	"DzMart/migration"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	DBerr := initializers.ConnectToDB()
	if DBerr != nil {
		log.Fatal(DBerr)
	}
	Migraterr := migration.MigrateAllTables()
	if Migraterr != nil {
		log.Fatal(DBerr)
	}
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	// user
	r.GET("/users", controllers.Getusers)
	r.POST("/users", controllers.Createuser)
	r.GET("/users/:id", controllers.Finduser)
	r.PUT("/users/:id", controllers.Updateuser)
	r.DELETE("/users/:id", controllers.Deleteuser)

	return r
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()
	r.Run()
}

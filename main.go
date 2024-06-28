package main

import (
	"DzMart/initializers"
	"DzMart/migration"
	"DzMart/routes"
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
	routes.UserRoutes(r)
	routes.CategoryRoutes(r)
	routes.ProductRoutes(r)
	routes.CommentRoutes(r)
	return r
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()
	r.Run()
}

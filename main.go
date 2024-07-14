package main

import (
	"DzMart/initializers"
	"DzMart/migration"
	"DzMart/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	DBerr := initializers.ConnectToDB()
	if DBerr != nil {
		fmt.Printf("Failed to Connect to database: %v\n", DBerr)
		return
	}
	Migraterr := migration.MigrateAllTables()
	if Migraterr != nil {
		fmt.Printf("Failed to migrate: %v\n", Migraterr)
		return
	}
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.UserRoutes(r)
	routes.CategoryRoutes(r)
	routes.ProductRoutes(r)
	routes.CommentRoutes(r)
	routes.TransactionRoutes(r)
	return r
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()
	r.Run()
}

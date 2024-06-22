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
	users := r.Group("/users")
	{
		users.GET("/", controllers.Getusers)
		users.POST("/", controllers.Createuser)
		users.GET("/:id", controllers.Finduser)
		users.PUT("/:id", controllers.Updateuser)
		users.DELETE("/:id", controllers.Deleteuser)
	}

	// category
	categories := r.Group("/categories")
	{
		categories.GET("/", controllers.Getcategories)
		categories.POST("/", controllers.Createcategory)
		categories.GET("/:name", controllers.Findcategory)
		categories.PUT("/:name", controllers.Updatecategory)
		categories.DELETE("/:name", controllers.Deletecategory)
		// categoriesGET("/:category/products", controllers.GetproductCategory)
		// categoriesGET("/:category/products/:name", controllers.Getproducts)
	}

	// product
	products := r.Group("/products")
	{
		products.GET("/", controllers.Getproducts)
		products.POST("/", controllers.Addproduct)
		products.GET("/:name", controllers.Findproduct)
		products.PUT("/:name", controllers.Updateproduct)
		products.DELETE("/:name", controllers.Deleteproduct)
	}

	return r
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()
	r.Run()
}

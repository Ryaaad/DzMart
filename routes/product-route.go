package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.GET("/", controllers.Getproducts)
		products.POST("/", controllers.Addproduct)
		products.GET("/:name", controllers.Findproduct)
		products.PUT("/:name", controllers.Updateproduct)
		products.DELETE("/:name", controllers.Deleteproduct)
		products.POST("/images/:name", controllers.AddProductImage)
		products.DELETE("/images/:name/:id", controllers.DeleteProductImage)
	}
}

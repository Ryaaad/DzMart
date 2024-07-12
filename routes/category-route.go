package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		categories.GET("/", controllers.GetCategories)
		categories.POST("/", controllers.CreateCategory)
		categories.GET("/images", controllers.GetCategoriesImage)
		categories.GET("/:name", controllers.GetCategory)
		categories.PUT("/:name", controllers.UpdateCategory)
		categories.DELETE("/:name", controllers.DeleteCategory)
		categories.GET("/:name/products", controllers.GetCategoryproducts)
	}
}

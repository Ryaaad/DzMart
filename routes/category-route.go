package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		categories.GET("/", controllers.Getcategories)
		categories.POST("/", controllers.Createcategory)
		categories.GET("/:name", controllers.Findcategory)
		categories.PUT("/:name", controllers.Updatecategory)
		categories.DELETE("/:name", controllers.Deletecategory)
		categories.GET("/:name/products", controllers.GetproductCategory)
	}
}

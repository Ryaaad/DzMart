package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	users := r.Group("/transactions")
	{
		users.GET("/", controllers.GetAllTransaction)
		users.POST("/", controllers.CreateTransaction)
		users.GET("/:id", controllers.GetTransactionById)
		users.PUT("/:id", controllers.UpdateTransaction)
		users.DELETE("/:id", controllers.DeleteTransaction)

		// users.POST(":id/Favorites", controllers.AddFavorite)
		// users.DELETE(":id/Favorites", controllers.DeleteAllFavorite)
		// users.DELETE(":id/Favorites/:productid", controllers.DeleteFavorite)
	}
}

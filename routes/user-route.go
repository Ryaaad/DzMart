package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.POST("/", controllers.CreateUser)
		users.GET("/:id", controllers.FindUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)

		users.POST(":id/Favorites", controllers.AddFavorite)
		users.DELETE(":id/Favorites", controllers.DeleteAllFavorite)
		users.DELETE(":id/Favorites/:productid", controllers.DeleteFavorite)
	}
}

package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", controllers.Getusers)
		users.POST("/", controllers.Createuser)
		users.GET("/:id", controllers.Finduser)
		users.PUT("/:id", controllers.Updateuser)
		users.DELETE("/:id", controllers.Deleteuser)

		users.POST(":id/Favorites", controllers.AddFavorite)
		users.GET(":id/Favorites", controllers.GetFavorites)
		users.DELETE(":id/Favorites/:productid", controllers.DeleteFavorite)
	}
}

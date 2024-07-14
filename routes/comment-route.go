package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	comments := r.Group("/comments")
	{
		comments.GET("/", controllers.GetAllComments)
		comments.POST("/", controllers.CreateComment)
		comments.GET("/:id", controllers.GetComment)
		comments.DELETE("/:id", controllers.Deletecomment)
	}
}

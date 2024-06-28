package routes

import (
	"DzMart/controllers"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	comments := r.Group("/comments")
	{
		comments.GET("/", controllers.Getcomments)
		comments.POST("/", controllers.CreateComment)
		comments.GET("/:id", controllers.Findcomment)
		comments.DELETE("/:id", controllers.Deletecomment)
	}
}

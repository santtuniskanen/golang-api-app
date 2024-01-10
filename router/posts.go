package router

import (
	"github.com/gin-gonic/gin"
	"github.com/santtuniskanen/golang-restapi/controllers"
)

func PostsRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.DELETE("/:id", controllers.DeletePost)
	posts.POST("/", controllers.CreatePost)
	posts.PUT("/:id", controllers.UpdatePost)
	posts.GET("/", controllers.GetPosts)
	posts.GET("/:id", controllers.GetPostByID)
}

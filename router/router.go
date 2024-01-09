package router

import (
	"github.com/gin-gonic/gin"
	"github.com/santtuniskanen/golang-restapi/controllers"
)

func Routes() {
	r := gin.Default()

	r.DELETE("/posts/:id", controllers.DeletePost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)

	r.Run()
}

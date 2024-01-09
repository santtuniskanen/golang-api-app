package router

import (
	"github.com/gin-gonic/gin"
	"github.com/santtuniskanen/golang-restapi/controllers"
)

func Routes() {
	r := gin.Default()

	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsGetAll)
	r.GET("/posts/:id", controllers.PostsGetOne)

	r.Run()
}

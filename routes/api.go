package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/controller"
)

var BlogController = &controller.BlogController{}

func SetupApiRoutes(r *gin.Engine) {
	r.GET("/api/xxx", func(ctx *gin.Context) {
		fmt.Println("test")
	})

	r.GET("/api/blog", BlogController.ListBlogs)
}

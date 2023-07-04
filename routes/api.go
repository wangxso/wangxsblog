package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/controller"
)

var BlogController = &controller.BlogController{}
var UserController = &controller.UserController{}

func SetupApiRoutes(r *gin.Engine) {
	r.GET("/api/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	r.GET("/api/blog", BlogController.ListBlogsByPage)
	r.POST("/api/blog", BlogController.CreateBlog)
	r.POST("/api/login", UserController.Login)
}

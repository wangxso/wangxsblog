package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/middleway"
)

func SetupAuthRoutes(r *gin.Engine) {
	r.GET("/auth/get", middleway.AuthMiddleware("wangxso_secret_key"), func(c *gin.Context) {
		userRole := c.GetString("user_role")
		if userRole != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access: admin only"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "admin access"})
	})
}

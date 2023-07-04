package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/models"
	"github.com/wangxso/wangxsoblog/services"
)

type UserController struct {
	u *services.UserService
}

func (ctrl *UserController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := ctrl.u.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": token})
}

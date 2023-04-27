package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/models"
	"github.com/wangxso/wangxsoblog/services"
)

type BlogController struct {
	s *services.BlogService
}

func (ctrl *BlogController) ListBlogs(c *gin.Context) {
	var blogs *[]models.Blog
	blogs, err := ctrl.s.ListBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

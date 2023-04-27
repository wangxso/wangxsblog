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

func (ctrl *BlogController) CreateBlog(c *gin.Context) {
	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.s.CreateBlog(&blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blog, "message": "success"})
}

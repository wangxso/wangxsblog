package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxso/wangxsoblog/services"
)

type CommentController struct {
	c *services.CommentService
}

func (ctrl *CommentController) ListCommentsByBlogId(c *gin.Context) {
	blogId := c.Query("blogId")
	if blogId == "" {
		c.JSON(400, gin.H{"error": "blogId is required"})
		return
	}
	comments, err := ctrl.c.ListCommentsByBlogId(blogId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": comments})
}

func (ctrl *CommentController) CreateCommentByBlogId(c *gin.Context) {
	// TODO
}

func (ctrl *CommentController) DeleteCommentByID(c *gin.Context) {
	// TODO
}

func (ctrl *CommentController) UpdateCommentByID(c *gin.Context) {
	// TODO
}

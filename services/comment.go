package services

import (
	"github.com/jinzhu/gorm"
	"github.com/wangxso/wangxsoblog/models"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db}
}

func (s *CommentService) ListCommentsByBlogId(blogId string) ([]*models.Comment, error) {
	comments := []*models.Comment{}
	if err := s.db.Where("blog_id = ?", blogId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

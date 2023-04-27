package services

import (
	"github.com/jinzhu/gorm"
	"github.com/wangxso/wangxsoblog/models"
)

type BlogService struct {
	db *gorm.DB
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{db}
}
func (s *BlogService) CreateBlog(blog *models.Blog) error {
	return models.CreateBlog(blog)
}
func (s *BlogService) GetBlog(id uint) (*models.Blog, error) {
	blog, err := models.GetBlogByID(id)
	return blog, err
}
func (s *BlogService) ListBlogs() (*[]models.Blog, error) {
	blogs, err := models.GetBlogs()
	return &blogs, err
}
func (s *BlogService) UpdateBlog(id uint, blog *models.Blog) error {
	oldBlog, err := s.GetBlog(id)
	if err != nil {
		return err
	}
	oldBlog.Title = blog.Title
	oldBlog.Content = blog.Content
	return models.UpdateBlogByID(oldBlog.ID, oldBlog)
}

func (s *BlogService) DeleteBlog(id uint) error {
	return models.DeleteBlogByID(id)
}

func (s *BlogService) CreateComment(blogID uint, comment *models.Comment) error {
	blog, err := s.GetBlog(blogID)
	if err != nil {
		return err
	}
	comment.BlogID = blogID
	comment.UserID = blog.UserID // 从 blog 中获取作者的 ID
	err = models.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogService) DeleteComment(id uint) error {
	return models.DeleteCommentByID(id)
}

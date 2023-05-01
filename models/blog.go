package models

import (
	"time"
)

type Blog struct {
	BaseModel
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   uint      `json:"userid"`
	User     User      `json:"user"`
	Comments []Comment `json:"comments"`
	CreateAt time.Time
	UpdateAt time.Time
}

type BlogsResponse struct {
	TotalCount int64
	TotalPage  int
	Blogs      []Blog
}

func CreateBlog(blog *Blog) error {
	return db.Create(blog).Error
}

func GetBlogByID(id uint) (*Blog, error) {
	var blog Blog
	err := db.Preload("Comments").Where("id = ?", id).First(&blog).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func GetBlogs() ([]Blog, error) {
	var blogs []Blog
	err := db.Preload("Comments").Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func UpdateBlogByID(id uint, data *Blog) error {
	err := db.Model(&Blog{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBlogByID(id uint) error {
	err := db.Where("id = ?", id).Delete(&Blog{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 分页查询博客列表
func GetBlogsByPage(page, size uint) ([]Blog, error) {
	var blogs []Blog
	// 判断 page 和 size 是否为 0
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	offset := (page - 1) * size
	err := db.Preload("Comments").Limit(size).Offset(offset).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetBlogCount() int64 {
	var count int64
	err := db.Model(&Blog{}).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

package models

import "time"

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

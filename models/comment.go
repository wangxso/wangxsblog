package models

import "time"

type Comment struct {
	BaseModel
	Content  string `json:"content"`
	UserID   uint   `json:"userid"`
	User     User   `json:"user"`
	BlogID   uint   `json:"blogid"`
	CreateAt time.Time
	UpdateAt time.Time
}

func CreateComment(c *Comment) error {
	return db.Create(&c).Error
}

func GetCommentById(id uint) (Comment, error) {
	var c Comment
	err := db.First(&c, id).Error
	return c, err
}

func GetAllComment(c *Comment) ([]Comment, error) {
	var comments []Comment
	err := db.Find(&comments).Error
	return comments, err
}

func UpdateComment(c *Comment) error {
	return db.Save(&c).Error
}

func DeleteComment(c *Comment) error {
	return db.Delete(&c).Error
}

func DeleteCommentByID(id uint) error {
	err := db.Where("id = ?", id).Delete(&Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

func ListCommentsByBlogId(blogId string) ([]Comment, error) {
	var comments []Comment
	err := db.Where("blog_id = ?", blogId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

package tests

import (
	"log"
	"testing"

	"github.com/wangxso/wangxsoblog/models"
)

func TestBlogCreate(t *testing.T) {

	user, err := models.FindUserByEmail("wangxs@ww.com")
	if err != nil {
		log.Fatal(err)
	}

	blog := &models.Blog{
		Title:   "测试文章",
		Content: "测试内容",
		UserID:  user.ID,
		User:    *user,
	}

	models.CreateBlog(blog)

	blog_tmp, err := models.GetBlogByID(blog.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(blog_tmp)
}

package tests

import (
	"log"
	"testing"

	"github.com/wangxso/wangxsoblog/models"
)

func TestUserCreate(t *testing.T) {

	user := models.User{
		Username: "wangxs",
		Password: "12345678",
		Email:    "wangxs@ww.com",
	}

	models.CreateUser(&user)
}

func TestUserDelete(t *testing.T) {
	user, err := models.FindUserByID(2)
	if err != nil {
		log.Fatal("read failed at err, ", err)
		t.Fail()
	}

	models.DeleteUser(user)
}

func TestFindUserByEmail(t *testing.T) {
	user, err := models.FindUserByEmail("wangxs@ww.com")
	if err != nil {
		log.Fatal("not found user, err=", err)
		t.Fail()
	}
	log.Println(user)
}

func TestCheckPassword(t *testing.T) {
	user, err := models.FindUserByEmail("wangxs@ww.com")
	if err != nil {
		t.Fail()
	}

	res := user.CheckPassword("12345678")
	if res != nil {
		log.Fatal("error password, err=", err)
	} else {
		log.Println("Success")
	}
}

func TestGenrateToken(t *testing.T) {
	user, err := models.FindUserByEmail("wangxs@ww.com")
	if err != nil {
		log.Fatal(err)
	}
	token, err := user.GenerateJWTToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(token)
}

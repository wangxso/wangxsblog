package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/wangxso/wangxsoblog/utils"
)

var db *gorm.DB

func init() {
	config, err := utils.GetConfig()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err.Error())
	}

	db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Db.User,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Name,
	),
	)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err.Error())
	}
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Comment{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

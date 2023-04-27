package models

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
	CreateAt time.Time
	UpdateAt time.Time
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	scope.SetColumn("Password", string(hashedPassword))
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type JWTToken struct {
	UserID uint
	jwt.StandardClaims
}

func (u *User) GenerateJWTToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTToken{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "blog_app",
		},
	})
	return token.SignedString([]byte("wangxso_secret_key"))
}

func CreateUser(u *User) error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByID(id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(u *User) error {
	if err := db.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(u *User) error {
	if err := db.Delete(u).Error; err != nil {
		return err
	}
	return nil
}

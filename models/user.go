package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const secret = "wangxso_secret_key"

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
	Roles    string `json:"roles"`
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
	return token.SignedString([]byte(secret))
}

// 解析token
func (u *User) CheckToken(authHeader string) (jwt.MapClaims, error) {
	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 判断token的加密方式是否为HMAC
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // 判断token是否有效
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}

func (u *User) GetUserIDFromCLaims(claims jwt.MapClaims) (uint, error) {
	userID, ok := claims["UserID"].(int64)
	if !ok {
		return 0, errors.New("user id not found in claims")
	}
	return uint(userID), nil
}

func (u *User) CheckUserHasRole(claims jwt.MapClaims, role string) bool {
	roles, ok := claims["Roles"].([]interface{})
	if !ok {
		return false
	}
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
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

package services

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/wangxso/wangxsoblog/models"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) Login(username, password string) (string, error) {
	user := models.User{}
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("invalid email/password combination")
	}

	token, err := user.GenerateJWTToken()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) VerifyJWTToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTToken{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("wangxso_secret"), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*models.JWTToken); ok {
		return claims.UserID, nil
	}
	return 0, errors.New("invalid jwt token")
}

package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.User
}

func NewAuthService(repo repository.User) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user, err := s.repo.Get(user.Username)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	} else if err == nil {
		return 0, fmt.Errorf("username is invalid or already taken")
	}

	password, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = password

	return s.repo.Create(user)
}

func (s *AuthService) CheckUser(username, password string) (int, error) {
	user, err := s.repo.Get(username)
	if err != nil {
		return 0, err
	}

	if !checkPasswordHash(password, user.Password) {
		return 0, fmt.Errorf("incorrect username or password")
	}

	return user.ID, nil
}

func (s *AuthService) GenerateToken(userID int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

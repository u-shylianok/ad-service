package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
	"github.com/u-shylianok/ad-service/internal/secure"
)

const (
	signingKey = "nn6gzv&xTae8bqO!Rhrd8Po$30_XAk"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo   repository.User
	hasher secure.Hasher
}

func NewAuthService(repo repository.User, hasher secure.Hasher) *AuthService {
	return &AuthService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	_, err := s.repo.Get(user.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	} else if err == nil {
		return 0, fmt.Errorf("username is invalid or already taken")
	}

	password, err := s.hasher.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = password

	return s.repo.Create(user)
}

func (s *AuthService) CheckUser(username, password string) (int, error) {
	user, err := s.repo.Get(username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("incorrect username or password")
	}

	if !s.hasher.CheckPasswordHash(password, user.Password) {
		return 0, fmt.Errorf("incorrect username or password")
	}

	return user.ID, nil
}

func (s *AuthService) GenerateToken(userID int) (string, int64, error) {
	expiresAt := time.Now().Add(tokenTTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	tokenStr, err := token.SignedString([]byte(signingKey))
	return tokenStr, expiresAt, err
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {

	logrus.WithField("token", accessToken).Info("token")
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

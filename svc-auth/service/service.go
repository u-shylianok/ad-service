package service

import (
	"github.com/u-shylianok/ad-service/svc-auth/internal/secure"
	"github.com/u-shylianok/ad-service/svc-auth/model"
	"github.com/u-shylianok/ad-service/svc-auth/repository"
)

type Service struct {
	Auth
}

type Auth interface {
	CreateUser(user model.User) (int, error)
	CheckUser(username, password string) (int, error)
	GenerateToken(userID int) (string, int64, error)
	ParseToken(accessToken string) (int, error)
}

func NewService(repos *repository.Repository, secure *secure.Secure) *Service {
	return &Service{
		Auth: NewAuthService(repos.User, secure.Hasher),
	}
}

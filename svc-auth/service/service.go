package service

import (
	"github.com/u-shylianok/ad-service/svc-auth/domain/model"
	"github.com/u-shylianok/ad-service/svc-auth/internal/secure"
	"github.com/u-shylianok/ad-service/svc-auth/repository"
)

type Service struct {
	Auth
}

type Auth interface {
	CreateUser(user model.User) (uint32, error)
	CheckUser(username, password string) (uint32, error)
	GenerateToken(userID uint32) (string, int64, error)
	ParseToken(accessToken string) (uint32, error)

	GetUser(userID uint32) (model.UserResponse, error)
}

func NewService(repos *repository.Repository, secure *secure.Secure) *Service {
	return &Service{
		Auth: NewAuthService(repos.User, secure.Hasher),
	}
}

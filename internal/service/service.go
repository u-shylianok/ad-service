package service

import (
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Ad interface {
	CreateAd(ad model.AdRequest) (int, error)
	ListAds(sortBy, order string) ([]model.AdResponse, error)
	GetAd(adId int, fields []string) (model.AdResponse, error)
}

type Service struct {
	Auth
	Ad
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
		Ad:   NewAdService(repos.Ad),
	}
}

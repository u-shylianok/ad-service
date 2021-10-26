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
	ListAds(params []model.AdsSortingParam) ([]model.AdResponse, error)
	GetAd(adID int, fields model.AdOptionalFieldsParam) (model.AdResponse, error)
	UpdateAd(adID int, ad model.AdRequest) error
	DeleteAd(adID int) error
}

type Photo interface {
	ListPhotos(adID int) ([]string, error)
}

type Tag interface {
	ListTags(adID int) ([]string, error)
}

type Service struct {
	Auth
	Ad
	Photo
	Tag
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:  NewAuthService(repos.Auth),
		Ad:    NewAdService(repos.Ad, repos.Photo, repos.Tag),
		Photo: NewPhotoService(repos.Photo),
		Tag:   NewTagService(repos.Tag),
	}
}

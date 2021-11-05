package service

import (
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	CheckUser(username, password string) (int, error)
	GenerateToken(userID int) (string, error)
	ParseToken(token string) (int, error)
}

type Ad interface {
	CreateAd(userID int, ad model.AdRequest) (int, error)
	ListAds(params []model.AdsSortingParam) ([]model.AdResponse, error)
	SearchAds(filter model.AdFilter) ([]model.AdResponse, error)
	GetAd(adID int, fields model.AdOptionalFieldsParam) (model.AdResponse, error)
	UpdateAd(userID, adID int, ad model.AdRequest) error
	DeleteAd(userID, adID int) error
}

type Photo interface {
	ListPhotos() ([]string, error)
	ListAdPhotos(adID int) ([]string, error)
}

type Tag interface {
	ListTags() ([]string, error)
	ListAdTags(adID int) ([]string, error)
}

type Service struct {
	Auth
	Ad
	Photo
	Tag
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:  NewAuthService(repos.User),
		Ad:    NewAdService(repos.Ad, repos.User, repos.Photo, repos.Tag),
		Photo: NewPhotoService(repos.Photo),
		Tag:   NewTagService(repos.Tag),
	}
}

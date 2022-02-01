package service

import (
	"github.com/u-shylianok/ad-service/svc-ads/model"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
)

type Service struct {
	Ad
	Photo
	Tag
}

type Ad interface {
	// ListAds(params []model.AdsSortingParam) ([]model.AdResponse, error)
	// SearchAds(filter model.AdFilter) ([]model.AdResponse, error)
	GetAd(adID uint32, fields model.GetAdOptional) (model.AdResponse, error)
	CreateAd(userID int, ad model.AdRequest) (int, error)
	// UpdateAd(userID, adID int, ad model.AdRequest) error
	// DeleteAd(userID, adID int) error
}

type Photo interface {
	ListPhotos() ([]string, error)
	ListAdPhotos(adID int) ([]string, error)
}

type Tag interface {
	ListTags() ([]string, error)
	ListAdTags(adID int) ([]string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Ad:    NewAdService(repos.Ad, repos.Photo, repos.Tag),
		Photo: NewPhotoService(repos.Photo),
		Tag:   NewTagService(repos.Tag),
	}
}

package service

import (
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
)

type Service struct {
	Ad
	Photo
	Tag
}

type Ad interface {
	GetAd(adID uint32, fields model.AdsOptional) (model.Ad, error)
	ListAds(params []model.AdsSortingParam) ([]model.Ad, error)
	SearchAds(filter model.AdFilter) ([]model.Ad, error)

	CreateAd(userID uint32, ad model.AdRequest) (uint32, error)
	UpdateAd(userID, adID uint32, ad model.AdRequest) (uint32, error)
	DeleteAd(userID, adID uint32) error
}

type Photo interface {
	ListPhotos() ([]string, error)
	ListAdPhotos(adID uint32) ([]string, error)
}

type Tag interface {
	ListTags() ([]string, error)
	ListAdTags(adID uint32) ([]string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Ad:    NewAdService(repos.Ad, repos.Photo, repos.Tag),
		Photo: NewPhotoService(repos.Photo),
		Tag:   NewTagService(repos.Tag),
	}
}

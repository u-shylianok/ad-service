package service

import (
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	"github.com/u-shylianok/ad-service/svc-ads/internal/secure"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
)

type Service struct {
	Auth
	Ad
	Photo
	Tag
}

// ListAds(context.Context, *ListAdsRequest) (*ListAdsResponse, error)
// SearchAds(context.Context, *SearchAdsRequest) (*SearchAdsResponse, error)
// GetAd(context.Context, *GetAdRequest) (*Ad, error)
// CreateAd(context.Context, *CreateAdRequest) (*Ad, error)
// UpdateAd(context.Context, *UpdateAdRequest) (*Ad, error)
// DeleteAd(context.Context, *DeleteAdRequest) (*empty.Empty, error)
// ListPhotos(context.Context, *ListPhotosRequest) (*ListPhotosResponse, error)
// ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
type Auth interface {
	CreateUser(user model.User) (int, error)
	CheckUser(username, password string) (int, error)
	GenerateToken(userID int) (string, int64, error)
	ParseToken(token string) (int, error)
}

type Ad interface {
	ListAds(params []model.AdsSortingParam) ([]model.AdResponse, error)
	SearchAds(filter model.AdFilter) ([]model.AdResponse, error)
	GetAd(adID int, fields model.AdOptionalFieldsParam) (model.AdResponse, error)
	CreateAd(userID int, ad model.AdRequest) (int, error)
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

func NewService(repos *repository.Repository, secure *secure.Secure) *Service {
	return &Service{
		Auth:  NewAuthService(repos.User, secure.Hasher),
		Ad:    NewAdService(repos.Ad, repos.User, repos.Photo, repos.Tag),
		Photo: NewPhotoService(repos.Photo),
		Tag:   NewTagService(repos.Tag),
	}
}

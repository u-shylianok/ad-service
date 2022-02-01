package repository

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
)

type Repository struct {
	Ad
	Photo
	Tag
}

//counterfeiter:generate --fake-name AdMock -o ../testing/mocks/repository/ad.go . Ad
type Ad interface {
	Create(userID uint32, ad model.AdRequest) (uint32, error)
	Get(adID uint32, fields model.AdsOptional) (model.Ad, error)
	List(params []model.AdsSortingParam) ([]model.Ad, error)
	ListWithFilter(filter model.AdFilter) ([]model.Ad, error)
	Update(userID, adID uint32, ad model.AdRequest) error
	Delete(userID, adID uint32) error
}

//counterfeiter:generate --fake-name PhotoMock -o ../testing/mocks/repository/photo.go . Photo
type Photo interface {
	Create(adID uint32, link string) (uint32, error)
	CreateList(adID uint32, photos []string) error
	ListLinks() ([]string, error)
	ListLinksByAd(adID uint32) ([]string, error)
	DeleteAllByAd(adID uint32) error
}

//counterfeiter:generate --fake-name TagMock -o ../testing/mocks/repository/tag.go . Tag
type Tag interface {
	Create(name string) (uint32, error)
	GetByName(name string) (model.Tag, error)
	ListNames() ([]string, error)
	ListNamesByAd(adID uint32) ([]string, error)

	AttachToAd(adID, tagID uint32) error
	DetachFromAd(adID, tagID uint32) error
	DetachAllFromAd(adID uint32) error

	GetIDOrCreateIfNotExists(tagName string) (uint32, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Ad:    NewAdPostrgres(db),
		Photo: NewPhotoPostrgres(db),
		Tag:   NewTagPostrgres(db),
	}
}

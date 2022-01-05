package repository

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/svc-ads/model"
)

type Repository struct {
	Ad
	Photo
	Tag
}

//counterfeiter:generate --fake-name AdMock -o ../testing/mocks/repository/ad.go . Ad
type Ad interface {
	Create(userID int, ad model.AdRequest) (int, error)
	Get(adID uint32, fields model.AdOptionalFieldsParam) (model.Ad, error)
	List(params []model.AdsSortingParam) ([]model.Ad, error)
	ListWithFilter(filter model.AdFilter) ([]model.Ad, error)
	Update(userID, adID int, ad model.AdRequest) error
	Delete(userID, adID int) error
}

//counterfeiter:generate --fake-name PhotoMock -o ../testing/mocks/repository/photo.go . Photo
type Photo interface {
	Create(adID int, link string) (int, error)
	CreateList(adID int, photos []string) error
	ListLinks() ([]string, error)
	ListLinksByAd(adID int) ([]string, error)
	DeleteAllByAd(adID int) error
}

//counterfeiter:generate --fake-name TagMock -o ../testing/mocks/repository/tag.go . Tag
type Tag interface {
	Create(name string) (int, error)
	GetByName(name string) (model.Tag, error)
	ListNames() ([]string, error)
	ListNamesByAd(adID int) ([]string, error)

	AttachToAd(adID int, tagID int) error
	DetachFromAd(adID int, tagID int) error
	DetachAllFromAd(adID int) error

	GetIDOrCreateIfNotExists(tagName string) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Ad:    NewAdPostrgres(db),
		Photo: NewPhotoPostrgres(db),
		Tag:   NewTagPostrgres(db),
	}
}

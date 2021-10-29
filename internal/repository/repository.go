package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type Repository struct {
	User
	Ad
	Photo
	Tag
}

type User interface {
	Create(user model.User) (int, error)
	Get(username, password string) (model.User, error)
	GetByID(id int) (model.User, error)
	ListInIDs(ids []int) ([]model.User, error)
}

type Ad interface {
	Create(userID int, ad model.AdRequest) (int, error)
	Get(adID int, fields model.AdOptionalFieldsParam) (model.Ad, error)
	List(params []model.AdsSortingParam) ([]model.Ad, error)
	ListWithFilter(filter model.AdFilter) ([]model.Ad, error)
	Update(adID int, ad model.AdRequest) error
	Delete(adID int) error
}

type Photo interface {
	Create(adID int, link string) (int, error)
	CreateList(adID int, photos []string) error
	ListLinks() ([]string, error)
	ListLinksByAd(adID int) ([]string, error)
	DeleteAllByAd(adID int) error
}

type Tag interface {
	Create(name string) (int, error)
	GetByName(name string) (model.Tag, error)
	ListNames() ([]string, error)
	ListNamesByAd(adID int) ([]string, error)

	AttachToAd(adID int, tagID int) error
	DetachFromAd(adID int, tagID int) error
	DetachAllFromAd(adID int) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Ad:    NewAdPostrgres(db),
		Photo: NewPhotoPostrgres(db),
		Tag:   NewTagPostrgres(db),
	}
}

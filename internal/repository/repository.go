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
	List(params []model.AdsSortingParam) ([]model.Ad, error)
	Get(adID int, fields model.AdOptionalFieldsParam) (model.Ad, error)
	Update(adID int, ad model.AdRequest) error
	Delete(adID int) error
}

type Photo interface {
	Create(adID int, link string) (int, error)
	CreateList(adID int, photos []string) error
	ListPhotoLinks(adID int) ([]string, error)
	DeleteAllAdPhotos(adID int) error
}

type Tag interface {
	Create(name string) (int, error)
	AttachTagToAd(adID int, tagID int) error
	ListTagNames(adID int) ([]string, error)
	FindByName(name string) (model.Tag, error)
	DetachTagFromAd(adID int, tagID int) error
	DetachAllTagsFromAd(adID int) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Ad:    NewAdPostrgres(db),
		Photo: NewPhotoPostrgres(db),
		Tag:   NewTagPostrgres(db),
	}
}

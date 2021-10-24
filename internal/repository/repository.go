package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type Repository struct {
	Auth
	Ad
	Photo
	Tag
}

type Auth interface {
	Create(user model.User) (int, error)
	Get(username, password string) (model.User, error)
}

type Ad interface {
	Create(userID int, ad model.AdRequest) (int, error)
	List(sortBy, order string) ([]model.AdResponse, error)
	Get(adID int, fields []string) (model.AdResponse, error)
	Update(ad model.AdRequest) error
	Delete(adID int) error
}

type Photo interface {
	Create(adID int, link string) (int, error)
	CreateList(adID int, photos []string) error
}

type Tag interface {
	Create(adID int, name string) (int, error)
	FindByName(name string) (model.Tag, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:  NewAuthPostgres(db),
		Ad:    NewAdPostrgres(db),
		Photo: NewPhotoPostrgres(db),
		Tag:   NewTagPostrgres(db),
	}
}

package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type Repository struct {
	Auth
	Ad
}

type Auth interface {
	Create(user model.User) (int, error)
	Get(username, password string) (model.User, error)
}

type Ad interface {
	Create(ad model.AdRequest) (int, error)
	List(sortBy, order string) ([]model.AdResponse, error)
	Get(adId int, fields []string) (model.AdResponse, error)
	Update(ad model.AdRequest) error
	Delete(adId int) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		Ad:   NewAdPostrgres(db),
	}
}

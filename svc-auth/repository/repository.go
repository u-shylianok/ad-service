package repository

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/svc-auth/model"
)

type Repository struct {
	User
}

//counterfeiter:generate --fake-name UserMock -o ../testing/mocks/repository/user.go . User
type User interface {
	Create(user model.User) (int, error)
	Get(username string) (model.User, error)
	GetByID(id int) (model.User, error)
	ListInIDs(ids []int) ([]model.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}

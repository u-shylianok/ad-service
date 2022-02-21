package repository

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/svc-auth/domain/model"
)

type Repository struct {
	User
}

//counterfeiter:generate --fake-name UserMock -o ../testing/mocks/repository/user.go . User
type User interface {
	Create(user model.User) (uint32, error)
	Get(username string) (model.User, error)
	GetByID(id uint32) (model.User, error)
	GetIDByUsername(username string) (uint32, error)
	ListInIDs(ids []uint32) ([]model.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}

package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(user model.User) (int, error) {
	var id int
	createUserQuery := "INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(createUserQuery, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) Get(username, password string) (model.User, error) {
	var user model.User
	getUserQuery := "SELECT id, name, username, password FROM users WHERE username=$1 AND password=$2"

	err := r.db.Get(&user, getUserQuery, username, password)

	return user, err
}

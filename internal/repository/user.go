package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user model.User) (int, error) {
	var id int
	createUserQuery := "INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(createUserQuery, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) Get(username string) (model.User, error) {
	var user model.User
	getUserQuery := "SELECT id, name, username, password FROM users WHERE username = $1"

	err := r.db.Get(&user, getUserQuery, username)

	return user, err
}

// Gets only name and username
func (r *UserPostgres) GetByID(id int) (model.User, error) {
	var user model.User
	getUserQuery := "SELECT name, username FROM users WHERE id = $1"

	err := r.db.Get(&user, getUserQuery, id)

	return user, err
}

// Gets user slice with (id, name, username)
func (r *UserPostgres) ListInIDs(ids []int) ([]model.User, error) {
	var users []model.User

	if len(ids) == 0 {
		return users, nil
	}

	stringIDs := make([]string, len(ids))
	for i := range ids {
		stringIDs[i] = fmt.Sprint(ids[i])
	}
	listUsersQuery := fmt.Sprintf("SELECT id, name, username FROM users WHERE id IN (%s)", strings.Join(stringIDs, ","))

	if err := r.db.Select(&users, listUsersQuery); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return users, nil
}

package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type PhotoPostgres struct {
	db *sqlx.DB
}

func NewPhotoPostrgres(db *sqlx.DB) *PhotoPostgres {
	return &PhotoPostgres{db: db}
}

func (r *PhotoPostgres) Create(adID int, link string) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		//logrus.Errorf("[create photo]: error: %s", err.Error())
		return 0, err
	}

	var photoID int
	createPhotoQuery := "INSERT INTO photos (ad_id, link) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRow(createPhotoQuery, adID, link)
	if err := row.Scan(&photoID); err != nil {
		//logrus.Errorf("[create photo]: error: %s", err.Error())
		tx.Rollback()
		return 0, err
	}

	return photoID, tx.Commit()
}

func (r *PhotoPostgres) CreateList(adID int, photos []string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		//logrus.Errorf("[create photos]: error: %s", err.Error())
		return err
	}

	values := []string{}
	args := []interface{}{}

	args = append(args, adID)
	argID := 2
	for _, photo := range photos {
		args = append(args, photo)
		values = append(values, fmt.Sprintf("($1, $%d, FALSE)", argID))
		argID++
	}

	createPhotosQuery := fmt.Sprintf("INSERT INTO photos (ad_id, link) VALUES %s", strings.Join(values, ","))
	if _, err := tx.Exec(createPhotosQuery, args...); err != nil {
		//logrus.Errorf("[create photos] error: %s", err.Error())
		tx.Rollback()
		return err
	}
	return nil
}

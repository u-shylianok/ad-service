package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PhotoPostgres struct {
	db *sqlx.DB
}

func NewPhotoPostrgres(db *sqlx.DB) *PhotoPostgres {
	return &PhotoPostgres{db: db}
}

func (r *PhotoPostgres) Create(adID uint32, link string) (uint32, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		//logrus.Errorf("[create photo]: error: %s", err.Error())
		return 0, err
	}

	var photoID uint32
	createPhotoQuery := "INSERT INTO photos (ad_id, link) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRow(createPhotoQuery, adID, link)
	if err := row.Scan(&photoID); err != nil {
		//logrus.Errorf("[create photo]: error: %s", err.Error())
		if err := tx.Rollback(); err != nil {
			logrus.WithError(err).Error("rollback error")
		}
		return 0, err
	}

	return photoID, tx.Commit()
}

func (r *PhotoPostgres) CreateList(adID uint32, photos []string) error {
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
		values = append(values, fmt.Sprintf("($1, $%d)", argID))
		argID++
	}

	createPhotosQuery := fmt.Sprintf("INSERT INTO photos (ad_id, link) VALUES %s", strings.Join(values, ","))
	if _, err := tx.Exec(createPhotosQuery, args...); err != nil {
		//logrus.Errorf("[create photos] error: %s", err.Error())
		if err := tx.Rollback(); err != nil {
			logrus.WithError(err).Error("rollback error")
		}
		return err
	}
	return tx.Commit()
}

func (r *PhotoPostgres) ListLinks() ([]string, error) {
	var photoLinks []string

	listAdsQuery := "SELECT link FROM photos"
	if err := r.db.Select(&photoLinks, listAdsQuery); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return photoLinks, nil
}

func (r *PhotoPostgres) ListLinksByAd(adID uint32) ([]string, error) {
	var photoLinks []string

	listAdsQuery := "SELECT link FROM photos WHERE ad_id = $1"
	if err := r.db.Select(&photoLinks, listAdsQuery, adID); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return photoLinks, nil
}

func (r *PhotoPostgres) DeleteAllByAd(adID uint32) error {
	deletePhotosQuery := "DELETE FROM photos WHERE ad_id = $1"
	_, err := r.db.Exec(deletePhotosQuery, adID)
	return err
}

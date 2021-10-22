package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (r *AdPostgres) createPhotos(tx *sqlx.Tx, adId int, mainPhoto model.PhotoRequest, otherPhotos []model.PhotoRequest) error {

	values := []string{}
	args := []interface{}{}

	args = append(args, adId, mainPhoto.Link)
	values = append(values, fmt.Sprintf("($1, $2, TRUE)"))

	argId := 3
	for _, photo := range otherPhotos {
		args = append(args, photo.Link)
		values = append(values, fmt.Sprintf("($1, $%d, FALSE)", argId))
		argId++
	}

	createPhotosQuery := fmt.Sprintf("INSERT INTO photos (ad_id, link, is_main) VALUES %s", strings.Join(values, ","))
	_, err := tx.Exec(createPhotosQuery, args...)
	if err != nil {
		logrus.Errorf("[create photos] error: %s", err.Error())
		tx.Rollback()
		return err
	}
	return nil
}

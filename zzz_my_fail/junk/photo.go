package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/blockloop/scan"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

type CreatePhotoInput struct {
	Link   string
	IsMain bool
}

func PhotoToCreateInput(p *model.Photo, isMain bool) *CreatePhotoInput {
	return &CreatePhotoInput{Link: p.Link, IsMain: isMain}
}

func (r *AdPostgres) createPhotos(tx *sql.Tx, adID int, photos []CreatePhotoInput) error {
	values := []string{}
	args := []interface{}{}

	argID := 1
	for _, photo := range photos {
		args = append(args, photo.Link)
		values = append(values, fmt.Sprintf("($%d)", argID))
		argID++
	}

	createPhotosQuery := fmt.Sprintf("INSERT INTO photos (link) VALUES %s RETURNING id", strings.Join(values, ","))
	rows, err := tx.Query(createPhotosQuery, args...)
	if err != nil {
		logrus.Errorf("[Create Ad] create photos error: %s", err.Error())
		tx.Rollback()
		return err
	}

	var photosIDs []int
	err = scan.Rows(&photosIDs, rows)
	if err != nil {
		logrus.Errorf("[Create Ad] scanning created photos error: %s", err.Error())
		tx.Rollback()
		return err
	}

	if len(photos) != len(photosIDs) {
		logrus.Error("[Create Ad] created photos should be equal to input photos.")
		tx.Rollback()
		return err
	}

	values = []string{}
	args = []interface{}{}

	args = append(args, adID)
	argID = 2
	for i, photo := range photos {
		args = append(args, photosIDs[i], photo.IsMain)
		values = append(values, fmt.Sprintf("($1, $%d, $%d)", argID, argID+1))
		argID += 2
	}

	createAdsPhotosQuery := fmt.Sprintf("INSERT INTO ads_photos (ad_id, photo_id, is_main) VALUES %s", strings.Join(values, ","))
	_, err = tx.Exec(createAdsPhotosQuery, args...)
	if err != nil {
		logrus.Errorf("[Create Ad] create AdsPhotos error: %s", err.Error())
		tx.Rollback()
		return err
	}

	return nil
}

func (r *AdPostgres) createPhotos(tx *sqlx.Tx, adID int, mainPhoto model.PhotoRequest, otherPhotos []model.PhotoRequest) error {

	values := []string{}
	args := []interface{}{}

	args = append(args, adID, mainPhoto.Link)
	values = append(values, fmt.Sprintf("($1, $2, TRUE)"))

	argID := 3
	for _, photo := range otherPhotos {
		args = append(args, photo.Link)
		values = append(values, fmt.Sprintf("($1, $%d, FALSE)", argID))
		argID++
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

package repository

import (
	"github.com/u-shylianok/ad-service/internal/model"
)

// AdTx createPhoto rolls back the transaction if something goes wrong
func (tx AdTx) createPhoto(photo model.Photo, adId int, isMain bool) (int, error) {
	var photoId int
	createPhotoQuery := "INSERT INTO photos (link) VALUES ($1) RETURNING id"
	row1 := tx.QueryRow(createPhotoQuery, photo.Link)
	err := row1.Scan(&photoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var adsPhotosId int
	createAdsPhotosQuery := "INSERT INTO ads_photos (ad_id, photo_id, is_main) VALUES ($1, $2, $3) RETURNING ad_id"
	row2 := tx.QueryRow(createAdsPhotosQuery, adId, photoId, isMain)
	err = row2.Scan(&adsPhotosId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return photoId, nil
}

package repository

import (
	"github.com/u-shylianok/ad-service/internal/model"
)

// AdTx createTag rolls back the transaction if something goes wrong
func (tx *AdTx) createTag(tag model.Tag, adId int) (int, error) {
	var tagId int
	getTagQuery := "SELECT id FROM tags WHERE name=$1"
	if err := tx.Get(&tagId, getTagQuery, tag.Name); err != nil {
		tx.Rollback()
		return 0, err
	}
	if tagId == 0 {
		createTagQuery := "INSERT INTO tags (name) VALUES ($1) RETURNING id"
		row1 := tx.QueryRow(createTagQuery, tag.Name)
		err := row1.Scan(&tagId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	var adsTagsId int
	createAdsPhotosQuery := "INSERT INTO ads_tags (ad_id, tag_id) VALUES ($1, $2) RETURNING ad_id"
	row2 := tx.QueryRow(createAdsPhotosQuery, adId, tagId)
	err := row2.Scan(&adsTagsId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return tagId, nil
}

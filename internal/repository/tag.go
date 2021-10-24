package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type TagPostgres struct {
	db *sqlx.DB
}

func NewTagPostrgres(db *sqlx.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) Create(adID int, name string) (int, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		//logrus.Errorf("[create tag]: error: %s", err.Error())
		return 0, err
	}

	var tagID int
	createTagQuery := "INSERT INTO tags (name) VALUES ($1) RETURNING id"
	row := tx.QueryRow(createTagQuery, name)
	if err := row.Scan(&tagID); err != nil {
		//logrus.Errorf("[create tag]: error: %s", err.Error())
		tx.Rollback()
		return 0, err
	}

	createAdsTagQuery := "INSERT INTO ads_tags (ad_id, tag_id) VALUES ($1, $2)"
	if _, err := tx.Exec(createAdsTagQuery, adID, tagID); err != nil {
		// logrus.Errorf("[create adstag]: error: %s", err.Error())
		tx.Rollback()
		return 0, err
	}

	return tagID, tx.Commit()
}

func (r *TagPostgres) FindByName(name string) (model.Tag, error) {
	var tag model.Tag

	getTagQuery := "SELECT id, name FROM tags WHERE name=$1"
	r.db.Get(&tag, getTagQuery, name)

	return tag, nil
}

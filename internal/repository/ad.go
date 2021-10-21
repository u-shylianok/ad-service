package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/ad-service/internal/model"
)

type AdPostgres struct {
	db *sqlx.DB
}

type AdTx struct {
	*sqlx.Tx
}

func NewAdPostrgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{db: db}
}

func (r *AdPostgres) Create(ad model.Ad) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	adTx := AdTx{tx}

	var adId int
	createAdQuery := "INSERT INTO ads (name, date, price, description) VALUES ($1, $2, $3, $4) RETURNING id"
	row := adTx.QueryRow(createAdQuery, ad.Name, ad.Date, ad.Price, ad.Description)
	err = row.Scan(&adId)
	if err != nil {
		adTx.Rollback()
		return 0, err
	}

	if _, err := adTx.createPhoto(ad.MainPhoto, adId, true); err != nil {
		return 0, err
	}
	for _, photo := range *ad.OtherPhotos {
		if _, err := adTx.createPhoto(photo, adId, false); err != nil {
			return 0, err
		}
	}
	for _, tag := range *ad.Tags {
		if _, err := adTx.createTag(tag, adId); err != nil {
			return 0, err
		}
	}

	return adId, tx.Commit()
}

func (r *AdPostgres) List(filter model.Ad) ([]model.Ad, error) {
	var ads []model.Ad

	listAdsQuery := "SELECT id, name, price FROM ads"
	if err := r.db.Select(&ads, listAdsQuery); err != nil {
		return nil, err
	}

	return ads, nil
}

func (r *AdPostgres) Get(adId int) (model.Ad, error) {
	var ad model.Ad

	getAdQuery := "SELECT id, name, date, price FROM ads WHERE id = $1"
	if err := r.db.Get(&ad, getAdQuery, adId); err != nil {
		return ad, err
	}
	return ad, nil
}

func (r *AdPostgres) Update(ad model.Ad) error {
	return nil
}

func (r *AdPostgres) Delete(adId int) error {
	return nil
}

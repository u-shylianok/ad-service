package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostrgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{db: db}
}

func (r *AdPostgres) Create(userId int, ad model.AdRequest) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		logrus.Errorf("[create ad]: error: %s", err.Error())
		return 0, err
	}

	currentDate := time.Now().Format("2006-01-02")

	var adId int
	createAdQuery := "INSERT INTO ads (user_id, name, date, price, description) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := tx.QueryRow(createAdQuery, userId, ad.Name, currentDate, ad.Price, ad.Description)
	if err := row.Scan(&adId); err != nil {
		logrus.Errorf("[create ad]: error: %s", err.Error())
		tx.Rollback()
		return 0, err
	}

	// (рус) Подготавливаем всё для создания фото (кажется, где-то я перемудрил или ступил)
	links := make([]string, len(*ad.OtherPhotos)+1)
	isMains := make([]bool, len(*ad.OtherPhotos)+1)
	links[0] = ad.MainPhoto.Link
	isMains[0] = true
	for i := 0; i < len(*ad.OtherPhotos); i++ {
		links[i+1] = (*ad.OtherPhotos)[i].Link
	}

	if err := r.createPhotos(tx, adId, links, isMains); err != nil {
		return 0, err
	}

	for _, tag := range *ad.Tags {
		if _, err := adTx.createTag(tag, adId); err != nil {
			return 0, err
		}
	}

	return adId, tx.Commit()
}

func (r *AdPostgres) List(sortBy, order string) ([]model.Ad, error) {
	var ads []model.Ad

	var addQuerySortBy, addQueryOrder string
	if sortBy == "price" || sortBy == "date" {
		addQuerySortBy = " ORDER BY ads." + sortBy

		if order == "dsc" {
			addQueryOrder = " DESC"
		}
	}

	query := fmt.Sprintf("SELECT * FROM ads%s%s", addQuerySortBy, addQueryOrder)
	logrus.Info("Все ок 1")
	if err := r.db.Select(&ads, query); err != nil {
		return nil, err
	}
	logrus.Info("Все ок 2")

	query2 := "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND ads_photos.is_main"
	for i := 0; i < len(ads); i++ {

		var photo model.Photo

		if err := r.db.Get(&photo, query2, ads[i].Id); err != nil {
			return nil, err
		}
		ads[i].MainPhoto = photo
	}
	logrus.Info("Все ок 3")

	return ads, nil
}

func (r *AdPostgres) Get(adId int, fields []string) (model.Ad, error) {
	var ad model.Ad

	fieldsToQueries := make(map[string]string)
	for _, a := range fields {
		fieldsToQueries[a] = a
	}

	if dsc, ok := fieldsToQueries["description"]; ok {
		fieldsToQueries["description"] = " , " + dsc
	}

	getAdQuery := fmt.Sprintf("SELECT id, name, date, price%s FROM ads WHERE id = $1", fieldsToQueries["description"])
	if err := r.db.Get(&ad, getAdQuery, adId); err != nil {
		return ad, err
	}

	var mainPhoto model.Photo
	getMainPhotoQuery := "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND NOT ads_photos.is_main"
	if err := r.db.Get(&mainPhoto, getMainPhotoQuery, adId); err != nil {
		return ad, err
	}
	ad.MainPhoto = mainPhoto

	var query2 string
	if _, ok := fieldsToQueries["photos"]; ok {
		query2 = "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND NOT ads_photos.is_main"

		var otherPhotos []model.Photo
		if err := r.db.Select(&otherPhotos, query2, ad.Id); err != nil {
			return ad, err
		}
		ad.OtherPhotos = &otherPhotos
	}

	return ad, nil
}

func (r *AdPostgres) Update(ad model.Ad) error {
	return nil
}

func (r *AdPostgres) Delete(adId int) error {
	return nil
}

package repository

import (
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

func (r *AdPostgres) Create(userID int, ad model.AdRequest) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		//logrus.Errorf("[create ad]: error: %s", err.Error())
		return 0, err
	}

	var adID int
	createAdQuery := "INSERT INTO ads (user_id, name, price, photo, description) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := tx.QueryRow(createAdQuery, userID, ad.Name, ad.Price, ad.MainPhoto, ad.Description)
	if err := row.Scan(&adID); err != nil {
		//logrus.Errorf("[create ad]: error: %s", err.Error())
		tx.Rollback()
		return 0, err
	}

	return adID, tx.Commit()
}

func (r *AdPostgres) List(sortBy, order string) ([]model.Ad, error) {
	var ads []model.Ad

	// var addQuerySortBy, addQueryOrder string
	// if sortBy == "price" || sortBy == "date" {
	// 	addQuerySortBy = " ORDER BY ads." + sortBy

	// 	if order == "dsc" {
	// 		addQueryOrder = " DESC"
	// 	}
	// }

	// query := fmt.Sprintf("SELECT * FROM ads%s%s", addQuerySortBy, addQueryOrder)
	// logrus.Info("Все ок 1")
	// if err := r.db.Select(&ads, query); err != nil {
	// 	return nil, err
	// }
	// logrus.Info("Все ок 2")

	// query2 := "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND ads_photos.is_main"
	// for i := 0; i < len(ads); i++ {

	// 	var photo model.Photo

	// 	if err := r.db.Get(&photo, query2, ads[i].ID); err != nil {
	// 		return nil, err
	// 	}
	// 	ads[i].MainPhoto = photo
	// }
	// logrus.Info("Все ок 3")

	listAdsQuery := "SELECT * FROM ads"
	if err := r.db.Select(&ads, listAdsQuery); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return ads, nil
}

func (r *AdPostgres) Get(adID int, fields []string) (model.AdResponse, error) {
	var ad model.AdResponse

	// fieldsToQueries := make(map[string]string)
	// for _, a := range fields {
	// 	fieldsToQueries[a] = a
	// }

	// if dsc, ok := fieldsToQueries["description"]; ok {
	// 	fieldsToQueries["description"] = " , " + dsc
	// }

	// getAdQuery := fmt.Sprintf("SELECT id, name, date, price%s FROM ads WHERE id = $1", fieldsToQueries["description"])
	// if err := r.db.Get(&ad, getAdQuery, adID); err != nil {
	// 	return ad, err
	// }

	// var mainPhoto model.Photo
	// getMainPhotoQuery := "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND NOT ads_photos.is_main"
	// if err := r.db.Get(&mainPhoto, getMainPhotoQuery, adID); err != nil {
	// 	return ad, err
	// }
	// ad.MainPhoto = mainPhoto

	// var query2 string
	// if _, ok := fieldsToQueries["photos"]; ok {
	// 	query2 = "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND NOT ads_photos.is_main"

	// 	var otherPhotos []model.Photo
	// 	if err := r.db.Select(&otherPhotos, query2, ad.ID); err != nil {
	// 		return ad, err
	// 	}
	// 	ad.OtherPhotos = &otherPhotos
	// }

	return ad, nil
}

func (r *AdPostgres) Update(ad model.AdRequest) error {
	return nil
}

func (r *AdPostgres) Delete(adID int) error {
	return nil
}

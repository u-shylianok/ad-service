package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	"github.com/u-shylianok/ad-service/svc-ads/repository/postgres/query"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostrgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{db: db}
}

func (r *AdPostgres) Get(adID uint32, fields model.AdsOptional) (model.Ad, error) {
	var ad model.Ad

	var fieldsQuery string
	if fields.Description {
		fieldsQuery = ", description"
	}

	getAdQuery := fmt.Sprintf("SELECT id, user_id, name, date, price, photo %s FROM ads WHERE id = $1", fieldsQuery)
	if err := r.db.Get(&ad, getAdQuery, adID); err != nil {
		return ad, err
	}

	return ad, nil
}

func (r *AdPostgres) List(params []model.AdsSortingParam) ([]model.Ad, error) {
	var ads []model.Ad

	var orderbyQuery string
	if params != nil {
		queryPart := make([]string, len(params))
		for i, param := range params {
			if param.IsDesc {
				queryPart[i] = fmt.Sprintf("%s DESC", param.Field)
			} else {
				queryPart[i] = fmt.Sprintf("%s ASC", param.Field)
			}
		}
		orderbyQuery = fmt.Sprintf("ORDER BY %s", strings.Join(queryPart, ","))
	}

	listAdsQuery := fmt.Sprintf("SELECT * FROM ads %s", orderbyQuery)
	if err := r.db.Select(&ads, listAdsQuery); err != nil {
		//logrus.Error(err)
		return nil, err
	}

	return ads, nil
}

func (r *AdPostgres) ListWithFilter(filter model.AdFilter) ([]model.Ad, error) {
	var ads []model.Ad

	listAdsWithFilterQuery, args := query.BuildAdFilterQuery(filter)
	logrus.WithFields(logrus.Fields{
		"query": listAdsWithFilterQuery,
		"args":  args}).Debug("building query successfully")

	if err := r.db.Select(&ads, listAdsWithFilterQuery, args...); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return ads, nil
}

func (r *AdPostgres) Create(userID uint32, ad model.AdRequest) (uint32, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		//logrus.Errorf("[create ad]: error: %s", err.Error())
		return 0, err
	}

	var adID uint32
	createAdQuery := "INSERT INTO ads (user_id, name, price, photo, description) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := tx.QueryRow(createAdQuery, userID, ad.Name, ad.Price, ad.MainPhoto, ad.Description)
	if err := row.Scan(&adID); err != nil {
		//logrus.Errorf("[create ad]: error: %s", err.Error())
		if err := tx.Rollback(); err != nil {
			logrus.WithError(err).Error("rollback error")
		}
		return 0, err
	}

	return adID, tx.Commit()
}

func (r *AdPostgres) Update(userID, adID uint32, ad model.AdRequest) (uint32, error) {
	updateAdQuery := "UPDATE ads SET name = $1, price = $2, description = $3, photo = $4 WHERE user_id = $5 AND id = $6 RETURNING id"

	var id uint32
	row := r.db.QueryRow(updateAdQuery, ad.Name, ad.Price, ad.Description, ad.MainPhoto, userID, adID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AdPostgres) Delete(userID, adID uint32) error {

	deleteAdQuery := "DELETE FROM ads WHERE user_id = $1 AND id = $2"
	_, err := r.db.Exec(deleteAdQuery, userID, adID)

	return err
}

// func (r *AdPostgres) CheckUser(userID, adID int) error {

// 	checkUserQuery := "SELECT id FROM ads WHERE ads.user_id = $1 AND ads.id = $2"

// 	var id int
// 	if err := r.db.Get(&id, checkUserQuery, userID, adID); err != nil {
// 		return err
// 	}

// 	if id != adID { // just for id usage
// 		return fmt.Errorf("unexpected error")
// 	}

// 	return nil
// }

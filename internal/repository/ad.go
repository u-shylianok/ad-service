package repository

import (
	"fmt"
	"strings"

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

func (r *AdPostgres) Get(adID int, fields model.AdOptionalFieldsParam) (model.Ad, error) {
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

	arg := 1
	// var isWhereExist bool
	var filterQuery string
	{
		var sb strings.Builder

		if filter.Username != "" {
			sb.WriteString(fmt.Sprintf("INNER JOIN users ON ads.user_id = users.id AND username = $%d ", arg))
			arg++
		}
		if len(filter.Tags) > 0 {
			sb.WriteString(fmt.Sprintf("INNER JOIN ads_tags ON ads.id = ads_tags.ad_id INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name = $%d", arg))
			arg++
		}
		if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
			sb.WriteString("WHERE ")
			if !filter.StartDate.IsZero() {
				sb.WriteString(fmt.Sprintf("ads.date >= $%d", arg))
				arg++
			}
			if !filter.EndDate.IsZero() {
				sb.WriteString(fmt.Sprintf("ads.date <= $%d", arg))
				arg++
			}
		}
		filterQuery = sb.String()
	}
	// if params != nil {
	// 	queryPart := make([]string, len(params))
	// 	for i, param := range params {
	// 		if param.IsDesc {
	// 			queryPart[i] = fmt.Sprintf("%s DESC", param.Field)
	// 		} else {
	// 			queryPart[i] = fmt.Sprintf("%s ASC", param.Field)
	// 		}
	// 	}
	// }

	// SELECT ads.* FROM ads
	//INNER JOIN users ON ads.user_id = users.id AND username='test'
	//INNER JOIN ads_tags ON ads.id = ads_tags.ad_id
	//INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name = 'ВАЖНОЕ'
	//WHERE ads.date > '2021-10-9' AND ads.date < '2021-10-25'
	logrus.Info(filterQuery)
	listAdsWithFilterQuery := fmt.Sprint(filterQuery)
	if err := r.db.Select(&ads, listAdsWithFilterQuery); err != nil {
		//logrus.Error(err)
		return nil, err

	}
	return ads, nil
}

func (r *AdPostgres) Update(adID int, ad model.AdRequest) error {

	updateAdQuery := "UPDATE ads SET name = $1, price = $2, description = $3, photo = $4 WHERE id = $5"
	_, err := r.db.Exec(updateAdQuery, ad.Name, ad.Price, ad.Description, ad.MainPhoto, adID)

	return err
}

func (r *AdPostgres) Delete(adID int) error {

	deleteAdQuery := "DELETE FROM ads WHERE id = $1"
	_, err := r.db.Exec(deleteAdQuery, adID)

	return err
}

package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (r *AdPostgres) createPhotos(tx *sqlx.Tx, adId int, links []string, isMains []bool) error {
	if len(links) != len(isMains) {
		err := errors.New("(un)expected error")
		logrus.Errorf("[create photos] error: %s", err.Error())
		tx.Rollback()
		return err
	}

	values := []string{}
	args := []interface{}{}

	args = append(args, adId)
	argId := 2
	for i := 0; i < len(links); i++ {
		args = append(args, links[i], isMains[i])
		values = append(values, fmt.Sprintf("($1, $%d, $%d)", argId, argId+1))
		argId += 2
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

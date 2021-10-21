package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/blockloop/scan"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (r *AdPostgres) listTagsInNames(names []string) (model.Tags, error) {
	var tags model.Tags

	listTagsQuery := fmt.Sprintf("SELECT id, name FROM tags WHERE name IN (%s)", strings.Join(names, ","))
	r.db.Select(&tags, listTagsQuery)

	return tags, nil
}

func (r *AdPostgres) createTags(tx *sql.Tx, adId int, tags model.Tags) error {
	existTags, err := r.listTagsInNames(tags.ListNames())
	if err != nil {
		logrus.Errorf("[create Ad] find existing tags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	tagsToCreate := tags.GetUnexistTags(existTags)

	values := []string{}
	args := []interface{}{}

	argId := 1
	for _, tag := range tagsToCreate {

		args = append(args, tag.Name)
		values = append(values, fmt.Sprintf("($%d)", argId))
		argId++
	}

	createTagsQuery := fmt.Sprintf("INSERT INTO tags (name) VALUES %s RETURNING id", strings.Join(values, ","))
	rows, err := tx.Query(createTagsQuery, args...)
	if err != nil {
		logrus.Errorf("[Create Ad] create tags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	var createdTagIds []int
	err = scan.Rows(&createdTagIds, rows)
	if err != nil {
		logrus.Errorf("[Create Ad] scanning created tags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	tagIds := append(createdTagIds, existTags.ListIds()...)
	if len(tags) != len(tagIds) {
		logrus.Error("[Create Ad] created and existed tags should be equal to input tags.")
		tx.Rollback()
		return err
	}

	values = []string{}
	args = []interface{}{}

	args = append(args, adId)
	argId = 2
	for _, tagId := range tagIds {
		args = append(args, tagId)
		values = append(values, fmt.Sprintf("($1, $%d)", argId))
		argId++
	}

	createAdsTagsQuery := fmt.Sprintf("INSERT INTO ads_tags (ad_id, tag_id) VALUES %s", strings.Join(values, ","))
	_, err = tx.Exec(createAdsTagsQuery, args...)
	if err != nil {
		logrus.Errorf("[Create Ad] create AdsTags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	return nil
}

func (r *AdPostgres) createAdsTags(tx *sql.Tx) error {

	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/blockloop/scan"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (r *AdPostgres) listTagsFromRequest(tags []model.TagRequest) ([]model.Tag, error) {
	var result []model.Tag
	names := model.TagsWithNames(tags).ToListNames()

	listTagsQuery := fmt.Sprintf("SELECT id, name FROM tags WHERE name IN (%s)", strings.Join(names, ","))
	r.db.Select(&tags, listTagsQuery)

	return result, nil
}

func (r *AdPostgres) createTags(tx *sqlx.Tx, tags []model.TagRequest) ([]int, error) {
	existTags, err := r.listTagsFromRequest(tags)
	if err != nil {
		logrus.Errorf("[create tag] find existing tags error: %s", err.Error())
		tx.Rollback()
		return nil, err
	}

	tagsToCreate := model.GetUnexistTags(tags, existTags)

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
		logrus.Errorf("[create tag] create tags error: %s", err.Error())
		tx.Rollback()
		return nil, err
	}

	var createdTagIds []int
	err = scan.Rows(&createdTagIds, rows)
	if err != nil {
		logrus.Errorf("[create tag] scanning created tags error: %s", err.Error())
		tx.Rollback()
		return nil, err
	}

	tagIds := append(createdTagIds, existTags.ListIds()...)

	return tagIds, nil
}

func (r *AdPostgres) createAdsTags(tx *sql.Tx, adId int, tagIds []int) error {
	values := []string{}
	args := []interface{}{}

	args = append(args, adId)
	argId := 2
	for _, tagId := range tagIds {
		args = append(args, tagId)
		values = append(values, fmt.Sprintf("($1, $%d)", argId))
		argId++
	}

	createAdsTagsQuery := fmt.Sprintf("INSERT INTO ads_tags (ad_id, tag_id) VALUES %s", strings.Join(values, ","))
	_, err := tx.Exec(createAdsTagsQuery, args...)
	if err != nil {
		logrus.Errorf("[Create Ad] create AdsTags error: %s", err.Error())
		tx.Rollback()
		return err
	}
	return nil
}

func listTagNamesToCreate(alltags []model.TagRequest, exist []model.Tag) []string {
	if len(exist) == 0 {
		return alltags
	}
	if len(alltags) == len(exist) {
		return nil
	}

	var result []model.TagRequest

	for _, tag1 := range alltags {
		for _, tag2 := range exist {
			if tag1.Name == tag2.Name {
				result = append(result, model.TagRequest{Name: tag1.Name})
			}
		}
	}
	return result
}

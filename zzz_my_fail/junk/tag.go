package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/blockloop/scan"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

type Tags []model.Tag

func (r *AdPostgres) listTagsInNames(names []string) (Tags, error) {
	var tags Tags

	listTagsQuery := fmt.Sprintf("SELECT id, name FROM tags WHERE name IN (%s)", strings.Join(names, ","))
	r.db.Select(&tags, listTagsQuery)

	return tags, nil
}

func (r *AdPostgres) createTags(tx *sql.Tx, adId int, tags Tags) error {
	existTags, err := r.listTagsInNames(tags.NameList())
	if err != nil {
		logrus.Errorf("[Create Ad] find existing tags error: %s", err.Error())
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

	tagIds := append(createdTagIds, existTags.IdList()...)
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

func prepareTagsInsertQuery() (string, []interface{}) {

	return "", nil
}

func (t Tags) NameList() []string {
	list := make([]string, len(t))
	for _, tag := range t {
		list = append(list, tag.Name)
	}
	return list
}

func (t Tags) IdList() []int {
	list := make([]int, len(t))
	for _, tag := range t {
		list = append(list, tag.Id)
	}
	return list
}

func (t Tags) GetUnexistTags(existed Tags) Tags {
	if len(existed) == 0 {
		return t
	}
	if len(t) == len(existed) {
		return nil
	}

	var result Tags

	for _, tag1 := range t {
		for _, tag2 := range existed {
			if tag1.Id == tag2.Id && tag1.Name == tag2.Name {
				result = append(result, model.Tag{Id: tag1.Id, Name: tag1.Name})
			}
		}
	}

	return result
}

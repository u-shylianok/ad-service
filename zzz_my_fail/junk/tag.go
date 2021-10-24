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

type Tags []model.Tag

func (r *AdPostgres) listTagsInNames(names []string) (Tags, error) {
	var tags Tags

	listTagsQuery := fmt.Sprintf("SELECT id, name FROM tags WHERE name IN (%s)", strings.Join(names, ","))
	r.db.Select(&tags, listTagsQuery)

	return tags, nil
}

func (r *AdPostgres) createTags(tx *sql.Tx, adID int, tags Tags) error {
	existTags, err := r.listTagsInNames(tags.NameList())
	if err != nil {
		logrus.Errorf("[Create Ad] find existing tags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	tagsToCreate := tags.GetUnexistTags(existTags)

	values := []string{}
	args := []interface{}{}

	argID := 1
	for _, tag := range tagsToCreate {

		args = append(args, tag.Name)
		values = append(values, fmt.Sprintf("($%d)", argID))
		argID++
	}

	createTagsQuery := fmt.Sprintf("INSERT INTO tags (name) VALUES %s RETURNING id", strings.Join(values, ","))
	rows, err := tx.Query(createTagsQuery, args...)
	if err != nil {
		logrus.Errorf("[Create Ad] create tags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	var createdTagIDs []int
	err = scan.Rows(&createdTagIDs, rows)
	if err != nil {
		logrus.Errorf("[Create Ad] scanning created tags error: %s", err.Error())
		tx.Rollback()
		return err
	}

	tagIDs := append(createdTagIDs, existTags.IDList()...)
	if len(tags) != len(tagIDs) {
		logrus.Error("[Create Ad] created and existed tags should be equal to input tags.")
		tx.Rollback()
		return err
	}

	values = []string{}
	args = []interface{}{}

	args = append(args, adID)
	argID = 2
	for _, tagID := range tagIDs {
		args = append(args, tagID)
		values = append(values, fmt.Sprintf("($1, $%d)", argID))
		argID++
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

func (t Tags) IDList() []int {
	list := make([]int, len(t))
	for _, tag := range t {
		list = append(list, tag.ID)
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
			if tag1.ID == tag2.ID && tag1.Name == tag2.Name {
				result = append(result, model.Tag{ID: tag1.ID, Name: tag1.Name})
			}
		}
	}

	return result
}

/// PART 2

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

	argID := 1
	for _, tag := range tagsToCreate {

		args = append(args, tag.Name)
		values = append(values, fmt.Sprintf("($%d)", argID))
		argID++
	}

	createTagsQuery := fmt.Sprintf("INSERT INTO tags (name) VALUES %s RETURNING id", strings.Join(values, ","))
	rows, err := tx.Query(createTagsQuery, args...)
	if err != nil {
		logrus.Errorf("[create tag] create tags error: %s", err.Error())
		tx.Rollback()
		return nil, err
	}

	var createdTagIDs []int
	err = scan.Rows(&createdTagIDs, rows)
	if err != nil {
		logrus.Errorf("[create tag] scanning created tags error: %s", err.Error())
		tx.Rollback()
		return nil, err
	}

	tagIDs := append(createdTagIDs, existTags.ListIDs()...)

	return tagIDs, nil
}

func (r *AdPostgres) createAdsTags(tx *sql.Tx, adID int, tagIDs []int) error {
	values := []string{}
	args := []interface{}{}

	args = append(args, adID)
	argID := 2
	for _, tagID := range tagIDs {
		args = append(args, tagID)
		values = append(values, fmt.Sprintf("($1, $%d)", argID))
		argID++
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

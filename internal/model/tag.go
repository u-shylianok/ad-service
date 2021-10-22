package model

import "errors"

// TagRequest - структура, получаемая в запросе
type TagRequest struct {
	Name string `json:"name"`
}

// Validate - валидация структуры
func (r TagRequest) Validate() error {
	if r.Name == "" {
		return errors.New("tag name should not be empty")
	}
	return nil
}

// TagRequest - структура, возвращаемая в ответе
type TagResponse struct {
	Name string `json:"name"`
}

// Tag - структура, используемая в запросах к БД
type Tag struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func TagsToListIds(tags []Tag) []int {
	result := make([]int, len(tags))
	for i, tag := range tags {
		result[i] = tag.Id
	}
	return result
}

func ToListTagNames(tags []TagRequest) []string {
	result := make([]string, len(tags))
	for i, tag := range tags {
		result[i] = tag.Name
	}
	return result
}

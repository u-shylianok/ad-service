package model

import "errors"

type TagRequest struct {
	Name string `json:"name"`
}

func (r TagRequest) Validate() error {
	if r.Name == "" {
		return errors.New("tag name should not be empty")
	}
	return nil
}

type Tag struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type TagResponse struct {
	Name string `json:"name"`
}

type Tags []Tag

func (t Tags) ListNames() []string {
	list := make([]string, len(t))
	for _, tag := range t {
		list = append(list, tag.Name)
	}
	return list
}

func (t Tags) ListIds() []int {
	list := make([]int, len(t))
	for _, tag := range t {
		list = append(list, tag.Id)
	}
	return list
}

func (t Tags) GetUnexistTags(exist Tags) Tags {
	if len(exist) == 0 {
		return t
	}
	if len(t) == len(exist) {
		return nil
	}

	var result Tags

	for _, tag1 := range t {
		for _, tag2 := range exist {
			if tag1.Id == tag2.Id && tag1.Name == tag2.Name {
				result = append(result, Tag{Id: tag1.Id, Name: tag1.Name})
			}
		}
	}

	return result
}

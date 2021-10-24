package model

type Tag struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func TagsToListIDs(tags []Tag) []int {
	result := make([]int, len(tags))
	for i, tag := range tags {
		result[i] = tag.ID
	}
	return result
}

func TagsToListNames(tags []Tag) []string {
	result := make([]string, len(tags))
	for i, tag := range tags {
		result[i] = tag.Name
	}
	return result
}

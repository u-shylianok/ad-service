package model

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagFilter struct {
	AdId int
}

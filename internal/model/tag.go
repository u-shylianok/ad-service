package model

type Tag struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

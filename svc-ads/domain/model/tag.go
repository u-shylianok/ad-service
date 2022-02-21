package model

type Tag struct {
	ID   uint32 `db:"id"`
	Name string `db:"name"`
}

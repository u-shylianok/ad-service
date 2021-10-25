package model

type Photo struct {
	ID   int    `db:"id"`
	AdID int    `db:"ad_id"`
	Link string `db:"link"`
}

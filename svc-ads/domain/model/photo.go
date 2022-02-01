package model

type Photo struct {
	ID   uint32 `db:"id"`
	AdID uint32 `db:"ad_id"`
	Link string `db:"link"`
}

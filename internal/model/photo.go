package model

import "errors"

type PhotoRequest struct {
	Link string `json:"link"`
}

func (r PhotoRequest) Validate() error {
	if r.Link == "" {
		return errors.New("photo link should not be empty")
	}
	return nil
}

type Photo struct {
	Id     int    `db:"id"`
	AdId   int    `db:"ad_id"`
	Link   string `db:"link"`
	IsMain bool   `db:"is_main"`
}

type PhotoResponse struct {
	Link string `json:"link"`
}

package model

import (
	"errors"
	"time"
)

type AdRequest struct {
	Name        string          `json:"name"`
	Price       int             `json:"price"`
	Description string          `json:"description"`
	MainPhoto   PhotoRequest    `json:"main_photo"`
	OtherPhotos *[]PhotoRequest `json:"other_photos"`
	Tags        *[]TagRequest   `json:"tags"`
}

func (r AdRequest) Validate() error {
	if r.Name == "" {
		return errors.New("ad name should not be empty")
	}
	if r.Price < 0 {
		return errors.New("ad price should not be negative")
	}
	if err := r.MainPhoto.Validate(); err != nil {
		return err
	}
	for _, photo := range *r.OtherPhotos {
		if err := photo.Validate(); err != nil {
			return err
		}
	}
	for _, tag := range *r.Tags {
		if err := tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Ad struct {
	Id          int       `db:"id"`
	UserId      int       `db:"user_id"`
	Name        string    `db:"name"`
	Date        time.Time `db:"date"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
}

type AdResponse struct {
	Id          int              `json:"id"`
	User        User             `json:"user"`
	Name        string           `json:"name"`
	Date        time.Time        `json:"date"`
	Price       int              `json:"price"`
	Description string           `json:"description"`
	MainPhoto   PhotoResponse    `json:"main_photo"`
	OtherPhotos *[]PhotoResponse `json:"other_photos"`
	Tags        *[]TagResponse   `json:"tags"`
}

package model

import (
	"errors"
	"time"
	"unicode/utf8"
)

type AdRequest struct {
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	MainPhoto   string    `json:"main_photo"`
	OtherPhotos *[]string `json:"other_photos"`
	Tags        *[]string `json:"tags"`
}

func (r AdRequest) Validate() error {

	// busines rules START
	if utf8.RuneCountInString(r.Name) > 200 {
		return errors.New("name should be no more than 200 symbols")
	}

	if utf8.RuneCountInString(r.Description) > 1000 {
		return errors.New("description should be no more than 1000 symbols")
	}

	if r.MainPhoto == "" {
		return errors.New("main photo must exist")
	}

	if r.OtherPhotos != nil && len(*r.OtherPhotos) > 2 {
		return errors.New("should be no more than 2 other photos")
	}
	// busines rules END

	// my rules START
	if r.Name == "" {
		return errors.New("name should not be empty")
	}
	if r.Price < 0 {
		return errors.New("price should not be negative")
	}
	for _, photo := range *r.OtherPhotos {
		if photo == "" {
			return errors.New("photo link should not be empty")
		}
	}
	if r.Tags != nil {
		for _, tag := range *r.Tags {
			if tag == "" {
				return errors.New("tag name should not be empty")
			}
		}
	}
	// my rules END

	return nil
}

type Ad struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	Name        string    `db:"name"`
	Date        time.Time `db:"date"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
}

type AdResponse struct {
	ID          int       `json:"id"`
	User        User      `json:"user"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	MainPhoto   string    `json:"main_photo"`
	OtherPhotos *[]string `json:"other_photos"`
	Tags        *[]string `json:"tags"`
}

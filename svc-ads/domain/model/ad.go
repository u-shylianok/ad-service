package model

import (
	"time"
)

type Ad struct {
	ID        uint32    `db:"id"`
	UserID    uint32    `db:"user_id"`
	Name      string    `db:"name"`
	Date      time.Time `db:"date"`
	Price     int       `db:"price"`
	MainPhoto string    `db:"photo"`

	Description *string   `db:"description"`
	OtherPhotos *[]string `db:"other_photos"`
	Tags        *[]string `db:"tags"`
}

type AdsOptional struct {
	Description bool
	Photos      bool
	Tags        bool
}

type AdsSortingParam struct {
	Field  string
	IsDesc bool
}

type AdFilter struct {
	UserID    uint32
	StartDate time.Time
	EndDate   time.Time
	Tags      []string
}

type AdRequest struct {
	Name        string
	Price       int
	Description string
	MainPhoto   string
	OtherPhotos *[]string
	Tags        *[]string
}

// type AdResponse struct {
// 	ID          uint32       `json:"id"`
// 	User        UserResponse `json:"user"`
// 	Name        string       `json:"name"`
// 	Date        time.Time    `json:"date"`
// 	Price       int          `json:"price"`
// 	Description string       `json:"description,omitempty"`
// 	MainPhoto   string       `json:"main_photo"`
// 	OtherPhotos *[]string    `json:"other_photos,omitempty"`
// 	Tags        *[]string    `json:"tags,omitempty"`
// }

// func (m Ad) ToResponse(user User, photos *[]string, tags *[]string) AdResponse {
// 	return AdResponse{
// 		ID:          m.ID,
// 		User:        user.ToResponse(),
// 		Name:        m.Name,
// 		Date:        m.Date,
// 		Price:       m.Price,
// 		Description: m.Description,
// 		MainPhoto:   m.MainPhoto,
// 		OtherPhotos: photos,
// 		Tags:        tags,
// 	}
// }

// func ConvertAdsToResponse(ads []Ad, usersMap map[uint32]User) []AdResponse {
// 	result := make([]AdResponse, len(ads))

// 	for i, ad := range ads {
// 		result[i] = ad.ToResponse(usersMap[ad.UserID], nil, nil)
// 	}
// 	return result
// }

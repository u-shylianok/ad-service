package model

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
)

type AdRequest struct {
	Name        string    `json:"name" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	Description string    `json:"description" binding:"required"`
	MainPhoto   string    `json:"main_photo" binding:"required"`
	OtherPhotos *[]string `json:"other_photos"`
	Tags        *[]string `json:"tags"`
}

func (r AdRequest) Validate() error {

	// my rules START
	if r.Name == "" {
		return fmt.Errorf("name should not be empty")
	}
	if r.Description == "" {
		return fmt.Errorf("description should not be empty")
	}
	if r.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}
	if r.OtherPhotos != nil {
		for _, photo := range *r.OtherPhotos {
			if photo == "" {
				return fmt.Errorf("photo link should not be empty")
			}
		}
	}
	if r.Tags != nil {
		for _, tag := range *r.Tags {
			if tag == "" {
				return fmt.Errorf("tag name should not be empty")
			}
		}
	}
	// my rules END

	// busines rules START
	if utf8.RuneCountInString(r.Name) > 200 {
		return fmt.Errorf("name should be no more than 200 symbols")
	}
	if utf8.RuneCountInString(r.Description) > 1000 {
		return fmt.Errorf("description should be no more than 1000 symbols")
	}
	if r.MainPhoto == "" {
		return fmt.Errorf("main photo must exist")
	}
	if r.OtherPhotos != nil && len(*r.OtherPhotos) > 2 {
		return fmt.Errorf("should be no more than 2 other photos")
	}
	// busines rules END
	return nil
}

type Ad struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	Name        string    `db:"name"`
	Date        time.Time `db:"date"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
	MainPhoto   string    `db:"photo"`
}

type AdResponse struct {
	ID          int          `json:"id"`
	User        UserResponse `json:"user"`
	Name        string       `json:"name"`
	Date        time.Time    `json:"date"`
	Price       int          `json:"price"`
	Description string       `json:"description,omitempty"`
	MainPhoto   string       `json:"main_photo"`
	OtherPhotos *[]string    `json:"other_photos,omitempty"`
	Tags        *[]string    `json:"tags,omitempty"`
}

func (m Ad) ToResponse(user User, photos *[]string, tags *[]string) AdResponse {
	return AdResponse{
		ID:          m.ID,
		User:        user.ToResponse(),
		Name:        m.Name,
		Date:        m.Date,
		Price:       m.Price,
		Description: m.Description,
		MainPhoto:   m.MainPhoto,
		OtherPhotos: photos,
		Tags:        tags,
	}
}

func ConvertAdsToResponse(ads []Ad, usersMap map[int]User) []AdResponse {
	result := make([]AdResponse, len(ads))

	for i, ad := range ads {
		result[i] = ad.ToResponse(usersMap[ad.UserID], nil, nil)
	}
	return result
}

type AdsSortingParam struct {
	Field  string
	IsDesc bool
}

// TODO : Change this comment
// Функция возвращает массив структур AdsSortingParam, которые формируются из запроса.
// Параметры чувствительны к порядку, в котором они написаны. sort_by[i] соответствует order[i].
// Если параметр указан неверно (напрмер, "AAA"), то он будет пропущен, как и Order, соответствующий ему.
func ListAdsSortingParamsFromURL(values url.Values) []AdsSortingParam {
	var log = logrus.WithFields(logrus.Fields{
		"method": "ListAdsSortingParamsFromURL",
	})

	if values == nil {
		return nil
	}
	sortParams := values["sortby"]
	orderParams := values["order"]

	ordersLen := len(orderParams)

	var result []AdsSortingParam

	for i, sortParam := range sortParams {
		sortParam := strings.ToLower(sortParam)

		if !IsAdsSortingParamAvailable(sortParam) {
			log.WithFields(logrus.Fields{
				"sortParam": sortParam,
				"paramNum":  i,
			}).Info("sort param is not available and will be skipped")
			continue
		}

		var isDesc bool
		if i < ordersLen {
			isDesc = strings.ToLower(orderParams[i]) == "dsc"
		} else {
			log.WithFields(logrus.Fields{
				"paramNum": i,
			}).Info("order param is missed and will be set by default (dsc)")
		}
		result = append(result, AdsSortingParam{Field: sortParam, IsDesc: isDesc})
	}

	return result
}

func IsAdsSortingParamAvailable(param string) bool {
	switch strings.ToLower(param) {
	case "name":
		return true
	case "date":
		return true
	case "price":
		return true
	case "description":
		return true
	default:
		return false
	}
}

type AdOptionalFieldsParam struct {
	Description bool
	Photos      bool
	Tags        bool
}

func GetAdOptionalFieldsFromURL(values url.Values) AdOptionalFieldsParam {
	var result AdOptionalFieldsParam

	if values == nil {
		return result
	}

	fields := values["fields"]

	for _, field := range fields {
		switch strings.ToLower(field) {
		case "description":
			result.Description = true
		case "photos":
			result.Photos = true
		case "tags":
			result.Tags = true
		}
	}
	return result
}

type AdFilter struct {
	Username  string
	StartDate time.Time
	EndDate   time.Time
	Tags      []string
}

const defaultDateFormat = "2006-01-02"

func GetAdFilterFromURL(values url.Values) AdFilter {
	var result AdFilter

	if values == nil {
		return result
	}

	result.Username = values.Get("username")

	if values.Get("startdate") != "" {
		startDate, err := time.Parse(defaultDateFormat, values.Get("startdate"))
		if err != nil {
			logrus.WithError(err).Warn("failed to parse startdate param")
		}
		result.StartDate = startDate
	}

	if values.Get("enddate") != "" {
		endDate, err := time.Parse(defaultDateFormat, values.Get("enddate"))
		if err != nil {
			logrus.WithError(err).Warn("failed to parse enddate param")
		}
		result.EndDate = endDate
	}

	result.Tags = values["tags"]

	return result
}

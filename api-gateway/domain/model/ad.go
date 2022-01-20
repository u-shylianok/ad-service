package model

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetAdOptionalRequestFromURL(values url.Values) *pbAds.GetAdOptionalRequest {
	result := &pbAds.GetAdOptionalRequest{}
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

// Параметры чувствительны к порядку, в котором они написаны. sort_by[i] соответствует order[i].
// Если параметр указан неверно (напрмер, "AAA"), то он будет пропущен, как и Order, соответствующий ему.
func ListAdsSortingParamsFromURL(values url.Values) []*pbAds.SortingParam {

	if values == nil {
		return nil
	}
	sortParams := values["sortby"]
	orderParams := values["order"]

	ordersLen := len(orderParams)

	var result []*pbAds.SortingParam

	for i, sortParam := range sortParams {
		sortParam := strings.ToLower(sortParam)
		if !IsAdsSortingParamAvailable(sortParam) {
			continue
		}

		var isDesc bool
		if i < ordersLen {
			isDesc = strings.ToLower(orderParams[i]) == "dsc"
		}
		result = append(result, &pbAds.SortingParam{Field: sortParam, IsDesc: isDesc})
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

type AdFilter struct {
	Username  string
	StartDate time.Time
	EndDate   time.Time
	Tags      []string
}

const defaultDateFormat = "2006-01-02"

func GetAdFilterFromURL(values url.Values) *pbAds.AdFilter {
	result := &pbAds.AdFilter{}

	if values == nil {
		return result
	}

	result.Username = values.Get("username")

	startDateRaw := values.Get("startdate")
	if startDateRaw != "" {
		startDate, err := time.Parse(defaultDateFormat, startDateRaw)
		if err != nil {
			logrus.WithError(err).Warn("failed to parse startdate param")
		}
		result.StartDate = timestamppb.New(startDate)
	}

	endDateRaw := values.Get("enddate")
	if endDateRaw != "" {
		endDate, err := time.Parse(defaultDateFormat, endDateRaw)
		if err != nil {
			logrus.WithError(err).Warn("failed to parse enddate param")
		}
		result.EndDate = timestamppb.New(endDate)
	}

	result.Tags = values["tags"]

	return result
}

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

func (r *AdRequest) ToPb() *pbAds.AdRequest {
	return &pbAds.AdRequest{
		Name:        r.Name,
		Price:       int32(r.Price),
		Description: r.Description,
		Photo:       r.MainPhoto,
		Photos:      *r.OtherPhotos,
		Tags:        *r.Tags,
	}
}

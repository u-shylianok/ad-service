package dto

import (
	"github.com/u-shylianok/ad-service/api-gateway/domain/model"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type pbAdsConvert struct{}

var PbAds pbAdsConvert

func (c *pbAdsConvert) ToGetAdRequest(id uint32, opt *model.AdsOptional) *pbAds.GetAdRequest {
	return &pbAds.GetAdRequest{
		Id:       id,
		Optional: c.ToGetAdOptionalRequest(opt),
	}
}

func (c *pbAdsConvert) ToGetAdOptionalRequest(opt *model.AdsOptional) *pbAds.GetAdOptionalRequest {
	return &pbAds.GetAdOptionalRequest{
		Description: opt.Description,
		Photos:      opt.Photos,
		Tags:        opt.Tags,
	}
}

func (c *pbAdsConvert) FromGetAdResponse(res *pbAds.GetAdResponse) *model.AdResponse {
	return c.FromAdResponse(res.Ad)
}

func (c *pbAdsConvert) FromAdResponse(res *pbAds.AdResponse) *model.AdResponse {
	ad := res.Ad
	user := res.User
	result := &model.AdResponse{
		ID:        ad.Id,
		User:      PbAuth.FromUser(user),
		Name:      ad.Name,
		Date:      ad.Date.AsTime(),
		Price:     int(ad.Price),
		MainPhoto: ad.Photo,
	}
	if ad.Description != "" {
		result.Description = &ad.Description
	}
	if ad.Photos != nil {
		result.OtherPhotos = &ad.Photos
	}
	if ad.Tags != nil {
		result.Tags = &ad.Tags
	}
	return result
}

func (c *pbAdsConvert) ToListAdsRequest(params []model.AdsSortingParam) *pbAds.ListAdsRequest {
	return &pbAds.ListAdsRequest{
		SortingParams: c.ToSortingParams(params),
	}
}

func (c *pbAdsConvert) ToSortingParams(params []model.AdsSortingParam) []*pbAds.SortingParam {
	result := make([]*pbAds.SortingParam, len(params))
	for i, param := range params {
		result[i] = &pbAds.SortingParam{
			Field:  param.Field,
			IsDesc: param.IsDesc,
		}
	}
	return result
}

func (c *pbAdsConvert) FromListAdsResponse(res *pbAds.ListAdsResponse) []model.AdResponse {
	result := make([]model.AdResponse, len(res.Ads))
	for i, ad := range res.Ads {
		result[i] = *c.FromAdResponse(ad)
	}
	return result
}

func (c *pbAdsConvert) ToSearchAdsRequest(filter model.AdFilter) *pbAds.SearchAdsRequest {
	return &pbAds.SearchAdsRequest{
		Filter: c.ToAdFilter(filter),
	}
}

func (c *pbAdsConvert) ToAdFilter(filter model.AdFilter) *pbAds.AdFilter {
	return &pbAds.AdFilter{
		Username:  filter.Username,
		StartDate: timestamppb.New(filter.StartDate),
		EndDate:   timestamppb.New(filter.EndDate),
		Tags:      filter.Tags,
	}
}

func (c *pbAdsConvert) FromSearchAdsResponse(res *pbAds.SearchAdsResponse) []model.AdResponse {
	result := make([]model.AdResponse, len(res.Ads))
	for i, ad := range res.Ads {
		result[i] = *c.FromAdResponse(ad)
	}
	return result
}

func (c *pbAdsConvert) ToCreateAdRequest(userID uint32, ad model.AdRequest) *pbAds.CreateAdRequest {
	return &pbAds.CreateAdRequest{
		UserId: userID,
		Ad:     c.ToAdRequest(ad),
	}
}

func (c *pbAdsConvert) ToAdRequest(req model.AdRequest) *pbAds.AdRequest {
	return &pbAds.AdRequest{
		Name:        req.Name,
		Price:       int32(req.Price),
		Description: req.Description,
		Photo:       req.MainPhoto,
		Photos:      *req.OtherPhotos,
		Tags:        *req.Tags,
	}
}

func (c *pbAdsConvert) ToUpdateAdRequest(userID, adID uint32, ad model.AdRequest) *pbAds.UpdateAdRequest {
	return &pbAds.UpdateAdRequest{
		UserId: userID,
		AdId:   adID,
		Ad:     c.ToAdRequest(ad),
	}
}

func (c *pbAdsConvert) ToDeleteAdRequest(userID, adID uint32) *pbAds.DeleteAdRequest {
	return &pbAds.DeleteAdRequest{
		UserId: userID,
		AdId:   adID,
	}
}

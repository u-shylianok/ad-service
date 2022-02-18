package dto

import (
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type pbAdsConvert struct{}

var PbAds pbAdsConvert

func (c *pbAdsConvert) FromGetAdRequest(req *pbAds.GetAdRequest) (uint32, model.AdsOptional) {
	opt := req.Optional
	return req.GetId(), model.AdsOptional{
		Description: opt.Description,
		Photos:      opt.Photos,
		Tags:        opt.Tags,
	}
}

func (c *pbAdsConvert) ToGetAdResponse(ad model.Ad, user *pbAuth.GetUserResponse) *pbAds.GetAdResponse {
	return &pbAds.GetAdResponse{
		Ad: c.ToAdResponse(ad, user.User),
	}
}

func (c *pbAdsConvert) ToAdResponse(ad model.Ad, user *pbAuth.UserResponse) *pbAds.AdResponse {
	return &pbAds.AdResponse{
		Ad:   c.ToAd(ad),
		User: user,
	}
}

func (c *pbAdsConvert) ToAd(ad model.Ad) *pbAds.Ad {
	result := &pbAds.Ad{
		Id:     ad.ID,
		UserId: ad.UserID,
		Name:   ad.Name,
		Date:   timestamppb.New(ad.Date),
		Price:  int32(ad.Price),
		Photo:  ad.MainPhoto,
	}
	if ad.Description != nil {
		result.Description = *ad.Description
	}
	if ad.OtherPhotos != nil {
		result.Photos = *ad.OtherPhotos
	}
	if ad.Tags != nil {
		result.Tags = *ad.Tags
	}
	return result
}

func (c *pbAdsConvert) FromListAdsRequest(req *pbAds.ListAdsRequest) []model.AdsSortingParam {
	result := make([]model.AdsSortingParam, len(req.SortingParams))
	for i, param := range req.SortingParams {
		result[i] = c.FromSortingParam(param)
	}
	return result
}

func (c *pbAdsConvert) FromSortingParam(param *pbAds.SortingParam) model.AdsSortingParam {
	return model.AdsSortingParam{
		Field:  param.Field,
		IsDesc: param.IsDesc,
	}
}

func (c *pbAdsConvert) ToListAdsResponse(ads []model.Ad, users *pbAuth.ListUsersInIDsResponse) *pbAds.ListAdsResponse {
	usersMap := make(map[uint32]*pbAuth.UserResponse)
	for _, user := range users.Users {
		usersMap[user.Id] = user
	}

	result := make([]*pbAds.AdResponse, len(ads))
	for i, ad := range ads {
		result[i] = c.ToAdResponse(ad, usersMap[ad.UserID])
	}
	return &pbAds.ListAdsResponse{
		Ads: result,
	}
}

func (c *pbAdsConvert) FromSearchAdsRequest(req *pbAds.SearchAdsRequest) (username string, adFilter model.AdFilter) {
	return c.FromAdFilter(req.Filter)
}

func (c *pbAdsConvert) FromAdFilter(filter *pbAds.AdFilter) (username string, adFilter model.AdFilter) {
	return filter.Username,
		model.AdFilter{
			StartDate: filter.StartDate.AsTime(),
			EndDate:   filter.EndDate.AsTime(),
			Tags:      filter.Tags,
		}
}

func (c *pbAdsConvert) ToSearchAdsResponse(ads []model.Ad,
	users *pbAuth.ListUsersInIDsResponse) *pbAds.SearchAdsResponse {

	usersMap := make(map[uint32]*pbAuth.UserResponse)
	for _, user := range users.Users {
		usersMap[user.Id] = user
	}

	result := make([]*pbAds.AdResponse, len(ads))
	for i, ad := range ads {
		result[i] = c.ToAdResponse(ad, usersMap[ad.UserID])
	}
	return &pbAds.SearchAdsResponse{
		Ads: result,
	}
}

func (c *pbAdsConvert) FromCreateAdRequest(req *pbAds.CreateAdRequest) (uint32, model.AdRequest) {
	return req.UserId, c.FromAdRequest(req.Ad)
}

func (c *pbAdsConvert) FromAdRequest(req *pbAds.AdRequest) model.AdRequest {
	return model.AdRequest{
		Name:        req.Name,
		Price:       int(req.Price),
		Description: req.Description,
		MainPhoto:   req.Photo,
		OtherPhotos: &req.Photos,
		Tags:        &req.Tags,
	}
}

func (c *pbAdsConvert) FromUpdateAdRequest(req *pbAds.UpdateAdRequest) (uint32, uint32, model.AdRequest) {
	return req.UserId, req.AdId, c.FromAdRequest(req.Ad)
}

func (c *pbAdsConvert) FromDeleteAdRequest(req *pbAds.DeleteAdRequest) (uint32, uint32) {
	return req.UserId, req.AdId
}

func (c *pbAdsConvert) FromListPhotosRequest(req *pbAds.ListPhotosRequest) uint32 {
	return req.AdId
}

func (c *pbAdsConvert) ToListPhotosResponse(photos []string) *pbAds.ListPhotosResponse {
	return &pbAds.ListPhotosResponse{
		Photos: photos,
	}
}

func (c *pbAdsConvert) FromListTagsRequest(req *pbAds.ListTagsRequest) uint32 {
	return req.AdId
}

func (c *pbAdsConvert) ToListTagsResponse(tags []string) *pbAds.ListTagsResponse {
	return &pbAds.ListTagsResponse{
		Tags: tags,
	}
}

package dto

import (
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromPbAds_GetAdRequest(req *pbAds.GetAdRequest) (uint32, model.AdsOptional) {
	opt := req.Optional
	return req.GetId(), model.AdsOptional{
		Description: opt.Description,
		Photos:      opt.Photos,
		Tags:        opt.Tags,
	}
}

func ToPbAds_GetAdResponse(ad model.Ad, user *pbAuth.GetUserResponse) *pbAds.GetAdResponse {
	return &pbAds.GetAdResponse{
		Ad: ToPbAds_AdResponse(ad, user.User),
	}
}

func ToPbAds_AdResponse(ad model.Ad, user *pbAuth.UserResponse) *pbAds.AdResponse {
	return &pbAds.AdResponse{
		Ad:   ToPbAds_Ad(ad),
		User: user,
	}
}

func ToPbAds_Ad(ad model.Ad) *pbAds.Ad {
	return &pbAds.Ad{
		Id:          ad.ID,
		UserId:      ad.UserID,
		Name:        ad.Name,
		Date:        timestamppb.New(ad.Date),
		Price:       int32(ad.Price),
		Description: *ad.Description,
		Photo:       ad.MainPhoto,
		Photos:      *ad.OtherPhotos,
		Tags:        *ad.Tags,
	}
}

func FromPbAds_ListAdsRequest(req *pbAds.ListAdsRequest) []model.AdsSortingParam {
	result := make([]model.AdsSortingParam, len(req.SortingParams))
	for i, param := range req.SortingParams {
		result[i] = FromPbAds_SortingParam(param)
	}
	return result
}

func FromPbAds_SortingParam(param *pbAds.SortingParam) model.AdsSortingParam {
	return model.AdsSortingParam{
		Field:  param.Field,
		IsDesc: param.IsDesc,
	}
}

func ToPbAds_ListAdsResponse(ads []model.Ad, users map[uint32]*pbAuth.UserResponse) *pbAds.ListAdsResponse {
	result := make([]*pbAds.AdResponse, len(ads))
	for i, ad := range ads {
		result[i] = ToPbAds_AdResponse(ad, users[ad.ID])
	}
	return &pbAds.ListAdsResponse{
		Ads: result,
	}
}

func FromPbAds_SearchAdsRequest(req *pbAds.SearchAdsRequest) model.AdFilter {
	return FromPbAds_AdFilter(req.Filter)
}

func FromPbAds_AdFilter(filter *pbAds.AdFilter) model.AdFilter {
	return model.AdFilter{
		Username:  filter.Username,
		StartDate: filter.StartDate.AsTime(),
		EndDate:   filter.EndDate.AsTime(),
		Tags:      filter.Tags,
	}
}

func ToPbAds_SearchAdsResponse(ads []model.Ad, users map[uint32]*pbAuth.UserResponse) *pbAds.SearchAdsResponse {
	result := make([]*pbAds.AdResponse, len(ads))
	for i, ad := range ads {
		result[i] = ToPbAds_AdResponse(ad, users[ad.ID])
	}
	return &pbAds.SearchAdsResponse{
		Ads: result,
	}
}

func FromPbAds_CreateAdRequest(req *pbAds.CreateAdRequest) (uint32, model.AdRequest) {
	return req.UserId, FromPbAds_AdRequest(req.Ad)
}

func FromPbAds_AdRequest(req *pbAds.AdRequest) model.AdRequest {
	return model.AdRequest{
		Name:        req.Name,
		Price:       int(req.Price),
		Description: req.Description,
		MainPhoto:   req.Photo,
		OtherPhotos: &req.Photos,
		Tags:        &req.Tags,
	}
}

func FromPbAds_UpdateAdRequest(req *pbAds.UpdateAdRequest) (uint32, uint32, model.AdRequest) {
	return req.UserId, req.AdId, FromPbAds_AdRequest(req.Ad)
}

func FromPbAds_DeleteAdRequest(req *pbAds.DeleteAdRequest) (uint32, uint32) {
	return req.UserId, req.AdId
}

func FromPbAds_ListPhotosRequest(req *pbAds.ListPhotosRequest) uint32 {
	return req.AdId
}

func ToPbAds_ListPhotosResponse(photos []string) *pbAds.ListPhotosResponse {
	return &pbAds.ListPhotosResponse{
		Photos: photos,
	}
}

func FromPbAds_ListTagsRequest(req *pbAds.ListTagsRequest) uint32 {
	return req.AdId
}

func ToPbAds_ListTagsResponse(tags []string) *pbAds.ListTagsResponse {
	return &pbAds.ListTagsResponse{
		Tags: tags,
	}
}

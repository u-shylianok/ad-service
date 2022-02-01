package dto

import (
	"github.com/u-shylianok/ad-service/api-gateway/domain/model"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbAds_GetAdRequest(id uint32, opt *model.AdsOptional) *pbAds.GetAdRequest {
	return &pbAds.GetAdRequest{
		Id:       id,
		Optional: ToPbAds_GetAdOptionalRequest(opt),
	}
}

func ToPbAds_GetAdOptionalRequest(opt *model.AdsOptional) *pbAds.GetAdOptionalRequest {
	return &pbAds.GetAdOptionalRequest{
		Description: opt.Description,
		Photos:      opt.Photos,
		Tags:        opt.Tags,
	}
}

func FromPbAds_GetAdResponse(res *pbAds.GetAdResponse) *model.AdResponse {
	return FromPbAds_AdResponse(res.Ad)
}

func FromPbAds_AdResponse(res *pbAds.AdResponse) *model.AdResponse {
	ad := res.Ad
	user := res.User
	return &model.AdResponse{
		ID:          ad.Id,
		User:        FromPbAuth_User(user),
		Name:        ad.Name,
		Date:        ad.Date.AsTime(),
		Price:       int(ad.Price),
		MainPhoto:   ad.Photo,
		Description: &ad.Description,
		OtherPhotos: &ad.Photos,
		Tags:        &ad.Tags,
	}
}

func ToPbAds_ListAdsRequest(params []model.AdsSortingParam) *pbAds.ListAdsRequest {
	return &pbAds.ListAdsRequest{
		SortingParams: ToPbAds_SortingParams(params),
	}
}

func ToPbAds_SortingParams(params []model.AdsSortingParam) []*pbAds.SortingParam {
	result := make([]*pbAds.SortingParam, len(params))
	for i, param := range params {
		result[i] = &pbAds.SortingParam{
			Field:  param.Field,
			IsDesc: param.IsDesc,
		}
	}
	return result
}

func FromPbAds_ListAdsResponse(res *pbAds.ListAdsResponse) []model.AdResponse {
	result := make([]model.AdResponse, len(res.Ads))
	for i, ad := range res.Ads {
		result[i] = *FromPbAds_AdResponse(ad)
	}
	return result
}

func ToPbAds_SearchAdsRequest(filter model.AdFilter) *pbAds.SearchAdsRequest {
	return &pbAds.SearchAdsRequest{
		Filter: ToPbAds_AdFilter(filter),
	}
}

func ToPbAds_AdFilter(filter model.AdFilter) *pbAds.AdFilter {
	return &pbAds.AdFilter{
		Username:  filter.Username,
		StartDate: timestamppb.New(filter.StartDate),
		EndDate:   timestamppb.New(filter.EndDate),
		Tags:      filter.Tags,
	}
}

func FromPbAds_SearchAdsResponse(res *pbAds.SearchAdsResponse) []model.AdResponse {
	result := make([]model.AdResponse, len(res.Ads))
	for i, ad := range res.Ads {
		result[i] = *FromPbAds_AdResponse(ad)
	}
	return result
}

func ToPbAds_CreateAdRequest(userID uint32, ad model.AdRequest) *pbAds.CreateAdRequest {
	return &pbAds.CreateAdRequest{
		UserId: userID,
		Ad:     ToPbAds_AdRequest(ad),
	}
}

func ToPbAds_AdRequest(req model.AdRequest) *pbAds.AdRequest {
	return &pbAds.AdRequest{
		Name:        req.Name,
		Price:       int32(req.Price),
		Description: req.Description,
		Photo:       req.MainPhoto,
		Photos:      *req.OtherPhotos,
		Tags:        *req.Tags,
	}
}

func ToPbAds_UpdateAdRequest(userID, adID uint32, ad model.AdRequest) *pbAds.UpdateAdRequest {
	return &pbAds.UpdateAdRequest{
		UserId: userID,
		AdId:   adID,
		Ad:     ToPbAds_AdRequest(ad),
	}
}

func ToPbAds_DeleteAdRequest(userID, adID uint32) *pbAds.DeleteAdRequest {
	return &pbAds.DeleteAdRequest{
		UserId: userID,
		AdId:   adID,
	}
}

package dto

import (
	"github.com/u-shylianok/ad-service/svc-ads/client/ads"
	"github.com/u-shylianok/ad-service/svc-ads/model"
)

func FromGetAdRequest(in *ads.GetAdRequest) (uint32, model.GetAdOptional) {
	opt := in.GetOptional()
	return in.GetId(), model.GetAdOptional{
		Description: opt.GetDescription(),
		Photos:      opt.GetPhotos(),
		Tags:        opt.GetTags(),
	}
}

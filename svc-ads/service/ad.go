package service

import (
	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
)

type AdService struct {
	adRepo    repository.Ad
	photoRepo repository.Photo
	tagRepo   repository.Tag
}

func NewAdService(adRepo repository.Ad, photoRepo repository.Photo,
	tagRepo repository.Tag) *AdService {

	return &AdService{
		adRepo:    adRepo,
		photoRepo: photoRepo,
		tagRepo:   tagRepo,
	}
}

func (s *AdService) GetAd(adID uint32, fields model.AdsOptional) (model.Ad, error) {
	ad, err := s.adRepo.Get(adID, fields)
	if err != nil {
		return model.Ad{}, err
	}

	var photos *[]string
	if fields.Photos {
		photoLinks, err := s.photoRepo.ListLinksByAd(adID)
		if err == nil {
			photos = &photoLinks
		}
	}

	var tags *[]string
	if fields.Tags {
		tagNames, err := s.tagRepo.ListNamesByAd(adID)
		if err == nil {
			tags = &tagNames
		}
	}
	ad.OtherPhotos = photos
	ad.Tags = tags
	return ad, nil
}

func (s *AdService) ListAds(params []model.AdsSortingParam) ([]model.Ad, error) {
	ads, err := s.adRepo.List(params)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (s *AdService) SearchAds(filter model.AdFilter) ([]model.Ad, error) {
	ads, err := s.adRepo.ListWithFilter(filter)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (s *AdService) CreateAd(userID uint32, ad model.AdRequest) (uint32, error) {
	adID, err := s.adRepo.Create(userID, ad)
	if err != nil {
		return adID, err
	}

	if ad.OtherPhotos != nil {
		if err := s.photoRepo.CreateList(adID, *ad.OtherPhotos); err != nil {
			return adID, err
		}
	}

	if ad.Tags != nil {
		for _, tagName := range *ad.Tags {
			tagID, err := s.tagRepo.GetIDOrCreateIfNotExists(tagName)
			if err != nil {
				continue
			}

			if err := s.tagRepo.AttachToAd(adID, tagID); err != nil {
				return adID, err
			}
		}
	}

	return adID, err
}

func (s *AdService) UpdateAd(userID, adID uint32, ad model.AdRequest) (uint32, error) {
	adID, err := s.adRepo.Update(userID, adID, ad)
	if err != nil {
		return adID, err
	}

	if err := s.photoRepo.DeleteAllByAd(adID); err != nil {
		return adID, err
	}
	if ad.OtherPhotos != nil {
		if err := s.photoRepo.CreateList(adID, *ad.OtherPhotos); err != nil {
			return adID, err
		}
	}

	if err := s.tagRepo.DetachAllFromAd(adID); err != nil {
		return adID, err
	}
	if ad.Tags != nil {
		for _, tagName := range *ad.Tags {
			tagID, err := s.tagRepo.GetIDOrCreateIfNotExists(tagName)
			if err != nil {
				continue
			}

			if err := s.tagRepo.AttachToAd(adID, tagID); err != nil {
				return adID, err
			}
		}
	}

	return adID, nil
}

func (s *AdService) DeleteAd(userID, adID uint32) error {
	return s.adRepo.Delete(userID, adID)
}

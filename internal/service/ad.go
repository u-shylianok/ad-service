package service

import (
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
)

type AdService struct {
	adRepo    repository.Ad
	photoRepo repository.Photo
	tagRepo   repository.Tag
}

func NewAdService(adRepo repository.Ad, photoRepo repository.Photo, tagRepo repository.Tag) *AdService {
	return &AdService{
		adRepo:    adRepo,
		photoRepo: photoRepo,
		tagRepo:   tagRepo,
	}
}

func (s *AdService) CreateAd(ad model.AdRequest) (int, error) {
	adID, err := s.adRepo.Create(1, ad)
	if err != nil {
		// comment
		return adID, err
	}

	if ad.OtherPhotos != nil {

		if err := s.photoRepo.CreateList(adID, *ad.OtherPhotos); err != nil {
			// comment
			return adID, err
		}
	}

	if ad.Tags != nil {
		// tags logic
	}

	return adID, err
}

func (s *AdService) ListAds(sortBy, order string) ([]model.AdResponse, error) {
	return s.adRepo.List(sortBy, order)
}

func (s *AdService) GetAd(adID int, fields []string) (model.AdResponse, error) {
	return s.adRepo.Get(adID, fields)
}

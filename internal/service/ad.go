package service

import (
	"github.com/jackc/pgx/v4"
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
	adID, err := s.adRepo.Create(1, ad) // TODO : add user id
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
		for _, tagName := range *ad.Tags {
			var tagID int
			tag, err := s.tagRepo.FindByName(tagName)
			if err != nil && err != pgx.ErrNoRows {
				continue
			} else if err != nil && err == pgx.ErrNoRows {
				tagID, err = s.tagRepo.Create(tagName)
			} else if err == nil {
				tagID = tag.ID
			}

			if err := s.tagRepo.AttachTagToAd(adID, tagID); err != nil {
				return adID, err
			}
		}
	}

	return adID, err
}

func (s *AdService) ListAds(sortBy, order string) ([]model.AdResponse, error) {
	ads, err := s.adRepo.List(sortBy, order)
	if err != nil {
		return nil, err
	}

	adsResponse := model.ConvertAdsToResponse(ads)

	return adsResponse, nil
}

func (s *AdService) GetAd(adID int, fields []string) (model.AdResponse, error) {
	ad, err := s.adRepo.Get(adID, fields)
	if err != nil {
		return model.AdResponse{}, err
	}

	adResponse := ad.ToResponse(nil, nil)

	return adResponse, nil
}

func (s *AdService) UpdateAd(adID int, ad model.AdRequest) error {
	if err := s.adRepo.Update(adID, ad); err != nil {
		return err
	}

	if ad.OtherPhotos != nil {
		// NOT OPTIMIZED
		if err := s.photoRepo.DeleteAllAdPhotos(adID); err != nil {
			return err
		}

		if err := s.photoRepo.CreateList(adID, *ad.OtherPhotos); err != nil {
			return err
		}
	}

	if ad.Tags != nil {
		for _, tagName := range *ad.Tags {
			// NOT OPTIMIZED
			if err := s.tagRepo.DetachAllTagsFromAd(adID); err != nil {
				return err
			}

			var tagID int
			tag, err := s.tagRepo.FindByName(tagName)
			if err != nil && err != pgx.ErrNoRows {
				continue
			} else if err != nil && err == pgx.ErrNoRows {
				tagID, err = s.tagRepo.Create(tagName)
			} else if err == nil {
				tagID = tag.ID
			}

			if err := s.tagRepo.AttachTagToAd(adID, tagID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *AdService) DeleteAd(adID int) error {
	return s.adRepo.Delete(adID)
}

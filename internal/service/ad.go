package service

import (
	"database/sql"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
)

type AdService struct {
	adRepo    repository.Ad
	userRepo  repository.User
	photoRepo repository.Photo
	tagRepo   repository.Tag
}

func NewAdService(adRepo repository.Ad, userRepo repository.User, photoRepo repository.Photo, tagRepo repository.Tag) *AdService {
	return &AdService{
		adRepo:    adRepo,
		userRepo:  userRepo,
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
			tag, err := s.tagRepo.GetByName(tagName)
			if err != nil && err != sql.ErrNoRows {
				continue
			} else if err != nil && err == sql.ErrNoRows {
				if tagID, err = s.tagRepo.Create(tagName); err != nil {
					continue
				}
				logrus.Infof("Tag: %s created with id = %d", tagName, tagID)
			} else if err == nil {
				tagID = tag.ID
			}

			if err := s.tagRepo.AttachToAd(adID, tagID); err != nil {
				return adID, err
			}
		}
	}

	return adID, err
}

func (s *AdService) ListAds(params []model.AdsSortingParam) ([]model.AdResponse, error) {
	ads, err := s.adRepo.List(params)
	if err != nil {
		return nil, err
	}

	usersMap := make(map[int]model.User)
	var usersIDs []int
	for _, ad := range ads {
		if _, ok := usersMap[ad.UserID]; !ok {
			usersMap[ad.UserID] = model.User{}
			usersIDs = append(usersIDs, ad.UserID)
		}
	}

	users, err := s.userRepo.ListInIDs(usersIDs)
	if err != nil {
		//logrus.Error() // Просто пока пишем ошибку
		return model.ConvertAdsToResponse(ads, nil), nil // Даже если пользователи не прогрузились, важно вернуть полученные объявления (мне кажется так)
	}
	for _, user := range users {
		usersMap[user.ID] = user
	}
	adsResponse := model.ConvertAdsToResponse(ads, usersMap)

	return adsResponse, nil
}

func (s *AdService) SearchAds(filter model.AdFilter) ([]model.AdResponse, error) {
	ads, err := s.adRepo.ListWithFilter(filter)
	if err != nil {
		return nil, err
	}

	usersMap := make(map[int]model.User)
	var usersIDs []int
	for _, ad := range ads {
		if _, ok := usersMap[ad.UserID]; !ok {
			usersMap[ad.UserID] = model.User{}
			usersIDs = append(usersIDs, ad.UserID)
		}
	}

	users, err := s.userRepo.ListInIDs(usersIDs)
	if err != nil {
		//logrus.Error() // Просто пока пишем ошибку
		return model.ConvertAdsToResponse(ads, nil), nil // Даже если пользователи не прогрузились, важно вернуть полученные объявления (мне кажется так)
	}
	for _, user := range users {
		usersMap[user.ID] = user
	}
	adsResponse := model.ConvertAdsToResponse(ads, usersMap)

	return adsResponse, nil
}

func (s *AdService) GetAd(adID int, fields model.AdOptionalFieldsParam) (model.AdResponse, error) {
	ad, err := s.adRepo.Get(adID, fields)
	if err != nil {
		return model.AdResponse{}, err
	}

	var adUser model.User
	if user, err := s.userRepo.GetByID(ad.UserID); err == nil {
		adUser = user
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

	adResponse := ad.ToResponse(adUser, photos, tags)

	return adResponse, nil
}

func (s *AdService) UpdateAd(adID int, ad model.AdRequest) error {
	if err := s.adRepo.Update(adID, ad); err != nil {
		return err
	}

	if ad.OtherPhotos != nil {
		// NOT OPTIMIZED
		if err := s.photoRepo.DeleteAllByAd(adID); err != nil {
			return err
		}

		if err := s.photoRepo.CreateList(adID, *ad.OtherPhotos); err != nil {
			return err
		}
	}

	if ad.Tags != nil {
		for _, tagName := range *ad.Tags {
			// NOT OPTIMIZED
			if err := s.tagRepo.DetachAllFromAd(adID); err != nil {
				return err
			}

			var tagID int
			tag, err := s.tagRepo.GetByName(tagName)
			if err != nil && err != pgx.ErrNoRows {
				continue
			} else if err != nil && err == pgx.ErrNoRows {
				tagID, err = s.tagRepo.Create(tagName)
			} else if err == nil {
				tagID = tag.ID
			}

			if err := s.tagRepo.AttachToAd(adID, tagID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *AdService) DeleteAd(adID int) error {
	return s.adRepo.Delete(adID)
}

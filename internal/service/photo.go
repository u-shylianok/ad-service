package service

import "github.com/u-shylianok/ad-service/internal/repository"

type PhotoService struct {
	photoRepo repository.Photo
}

func NewPhotoService(photoRepo repository.Photo) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepo,
	}
}

func (s *PhotoService) ListAdPhotos(adID int) ([]string, error) {
	return s.photoRepo.ListLinksByAd(adID)
}

func (s *PhotoService) ListPhotos() ([]string, error) {
	return s.photoRepo.ListLinks()
}

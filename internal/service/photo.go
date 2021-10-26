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

func (s *PhotoService) ListPhotos(adID int) ([]string, error) {
	return s.photoRepo.ListPhotoLinks(adID)
}

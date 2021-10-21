package service

import (
	"github.com/u-shylianok/ad-service/internal/model"
	"github.com/u-shylianok/ad-service/internal/repository"
)

type AdService struct {
	repo repository.Ad
}

func NewAdService(repo repository.Ad) *AdService {
	return &AdService{repo: repo}
}

func (s *AdService) CreateAd(ad model.AdRequest) (int, error) {
	return s.repo.Create(ad)
}

func (s *AdService) ListAds(sortBy, order string) ([]model.AdResponse, error) {
	return s.repo.List(sortBy, order)
}

func (s *AdService) GetAd(adId int, fields []string) (model.AdResponse, error) {
	return s.repo.Get(adId, fields)
}

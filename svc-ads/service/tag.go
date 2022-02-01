package service

import "github.com/u-shylianok/ad-service/svc-ads/repository"

type TagService struct {
	tagRepo repository.Tag
}

func NewTagService(tagRepo repository.Tag) *TagService {
	return &TagService{
		tagRepo: tagRepo,
	}
}

func (s *TagService) ListAdTags(adID uint32) ([]string, error) {
	return s.tagRepo.ListNamesByAd(adID)
}

func (s *TagService) ListTags() ([]string, error) {
	return s.tagRepo.ListNames()
}

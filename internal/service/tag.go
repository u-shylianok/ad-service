package service

import "github.com/u-shylianok/ad-service/internal/repository"

type TagService struct {
	tagRepo repository.Tag
}

func NewTagService(tagRepo repository.Tag) *TagService {
	return &TagService{
		tagRepo: tagRepo,
	}
}

func (s *TagService) ListTags(adID int) ([]string, error) {
	return s.tagRepo.ListTagNames(adID)
}

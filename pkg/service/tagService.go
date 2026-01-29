package service

import (
	"errors"

	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
)

type TagService struct {
	repo repository.Tag
}

func NewTagService(tagRepo repository.Tag) Tag {
	return &TagService{
		repo: tagRepo,
	}
}

func (s *TagService) CreateTag(input request.CreateTagInput) (models.Tag, error) {
	var existingTags []models.Tag
	existingTags, _ = s.repo.GetAllTags()
	for _, tag := range existingTags {
		if tag.Name == input.Name {
			return models.Tag{}, errors.New("tag already exists")
		}
	}

	tag := models.Tag{
		Name: input.Name,
	}

	_, err := s.repo.CreateTag(tag)
	if err != nil {
		return models.Tag{}, err
	}

	return tag, nil
}

func (s *TagService) GetAllTags() ([]models.Tag, error) {
	return s.repo.GetAllTags()
}

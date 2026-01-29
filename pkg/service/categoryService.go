package service

import (
	"errors"

	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(categoryRepo repository.Category) Category {
	return &CategoryService{
		repo: categoryRepo,
	}
}

func (s *CategoryService) CreateCategory(input request.CreateCategoryInput) (models.Category, error) {
	var existingCategories []models.Category
	existingCategories, _ = s.repo.GetAllCategories()
	for _, cat := range existingCategories {
		if cat.Name == input.Name {
			return models.Category{}, errors.New("category already exists")
		}
	}

	category := models.Category{
		Name: input.Name,
	}

	_, err := s.repo.CreateCategory(category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

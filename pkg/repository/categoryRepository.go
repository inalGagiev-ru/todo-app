package repository

import (
	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category models.Category) (uint, error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil
}

func (r *CategoryRepository) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	return category, result.Error
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}

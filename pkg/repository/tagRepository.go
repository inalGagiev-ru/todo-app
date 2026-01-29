package repository

import (
	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) CreateTag(tag models.Tag) (uint, error) {
	result := r.db.Create(&tag)
	if result.Error != nil {
		return 0, result.Error
	}
	return tag.ID, nil
}

func (r *TagRepository) GetTagsByIDs(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag
	result := r.db.Where("id IN ?", ids).Find(&tags)
	return tags, result.Error
}

func (r *TagRepository) GetAllTags() ([]models.Tag, error) {
	var tags []models.Tag
	result := r.db.Find(&tags)
	return tags, result.Error
}

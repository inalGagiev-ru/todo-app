package repository

import (
	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"gorm.io/gorm"
)

type User interface {
	Create(user models.User) (uint, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id uint) (models.User, error)
	Update(user models.User) error
	Delete(id uint) error
}

type Task interface {
	CreateTask(task models.Task) (uint, error)
	GetTaskByID(id, userID uint) (models.Task, error)
	GetAllTasks(userID uint, filters TaskFilters) ([]models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(id, userID uint) error
}

type Category interface {
	CreateCategory(category models.Category) (uint, error)
	GetCategoryByID(id uint) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
}

type Tag interface {
	CreateTag(tag models.Tag) (uint, error)
	GetTagsByIDs(ids []uint) ([]models.Tag, error)
	GetAllTags() ([]models.Tag, error)
}

type TaskFilters struct {
	Status     string
	CategoryID *uint
	TagIDs     []uint
}

type Repository struct {
	User
	Task
	Category
	Tag
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Task:     NewTaskRepository(db),
		Category: NewCategoryRepository(db),
		Tag:      NewTagRepository(db),
	}
}

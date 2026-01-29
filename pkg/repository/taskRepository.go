package repository

import (
	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task models.Task) (uint, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil
}

func (r *TaskRepository) GetTaskByID(id, userID uint) (models.Task, error) {
	var task models.Task
	result := r.db.Preload("Category").Preload("Tags").
		Where("id = ? AND user_id = ?", id, userID).
		First(&task)
	return task, result.Error
}

func (r *TaskRepository) GetAllTasks(userID uint, filters TaskFilters) ([]models.Task, error) {
	var tasks []models.Task

	query := r.db.Preload("Category").Preload("Tags").
		Where("user_id = ?", userID)

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	if filters.CategoryID != nil {
		query = query.Where("category_id = ?", *filters.CategoryID)
	}

	if len(filters.TagIDs) > 0 {
		query = query.Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
			Where("task_tags.tag_id IN ?", filters.TagIDs)
	}

	result := query.Find(&tasks)
	return tasks, result.Error
}

func (r *TaskRepository) UpdateTask(task models.Task) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&task).Error; err != nil {
			return err
		}

		if len(task.Tags) > 0 {
			if err := tx.Model(&task).Association("Tags").Replace(task.Tags); err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *TaskRepository) DeleteTask(id, userID uint) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{})
	return result.Error
}

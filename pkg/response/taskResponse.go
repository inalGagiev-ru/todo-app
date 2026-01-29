package response

import (
	"time"

	"github.com/inalGagiev-ru/todo-app/pkg/dto"
	"github.com/inalGagiev-ru/todo-app/pkg/models"
)

type TaskResponse struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	Priority    string             `json:"priority"`
	DueDate     *time.Time         `json:"due_date,omitempty"`
	UserID      uint               `json:"user_id"`
	Category    *dto.CategoryShort `json:"category,omitempty"`
	Tags        []dto.TagShort     `json:"tags,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func ToTaskResponse(task models.Task) TaskResponse {
	var category *dto.CategoryShort
	if task.Category != nil {
		category = &dto.CategoryShort{
			ID:   task.Category.ID,
			Name: task.Category.Name,
		}
	}

	tags := make([]dto.TagShort, len(task.Tags))
	for i, tag := range task.Tags {
		tags[i] = dto.TagShort{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}

	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
		UserID:      task.UserID,
		Category:    category,
		Tags:        tags,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

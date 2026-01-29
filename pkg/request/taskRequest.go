package request

import "time"

type CreateTaskInput struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" enums:"pending,in_progress,completed"`
	Priority    string    `json:"priority" enums:"low,medium,high"`
	DueDate     time.Time `json:"due_date,omitempty"`
	CategoryID  *uint     `json:"category_id,omitempty"`
	TagIDs      []uint    `json:"tag_ids,omitempty"`
}

type UpdateTaskInput struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Status      *string    `json:"status,omitempty" enums:"pending,in_progress,completed"`
	Priority    *string    `json:"priority,omitempty" enums:"low,medium,high"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CategoryID  *uint      `json:"category_id,omitempty"`
	TagIDs      []uint     `json:"tag_ids,omitempty"`
}

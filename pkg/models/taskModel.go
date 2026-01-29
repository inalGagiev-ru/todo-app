package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Status      string         `gorm:"default:'pending'" json:"status"`  // pending, in_progress, completed
	Priority    string         `gorm:"default:'medium'" json:"priority"` // low, medium, high
	DueDate     *time.Time     `json:"due_date,omitempty"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	CategoryID  *uint          `json:"category_id,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	User     User      `gorm:"foreignKey:UserID" json:"-"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags     []Tag     `gorm:"many2many:task_tags;" json:"tags,omitempty"`
}

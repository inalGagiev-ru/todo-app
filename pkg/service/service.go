package service

import (
	"github.com/inalGagiev-ru/todo-app/pkg/dto"
	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
	"github.com/inalGagiev-ru/todo-app/pkg/response"
)

type Auth interface {
	SignUp(input request.SignUpInput) (response.AuthResponse, error)
	SignIn(input request.SignInInput) (response.AuthResponse, error)
}

type User interface {
	GetProfile(userID uint) (dto.UserProfile, error)
	UpdateProfile(userID uint, input request.UpdateUserInput) (dto.UserProfile, error)
	DeleteAccount(userID uint) error
}

type Task interface {
	CreateTask(userID uint, input request.CreateTaskInput) (response.TaskResponse, error)
	GetTaskByID(userID, taskID uint) (response.TaskResponse, error)
	GetAllTasks(userID uint, filters repository.TaskFilters) ([]response.TaskResponse, error)
	UpdateTask(userID, taskID uint, input request.UpdateTaskInput) (response.TaskResponse, error)
	DeleteTask(userID, taskID uint) error
}

type Category interface {
	CreateCategory(input request.CreateCategoryInput) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
}

type Tag interface {
	CreateTag(input request.CreateTagInput) (models.Tag, error)
	GetAllTags() ([]models.Tag, error)
}

type Service struct {
	Auth
	User
	Task
	Category
	Tag
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:     NewAuthService(repos.User),
		User:     NewUserService(repos.User),
		Task:     NewTaskService(repos.Task, repos.Tag, repos.Category),
		Category: NewCategoryService(repos.Category),
		Tag:      NewTagService(repos.Tag),
	}
}

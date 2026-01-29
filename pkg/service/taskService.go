package service

import (
	"errors"
	"time"

	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
	"github.com/inalGagiev-ru/todo-app/pkg/response"
)

type TaskService struct {
	taskRepo repository.Task
	tagRepo  repository.Tag
	catRepo  repository.Category
}

func NewTaskService(taskRepo repository.Task, tagRepo repository.Tag, catRepo repository.Category) Task {
	return &TaskService{
		taskRepo: taskRepo,
		tagRepo:  tagRepo,
		catRepo:  catRepo,
	}
}

func (s *TaskService) CreateTask(userID uint, input request.CreateTaskInput) (response.TaskResponse, error) {
	validStatuses := map[string]bool{"pending": true, "in_progress": true, "completed": true}
	if input.Status != "" && !validStatuses[input.Status] {
		return response.TaskResponse{}, errors.New("invalid status")
	}

	validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
	if input.Priority != "" && !validPriorities[input.Priority] {
		return response.TaskResponse{}, errors.New("invalid priority")
	}

	var tags []models.Tag
	if len(input.TagIDs) > 0 {
		tags, _ = s.tagRepo.GetTagsByIDs(input.TagIDs)
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
		UserID:      userID,
		Tags:        tags,
	}

	if !input.DueDate.IsZero() {
		task.DueDate = &input.DueDate
	}

	if input.CategoryID != nil {
		_, err := s.catRepo.GetCategoryByID(*input.CategoryID)
		if err != nil {
			return response.TaskResponse{}, errors.New("category not found")
		}
		task.CategoryID = input.CategoryID
	}

	taskID, err := s.taskRepo.CreateTask(task)
	if err != nil {
		return response.TaskResponse{}, err
	}

	createdTask, err := s.taskRepo.GetTaskByID(taskID, userID)
	if err != nil {
		return response.TaskResponse{}, err
	}

	return response.ToTaskResponse(createdTask), nil
}

func (s *TaskService) GetTaskByID(userID, taskID uint) (response.TaskResponse, error) {
	task, err := s.taskRepo.GetTaskByID(taskID, userID)
	if err != nil {
		return response.TaskResponse{}, errors.New("task not found")
	}

	return response.ToTaskResponse(task), nil
}

func (s *TaskService) GetAllTasks(userID uint, filters repository.TaskFilters) ([]response.TaskResponse, error) {
	tasks, err := s.taskRepo.GetAllTasks(userID, filters)
	if err != nil {
		return nil, err
	}

	responses := make([]response.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = response.ToTaskResponse(task)
	}

	return responses, nil
}

func (s *TaskService) UpdateTask(userID, taskID uint, input request.UpdateTaskInput) (response.TaskResponse, error) {
	task, err := s.taskRepo.GetTaskByID(taskID, userID)
	if err != nil {
		return response.TaskResponse{}, errors.New("task not found")
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description
	}
	if input.Status != nil {
		validStatuses := map[string]bool{"pending": true, "in_progress": true, "completed": true}
		if !validStatuses[*input.Status] {
			return response.TaskResponse{}, errors.New("invalid status")
		}
		task.Status = *input.Status
	}
	if input.Priority != nil {
		validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
		if !validPriorities[*input.Priority] {
			return response.TaskResponse{}, errors.New("invalid priority")
		}
		task.Priority = *input.Priority
	}
	if input.DueDate != nil {
		task.DueDate = input.DueDate
	}
	if input.CategoryID != nil {
		_, err := s.catRepo.GetCategoryByID(*input.CategoryID)
		if err != nil {
			return response.TaskResponse{}, errors.New("category not found")
		}
		task.CategoryID = input.CategoryID
	}
	if input.TagIDs != nil {
		tags, _ := s.tagRepo.GetTagsByIDs(input.TagIDs)
		task.Tags = tags
	}

	task.UpdatedAt = time.Now()

	err = s.taskRepo.UpdateTask(task)
	if err != nil {
		return response.TaskResponse{}, err
	}

	updatedTask, err := s.taskRepo.GetTaskByID(taskID, userID)
	if err != nil {
		return response.TaskResponse{}, err
	}

	return response.ToTaskResponse(updatedTask), nil
}

func (s *TaskService) DeleteTask(userID, taskID uint) error {
	return s.taskRepo.DeleteTask(taskID, userID)
}

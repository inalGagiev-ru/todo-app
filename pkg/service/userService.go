package service

import (
	"errors"

	"github.com/inalGagiev-ru/todo-app/pkg/dto"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
	"github.com/inalGagiev-ru/todo-app/pkg/response"
)

type UserService struct {
	repo repository.User
}

func NewUserService(userRepo repository.User) User {
	return &UserService{
		repo: userRepo,
	}
}

func (s *UserService) GetProfile(userID uint) (dto.UserProfile, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return dto.UserProfile{}, errors.New("user not found")
	}

	return response.ToUserProfile(user), nil
}

func (s *UserService) UpdateProfile(userID uint, input request.UpdateUserInput) (dto.UserProfile, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return dto.UserProfile{}, errors.New("user not found")
	}

	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil && *input.Email != user.Email {
		_, err := s.repo.GetByEmail(*input.Email)
		if err == nil {
			return dto.UserProfile{}, errors.New("email already taken")
		}
		user.Email = *input.Email
	}

	err = s.repo.Update(user)
	if err != nil {
		return dto.UserProfile{}, err
	}

	return response.ToUserProfile(user), nil
}

func (s *UserService) DeleteAccount(userID uint) error {
	return s.repo.Delete(userID)
}

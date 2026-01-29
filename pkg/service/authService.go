package service

import (
	"errors"

	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
	"github.com/inalGagiev-ru/todo-app/pkg/response"
	"github.com/inalGagiev-ru/todo-app/pkg/utils"
)

type AuthService struct {
	repo repository.User
}

func NewAuthService(userRepo repository.User) Auth {
	return &AuthService{
		repo: userRepo,
	}
}

func (s *AuthService) SignUp(input request.SignUpInput) (response.AuthResponse, error) {
	_, err := s.repo.GetByEmail(input.Email)
	if err == nil {
		return response.AuthResponse{}, errors.New("user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return response.AuthResponse{}, err
	}

	user := models.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Name:         input.Name,
	}

	userID, err := s.repo.Create(user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	token, err := utils.GenerateToken(userID)
	if err != nil {
		return response.AuthResponse{}, err
	}

	createdUser, err := s.repo.GetByID(userID)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return response.AuthResponse{
		Token: token,
		User:  response.ToUserProfile(createdUser),
	}, nil
}

func (s *AuthService) SignIn(input request.SignInInput) (response.AuthResponse, error) {
	user, err := s.repo.GetByEmail(input.Email)
	if err != nil {
		return response.AuthResponse{}, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(input.Password, user.PasswordHash) {
		return response.AuthResponse{}, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return response.AuthResponse{
		Token: token,
		User:  response.ToUserProfile(user),
	}, nil
}

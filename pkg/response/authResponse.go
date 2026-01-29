package response

import (
	"github.com/inalGagiev-ru/todo-app/pkg/dto"
	"github.com/inalGagiev-ru/todo-app/pkg/models"
)

type AuthResponse struct {
	Token string          `json:"token"`
	User  dto.UserProfile `json:"user"`
}

func ToUserProfile(user models.User) dto.UserProfile {
	return dto.UserProfile{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

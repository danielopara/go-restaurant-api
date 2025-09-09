package user

import (
	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/response"
)

type UserService interface {
	GetAllUsers() ([] response.UserResponse, error)
	RegisterUser(user *models.User) (response.RegisterResponse, error)
	GetUserByEmail(email string) (*models.User, error)
}
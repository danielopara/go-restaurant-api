package user

import (
	"errors"

	"github.com/danielopara/restaurant-api/claims"
	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/response"
)

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}


func (s *userService) GetAllUsers() ([] response.UserResponse, error){
	users, err := s.userRepo.FindAll()
	if err != nil{
		return nil, err
	}

	var res []response.UserResponse
	for _, u := range users {
		res = append(res, response.UserResponse{
			FirstName: u.FirstName,
			LastName: u.LastName,
			Email: u.Email,
			Role: u.Role,
		})
	}

	return res, nil
}

// register user
func (s *userService) RegisterUser(user *models.User) (response.RegisterResponse, error) {
	existingUser, err := s.userRepo.FindByEmail(user.Email)

	if err != nil {
		return response.RegisterResponse{}, nil
	}
	
	if existingUser != nil {
		return  response.RegisterResponse{}, errors.New("user with this email already exists",)
	}
	
	hashPassword, err := claims.HashPassword(user.Password)

	if err != nil{
		return response.RegisterResponse{}, errors.New("failed to hash password")
	}
	
	user.Password = hashPassword
	if err := s.userRepo.Create(user); err != nil{
		return response.RegisterResponse{}, err
	}

	token, err := claims.GenerateToken(*user)
	if err!=nil{
		return response.RegisterResponse{}, err
	}

	res := response.RegisterResponse{
		Success: true,
		Token: token,
		User: *user,
	}

	return res, nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	user = &models.User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Role: user.Role,
	}
	return user, nil
}
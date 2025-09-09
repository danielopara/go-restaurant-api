package auth

import (
	"errors"

	"github.com/danielopara/restaurant-api/claims"
	"github.com/danielopara/restaurant-api/internal/user"
	"github.com/danielopara/restaurant-api/response"
)

type AuthService interface{
	Login(email, password string) (response.LoginResponse, error)
}


type authServiceImpl struct {
	userRepo user.UserRepository
}

func NewAuthService(userRepo user.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (a *authServiceImpl) Login(email, password string) (response.LoginResponse, error){
	user, err := a.userRepo.FindByEmail(email)

	if err != nil{
		return response.LoginResponse{}, errors.New("user not found")
	}

	if !claims.CheckPassword(password, user.Password){
		return response.LoginResponse{}, errors.New("invalid email or password")
	}

	token, err := claims.GenerateToken(*user)
	if err!= nil{
		return response.LoginResponse{}, errors.New("unable to generate token")
	}

	res := response.LoginResponse{
		Token: token,
	}


	return res, nil
}
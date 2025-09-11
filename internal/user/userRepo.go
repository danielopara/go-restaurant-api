package user

import (
	"errors"

	"github.com/danielopara/restaurant-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	FindAll() ([] models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepo{db: db}
}


func (r *userRepo) FindById(id uint) (*models.User, error){
	var user *models.User
	err := r.db.Where("ID = ?", id).First(&user).Error

	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, errors.New("User does not exist")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*models.User, error){
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return  nil, err
	}
	return &user, err
}

func (r *userRepo) FindAll()([]models.User, error){
	var users []models.User
	err := r.db.Find(&users).Error

	if err != nil{
		return nil, err
	}

	return users, nil
}
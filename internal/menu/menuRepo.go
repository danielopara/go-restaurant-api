package menu

import (
	"github.com/danielopara/restaurant-api/models"
	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateFood(menu *models.Menu) error
	FindFood(food string)(*models.Menu, error)
	Foods()([] *models.Menu, error)
	UpdateMenuItem(id uint, updates map[string]interface{})error
	DeleteMenuItem(id uint) error
}

type menuRepo struct{
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository{
	return &menuRepo{db: db}
}

//delete todo by id
func (m *menuRepo) DeleteMenuItem(id uint) error {
	return m.db.Delete(&models.Menu{}, id).Error
}

//find all food
func (m *menuRepo) Foods() ( [] *models.Menu, error){
	var menu []*models.Menu
	err := m.db.Find(&menu).Error

	if err != nil{
		return nil, err
	}

	return menu, nil
}

//update food
func (m *menuRepo) UpdateMenuItem(id uint, updates map[string]interface{}) error{
	return m.db.Model(&models.Menu{}).Where("id=?", id).Updates(updates).Error
}


// create food
func (m *menuRepo) CreateFood(menu *models.Menu) error{
	return m.db.Create(menu).Error
}

// find food
func (m *menuRepo) FindFood(food string) (*models.Menu, error){

	var menuItem models.Menu
	err := m.db.Where("name = ?", food).First(&menuItem).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return nil, err
	}

	return &menuItem, nil
}
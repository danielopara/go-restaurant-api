package menu

import (
	"errors"

	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/response"
)

type MenuService interface {
	CreateFood(menu *models.Menu) (*models.Menu, error)
	GetFoodByName(name string) (*models.Menu, error)
	GetMenu() ([] *response.MenuResponse, error)
	UpdateMenuItem(id uint, updates map[string]interface{}) error
	DeleteMenuItem(id uint) error
}


type menuServiceImpl struct{
	menuRepo MenuRepository
}

func NewMenuService(menuRepo MenuRepository) MenuService{
	return &menuServiceImpl{menuRepo: menuRepo}
}

func (m *menuServiceImpl) DeleteMenuItem(id uint) error{
	return m.menuRepo.DeleteMenuItem(id)
}

// update
func (m *menuServiceImpl) UpdateMenuItem(id uint, updates map[string]interface{}) error{
	return m.menuRepo.UpdateMenuItem(id, updates)
}

// get get menu item
func (m *menuServiceImpl) GetMenu() ([] *response.MenuResponse, error){
	menu, err := m.menuRepo.Foods()
	
	if err != nil{
		return nil, err
	}

	var res []*response.MenuResponse
	for _, menuItem := range menu{
		res = append(res, &response.MenuResponse{
			Name: menuItem.Name,
			Price: menuItem.Price,
			Category: menuItem.Category,
			Available: menuItem.Available,
		})
	}
	
	return res, nil
}

//create food
func (m *menuServiceImpl) CreateFood(menu *models.Menu) (*models.Menu, error){
	existingMenuItem, err := m.menuRepo.FindFood(menu.Name)

	if err != nil{
		return nil, err
	}

	if existingMenuItem != nil{
		return nil, errors.New("menu Item already exist")
	}

	if err := m.menuRepo.CreateFood(menu); err != nil{
		return nil, err
	}
	
	return menu, nil
}

// get food by name
func (m *menuServiceImpl) GetFoodByName(name string) (*models.Menu, error){

	menuItem, err := m.menuRepo.FindFood(name)

	if err != nil{
		return nil, err
	}

	menu := &models.Menu{
		Name: menuItem.Name,
		Price: menuItem.Price,
		Category: menuItem.Category,
		Available: menuItem.Available,
	}
	

	return menu, nil
}
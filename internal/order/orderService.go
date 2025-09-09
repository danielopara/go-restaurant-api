package order

import (
	"errors"

	"github.com/danielopara/restaurant-api/internal/menu"
	"github.com/danielopara/restaurant-api/models"
)


type OrderItemRequest struct {
	MenuItemID uint `json:"menu_item_id"`
	Name string `json:"name"`
	Quantity   int  `json:"quantity"`
}


type OrderService interface {
	MakeOrder(waiterID uint, tableNo int, items []OrderItemRequest) ( *models.Order ,error)
}


type orderServiceImpl struct{
	orderRepo OrderRepository
	menuRepo menu.MenuRepository
}

func NewOrderService(orderRepo OrderRepository, menuRepo menu.MenuRepository) OrderService{
	return &orderServiceImpl{orderRepo: orderRepo, menuRepo: menuRepo}
}


func (o *orderServiceImpl) MakeOrder(waiterID uint, tableNo int, items []OrderItemRequest) ( *models.Order ,error){

	if len(items) == 0 {
		return nil, errors.New("item cannot be empty")
	}

	var orderItems []models.OrderItem

	for _, i := range items{
		menuItem, err := o.menuRepo.FindFood(i.Name)
		if err != nil {
			return nil, errors.New("menu item not found")
		}

		if !menuItem.Available {
			return nil, errors.New("menu item is not available")
		}

		if i.Quantity < 0{
			return nil, errors.New("quantity must be greater than 0")
		}
		
		orderItems = append(orderItems, models.OrderItem{
			MenuItemID: i.MenuItemID,
			Quantity: uint(i.Quantity),
			// MenuItem: ,
			
		})
	}

	order := &models.Order{
		TableNo: uint(tableNo),
		WaiterID: waiterID,
		Items: orderItems,
		Status: models.StatusPending,
	}

	if err := o.orderRepo.CreateOrder(order); err != nil{
		return nil, err
	}

	return order, nil
 }
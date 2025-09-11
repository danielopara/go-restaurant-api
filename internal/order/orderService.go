package order

import (
	"errors"

	"github.com/danielopara/restaurant-api/internal/menu"
	"github.com/danielopara/restaurant-api/internal/user"
	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/request"
	"github.com/danielopara/restaurant-api/response"
)




type OrderService interface {
	MakeOrder(waiterID uint, tableNo int, items []request.OrderItemRequest) ( *response.OrderResponse ,error)
}


type orderServiceImpl struct{
	orderRepo OrderRepository
	menuRepo menu.MenuRepository
	userRepo user.UserRepository
}

func NewOrderService(orderRepo OrderRepository, menuRepo menu.MenuRepository, userRepo user.UserRepository) OrderService{
	return &orderServiceImpl{orderRepo: orderRepo, menuRepo: menuRepo, userRepo: userRepo}
}


func (o *orderServiceImpl) MakeOrder(waiterID uint, tableNo int, items []request.OrderItemRequest) ( *response.OrderResponse ,error){

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
			MenuItemID: menuItem.ID,
			Quantity: uint(i.Quantity),
			MenuItem: *menuItem,
			
		})
	}

	waiter, err := o.userRepo.FindById(waiterID)
	if err != nil {
		return  nil, errors.New("waiter not found")
	}

	order := &models.Order{
		TableNo: uint(tableNo),
		WaiterID: waiter.ID,
		Items: orderItems,
		Status: models.StatusPending,
	}

	if err := o.orderRepo.CreateOrder(order); err != nil{
		return nil, errors.New("error creating order")
	}

	
	var menuItems []response.OrderMenuItemsResponse

	for _, i := range orderItems{
		menuItems = append(menuItems, response.OrderMenuItemsResponse{
			Name: i.MenuItem.Name,
			Quantity: i.Quantity,
		})
	}
	orderRes:=&response.OrderResponse{
		TableNo: order.TableNo,
		WaiterID: order.WaiterID,
		WaiterName: waiter.FirstName + waiter.LastName,
		MenuItems: menuItems,
		Status: string(order.Status),
	}

	return orderRes, nil
 }
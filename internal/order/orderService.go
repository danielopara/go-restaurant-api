package order

import (
	"errors"
	"fmt"

	"github.com/danielopara/restaurant-api/internal/menu"
	"github.com/danielopara/restaurant-api/internal/user"
	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/request"
	"github.com/danielopara/restaurant-api/response"
)




type OrderService interface {
	MakeOrder(waiterID uint, tableNo int, items []request.OrderItemRequest) ( *response.OrderResponse ,error)
	FindOrderById(id uint) (*response.OrderResponse, error)
	DeleteOrderById(id uint) error
	FindOrders() ([]*models.Order , error)
	UpdateOrderStatus(id uint, status models.OrderStatus, role models.Role, userId uint) error
}


type orderServiceImpl struct{
	orderRepo OrderRepository
	menuRepo menu.MenuRepository
	userRepo user.UserRepository
}

func NewOrderService(orderRepo OrderRepository, menuRepo menu.MenuRepository, userRepo user.UserRepository) OrderService{
	return &orderServiceImpl{orderRepo: orderRepo, menuRepo: menuRepo, userRepo: userRepo}
}

// find all orders
func( o *orderServiceImpl) FindOrders() ([]*models.Order,error){
	orders, err := o.orderRepo.FindOrders()

	if err != nil{
		return nil,err
	}

	return orders, nil
}

//updating status
func (o *orderServiceImpl) UpdateOrderStatus(id uint, status models.OrderStatus, role models.Role, userId uint)error{
	order, err := o.orderRepo.FindOrderById(id)
	
	if err != nil {
		return err
	}

	if !status.IsValid(){
		return errors.New("not a valid status")
	}

	switch role{
	case models.RoleWaiter:
		if order.WaiterID == userId{
		return errors.New("user does not own the order")
	}
		if !(order.Status == models.StatusReady && status == models.StatusServed){
			return errors.New("waiter can only updated orders to served")
		}
	case models.RoleChef:
		if order.ChefID == nil{
			order.ChefID = &userId
		} else if *order.ChefID != userId{
			return errors.New("another chef is assigned to the order")
		}
		if !((order.Status == models.StatusPending && status == models.StatusInProgress) || 
		(order.Status == models.StatusInProgress && status == models.StatusReady)) {
			return errors.New("chef can only move order from Pending -> InProgress or InProgress -> Ready")
		}
	case models.RoleCashier:
		if status != models.StatusClosed{
			return errors.New("cashier can only close orders")
		}
	case models.RoleManager:
		return o.orderRepo.UpdateOrderStatus(id, status)
	case models.RoleAdmin:
		return o.orderRepo.UpdateOrderStatus(id, status)
	default:
		return errors.New("unauthorized role")
	}

	return o.orderRepo.UpdateOrderStatus(order.ID, status)
}


// delete order by id
func (o *orderServiceImpl) DeleteOrderById(id uint)error{
	return o.orderRepo.DeleteOrderById(id)
}

// find order by id
func (o *orderServiceImpl) FindOrderById(id uint)(*response.OrderResponse, error){

	order, err := o.orderRepo.FindOrderById(id)

	if err != nil{
		return nil, err
	}

	waiterName, err :=o.userRepo.FindById(order.WaiterID)
	if err != nil{
		return nil, err
	}

	var menuItems []response.OrderMenuItemsResponse
	for _, i := range order.Items{
		menuItems = append(menuItems, response.OrderMenuItemsResponse{
			Name: i.MenuItem.Name,
			Quantity: i.Quantity,
		})
	}

	orderRes := response.OrderResponse{
		TableNo: uint(order.TableNo),
		WaiterID: order.WaiterID,
		WaiterName: fmt.Sprintf("%s %s", waiterName.FirstName, waiterName.LastName),
		Status: string(order.Status) ,
		MenuItems: menuItems,
	}

	return &orderRes, err
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
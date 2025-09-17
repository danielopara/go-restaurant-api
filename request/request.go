package request

import "github.com/danielopara/restaurant-api/models"

type Register struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role" binding:"required,oneof=waiter chef cashier manager admin"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// menu

type CreateMenuItem struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Available   bool    `json:"available"`
}

type OrderItemRequest struct {
	// MenuItemID uint   `json:"menu_item_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CreateOrder struct {
	// MenuItemId uint `json:"menu_item_id" binding:"required"`
	TableNo  uint               `json:"table_no" binding:"required"`
	WaiterID uint               `json:"waiter_id" binding:"required"`
	Items    []OrderItemRequest `json:"items" binding:"required"`
}

type UpdateOrderStatus struct {
	UserId uint `json:"user_id" binding:"required"`
	Status  models.OrderStatus `json:"status" binding:"required"`
}
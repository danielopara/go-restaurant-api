package response

import "github.com/danielopara/restaurant-api/models"

type RegisterResponse struct {
	Success bool `json:"success"`
	Token   string `json:"token"`
	User    models.User `json:"user"`
}


type UserResponse struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Role models.Role `json:"role"`
}


type LoginResponse struct{
	Token string `json:"token"`
	// FirstName string `json:"first_name"`
	// LastName string `json:"last_name"`
	// Email string `json:"email"`
}

type OrderMenuItemsResponse struct{
	Name string `json:"name" binding:"required"`
	Quantity uint	`json:"quantity" binding:"required"`
}

type OrderResponse struct{
	TableNo uint
	WaiterID uint
	WaiterName string
	MenuItems []OrderMenuItemsResponse
	Status string
}

type MenuResponse struct{
	Name string `json:"name"`
	Category models.Category `json:"category"`
	Price float64 `json:"price"`
	Available bool `json:"available"`
}
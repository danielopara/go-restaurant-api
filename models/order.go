package models

import "gorm.io/gorm"

type OrderStatus string

const (
	StatusPending OrderStatus = "Pending"
	StatusInProgress OrderStatus = "In Progress"
	StatusReady OrderStatus = "Ready"
	StatusServed OrderStatus = "Served"
	StatusClosed OrderStatus = "Closed"
)

type OrderItem struct {
	gorm.Model
	OrderId uint	`json:"order_id"`
	MenuItemID uint `json:"menu_item_id"`
	MenuItem Menu	`json:"menu_item" gorm:"foreignKey:MenuItemID"`
	Quantity uint	`json:"quantity"`
}

type Order struct {
	gorm.Model
	TableNo	uint	`json:"table_no"`
	WaiterID uint	`json:"waiter_id"`
	ChefID *uint	`json:"chef_id"`
	Status OrderStatus	`json:"status"`
	Items []OrderItem	`json:"items"`
}
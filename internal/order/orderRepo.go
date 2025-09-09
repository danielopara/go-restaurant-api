package order

import (
	"github.com/danielopara/restaurant-api/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(*models.Order) error
}

type orderRepo struct{
	db *gorm.DB
}


func NewOrderRepository ( db *gorm.DB) OrderRepository{
	return &orderRepo{db:db}
}

func (o *orderRepo) CreateOrder(order *models.Order) error {
	return o.db.Create(order).Error
}
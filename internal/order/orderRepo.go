package order

import (
	"errors"
	"fmt"

	"github.com/danielopara/restaurant-api/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(*models.Order) error
	FindOrderById(id uint) ( *models.Order, error)
	DeleteOrderById(id uint)error
	UpdateOrderStatus(id uint, status models.OrderStatus) error
	FindOrders()([]*models.Order, error)
}

type orderRepo struct{
	db *gorm.DB
}


func NewOrderRepository ( db *gorm.DB) OrderRepository{
	return &orderRepo{db:db}
}

func (o *orderRepo) FindOrders()([]*models.Order, error){
	var orders []*models.Order

	err := o.db.Preload("Items.MenuItem").Find(&orders).Error

	if err != nil{
		return nil, errors.New("could not fetch orders")
	}

	return orders, nil
}

//update status
func (o *orderRepo) UpdateOrderStatus(id uint, status models.OrderStatus) error{
	return o.db.Model(&models.Order{}).Where("id=?", id).Update("status", status).Error
}

func (o *orderRepo) CreateOrder(order *models.Order) error {
	return o.db.Create(order).Error
}

func (o *orderRepo) DeleteOrderById(id uint) error {
	return o.db.Delete(&models.Order{}, id).Error
}

func (o *orderRepo) FindOrderById(id uint) (*models.Order, error){
	var order *models.Order

	if err := o.db.Preload("Items.MenuItem").First(&order, id).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	return order, nil
}
package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/danielopara/restaurant-api/internal/order"
	"github.com/danielopara/restaurant-api/models"
)

type OrderCache struct {
	order.OrderRepository
	cache *Cache
}

func NewOrderCache(innerRepo order.OrderRepository, cache *Cache) *OrderCache{
	return &OrderCache{
		OrderRepository: innerRepo,
		cache: cache,
	}
}

func (r *OrderCache) FindOrders() ([]*models.Order, error){
	if cached, ok := r.cache.Get("all_orders"); ok{
		if cacheStr, ok := cached.(string); ok {
			var orders []*models.Order
			if err := json.Unmarshal([]byte(cacheStr), &orders); err == nil{
				log.Println("fetched orders from cache")
				return orders, nil
			}
		}
		r.cache.Delete("all_orders")
	}

	orders, err := r.OrderRepository.FindOrders()
	if err != nil{
		return nil, err
	}

	if data, err := json.Marshal(orders); err == nil{
		r.cache.Set("all_orders", string(data), 5*time.Minute)
		log.Println("cached orders from db")
	}

	return orders, err
}

func (r *OrderCache) FindOrderById(id uint) (*models.Order, error){
	key := "order_" + string(rune(id))

	if cached, ok := r.cache.Get(key); ok{
		if cacheStr, ok := cached.(string); ok{
			var order models.Order
			if err := json.Unmarshal([]byte(cacheStr), &order); err == nil{
				log.Println("fetched order from cache")
				return &order, nil
			}
			r.cache.Delete(key)
		}
	}

	order, err := r.OrderRepository.FindOrderById(id)
	if err != nil{
		return nil, err
	}

	if data, err := json.Marshal(order); err == nil{
		log.Println("fetched order from db")
		r.cache.Set(key, string(data), 5*time.Minute)
	}

	return order, nil
}
package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/danielopara/restaurant-api/internal/menu"
	"github.com/danielopara/restaurant-api/models"
)

type CacheRepo struct {
	innerRepo menu.MenuRepository
	cache     *Cache
}

func NewCacheRepo(innerRepo menu.MenuRepository, cache *Cache) *CacheRepo{
	return &CacheRepo{
		innerRepo: innerRepo,
		cache: cache,
	}
}


func (r *CacheRepo) Foods() ([]*models.Menu, error) {
    if cached, ok := r.cache.Get("all_menu_items"); ok {
        if cacheStr, ok := cached.(string); ok {
            var menu []*models.Menu
            if err := json.Unmarshal([]byte(cacheStr), &menu); err == nil {
				log.Println("fetched from cache")
                return menu, nil
            }
			log.Println(" cache corrupted, deleting and fetching from DB")
            r.cache.Delete("all_menu_items")
        }
    }
    
	log.Println("fetching from db")
    menu, err := r.innerRepo.Foods()
    if err != nil {
		log.Println("error fetching from db:", err)
        return nil, err
    }
    
    if data, err := json.Marshal(menu); err == nil {
        r.cache.Set("all_menu_items", string(data), 5*time.Minute)
		log.Println("cached db results")
    }
    
    return menu, nil
}

func (r *CacheRepo) CreateFood(m *models.Menu) ( error) {
	err := r.innerRepo.CreateFood(m)
	if err == nil {
		r.cache.Delete("all_menu_items") 
	}
	return err
}

func (r *CacheRepo) DeleteMenuItem(id uint) error{
	err := r.innerRepo.DeleteMenuItem(id)

	if err == nil{
		r.cache.Delete("all_menu_items")
	}

	return  err
}

func (r *CacheRepo) FindFood(food string) (*models.Menu, error){
	key := "menu_item_" + food

	if cached, ok := r.cache.Get(key); ok{
		if cachestr, ok := cached.(string); ok{
			var menuItem models.Menu
			if err := json.Unmarshal([]byte(cachestr), &menuItem); err == nil{
				log.Println("fetched from cache")
				return &menuItem, nil
			}
			r.cache.Delete(key)
		}
	}
	 menuItem, err := r.innerRepo.FindFood(food)
    if err != nil {
        return nil, err
    }
    
    if data, err := json.Marshal(menuItem); err == nil {
        r.cache.Set(key, string(data), 5*time.Minute)
		log.Println("cached db results")
    }
    
    return menuItem, nil
}

func (r *CacheRepo) UpdateMenuItem(id uint, updates map[string]interface{}) error{
	err := r.innerRepo.UpdateMenuItem(id, updates)

	if err == nil{
		r.cache.Delete("all_menu_items")
	}
	return err
}
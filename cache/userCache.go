package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/danielopara/restaurant-api/internal/user"
	"github.com/danielopara/restaurant-api/models"
)

type UserCache struct {
	user.UserRepository
	cache *Cache
}

func NewUserCache(innerRepo user.UserRepository, cache *Cache) *UserCache{
	return &UserCache{
		UserRepository: innerRepo,
		cache: cache,
	}
}

func (r *UserCache) FindByEmail(email string) (*models.User, error){
	key := "user_" + email

	if cached, ok := r.cache.Get(key); ok{
		if cacheStr, ok := cached.(string); ok{
			var user models.User
			if err := json.Unmarshal([]byte(cacheStr), &user); err == nil{
				log.Println("fetched user from cache")
				return &user, nil
			}

			r.cache.Delete(key)
		}
	}

	user, err := r.UserRepository.FindByEmail(email)
	if err != nil{
		return nil, err
	}

	if data, err := json.Marshal(user ); err == nil{
		log.Println("failed to fetch from db")
		r.cache.Set(key, string(data), 5 *time.Minute)
	}

	return user, nil
}

func (r *UserCache) FindAll()([]*models.User, error){
	if cached, ok := r.cache.Get("all_users"); ok{
		if cacheStr, ok := cached.(string); ok{
			var users []*models.User
			if err := json.Unmarshal([]byte(cacheStr), &users); err == nil{
				log.Println("fetched users from cache")
				return users, nil
			}
		}	
		log.Println("failed to fetch users from cache, fetching from db")
		r.cache.Delete("all_users")
	}

	users, err := r.UserRepository.FindAll()
	if err != nil{
		return nil, err
	}

	if data, err := json.Marshal(users); err == nil{
		r.cache.Set("all_users", string(data), 5*time.Minute)
		log.Println("cached users from db")
	}

	return users, nil
}
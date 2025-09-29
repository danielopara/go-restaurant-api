package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	Items map[string]CacheItem
	mu    sync.RWMutex
}

func NewCache() *Cache{
	return &Cache{
		Items: make(map[string]CacheItem),
	}
}


func (c *Cache) Set(key string, value interface{}, duration time.Duration){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Items[key] = CacheItem{
		Value: value,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
} 

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.Items[key]
	
	if !found || (item.Expiration > 0 && time.Now().UnixNano() > item.Expiration){
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Items, key)
}
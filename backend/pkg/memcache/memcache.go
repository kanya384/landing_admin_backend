//simple memory cache
package memcache

import (
	"errors"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string) error
}

type cache struct {
	mu    sync.RWMutex
	items map[string]Item
}

type Item struct {
	Value   interface{}
	Created time.Time
}

func New() Cache {
	return &cache{
		items: make(map[string]Item),
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.items[key] = Item{
		Value:   value,
		Created: time.Now(),
	}
	c.mu.Unlock()
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.items[key]
	return item, ok
}

func (c *cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.items[key]; !ok {
		return errors.New("key not found")
	}

	delete(c.items, key)
	return nil
}

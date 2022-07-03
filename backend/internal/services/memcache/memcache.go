//simple memory cache
package memcache

import (
	"errors"
	"landing_admin_backend/internal/config"
	"net/http"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (Item, bool)
	Delete(key string) error
	SetAppContent(content map[string]interface{})
	GetAppContent() (content map[string]interface{}, err error)
}

type cache struct {
	mu      sync.RWMutex
	items   map[string]Item
	content map[string]interface{}
	cfg     *config.Config
}

type Item struct {
	Value   interface{}
	Created time.Time
}

func New(cfg *config.Config) Cache {
	return &cache{
		items: make(map[string]Item),
		cfg:   cfg,
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.items[key] = Item{
		Value:   value,
		Created: time.Now(),
	}
	c.content = nil
	c.mu.Unlock()
}

func (c *cache) Get(key string) (Item, bool) {
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
	c.content = nil
	go func() {
		//updates cache on landing
		time.Sleep(time.Second * 2)
		http.Get("http://" + c.cfg.AppHost + "/update")
	}()
	return nil
}

func (c *cache) SetAppContent(content map[string]interface{}) {
	c.mu.Lock()
	c.content = content
	c.mu.Unlock()
}

func (c *cache) GetAppContent() (content map[string]interface{}, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.content) == 0 {
		return nil, errors.New("no cache")
	}
	return c.content, nil
}

package cacheinmemory

import (
	"banner-server/pkg/cache/cacheInMemory/model"
	"os"
	"sync"
	"time"
)

type CacheInMemory struct {
	sync.RWMutex
	cleanupInterval   time.Duration
	defaultExpiration time.Duration
	items             map[model.KeyCache]Item
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

func New() *CacheInMemory {
	items := make(map[model.KeyCache]Item)
	expiration, _ := time.ParseDuration(os.Getenv("CACHE_EXPIRATION"))
	cleanup, _ := time.ParseDuration(os.Getenv("CACHE_CLEANUP_INTERVAL"))
	cache := CacheInMemory{
		items:             items,
		cleanupInterval:   cleanup,
		defaultExpiration: expiration,
	}

	if cleanup > 0 {
		go cache.GC()
	}
	return &cache
}

func (c *CacheInMemory) Set(key model.KeyCache, value interface{}) {
	var expiration int64
	if c.defaultExpiration > 0 {
		expiration = time.Now().Add(c.defaultExpiration).UnixNano()
	}
	c.Lock()
	c.items[key] = Item{
		Value:      value,
		Created:    time.Now(),
		Expiration: expiration,
	}
	c.Unlock()
}

func (c *CacheInMemory) Get(key model.KeyCache) (interface{}, bool) {
	c.RLock()
	item, found := c.items[key]
	c.RUnlock()
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Value, true
}

func (c *CacheInMemory) GC() {
	for {
		<-time.After(c.cleanupInterval)
		if c.items == nil {
			return
		}
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

func (c *CacheInMemory) expiredKeys() []model.KeyCache {
	c.RLock()
	keys := make([]model.KeyCache, 0)
	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}
	c.RUnlock()
	return keys
}

func (c *CacheInMemory) clearItems(keys []model.KeyCache) {
	c.Lock()
	for _, k := range keys {
		delete(c.items, k)
	}
	c.Unlock()
}

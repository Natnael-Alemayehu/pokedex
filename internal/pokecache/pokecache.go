package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return data.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) > interval {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}

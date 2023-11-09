package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

type cacheEntry struct {
	lastRetrieved time.Time
	val           []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		lastRetrieved: time.Now(),
		val:           val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key *string) (val []byte, got bool) {
	if key == nil {
		return []byte{}, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.entries[*key]
	if !ok {
		return []byte{}, false
	}
	c.entries[*key] = cacheEntry{
		lastRetrieved: time.Now(),
		val:           v.val,
	}
	return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for {
		<-ticker.C

		for k, v := range c.entries {
			c.mu.Lock()
			if time.Since(v.lastRetrieved) > interval {
				delete(c.entries, k)
			}
			c.mu.Unlock()
		}
	}
}

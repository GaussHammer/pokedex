package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheMap[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.cacheMap[key]
	if !found {
		fmt.Println("Cache: Key not found:", key) // Debug
		return nil, false
	}

	// Check if entry has expired
	if time.Since(entry.createdAt) > c.interval {
		fmt.Println("Cache: Entry expired:", key) // Debug
		delete(c.cacheMap, key)
		return nil, false
	}

	fmt.Println("Cache: Entry found:", key) // Debug
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		now := time.Now()
		c.mu.Lock()
		for key := range c.cacheMap {
			if now.Sub(c.cacheMap[key].createdAt) > interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}

}

package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	CacheList map[string]CacheEntry
	Mu        *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {

	cacheList := make(map[string]CacheEntry)
	mu := sync.RWMutex{}

	cache := &Cache{
		CacheList: cacheList,
		Mu:        &mu,
	}
	go func() {
		cache.reapLoop(interval)
	}()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.CacheList[key] = CacheEntry{CreatedAt: time.Now(), Val: val}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	if a, exists := c.CacheList[key]; exists {
		return a.Val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.Mu.Lock()
		for key, entry := range c.CacheList {
			if t.Sub(entry.CreatedAt) > interval {
				delete(c.CacheList, key)
			}
		}
		c.Mu.Unlock()
	}
}

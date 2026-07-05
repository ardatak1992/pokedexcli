package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		entries: map[string]cacheEntry{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {

	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		timeAgo := time.Now().UTC().Add(-interval)
		for k, v := range c.entries {
			if v.createdAt.Before(timeAgo) {
				delete(c.entries, k)
			}
		}
	}

}

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

func NewCache(mut *sync.RWMutex) Cache {

	return Cache{
		entries: map[string]cacheEntry{},
		mu:      mut,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c* Cache) reapLoop(){
	
}

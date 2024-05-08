package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	Cache map[string]CacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cache: map[string]CacheEntry{},
	}
	go c.Purgeloop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Cache[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exits := c.Cache[key]
	if exits {
		return entry.val, true
	} else {
		return []byte{}, false
	}
}
func (c *Cache) Purge(interval time.Duration) {
	timeBeforeValidCache := time.Now().UTC().Add(-interval)
	for k, v := range c.Cache {
		if v.createdAt.Before(timeBeforeValidCache) {
			delete(c.Cache, k)
		}
	}
}
func (c *Cache) Purgeloop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Purge(interval)
	}
}

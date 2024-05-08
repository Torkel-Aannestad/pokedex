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
	mu    sync.Mutex
	items map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		items: map[string]cacheEntry{},
	}
	go c.Purgeloop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cEntry, exists := c.items[key]
	if !exists {
		return []byte{}, false
	}
	return cEntry.val, true
}

func (c *Cache) Purgeloop(internal time.Duration) {
	ticker := time.NewTicker(internal)
	for range ticker.C {
		c.Purgeloop(internal)
	}

}
func (c *Cache) Purge(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.items {
		if v.createdAt.Before(timeAgo) {
			delete(c.items, k)
		}
	}
}

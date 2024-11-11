package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	Mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	v, k := c.Entries[key]
	return v.val, k
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for k, v := range c.Entries {
		if time.Now().Second()-v.createdAt.Second() > int(interval.Seconds()) {
			delete(c.Entries, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Entries: make(map[string]cacheEntry),
		Mu:      &sync.Mutex{},
	}

	c.reapLoop(5)

	return c

}

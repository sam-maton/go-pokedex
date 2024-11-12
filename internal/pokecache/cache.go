package pokecache

import (
	"fmt"
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
		createdAt: time.Now().UTC(),
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
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for k, v := range c.Entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.Entries, k)
		}
	}

	fmt.Print(c.Entries)
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Entries: make(map[string]cacheEntry),
		Mu:      &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c

}

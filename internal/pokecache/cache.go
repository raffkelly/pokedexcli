package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache        map[string]cacheEntry
	mux          *sync.Mutex
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	if ok {
		return val.val, ok
	}
	return nil, ok
}

func (c *Cache) reapLoop() {
	clock := time.NewTicker(c.reapInterval)
	for {
		<-clock.C
		c.mux.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) >= c.reapInterval {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache:        make(map[string]cacheEntry),
		mux:          &sync.Mutex{},
		reapInterval: interval,
	}
	go c.reapLoop()
	return c
}

package pokecache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type storage map[string]CacheEntry

type Cache struct {
	mu       sync.Mutex
	storage  storage
	Interval time.Duration
	Ticker   *time.Ticker
}

func NewCache(interval time.Duration) *Cache {
	storage := make(storage)

	cache := &Cache{
		mu:       sync.Mutex{},
		storage:  storage,
		Interval: interval,
	}
	go cache.ReapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.storage[key]; ok {
		return errors.New("already exists")
	}
	c.storage[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return nil
}

func (c *Cache) Read(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	storage := c.storage
	if _, ok := storage[key]; ok {
		return storage[key].val, true
	}
	return nil, false
}

func (c *Cache) ReapLoop() {
	storage := c.storage
	ticker := time.NewTicker(c.Interval)
	for {
		<-ticker.C
		c.mu.Lock()
		fmt.Println("im being called")
		for key, entry := range storage {
			if time.Since(entry.createdAt) > c.Interval {
				delete(storage, key)
			}
		}
		c.mu.Unlock()
	}
}

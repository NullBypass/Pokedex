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
	store    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(duration time.Duration) *Cache {
	cache := &Cache{
		store:    make(map[string]cacheEntry),
		interval: duration,
	}
	cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := cache.store[key]
	return cacheEntry.val, ok
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	go func() {
		for range ticker.C {
			cache.mu.Lock()
			for key, entry := range cache.store {
				if time.Since(entry.createdAt) > cache.interval {
					delete(cache.store, key)
				}
			}
			cache.mu.Unlock()
		}
	}()

}

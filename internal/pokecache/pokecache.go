package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type CacheMethods interface {
	Get(string) []byte
	Add(string, []byte)
	reapLoop()
}

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.RWMutex
}

func (this *Cache) Get(key string) ([]byte, bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	x, ok := this.data[key]
	return x.Val, ok
}

func (this *Cache) Add(key string, val []byte) {
	this.data[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (this *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		this.mu.Lock()
		ttl := time.Now().Add(-interval)

		for key, entry := range this.data {
			if ttl.After(entry.CreatedAt) {
				delete(this.data, key)
			}
		}
		this.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {

	cache := &Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache

}
